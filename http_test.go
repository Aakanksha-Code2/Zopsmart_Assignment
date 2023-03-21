package patient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aakanksha/ppms/internal/models"
	"github.com/aakanksha/ppms/internal/service"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"
	"time"
)

var current_time = time.Now()
var patient = models.Patient{
	Id:          5,
	Name:        "ZopSmart",
	Phone:       "+919172681679",
	Discharge:   true,
	CreatedAt:   current_time,
	UpdatedAt:   current_time,
	BloodGroup:  "+A",
	Description: "patient description",
}

func Test_Insert(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockPatientService := service.NewMockServiceInterface(mockCtrl)

	testCases := []struct {
		body          []byte
		input         models.Patient
		mockCall      *gomock.Call
		expectedError error
		status        int
	}{
		//Success
		{
			body: []byte(`{
		
				"name": "Zopsmart",
				"phone": "+919172681679",
				"discharge": true,
				"bloodGroup": "+A",
				"description": "patient description"
				}`),
			input:         patient,
			mockCall:      mockPatientService.EXPECT().Insert(gomock.Any()).Return(&patient, nil),
			expectedError: nil,
			status:        200,
		},
		//Failure
		{
			body: []byte(`{
				"name": "Zopsmart",
				"phone": "+919172681679",
				"discharged": true,
				"bloodGroup": "+A",
				"description": "patient description"
				}`),
			input:         patient,
			mockCall:      mockPatientService.EXPECT().Insert(gomock.Any()).Return(nil, errors.New("error")),
			expectedError: errors.New("error"),
			status:        400,
		},
		{
			body: []byte(`{{
                
				"name": "ak",
				"phone": "+919172681679",
				"discharged": true,
				"bloodGroup": "+A",
				"description": "patient description"
				}`),
			input:         patient,
			expectedError: errors.New("id does not exists"),
			status:        400,
		},
	}
	p := New(mockPatientService)
	for _, testCase := range testCases {
		r := httptest.NewRequest("POST", "/patients", bytes.NewBuffer(testCase.body))
		w := httptest.NewRecorder()
		p.Insert(w, r)
		if !reflect.DeepEqual(testCase.status, w.Result().StatusCode) {
			t.Errorf("Expected error: %v Got %v", testCase.status, w.Result().StatusCode)
		}
	}
}
func Test_GetAll(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockPatientService := service.NewMockServiceInterface(mockCtrl)
	testCases := []struct {
		//body          []byte
		mockCall      *gomock.Call
		expectedError error
		status        int
	}{
		// Success
		{
			mockCall:      mockPatientService.EXPECT().GetAll().Return([]*models.Patient{&patient}, nil),
			expectedError: nil,
			status:        200,
		},
		//Failure
		{
			mockCall:      mockPatientService.EXPECT().GetAll().Return(nil, errors.New("error")),
			expectedError: errors.New("error"),
			status:        400,
		},
	}
	p := New(mockPatientService)
	for _, testCase := range testCases {
		l, _ := json.Marshal(patient)
		m := bytes.NewBuffer(l)
		r := httptest.NewRequest("GET", "/patients", m)
		w := httptest.NewRecorder()
		p.GetAll(w, r)
		if !reflect.DeepEqual(testCase.status, w.Result().StatusCode) {
			t.Errorf("Expected error: %v Got %v", testCase.status, w.Result().StatusCode)
		}
	}
}

func TestGetByID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	idString := strconv.Itoa(patient.Id)
	mockPatientService := service.NewMockServiceInterface(mockCtrl)
	testCases := []struct {
		id string
		//body          []byte
		mockCall      *gomock.Call
		expectedError error
		status        int
	}{
		//Success
		{
			id:            idString,
			mockCall:      mockPatientService.EXPECT().GetByID(patient.Id).Return(&patient, nil),
			expectedError: nil,
			status:        200,
		},
		// Failure
		{
			id:            idString,
			mockCall:      mockPatientService.EXPECT().GetByID(patient.Id).Return(&models.Patient{}, errors.New("error")),
			expectedError: errors.New("error"),
			status:        400,
		},
	}
	p := New(mockPatientService)
	for _, testCase := range testCases {
		l, _ := json.Marshal(patient)
		m := bytes.NewBuffer(l)
		r := httptest.NewRequest("GET", fmt.Sprintf("/patients/%s", testCase.id), m)
		r = mux.SetURLVars(r, map[string]string{
			"id": testCase.id,
		})
		w := httptest.NewRecorder()
		p.GetByID(w, r)
		if !reflect.DeepEqual(testCase.status, w.Result().StatusCode) {
			t.Errorf("Expected error: %v Got %v", testCase.status, w.Result().StatusCode)
		}
	}
}

func Test_DeletePatientService(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	idString := strconv.Itoa(patient.Id)
	mockPatientService := service.NewMockServiceInterface(mockCtrl)
	testCases := []struct {
		body          []byte
		id            string
		mockCall      *gomock.Call
		expectedError error
		status        int
	}{
		// Success
		{
			id:            idString,
			mockCall:      mockPatientService.EXPECT().Delete(patient.Id).Return(nil),
			expectedError: nil,
			status:        200,
		},
		//Failure
		{
			id:            idString,
			mockCall:      mockPatientService.EXPECT().Delete(patient.Id).Return(errors.New("error")),
			expectedError: errors.New("error"),
			status:        400,
		},
	}
	p := New(mockPatientService)
	for _, testCase := range testCases {
		l, _ := json.Marshal(patient)
		m := bytes.NewBuffer(l)
		r := httptest.NewRequest("DELETE", fmt.Sprintf("/patients/%s", testCase.id), m)
		r = mux.SetURLVars(r, map[string]string{
			"id": testCase.id,
		})
		w := httptest.NewRecorder()
		p.Delete(w, r)
		if !reflect.DeepEqual(testCase.status, w.Result().StatusCode) {
			t.Errorf("Expected error: %v Got %v", testCase.status, w.Result().StatusCode)
		}
	}
}

func Test_UpdatePatientService(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	idString := strconv.Itoa(patient.Id)
	mockPatientService := service.NewMockServiceInterface(mockCtrl)
	testCases := []struct {
		body          []byte
		id            string
		mockCall      *gomock.Call
		expectedError error
		status        int
	}{
		// Success
		{
			body: []byte(`{
		"id": 55,
		   "name": "ak",
		   "phone": "123566780011133",
		   "discharge": false,
		   "bloodGroup": "A+1133",
		   "description": "goo1133d"
					}`),
			id:            idString,
			mockCall:      mockPatientService.EXPECT().Update(gomock.Any(), patient.Id).Return(&patient, nil),
			expectedError: nil,
			status:        200,
		},

		{
			body: []byte(`{
				}`),
			id:            idString,
			mockCall:      mockPatientService.EXPECT().Update(gomock.Any(), patient.Id).Return(&models.Patient{}, errors.New("error")),
			expectedError: errors.New("error"),
			status:        400,
		},
		//failure
		{
			body: []byte(`{{
		"id": 55,
		   "name": "ak",
		   "phone": "123566780011133",
		   "discharge": false,
		   "bloodGroup": "A+1133",
		   "description": "goo1133d"
					}`),
			id:            idString,
			expectedError: errors.New("error"),
			status:        400,
		},
	}
	p := New(mockPatientService)
	for _, testCase := range testCases {
		r := httptest.NewRequest("PUT", fmt.Sprintf("/patients/%s", testCase.id), bytes.NewBuffer(testCase.body))
		r = mux.SetURLVars(r, map[string]string{
			"id": testCase.id,
		})
		w := httptest.NewRecorder()
		p.Update(w, r)
		if !reflect.DeepEqual(testCase.status, w.Result().StatusCode) {
			t.Errorf("Expected error: %v Got %v", testCase.status, w.Result().StatusCode)
		}
	}
}


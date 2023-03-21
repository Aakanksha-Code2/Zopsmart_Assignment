package patient

import (
	"encoding/json"
	"github.com/aakanksha/ppms/internal/models"
	"github.com/aakanksha/ppms/internal/service"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type https struct {
	svc service.ServiceInterface
}

func New(svc service.ServiceInterface) *https {
	return &https{svc}
}

type ErrorStruct struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"Message"`
}
type ResponseStruct struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func Writer(w http.ResponseWriter, response interface{}, status int) {
	res, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(res))
}

type data struct {
	Patient interface{}
}

func (p *https) GetByID(w http.ResponseWriter, r *http.Request) {
	var response interface{}
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	patient, err := p.svc.GetByID(id)
	if err != nil {
		response = ErrorStruct{
			Code:    http.StatusBadRequest,
			Status:  "Error",
			Message: "Invalid ID",
		}
		Writer(w, response, http.StatusBadRequest)
		return
	}
	response = ResponseStruct{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   data{patient},
	}
	Writer(w, response, http.StatusOK)
}

func (p *https) GetAll(w http.ResponseWriter, r *http.Request) {
	var response interface{}
	patients, err := p.svc.GetAll()
	if err != nil {
		response = ErrorStruct{
			Code:    http.StatusBadRequest,
			Status:  "Error",
			Message: "Invalid ID",
		}
		Writer(w, response, http.StatusBadRequest)
		return
	}
	response = ResponseStruct{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   data{patients},
	}
	Writer(w, response, http.StatusOK)
}

func (p *https) Insert(w http.ResponseWriter, r *http.Request) {
	var response interface{}
	var patient models.Patient
	err := json.NewDecoder(r.Body).Decode(&patient)
	if err != nil {
		response = ErrorStruct{
			Code:    http.StatusBadRequest,
			Status:  "Error",
			Message: err.Error(),
		}
		Writer(w, response, http.StatusBadRequest)
		return
	}
	patientvalue, err := p.svc.Insert(&patient)
	if err != nil {
		response = ErrorStruct{
			Code:    http.StatusBadRequest,
			Status:  "Error",
			Message: err.Error(),
		}
		Writer(w, response, http.StatusBadRequest)
		return
	}
	patientvalue = &models.Patient{Id: patientvalue.Id, Name: patientvalue.Name, Phone: patientvalue.Phone, Discharge: patientvalue.Discharge, BloodGroup: patientvalue.BloodGroup, Description: patientvalue.Description}
	response = ResponseStruct{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   data{*patientvalue},
	}
	Writer(w, response, http.StatusOK)
}

func (p *https) Update(w http.ResponseWriter, r *http.Request) {
	var response interface{}
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var patient *models.Patient
	err := json.NewDecoder(r.Body).Decode(&patient)
	if err != nil {
		response = ErrorStruct{
			Code:    http.StatusBadRequest,
			Status:  "Error",
			Message: err.Error(),
		}
		Writer(w, response, http.StatusBadRequest)
		return
	}
	patient, err = p.svc.Update(patient, id)

	if err != nil {
		response = ErrorStruct{
			Code:    http.StatusBadRequest,
			Status:  "Error",
			Message: err.Error(),
		}
		Writer(w, response, http.StatusBadRequest)
		return
	}

	response = ResponseStruct{
		Code:   http.StatusOK,
		Status: "Success",

		Data: data{patient},
	}
	Writer(w, response, http.StatusOK)
}

func (p *https) Delete(w http.ResponseWriter, r *http.Request) {
	var response interface{}
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	err := p.svc.Delete(id)
	if err != nil {
		response = ErrorStruct{
			Code:    http.StatusBadRequest,
			Status:  "Error",
			Message: err.Error(),
		}
		Writer(w, response, http.StatusBadRequest)
		return
	}
	response = ResponseStruct{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   "Patient deleted Successfully",
	}
	Writer(w, response, http.StatusOK)
}


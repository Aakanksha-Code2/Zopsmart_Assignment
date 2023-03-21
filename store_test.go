package patient

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aakanksha/ppms/internal/models"
	"testing"
	"time"
)

var current_time = time.Now()

func TestInsert(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	testcases := []struct {
		desc        string
		input       *models.Patient
		output      *models.Patient
		mockQuery   []interface{}
		expectError error
	}{
		{
			desc:   "success",
			input:  &models.Patient{Id: 1, Name: "ZopSmart", Phone: "+919172681679", Discharge: true, BloodGroup: "+A", Description: "description"},
			output: &models.Patient{Id: 1, Name: "ZopSmart", Phone: "+919172681679", Discharge: true, CreatedAt: current_time, UpdatedAt: current_time, BloodGroup: "+A", Description: "description"},
			mockQuery: []interface{}{mock.ExpectExec("insert into patient (name,phone,discharge,bloodgroup,description) values (?, ?, ?, ?, ?)").
				WithArgs("ZopSmart", "+919172681679", true, "+A", "description").
				WillReturnResult(sqlmock.NewResult(1, 1)),
				mock.ExpectQuery("select id,name,phone,discharge,createdat,udatedat,bloodgroup,description from patient where deletedat IS NULL and id=?").WithArgs(1).
					WillReturnRows(mock.NewRows([]string{"id", "name", "phone", "discharge", "createdat", "udatedat", "bloodgroup", "description"}).
						AddRow(1, "ZopSmart", "+919172681679", true, current_time, current_time, "+A", "description")),
			},
			expectError: nil,
		},
		{
			desc:   "failure",
			input:  &models.Patient{Id: 1, Name: "ZopSmart", Phone: "+919172681679", Discharge: true, BloodGroup: "+A", Description: "description"},
			output: &models.Patient{Id: 1, Name: "ZopSmart", Phone: "+919172681679", Discharge: true, CreatedAt: current_time, UpdatedAt: current_time, BloodGroup: "+A", Description: "description"},
			mockQuery: []interface{}{mock.ExpectExec("insert into patient (name,phone,discharge,bloodgroup,description) values (?, ?, ?, ?, ?)").
				WithArgs("ZopSmart", "+919172681679", true, "+A", "description").WillReturnError(errors.New("error in executing insert")),
				mock.ExpectQuery("select id,name,phone,discharge,createdat,udatedat,bloodgroup,description from patient where deletedat IS NULL and id=?").WithArgs(1).
					WillReturnError(errors.New("error in executing insert")),
			},
			expectError: errors.New("error in executing insert"),
		},
	}

	for _, testCase := range testcases {
		t.Run(testCase.desc, func(t *testing.T) {
			a := New(db)
			_, err := a.Insert(testCase.input)
			if err != nil && err.Error() != testCase.expectError.Error() {
				t.Errorf("expected error :%v, got :%v ", testCase.expectError, err)
			}

		})
	}
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	tests := []struct {
		desc  string
		id    int
		input *models.Patient
		//output      *models.Patient
		mockQuery   []interface{}
		expectError error
	}{
		{
			desc:  "success",
			id:    1,
			input: &models.Patient{Id: 1, Name: "ZopSmart", Phone: "+919172681679", Discharge: true, UpdatedAt: time.Now(), BloodGroup: "+A", Description: "description"},
			//output: &models.Patient{Id: 1, Name: "ZopSmart", Phone: "+919172681679", Discharge: true, CreatedAt: current_time, UpdatedAt: current_time, BloodGroup: "+A", Description: "description"},
			mockQuery: []interface{}{mock.ExpectExec("update patient SET name = ?, phone=?, discharge=?,udatedat=?,bloodgroup=?,description=? where deletedat IS NULL and id=?").
				WithArgs("ZopSmart", "+919172681679", true, sqlmock.AnyArg(), "+A", "description", int64(1)).
				WillReturnResult(sqlmock.NewResult(1, 1)),
				mock.ExpectQuery("select id,name,phone,discharge,createdat,udatedat,bloodgroup,description from patient where deletedat IS NULL and id=?").WithArgs(1).
					WillReturnRows(mock.NewRows([]string{"id", "name", "phone", "discharge", "createdat", "udatedat", "bloodgroup", "description"}).
						AddRow(1, "ZopSmart", "+919172681679", true, time.Now(), "2022-02-22 13:23:22", "+A", "description")),
			},
			expectError: nil,
		},
		{
			desc:  "FAILURE",
			id:    1,
			input: &models.Patient{Id: 1, Name: "ZopSmart", Phone: "+919172681679", Discharge: true, UpdatedAt: time.Now(), BloodGroup: "+A", Description: "description"},
			//output: &models.Patient{Id: 1, Name: "ZopSmart", Phone: "+919172681679", Discharge: true, CreatedAt: current_time, UpdatedAt: current_time, BloodGroup: "+A", Description: "description"},
			mockQuery: []interface{}{mock.ExpectExec("update patient SET name = ?, phone=?, discharge=?,udatedat=?,bloodgroup=?,description=? where deletedat IS NULL and id=?").
				WithArgs("ZopSmart", "+919172681679", true, sqlmock.AnyArg(), "+A", "description", int64(1)).
				WillReturnError(errors.New("error in update")),
				mock.ExpectQuery("select id,name,phone,discharge,createdat,udatedat,bloodgroup,description from patient where deletedat IS NULL and id=?").WithArgs(1).
					WillReturnError(errors.New("error in update")),
			},
			expectError: errors.New("error in update"),
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.desc, func(t *testing.T) {
			a := New(db)
			_, err := a.Update(testCase.input, testCase.id)
			if err != nil && err.Error() != testCase.expectError.Error() {
				t.Errorf("expected error :%v, got :%v ", testCase.expectError, err)
			}

		})
	}
}

func TestGetById(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	tests := []struct {
		desc        string
		id          int
		output      *models.Patient
		mockQuery   interface{}
		expectError error
	}{
		{
			desc:   "success",
			id:     1,
			output: &models.Patient{Id: 1, Name: "ZopSmart", Phone: "+919172681679", Discharge: true, CreatedAt: current_time, UpdatedAt: current_time, BloodGroup: "+A", Description: "description"},
			mockQuery: mock.ExpectQuery("select id,name,phone,discharge,createdat,udatedat,bloodgroup,description from patient where deletedat IS NULL and id=?").
				WithArgs(1).WillReturnRows(mock.NewRows([]string{"id", "name", "phone", "discharge", "createdat", "udatedat", "bloodgroup", "description"}).
				AddRow(1, "ZopSmart", "+919172681679", true, current_time, current_time, "+A", "description")),
			expectError: nil,
		},
		{
			desc:   "failure",
			id:     1,
			output: &models.Patient{Id: 1, Name: "ZopSmart", Phone: "+919172681679", Discharge: true, CreatedAt: current_time, UpdatedAt: current_time, BloodGroup: "+A", Description: "description"},
			mockQuery: mock.ExpectQuery("select id,name,phone,discharge,createdat,udatedat,bloodgroup,description from patient where deletedat IS NULL and id=?").
				WithArgs(1).WillReturnError(errors.New("error in fetching row")),
			expectError: errors.New("error in fetching row"),
		},
		{
			desc:   "failure",
			id:     1,
			output: &models.Patient{Id: 1, Name: "ZopSmart", Phone: "+919172681679", Discharge: true, CreatedAt: current_time, UpdatedAt: current_time, BloodGroup: "+A", Description: "description"},
			mockQuery: mock.ExpectQuery("select id,name,phone,discharge,createdat,udatedat,bloodgroup,description from patient where deletedat IS NULL and id=?").
				WithArgs(1).WillReturnError(sql.ErrNoRows),
			expectError: sql.ErrNoRows,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.desc, func(t *testing.T) {
			a := New(db)
			_, err := a.GetByID(testCase.id)
			if err != nil && err.Error() != testCase.expectError.Error() {
				t.Errorf("expected error :%v, got :%v ", testCase.expectError, err)
			}

		})
	}
}

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	createat := time.Now()
	updateat := time.Now()
	rows := sqlmock.NewRows([]string{"id", "name", "phone", "discharge", "createdAt", "updatedAt", "bloodGroup", "description"}).
		AddRow(1, "P", "+916354346285", false, createat, updateat, "+b", "Cold").
		AddRow(2, "a", "+916666555653", false, createat, updateat, "+o", "Cold")

	tests := []struct {
		desc        string
		output      []*models.Patient
		mockQuery   interface{}
		expectError error
	}{
		{
			desc:        "success",
			output:      []*models.Patient{{Id: 1, Name: "aakanksha3", Phone: "123", Discharge: true, BloodGroup: "A+", Description: "abc"}},
			mockQuery:   mock.ExpectQuery("select id,name,phone,discharge,createdat,udatedat,bloodgroup,description from patient where deletedat IS NULL;").WillReturnRows(rows),
			expectError: nil,
		},
		{
			desc:        "failure",
			output:      []*models.Patient{{Id: 3, Name: "aakanksha3", Phone: "123", Discharge: true, BloodGroup: "A+", Description: "abc"}},
			mockQuery:   mock.ExpectQuery("select id,name,phone,discharge,createdat,udatedat,bloodgroup,description from patient where deletedat IS NULL;").WillReturnError(errors.New("not passesd correct data")),
			expectError: errors.New("not passesd correct data"),
		},
		{
			desc:        "failures",
			output:      []*models.Patient{{Id: 1, Name: "aakanksha3", Phone: "123", Discharge: true, BloodGroup: "A+", Description: "abc"}},
			mockQuery:   mock.ExpectQuery("select id,name,phone,discharge,createdat,udatedat,bloodgroup,description from patient where deletedat IS NULL;").WillReturnError(errors.New("error in row scan")),
			expectError: errors.New("error in row scan"),
		},
	}

	for _, testCase := range tests {
		t.Run("", func(t *testing.T) {

			a := New(db)
			_, err := a.GetAll()

			if err != nil && err.Error() != testCase.expectError.Error() {
				t.Errorf("expected error :%v, got :%v ", testCase.expectError, err)
			}

		})
	}
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	tests := []struct {
		id          int
		mockQuery   interface{}
		expectError error
	}{
		{
			id:          1,
			mockQuery:   mock.ExpectExec("UPDATE patient SET deletedat=? WHERE id=? AND deletedat IS NULL").WithArgs(sqlmock.AnyArg(), 1).WillReturnResult(sqlmock.NewResult(1, 1)),
			expectError: nil,
		},
		{
			id:          4,
			mockQuery:   mock.ExpectExec("UPDATE patient SET deletedat=? WHERE id=? AND deletedat IS NULL").WithArgs(sqlmock.AnyArg(), 4).WillReturnError(errors.New("error of delete")),
			expectError: errors.New("error of delete"),
		},
	}

	for _, testCase := range tests {
		t.Run("", func(t *testing.T) {

			a := New(db)

			err := a.Delete(testCase.id)
			fmt.Println(err)

			if err != nil && err.Error() != testCase.expectError.Error() {
				t.Errorf("expected error :%v, got :%v ", testCase.expectError, err)
			}

		})
	}
}


package patient

import (
	"database/sql"
	"github.com/aakanksha/ppms/internal/models"
	"time"
)

type store struct {
	db *sql.DB
}

func New(db *sql.DB) *store {
	return &store{db: db}
}
func (s *store) Insert(pt *models.Patient) (*models.Patient, error) {
	query := "insert into patient (name,phone,discharge,bloodgroup,description) values (?, ?, ?, ?, ?)"
	res, err := s.db.Exec(query, pt.Name, pt.Phone, pt.Discharge, pt.BloodGroup, pt.Description)
	if err != nil {
		return nil, err
	}
	lastinserted, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return s.GetByID(int(lastinserted))
}

func (s *store) GetByID(gid int) (*models.Patient, error) {
	var pt models.Patient
	query := "select id,name,phone,discharge,createdat,udatedat,bloodgroup,description from patient where deletedat IS NULL and id=?"
	row := s.db.QueryRow(query, gid)
	err := row.Scan(&pt.Id, &pt.Name, &pt.Phone, &pt.Discharge, &pt.CreatedAt, &pt.UpdatedAt, &pt.BloodGroup, &pt.Description)
	if err == sql.ErrNoRows {
		return nil, err
	}
	return &pt, nil
}

func (s *store) GetAll() ([]*models.Patient, error) {
	query := "select id,name,phone,discharge,createdat,udatedat,bloodgroup,description from patient where deletedat IS NULL;"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	var patients []*models.Patient
	defer rows.Close()
	for rows.Next() {
		var pt models.Patient
		err := rows.Scan(&pt.Id, &pt.Name, &pt.Phone, &pt.Discharge, &pt.CreatedAt, &pt.UpdatedAt, &pt.BloodGroup, &pt.Description)
		if err != nil {
			return nil, err
		}
		patients = append(patients, &pt)
	}
	return patients, nil
}

func (s *store) Update(pt *models.Patient, uid int) (*models.Patient, error) {

	query := "update patient SET name = ?, phone=?, discharge=?,udatedat=?,bloodgroup=?,description=? where deletedat IS NULL and id=?"
	_, err := s.db.Exec(query, &pt.Name, &pt.Phone, &pt.Discharge, time.Now(), &pt.BloodGroup, &pt.Description, uid)

	if err != nil {
		return nil, err
	}
	return s.GetByID(uid)
}

func (s *store) Delete(did int) error {
	query := "UPDATE patient SET deletedat=? WHERE id=? AND deletedat IS NULL"
	uDeletedAt := time.Now()
	_, err := s.db.Exec(query, uDeletedAt, did)

	return err
}


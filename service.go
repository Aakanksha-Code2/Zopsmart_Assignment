package patient

import (
	"errors"
	"github.com/aakanksha/ppms/internal/models"
	"github.com/aakanksha/ppms/internal/stores"
)

type Svc struct {
	stores stores.StoreInterface
}

func New(stores stores.StoreInterface) *Svc {
	return &Svc{stores}
}

func (ps *Svc) GetAll() ([]*models.Patient, error) {
	res, err := ps.stores.GetAll()
	return res, err
}

func (ps *Svc) GetByID(id int) (*models.Patient, error) {
	if !validId(id) {
		return nil, errors.New("invalid id")
	}
	patient, err := ps.stores.GetByID(id)
	if err != nil {
		return nil, err
	}
	return patient, err
}

func (ps *Svc) Insert(p *models.Patient) (*models.Patient, error) {
	if !validatename(p.Name) {
		return &models.Patient{}, errors.New("invalid name")
	}
	res, err := ps.stores.Insert(p)
	return res, err
}
func (ps *Svc) Update(p *models.Patient, id int) (*models.Patient, error) {
	if !validId(id) {
		return nil, errors.New("invalid id")
	}
	result, err := ps.stores.GetByID(id)

	if result == nil {
		return nil, err
	}
	update, err1 := ps.stores.Update(p, id)

	return update, err1
}
func (ps *Svc) Delete(id int) error {
	if !validId(id) {
		return errors.New("invalid id")
	}
	_, err := ps.stores.GetByID(id)
	if err != nil {
		return err
	}
	err = ps.stores.Delete(id)
	return err
}


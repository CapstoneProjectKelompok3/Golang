package service

import (
	"errors"
	"project-capston/features/unit"

	"github.com/go-playground/validator/v10"
)

type UnitService struct {
	unitService unit.UnitDataInterface
	validate    *validator.Validate
}

// GetAll implements unit.UnitServiceInterface.
func (service *UnitService) GetAll(param unit.QueryParams, token string) (bool, []unit.UnitEntity, error) {
	var totalPages int64
	nextPage := true
	count, data, err := service.unitService.SelectAll(param, token)
	if err != nil {
		return true, nil, err
	}
	if count == 0 {
		nextPage = false
	}

	if param.IsClassDashboard {
		totalPages = count / int64(param.ItemsPerPage)
		if count%int64(param.ItemsPerPage) != 0 {
			totalPages += 1
		}
		if param.Page == int(totalPages) {
			nextPage = false
		}
		if param.Page < param.ItemsPerPage {
			nextPage = false
		}

		if data == nil {
			nextPage = false
		}
	}
	return nextPage, data, nil
}

// GetById implements unit.UnitServiceInterface.
func (service *UnitService) GetById(id uint, token string) (unit.UnitEntity, error) {
	data, err := service.unitService.SelectById(id, token)
	if err != nil {
		return unit.UnitEntity{}, err
	}
	return data, nil
}

// Edit implements unit.UnitServiceInterface.
func (repo *UnitService) Edit(input unit.UnitEntity, id uint) error {
	err := repo.unitService.Update(input, id)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements unit.UnitServiceInterface.
func (service *UnitService) Delete(id uint) error {
	err := service.unitService.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// Add implements unit.UnitServiceInterface.
func (service *UnitService) Add(input unit.UnitEntity) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errors.New("error validate")
	}

	if input.GovermentType != "Polisi" && input.GovermentType != "Rumah Sakit" && input.GovermentType != "Damkar" && input.GovermentType != "DISHUB" && input.GovermentType != "SAR" {
		return errors.New("input is wrong")
	}
	err := service.unitService.Insert(input)
	if err != nil {
		return err
	}
	return nil
}

func New(service unit.UnitDataInterface) unit.UnitServiceInterface {
	return &UnitService{
		unitService: service,
		validate:    validator.New(),
	}
}

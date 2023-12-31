package service

import (
	"errors"
	"project-capston/features/goverment"

	"github.com/go-playground/validator/v10"
)

type govermentService struct {
	governmentData goverment.GovernmentDataInterface
	validate       *validator.Validate
}



func New(repo goverment.GovernmentDataInterface) goverment.GovernmentServiceInterface {
	return &govermentService{
		governmentData: repo,
		validate:       validator.New(),
	}
}

// Create implements goverment.GovernmentServiceInterface.
func (service *govermentService) Create(input goverment.Core) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errors.New("validation error" + errValidate.Error())
	}

	err := service.governmentData.Insert(input)
	return err
}

// GetAll implements goverment.GovernmentServiceInterface.
func (service *govermentService) GetAll(pageNumber int, pageSize int) ([]goverment.Core, error) {
	result, err := service.governmentData.SelectAll(pageNumber, pageSize)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetById implements goverment.GovernmentServiceInterface.
func (service *govermentService) GetById(id uint) (goverment.Core, error) {
	result, err := service.governmentData.Select(id)
	if err != nil {
		return goverment.Core{}, err
	}
	return result, nil
}

// EditById implements goverment.GovernmentServiceInterface.
func (service *govermentService) EditById(id uint, input goverment.Core) error {
	err := service.governmentData.Update(id, input)
	return err
}

// DeleteById implements goverment.GovernmentServiceInterface.
func (service *govermentService) DeleteById(id uint) error {
	err := service.governmentData.Delete(id)
	return err
}

// GetNearestLocation implements goverment.GovernmentServiceInterface.
func (service *govermentService) GetNearestLocation(latitude float64, longitude float64, radius float64) ([]goverment.Location, error) {
	result, err := service.governmentData.SelectNearestLocation(latitude, longitude, radius)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCountUnit implements goverment.GovernmentServiceInterface.
func (service *govermentService) GetCountUnit(level string) (goverment.UnitCount, error) {
	if level !="superadmin"{
		return goverment.UnitCount{},errors.New("hanya superadmin yang dapat melihat jumlah unit")
	}
	data,err:=service.governmentData.SelectCountUnit()
	if err != nil{
		return goverment.UnitCount{},err
	}
	return data,nil
}
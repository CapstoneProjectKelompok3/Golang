package service

import (
	"errors"
	"project-capston/features/emergency"

	"github.com/go-playground/validator/v10"
)

type EmergencyService struct {
	emergencyService emergency.EmergencyDataInterface
	validate         *validator.Validate
}

// GetById implements emergency.EmergencyServiceInterface.
func (service *EmergencyService) GetById(id uint) (emergency.EmergencyEntity, error) {
	data,err:=service.emergencyService.SelectById(id)
	if err != nil{
		return emergency.EmergencyEntity{},err
	}
	return data,nil
}

// Edit implements emergency.EmergencyServiceInterface.
func (repo *EmergencyService) Edit(input emergency.EmergencyEntity, id uint) error {
	err := repo.emergencyService.Update(input, id)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements emergency.EmergencyServiceInterface.
func (service *EmergencyService) Delete(id uint) error {
	err := service.emergencyService.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// Add implements emergency.EmergencyServiceInterface.
func (service *EmergencyService) Add(input emergency.EmergencyEntity) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errors.New("error validate, receiver_id/longitude/latitude require")
	}
	err := service.emergencyService.Insert(input)
	if err != nil {
		return err
	}
	return nil
}

func New(service emergency.EmergencyDataInterface) emergency.EmergencyServiceInterface {
	return &EmergencyService{
		emergencyService: service,
		validate:         validator.New(),
	}
}

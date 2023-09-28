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

// Delete implements emergency.EmergencyServiceInterface.
func (service *EmergencyService) Delete(id uint) error {
	err:=service.emergencyService.Delete(id)
	if err != nil{
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

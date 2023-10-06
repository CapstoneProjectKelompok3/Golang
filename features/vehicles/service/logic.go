package service

import (
	"errors"
	"project-capston/features/vehicles"

	"github.com/go-playground/validator/v10"
)

type VehicleService struct {
	vehicleService vehicles.VehicleDataInterface
	validate       *validator.Validate
}

// Delete implements vehicles.VehicleServiceInterface.
func (service *VehicleService) Delete(id uint) error {
	err:=service.vehicleService.Delete(id)
	if err != nil{
		return err
	}
	return nil
}

// GetAll implements vehicles.VehicleServiceInterface.
func (service *VehicleService) GetAll() ([]vehicles.VehicleEntity, error) {
	data, err := service.vehicleService.SelectAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetById implements vehicles.VehicleServiceInterface.
func (service *VehicleService) GetById(id uint) (vehicles.VehicleEntity, error) {
	data, err := service.vehicleService.SelectById(id)
	return data, err
}

// Edit implements vehicles.VehicleServiceInterface.
func (service *VehicleService) Edit(input vehicles.VehicleEntity, id uint,level string) error {
	if level !="admin"{
		return errors.New("hanya admin yang dapat mengedit kendaraan")
	}
	err := service.vehicleService.Update(input, id)
	if err != nil {
		return err
	}
	return nil
}

// Add implements vehicles.VehicleServiceInterface.
func (service *VehicleService) Add(input vehicles.VehicleEntity,level string) error {

	if level !="admin"{
		return errors.New("hanya admin yang dapat menambah kendaraan")
	}
	errValide := service.validate.Struct(input)
	if errValide != nil {
		return errors.New("validate, plate dan goverment_id required")
	}
	err := service.vehicleService.Insert(input)
	if err != nil {
		return err
	}
	return nil
}

func New(service vehicles.VehicleDataInterface) vehicles.VehicleServiceInterface {
	return &VehicleService{
		vehicleService: service,
		validate:       validator.New(),
	}
}

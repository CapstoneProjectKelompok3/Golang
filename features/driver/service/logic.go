package service

import (
	"errors"
	"project-capston/features/driver"

	"github.com/go-playground/validator/v10"
)

type driverService struct {
	driverData driver.DriverDataInterface
	validate   *validator.Validate
}

func New(repo driver.DriverDataInterface) driver.DriverServiceInterface {
	return &driverService{
		driverData: repo,
		validate:   validator.New(),
	}
}

// Create implements driver.DriverServiceInterface.
func (service *driverService) Create(input driver.Core) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errors.New("validation error" + errValidate.Error())
	}

	err := service.driverData.Insert(input)
	return err
}

// GetAll implements driver.DriverServiceInterface.
func (service *driverService) GetAll(pageNumber int, pageSize int) ([]driver.Core, error) {
	result, err := service.driverData.SelectAll(pageNumber, pageSize)
	if err != nil {
		return nil, err
	}
	return result, nil
}

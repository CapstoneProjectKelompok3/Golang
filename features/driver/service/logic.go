package service

import (
	"errors"
	"project-capston/app/middlewares"
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
func (service *driverService) GetAll(pageNumber int, pageSize int) ([]driver.DriverCore, error) {
	result, err := service.driverData.SelectAll(pageNumber, pageSize)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Login implements driver.DriverServiceInterface.
func (service *driverService) Login(email string, password string) (dataLogin driver.Core, token string, err error) {
	dataLogin, err = service.driverData.Login(email, password)
	if err != nil {
		return driver.Core{}, "", err
	}
	token, err = middlewares.CreateToken(dataLogin.Id)
	if err != nil {
		return driver.Core{}, "", err
	}
	return dataLogin, token, nil
}

// KerahkanDriver implements driver.DriverServiceInterface.
func (service *driverService) KerahkanDriver(lat string, long string, police int, hospital int, firestation int, dishub int, SAR int) ([]driver.DriverCore, error) {
	result, err := service.driverData.KerahkanDriver(lat, long, police, hospital, firestation, dishub, SAR)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SelectProfile implements driver.DriverServiceInterface.
func (service *driverService) GetProfile(id int) (driver.Core, error) {
	result, err := service.driverData.SelectProfile(id)
	return result, err
}

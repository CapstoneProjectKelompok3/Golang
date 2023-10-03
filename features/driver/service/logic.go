package service

import (
	"errors"
	"fmt"
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
	fmt.Println("email", email)
	fmt.Println("email", password)
	fmt.Println("dataLogin", dataLogin)
	if err != nil {
		return driver.Core{}, "", err
	}
	token, err = middlewares.CreateToken(dataLogin.Id)
	if err != nil {
		return driver.Core{}, "", err
	}
	return dataLogin, token, nil
}

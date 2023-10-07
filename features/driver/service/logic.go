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

// AcceptOrRejectOrder implements driver.DriverServiceInterface.
func (service *driverService) AcceptOrRejectOrder(idEmergenci uint,IsAccepted bool, idDriver int) error {
	// errValidate := service.validate.Struct(IsAccepted)
	// if errValidate != nil {
	// 	return errors.New("validation error" + errValidate.Error())
	// }

	err := service.driverData.AcceptOrRejectOrder(IsAccepted, idDriver)

	if IsAccepted {

		data,errData:=service.driverData.SelectProfile(idDriver)
		if errData != nil{
			return errData
		}
		if data.Token ==""{
			return errors.New("maaf anda telah mencancel orderan, token pindah ke driver lain")
		}
		id,tipeUnit,errUnit:=service.driverData.SelectUnit(idEmergenci)
		if errUnit != nil{
			return errUnit
		}
		var idUnit []uint
		for i:=0;i<len(tipeUnit);i++{
			if data.GovermentType == tipeUnit[i]{
				idUnit=append(idUnit, id[i])
			}
		}
		idHistory,errHis:=service.driverData.SelectHistori(idUnit[0])
		if errHis != nil{
			return errHis
		}
		errUpdate:=service.driverData.UpdateHistoryUnit(uint(idDriver),idHistory)
		if errUpdate !=nil{
			return errUpdate
		}
	
	}
	
	return err
}

// KerahkanDriver implements driver.DriverServiceInterface.
func (service *driverService) KerahkanDriver(id uint,lat string, long string, police int, hospital int, firestation int, dishub int, SAR int) ([]driver.DriverCore, error) {

	result, err := service.driverData.KerahkanDriver(lat, long, police, hospital, firestation, dishub, SAR)
	if err != nil {
		return nil, err
	}
	count := []int{police, hospital, firestation, dishub, SAR}
	tipe :=[]string{"police", "hospital", "firestation", "dishub", "SAR"}
	errUnit:=service.driverData.CreateUnit(id,tipe,count)
	if errUnit != nil{
		return nil,errUnit
	}
	errHistori:=service.driverData.CreateUnitHistori(id)
	if errHistori != nil{
		return nil,errUnit
	}

	return result, nil
}

// SelectProfile implements driver.DriverServiceInterface.
func (service *driverService) GetProfile(id int) (driver.DriverCore, error) {
	result, err := service.driverData.SelectProfile(id)
	return result, err
}

// DriverOnTrip implements driver.DriverServiceInterface.
func (service *driverService) DriverOnTrip(id int, lat float64, long float64) (driver.DriverCore, error) {
	// errValidate := service.validate.Struct(long)
	// if errValidate != nil {
	// 	return errors.New("validation error" + errValidate.Error())
	// }
	result, err := service.driverData.DriverOnTrip(id, lat, long)
	return result, err
}

// FinishTrip implements driver.DriverServiceInterface.
func (service *driverService) FinishTrip(id int) error {
	err := service.driverData.FinishTrip(id)
	return err
}

// Logout implements driver.DriverServiceInterface.
func (service *driverService) Logout(id int) error {
	err := service.driverData.Logout(id)
	return err
}

// GetCountDriver implements driver.DriverServiceInterface.
func (service *driverService) GetCountDriver() (int64, error) {
	count, err := service.driverData.SelectCountDriver()
	if err != nil {
		return 0, err
	}
	return count, nil
}

// Delete implements driver.DriverServiceInterface.
func (service *driverService) Delete(id uint) error {
	err := service.driverData.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

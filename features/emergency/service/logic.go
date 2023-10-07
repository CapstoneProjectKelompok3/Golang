package service

import (
	"errors"
	"fmt"
	"project-capston/features/emergency"
	"project-capston/helper"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type EmergencyService struct {
	emergencyService emergency.EmergencyDataInterface
	validate         *validator.Validate
}

// SumEmergency implements emergency.EmergencyServiceInterface.
func (service *EmergencyService) SumEmergency() (int64, error) {
	count,err:=service.emergencyService.SumEmergency()
	if err != nil{
		return 0,err
	}
	return count,nil
}

// ActionGmail implements emergency.EmergencyServiceInterface.
func (service *EmergencyService) ActionGmail(input string) error {
	err := service.emergencyService.ActionGmail(input)
	if err != nil {
		return err
	}
	return nil
}

// GetAll implements emergency.EmergencyServiceInterface.
func (service *EmergencyService) GetAll(param emergency.QueryParams, token string,idCall uint,level string) (bool, []emergency.EmergencyEntity, error) {
	var totalPages int64
	nextPage := true
	count, data, err := service.emergencyService.SelectAll(param, token,idCall,level)
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

// GetById implements emergency.EmergencyServiceInterface.
func (service *EmergencyService) GetById(id uint, token string) (emergency.EmergencyEntity, error) {
	data, err := service.emergencyService.SelectById(id, token)
	if err != nil {
		return emergency.EmergencyEntity{}, err
	}
	return data, nil
}

// Edit implements emergency.EmergencyServiceInterface.
func (repo *EmergencyService) Edit(input emergency.EmergencyEntity, id uint, level string, idUser uint) error {

	if level != "admin" {
		return errors.New("hanya admin yang dapat mengedit emergency")
	}
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
func (service *EmergencyService) Add(input emergency.EmergencyEntity, token string) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errors.New("error validate, receiver_id/longitude/latitude require")
	}

	idInsert, errInsert := service.emergencyService.Insert(input)
	if errInsert != nil {
		return errInsert
	}

	name := fmt.Sprintf("Kasus %d", idInsert)
	input.Name = name
	errUpdate := service.emergencyService.Update(input, idInsert)
	if errUpdate != nil {
		return errUpdate
	}

	idCall := strconv.Itoa(int(input.CallerID))
	dataUserCall, errUserCall := service.emergencyService.SelectUser(idCall, token)
	if errUserCall != nil {
		return errUserCall
	}
	notif := helper.MessageGomailE{
		EmailReceiver: dataUserCall.Email,
		Sucject:       name,
		Content:       "Kasus terbaru yang harus sedang ditangani, semoga user dapat tenang dan menunggu notifikasi selanjutnya",
		Name:          dataUserCall.Name,
		Email:         dataUserCall.Email,
	}
	status, errEmail := service.emergencyService.SendNotification(notif)
	if errEmail != nil {
		return errors.New("gagal send email from admin")
	}
	fmt.Println("status email", status)

	return nil
}

func New(service emergency.EmergencyDataInterface) emergency.EmergencyServiceInterface {
	return &EmergencyService{
		emergencyService: service,
		validate:         validator.New(),
	}
}

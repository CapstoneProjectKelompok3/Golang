package service

import (
	"errors"
	history "project-capston/features/history"

	"github.com/go-playground/validator/v10"
)

type HistoryService struct {
	historyService history.HistoryDataInterface
	validate       *validator.Validate
}

// GetAll implements history.HistoryServiceInterface.
func (service *HistoryService) GetAll(param history.QueryParams, token string) (bool, []history.HistoryEntity, error) {
	var totalPages int64
	nextPage := true
	count, data, err := service.historyService.SelectAll(param, token)
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

// GetById implements history.HistoryServiceInterface.
func (service *HistoryService) GetById(id uint, token string) (history.HistoryEntity, error) {
	data, err := service.historyService.SelectById(id, token)
	if err != nil {
		return history.HistoryEntity{}, err
	}
	return data, nil
}

// Edit implements history.HistoryServiceInterface.
func (repo *HistoryService) Edit(input history.HistoryEntity, id uint) error {
	err := repo.historyService.Update(input, id)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements history.HistoryServiceInterface.
func (service *HistoryService) Delete(id uint) error {
	err := service.historyService.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// Add implements history.HistoryServiceInterface.
func (service *HistoryService) Add(input history.HistoryEntity) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errors.New("error validate")
	}
	err := service.historyService.Insert(input)
	if err != nil {
		return err
	}
	return nil
}

func New(service history.HistoryDataInterface) history.HistoryServiceInterface {
	return &HistoryService{
		historyService: service,
		validate:       validator.New(),
	}
}

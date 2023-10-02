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

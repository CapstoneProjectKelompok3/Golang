package handler

import (
	"net/http"
	"project-capston/features/goverment"
	"project-capston/helper"
	"strings"

	"github.com/labstack/echo/v4"
)

type governmentHandler struct {
	governmentService goverment.GovernmentServiceInterface
}

func New(service goverment.GovernmentServiceInterface) *governmentHandler {
	return &governmentHandler{
		governmentService: service,
	}
}

func (handler *governmentHandler) CreateGovernment(c echo.Context) error {
	userInput := new(GovernmentRequest)
	errBind := c.Bind(&userInput)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	governmentCore := RequestToCore(*userInput)
	err := handler.governmentService.Create(governmentCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else if strings.Contains(err.Error(), "for key 'governments.name'") {
			return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusConflict, "Government with this name already exists", nil))
		}
	}
	return c.JSON(http.StatusCreated, helper.WebResponse(http.StatusCreated, "success insert data", nil))
}

package handler

import (
	"net/http"
	"project-capston/features/goverment"
	"project-capston/helper"
	"strconv"
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

func (handler *governmentHandler) GetAllGovernment(c echo.Context) error {
	pageNumber, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("size"))

	if pageNumber <= 0 {
		pageNumber = 1
	}
	if pageSize <= 0 {
		pageSize = 100
	}

	result, err := handler.governmentService.GetAll(int(pageNumber), int(pageSize))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "error read data", nil))
	}

	var governmentResponse []GovernmentResponse
	for _, value := range result {
		governmentResponse = append(governmentResponse, GovernmentResponse{
			ID:        value.ID,
			Name:      value.Name,
			Type:      value.Type,
			Address:   value.Address,
			Latitude:  value.Latitude,
			Longitude: value.Longitude,
		})
	}
	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "success read data", governmentResponse))
}

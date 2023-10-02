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

func (handler *governmentHandler) GetGovernmentById(c echo.Context) error {
	id := c.Param("government_id")

	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "error id not valid", nil))
	}

	result, err := handler.governmentService.GetById(uint(idConv))
	if err != nil {

		return c.JSON(http.StatusNotFound, helper.WebResponse(http.StatusNotFound, "data not found", nil))
	}

	// _, _, userCompanyId := middlewares.ExtractTokenUserId(c)
	// if idConv != userCompanyId {
	// 	return c.JSON(http.StatusForbidden, helpers.WebResponse(http.StatusForbidden, exception.ErrForbiddenAccess.Error(), nil))
	// } else {
	resultResponse := GovernmentResponse{
		ID:        result.ID,
		Name:      result.Name,
		Type:      result.Type,
		Address:   result.Address,
		Latitude:  result.Latitude,
		Longitude: result.Longitude,
	}
	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "success read data", resultResponse))
	// }
}

func (handler *governmentHandler) UpdateGovernment(c echo.Context) error {
	id := c.Param("government_id")

	idParam, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "error data id. data not valid", nil))
	}

	// _, roleName, userCompanyId := middlewares.ExtractTokenUserId(c)
	// if roleName == "Non-HR" {
	// 	return c.JSON(http.StatusForbidden, helpers.WebResponse(http.StatusForbidden, exception.ErrForbiddenAccess.Error(), nil))
	// }
	// if idParam != userCompanyId {
	// 	return c.JSON(http.StatusForbidden, helpers.WebResponse(http.StatusForbidden, exception.ErrForbiddenAccess.Error(), nil))
	// } else {
	userInput := new(GovernmentRequest)
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}
	governmentCore := RequestToCore(*userInput)
	err := handler.governmentService.EditById(uint(idParam), governmentCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else if strings.Contains(err.Error(), "for key 'governments.name'") {
			return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "This government name already exist please try again", nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "error update data government", nil))
		}
	}
	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "success update data government", nil))
}

// }

func (handler *governmentHandler) DeleteGovernment(c echo.Context) error {
	id := c.Param("government_id")
	idCompany, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "error id not valid", nil))
	}

	err := handler.governmentService.DeleteById(uint(idCompany))
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "error delete data", nil))

		}
	}
	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "success delete government", nil))
}

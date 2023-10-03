package handler

import (
	"fmt"
	"net/http"
	"project-capston/features/driver"
	"project-capston/helper"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type DriverHandler struct {
	driverService driver.DriverServiceInterface
}

func New(service driver.DriverServiceInterface) *DriverHandler {
	return &DriverHandler{
		driverService: service,
	}
}

func (handler *DriverHandler) CreateDriver(c echo.Context) error {
	driverInput := new(DriverRequest)

	fmt.Println(driverInput)

	errBind := c.Bind(&driverInput) // mendapatkan data yang dikirim oleh FE melalui request body
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	driverCore := RequestToCore(*driverInput)

	err := handler.driverService.Create(driverCore)

	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, err.Error(), nil))
		}
	}

	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusCreated, "success insert data", nil))
}

func (handler *DriverHandler) GetAllDriver(c echo.Context) error {
	pageNumber, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("size"))

	if pageNumber <= 0 {
		pageNumber = 1
	}
	if pageSize <= 0 {
		pageSize = 100
	}

	// _, roleName, _ := middlewares.ExtractTokenUserId(c)
	// if roleName != "Superadmin" {
	// 	return c.JSON(http.StatusForbidden, helpers.WebResponse(http.StatusForbidden, exception.ErrForbiddenAccess.Error(), nil))
	// }

	result, err := handler.driverService.GetAll(int(pageNumber), int(pageSize))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "error read data", nil))
	}
	var driverResponse []DriverResponse

	for _, value := range result {

		var status string
		if value.Status {
			status = "available"
		} else {
			status = "not_available"
		}
		driverResponse = append(driverResponse, DriverResponse{
			Id: value.Id,

			GovermentName: value.GovermentName,
			Fullname:      value.Fullname,
			Token:         value.Token,
			Status:        status,
			DrivingStatus: value.DrivingStatus,
			VehicleID:     value.VehicleID,
			Latitude:      value.Latitude,
			Longitude:     value.Longitude,
		})
	}
	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "success read data", driverResponse))
}

func (handler *DriverHandler) Login(c echo.Context) error {
	userInput := new(LoginDriverRequest)

	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	dataLogin, token, err := handler.driverService.Login(userInput.Email, userInput.Password)

	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "error login", nil))

		}
	}

	response := LoginResponse{
		Id: dataLogin.Id,
		// GovermentName: dataLogin.G,
		Fullname:      dataLogin.Fullname,
		Token:         token,
		Status:        dataLogin.Status,
		DrivingStatus: dataLogin.DrivingStatus,
		VehicleID:     dataLogin.VehicleID,
		Latitude:      dataLogin.Latitude,
		Longitude:     dataLogin.Longitude,
	}

	// response := map[string]any{
	// 	"token":   token,
	// 	"user_id": dataLogin.Id,
	// 	"name":    dataLogin.Fullname,
	// }

	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusCreated, "driver success login", response))
}

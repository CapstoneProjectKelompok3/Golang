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

	totalData := 0
	for _, value := range result {
		totalData++
		driverResponse = append(driverResponse, DriverResponse{
			Id:            value.Id,
			GovermentName: value.GovermentName,
			GovermentType: value.GovermentType,
			Email:         value.Email,
			Fullname:      value.Fullname,
			Token:         value.Token,
			Status:        value.Status,
			DrivingStatus: value.DrivingStatus,
			VehicleID:     value.VehicleID,
			Latitude:      value.Latitude,
			Longitude:     value.Longitude,
		})
	}
	return c.JSON(http.StatusOK, helper.WebResponsePagination(http.StatusOK, totalData, "success read data", driverResponse))
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

func (handler *DriverHandler) KerahkanDriver(c echo.Context) error {
	totalPolice := c.QueryParam("police")
	totalPoliceConv, _ := strconv.Atoi(totalPolice)

	totalHospital := c.QueryParam("hospital")
	totalHospitalConv, _ := strconv.Atoi(totalHospital)

	totalFirestation := c.QueryParam("firestation")
	totalFirestationConv, _ := strconv.Atoi(totalFirestation)

	totalDishub := c.QueryParam("dishub")
	totalDishubConv, _ := strconv.Atoi(totalDishub)

	totalSAR := c.QueryParam("SAR")
	totalSARConv, _ := strconv.Atoi(totalSAR)

	result, err := handler.driverService.KerahkanDriver(totalPoliceConv, totalHospitalConv, totalFirestationConv, totalDishubConv, totalSARConv)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "error read data", nil))
	}

	var driverResponse []DriverAvailableResponse

	totalData := 0

	for _, value := range result {
		totalData++
		driverResponse = append(driverResponse, DriverAvailableResponse{
			Id:            value.Id,
			GovermentName: value.GovermentName,
			GovermentType: value.GovermentType,
			Email:         value.Email,
			Fullname:      value.Fullname,
			Token:         value.Token,
			Status:        value.Status,
			DrivingStatus: value.DrivingStatus,
			VehicleID:     value.VehicleID,
			Latitude:      value.Latitude,
			Longitude:     value.Longitude,
			Radius:        value.Distance,
		})
	}
	return c.JSON(http.StatusOK, helper.WebResponsePagination(http.StatusOK, totalData, "kerahkan driver", driverResponse))
}

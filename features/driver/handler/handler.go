package handler

import (
	"context"
	"fmt"
	"net/http"
	"project-capston/app/middlewares"
	"project-capston/features/driver"
	"project-capston/helper"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

var ctx = context.Background()

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

	errBind := c.Bind(&driverInput) // mendapatkan data yang dikirim oleh FE melalui request body
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	driverCore := RequestToCore(*driverInput)

	err := handler.driverService.Create(driverCore)

	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "error validation", nil))
		} else if strings.Contains(err.Error(), "for key 'drivers.email'") {
			return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "Drivers with this account already exist", nil))
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
			return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, err.Error(), nil))

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
	latLonInput := new(LatLonRequest)

	errBind := c.Bind(&latLonInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	lat := fmt.Sprintf("%f", latLonInput.Lat)
	lon := fmt.Sprintf("%f", latLonInput.Longitude)

	totalPolice := c.QueryParam("police")
	totalPoliceConv, _ := strconv.Atoi(totalPolice)

	emergencyId := c.QueryParam("emergency_id")
	emergencyIdConv, _ := strconv.Atoi(emergencyId)

	totalHospital := c.QueryParam("hospital")
	totalHospitalConv, _ := strconv.Atoi(totalHospital)

	totalFirestation := c.QueryParam("firestation")
	totalFirestationConv, _ := strconv.Atoi(totalFirestation)

	totalDishub := c.QueryParam("dishub")
	totalDishubConv, _ := strconv.Atoi(totalDishub)

	totalSAR := c.QueryParam("SAR")
	totalSARConv, _ := strconv.Atoi(totalSAR)

	idEmergency := c.QueryParam("emergency_id")
	idEmergencyConv, _ := strconv.Atoi(idEmergency)

	result, err := handler.driverService.KerahkanDriver(uint(idEmergencyConv), lat, lon, totalPoliceConv, totalHospitalConv, totalFirestationConv, totalDishubConv, totalSARConv, emergencyIdConv)
	fmt.Println("Result", idEmergencyConv)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "error read data", nil))
	}

	var driverResponse []DriverAvailableResponse

	totalData := 0

	for _, value := range result {
		totalData++
		fmt.Println("Handler", value.Distance)
		driverResponse = append(driverResponse, DriverAvailableResponse{
			Id:            value.Id,
			GovermentName: value.GovermentName,
			GovermentType: value.GovermentType,
			Email:         value.Email,
			Fullname:      value.Fullname,
			Token:         value.Token,
			Status:        value.Status,
			EmergencyID:   value.EmergencyID,
			EmergencyName: value.EmergencyName,
			DrivingStatus: value.DrivingStatus,
			VehicleID:     value.VehicleID,
			Latitude:      value.Latitude,
			Longitude:     value.Longitude,
			Distance:      value.Distance,
		})
	}
	if totalData == 0 {
		return c.JSON(http.StatusOK, helper.WebResponsePagination(http.StatusOK, totalData, "All personil are on duty", nil))

	} else {
		return c.JSON(http.StatusOK, helper.WebResponsePagination(http.StatusOK, totalData, "Success Find and assign driver", driverResponse))

	}

}

func (handler *DriverHandler) GetProfileDriver(c echo.Context) error {
	fmt.Println("ID TOKEN")
	idToken := middlewares.ExtractTokenDriverId(c)

	fmt.Println("ID TOKEN", idToken)

	result, err := handler.driverService.GetProfile(idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "error read data", nil))
	}

	driverResponse := DriverResponse{
		Id:            result.Id,
		GovermentName: result.GovermentName,
		GovermentType: result.GovermentType,
		Email:         result.Email,
		Fullname:      result.Fullname,
		Token:         result.Token,
		Status:        result.Status,
		DrivingStatus: result.DrivingStatus,
		VehicleID:     result.VehicleID,
		EmergenciesID: result.EmergenciesID,
		EmergencyName: result.EmergencyName,
		Latitude:      result.Latitude,
		Longitude:     result.Longitude,
	}

	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "success get my profile", driverResponse))
}

func (handler *DriverHandler) DriverOnTrip(c echo.Context) error {
	idToken := middlewares.ExtractTokenDriverId(c)

	driverInput := new(LatLonRequest)

	errBind := c.Bind(&driverInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	result, err := handler.driverService.DriverOnTrip(idToken, driverInput.Lat, driverInput.Longitude)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "error read data", nil))
	}

	fmt.Println("Result", result.Distance)

	driverResponse := DriverAvailableResponse{
		Id:            result.Id,
		GovermentName: result.GovermentName,
		GovermentType: result.GovermentType,
		Email:         result.Email,
		Fullname:      result.Fullname,
		Token:         result.Token,
		Status:        result.Status,
		DrivingStatus: result.DrivingStatus,
		VehicleID:     result.VehicleID,
		Latitude:      result.Latitude,
		Longitude:     result.Longitude,
		Distance:      result.Distance,
	}

	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "success get my position", driverResponse))
}

var message string

func (handler *DriverHandler) DriverAcceptOrRejectOrder(c echo.Context) error {
	idToken := middlewares.ExtractTokenDriverId(c)

	fmt.Println(idToken)
	idEmergensi := c.QueryParam("emergensi_id")
	idConv, _ := strconv.Atoi(idEmergensi)
	driverInput := new(AcceptOrRejectOrderRequest)

	errBind := c.Bind(&driverInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	if driverInput.IsAccepted {
		message = "Success Accepted order"
	} else {
		message = "Success Rejected order"
	}

	err := handler.driverService.AcceptOrRejectOrder(uint(idConv), driverInput.IsAccepted, idToken)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusNotFound, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, message, nil))
}

func (handler *DriverHandler) GetCountDriver(c echo.Context) error {
	count, err := handler.driverService.GetCountDriver()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]any{
		"status":         "success",
		"jumlah_petugas": count,
	})
}

func (handler *DriverHandler) DriverLogout(c echo.Context) error {
	idToken := middlewares.ExtractTokenDriverId(c)

	err := handler.driverService.Logout(idToken)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusNotFound, err.Error(), nil))
	}
	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "Success Logout", nil))
}

func (handler *DriverHandler) DriverFinishedTrip(c echo.Context) error {
	idToken := middlewares.ExtractTokenDriverId(c)
	id := c.QueryParam("emergenci_id")

	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return errConv
	}

	err := handler.driverService.FinishTrip(idToken, uint(idConv))

	fmt.Println(err)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusNotFound, err.Error(), nil))
	}
	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "Success Finished your trip", nil))
}

func (handler *DriverHandler) Delete(c echo.Context) error {
	id := c.Param("driver_id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, "id not valid")
	}
	err := handler.driverService.Delete(uint(idConv))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "success delete driver")
}

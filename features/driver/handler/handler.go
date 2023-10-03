package handler

import (
	"fmt"
	"net/http"
	"project-capston/features/driver"
	"project-capston/helper"
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

package handler

import (
	"net/http"
	"project-capston/app/middlewares"
	usernodejs "project-capston/features/UserNodeJs"
	"project-capston/features/unit"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type UnitHandler struct {
	unitHandler unit.UnitServiceInterface
}

func New(handler unit.UnitServiceInterface) *UnitHandler {
	return &UnitHandler{
		unitHandler: handler,
	}
}

func (handler *UnitHandler) Add(c echo.Context) error {
	idEmergencies, _ := middlewares.ExtractTokenUserId(c)
	idVehicle := c.Param("vehicles_id")
	idConv, errConv := strconv.Atoi(idVehicle)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, "id not valid")
	}
	var input UnitRequest
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, "error bind data")
	}
	entity := RequestToEntity(input)
	entity.EmergenciesID = uint(idEmergencies)
	entity.VehicleID = uint(idConv)

	err := handler.unitHandler.Add(entity)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, err.Error())
		} else {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
	}
	return c.JSON(http.StatusCreated, "success create data unit")

}

func (handler *UnitHandler) Delete(c echo.Context) error {
	id := c.Param("unit_id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, "id not valid")
	}
	err := handler.unitHandler.Delete(uint(idConv))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "success delete unit")
}

func (handler *UnitHandler) Edit(c echo.Context) error {
	id := c.Param("unit_id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, "id not valid")
	}
	var input UnitRequest
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, "error bind data")
	}

	Entity := RequestToEntity(input)
	err := handler.unitHandler.Edit(Entity, uint(idConv))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "success update unit")
}

func (handler *UnitHandler) GetById(c echo.Context) error {

	id := c.Param("unit_id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, "id not valid")
	}
	token, errToken := usernodejs.GetTokenHandler(c)
	if errToken != nil {
		return c.JSON(http.StatusUnauthorized, "fail get token")
	}
	data, err := handler.unitHandler.GetById(uint(idConv), token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	response := EntityToResponse(data)
	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get unit by id",
		"data":    response,
	})
}

func (handler *UnitHandler) GetAll(c echo.Context) error {
	var qparams unit.QueryParams
	page := c.QueryParam("page")
	itemsPerPage := c.QueryParam("itemsPerPage")

	if itemsPerPage == "" {
		qparams.IsClassDashboard = false
	} else {
		qparams.IsClassDashboard = true
		itemsConv, errItem := strconv.Atoi(itemsPerPage)
		if errItem != nil {
			return c.JSON(http.StatusBadRequest, "item per page not valid")
		}
		qparams.ItemsPerPage = itemsConv
	}

	if page == "" {
		qparams.Page = 1
	} else {
		pageConv, errPage := strconv.Atoi(page)
		if errPage != nil {
			return c.JSON(http.StatusBadRequest, "page not valid")
		}
		qparams.Page = pageConv
	}
	// name:=c.QueryParam("searchName")
	// qparams.SearchName = name
	token, errToken := usernodejs.GetTokenHandler(c)
	if errToken != nil {
		return c.JSON(http.StatusUnauthorized, "fail get token")
	}

	bol, data, err := handler.unitHandler.GetAll(qparams, token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var response []UnitResponse
	for _, v := range data {
		response = append(response, EntityToResponse(v))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":   "success get all unit",
		"data":      response,
		"next_page": bol,
	})
}

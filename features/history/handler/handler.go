package handler

import (
	"net/http"
	"project-capston/app/middlewares"
	usernodejs "project-capston/features/UserNodeJs"
	"project-capston/features/history"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type HistoryHandler struct {
	historyHandler history.HistoryServiceInterface
}

func New(handler history.HistoryServiceInterface) *HistoryHandler {
	return &HistoryHandler{
		historyHandler: handler,
	}
}

func (handler *HistoryHandler) Add(c echo.Context) error {
	idUnit, _ := middlewares.ExtractTokenUserId(c)
	idDeiver := c.Param("driver_id")
	idConv, errConv := strconv.Atoi(idDeiver)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, "id not valid")
	}
	var input HistoryRequest
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, "error bind data")
	}
	entity := RequestToEntity(input)
	entity.UnitID = uint(idUnit)
	entity.DriverID = uint(idConv)

	err := handler.historyHandler.Add(entity)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, err.Error())
		} else {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
	}
	return c.JSON(http.StatusCreated, "success create data history")

}

func (handler *HistoryHandler) Delete(c echo.Context) error {
	id := c.Param("history_id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, "id not valid")
	}
	err := handler.historyHandler.Delete(uint(idConv))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "success delete history")
}

func (handler *HistoryHandler) Edit(c echo.Context) error {
	id := c.Param("history_id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, "id not valid")
	}
	var input HistoryRequest
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, "error bind data")
	}

	Entity := RequestToEntity(input)
	err := handler.historyHandler.Edit(Entity, uint(idConv))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "success update history")
}

func (handler *HistoryHandler) GetById(c echo.Context) error {

	id := c.Param("history_id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, "id not valid")
	}
	token, errToken := usernodejs.GetTokenHandler(c)
	if errToken != nil {
		return c.JSON(http.StatusUnauthorized, "fail get token")
	}
	data, err := handler.historyHandler.GetById(uint(idConv), token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	response := EntityToResponse(data)
	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get history by id",
		"data":    response,
	})
}

func (handler *HistoryHandler) GetAll(c echo.Context) error {
	var qparams history.QueryParams
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

	bol, data, err := handler.historyHandler.GetAll(qparams, token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var response []HistoryResponse
	for _, v := range data {
		response = append(response, EntityToResponse(v))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":   "success get all history",
		"data":      response,
		"next_page": bol,
	})
}

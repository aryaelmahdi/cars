package controller

import (
	"net/http"
	"project/helper"
	"project/model"

	"github.com/labstack/echo/v4"
)

type TypesController struct {
	model model.TypesModel
}

func (tc *TypesController) Init(tm model.TypesModel) {
	tc.model = tm
}

func (tc *TypesController) InsertTypes() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := model.Types{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}
		res := tc.model.InsertTypes(input)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}
		return c.JSON(http.StatusCreated, helper.SetResponse("success", res))
	}
}

func (tc *TypesController) GetAllTypes() echo.HandlerFunc {
	return func(c echo.Context) error {
		res := tc.model.GetAllTypes()
		return c.JSON(http.StatusOK, helper.SetResponse("success", res))
	}
}

func (tc *TypesController) GetTypeByName() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.Param("name")
		res := tc.model.GetTypeByName(name)
		if res == nil {
			return c.JSON(http.StatusBadRequest, nil)
		}
		return c.JSON(http.StatusOK, helper.SetResponse("success", res))
	}
}

func (tc *TypesController) DeleteTypes() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.Param("name")
		if err := tc.model.DeleteType(name); err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}
		return c.JSON(http.StatusNoContent, nil)
	}
}

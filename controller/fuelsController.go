package controller

import (
	"net/http"
	"project/helper"
	"project/model"

	"github.com/labstack/echo/v4"
)

type FuelController struct {
	model model.FuelModel
}

func (fc *FuelController) Init(fm model.FuelModel) {
	fc.model = fm
}

func (fc *FuelController) GetAllFuel() echo.HandlerFunc {
	return func(c echo.Context) error {
		res := fc.model.GetAllFuel()
		return c.JSON(http.StatusOK, helper.SetResponse("success", res))
	}
}

func (fc *FuelController) GetFuelByName() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.Param("name")
		res := fc.model.GetFuelByName(name)
		if res == nil {
			return c.JSON(http.StatusBadRequest, nil)
		}
		return c.JSON(http.StatusOK, helper.SetResponse("success", res))
	}
}

func (fc *FuelController) DeleteFuel() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.Param("name")
		if err := fc.model.DeleteFuel(name); err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}
		return c.JSON(http.StatusNoContent, nil)
	}
}

func (fc *FuelController) InsertFuel() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := model.Fuel{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}
		res := fc.model.InsertFuel(input)

		if res == nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}
		return c.JSON(http.StatusCreated, helper.SetResponse("success", res))
	}
}

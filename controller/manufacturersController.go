package controller

import (
	"net/http"
	"project/config"
	"project/helper"
	"project/model"

	"github.com/labstack/echo/v4"
)

type ManufacturersController struct {
	config config.Config
	model  model.ManufacturersModel
}

func (mc *ManufacturersController) Init(m model.ManufacturersModel, cfg config.Config) {
	mc.model = m
	mc.config = cfg
}

func (mc *ManufacturersController) GetAllManufacturers() echo.HandlerFunc {
	return func(c echo.Context) error {
		res := mc.model.GetAllManufacturers()
		return c.JSON(http.StatusOK, helper.SetResponse("success", res))
	}
}

func (mc *ManufacturersController) GetManufacturersByName() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.Param("name")
		res := mc.model.GetManufacturersByName(name)
		if res == nil {
			return c.JSON(http.StatusBadRequest, nil)
		}
		return c.JSON(http.StatusOK, helper.SetResponse("success", res))
	}
}

func (mc *ManufacturersController) DeleteManufacturer() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.Param("name")
		if err := mc.model.DeleteManufacturer(name); err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}
		return c.JSON(http.StatusNoContent, nil)
	}
}

func (mc *ManufacturersController) InsertManufacturer() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := model.Manufacturers{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("invalid input", nil))
		}
		res := mc.model.InsertManufacturer(input)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}
		return c.JSON(http.StatusCreated, helper.SetResponse("success", res))
	}
}

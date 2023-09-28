package controller

import (
	"io"
	"net/http"
	"project/helper"
	"project/model"

	"github.com/labstack/echo/v4"
)

type CarsController struct {
	model model.CarModel
}

func (cc *CarsController) Init(cm model.CarModel) {
	cc.model = cm
}

func (cc *CarsController) InsertCars() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := model.Car{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("invalid input", nil))
		}

		file, _, err := c.Request().FormFile("image")
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("failed to read image", nil))
		}

		name := c.FormValue("name")
		if name == "" {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("invalid input", nil))
		}
		make := c.FormValue("make")
		if make == "" {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("invalid input", nil))
		}
		aspiration := c.FormValue("aspiration")
		if aspiration == "" {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("invalid input", nil))
		}
		fuel := c.FormValue("fuel")
		if fuel == "" {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("invalid input", nil))
		}
		transmission := c.FormValue("transmission")
		if transmission == "" {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("invalid input", nil))
		}
		drivetrain := c.FormValue("drivetrain")
		if drivetrain == "" {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("invalid input", nil))
		}
		types := c.FormValue("type")
		if transmission == "" {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("invalid input", nil))
		}

		imageBytes, err := io.ReadAll(file)
		input.Image = imageBytes
		input.Name = name
		input.Make = make
		input.Aspiration = aspiration
		input.Fuel = fuel
		input.Transmission = transmission
		input.Drivetrain = drivetrain
		input.Type = types

		res := cc.model.InsertCar(input)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, helper.SetResponse("something went wrong", nil))
		}

		return c.JSON(http.StatusCreated, helper.SetResponse("success", res))
	}
}

func (cc *CarsController) GetAllCars() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := model.Car{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}
		res := cc.model.GetAllCars()
		return c.JSON(http.StatusOK, helper.SetResponse("success", res))
	}
}

func (cc *CarsController) GetCarByName() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.Param("name")
		res := cc.model.GetCarByName(name)
		if res == nil {
			return c.JSON(http.StatusBadRequest, nil)
		}
		return c.JSON(http.StatusOK, helper.SetResponse("success", res))
	}
}

func (cc *CarsController) DeleteCar() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.Param("name")
		if err := cc.model.DeleteCar(name); err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}
		return c.JSON(http.StatusNoContent, nil)
	}
}

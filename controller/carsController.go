package controller

import (
	"fmt"
	"io"
	"net/http"
	"project/config"
	"project/helper"
	"project/model"

	"github.com/labstack/echo/v4"
)

type CarsController struct {
	model  model.CarModel
	config config.Config
}

func (cc *CarsController) Init(cm model.CarModel, cfg config.Config) {
	cc.model = cm
	cc.config = cfg
}

func (cc *CarsController) InsertCars() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := model.Car{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("invalid input", err))
		}

		file, _, err := c.Request().FormFile("image")
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("failed to read image", err))
		}

		name := c.FormValue("name")
		if name == "" {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("invalid input", err))
		}
		make := c.FormValue("make")
		if make == "" {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("invalid input", err))
		}
		powerplant := c.FormValue("powerplant")
		if powerplant == "" {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("invalid input", err))
		}
		drivetrain := c.FormValue("drivetrain")
		if drivetrain == "" {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("invalid input", err))
		}

		imageBytes, err := io.ReadAll(file)
		input.Image = imageBytes
		input.Name = name
		input.Make = make
		input.Powerplant = powerplant
		input.Drivetrain = drivetrain

		fmt.Println("inserted model :", input.Drivetrain, input.Powerplant, input.Name, input.Make)

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

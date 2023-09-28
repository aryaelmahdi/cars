package controller

import (
	"net/http"
	"project/helper"
	"project/model"

	"github.com/labstack/echo/v4"
)

type DrivetrainController struct {
	model model.DrivetrainModel
}

func (dc *DrivetrainController) Init(dm model.DrivetrainModel) {
	dc.model = dm
}

func (dc *DrivetrainController) GetAllDrivetrain() echo.HandlerFunc {
	return func(c echo.Context) error {
		res := dc.model.GetAllDrivetrain()
		return c.JSON(http.StatusOK, helper.SetResponse("success", res))
	}
}

func (dc *DrivetrainController) GetDrivetrainByName() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.Param("name")
		res := dc.model.GetDrivetrainByName(name)
		if res == nil {
			return c.JSON(http.StatusBadRequest, nil)
		}
		return c.JSON(http.StatusOK, helper.SetResponse("success", res))
	}
}

func (dc *DrivetrainController) DeletedriveTrain() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.Param("name")
		if err := dc.model.DeletedriveTrain(name); err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}
		return c.JSON(http.StatusNoContent, nil)
	}
}

func (dc *DrivetrainController) InsertDrivetrain() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := model.Drivetrain{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}
		res := dc.model.InsertDrivetrain(input)

		if res == nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}
		return c.JSON(http.StatusCreated, helper.SetResponse("success", res))
	}
}

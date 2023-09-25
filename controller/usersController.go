package controller

import (
	"net/http"
	"project/helper"
	"project/model"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	model model.UserModel
}

func (uc *UserController) Init(um model.UserModel) {
	uc.model = um
}

func (uc *UserController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := model.Users{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("invalid input", nil))
		}

		res, err := uc.model.Register(input)

		if err != nil {
			logrus.Error("Failed to register user:", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.SetResponse("something went wrong", nil))
		}

		return c.JSON(http.StatusCreated, helper.SetResponse("success", res))
	}
}

func (uc *UserController) GetAllUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := uc.model.GetAllUsers()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}
		return c.JSON(http.StatusOK, helper.SetResponse("success", res))
	}
}

func (uc *UserController) GetUserByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := c.Param("id")
		id, err := strconv.Atoi(param)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("invalid id", nil))
		}
		res, err := uc.model.GetUserByID(id)

		if res == nil {
			return c.JSON(http.StatusInternalServerError, helper.SetResponse("something went wrong", nil))
		}

		return c.JSON(http.StatusOK, helper.SetResponse("success", res))
	}
}

func (uc *UserController) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := c.Param("id")
		id, err := strconv.Atoi(param)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("invalid input", nil))
		}

		uc.model.DeleteUser(id)
		return c.JSON(http.StatusNoContent, nil)
	}
}

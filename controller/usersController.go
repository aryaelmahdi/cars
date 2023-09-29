package controller

import (
	"fmt"
	"net/http"
	"project/config"
	"project/helper"
	mid "project/middleware"
	"project/model"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	model  model.UserModel
	config config.Config
}

func (uc *UserController) Init(um model.UserModel, cfg config.Config) {
	uc.model = um
	uc.config = cfg
}

func (uc *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = model.Login{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("invalid user input", nil))
		}

		res := uc.model.Login(input.Email, input.Password)

		if res == nil {
			return c.JSON(http.StatusInternalServerError, helper.SetResponse("something went wrong", nil))
		}

		if res.Id == 0 {
			return c.JSON(http.StatusNotFound, helper.SetResponse("data not found", nil))
		}

		jwtToken := mid.GenerateJWT(uc.config.Secret, res.Name, fmt.Sprint(res.Id))

		if jwtToken == nil {
			return c.JSON(http.StatusInternalServerError, helper.SetResponse("cannot process data", nil))
		}

		jwtToken["info"] = res

		return c.JSON(http.StatusOK, helper.SetResponse("success", jwtToken))
	}
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

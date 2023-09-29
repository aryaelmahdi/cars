package controller

import (
	"net/http"
	"project/config"
	"project/helper"
	"project/model"

	"github.com/labstack/echo/v4"
)

type EngineController struct {
	model  model.EngineModel
	config config.Config
}

func (ec *EngineController) Init(em model.EngineModel, cfg config.Config) {
	ec.model = em
	ec.config = cfg
}

func (ec *EngineController) InsertEngine() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := model.Engine{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("invalid input", nil))
		}
		res := ec.model.InsertEngine(input)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, helper.SetResponse("something went wrong", nil))
		}
		return c.JSON(http.StatusCreated, helper.SetResponse("success", res))
	}
}

func (ec *EngineController) GetAllEngines() echo.HandlerFunc {
	return func(c echo.Context) error {
		res := ec.model.GetAllEngine()
		if res == nil {
			return c.JSON(http.StatusInternalServerError, helper.SetResponse("something went wrong", nil))
		}
		return c.JSON(http.StatusOK, helper.SetResponse("success", res))
	}
}

func (ec *EngineController) GetEngineByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		res := ec.model.GetEngineByID(id)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}
		return c.JSON(http.StatusOK, helper.SetResponse("success", res))
	}
}

func (ec *EngineController) DeleteEngine() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		if err := ec.model.DeleteEngine(id); err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}
		return c.JSON(http.StatusNoContent, nil)
	}
}

package main

import (
	"fmt"
	"project/config"
	"project/controller"
	"project/model"
	"project/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	cfg := config.Init()
	db := config.LoadDB(cfg)

	userModel := model.UserModel{}
	userModel.Init(db.Database)
	userController := controller.UserController{}
	userController.Init(userModel)

	aspirtaionModel := model.AspirationsModel{}
	aspirtaionModel.Init(db.Database)
	aspirationController := controller.AspirationsController{}
	aspirationController.Init(model.AspirationsModel(userModel))

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		}))

	if db.Database != nil {
		e.Logger.Info("Database connection established successfully")
	} else {
		e.Logger.Fatal("Failed to establish a database connection")
	}

	routes.UserRoutes(e, userController)
	routes.AspirationRoutes(e, aspirationController)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.ServerPort)).Error())
}

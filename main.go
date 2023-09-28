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

	drivetrainModel := model.DrivetrainModel{}
	drivetrainModel.Init(db.Database)
	drivetrainController := controller.DrivetrainController{}
	drivetrainController.Init(drivetrainModel)

	fuelModel := model.FuelModel{}
	fuelModel.Init(db.Database)
	fuelController := controller.FuelController{}
	fuelController.Init(fuelModel)

	transmissionModel := model.TransmissionModel{}
	transmissionModel.Init(db.Database)
	transmissionController := controller.TransmissionsController{}
	transmissionController.Init(transmissionModel)

	manufacturerModel := model.ManufacturersModel{}
	manufacturerModel.Init(db.Database)
	manufacturerController := controller.ManufacturersController{}
	manufacturerController.Init(manufacturerModel)

	typesModel := model.TypesModel{}
	typesModel.Init(db.Database)
	typesController := controller.TypesController{}
	typesController.Init(typesModel)

	carModel := model.CarModel{}
	carModel.Init(db.Database)
	carController := controller.CarsController{}
	carController.Init(carModel)

	engineModel := model.EngineModel{}
	engineModel.Init(db.Database)
	engineController := controller.EngineController{}
	engineController.Init(engineModel)

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
	routes.DrivertrainRoutes(e, drivetrainController)
	routes.FuelRoutes(e, fuelController)
	routes.TransmissionRoutes(e, transmissionController)
	routes.ManufacturersRoutes(e, manufacturerController)
	routes.TypesRoutes(e, typesController)
	routes.CarRoutes(e, carController)
	routes.EngineRoutes(e, engineController)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.ServerPort)).Error())
}

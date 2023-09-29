package routes

import (
	"project/config"
	"project/controller"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo, uc controller.UserController, config *config.Config) {
	// var user = e.Group("/users")
	// user.POST(c.Register(), )
	e.POST("/Login", uc.Login())
	e.POST("/users", uc.Register())
	// e.GET("/users", uc.GetAllUsers(), echojwt.JWT([]byte(config.Secret)))
	e.GET("/users/:id", uc.GetUserByID(), echojwt.JWT([]byte(config.Secret)))
	e.DELETE("/users/:id", uc.DeleteUser(), echojwt.JWT([]byte(config.Secret)))
}

func AspirationRoutes(e *echo.Echo, ac controller.AspirationsController, config *config.Config) {
	e.POST("/aspirations", ac.InsertAspiration(), echojwt.JWT([]byte(config.Secret)))
	e.DELETE("/aspirations/:id", ac.DeleteAspiration(), echojwt.JWT([]byte(config.Secret)))
	e.GET("/aspirations", ac.GetAllAspirations())
	e.GET("/aspirations/:id", ac.GetAspirationByID())
}

func DrivertrainRoutes(e *echo.Echo, dc controller.DrivetrainController, config *config.Config) {
	e.POST("/drivetrains", dc.InsertDrivetrain(), echojwt.JWT([]byte(config.Secret)))
	e.GET("/drivetrains", dc.GetAllDrivetrain())
	e.GET("/drivetrains/:name", dc.GetDrivetrainByName())
	e.DELETE("/drivetrains/:name", dc.DeletedriveTrain(), echojwt.JWT([]byte(config.Secret)))
}

func FuelRoutes(e *echo.Echo, fc controller.FuelController, config *config.Config) {
	e.POST("/fuels", fc.InsertFuel())
	e.GET("/fuels", fc.GetAllFuel())
	e.GET("/fuels/:name", fc.GetFuelByName())
	e.DELETE("/fuels/:name", fc.DeleteFuel())
}

func TransmissionRoutes(e *echo.Echo, tc controller.TransmissionsController, config *config.Config) {
	e.GET("/transmissions", tc.GetAllTransmissions())
	e.POST("/transmissions", tc.InsertTransmission(), echojwt.JWT([]byte(config.Secret)))
	e.GET("/transmissions/:name", tc.GetTransmissionByName())
	e.DELETE("/transmissions/:name", tc.DeleteTransmission(), echojwt.JWT([]byte(config.Secret)))
}

func ManufacturersRoutes(e *echo.Echo, mc controller.ManufacturersController, config *config.Config) {
	e.GET("/manufacturers", mc.GetAllManufacturers())
	e.GET("/manufacturers/:name", mc.GetManufacturersByName())
	e.POST("/manufacturers", mc.InsertManufacturer(), echojwt.JWT([]byte(config.Secret)))
	e.DELETE("/manufacturers/:name", mc.DeleteManufacturer(), echojwt.JWT([]byte(config.Secret)))
}

func TypesRoutes(e *echo.Echo, tc controller.TypesController, config *config.Config) {
	e.GET("/types", tc.GetAllTypes())
	e.GET("/types/:name", tc.GetTypeByName())
	e.POST("/types", tc.InsertTypes(), echojwt.JWT([]byte(config.Secret)))
	e.DELETE("/types/:name", tc.DeleteTypes(), echojwt.JWT([]byte(config.Secret)))
}

func CarRoutes(e *echo.Echo, cc controller.CarsController, config *config.Config) {
	e.GET("/cars", cc.GetAllCars())
	e.GET("/cars/:name", cc.GetCarByName())
	e.POST("/cars", cc.InsertCars(), echojwt.JWT([]byte(config.Secret)))
	e.DELETE("/cars/:name", cc.DeleteCar(), echojwt.JWT([]byte(config.Secret)))
}

func EngineRoutes(e *echo.Echo, ec controller.EngineController, config *config.Config) {
	e.GET("/engines", ec.GetAllEngines())
	e.GET("/engines/:id", ec.GetEngineByID())
	e.POST("/engines", ec.InsertEngine(), echojwt.JWT([]byte(config.Secret)))
	e.DELETE("engines/:id", ec.DeleteEngine(), echojwt.JWT([]byte(config.Secret)))
}

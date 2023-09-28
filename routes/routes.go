package routes

import (
	"project/controller"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo, uc controller.UserController) {
	// var user = e.Group("/users")
	// user.POST(c.Register(), )
	e.POST("/users", uc.Register())
	e.GET("/users", uc.GetAllUsers())
	e.GET("/users/:id", uc.GetUserByID())
	e.DELETE("/users/:id", uc.DeleteUser())
}

func AspirationRoutes(e *echo.Echo, ac controller.AspirationsController) {
	e.POST("/aspirations", ac.InsertAspiration())
	e.DELETE("/aspirations/:id", ac.DeleteAspiration())
	e.GET("/aspirations", ac.GetAllAspirations())
	e.GET("/aspirations/:id", ac.GetAspirationByID())
}

func DrivertrainRoutes(e *echo.Echo, dc controller.DrivetrainController) {
	e.POST("/drivetrains", dc.InsertDrivetrain())
	e.GET("/drivetrains", dc.GetAllDrivetrain())
	e.GET("/drivetrains/:name", dc.GetDrivetrainByName())
	e.DELETE("/drivetrains/:name", dc.DeletedriveTrain())
}

func FuelRoutes(e *echo.Echo, fc controller.FuelController) {
	e.POST("/fuels", fc.InsertFuel())
	e.GET("/fuels", fc.GetAllFuel())
	e.GET("/fuels/:name", fc.GetFuelByName())
	e.DELETE("/fuels/:name", fc.DeleteFuel())
}

func TransmissionRoutes(e *echo.Echo, tc controller.TransmissionsController) {
	e.GET("/transmissions", tc.GetAllTransmissions())
	e.POST("/transmissions", tc.InsertTransmission())
	e.GET("/transmissions/:name", tc.GetTransmissionByName())
	e.DELETE("/transmissions/:name", tc.DeleteTransmission())
}

func ManufacturersRoutes(e *echo.Echo, mc controller.ManufacturersController) {
	e.GET("/manufacturers", mc.GetAllManufacturers())
	e.GET("/manufacturers/:name", mc.GetManufacturersByName())
	e.POST("/manufacturers", mc.InsertManufacturer())
	e.DELETE("/manufacturers/:name", mc.DeleteManufacturer())
}

func TypesRoutes(e *echo.Echo, tc controller.TypesController) {
	e.GET("/types", tc.GetAllTypes())
	e.GET("/types/:name", tc.GetTypeByName())
	e.POST("/types", tc.InsertTypes())
	e.DELETE("/types/:name", tc.DeleteTypes())
}

func CarRoutes(e *echo.Echo, cc controller.CarsController) {
	e.GET("/cars", cc.GetAllCars())
	e.GET("/cars/:name", cc.GetCarByName())
	e.POST("/cars", cc.InsertCars())
	e.DELETE("/cars/:name", cc.DeleteCar())
}

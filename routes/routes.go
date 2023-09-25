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

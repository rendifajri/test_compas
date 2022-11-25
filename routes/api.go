package routes

import (
	"net/http"

	"test_compas/controllers"

	"github.com/labstack/echo/v4"
)

func Build(e *echo.Echo) {
	var resp = map[string]interface{}{
		"status":  "",
		"message": "",
		"data":    nil,
	}
	echo.NotFoundHandler = func(c echo.Context) error {
		resp["status"] = "not_found"
		resp["message"] = "Route Not Found"
		return c.JSON(http.StatusNotFound, resp)
	}
	echo.MethodNotAllowedHandler = func(c echo.Context) error {
		resp["status"] = "not_allowed"
		resp["message"] = "Route Not Allowed"
		return c.JSON(http.StatusMethodNotAllowed, resp)
	}
	e.GET("/user/index", controllers.UserIndex)
	e.GET("/user/:id", controllers.UserDetail)
	e.POST("/user/:id/point", controllers.UserPoint)
	e.POST("/user/:id/point/minus", controllers.UserPointMinus)
	e.POST("/user/:id/process/status", controllers.UserProcessStatus)
}

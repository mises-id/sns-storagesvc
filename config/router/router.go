package router

import (
	"github.com/labstack/echo"
	"github.com/mises-id/sns-storagesvc/app/ctrl"
)

func SetRouter(e *echo.Echo) {
	//image
	api := &ctrl.ImageCtrl{}
	e.GET("*", api.Handler)
}

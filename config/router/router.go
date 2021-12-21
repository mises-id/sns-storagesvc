package router

import (
	"github.com/labstack/echo"
	"github.com/mises-id/sns-storagesvc/app/ctrl"
)

func SetRouter(e *echo.Echo) {
	//image
	ImageCtrl := &ctrl.ImageCtrl{}
	e.GET("*", ImageCtrl.Handler)
}

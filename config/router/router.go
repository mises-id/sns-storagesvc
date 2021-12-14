package router

import (
	"github.com/labstack/echo"
	"github.com/mises-id/sns-storagesvc/app/services/views/image"
)

func SetRouter(e *echo.Echo) {

	e.GET("*", image.Handler)
}

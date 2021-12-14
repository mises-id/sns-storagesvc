package middleware

import (
	"net/http"

	"github.com/labstack/echo"
)

var ErrorResponseMiddleware = func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}
		return nil
	}
}

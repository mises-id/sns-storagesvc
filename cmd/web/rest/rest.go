package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo"
	"github.com/mises-id/storagesvc/app/services/views/image"
	"github.com/mises-id/storagesvc/config/env"
)

func Start(ctx context.Context) error {
	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
			}
			return nil
		}

	})
	// Route
	e.GET("*", image.Handler)
	go func() {
		if err := e.Start(fmt.Sprintf(":%d", env.Envs.Port)); err != nil {
			log.Fatal(err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	return e.Shutdown(ctx)
}

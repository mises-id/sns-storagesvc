package rest

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo"
	"github.com/mises-id/sns-storagesvc/lib/middleware"

	"github.com/mises-id/sns-storagesvc/config/env"
	"github.com/mises-id/sns-storagesvc/config/router"
)

func Start(ctx context.Context) error {
	e := echo.New()
	e.Use(middleware.ErrorResponseMiddleware)
	// Router
	router.SetRouter(e)
	go func() {
		if err := e.Start(fmt.Sprintf(":%d", env.Envs.WebPort)); err != nil {
			log.Fatal(err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	return e.Shutdown(ctx)
}

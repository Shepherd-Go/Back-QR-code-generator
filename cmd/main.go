package main

import (
	"fmt"
	"log"

	"github.com/andresxlp/qr-system/cmd/providers"
	"github.com/andresxlp/qr-system/config"
	"github.com/andresxlp/qr-system/internal/infra/api/router"
	"github.com/labstack/echo/v4"
)

func main() {

	container := providers.BuildContainer()
	port := config.Environments().Server.Port

	if err := container.Invoke(func(server *echo.Echo, router *router.Router) {
		router.Init()

		server.Logger.Fatal(server.Start(fmt.Sprintf(":%d", port)))

	}); err != nil {
		log.Fatal(err)
	}

}

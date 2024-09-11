package main

import (
	"fmt"
	"log"

	"github.com/andresxlp/qr-system/cmd/providers"
	"github.com/andresxlp/qr-system/config"
	"github.com/andresxlp/qr-system/internal/infra/api/router"
	"github.com/labstack/echo/v4"
)

//	@title			API de Generación y Validación de Códigos QR
//	@version		1.0
//	@description	Esta API permite generar y validar códigos QR para invitaciones a eventos.
//	@BasePath		/api/qr-code
//	@schemas		http

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

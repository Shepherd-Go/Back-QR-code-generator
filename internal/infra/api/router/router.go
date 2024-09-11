package router

import (
	"net/http"

	"github.com/andresxlp/qr-system/internal/infra/api/handler"
	"github.com/andresxlp/qr-system/internal/infra/api/router/groups"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/andresxlp/qr-system/docs"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
)

type Router struct {
	server  *echo.Echo
	qrGroup groups.QR
}

func New(server *echo.Echo, qr groups.QR) *Router {
	return &Router{
		server,
		qr,
	}
}

func (r *Router) Init() {
	r.server.Use(middleware.Recover())

	r.server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, method=${method}, uri=${uri}, latency=${latency_human}, status=${status}\n",
	}))
	r.server.GET("/docs/*", echoSwagger.WrapHandler)

	r.server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}))

	basePath := r.server.Group("/api/qr-code")
	basePath.GET("/health", handler.HealthCheck)

	r.qrGroup.Resource(basePath)
}

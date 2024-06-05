package router

import (
	"net/http"

	"github.com/andresxlp/qr-system/internal/infra/api/handler"
	"github.com/andresxlp/qr-system/internal/infra/api/router/groups"
	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
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
		Format: "time=${time_rfc3339}, method=${method}, uri=${uri}, latency=${latency_human}\n",
	}))

	r.server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}))

	basePath := r.server.Group("/api/qr-code") //customize your basePath
	basePath.GET("/health", handler.HealthCheck)

	r.qrGroup.Resource(basePath)
}

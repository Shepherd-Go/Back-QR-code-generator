package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/andresxlp/qr-system/internal/app"
	"github.com/labstack/echo/v4"
)

type Admin interface {
	OnlyAdmins(next echo.HandlerFunc) echo.HandlerFunc
}

type admin struct {
	adminServices app.Admin
}

func NewAdmin(adminService app.Admin) Admin {
	return &admin{adminService}
}

func (a *admin) OnlyAdmins(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		email := strings.TrimSpace(strings.ToLower(c.Request().Header.Get("email")))
		if email == "" {
			err := errors.New("you must provide an email")
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		_, err := a.adminServices.GetByEmail(ctx, email)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}

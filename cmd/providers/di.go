package providers

import (
	"github.com/andresxlp/qr-system/internal/app"
	"github.com/andresxlp/qr-system/internal/infra/adapters/mongo"
	"github.com/andresxlp/qr-system/internal/infra/adapters/mongo/repo"
	"github.com/andresxlp/qr-system/internal/infra/api/handler"
	"github.com/andresxlp/qr-system/internal/infra/api/middleware"
	"github.com/andresxlp/qr-system/internal/infra/api/router"
	"github.com/andresxlp/qr-system/internal/infra/api/router/groups"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

var Container *dig.Container

func BuildContainer() *dig.Container {

	Container := dig.New()

	_ = Container.Provide(func() *echo.Echo {
		return echo.New()
	})

	_ = Container.Provide(mongo.ConnInstance)

	_ = Container.Provide(middleware.NewAdmin)

	_ = Container.Provide(router.New)

	_ = Container.Provide(groups.NewQr)

	_ = Container.Provide(handler.NewInvitation)

	_ = Container.Provide(app.NewQr)
	_ = Container.Provide(app.NewAdmin)

	_ = Container.Provide(repo.NewQr)
	_ = Container.Provide(repo.NewAdmin)

	return Container
}

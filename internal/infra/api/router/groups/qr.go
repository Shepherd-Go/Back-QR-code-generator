package groups

import (
	"github.com/andresxlp/qr-system/internal/infra/api/handler"
	"github.com/andresxlp/qr-system/internal/infra/api/middleware"
	"github.com/labstack/echo/v4"
)

type QR interface {
	Resource(group *echo.Group)
}

type qr struct {
	qrHandler       handler.Invitation
	adminMiddleware middleware.Admin
}

func NewQr(qrHand handler.Invitation, adminMiddleware middleware.Admin) QR {
	return qr{
		qrHand,
		adminMiddleware,
	}
}

func (groups qr) Resource(c *echo.Group) {
	groupPath := c.Group("")
	groupPath.POST("/generate", groups.qrHandler.CreateInvitation)
	groupPath.GET("/validate/:id", groups.qrHandler.ValidateInvitation)
	groupPath.PUT("/confirm/:id", groups.qrHandler.ConfirmInvitation)
}

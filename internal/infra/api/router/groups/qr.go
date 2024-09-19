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
	qrHandler       handler.QR
	adminMiddleware middleware.Admin
}

func NewQr(qrHand handler.QR, adminMiddleware middleware.Admin) QR {
	return qr{
		qrHand,
		adminMiddleware,
	}
}

func (groups qr) Resource(c *echo.Group) {
	groupPath := c.Group("")
	groupPath.POST("/generate", groups.qrHandler.GenerateQRCode)
	groupPath.GET("/validate/:id", groups.qrHandler.ValidateQRCode)
	groupPath.GET("/lottery", groups.qrHandler.GetGuestFromLoterry)
	groupPath.DELETE("/lottery/delete/:id", groups.qrHandler.DeleteGuestFromLoterry)
	//groupPath.POST("/generate_batch", groups.qrHandler.GenerateQRCodeBatch)
	//groupPath.PUT("/confirm/:id", groups.qrHandler.ConfirmInvitation)
	//groupPath.GET("/count/:email", groups.qrHandler.CountQRCodeUsed)
}

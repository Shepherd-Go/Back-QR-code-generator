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
	//groupPath.POST("/download/:serial", groups.qrHandler.DownloadQRCode)
	//groupPath.POST("/validate/:serial", groups.qrHandler.ValidateQRCode)
	groupPath.GET("/count/:email", groups.qrHandler.CountQRCodeUsed)
}

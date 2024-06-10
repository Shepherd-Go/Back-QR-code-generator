package handler

import (
	"net/http"

	"github.com/andresxlp/qr-system/internal/app"
	"github.com/andresxlp/qr-system/internal/domain/dto"
	"github.com/andresxlp/qr-system/internal/domain/entity"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QR interface {
	GenerateQRCode(c echo.Context) error
	ValidateQRCode(c echo.Context) error
	ConfirmInvitation(c echo.Context) error
	//CountQRCodeUsed(cntx echo.Context) error
}

type qr struct {
	qrService app.QR
}

func NewQr(qrService app.QR) QR {
	return &qr{qrService}
}

func (q *qr) GenerateQRCode(c echo.Context) error {
	ctx := c.Request().Context()

	requestQr := dto.QRManagement{}
	if err := c.Bind(&requestQr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, entity.Error{
			Message: "Error",
			Data:    err.Error(),
		})
	}

	go q.qrService.GenerateQRCodes(ctx, requestQr)

	return c.JSON(http.StatusOK, "QR Codes are being generated")
}

/*func (q *qr) DownloadQRCode(cntx echo.Context) error {
	ctx := context.Background()
	requestQr := dto.QrRequestCommon{}
	if err := cntx.Bind(&requestQr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, entity.Error{
			Message: "Error",
			Data:    err.Error(),
		})
	}

	if err := requestQr.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, entity.Error{
			Message: "Error",
			Data:    err.Error(),
		})
	}

	qrCodeByte, err := q.qrService.DownloadQRCode(ctx, requestQr)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, entity.Error{
			Message: "Error",
			Data:    err.Error(),
		})
	}

	cntx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	return cntx.JSON(http.StatusOK, entity.Success{
		Message: "Success",
		Data:    qrCodeByte,
	})
}*/

func (q *qr) ValidateQRCode(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, entity.Error{Message: err.Error()})
	}

	infoGuest, err := q.qrService.ValidateQRCode(ctx, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, entity.Success{
		Message: "Success",
		Data:    infoGuest,
	})
}

func (q *qr) ConfirmInvitation(c echo.Context) error {

	ctx := c.Request().Context()

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, entity.Error{Message: err.Error()})
	}

	err = q.qrService.ConfirmInvitation(ctx, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, entity.Success{Message: "invitation confirmed successfully"})
}

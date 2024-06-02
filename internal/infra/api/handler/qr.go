package handler

import (
	"context"
	"net/http"

	"github.com/andresxlp/qr-system/internal/app"
	"github.com/andresxlp/qr-system/internal/domain/dto"
	"github.com/andresxlp/qr-system/internal/domain/entity"
	"github.com/labstack/echo/v4"
)

type QR interface {
	GenerateQRCode(cntx echo.Context) error
	DownloadQRCode(cntx echo.Context) error
	ValidateQRCode(cntx echo.Context) error
	CountQRCodeUsed(cntx echo.Context) error
}

type qr struct {
	qrService app.QR
}

func NewQr(qrService app.QR) QR {
	return &qr{qrService}
}

func (q *qr) GenerateQRCode(cntx echo.Context) error {
	ctx := context.Background()

	requestQr := dto.CreateQrRequest{}
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

	go q.qrService.GenerateQRCodes(ctx, requestQr)

	return cntx.JSON(http.StatusOK, "QR Codes are being generated")
}

func (q *qr) DownloadQRCode(cntx echo.Context) error {
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
}

func (q *qr) ValidateQRCode(cntx echo.Context) error {
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

	err := q.qrService.ValidateQRCode(ctx, requestQr)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, entity.Error{
			Message: "Error",
			Data:    err.Error(),
		})
	}

	return cntx.JSON(http.StatusOK, entity.Success{
		Message: "Success",
		Data:    "QR Code Validado Correctamente",
	})
}

func (q *qr) CountQRCodeUsed(cntx echo.Context) error {
	ctx := context.Background()
	email := cntx.Param("email")
	if email == "" {
		return echo.NewHTTPError(http.StatusBadRequest, entity.Error{
			Message: "Error",
			Data:    "email param are required",
		})
	}

	totalQrCodeUsed, err := q.qrService.CountQRCodeUsed(ctx, email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, entity.Error{
			Message: "Error",
			Data:    err.Error(),
		})
	}

	return cntx.JSON(http.StatusOK, entity.Success{
		Message: "Total QR Code Used",
		Data:    totalQrCodeUsed,
	})
}

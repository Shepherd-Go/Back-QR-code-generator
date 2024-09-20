package handler

import (
	"net/http"

	"github.com/andresxlp/qr-system/internal/app"
	"github.com/andresxlp/qr-system/internal/domain/dto"
	"github.com/andresxlp/qr-system/internal/domain/entity"
	"github.com/andresxlp/qr-system/pkg"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QR interface {
	GenerateQRCode(c echo.Context) error
	ValidateQRCode(c echo.Context) error
	GetGuestFromLottery(c echo.Context) error
	GenerateQRCodeBatch(c echo.Context) error
}

type qr struct {
	qrService app.QR
}

func NewQr(qrService app.QR) QR {
	return &qr{qrService}
}

// GenerateQRCode
//
//	@Summary		Generar código QR
//	@Description	Genera un código QR basado en la información proporcionada
//	@Tags			QR
//	@Accept			json
//	@Produce		json
//	@Param			body	body		dto.QRManagement	true	"Datos para generar el código QR"
//	@Success		200		{string}	string				"Los códigos QR se están generando"
//	@Failure		400		{object}	entity.Error
//	@Router			/generate [post]
func (q *qr) GenerateQRCode(c echo.Context) error {
	ctx := c.Request().Context()

	requestQr := dto.QRManagement{}
	if err := c.Bind(&requestQr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, entity.Error{
			Message: "Error",
			Data:    err.Error(),
		})
	}

	q.qrService.GenerateQRCodes(ctx, requestQr)

	return c.JSON(http.StatusOK, "QR Codes are being generated")
}

// GetGuestFromLottery
//
//	@Summary		Obtener invitado de la lotería
//	@Description	Obtiene un invitado aleatorio de la lotería
//	@Tags			QR
//	@Produce		json
//	@Success		200	{object} dto.QRManagement
//	@Router			/lottery [get]
func (q *qr) GetGuestFromLottery(c echo.Context) error {
	ctx := c.Request().Context()

	guest := q.qrService.GetGuestFromLottery(ctx)

	return c.JSON(http.StatusOK, guest)
}

// GenerateQRCodeBatch
//
//	@Summary		Generar códigos QR en lote
//	@Description	Genera múltiples códigos QR basados en la información proporcionada en un archivo CSV
//	@Tags			QR
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			Invitaciones	formData	file	true	"Archivo CSV con los datos de los invitados | Formato CSV (nombre, invitado_por, parentesco, sorteo)"
//	@Success		200				{string}	string	"Los códigos QR se están generando"
//	@Failure		400				{object}	entity.Error
//	@Router			/generate_batch [post]
func (q *qr) GenerateQRCodeBatch(c echo.Context) error {
	ctx := c.Request().Context()

	file, err := c.FormFile("Invitaciones")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, entity.Error{
			Message: "Error",
			Data:    err.Error(),
		})
	}

	var requestQr []dto.QRManagement
	if err = pkg.BindFile(file, &requestQr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, entity.Error{
			Message: "Error",
			Data:    err.Error(),
		})
	}

	for _, guest := range requestQr[1:] {
		q.qrService.GenerateQRCodes(ctx, guest)
	}

	return c.JSON(http.StatusOK, "QR Codes are being generated")
}

// ValidateQRCode
//
//	@Summary		Validar código QR
//	@Description	Valida un código QR basado en el ID proporcionado
//	@Tags			QR
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"ID del código QR"
//	@Success		200	{object}	entity.Success{message=string,data=dto.QRManagement}
//	@Failure		400	{object}	entity.Error
//	@Failure		500	{object}	entity.Error
//	@Router			/validate/{id} [get]
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

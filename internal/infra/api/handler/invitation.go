package handler

import (
	"context"
	"net/http"

	"github.com/andresxlp/qr-system/internal/app"
	"github.com/andresxlp/qr-system/internal/domain/dto"
	"github.com/andresxlp/qr-system/internal/domain/entity"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invitation interface {
	CreateInvitation(c echo.Context) error
	ValidateInvitation(c echo.Context) error
	ConfirmInvitation(c echo.Context) error
}

type invitation struct {
	qrService app.Invitation
}

func NewInvitation(qrService app.Invitation) Invitation {
	return &invitation{qrService}
}

func (q *invitation) CreateInvitation(c echo.Context) error {
	ctx := context.Background()

	guest := dto.Guest{}
	if err := c.Bind(&guest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, entity.Error{
			Message: "Error",
			Data:    err.Error(),
		})
	}

	go q.qrService.GenerateInvitation(ctx, guest)

	return c.JSON(http.StatusOK, "QR Codes are being generated")
}

func (q *invitation) ValidateInvitation(c echo.Context) error {
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

func (q *invitation) ConfirmInvitation(c echo.Context) error {

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

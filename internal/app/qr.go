package app

import (
	"bytes"
	"context"
	"fmt"
	"image/color"
	"image/png"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/andresxlp/qr-system/internal/domain/dto"
	"github.com/andresxlp/qr-system/internal/domain/entity"
	"github.com/andresxlp/qr-system/internal/domain/ports/repo"
	"github.com/charmbracelet/log"
	"github.com/fogleman/gg"
	"github.com/labstack/echo/v4"
	"github.com/skip2/go-qrcode"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/exp/rand"
)

type QR interface {
	GenerateQRCodes(ctx context.Context, request []dto.QRManagement)
	ValidateQRCode(ctx context.Context, id primitive.ObjectID) (dto.QRManagement, error)
	ConfirmInvitation(ctx context.Context, id primitive.ObjectID) error
	GetGuestFromLoterry(ctx context.Context) dto.QRManagement
	DeleteGusteFromLoterry(ctx context.Context, idGuest primitive.ObjectID) error
	//CountQRCodeUsed(ctx context.Context, emailOwner string) (int64, error)
}
type qr struct {
	mongo repo.QR
}

func NewQr(mongo repo.QR) QR {
	return &qr{
		mongo,
	}
}

func (q *qr) GenerateQRCodes(ctx context.Context, request []dto.QRManagement) {

	for _, v := range request {

		id, err := q.mongo.Create(ctx, v)
		if err != nil {
			log.Error("")
			return
		}

		qrImg := q.createQrCode(id)

		fmt.Sprintf("Creating inviation: %v\n", v.Nombre)
		q.createTicketWithQR(qrImg, v.Nombre)

	}

}

func (q *qr) createQrCode(code string) entity.QrImage {
	QRCode, err := qrcode.New(code, qrcode.Medium)
	if err != nil {
		log.Error(err)
	}

	QRCode.DisableBorder = true

	return entity.QrImage{
		Serial:  code,
		ImgFile: QRCode.Image(350),
	}
}

func (q *qr) createTicketWithQR(qrImg entity.QrImage, guestName string) {
	imgTicket, err := gg.LoadPNG("../tmp/V-Neifer.png")
	if err != nil {
		log.Error(err)
		return
	}

	dc := gg.NewContextForImage(imgTicket)
	dc.Clear()
	dc.SetColor(color.RGBA{
		R: 203,
		G: 167,
		B: 122,
		A: 255,
	})
	dc.DrawImage(imgTicket, 0, 0)

	if err = dc.LoadFontFace("../tmp/fonts/higuen_serif.ttf", 55); err != nil {
		panic(err)
	}
	dc.DrawImage(qrImg.ImgFile, 445, 1079)
	dc.DrawStringWrapped(strings.ToUpper(guestName), 220, 1620, 0, 0, 800, 1, 1)
	dc.Clip()

	ticketWithQR := dc.Image()

	buff := new(bytes.Buffer)

	if err = png.Encode(buff, ticketWithQR); err != nil {
		log.Error(err)
	}

	invitationPath := "../tmp/invitations/"
	dir := filepath.Dir(invitationPath)
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			log.Fatalf("an error occurred when trying to create directory: %v", err)
		}
	}

	if err = os.WriteFile(fmt.Sprintf("%s%s.png", invitationPath, strings.Replace(guestName, " ", "-", -1)), buff.Bytes(), 0644); err != nil {
		log.Error(err)
	}

}

func (q *qr) ValidateQRCode(ctx context.Context, id primitive.ObjectID) (dto.QRManagement, error) {

	infoGuest, err := q.mongo.ValidateQrCode(ctx, id)
	if err != nil {
		if err.Error() == "El qr no es una invitación valida" {
			return dto.QRManagement{}, echo.NewHTTPError(http.StatusNotFound, entity.Error{Message: err.Error()})
		}
		log.Errorf(err.Error())
		return dto.QRManagement{}, echo.NewHTTPError(http.StatusInternalServerError, entity.Error{Message: "an internal error has occurred"})
	}

	if infoGuest.Status == "Used" {
		return dto.QRManagement{}, echo.NewHTTPError(http.StatusUnauthorized, entity.Error{Message: "La Invitación ya fue utilizada"})
	}

	infoGuest.ID = ""

	if err = q.ConfirmInvitation(ctx, id); err != nil {
		return dto.QRManagement{}, err
	}

	if infoGuest.Sorteo == "Si" {
		guestsOfLoterry = append(guestsOfLoterry, infoGuest)
		fmt.Println("se agg", infoGuest)
	}

	return infoGuest, nil
}

func (q *qr) ConfirmInvitation(ctx context.Context, id primitive.ObjectID) error {
	if err := q.mongo.ConfirmInvitation(ctx, id); err != nil {
		log.Errorf(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, entity.Error{Message: "an internal error has occurred"})
	}

	return nil
}

var guestsOfLoterry []dto.QRManagement

func (q *qr) GetGuestFromLoterry(ctx context.Context) dto.QRManagement {

	if guestsOfLoterry == nil {
		return dto.QRManagement{}
	}

	randGuest := rand.Intn(len(guestsOfLoterry)-0) + 0

	return guestsOfLoterry[randGuest]
}

func (q *qr) DeleteGusteFromLoterry(ctx context.Context, idGuest primitive.ObjectID) error {

	var guestsOfLoterryAddDelete []dto.QRManagement

	for _, v := range guestsOfLoterry {

		if v.ID == idGuest.String() {
			continue
		}

		guestsOfLoterryAddDelete = append(guestsOfLoterryAddDelete, v)
	}

	guestsOfLoterry = guestsOfLoterryAddDelete

	return nil
}

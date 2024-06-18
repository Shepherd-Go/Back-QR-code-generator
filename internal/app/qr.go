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
	"github.com/andresxlp/qr-system/internal/infra/adapters/mongo/models"
	"github.com/fogleman/gg"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/skip2/go-qrcode"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QR interface {
	GenerateQRCodes(ctx context.Context, request dto.QRManagement)
	ValidateQRCode(ctx context.Context, id primitive.ObjectID) (dto.QRManagement, error)
	ConfirmInvitation(ctx context.Context, id primitive.ObjectID) error
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

func (q *qr) GenerateQRCodes(ctx context.Context, request dto.QRManagement) {

	qrData := models.Qr{
		N_Table:    request.N_Table,
		N_Seat:     request.N_Seat,
		Guest_Name: request.Guest_Name,
		Rol:        request.Rol,
		Status:     "Created",
	}

	id, err := q.mongo.Create(ctx, qrData)
	if err != nil {
		log.Error(err)
	}

	qrImg := q.createQrCode(id)

	q.createTicketWithQR(qrImg, request.Guest_Name)

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
	imgTicket, err := gg.LoadPNG("tmp/invitation_base.png")
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

	if err = dc.LoadFontFace("tmp/fonts/higuen_serif.ttf", 55); err != nil {
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

	invitationPath := "tmp/invitations/"
	dir := filepath.Dir(invitationPath)
	log.Infof("Directorio: %s", dir)
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
		if err.Error() == "this qr-code not exist" {
			return dto.QRManagement{}, echo.NewHTTPError(http.StatusNotFound, entity.Error{Message: "this qr-code not exist"})
		}
		log.Errorf(err.Error())
		return dto.QRManagement{}, echo.NewHTTPError(http.StatusInternalServerError, entity.Error{Message: "an internal error has occurred"})
	}

	return infoGuest, nil
}

func (q *qr) ConfirmInvitation(ctx context.Context, id primitive.ObjectID) error {

	_, err := q.ValidateQRCode(ctx, id)
	if err != nil {
		return err
	}

	if err = q.mongo.ConfirmInvitation(ctx, id); err != nil {
		log.Errorf(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, entity.Error{Message: "an internal error has occurred"})
	}

	return nil
}

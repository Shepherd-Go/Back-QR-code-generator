package app

import (
	"bytes"
	"context"
	"fmt"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/andresxlp/qr-system/internal/domain/dto"
	"github.com/andresxlp/qr-system/internal/domain/entity"
	"github.com/andresxlp/qr-system/internal/domain/ports/repo"
	"github.com/andresxlp/qr-system/internal/infra/adapters/mongo/models"
	"github.com/fogleman/gg"
	"github.com/labstack/gommon/log"
	"github.com/skip2/go-qrcode"
)

type QR interface {
	GenerateQRCodes(ctx context.Context, request dto.CreateQrRequest)
	//DownloadQRCode(ctx context.Context, downloadCode dto.QrRequestCommon) ([]byte, error)
	//ValidateQRCode(ctx context.Context, requestQr dto.QrRequestCommon) error
	CountQRCodeUsed(ctx context.Context, emailOwner string) (int64, error)
}
type qr struct {
	mongo repo.QR
}

func NewQr(mongo repo.QR) QR {
	return &qr{
		mongo,
	}
}

func (q *qr) GenerateQRCodes(ctx context.Context, request dto.CreateQrRequest) {

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

/*func (q *qr) DownloadQRCode(ctx context.Context, downloadCode dto.QrRequestCommon) ([]byte, error) {
	qrDB, err := q.mongo.GetQrCode(ctx, models.Qr{Serial: downloadCode.Serial /*, Pin: downloadCode.Pin})
	if err != nil {
		return nil, err
	}

	return qrDB.ImgBytes, nil
}

func (q *qr) ValidateQRCode(ctx context.Context, requestQr dto.QrRequestCommon) error {
	err := q.mongo.ValidateQrCode(ctx, models.Qr{Serial: requestQr.Serial /*, Pin: requestQr.Pin})
	if err != nil {
		return err
	}
	return err
}*/

func (q *qr) CountQRCodeUsed(ctx context.Context, emailOwner string) (int64, error) {
	totalQRUsed, err := q.mongo.CountQRCodeUsed(ctx, emailOwner)
	if err != nil {
		return 0, err
	}

	return totalQRUsed, err
}

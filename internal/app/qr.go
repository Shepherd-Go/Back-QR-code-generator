package app

import (
	"bytes"
	"context"
	"fmt"
	"image/color"
	"image/png"
	"os"
	"strings"
	"time"

	"github.com/andresxlp/qr-system/config"
	"github.com/andresxlp/qr-system/internal/domain/dto"
	"github.com/andresxlp/qr-system/internal/domain/entity"
	"github.com/andresxlp/qr-system/internal/domain/ports/repo"
	"github.com/andresxlp/qr-system/internal/infra/adapters/mongo/models"
	"github.com/fogleman/gg"
	"github.com/labstack/gommon/log"
	uuid "github.com/satori/go.uuid"
	"github.com/skip2/go-qrcode"
)

type QR interface {
	GenerateQRCodes(ctx context.Context, request dto.CreateQrRequest)
	DownloadQRCode(ctx context.Context, downloadCode dto.QrRequestCommon) ([]byte, error)
	ValidateQRCode(ctx context.Context, requestQr dto.QrRequestCommon) error
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
	for i := 0; i < request.TotalQR; i++ {
		id := uuid.NewV4()
		code := fmt.Sprintf("%s", id)

		qrImg := q.createQrCode(code)

		q.createTicketWithQR(qrImg, request.Zone, i)

		qrData := models.Qr{
			Serial: code,
			Status: "Created",
			//ImgBytes:  q.createQrCode(code),
			CreatedBy: request.Email,
			CreatedAt: time.Now().Local(),
		}

		if err := q.mongo.Create(ctx, qrData); err != nil {
			log.Error(err)
		}

	}
}

func (q *qr) createQrCode(code string) entity.QrImage {
	qrCode := fmt.Sprintf("%s/validate/%s", config.Environments().InternalPrivatePath, code)

	qrByte, err := qrcode.Encode(qrCode, qrcode.Medium, 235)
	if err != nil {
		log.Error(err)
	}

	qrImg := entity.QrImage{
		Serial:   code,
		PathName: fmt.Sprintf("tmp/qr-code-%s.png", code),
	}

	img, err := png.Decode(bytes.NewBuffer(qrByte))
	if err != nil {
		log.Errorf("an error occurred when try decode img: %v", err)
	}

	file, err := os.Create(qrImg.PathName)
	if err != nil {
		log.Errorf("an error occurred when try create file: %v", err)
	}
	defer file.Close()

	if err = png.Encode(file, img); err != nil {
		log.Errorf("an error occurred when try Encode file: %v", err)
	}

	qrImg.ImgFile, err = gg.LoadPNG(qrImg.PathName)
	if err != nil {
		log.Errorf("an error occurred when try load qr img", err)
	}

	return qrImg
}

func (q *qr) createTicketWithQR(qrImg entity.QrImage, zone string, i int) {
	imgTicket, err := gg.LoadPNG("tmp/ticket.png")
	if err != nil {
		log.Error(err)
	}

	dc := gg.NewContextForImage(imgTicket)
	dc.Clear()
	dc.SetColor(color.Black)
	dc.DrawImage(imgTicket, 0, 0)

	if err = dc.LoadFontFace("tmp/fonts/impact.ttf", 48); err != nil {
		panic(err)
	}
	dc.DrawImage(qrImg.ImgFile, 75, 550)
	dc.DrawString(strings.ToUpper(zone), 75, 830)
	if err = dc.LoadFontFace("tmp/fonts/impact.ttf", 24); err != nil {
		panic(err)
	}
	dc.DrawString(fmt.Sprintf("N° %04d", i+1), 75, 870)
	dc.Clip()

	ticketWithQR := dc.Image()

	buff := new(bytes.Buffer)

	if err = png.Encode(buff, ticketWithQR); err != nil {
		log.Error(err)
	}

	if err = os.WriteFile(fmt.Sprintf("tmp/tickets/ticket-%s.png", qrImg.Serial), buff.Bytes(), 0644); err != nil {
		log.Error(err)
	}

	os.Remove(qrImg.PathName)
}

func (q *qr) DownloadQRCode(ctx context.Context, downloadCode dto.QrRequestCommon) ([]byte, error) {
	qrDB, err := q.mongo.GetQrCode(ctx, models.Qr{Serial: downloadCode.Serial /*, Pin: downloadCode.Pin*/})
	if err != nil {
		return nil, err
	}

	return qrDB.ImgBytes, nil
}

func (q *qr) ValidateQRCode(ctx context.Context, requestQr dto.QrRequestCommon) error {
	err := q.mongo.ValidateQrCode(ctx, models.Qr{Serial: requestQr.Serial /*, Pin: requestQr.Pin*/})
	if err != nil {
		return err
	}
	return err
}

func (q *qr) CountQRCodeUsed(ctx context.Context, emailOwner string) (int64, error) {
	totalQRUsed, err := q.mongo.CountQRCodeUsed(ctx, emailOwner)
	if err != nil {
		return 0, err
	}

	return totalQRUsed, err
}

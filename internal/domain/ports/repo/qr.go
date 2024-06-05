package repo

import (
	"context"

	"github.com/andresxlp/qr-system/internal/infra/adapters/mongo/models"
)

type QR interface {
	Create(ctx context.Context, qr models.Qr) (string, error)
	//GetQrCode(ctx context.Context, qr models.Qr) (models.Qr, error)
	//ValidateQrCode(ctx context.Context, requestQr models.Qr) error
	CountQRCodeUsed(ctx context.Context, emailOwner string) (int64, error)
}

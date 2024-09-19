package repo

import (
	"context"

	"github.com/andresxlp/qr-system/internal/domain/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QR interface {
	Create(ctx context.Context, qr dto.QRManagement) (string, error)
	ValidateQrCode(ctx context.Context, id primitive.ObjectID) (dto.QRManagement, error)
	ConfirmInvitation(ctx context.Context, id primitive.ObjectID) error
	UpdateGuestFromSorteo(ctx context.Context, name string) error
	//CountQRCodeUsed(ctx context.Context, emailOwner string) (int64, error)
}

package repo

import (
	"context"

	"github.com/andresxlp/qr-system/internal/domain/dto"
	"github.com/andresxlp/qr-system/internal/infra/adapters/mongo/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QR interface {
	Create(ctx context.Context, qr models.Invitados) (string, error)
	ValidateQrCode(ctx context.Context, id primitive.ObjectID) (dto.QRManagement, error)
	ConfirmInvitation(ctx context.Context, id primitive.ObjectID) error
	//CountQRCodeUsed(ctx context.Context, emailOwner string) (int64, error)
}

package repo

import (
	"context"

	"github.com/andresxlp/qr-system/internal/infra/adapters/mongo/models"
)

type Admin interface {
	GetByEmail(ctx context.Context, email string) (models.Admin, error)
}

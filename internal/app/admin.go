package app

import (
	"context"

	"github.com/andresxlp/qr-system/internal/domain/ports/repo"
	"github.com/andresxlp/qr-system/internal/infra/adapters/mongo/models"
)

type Admin interface {
	GetByEmail(ctx context.Context, email string) (models.Admin, error)
}

type admin struct {
	mongo repo.Admin
}

func NewAdmin(mongo repo.Admin) Admin {
	return &admin{mongo}
}

func (a *admin) GetByEmail(ctx context.Context, email string) (models.Admin, error) {
	emailFromDB, err := a.mongo.GetByEmail(ctx, email)
	if err != nil {
		return models.Admin{}, err
	}

	return emailFromDB, nil
}

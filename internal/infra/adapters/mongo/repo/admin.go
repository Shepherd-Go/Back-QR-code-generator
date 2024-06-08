package repo

import (
	"context"
	"fmt"

	"github.com/andresxlp/qr-system/internal/domain/ports/repo"
	"github.com/andresxlp/qr-system/internal/infra/adapters/mongo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type admin struct {
	dbClient models.DBClientWrite
}

func NewAdmin(dbCliente models.DBClientWrite) repo.Admin {
	return admin{dbCliente}
}

func (a admin) GetByEmail(ctx context.Context, email string) (models.Admin, error) {
	db := a.dbClient.Collection("admins")

	var adminFromDB models.Admin

	filter := bson.D{{"email", email}}

	if err := db.FindOne(ctx, filter).Decode(&adminFromDB); err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Admin{}, fmt.Errorf("admin was not found with the email [%s]", email)
		}
		return models.Admin{}, err
	}

	return adminFromDB, nil
}

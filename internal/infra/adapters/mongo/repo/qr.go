package repo

import (
	"context"

	"github.com/andresxlp/qr-system/internal/domain/ports/repo"
	"github.com/andresxlp/qr-system/internal/infra/adapters/mongo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type qr struct {
	dbClient models.DBClientWrite
}

func NewQr(dbClient models.DBClientWrite) repo.QR {
	return &qr{dbClient}
}

func (q qr) Create(ctx context.Context, qr models.Qr) (string, error) {
	db := q.dbClient.Database("qr-code").Collection("qr")
	objectID, err := db.InsertOne(ctx, qr)

	id := objectID.InsertedID.(primitive.ObjectID).Hex()

	return id, err
}

/*func (q qr) GetQrCode(ctx context.Context, requestQr models.Qr) (models.Qr, error) {
	db := q.dbClient.Database("qr-code").Collection("qr")
	filter := bson.D{{"serial", requestQr.Serial}}

	var qrFromDb models.Qr
	if err := db.FindOne(ctx, filter).Decode(&qrFromDb); err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Qr{}, fmt.Errorf("this qr code not exist")
		}
		return models.Qr{}, err
	}

	if qrFromDb.Downloaded {
		return models.Qr{UpdatedAt: requestQr.UpdatedAt}, fmt.Errorf("this QR code has already been downloaded and assigned")
	}

	update := bson.D{{"$set", bson.D{{"status", "Assigned"}, {"downloaded", true},
		{"updated_at", time.Now().Local()}, {"pin", requestQr.Pin}}}}

	if _, err := db.UpdateOne(ctx, filter, update); err != nil {
		return models.Qr{}, err
	}

	return qrFromDb, nil
}

func (q qr) ValidateQrCode(ctx context.Context, requestQr models.Qr) error {
	db := q.dbClient.Database("qr-code").Collection("qr")
	filter := bson.D{{"serial", requestQr.Serial}}

	var qrFromDb models.Qr
	if err := db.FindOne(ctx, filter).Decode(&qrFromDb); err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("this qr-code not exist")
		}
		return err
	}

	//if !qrFromDb.CheckPin(requestQr.Pin) {
	//	return fmt.Errorf("the pin does not match the qr-code")
	//}

	if qrFromDb.Status == "Used" {
		return fmt.Errorf("this qr has already been used")
	}

	update := bson.D{{"$set", bson.D{{"status", "Used"}, {"used_at", time.Now().Local()}}}}

	if _, err := db.UpdateOne(ctx, filter, update); err != nil {
		return err
	}

	return nil
}*/

func (q qr) CountQRCodeUsed(ctx context.Context, emailOwner string) (int64, error) {
	db := q.dbClient.Database("qr-code").Collection("qr")
	filter := bson.D{{"created_by", emailOwner}, {"status", "Used"}}
	totalQRCodeUsed, err := db.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}

	return totalQRCodeUsed, nil
}

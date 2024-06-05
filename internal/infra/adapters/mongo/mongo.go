package mongo

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/andresxlp/qr-system/config"
	"github.com/andresxlp/qr-system/internal/infra/adapters/mongo/models"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	once          sync.Once
	instanceWrite models.DBClientWrite
)

func ConnInstance() models.DBClientWrite {
	once.Do(func() {
		instanceWrite = getConnection()
	})

	return instanceWrite
}

func getConnection() models.DBClientWrite {
	return models.DBClientWrite{Client: generateClient()}
}

func generateClient() *mongo.Client {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctxTimeout, options.Client().ApplyURI(config.Environments().Databse.MongoDBConnectionWrite))
	if err != nil {
		panic(fmt.Sprintf("mongoDB error in client configuration: %s", err.Error()))
	}

	if err = client.Ping(ctxTimeout, readpref.Primary()); err != nil {
		panic(fmt.Sprintf("mongoDB error in client connection: %s", err.Error()))
	}

	client.Database("qr-code")

	log.Info("Database Write Connection Successfully")

	return client
}

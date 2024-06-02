package models

import "go.mongodb.org/mongo-driver/mongo"

type DBClientWrite struct {
	*mongo.Client
}

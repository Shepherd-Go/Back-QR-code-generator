package models

import (
	"time"
)

type Qr struct {
	Serial     string    `json:"serial"`
	Status     string    `json:"status"`
	Downloaded bool      `json:"downloaded"`
	Pin        string    `json:"pin"`
	ImgBytes   []byte    `json:"img_bytes" bson:"img_bytes"`
	CreatedBy  string    `json:"created_by" bson:"created_by"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" bson:"updated_at"`
	UsedAt     time.Time `json:"used_at" bson:"used_at"`
}

func (q *Qr) CheckPin(pin string) bool {
	return q.Pin == pin
}

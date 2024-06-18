package models

import (
	"github.com/andresxlp/qr-system/internal/domain/dto"
)

type Qr struct {
	N_Table    int    `json:"n_table" bson:"n_table"`
	N_Seat     int    `json:"n_seat" bson:"n_seat"`
	Guest_Name string `json:"guest_name" bson:"guest_name"`
	Rol        string `json:"rol" bson:"rol"`
	Status     string `json:"status" bson:"status"`
}

func (qr *Qr) ToDomainDTO() dto.QRManagement {
	return dto.QRManagement{
		N_Table:    qr.N_Table,
		N_Seat:     qr.N_Seat,
		Guest_Name: qr.Guest_Name,
		Rol:        qr.Rol,
		Status:     qr.Status,
	}
}

package dto

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type CreateQrRequest struct {
	N_Table    int    `json:"n_table"`
	N_Seat     int    `json:"n_seat"`
	Guest_Name string `json:"guest_name"`
	Rol        string `json:"rol"`
}

type QrRequestCommon struct {
	Serial string `param:"serial" validate:"required"`
	//Pin    string `json:"pin" validate:"required,numeric,gt=0000,lt=9999,len=4"`
}

func (q *CreateQrRequest) Validate() error {
	return validate.Struct(q)
}

func (q *QrRequestCommon) Validate() error {
	return validate.Struct(q)
}

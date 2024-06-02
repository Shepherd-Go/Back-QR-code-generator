package dto

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type CreateQrRequest struct {
	TotalQR int    `json:"total_qr" validate:"min=1"`
	Email   string `json:"email" validate:"required,email"`
	Zone    string `json:"zone" validate:"required"`
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

package models

import (
	"github.com/andresxlp/qr-system/internal/domain/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invitados struct {
	ID          *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Nombre      string              `json:"nombre" bson:"nombre"`
	InvitadoPor string              `json:"invitado_por" bson:"invitado_por"`
	Parentesco  string              `json:"parentesco" bson:"parentesco"`
	Sorteo      string              `json:"sorteo" bson:"sorteo"`
	Creado      string              `json:"creado" bson:"creado"`
	Entregado   string              `json:"entregado" bson:"entregado"`
	Status      string              `json:"status" bson:"status"`
}

func (qr *Invitados) ToDomainDTO() dto.QRManagement {
	return dto.QRManagement{
		ID:          qr.ID.String(),
		Nombre:      qr.Nombre,
		InvitadoPor: qr.InvitadoPor,
		Parentesco:  qr.Parentesco,
		Sorteo:      qr.Sorteo,
		Creado:      qr.Creado,
		Entregado:   qr.Entregado,
		Status:      qr.Status,
	}
}

package dto

type QRManagement struct {
	ID          string `json:"id,omitempty" csv:"omitempty" swaggerignore:"true"`
	Nombre      string `json:"nombre" csv:"nombre" example:"Pedro Pacheco"`
	InvitadoPor string `json:"invitado_por" csv:"invitado_por" example:"Neifer"`
	Parentesco  string `json:"parentesco" csv:"parentesco" example:"Familia"`
	Sorteo      string `json:"sorteo" csv:"omitempty" example:"Si"`
	Creado      string `json:"creado" csv:"omitempty" example:""`
	Entregado   string `json:"entregado" csv:"omitempty" example:""`
	Status      string `json:"status" csv:"omitempty" example:"Created"`
}

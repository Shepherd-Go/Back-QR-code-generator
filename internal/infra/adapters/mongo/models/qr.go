package models

type Qr struct {
	N_Table    int    `json:"n_table" bson:"n_table"`
	N_Seat     int    `json:"n_seat" bson:"n_seat"`
	Guest_Name string `json:"guest_name" bson:"guest_name"`
	Rol        string `json:"rol"`
	Status     string `json:"status" bson:"status"`
}

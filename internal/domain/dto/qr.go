package dto

type QRManagement struct {
	ID         string `json:"id,omitempty"`
	N_Table    int    `json:"n_table"`
	N_Seat     int    `json:"n_seat"`
	Guest_Name string `json:"guest_name"`
	Rol        string `json:"rol"`
	Status     string `json:"status"`
}

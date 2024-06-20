package dto

type Guest struct {
	Guest_Name string `json:"guest_name"`
	N_Table    int    `json:"n_table"`
	N_Seat     int    `json:"n_seat"`
	Rol        string `json:"rol"`
	Lottery    *bool  `json:"lottery"`
	Status     string `json:"status"`
}

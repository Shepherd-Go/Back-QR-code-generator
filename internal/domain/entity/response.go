package entity

type Success struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Error struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

package models

type Equipment struct {
	Id          int     `json:"id"`
	Type        string  `json:"type"`
	IsAvailable bool    `json:"is_available"`
	Price       float64 `json:"price"`
}

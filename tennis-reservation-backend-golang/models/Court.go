package models

type Court struct {
	Id          int  `json:"id"`
	IsAvailable bool `json:"is_available"`
}

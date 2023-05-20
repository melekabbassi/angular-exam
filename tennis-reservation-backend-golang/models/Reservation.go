package models

type Reservation struct {
	Id        int         `json:"id"`
	User      User        `json:"user"`
	Court     Court       `json:"court"`
	Type      string      `json:"type"`
	Date      string      `json:"date"`
	Hour      int         `json:"hour"`
	Duration  int         `json:"duration"`
	Equipment []Equipment `json:"equipment"`
	Services  []Service   `json:"services"`
}

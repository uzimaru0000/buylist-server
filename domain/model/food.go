package model

type Food struct {
	Name     string `json:"name"`
	ImageURL string `json:"imageURL"`
	Exp      int64  `json:"exp"`
	Amount   int    `json:"amount"`
}

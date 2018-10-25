package model

import "time"

type BuyList struct {
	ID          string    `json:"id"`
	Ingredients []string  `json:"ingredients"`
	CreatedAt   time.Time `json:"-"`
}

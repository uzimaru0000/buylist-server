package model

import "time"

type BuyList struct {
	ID          string            `json:"id"`
	Ingredients map[string]string `json:"ingredients"`
	CreatedAt   time.Time         `json:"-"`
}

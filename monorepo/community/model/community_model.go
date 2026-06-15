package model

import "time"

type Community struct {
	Id          uint64    `json:"id,omitempty"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"imageUrl"`
	CreatedAt   time.Time `json:"createdAt"`
	Members     uint64    `json:"members"`
	CreatedBy   uint64    `json:"createdBy,omitempty"`
}

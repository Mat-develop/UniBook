package model

import "time"

type Tag struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name"`
	CreatedBy uint64    `json:"createdBy,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
}

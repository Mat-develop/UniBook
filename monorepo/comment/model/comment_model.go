package model

import (
	"errors"
	"strings"
	"time"
)

type Comment struct {
	ID        uint64    `json:"id"`
	PostID    uint64    `json:"postId,omitempty"`
	UserID    uint64    `json:"userId,omitempty"`
	UserNick  string    `json:"userNick"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
}

type CommentDTO struct {
	Body string `json:"body"`
}

func (c *CommentDTO) Prepare() error {
	c.Body = strings.TrimSpace(c.Body)
	if c.Body == "" {
		return errors.New("comment body is required")
	}
	return nil
}

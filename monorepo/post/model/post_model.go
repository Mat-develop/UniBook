package model

import (
	"errors"
	"strings"
	"time"
)

type Post struct {
	ID            uint64    `json:"id,omitempty"`
	CommunityId   int32     `json:"communityId"`
	CommunityName string    `json:"communityName"`
	UserId        uint64    `json:"UserId,omitempty"`
	UserNick      string    `json:"userNick,omitempty"`
	Title         string    `json:"title,omitempty"`
	Body          string    `json:"body,omitempty"`
	ImageUrl      string    `json:"imageUrl,omitempty"`
	Likes         int32     `json:"likes"`
	CreatedAt     time.Time `json:"createdAt"`
}

type PostDTO struct {
	CommunityId int32  `json:"communityId"`
	Title       string `json:"title"`
	Body        string `json:"body"`
}

func (p *PostDTO) Prepare() error {
	if err := p.validate(); err != nil {
		return err
	}

	p.format()
	return nil
}

func (p *PostDTO) validate() error {
	if p.Title == "" {
		return errors.New("the title is empty")
	}

	if p.Body == "" {
		return errors.New("the body is empty")
	}

	return nil
}

func (p *PostDTO) format() {
	p.Title = strings.TrimSpace(p.Title)
	p.Body = strings.TrimSpace(p.Body)
}

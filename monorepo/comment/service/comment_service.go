package service

import (
	"fmt"
	"v1/comment/model"
	"v1/comment/repository"
)

type CommentService interface {
	CreateComment(postID, userID uint64, dto model.CommentDTO) error
	GetComments(postID uint64) ([]model.Comment, error)
}

type commentService struct {
	repo repository.CommentRepository
}

func NewCommentService(repo repository.CommentRepository) CommentService {
	return &commentService{repo: repo}
}

func (s *commentService) CreateComment(postID, userID uint64, dto model.CommentDTO) error {
	if err := dto.Prepare(); err != nil {
		return fmt.Errorf("invalid comment: %w", err)
	}
	return s.repo.Create(postID, userID, dto.Body)
}

func (s *commentService) GetComments(postID uint64) ([]model.Comment, error) {
	return s.repo.FindByPost(postID)
}

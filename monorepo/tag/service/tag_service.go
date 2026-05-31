package service

import (
	"errors"
	"strings"
	"v1/tag/model"
	"v1/tag/repository"
)

type TagService interface {
	Create(name string, userID uint64) (uint64, error)
	ListAll() ([]model.Tag, error)
	Search(name string) ([]model.Tag, error)
	FindOrCreate(name string, userID uint64) (uint64, error)
}

type tagService struct {
	repo repository.TagRepository
}

func NewTagService(repo repository.TagRepository) TagService {
	return &tagService{repo: repo}
}

func (s *tagService) Create(name string, userID uint64) (uint64, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return 0, errors.New("tag name is required")
	}
	return s.repo.Create(name, userID)
}

func (s *tagService) ListAll() ([]model.Tag, error) {
	return s.repo.FindAll()
}

func (s *tagService) Search(name string) ([]model.Tag, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return s.repo.FindAll()
	}
	return s.repo.Search(name)
}

func (s *tagService) FindOrCreate(name string, userID uint64) (uint64, error) {
	return s.repo.FindOrCreate(name, userID)
}

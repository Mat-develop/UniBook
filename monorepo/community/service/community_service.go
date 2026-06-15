package service

import (
	"errors"
	"strings"

	"v1/community/model"
	"v1/community/repository"
	userModel "v1/users/model"
)

type CommunityService interface {
	Create(community model.Community) (uint64, error)
	GetCommunityByID(id uint64) (model.Community, error)
	GetCommunityByName(name string) (model.Community, error)
	ListCommunities() ([]model.Community, error)
	GetCreatedCommunities(userID uint64) ([]model.Community, error)
	Search(q string) ([]model.Community, error)
	Delete(id uint64) error

	Follow(userID, communityID uint64) error
	Unfollow(userID, communityID uint64) error
	GetFollowers(communityID uint64) ([]userModel.User, error)
	GetJoinedCommunities(userID uint64) ([]uint64, error)
}

type communityService struct {
	repo repository.CommunityRepository
}

func NewCommunityService(repo repository.CommunityRepository) CommunityService {
	return &communityService{repo: repo}
}

func (s *communityService) Create(community model.Community) (uint64, error) {
	community.Name = strings.TrimSpace(community.Name)
	if community.Name == "" {
		return 0, errors.New("name required")
	}

	if existing, err := s.repo.FindByName(community.Name); err == nil && existing.Id != 0 {
		return 0, errors.New("community already exists")
	}
	return s.repo.Create(community)
}

func (s *communityService) GetCommunityByID(id uint64) (model.Community, error) {
	return s.repo.FindByID(id)
}

func (s *communityService) GetCommunityByName(name string) (model.Community, error) {
	return s.repo.FindByName(name)
}

func (s *communityService) ListCommunities() ([]model.Community, error) {
	return s.repo.FindAll()
}

func (s *communityService) Search(q string) ([]model.Community, error) {
	return s.repo.Search(q)
}

func (s *communityService) Delete(id uint64) error {
	return s.repo.Delete(id)
}

func (s *communityService) Follow(userID, communityID uint64) error {
	return s.repo.Follow(userID, communityID)
}

func (s *communityService) Unfollow(userID, communityID uint64) error {
	return s.repo.Unfollow(userID, communityID)
}

func (s *communityService) GetFollowers(communityID uint64) ([]userModel.User, error) {
	return s.repo.FindFollowers(communityID)
}

func (s *communityService) GetJoinedCommunities(userID uint64) ([]uint64, error) {
	return s.repo.FindJoinedByUser(userID)
}

func (s *communityService) GetCreatedCommunities(userID uint64) ([]model.Community, error) {
	return s.repo.FindByCreator(userID)
}

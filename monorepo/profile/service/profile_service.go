package service

import (
	"v1/profile/model"
	"v1/profile/repository"
)

type ProfileService interface {
	GetProfile(userID uint64) (model.UserProfile, error)
	AddEducation(userID uint64, e model.Education) (uint64, error)
	UpdateEducation(userID, id uint64, e model.Education) error
	DeleteEducation(userID, id uint64) error
	AddProject(userID uint64, p model.Project) (uint64, error)
	UpdateProject(userID, id uint64, p model.Project) error
	DeleteProject(userID, id uint64) error
	AddCourse(userID uint64, c model.Course) (uint64, error)
	UpdateCourse(userID, id uint64, c model.Course) error
	DeleteCourse(userID, id uint64) error
}

type profileService struct {
	repo repository.ProfileRepository
}

func NewProfileService(repo repository.ProfileRepository) ProfileService {
	return &profileService{repo: repo}
}

func (s *profileService) GetProfile(userID uint64) (model.UserProfile, error) {
	return s.repo.GetProfile(userID)
}

func (s *profileService) AddEducation(userID uint64, e model.Education) (uint64, error) {
	return s.repo.AddEducation(userID, e)
}

func (s *profileService) UpdateEducation(userID, id uint64, e model.Education) error {
	return s.repo.UpdateEducation(userID, id, e)
}

func (s *profileService) DeleteEducation(userID, id uint64) error {
	return s.repo.DeleteEducation(userID, id)
}

func (s *profileService) AddProject(userID uint64, p model.Project) (uint64, error) {
	return s.repo.AddProject(userID, p)
}

func (s *profileService) UpdateProject(userID, id uint64, p model.Project) error {
	return s.repo.UpdateProject(userID, id, p)
}

func (s *profileService) DeleteProject(userID, id uint64) error {
	return s.repo.DeleteProject(userID, id)
}

func (s *profileService) AddCourse(userID uint64, c model.Course) (uint64, error) {
	return s.repo.AddCourse(userID, c)
}

func (s *profileService) UpdateCourse(userID, id uint64, c model.Course) error {
	return s.repo.UpdateCourse(userID, id, c)
}

func (s *profileService) DeleteCourse(userID, id uint64) error {
	return s.repo.DeleteCourse(userID, id)
}

package service

import (
	"encoding/json"
	"errors"
	"v1/users/model"
	"v1/users/repository"
	"v1/util/authentication"
)

type UserService interface {
	Create(user *model.User) (uint64, error)
	Get(nameOrNick string) ([]model.User, error)
	Update(userID uint64, userModel *model.User, tokenID uint64) error
	UpdatePassword(requestBody []byte, userID uint64, userToken uint64) error
	UpdateImage(userID uint64, tokenID uint64, imageData string) error
	Follow(userId uint64, followerID uint64, follow bool) error
	Delete(userID uint64, tokenID uint64) error
	GetFollowers(userID uint64, tokenID uint64, my bool) ([]model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Create(user *model.User) (uint64, error) {
	if err := user.Prepare("register"); err != nil {
		return 0, err
	}

	return s.repo.Create(*user)
}

func (s *userService) Get(nameOrNick string) ([]model.User, error) {
	if nameOrNick != "" {
		return s.repo.FindUserByName(nameOrNick)
	}

	return nil, errors.New("name or nick search null, nothing to find")
}

func (s *userService) Update(userID uint64, updated *model.User, tokenID uint64) error {
	if userID != tokenID {
		return errors.New("account doesn't match")
	}

	if err := updated.Prepare("edit"); err != nil {
		return err
	}

	return s.repo.Update(userID, *updated)
}

func (s *userService) Delete(userID uint64, tokenID uint64) error {
	if userID != tokenID {
		return errors.New("account doesn't match")
	}

	return s.repo.Delete(userID)
}

func (s *userService) Follow(userID, followerID uint64, follow bool) error {
	if userID == followerID {
		return errors.New("same ID, not possible")
	}

	if follow {
		return s.repo.Follow(userID, followerID)
	}

	return s.repo.Unfollow(userID, followerID)
}

func (s *userService) GetFollowers(userID uint64, tokenID uint64, my bool) ([]model.User, error) {
	if userID != tokenID {
		return nil, errors.New("account doesn't match")
	}

	if my {
		return s.repo.FindFollowers(userID)
	}

	return s.repo.FindFollowing(userID)
}

func (s *userService) UpdateImage(userID uint64, tokenID uint64, imageData string) error {
	if userID != tokenID {
		return errors.New("account doesn't match")
	}
	return s.repo.UpdateImage(userID, imageData)
}

func (s *userService) UpdatePassword(requestBody []byte, userID uint64, userToken uint64) error {
	if userToken != userID {
		return errors.New("different account, action not possible")
	}

	password := model.Password{}
	if err := json.Unmarshal(requestBody, &password); err != nil {
		return errors.New("error unmarshaw body")
	}

	existingPassword, err := s.repo.FindPasswordById(userID)
	if err != nil {
		return errors.New("error finding existing password")
	}

	if err := authentication.Verify(existingPassword, password.New); err != nil {
		return errors.New("unauthorized")
	}

	hashPassword, err := authentication.Hash(password.New)
	if err != nil {
		return errors.New("error on transforming password")
	}

	if err := s.repo.UpdatePassword(userID, string(hashPassword)); err != nil {
		return errors.New("error updating password")
	}

	return nil
}

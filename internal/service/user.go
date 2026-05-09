package service

import (
	"github.com/bo-patrol/internal/domain"
	"github.com/bo-patrol/internal/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetByID(id uint) (*domain.UserResponse, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user.ToResponse(), nil
}

func (s *UserService) List(page, pageSize int) ([]*domain.UserResponse, int64, error) {
	users, total, err := s.userRepo.List(page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	result := make([]*domain.UserResponse, len(users))
	for i, user := range users {
		result[i] = user.ToResponse()
	}

	return result, total, nil
}

func (s *UserService) Update(id uint, updates map[string]interface{}) error {
	return s.userRepo.Update(id, updates)
}

func (s *UserService) Delete(id uint) error {
	return s.userRepo.Delete(id)
}

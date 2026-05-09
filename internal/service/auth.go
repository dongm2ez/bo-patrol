package service

import (
	"errors"

	"github.com/bo-patrol/internal/domain"
	"github.com/bo-patrol/internal/middleware"
	"github.com/bo-patrol/internal/repository"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Register(req *domain.RegisterRequest) (*domain.User, error) {
	_, err := s.userRepo.GetByUsername(req.Username)
	if err == nil {
		return nil, errors.New("用户名已存在")
	}

	user := &domain.User{
		Username: req.Username,
		Password: req.Password,
		Name:     req.Name,
		Phone:    req.Phone,
		Role:     domain.UserRole(req.Role),
		Status:   1,
	}

	if user.Role == "" {
		user.Role = domain.RoleInspector
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(req *domain.LoginRequest) (*domain.LoginResponse, error) {
	user, err := s.userRepo.GetByUsername(req.Username)
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	if !s.userRepo.VerifyPassword(user, req.Password) {
		return nil, errors.New("用户名或密码错误")
	}

	if user.Status != 1 {
		return nil, errors.New("用户已被禁用")
	}

	token, err := middleware.GenerateToken(user.ID, user.Username, string(user.Role))
	if err != nil {
		return nil, err
	}

	return &domain.LoginResponse{
		Token: token,
		User:  user,
	}, nil
}

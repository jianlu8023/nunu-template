package service

import (
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"
	"nunu-template/internal/model"
	"nunu-template/internal/pkg/request/user"
	"nunu-template/internal/pkg/response"
	"nunu-template/internal/repository"
)

type UserService interface {
	Register(ctx context.Context, req *user.RegisterRequest) error
	Login(ctx context.Context, req *user.LoginRequest) (string, error)
	GetProfile(ctx context.Context, userId string) (*model.User, error)
	UpdateProfile(ctx context.Context, userId string, req *user.UpdateProfileRequest) error
}

type userService struct {
	userRepo repository.UserRepository
	*Service
}

func NewUserService(service *Service, userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
		Service:  service,
	}
}

func (s *userService) Register(ctx context.Context, req *user.RegisterRequest) error {
	// 检查用户名是否已存在
	if user, err := s.userRepo.GetByUsername(ctx, req.Username); err == nil && user != nil {
		return response.ErrUsernameAlreadyUse
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return response.ErrUserPassWordEncrypt
	}
	// Generate user ID
	userId, err := s.sid.GenString()
	if err != nil {
		return response.ErrUserNotFound
	}
	// Create a user
	user := &model.User{
		UserId:   userId,
		Username: req.Username,
		Password: string(hashedPassword),
		Email:    req.Email,
	}
	if err = s.userRepo.Create(ctx, user); err != nil {
		return response.ErrCreateUser
	}

	return nil
}

func (s *userService) Login(ctx context.Context, req *user.LoginRequest) (string, error) {
	user, err := s.userRepo.GetByUsername(ctx, req.Username)
	if err != nil || user == nil {
		return "", response.ErrUsernameAlreadyUse
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", response.ErrEncryptPassword
	}
	token, err := s.jwt.GenToken(user.UserId, time.Now().Add(time.Hour*24*90))
	if err != nil {
		return "", response.ErrGenJWT
	}

	return token, nil
}

func (s *userService) GetProfile(ctx context.Context, userId string) (*model.User, error) {
	user, err := s.userRepo.GetByID(ctx, userId)
	if err != nil {
		return nil, response.ErrUsernameAlreadyUse
	}

	return user, nil
}

func (s *userService) UpdateProfile(ctx context.Context, userId string, req *user.UpdateProfileRequest) error {
	user, err := s.userRepo.GetByID(ctx, userId)
	if err != nil {
		return response.ErrUsernameAlreadyUse
	}

	user.Email = req.Email
	user.Nickname = req.Nickname

	if err = s.userRepo.Update(ctx, user); err != nil {
		return response.ErrUpdateUser
	}

	return nil
}

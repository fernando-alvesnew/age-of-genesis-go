package auth

import (
	"context"
	"errors"
	"strings"

	"github.com/alves/age-of-genesis/internal/domain/user"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrBannedUser         = errors.New("account banned")
)

type UserRepository interface {
	FindByLoginOrEmail(ctx context.Context, login string) (*user.User, error)
	UpdateLastIP(ctx context.Context, userID int64, ip string) error
}

type TokenService interface {
	Generate(userID int64, login string, userType string) (string, error)
}

type LoginService struct {
	userRepo UserRepository
	tokenSvc TokenService
}

type LoginInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
	IP       string
}

type LoginOutput struct {
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

func NewLoginService(userRepo UserRepository, tokenSvc TokenService) *LoginService {
	return &LoginService{
		userRepo: userRepo,
		tokenSvc: tokenSvc,
	}
}

func (s *LoginService) Execute(ctx context.Context, in LoginInput) (*LoginOutput, error) {
	u, err := s.userRepo.FindByLoginOrEmail(ctx, strings.TrimSpace(in.Login))
	if err != nil || u == nil {
		return nil, ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(in.Password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	if u.IsBanned {
		return nil, ErrBannedUser
	}

	if err := s.userRepo.UpdateLastIP(ctx, u.ID, in.IP); err != nil {
		return nil, err
	}

	token, err := s.tokenSvc.Generate(u.ID, u.Login, u.UserType)
	if err != nil {
		return nil, err
	}

	return &LoginOutput{
		UserID: u.ID,
		Token:  token,
	}, nil
}

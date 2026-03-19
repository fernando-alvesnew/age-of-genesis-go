package auth

import (
	"context"
	"testing"

	"github.com/alves/age-of-genesis/internal/domain/user"
	"golang.org/x/crypto/bcrypt"
)

type userRepoStub struct {
	user *user.User
}

func (s *userRepoStub) FindByLoginOrEmail(context.Context, string) (*user.User, error) {
	return s.user, nil
}

func (s *userRepoStub) UpdateLastIP(context.Context, int64, string) error {
	return nil
}

type tokenStub struct{}

func (s *tokenStub) Generate(int64, string, string) (string, error) {
	return "token123", nil
}

func TestLoginService_ExecuteSuccess(t *testing.T) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	repo := &userRepoStub{
		user: &user.User{
			ID:       1,
			Login:    "player",
			Password: string(hashed),
			UserType: "player",
		},
	}
	svc := NewLoginService(repo, &tokenStub{})

	out, err := svc.Execute(context.Background(), LoginInput{
		Login:    "player",
		Password: "123456",
		IP:       "127.0.0.1",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Token == "" {
		t.Fatalf("expected token")
	}
}

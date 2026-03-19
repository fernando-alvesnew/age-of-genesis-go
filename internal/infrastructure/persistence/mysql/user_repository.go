package mysql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/alves/age-of-genesis/internal/domain/user"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByLoginOrEmail(ctx context.Context, login string) (*user.User, error) {
	query := `
		SELECT id, login, email, password, user_type, is_banned, COALESCE(last_ip, '')
		FROM users
		WHERE login = ? OR email = ?
		LIMIT 1
	`

	row := r.db.QueryRowContext(ctx, query, login, login)
	var u user.User
	if err := row.Scan(&u.ID, &u.Login, &u.Email, &u.Password, &u.UserType, &u.IsBanned, &u.LastIP); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) UpdateLastIP(ctx context.Context, userID int64, ip string) error {
	_, err := r.db.ExecContext(ctx, "UPDATE users SET last_ip = ? WHERE id = ?", ip, userID)
	return err
}

package database

import (
	"database/sql"
	"github.com/tetsushi-yamada/hackathon_backend/internal/domain"
)

type UserDatabase struct{}

func NewUserDatabase() *UserDatabase { return &UserDatabase{} }

func (repo *UserDatabase) CreateUserTx(tx *sql.Tx, user domain.User) error {
	query := `INSERT INTO users (user_id, user_name, email) VALUES (?, ?, ?)`
	_, err := tx.Exec(query, user.UserId, user.UserName, user.Email)
	return err
}

func (repo *UserDatabase) GetUserTx(tx *sql.Tx, userID string) (*domain.User, error) {
	user := new(domain.User)
	query := `SELECT user_id, user_name, email, created_at, updated_at FROM users WHERE user_id = ?`
	err := tx.QueryRow(query, userID).Scan(&user.UserId, &user.UserName, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

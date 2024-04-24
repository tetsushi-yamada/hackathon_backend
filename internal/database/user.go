package database

import (
	"database/sql"
	"github.com/tetsushi-yamada/hackathon_backend/internal/domain/user"
	"time"
)

type UserDatabase struct{}

func NewUserDatabase() *UserDatabase { return &UserDatabase{} }

func (repo *UserDatabase) CreateUserTx(tx *sql.Tx, user user.User) error {
	query := `INSERT INTO users (user_id, user_name, email) VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE user_name = VALUES(user_name), email = VALUES(email)`
	_, err := tx.Exec(query, user.UserId, user.UserName, user.Email)
	if err != nil {
		return err
	}
	return err
}

func (repo *UserDatabase) GetUserTx(tx *sql.Tx, userID string) (*user.User, error) {
	user := new(user.User)
	query := `SELECT user_id, user_name, email, created_at, updated_at FROM users WHERE user_id = ?`
	var createdAt, updatedAt []byte
	err := tx.QueryRow(query, userID).Scan(&user.UserId, &user.UserName, &user.Email, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}
	user.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(createdAt))
	if err != nil {
		return nil, err
	}
	user.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", string(updatedAt))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserDatabase) DeleteUserTx(tx *sql.Tx, userID string) error {
	query := `DELETE FROM tweets WHERE user_id = ?`
	_, err := tx.Exec(query, userID)
	query = `DELETE FROM users WHERE user_id = ?`
	_, err = tx.Exec(query, userID)
	if err != nil {
		return err
	}
	return nil
}

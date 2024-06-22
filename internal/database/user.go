package database

import (
	"database/sql"
	"github.com/tetsushi-yamada/hackathon_backend/internal/domain/user"
	"time"
)

type UserDatabase struct{}

func NewUserDatabase() *UserDatabase { return &UserDatabase{} }

func (repo *UserDatabase) CreateUserTx(tx *sql.Tx, User user.User) error {
	var err error
	if User.UserDescription == nil || *User.UserDescription == "" {
		query := `INSERT INTO users (user_id, user_name, age) VALUES (?, ?, ?)`
		_, err = tx.Exec(query, User.UserID, User.UserName, User.Age)
	} else {
		query := `INSERT INTO users (user_id, user_name, age, user_description) VALUES (?, ?, ?, ?)`
		_, err = tx.Exec(query, User.UserID, User.UserName, User.Age, User.UserDescription)
	}
	if err != nil {
		return err
	}
	return err
}

func (repo *UserDatabase) GetUserTx(tx *sql.Tx, userID string) (*user.User, error) {
	user := new(user.User)
	query := `SELECT user_id, user_name, age, user_description, is_private, is_suspended, created_at, updated_at FROM users WHERE user_id = ?`
	var createdAt, updatedAt []byte
	var UserDescription sql.NullString
	err := tx.QueryRow(query, userID).Scan(&user.UserID, &user.UserName, &user.Age, &UserDescription, &user.IsPrivate, &user.IsSuspended, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}
	if UserDescription.Valid {
		user.UserDescription = &UserDescription.String
	} else {
		user.UserDescription = nil
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
	query := `DELETE FROM follows WHERE user_id = ?`
	_, err := tx.Exec(query, userID)
	query = `DELETE FROM tweets WHERE user_id = ?`
	_, err = tx.Exec(query, userID)
	query = `DELETE FROM users WHERE user_id = ?`
	_, err = tx.Exec(query, userID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserDatabase) UpdateUserTx(tx *sql.Tx, User user.User) error {
	var err error
	if User.UserDescription == nil || *User.UserDescription == "" {
		query := `UPDATE users SET user_name = ?, age = ?, is_private = ?, is_suspended = ? WHERE user_id = ?`
		_, err = tx.Exec(query, User.UserName, User.Age, User.IsPrivate, User.IsSuspended, User.UserID)
	} else {
		query := `UPDATE users SET user_name = ?, age = ?, user_description = ? , is_private = ?, is_suspended = ? WHERE user_id = ?`
		_, err = tx.Exec(query, User.UserName, User.Age, User.UserDescription, User.IsPrivate, User.IsSuspended, User.UserID)
	}
	if err != nil {
		return err
	}
	return nil

}

func (repo *UserDatabase) SearchUsersTx(tx *sql.Tx, userName string) ([]*user.User, error) {
	var users []*user.User
	query := `SELECT user_id, user_name, age, user_description, is_private, created_at, updated_at FROM users WHERE user_name LIKE ?`
	rows, err := tx.Query(query, "%"+userName+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		user := new(user.User)
		var createdAt, updatedAt []byte
		var UserDescription sql.NullString
		err = rows.Scan(&user.UserID, &user.UserName, &user.Age, &UserDescription, &user.IsPrivate, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}
		if UserDescription.Valid {
			user.UserDescription = &UserDescription.String
		} else {
			user.UserDescription = nil
		}
		user.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(createdAt))
		if err != nil {
			return nil, err
		}
		user.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", string(updatedAt))
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

package usecase

import (
	"database/sql"
	"github.com/tetsushi-yamada/hackathon_backend/internal/database"
	"github.com/tetsushi-yamada/hackathon_backend/internal/domain/user"
)

type UserUsecase struct {
	UserDatabase *database.UserDatabase
	sql          *sql.DB
}

func NewUserUsecase(db *sql.DB, ud *database.UserDatabase) *UserUsecase {
	return &UserUsecase{
		UserDatabase: ud,
		sql:          db,
	}
}

func (uu *UserUsecase) CreateUserUsecase(user user.User) error {
	db := uu.sql
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback() // ensures rollback if next steps fail

	if err = uu.UserDatabase.CreateUserTx(tx, user); err != nil {
		return err
	}
	return tx.Commit() // commit the transaction
}

func (uu *UserUsecase) GetUserUsecase(userID string) (*user.User, error) {
	db := uu.sql
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback() // ensures rollback if next steps fail

	user, err := uu.UserDatabase.GetUserTx(tx, userID)
	if err != nil {
		return nil, err
	}
	err = tx.Commit() // commit the transaction
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uu *UserUsecase) DeleteUserUsecase(userID string) error {
	db := uu.sql
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback() // ensures rollback if next steps fail

	err = uu.UserDatabase.DeleteUserTx(tx, userID)
	if err != nil {
		return err
	}
	err = tx.Commit() // commit the transaction
	if err != nil {
		return err
	}
	return nil
}

func (uu *UserUsecase) UpdateUserUsecase(user user.User) error {
	db := uu.sql
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback() // ensures rollback if next steps fail

	err = uu.UserDatabase.UpdateUserTx(tx, user)
	if err != nil {
		return err
	}
	err = tx.Commit() // commit the transaction
	if err != nil {
		return err
	}
	return nil

}

func (uu *UserUsecase) SearchUserUsecase(userName string) ([]*user.User, error) {
	db := uu.sql
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback() // ensures rollback if next steps fail

	users, err := uu.UserDatabase.SearchUsersTx(tx, userName)
	if err != nil {
		return nil, err
	}
	err = tx.Commit() // commit the transaction
	if err != nil {
		return nil, err
	}
	return users, nil
}

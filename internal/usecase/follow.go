package usecase

import (
	"database/sql"
	"github.com/tetsushi-yamada/hackathon_backend/internal/database"
	"github.com/tetsushi-yamada/hackathon_backend/internal/domain/follow"
)

type FollowUsecase struct {
	FollowDatabase *database.FollowDatabase
	sql            *sql.DB
}

func NewFollowUsecase(db *sql.DB, fd *database.FollowDatabase) *FollowUsecase {
	return &FollowUsecase{
		FollowDatabase: fd,
		sql:            db,
	}
}

func (fu *FollowUsecase) CreateFollowUsecase(follow follow.Follow) error {
	db := fu.sql
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback() // ensures rollback if next steps fail

	if err = fu.FollowDatabase.CreateFollowTx(tx, follow); err != nil {
		return err
	}
	return tx.Commit() // commit the transaction
}

func (fu *FollowUsecase) GetFollowsUsecase(userID string) ([]*follow.Follow, error) {
	db := fu.sql
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback() // ensures rollback if next steps fail

	follows, err := fu.FollowDatabase.GetFollowsTx(tx, userID)
	if err != nil {
		return nil, err
	}
	err = tx.Commit() // commit the transaction
	if err != nil {
		return nil, err
	}
	return follows, nil
}

func (fu *FollowUsecase) DeleteFollowUsecase(userID string, followID string) error {
	db := fu.sql
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback() // ensures rollback if next steps fail

	err = fu.FollowDatabase.DeleteFollowTx(tx, userID, followID)
	if err != nil {
		return err
	}
	err = tx.Commit() // commit the transaction
	if err != nil {
		return err
	}
	return nil
}

func (fu *FollowUsecase) GetFollowOrNotUsecase(userID string, followID string) (*follow.FollowOrNot, error) {
	db := fu.sql
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback() // ensures rollback if next steps fail

	follow_bool, err := fu.FollowDatabase.GetFollowOrNotTx(tx, userID, followID)
	if err != nil {
		return nil, err
	}
	err = tx.Commit() // commit the transaction
	if err != nil {
		return nil, err
	}
	return follow_bool, nil
}

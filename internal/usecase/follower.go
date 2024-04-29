package usecase

import (
	"database/sql"
	"github.com/tetsushi-yamada/hackathon_backend/internal/database"
	"github.com/tetsushi-yamada/hackathon_backend/internal/domain/follow"
)

type FollowerUsecase struct {
	FollowDatabase *database.FollowDatabase
	sql            *sql.DB
}

func NewFollowerUsecase(db *sql.DB, fd *database.FollowDatabase) *FollowerUsecase {
	return &FollowerUsecase{
		FollowDatabase: fd,
		sql:            db,
	}
}

func (fu *FollowerUsecase) GetFollowersUsecase(followID string) ([]*follow.Follow, error) {
	db := fu.sql
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback() // ensures rollback if next steps fail

	followers, err := fu.FollowDatabase.GetFollowersTx(tx, followID)
	if err != nil {
		return nil, err
	}
	err = tx.Commit() // commit the transaction
	if err != nil {
		return nil, err
	}
	return followers, nil
}

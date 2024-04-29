package usecase

import (
	"database/sql"
	"github.com/tetsushi-yamada/hackathon_backend/internal/database"
	"github.com/tetsushi-yamada/hackathon_backend/internal/domain/follow"
)

type FollowOrNotUsecase struct {
	FollowDatabase *database.FollowDatabase
	sql            *sql.DB
}

func NewFollowOrNotUsecase(db *sql.DB, fd *database.FollowDatabase) *FollowOrNotUsecase {
	return &FollowOrNotUsecase{FollowDatabase: fd, sql: db}
}

func (fu *FollowOrNotUsecase) GetFollowOrNotUsecase(userID string, followID string) (*follow.FollowOrNot, error) {
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

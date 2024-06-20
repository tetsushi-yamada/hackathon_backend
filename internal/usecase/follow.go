package usecase

import (
	"database/sql"
	"github.com/tetsushi-yamada/hackathon_backend/internal/database"
	"github.com/tetsushi-yamada/hackathon_backend/internal/domain/follow"
)

type FollowUsecase struct {
	FollowDatabase        *database.FollowDatabase
	FollowRequestDatabase *database.FollowRequestDatabase
	UserDatabase          *database.UserDatabase
	sql                   *sql.DB
}

func NewFollowUsecase(db *sql.DB, fd *database.FollowDatabase, frd *database.FollowRequestDatabase, ud *database.UserDatabase) *FollowUsecase {
	return &FollowUsecase{
		FollowDatabase:        fd,
		FollowRequestDatabase: frd,
		UserDatabase:          ud,
		sql:                   db,
	}
}

func (fu *FollowUsecase) CreateFollowUsecase(follow_domain follow.Follow) error {
	db := fu.sql
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback() // ensures rollback if next steps fail

	user, err := fu.UserDatabase.GetUserTx(tx, follow_domain.FollowID)
	if err != nil {
		return err
	}
	if user.IsPrivate {
		followRequest := follow.FollowRequest{
			UserID:   follow_domain.UserID,
			FollowID: follow_domain.FollowID,
		}
		if err = fu.FollowRequestDatabase.CreateFollowRequestTx(tx, followRequest); err != nil {
			return err
		}
	} else {
		if err = fu.FollowDatabase.CreateFollowTx(tx, follow_domain); err != nil {
			return err
		}
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

func (fu *FollowUsecase) UpdateFollowRequestUsecase(followRequest follow.FollowRequest) error {
	db := fu.sql
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback() // ensures rollback if next steps fail

	err = fu.FollowRequestDatabase.UpdateFollowRequestTx(tx, followRequest)
	if err != nil {
		return err
	}

	if followRequest.Status == "approved" {
		follow := follow.Follow{
			UserID:   followRequest.UserID,
			FollowID: followRequest.FollowID,
		}
		if err = fu.FollowDatabase.CreateFollowTx(tx, follow); err != nil {
			return err
		}
		if err = fu.FollowRequestDatabase.DeleteFollowRequestTx(tx, followRequest.UserID, followRequest.FollowID); err != nil {
			return err
		}
	} else if followRequest.Status == "rejected" {
		if err = fu.FollowRequestDatabase.DeleteFollowRequestTx(tx, followRequest.UserID, followRequest.FollowID); err != nil {
			return err
		}
	}

	return tx.Commit() // commit the transaction
}

func (fu *FollowUsecase) GetFollowRequestsUsecase(followID string) ([]*follow.FollowRequest, error) {
	db := fu.sql
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback() // ensures rollback if next steps fail

	followRequests, err := fu.FollowRequestDatabase.GetFollowRequestsTx(tx, followID)
	if err != nil {
		return nil, err
	}
	err = tx.Commit() // commit the transaction
	if err != nil {
		return nil, err
	}
	return followRequests, nil
}

func (fu *FollowUsecase) GetFollowRequestUsecase(userID string, followID string) (*follow.FollowRequestOrNot, error) {
	db := fu.sql
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback() // ensures rollback if next steps fail

	followRequest, err := fu.FollowRequestDatabase.GetFollowRequestTx(tx, userID, followID)
	if err != nil {
		return nil, err
	}
	err = tx.Commit() // commit the transaction
	if err != nil {
		return nil, err
	}
	return followRequest, nil
}

package usecase

import (
	"database/sql"
	"github.com/tetsushi-yamada/hackathon_backend/internal/database"
	"github.com/tetsushi-yamada/hackathon_backend/internal/domain/good"
)

type GoodUsecase struct {
	GoodDatabase *database.GoodDatabase
	sql          *sql.DB
}

func NewGoodUsecase(db *sql.DB, gd *database.GoodDatabase) *GoodUsecase {
	return &GoodUsecase{
		GoodDatabase: gd,
		sql:          db,
	}
}

func (gu *GoodUsecase) CreateGoodUsecase(good good.Good) error {
	db := gu.sql
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback() // ensures rollback if next steps fail

	if err = gu.GoodDatabase.CreateGoodTx(tx, good); err != nil {
		return err
	}
	return tx.Commit() // commit the transaction
}

func (gu *GoodUsecase) GetGoodsUsecaseByTweetID(tweetID string) ([]*good.Good, error) {
	db := gu.sql
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback() // ensures rollback if next steps fail

	goods, err := gu.GoodDatabase.GetGoodsTxByTweetID(tx, tweetID)
	if err != nil {
		return nil, err
	}
	err = tx.Commit() // commit the transaction
	if err != nil {
		return nil, err
	}
	return goods, nil
}

func (gu *GoodUsecase) GetGoodsUsecaseByUserID(userID string) ([]*good.Good, error) {
	db := gu.sql
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback() // ensures rollback if next steps fail

	goods, err := gu.GoodDatabase.GetGoodsTxByUserID(tx, userID)
	if err != nil {
		return nil, err
	}
	err = tx.Commit() // commit the transaction
	if err != nil {
		return nil, err
	}
	return goods, nil
}

func (gu *GoodUsecase) DeleteGoodUsecase(tweetID string, userID string) error {
	db := gu.sql
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback() // ensures rollback if next steps fail

	err = gu.GoodDatabase.DeleteGoodTx(tx, tweetID, userID)
	if err != nil {
		return err
	}
	err = tx.Commit() // commit the transaction
	if err != nil {
		return err
	}
	return nil
}

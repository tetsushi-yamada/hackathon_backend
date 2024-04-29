package usecase

import (
	"database/sql"
	"github.com/tetsushi-yamada/hackathon_backend/internal/database"
	"github.com/tetsushi-yamada/hackathon_backend/internal/domain/tweet"
)

type TweetUsecase struct {
	TweetDatabase *database.TweetDatabase
	sql           *sql.DB
}

func NewTweetUsecase(db *sql.DB, td *database.TweetDatabase) *TweetUsecase {
	return &TweetUsecase{
		TweetDatabase: td,
		sql:           db,
	}
}

func (tu *TweetUsecase) CreateTweetUsecase(tweet tweet.Tweet) error {
	db := tu.sql
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback() // ensures rollback if next steps fail

	if err = tu.TweetDatabase.CreateTweetTx(tx, tweet); err != nil {
		return err
	}
	return tx.Commit() // commit the transaction
}

func (tu *TweetUsecase) GetTweetsUsecase(userID string) ([]*tweet.Tweet, error) {
	db := tu.sql
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback() // ensures rollback if next steps fail

	tweets, err := tu.TweetDatabase.GetTweetsTx(tx, userID)
	if err != nil {
		return nil, err
	}
	err = tx.Commit() // commit the transaction
	if err != nil {
		return nil, err
	}
	return tweets, nil
}

func (tu *TweetUsecase) DeleteTweetUsecase(tweetID string) error {
	db := tu.sql
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback() // ensures rollback if next steps fail

	err = tu.TweetDatabase.DeleteTweetTx(tx, tweetID)
	if err != nil {
		return err
	}
	err = tx.Commit() // commit the transaction
	if err != nil {
		return err
	}
	return nil
}

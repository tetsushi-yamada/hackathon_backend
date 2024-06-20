package usecase

import (
	"database/sql"
	"github.com/tetsushi-yamada/hackathon_backend/internal/database"
	"github.com/tetsushi-yamada/hackathon_backend/internal/domain/tweet_picture"
)

type TweetPictureUsecase struct {
	TweetDatabase *database.TweetDatabase
	sql           *sql.DB
}

func NewTweetPictureUsecase(db *sql.DB, td *database.TweetDatabase) *TweetPictureUsecase {
	return &TweetPictureUsecase{
		TweetDatabase: td,
		sql:           db,
	}
}

func (tpu *TweetPictureUsecase) UploadTweetPicture(tweetPicture tweet_picture.TweetPicture) error {
	db := tpu.sql
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback() // ensures rollback if next steps fail

	if err = tpu.TweetDatabase.CreateTweetPictureTx(tx, tweetPicture); err != nil {
		return err
	}
	return tx.Commit() // commit the transaction
}

func (tpu *TweetPictureUsecase) GetTweetPicture(tweetID string) (*tweet_picture.TweetPicture, error) {
	db := tpu.sql
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback() // ensures rollback if next steps fail

	tweetPicture, err := tpu.TweetDatabase.GetTweetPictureTx(tx, tweetID)
	if err != nil {
		return nil, err
	}
	err = tx.Commit() // commit the transaction
	if err != nil {
		return nil, err
	}
	return tweetPicture, nil
}

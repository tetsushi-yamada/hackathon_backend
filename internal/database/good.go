package database

import (
	"database/sql"
	"github.com/tetsushi-yamada/hackathon_backend/internal/domain/good"
	"time"
)

type GoodDatabase struct {
}

func NewGoodDatabase() *GoodDatabase {
	return &GoodDatabase{}
}

func (gd *GoodDatabase) CreateGoodTx(tx *sql.Tx, good good.Good) error {
	_, err := tx.Exec("INSERT INTO goods (tweet_id, user_id) VALUES (?, ?)", good.TweetID, good.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *GoodDatabase) GetGoodsTxByTweetID(tx *sql.Tx, tweetID string) ([]*good.Good, error) {
	rows, err := tx.Query("SELECT * FROM goods WHERE tweet_id = ?", tweetID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var goods []*good.Good
	for rows.Next() {
		var good good.Good
		var createdAt []byte
		if err := rows.Scan(&good.TweetID, &good.UserID, &createdAt); err != nil {
			return nil, err
		}
		good.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(createdAt))
		if err != nil {
			return nil, err
		}
		goods = append(goods, &good)
	}
	return goods, nil
}

func (repo *GoodDatabase) GetGoodsTxByUserID(tx *sql.Tx, userID string) ([]*good.Good, error) {
	rows, err := tx.Query("SELECT * FROM goods WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var goods []*good.Good
	for rows.Next() {
		var good good.Good
		var createdAt []byte
		if err := rows.Scan(&good.TweetID, &good.UserID, &createdAt); err != nil {
			return nil, err
		}
		good.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(createdAt))
		if err != nil {
			return nil, err
		}
		goods = append(goods, &good)
	}
	return goods, nil
}

func (repo *GoodDatabase) DeleteGoodTx(tx *sql.Tx, tweetID string, userID string) error {
	_, err := tx.Exec("DELETE FROM goods WHERE tweet_id = ? AND user_id = ?", tweetID, userID)
	if err != nil {
		return err
	}
	return nil
}

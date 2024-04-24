package database

import (
	"database/sql"
	"github.com/tetsushi-yamada/hackathon_backend/internal/domain/tweet"
	"time"
)

type TweetDatabase struct{}

func NewTweetDatabase() *TweetDatabase { return &TweetDatabase{} }

func (repo *TweetDatabase) CreateTweetTx(tx *sql.Tx, tweet tweet.Tweet) error {
	query := `DELETE FROM tweets WHERE tweet_id = ?`
	_, err := tx.Exec(query, tweet.TweetID)
	if err != nil {
		return err
	}
	query = `INSERT INTO tweets (tweet_id, user_id, tweet_text) VALUES (?, ?, ?)`
	_, err = tx.Exec(query, tweet.TweetID, tweet.UserID, tweet.TweetText)
	if err != nil {
		return err
	}
	return err
}

func (repo *TweetDatabase) GetTweetsTx(tx *sql.Tx, userID string) ([]*tweet.Tweet, error) {
	var tweets []*tweet.Tweet
	query := `SELECT tweet_id, user_id, tweet_text, created_at, updated_at FROM tweets WHERE user_id = ?`
	rows, err := tx.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		tweet := new(tweet.Tweet)
		var createdAt, updatedAt []byte
		if err := rows.Scan(&tweet.TweetID, &tweet.UserID, &tweet.TweetText, &createdAt, &updatedAt); err != nil {
			return nil, err
		}
		tweet.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(createdAt))
		if err != nil {
			return nil, err
		}
		tweet.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", string(updatedAt))
		if err != nil {
			return nil, err
		}
		tweets = append(tweets, tweet)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tweets, nil
}

func (repo *TweetDatabase) DeleteTweetTx(tx *sql.Tx, tweetID string) error {
	query := `DELETE FROM tweets WHERE tweet_id = ?`
	_, err := tx.Exec(query, tweetID)
	if err != nil {
		return err
	}
	return nil
}

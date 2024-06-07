package database

import (
	"database/sql"
	"github.com/tetsushi-yamada/hackathon_backend/internal/domain/tweet"
	"github.com/tetsushi-yamada/hackathon_backend/internal/domain/tweet_picture"
	"time"
)

type TweetDatabase struct{}

func NewTweetDatabase() *TweetDatabase { return &TweetDatabase{} }

func (repo *TweetDatabase) CreateTweetTx(tx *sql.Tx, tweet tweet.Tweet) error {
	var err error
	if tweet.ParentID == nil || *tweet.ParentID == "" {
		query := `INSERT INTO tweets (tweet_id, user_id, tweet_text) VALUES (?, ?, ?)`
		_, err = tx.Exec(query, tweet.TweetID, tweet.UserID, tweet.TweetText)
	} else {
		query := `INSERT INTO tweets (tweet_id, user_id, tweet_text, parent_id) VALUES (?, ?, ?, ?)`
		_, err = tx.Exec(query, tweet.TweetID, tweet.UserID, tweet.TweetText, tweet.ParentID)
	}
	if err != nil {
		return err
	}
	return err
}

func (repo *TweetDatabase) GetTweetsTx(tx *sql.Tx, userID string) ([]*tweet.Tweet, error) {
	var tweets []*tweet.Tweet
	query := `SELECT tweet_id, user_id, tweet_text, parent_id, created_at, updated_at FROM tweets WHERE user_id = ?`
	rows, err := tx.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		tweet := new(tweet.Tweet)
		var createdAt, updatedAt []byte
		var parentID sql.NullString
		if err := rows.Scan(&tweet.TweetID, &tweet.UserID, &tweet.TweetText, &parentID, &createdAt, &updatedAt); err != nil {
			return nil, err
		}
		if parentID.Valid {
			tweet.ParentID = &parentID.String
		} else {
			tweet.ParentID = nil
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

func (repo *TweetDatabase) GetTweetByTweetIDTx(tx *sql.Tx, tweetID string) (*tweet.Tweet, error) {
	tweet := new(tweet.Tweet)
	var createdAt, updatedAt []byte
	var parentID sql.NullString
	query := `SELECT tweet_id, user_id, tweet_text, parent_id, created_at, updated_at FROM tweets WHERE tweet_id = ?`
	err := tx.QueryRow(query, tweetID).Scan(&tweet.TweetID, &tweet.UserID, &tweet.TweetText, &parentID, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}
	if parentID.Valid {
		tweet.ParentID = &parentID.String
	} else {
		tweet.ParentID = nil
	}
	tweet.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(createdAt))
	if err != nil {
		return nil, err
	}
	tweet.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", string(updatedAt))
	if err != nil {
		return nil, err
	}
	return tweet, nil
}

func (repo *TweetDatabase) UpdateTweetTx(tx *sql.Tx, tweet tweet.Tweet) error {
	query := `UPDATE tweets SET tweet_text = ? WHERE tweet_id = ?`
	_, err := tx.Exec(query, tweet.TweetText, tweet.TweetID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *TweetDatabase) DeleteTweetTx(tx *sql.Tx, tweetID string) error {
	query := `DELETE FROM tweets WHERE tweet_id = ?`
	_, err := tx.Exec(query, tweetID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *TweetDatabase) SearchTweetsTx(tx *sql.Tx, keyword string) ([]*tweet.Tweet, error) {
	var tweets []*tweet.Tweet
	query := `SELECT tweet_id, user_id, tweet_text, parent_id, created_at, updated_at FROM tweets WHERE tweet_text LIKE ?`
	rows, err := tx.Query(query, "%"+keyword+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		tweet := new(tweet.Tweet)
		var createdAt, updatedAt []byte
		var parentID sql.NullString
		if err := rows.Scan(&tweet.TweetID, &tweet.UserID, &tweet.TweetText, &parentID, &createdAt, &updatedAt); err != nil {
			return nil, err
		}
		if parentID.Valid {
			tweet.ParentID = &parentID.String
		} else {
			tweet.ParentID = nil
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

func (repo *TweetDatabase) GetRepliesTx(tx *sql.Tx, tweetID string) ([]*tweet.Tweet, error) {
	var tweets []*tweet.Tweet
	query := `SELECT tweet_id, user_id, tweet_text, parent_id, created_at, updated_at FROM tweets WHERE parent_id = ?`
	rows, err := tx.Query(query, tweetID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		tweet := new(tweet.Tweet)
		var createdAt, updatedAt []byte
		var parentID sql.NullString
		if err := rows.Scan(&tweet.TweetID, &tweet.UserID, &tweet.TweetText, &parentID, &createdAt, &updatedAt); err != nil {
			return nil, err
		}
		if parentID.Valid {
			tweet.ParentID = &parentID.String
		} else {
			tweet.ParentID = nil
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

func (repo *TweetDatabase) CreateTweetPictureTx(tx *sql.Tx, tweetPicture tweet_picture.TweetPicture) error {
	query := `INSERT INTO tweet_pictures (tweet_id, tweet_picture) VALUES (?, ?) ON DUPLICATE KEY UPDATE tweet_picture = VALUES(tweet_picture)`
	_, err := tx.Exec(query, tweetPicture.TweetID, tweetPicture.TweetPicture)
	if err != nil {
		return err
	}
	return nil
}

func (repo *TweetDatabase) GetTweetPictureTx(tx *sql.Tx, tweetID string) (*tweet_picture.TweetPicture, error) {
	var tweetPicture tweet_picture.TweetPicture
	query := `SELECT tweet_id, tweet_picture FROM tweet_pictures WHERE tweet_id = ?`
	row := tx.QueryRow(query, tweetID)
	err := row.Scan(&tweetPicture.TweetID, &tweetPicture.TweetPicture)
	if err != nil {
		return nil, err
	}
	return &tweetPicture, nil
}

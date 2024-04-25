package database

import (
	"database/sql"
	"github.com/tetsushi-yamada/hackathon_backend/internal/domain/follow"
	"time"
)

type FollowDatabase struct{}

func NewFollowDatabase() *FollowDatabase { return &FollowDatabase{} }

func (repo *FollowDatabase) CreateFollowTx(tx *sql.Tx, follow follow.Follow) error {
	query := `INSERT INTO follows (user_id, follow_id) VALUES (?, ?)`
	_, err := tx.Exec(query, follow.UserID, follow.FollowID)
	if err != nil {
		return err
	}
	return err
}

func (repo *FollowDatabase) GetFollowsTx(tx *sql.Tx, userID string) ([]*follow.Follow, error) {
	var follows []*follow.Follow
	query := `SELECT user_id, follow_id, created_at FROM follows WHERE user_id = ?`
	rows, err := tx.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		follow := new(follow.Follow)
		var createdAt []byte
		if err := rows.Scan(&follow.UserID, &follow.FollowID, &createdAt); err != nil {
			return nil, err
		}
		follow.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(createdAt))
		if err != nil {
			return nil, err
		}
		follows = append(follows, follow)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return follows, nil
}

func (repo *FollowDatabase) DeleteFollowTx(tx *sql.Tx, userID string, followID string) error {
	query := `DELETE FROM follows WHERE user_id = ? AND follow_id = ?`
	_, err := tx.Exec(query, userID, followID)
	if err != nil {
		return err
	}
	return nil
}

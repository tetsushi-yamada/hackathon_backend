package database

import (
	"database/sql"
	"github.com/tetsushi-yamada/hackathon_backend/internal/domain/follow"
	"time"
)

type FollowRequestDatabase struct{}

func NewFollowRequestDatabase() *FollowRequestDatabase { return &FollowRequestDatabase{} }

func (repo *FollowRequestDatabase) CreateFollowRequestTx(tx *sql.Tx, followRequest follow.FollowRequest) error {
	query := `DELETE FROM follow_requests WHERE user_id = ? AND follow_id = ?`
	_, err := tx.Exec(query, followRequest.UserID, followRequest.FollowID)
	if err != nil {
		return err
	}
	query = `INSERT INTO follow_requests (user_id, follow_id) VALUES (?, ?)`
	_, err = tx.Exec(query, followRequest.UserID, followRequest.FollowID)
	if err != nil {
		return err
	}
	return err
}

func (repo *FollowRequestDatabase) GetFollowRequestTx(tx *sql.Tx, userID string, followID string) (*follow.FollowRequestOrNot, error) {
	var followRequestOrNot follow.FollowRequestOrNot
	query := `SELECT EXISTS (SELECT * FROM follow_requests WHERE user_id = ? AND follow_id = ?) AS BOOL`
	row := tx.QueryRow(query, userID, followID)
	if err := row.Scan(&followRequestOrNot.Bool); err != nil {
		return nil, err
	}
	return &followRequestOrNot, nil
}

func (repo *FollowRequestDatabase) GetFollowRequestsTx(tx *sql.Tx, followID string) ([]*follow.FollowRequest, error) {
	var followRequests []*follow.FollowRequest
	query := `SELECT user_id, follow_id, status, created_at FROM follow_requests WHERE follow_id = ?`
	rows, err := tx.Query(query, followID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		followRequest := new(follow.FollowRequest)
		var createdAt []byte
		if err := rows.Scan(&followRequest.UserID, &followRequest.FollowID, &followRequest.Status, &createdAt); err != nil {
			return nil, err
		}
		followRequest.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(createdAt))
		followRequests = append(followRequests, followRequest)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return followRequests, nil
}

func (repo *FollowRequestDatabase) DeleteFollowRequestTx(tx *sql.Tx, userID string, followID string) error {
	query := `DELETE FROM follow_requests WHERE user_id = ? AND follow_id = ?`
	_, err := tx.Exec(query, userID, followID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *FollowRequestDatabase) UpdateFollowRequestTx(tx *sql.Tx, followRequest follow.FollowRequest) error {
	query := `UPDATE follow_requests SET status = ? WHERE user_id = ? AND follow_id = ?`
	_, err := tx.Exec(query, followRequest.Status, followRequest.UserID, followRequest.FollowID)
	if err != nil {
		return err
	}
	return nil
}

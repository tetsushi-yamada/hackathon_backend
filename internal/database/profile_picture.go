package database

import (
	"database/sql"
	"github.com/tetsushi-yamada/hackathon_backend/internal/domain/profile_picture"
)

type ProfilePictureDatabase struct{}

func NewProfilePictureDatabase() *ProfilePictureDatabase {
	return &ProfilePictureDatabase{}
}

func (pd *ProfilePictureDatabase) CreateProfilePictureTx(tx *sql.Tx, profilePicture profile_picture.ProfilePicture) error {
	query := `INSERT INTO profile_pictures (user_id, profile_picture) VALUES (?, ?) ON DUPLICATE KEY UPDATE profile_picture = VALUES(profile_picture)`
	_, err := tx.Exec(query, profilePicture.UserID, profilePicture.ProfilePicture)
	if err != nil {
		return err
	}
	return err
}

func (pd *ProfilePictureDatabase) GetProfilePictureTx(tx *sql.Tx, userID string) (*profile_picture.ProfilePicture, error) {
	var ProfilePicture profile_picture.ProfilePicture
	query := `SELECT user_id, profile_picture FROM profile_pictures WHERE user_id = ?`
	row := tx.QueryRow(query, userID)
	err := row.Scan(&ProfilePicture.UserID, &ProfilePicture.ProfilePicture)
	if err != nil {
		return nil, err
	}
	return &ProfilePicture, nil
}

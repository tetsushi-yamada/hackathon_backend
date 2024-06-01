package usecase

import (
	"database/sql"
	"github.com/tetsushi-yamada/hackathon_backend/internal/database"
	"github.com/tetsushi-yamada/hackathon_backend/internal/domain/profile_picture"
)

type ProfilePictureUsecase struct {
	ProfilePictureRepository *database.ProfilePictureDatabase
	sql                      *sql.DB
}

func NewProfilePictureUsecase(db *sql.DB, pd *database.ProfilePictureDatabase) *ProfilePictureUsecase {
	return &ProfilePictureUsecase{
		ProfilePictureRepository: pd,
		sql:                      db,
	}
}

func (pu *ProfilePictureUsecase) UploadProfilePicture(profilePicture profile_picture.ProfilePicture) error {
	db := pu.sql
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback() // ensures rollback if next steps fail

	if err = pu.ProfilePictureRepository.CreateProfilePictureTx(tx, profilePicture); err != nil {
		return err
	}
	return tx.Commit() // commit the transaction
}

func (pu *ProfilePictureUsecase) GetProfilePicture(userID string) (*profile_picture.ProfilePicture, error) {
	db := pu.sql
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback() // ensures rollback if next steps fail

	profilePicture, err := pu.ProfilePictureRepository.GetProfilePictureTx(tx, userID)
	if err != nil {
		return nil, err
	}
	err = tx.Commit() // commit the transaction
	if err != nil {
		return nil, err
	}
	return profilePicture, nil
}

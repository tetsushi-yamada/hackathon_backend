package usecase

import (
	"database/sql"
	"github.com/tetsushi-yamada/hackathon_backend/internal/database"
	"github.com/tetsushi-yamada/hackathon_backend/internal/domain/block"
)

type BlockUsecase struct {
	br  *database.BlockDatabase
	sql *sql.DB
}

func NewBlockUsecase(db *sql.DB, br *database.BlockDatabase) *BlockUsecase {
	return &BlockUsecase{
		br:  br,
		sql: db,
	}
}

func (bu *BlockUsecase) CreateBlockUsecase(block_domain block.Block) error {
	db := bu.sql
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback() // ensures rollback if next steps fail

	if err = bu.br.CreateBlockTx(tx, block_domain); err != nil {
		return err
	}
	return tx.Commit() // commit the transaction
}

func (bu *BlockUsecase) GetBlocksUsecase(userID string, blockID string) (*block.BlockOrNot, error) {
	db := bu.sql
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback() // ensures rollback if next steps fail

	blockOrNot, err := bu.br.GetBlockTx(tx, userID, blockID)
	if err != nil {
		return nil, err
	}
	return blockOrNot, nil
}

func (bu *BlockUsecase) DeleteBlockUsecase(userID string, blockID string) error {
	db := bu.sql
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback() // ensures rollback if next steps fail

	err = bu.br.DeleteBlockTx(tx, userID, blockID)
	if err != nil {
		return err
	}
	return tx.Commit() // commit the transaction
}

package database

import (
	"database/sql"
	"github.com/tetsushi-yamada/hackathon_backend/internal/domain/block"
)

type BlockDatabase struct{}

func NewBlockDatabase() *BlockDatabase {
	return &BlockDatabase{}
}

func (bd *BlockDatabase) CreateBlockTx(tx *sql.Tx, block_domain block.Block) error {
	_, err := tx.Exec("INSERT INTO blocks (user_id, block_id) VALUES (?, ?)", block_domain.UserID, block_domain.BlockID)
	if err != nil {
		return err
	}
	return nil
}

func (bd *BlockDatabase) GetBlockTx(tx *sql.Tx, userID string, blockID string) (*block.BlockOrNot, error) {
	row := tx.QueryRow("SELECT EXISTS (SELECT 1 FROM blocks WHERE user_id = ? AND block_id = ?)", userID, blockID)

	var blockOrNot block.BlockOrNot
	err := row.Scan(&blockOrNot.Bool)
	if err != nil {
		return nil, err
	}
	return &blockOrNot, nil
}

func (bd *BlockDatabase) DeleteBlockTx(tx *sql.Tx, userID string, blockID string) error {
	_, err := tx.Exec("DELETE FROM blocks WHERE user_id = ? AND block_id = ?", userID, blockID)
	if err != nil {
		return err
	}
	return nil
}

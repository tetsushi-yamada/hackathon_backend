package block

import "time"

type Block struct {
	UserID    string    `json:"user_id" db:"user_id"`
	BlockID   string    `json:"block_id" db:"block_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type BlockOrNot struct {
	Bool bool `json:"bool"`
}

package init_query

import "database/sql"

func DropBlockTable(db *sql.DB) error {
	// テーブル削除用のSQL文
	dropTableSQL := `DROP TABLE IF EXISTS blocks;`

	// SQL文の実行
	_, err := db.Exec(dropTableSQL)
	if err != nil {
		return err
	}
	return nil

}

func CreateBlockTable(db *sql.DB) error {
	// テーブル作成用のSQL文
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS blocks (
		user_id char(36) not null,
		block_id char(36) not null,
		created_at timestamp not null default current_timestamp,
		PRIMARY KEY (user_id, block_id),
		FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
	    FOREIGN KEY (block_id) REFERENCES users(user_id) ON DELETE CASCADE
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
	`

	// SQL文の実行
	_, err := db.Exec(createTableSQL)
	if err != nil {
		return err
	}
	return nil
}

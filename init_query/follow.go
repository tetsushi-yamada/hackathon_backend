package init_query

import "database/sql"

func CreateFollowTable(db *sql.DB) error {
	// テーブル作成用のSQL文
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS follows (
		user_id char(36) not null,
		follow_id char(36) not null,
		created_at timestamp not null default current_timestamp,
		PRIMARY KEY (user_id, follow_id),
		FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
	`

	// SQL文の実行
	_, err := db.Exec(createTableSQL)
	if err != nil {
		return err
	}
	return nil
}

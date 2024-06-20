package init_query

import "database/sql"

func DropFollowRequestTable(db *sql.DB) error {
	// テーブル削除用のSQL文
	dropTableSQL := `DROP TABLE IF EXISTS follow_requests;`

	// SQL文の実行
	_, err := db.Exec(dropTableSQL)
	if err != nil {
		return err
	}
	return nil

}

func CreateFollowRequestTable(db *sql.DB) error {
	// テーブル作成用のSQL文
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS follow_requests (
		user_id char(36) not null,
		follow_id char(36) not null,
	    status varchar(32) not null default 'pending',
		created_at timestamp not null default current_timestamp,
		PRIMARY KEY (user_id, follow_id),
		FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
	    FOREIGN KEY (follow_id) REFERENCES users(user_id) ON DELETE CASCADE
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
	`

	// SQL文の実行
	_, err := db.Exec(createTableSQL)
	if err != nil {
		return err
	}
	return nil
}

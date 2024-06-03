package init_query

import "database/sql"

func DropProfilePictureTable(db *sql.DB) error {
	// テーブル削除用のSQL文
	dropTableSQL := `DROP TABLE IF EXISTS profile_pictures;`

	// SQL文の実行
	_, err := db.Exec(dropTableSQL)
	if err != nil {
		return err
	}
	return nil
}

func CreateProfilePictureTable(db *sql.DB) error {
	// テーブル作成用のSQL文
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS profile_pictures (
		user_id char(36) not null,
		profile_picture varchar(255) not null,
		created_at timestamp not null default current_timestamp,
		updated_at timestamp not null default current_timestamp on update current_timestamp,
		PRIMARY KEY (user_id),
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

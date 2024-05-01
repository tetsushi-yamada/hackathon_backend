package init_query

import "database/sql"

func DropUserTable(db *sql.DB) error {
	// テーブル削除用のSQL文
	dropTableSQL := `DROP TABLE IF EXISTS users;`

	// SQL文の実行
	_, err := db.Exec(dropTableSQL)
	if err != nil {
		return err
	}
	return nil
}

func CreateUserTable(db *sql.DB) error {
	// テーブル作成用のSQL文
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		user_id char(36) not null,
		user_name varchar(32) not null,
		created_at timestamp not null default current_timestamp,
		updated_at timestamp not null default current_timestamp on update current_timestamp,
		PRIMARY KEY (user_id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
	`

	// SQL文の実行
	_, err := db.Exec(createTableSQL)
	if err != nil {
		return err
	}
	return nil
}

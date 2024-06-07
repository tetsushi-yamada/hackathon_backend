package init_query

import "database/sql"

func DropTweetPictureTable(db *sql.DB) error {
	// テーブル削除用のSQL文
	dropTableSQL := `DROP TABLE IF EXISTS tweet_pictures;`

	// SQL文の実行
	_, err := db.Exec(dropTableSQL)
	if err != nil {
		return err
	}
	return nil
}

func CreateTweetPictureTable(db *sql.DB) error {
	// テーブル作成用のSQL文
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS tweet_pictures (
		tweet_id char(36) not null,
		tweet_picture varchar(255) not null,
		created_at timestamp not null default current_timestamp,
		updated_at timestamp not null default current_timestamp on update current_timestamp,
		PRIMARY KEY (tweet_id),
		FOREIGN KEY (tweet_id) REFERENCES tweets(tweet_id) ON DELETE CASCADE
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
	`

	// SQL文の実行
	_, err := db.Exec(createTableSQL)
	if err != nil {
		return err
	}
	return nil
}

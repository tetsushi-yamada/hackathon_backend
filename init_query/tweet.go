package init_query

import "database/sql"

func DropTweetTable(db *sql.DB) error {
	// テーブル削除用のSQL文
	dropTableSQL := `DROP TABLE IF EXISTS tweets;`

	// SQL文の実行
	_, err := db.Exec(dropTableSQL)
	if err != nil {
		return err
	}
	return nil
}

func CreateTweetTable(db *sql.DB) error {
	// テーブル作成用のSQL文
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS tweets (
		tweet_id char(36) not null,
		user_id char(36) not null,
		tweet_text varchar(255) not null,
	    parent_id char(36) default null,
	    retweet_id char(36) default null,
	    is_inappropriate boolean not null default false,
		created_at timestamp not null default current_timestamp,
		updated_at timestamp not null default current_timestamp on update current_timestamp,
		PRIMARY KEY (tweet_id),
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

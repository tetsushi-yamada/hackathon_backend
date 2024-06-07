package init_query

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func StartDB() *sql.DB {
	// DB接続のための準備
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPwd := os.Getenv("MYSQL_PWD")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	connStr := fmt.Sprintf("%s:%s@%s/%s", mysqlUser, mysqlPwd, mysqlHost, mysqlDatabase)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Init_table(db *sql.DB) error {
	if err := DropGoodTable(db); err != nil {
		return err
	}

	if err := DropFollowTable(db); err != nil {
		return err
	}

	if err := DropTweetPictureTable(db); err != nil {
		return err
	}

	if err := DropTweetTable(db); err != nil {
		return err
	}

	if err := DropProfilePictureTable(db); err != nil {
		return err
	}

	if err := DropUserTable(db); err != nil {
		return err
	}

	if err := CreateUserTable(db); err != nil {
		return err
	}

	if err := CreateProfilePictureTable(db); err != nil {
		return err
	}

	if err := CreateTweetTable(db); err != nil {
		return err
	}

	if err := CreateTweetPictureTable(db); err != nil {
		return err
	}

	if err := CreateFollowTable(db); err != nil {
		return err
	}

	if err := CreateGoodTable(db); err != nil {
		return err
	}

	return nil
}

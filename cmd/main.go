package main

import (
	"database/sql"
	"fmt"
	"github.com/tetsushi-yamada/hackathon_backend/internal/server"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	router := server.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
	e.Logger.Fatal(e.Start(":8000"))
}

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

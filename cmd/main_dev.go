package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tetsushi-yamada/hackathon_backend/init_query"
	"github.com/tetsushi-yamada/hackathon_backend/internal/database"
	"github.com/tetsushi-yamada/hackathon_backend/internal/handler"
	"github.com/tetsushi-yamada/hackathon_backend/internal/server"
	"github.com/tetsushi-yamada/hackathon_backend/internal/usecase"
	"log"
	"net/http"
	"os"
)

func main() {

	mysqlUser := os.Getenv("DB_USER")
	mysqlPwd := os.Getenv("DB_PASSWORD")
	mysqlHost := os.Getenv("DB_HOST")
	mysqlDatabase := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s", mysqlUser, mysqlPwd, mysqlHost, mysqlDatabase)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = init_query.Init_table(db)
	if err != nil {
		log.Fatal(err)
	}

	//database層
	userDatabase := database.NewUserDatabase()
	tweetDatabase := database.NewTweetDatabase()
	followDatabase := database.NewFollowDatabase()

	//usecase層
	userUsecase := usecase.NewUserUsecase(db, userDatabase)
	tweetUsecase := usecase.NewTweetUsecase(db, tweetDatabase)
	followUsecase := usecase.NewFollowUsecase(db, followDatabase)
	followerUsecase := usecase.NewFollowerUsecase(db, followDatabase)
	followOrNotUsecase := usecase.NewFollowOrNotUsecase(db, followDatabase)

	//handler層
	userHandler := handler.NewUserHandler(userUsecase)
	tweetHandler := handler.NewTweetHandler(tweetUsecase)
	followHandler := handler.NewFollowHandler(followUsecase)
	followerHandler := handler.NewFollowerHandler(followerUsecase)
	followOrNotHandler := handler.NewFollowOrNotHandler(followOrNotUsecase)

	handlers := handler.Handlers{
		User:        userHandler,
		Tweet:       tweetHandler,
		Follow:      followHandler,
		Follower:    followerHandler,
		FollowOrNot: followOrNotHandler,
	}

	router := server.NewRouter(&handlers)
	err = http.ListenAndServe(":8001", router)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

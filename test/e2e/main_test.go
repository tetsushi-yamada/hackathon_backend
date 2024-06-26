package e2e

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tetsushi-yamada/hackathon_backend/internal/database"
	"github.com/tetsushi-yamada/hackathon_backend/internal/handler"
	"github.com/tetsushi-yamada/hackathon_backend/internal/server"
	"github.com/tetsushi-yamada/hackathon_backend/internal/usecase"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// DB
	mysqlUser := "user"
	mysqlPwd := "password"
	mysqlDatabase := "hackathon_test"

	connStr := fmt.Sprintf("%s:%s@(localhost:3306)/%s", mysqlUser, mysqlPwd, mysqlDatabase)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		os.Exit(1)
	}

	//database層
	userDatabase := database.NewUserDatabase()
	profilePictureDatabase := database.NewProfilePictureDatabase()
	tweetDatabase := database.NewTweetDatabase()
	followDatabase := database.NewFollowDatabase()
	followRequestDatabase := database.NewFollowRequestDatabase()
	goodDatabase := database.NewGoodDatabase()

	//usecase層
	userUsecase := usecase.NewUserUsecase(db, userDatabase)
	profilePictureUsecase := usecase.NewProfilePictureUsecase(db, profilePictureDatabase)
	tweetUsecase := usecase.NewTweetUsecase(db, tweetDatabase)
	followUsecase := usecase.NewFollowUsecase(db, followDatabase, followRequestDatabase, userDatabase)
	followerUsecase := usecase.NewFollowerUsecase(db, followDatabase)
	goodUsecase := usecase.NewGoodUsecase(db, goodDatabase)

	//handler層
	userHandler := handler.NewUserHandler(userUsecase)
	profilePictureHandler := handler.NewProfilePictureHandler(profilePictureUsecase)
	tweetHandler := handler.NewTweetHandler(tweetUsecase)
	followHandler := handler.NewFollowHandler(followUsecase)
	followerHandler := handler.NewFollowerHandler(followerUsecase)
	goodHandler := handler.NewGoodHandler(goodUsecase)

	handlers := handler.Handlers{
		User:           userHandler,
		ProfilePicture: profilePictureHandler,
		Tweet:          tweetHandler,
		Follow:         followHandler,
		Follower:       followerHandler,
		Good:           goodHandler,
	}

	router := server.NewRouter(&handlers)
	go func() {
		if err := http.ListenAndServe(":8000", router); err != nil {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()
	os.Exit(m.Run())
}

package main

import (
	"github.com/tetsushi-yamada/hackathon_backend/init_query"
	"github.com/tetsushi-yamada/hackathon_backend/internal/database"
	"github.com/tetsushi-yamada/hackathon_backend/internal/handler"
	"github.com/tetsushi-yamada/hackathon_backend/internal/server"
	"github.com/tetsushi-yamada/hackathon_backend/internal/usecase"
	"log"
	"net/http"
)

func CORSMiddlewareProd(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// プリフライトリクエストの応答
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// 次のミドルウェアまたはハンドラを呼び出す
		next.ServeHTTP(w, r)
	})
}

func main() {

	db := init_query.StartDB()
	err := init_query.Init_table(db)
	if err != nil {
		log.Fatal(err)
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
	tweetPictureUsecase := usecase.NewTweetPictureUsecase(db, tweetDatabase)
	followUsecase := usecase.NewFollowUsecase(db, followDatabase, followRequestDatabase, userDatabase)
	followerUsecase := usecase.NewFollowerUsecase(db, followDatabase)
	goodUsecase := usecase.NewGoodUsecase(db, goodDatabase)

	//handler層
	userHandler := handler.NewUserHandler(userUsecase)
	profilePictureHandler := handler.NewProfilePictureHandler(profilePictureUsecase)
	tweetHandler := handler.NewTweetHandler(tweetUsecase)
	tweetPictureHandler := handler.NewTweetPictureHandler(tweetPictureUsecase)
	followHandler := handler.NewFollowHandler(followUsecase)
	followerHandler := handler.NewFollowerHandler(followerUsecase)
	goodHandler := handler.NewGoodHandler(goodUsecase)

	handlers := handler.Handlers{
		User:           userHandler,
		ProfilePicture: profilePictureHandler,
		Tweet:          tweetHandler,
		TweetPicture:   tweetPictureHandler,
		Follow:         followHandler,
		Follower:       followerHandler,
		Good:           goodHandler,
	}

	router := server.NewRouter(&handlers)
	corsRouter := CORSMiddlewareProd(router)

	err = http.ListenAndServe(":8080", corsRouter)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

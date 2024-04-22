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

func main() {

	db := init_query.StartDB()
	err := init_query.Init_table(db)
	if err != nil {
		log.Fatal(err)
	}

	//database層
	userDatabase := database.NewUserDatabase()

	//usecase層
	userUsecase := usecase.NewUserUsecase(db, userDatabase)

	//handler層
	userHandler := handler.NewUserHandler(userUsecase)

	handlers := handler.Handlers{
		User: userHandler,
	}

	router := server.NewRouter(&handlers)
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

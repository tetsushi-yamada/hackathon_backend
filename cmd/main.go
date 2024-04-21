package main

import (
	"github.com/tetsushi-yamada/hackathon_backend/init_query"
	"github.com/tetsushi-yamada/hackathon_backend/internal/server"
	"log"
	"net/http"
)

func main() {
	db := init_query.StartDB()

	err := init_query.CreateUserTable(db)
	if err != nil {
		log.Fatal(err)
	}

	err = init_query.CreateTweetTable(db)
	if err != nil {
		log.Fatal(err)
	}

	err = init_query.CreateFollowTable(db)
	if err != nil {
		log.Fatal(err)
	}

	router := server.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

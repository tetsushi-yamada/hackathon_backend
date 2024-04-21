package main

import (
	"github.com/tetsushi-yamada/hackathon_backend/init_query"
	"github.com/tetsushi-yamada/hackathon_backend/internal/server"
	"log"
	"net/http"
)

func main() {
	router := server.NewRouter()
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

	db := init_query.StartDB()
	err = init_query.Init_table(db)
	if err != nil {
		log.Fatal(err)
	}
}

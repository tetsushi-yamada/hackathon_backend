package server

import (
	"github.com/tetsushi-yamada/hackathon_backend/internal/handler"
)

type server struct {
	userHandler     *userHandler
	tweetHandler    *tweetHandler
	followHandler   *followHandler
	followerHandler *followerHandler
}

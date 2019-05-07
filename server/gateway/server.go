package gateway

import (
	"log"
	"net/http"

	"github.com/rs/cors"
)

type Server struct {
	CORS *cors.Cors
}

func NewServer() *Server {
	s := &Server{}

	s.CORS = cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT", "PATCH"},
		AllowCredentials: true,
		MaxAge:           86400,
	})

	return s
}

// TODO: DO SOMETHING!
func (s *Server) ServeHTTP(_ http.ResponseWriter, _ *http.Request) {
	log.Println("Incoming Request")
}

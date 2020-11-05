package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"net/http"
	"os"
)

type server struct {
	router *mux.Router
	logger *zerolog.Logger
}

func newServer() *server{
	logger :=  zerolog.New(os.Stderr).With().Timestamp().Logger()
	return &server{
		router: mux.NewRouter(),
		logger: &logger,
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter,r *http.Request){
	s.router.ServeHTTP(w,r)
}

func (s *server) configureRouter(){

}
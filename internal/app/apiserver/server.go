package apiserver

import (
	"github.com/JetBrainer/BackOCRService/internal/app/model"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"net/http"
	"os"
)

type server struct {
	router *mux.Router
	logger zerolog.Logger
	config Config
}

func newServer() *server{
	// Put Log Level to Debug
	logLevel :=  zerolog.InfoLevel
	logLevel = zerolog.DebugLevel
	zerolog.SetGlobalLevel(logLevel)
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	// Configure Router
	s := &server{
		router: mux.NewRouter(),
		logger: logger,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter,r *http.Request){
	s.router.ServeHTTP(w,r)
}

func (s *server) configureRouter(){
	s.router.HandleFunc("/", s.getDocHandler()).Methods(http.MethodPost)
}

func (s *server) getDocHandler() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse MultiPart File
		err := r.ParseMultipartForm(10 << 20)
		if err != nil{
			http.Error(w,err.Error(),http.StatusBadRequest)
		}
		// Send Request to another Api and get text result
		res, err := s.config.ParseFromLocal(r.Body)
		if err != nil{
			s.logger.Err(err).Msg("Error parsing from Local")
		}

		model.RuleUsageLocal(res)
	}
}
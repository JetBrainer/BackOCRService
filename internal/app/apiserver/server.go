package apiserver

import (
	"github.com/JetBrainer/BackOCRService/internal/app"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"net/http"
	"os"
)

type server struct {
	router  *mux.Router
	logger  zerolog.Logger
	config  *Config
	//doc		*model.DocStruct
}

func newServer(config *Config) *server{
	// Put Log Level to Debug
	logLevel :=  zerolog.InfoLevel
	logLevel = zerolog.DebugLevel
	zerolog.SetGlobalLevel(logLevel)
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	// Configure Router
	s := &server{
		router: mux.NewRouter(),
		logger: logger,
		config: config,
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

		jValue := &OCRText{}
		// Send Request to another Api and get text result
		err = s.config.ParseFromPost(r.Body, jValue)
		if err != nil{
			s.logger.Err(err).Msg("Error parsing from Local")
		}

		app.RuleDocUsage(jValue.JustText())
	}
}
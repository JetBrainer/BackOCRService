package apiserver

import (
	"encoding/json"
	"github.com/JetBrainer/BackOCRService/internal/app"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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
	//logLevel :=  zerolog.InfoLevel
	logLevel := zerolog.DebugLevel
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
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	s.router.HandleFunc("/", s.DocJsonHandler()).Methods(http.MethodPost)
}

func (s *server) getDocHandler() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse MultiPart File
		//err := r.ParseMultipartForm(32 << 20)
		//if err != nil{
		//	http.Error(w,multipart.ErrMessageTooLarge.Error(),http.StatusBadRequest)
		//}

		jValue := &OCRText{}

		// Send Request to another Api and get text result
		err := s.config.ParseFromPost(r, jValue)
		if err != nil{
			s.logger.Err(err).Msg("Error parsing from Local")
			http.Error(w,err.Error(),http.StatusBadRequest)
		}
		defer r.Body.Close()
		//s.logger.Info().Msg(jValue.JustText())
		//val := app.RuleDocUsage(jValue.JustText())
		//s.logger.Info().Msg(val)
		s.logger.Info().Msg(jValue.JustText())
	}
}

func (s *server) DocJsonHandler() http.HandlerFunc{
	type req struct {
		Base string `json:"base"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		jValue := &OCRText{}
		val := &req{}
		err := json.NewDecoder(r.Body).Decode(&val)
		if err != nil{
			log.Info().Msg("Unmarshal error")
		}

		// Send JSON request
		err = s.config.ParseFromBase64(val.Base, jValue)
		if err != nil{
			s.logger.Err(err).Msg("Error parsing from Local")
			http.Error(w,err.Error(),http.StatusBadRequest)
		}
		val1, val2, val3, val4, val5, val6, val7, val8, val9 := app.RuleDocUsage(jValue.JustText())

		log.Info().Msg(val1)
		log.Info().Msg(val2)
		log.Info().Msg(val3)
		log.Info().Msg(val4)
		log.Info().Msg(val5)
		log.Info().Msg(val6)
		log.Info().Msg(val7)
		log.Info().Msg(val8)
		log.Info().Msg(val9)
	}
}
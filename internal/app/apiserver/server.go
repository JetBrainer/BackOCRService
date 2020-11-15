// Package classification of Document API
//
// Documentation for Document API
//
//  Schemes: http
//  BasePath: /
//  Version: 1.0.0
//
//  Consumes:
//  - application/json
//
//  Produces:
//  - application/json
// swagger:meta
package apiserver

import (
	"encoding/json"
	"github.com/JetBrainer/BackOCRService/internal/app"
	"github.com/go-openapi/runtime/middleware"
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
	s.router.HandleFunc("/image", s.DocJsonHandler()).Methods(http.MethodPost)
	s.router.HandleFunc("/form",s.getDocPartFormHandler()).Methods(http.MethodPost)

	ops := middleware.RedocOpts{SpecURL: "/api/v1/swagger.yaml"}
	sh := middleware.Redoc(ops, nil)
	s.router.Handle("/docs",sh)
	s.router.Handle("/api/v1/swagger.yaml",http.FileServer(http.Dir("./")))
}

func (s *server) getDocPartFormHandler() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
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

// swagger:route POST /image Image
// Returns particular document field
// responses
//	200: docStrRepoResp
func (s *server) DocJsonHandler() http.HandlerFunc{
	// Our base64 document
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

		// Document structure and we parse text to it
		docStruct := app.DocStr{}
		docStruct.RuleDocUsage(jValue.JustText())

		err = json.NewEncoder(w).Encode(&docStruct)
		if err != nil{
			s.logger.Print(err)
			s.logger.Info().Msg("error parsing json")
		}

		w.WriteHeader(http.StatusOK)
	}
}
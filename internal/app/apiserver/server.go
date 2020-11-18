// Package classification of Document API
//
// Documentation for Enterprise Intelligent Character Recognition API
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
	"github.com/JetBrainer/BackOCRService/internal/app/store"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"html/template"
	"net/http"
	"os"
)

type server struct {
	router  *mux.Router
	logger  *zerolog.Logger
	config  *Config
	store 	store.Store
}

func newServer(store store.Store, config *Config) *server{
	// Put Log Level to Debug

	logLevel := zerolog.DebugLevel
	zerolog.SetGlobalLevel(logLevel)
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	// Configure Router
	s := &server{
		router: mux.NewRouter(),
		logger: &logger,
		config: config,
		store: 	store,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter,r *http.Request){
	s.router.ServeHTTP(w,r)
}

func (s *server) configureRouter(){

	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))

	s.router.HandleFunc("/register",s.createUserHandler()).Methods(http.MethodPost)
	s.router.HandleFunc("/image", s.docJsonHandler()).Methods(http.MethodPost)
	s.router.HandleFunc("/form",s.getDocPartFormHandler()).Methods(http.MethodPost)

	// Openapi 2.0 spec generation handler
	ops := middleware.RedocOpts{SpecURL: "/api/v1/swagger.yaml"}
	sh := middleware.Redoc(ops, nil)
	s.router.Handle("/docs",sh)
	s.router.Handle("/api/v1/swagger.yaml",http.FileServer(http.Dir("./")))

	s.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("static/menu.html"))
		tmpl.Execute(w,nil)
	})
	s.router.PathPrefix("/").
		Handler(http.FileServer(http.Dir("./static")))
}

func (s *server) getDocPartFormHandler() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		jValue := &OCRText{}

		// Send Request to another Api and get text result
		err = s.config.ParseFromPost(r.Body, jValue)
		if err != nil{
			s.logger.Err(err).Msg("Error parsing from Local")
			http.Error(w,err.Error(),http.StatusBadRequest)
		}
		defer r.Body.Close()
		s.logger.Info().Msg(jValue.JustText())

		w.WriteHeader(http.StatusOK)
	}
}

// swagger:route POST /image Image docRequest
// Returns particular document field
//
// Document return
//
// Client sends Full Scanned Document and get's every need field
//
// responses:
//	200: docResponse

// Get JSON base64 image
func (s *server) docJsonHandler() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		jValue := &OCRText{}
		val := &req{}
		err := json.NewDecoder(r.Body).Decode(&val)
		if err != nil{
			log.Info().Msg("Unmarshal error")
			return
		}

		err = s.store.User().CheckToken(val.Token)
		if err != nil{
			http.Error(w,"Invalid Token",http.StatusBadRequest)
			return
		}

		// Send JSON request
		err = s.config.ParseFromBase64(val.Base, jValue)
		if err != nil{
			s.logger.Err(err).Msg("Error parsing from Local")
			http.Error(w,err.Error(),http.StatusUnprocessableEntity)
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

// Your Invoice structure
// swagger:response docResponse
type docResponse struct {
	// recognized fields
	// in: body
	Body []app.DocStr
}

// Our base64 document
type req struct {
	Token string `json:"token"`
	Base  string `json:"base64"`
}

// Get data for you
// swagger:parameters docRequest
type docRequest struct {
	// Need data
	// in: body
	// required: true
	Body req
}

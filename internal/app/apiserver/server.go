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
	"fmt"
	"html/template"
	"net/http"
	"os"

	client2 "github.com/Somatic-KZ/sso-client"
	"github.com/Somatic-KZ/sso-client/client"
	"github.com/go-chi/chi"
	middlware2 "github.com/go-chi/chi/middleware"

	"github.com/JetBrainer/BackOCRService/internal/app"
	"github.com/JetBrainer/BackOCRService/internal/app/store"

	"github.com/go-chi/jwtauth"
	"github.com/go-openapi/runtime/middleware"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type server struct {
	router *chi.Mux
	logger *zerolog.Logger
	config *Config
	store  store.Store
	ssoCli client.SSO
}

func newServer(store store.Store, config *Config) (*server, error) {
	// Put Log Level to Debug

	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	cli, err := client2.NewSSO(&client2.Config{Addr: config.GRPCPort, Protocol: "grpc"})
	if err != nil {
		return nil, err
	}

	// Configure Router
	s := &server{
		router: chi.NewRouter(),
		logger: &logger,
		config: config,
		store:  store,
		ssoCli: cli,
	}

	s.configureRouter()

	return s, nil
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {

	auth := NewAuthenticator([]byte(s.config.JWTKey))
	s.router.Use(middlware2.NoCache) // no-cache
	//r.Use(middleware.RequestID) // вставляет request ID в контекст каждого запроса
	s.router.Use(middlware2.Logger)    // логирует начало и окончание каждого запроса с указанием времени обработки
	s.router.Use(middlware2.Recoverer) // управляемо обрабатывает паники и выдает stack trace при их возникновении
	s.router.Use(middlware2.RealIP)    // устанавливает RemoteAddr для каждого запроса с заголовками X-Forwarded-For или X-Real-IP
	s.router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}).Handler)

	s.router.Use(jwtauth.Verifier(auth.TokenAuth()))
	s.router.Use(NewUserAccessCtx([]byte(s.config.JWTKey)).ChiMiddleware)

	s.router.Post("/image", s.docJsonHandler())
	s.router.Post("/form", s.getDocPartFormHandler())

	// Openapi 2.0 spec generation handler
	ops := middleware.RedocOpts{SpecURL: "/api/v1/swagger.yaml"}
	sh := middleware.Redoc(ops, nil)
	s.router.Handle("/docs", sh)
	s.router.Handle("/api/v1/swagger.yaml", http.FileServer(http.Dir("./")))

	s.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("static/menu.html"))
		if err := tmpl.Execute(w, nil); err != nil {
			s.logger.Err(err).Msg("Error loading template")
		}
	})

}

func (s *server) getDocPartFormHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		jValue := &OCRText{}

		// Send Request to another Api and get text result
		err = s.config.ParseFromPost(r.Body, jValue)
		if err != nil {
			s.logger.Err(err).Msg("Error parsing from Local")
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		defer r.Body.Close()
		s.logger.Info().Msg(jValue.JustText())

		w.WriteHeader(http.StatusOK)
	}
}

// swagger:route POST /image Document docRequest
// Returns particular document field
//
// Document return
//
// Client sends Full Scanned Document and get's every need field
//
// responses:
//	200: docResponse

// Get JSON base64 image
func (s *server) docJsonHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id, ok := ctx.Value("tdid").(string)
		if !ok {
			http.Error(w, "no id", http.StatusUnauthorized)
			return
		}

		jValue := &OCRText{}
		val := &req{}
		err := json.NewDecoder(r.Body).Decode(&val)
		if err != nil {
			log.Info().Msg("Unmarshal error")
			http.Error(w,err.Error(), http.StatusBadRequest)
			return
		}

		if val.Base == "" {
			log.Info().Msg("Not Valid Image")
			http.Error(w,err.Error(), http.StatusBadRequest)
		}

		err = s.ssoCli.UserToken(ctx, id)
		if err != nil {
			http.Error(w, client.ErrTokenNotFound.Error(), http.StatusNoContent)
			return
		}

		// Send JSON request
		err = s.config.ParseFromBase64(val.Base, jValue)
		if err != nil {
			s.logger.Err(err).Msg("Error parsing from Local")
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		// Document structure and we parse text to it
		docType := app.GetMultiplexer(jValue.JustText())
		if docType == nil {
			s.logger.Err(err).Msg("No recognized document")
			http.Error(w, "NOT RECOGNIZED", http.StatusUnprocessableEntity)
			return
		}
		docType.RuleDocUsage(jValue.JustText())

		fmt.Println(jValue.JustText())
		err = json.NewEncoder(w).Encode(&docType)
		if err != nil {
			s.logger.Print(err)
			s.logger.Info().Msg("error parsing json")
			http.Error(w, "ERR JSON", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// Your Invoice structure
// swagger:response docResponse
type docResponse struct {
	// recognized fields
	// in: body
	Body []interface{}
}

// Our base64 document
type req struct {
	Base string `json:"base64"`
}

// Get data, put your token in header
// swagger:parameters docRequest
type docRequest struct {
	// Need data
	// in: body
	// required: true
	Body req
}

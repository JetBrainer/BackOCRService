package apiserver

import (
	"encoding/json"
	"github.com/JetBrainer/BackOCRService/internal/app/model"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (s *server) createUserHandler() http.HandlerFunc{
	type request struct {
		Email 			string `json:"email"`
		Password		string `json:"password"`
		Organization	string `json:"organization"`
	}
	type response struct {
		ID		int 	`json:"id"`
		Token	string 	`json:"token"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil{
			http.Error(w,"Error parsing json",http.StatusBadRequest)
			return
		}
		u := &model.User{
			Email: req.Email,
			Password: req.Password,
			Organization: req.Organization,
		}

		if err := s.store.User().Create(u); err != nil{
			log.Print(err)
			http.Error(w,"Error adding values",http.StatusUnprocessableEntity)
			return
		}
		resp := &response{
			ID: u.ID,
			Token: u.Token,
		}

		u.Sanitize()
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(resp); err != nil{
			s.logger.Info().Msg("Error JSON Encode")
		}
	}
}

func (s *server) handleUserDelete() http.HandlerFunc{
	type request struct {
		Email string `json:"email"`
	}
	return func(w http.ResponseWriter,r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil{
			http.Error(w,"Unable Decode JSON",http.StatusBadRequest)
			return
		}

		if err := s.store.User().DeleteUser(req.Email); err != nil{
			s.logger.Info().Msg("Unable Delete user")
			http.Error(w,"Incorrect Email Or Password",http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
package apiserver

import (
	"encoding/json"
	"github.com/JetBrainer/BackOCRService/internal/app/model"
	"net/http"
)

func (s *server) createUserHandler() http.HandlerFunc{
	type request struct {
		Email 			string `json:"email"`
		Password		string `json:"password"`
		Organization	string `json:"organization"`
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
			http.Error(w,"Error adding values",http.StatusUnprocessableEntity)
			return
		}

		u.Sanitize()
		w.WriteHeader(http.StatusOK)
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
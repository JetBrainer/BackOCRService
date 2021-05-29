package apiserver

import (
	"context"
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
)

var ErrInvalidToken = errors.New("token is incorrect or expired")

type UserAccessCtx struct {
	jwtKey []byte
}

// Claims структура, хранящая закодированный JWT авторизации.
// Встраивание для предоставления поля expiry.
type Claims struct {
	TDID      string   `json:"tdid"`
	IsRefresh bool     `json:"is_refresh"`
	Roles     []string `json:"roles"`
	jwt.StandardClaims
}

func NewUserAccessCtx(jwtKey []byte) *UserAccessCtx {
	return &UserAccessCtx{
		jwtKey: jwtKey,
	}
}

func (ua UserAccessCtx) ChiMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _, err := jwtauth.FromContext(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if token == nil || !token.Valid {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// инициализируем новый инстанс Claims
		claims := new(Claims)
		tkn, err := jwt.ParseWithClaims(token.Raw, claims, func(token *jwt.Token) (interface{}, error) {
			return ua.jwtKey, nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if claims.IsRefresh {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "tdid", claims.TDID)

		// токен валидный, пропускаем его
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

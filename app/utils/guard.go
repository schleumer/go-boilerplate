package utils

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"../database"
	"../models"
	"../system"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

type GuardedRouter struct {
	MuxRouter *mux.Router
}

func NewGuardedRouter(router *mux.Router) GuardedRouter {
	return GuardedRouter{router}
}

type SimpleRoute func(rs http.ResponseWriter, rq *http.Request)

type GuardedRoute func(user models.User, rs http.ResponseWriter, rq *http.Request)

func (r *GuardedRouter) ProtectFunc(path string, fn GuardedRoute, methods ...string) *mux.Route {
	if len(methods) > 0 {
		return r.MuxRouter.HandleFunc(path, Protect(fn)).Methods(methods...)
	}

	return r.MuxRouter.HandleFunc(path, Protect(fn))
}

func (r *GuardedRouter) Func(path string, fn SimpleRoute, methods ...string) *mux.Route {
	if len(methods) > 0 {
		return r.MuxRouter.HandleFunc(path, fn).Methods(methods...)
	}

	return r.MuxRouter.HandleFunc(path, fn)
}

type Claim struct {
	uid float64
}

func ParseToken(token string) (Claim, error) {
	result, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return system.JwtSecret, nil
	})

	if err != nil {
		return Claim{}, err
	}

	if claims, ok := result.Claims.(jwt.MapClaims); ok && result.Valid {
		return Claim{claims["sub"].(float64)}, nil

	} else {
		return Claim{}, err
	}

	return Claim{}, err
}

func Protect(fn GuardedRoute) SimpleRoute {
	return func(rs http.ResponseWriter, rq *http.Request) {
		var token = rq.Header.Get("Authorization")
		r, _ := regexp.Compile("^Bearer .*")

		if token == "" {
			token = rq.FormValue("token")
		} else if r.Match([]byte(token)) {
			token = strings.SplitN(token, " ", 2)[1]
		} else {
			token = ""
		}

		if token == "" {
			rs.WriteHeader(http.StatusBadRequest)
			rs.Write([]byte("Invalid request, Authorization not found."))
			return
		}

		claim, err := ParseToken(token)

		if err != nil {
			rs.WriteHeader(http.StatusBadRequest)
			rs.Write([]byte("Invalid request, invalid token."))
			return
		}

		count := 0
		database.LocalDb.Table("users").Where("id = ?", claim.uid).Count(&count)

		if count < 1 {
			rs.WriteHeader(http.StatusBadRequest)
			rs.Write([]byte("Invalid request, Authorization not found."))
			return
		}

		user := models.User{}

		database.LocalDb.First(&user, "id = ?", claim.uid).Count(&count)

		fn(user, rs, rq)
	}
}

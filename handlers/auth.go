package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/era-n/taskbackdev/token"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	guid := r.URL.Query().Get("guid")

	w.Header().Set("Content-Type", "application/json")

	response, err := token.NewPairOfToken(guid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(response)

	w.Write(resp)
}

func RefreshMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		guid := r.URL.Query().Get("guid")
		refreshToken := r.URL.Query().Get("token")

		err := token.ValidateRefreshToken(guid, refreshToken)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(w, r)
	})
}

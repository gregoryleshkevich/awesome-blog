package middleware

import (
	"awesome-blog/internal/services"
	"encoding/json"
	"net/http"
	"strings"
)

type jwtError struct {
	Message string
}

func JWTVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := strings.TrimSpace(r.Header.Get("Authorization"))

		if header == "" {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(jwtError{Message: "Missing auth token"})
			return
		}

		err := services.JWTService.Verify(header)

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(jwtError{Message: err.Error()})
			return
		}

		next.ServeHTTP(w, r)
	})
}

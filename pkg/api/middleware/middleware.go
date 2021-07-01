package middleware

import (
	"net/http"

	"github.com/tonymontanapaffpaff/golang-training-university/pkg/data/auth"

	log "github.com/sirupsen/logrus"
)

func TokenAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			log.Errorf("failed to authorize: %s", err)
			w.WriteHeader(http.StatusUnauthorized)
			_, err = w.Write([]byte("unauthorized"))
			if err != nil {
				log.Error(err)
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
		next.ServeHTTP(w, r)
	})
}

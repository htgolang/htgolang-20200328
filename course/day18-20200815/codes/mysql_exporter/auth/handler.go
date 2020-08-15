package auth

import (
	"mysql_exporter/config"
	"net/http"
)

func BasicAuth(config *config.AuthConfig, handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		secret := r.Header.Get("authorization")
		// 认证
		if isAuth(secret, config) {
			handler.ServeHTTP(w, r)
		} else {
			w.Header().Set("www-authenticate", `basic realm="my site"`)
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}

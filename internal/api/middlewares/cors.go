package middlewares

import "net/http"

var allowedOrigins = []string{
	"https://localhost:3000",
	"https://myproject.com",
}

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		origin := r.Header.Get("Origin")

		if !isOriginAllow(origin) {
			http.Error(w, "Now allow by CORS", http.StatusForbidden)
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Max-Age", "3600")

		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func isOriginAllow(origin string) bool {
	for _, allowOrigin := range allowedOrigins {
		if origin == allowOrigin {
			return true
		}
	}

	return false
}

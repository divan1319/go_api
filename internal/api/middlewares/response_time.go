package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func ResponseTime(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received request in ResponseTime")
		start := time.Now()
		//create custom responsewriter
		wrappedWriter := &responseWritter{ResponseWriter: w, status: http.StatusOK}

		next.ServeHTTP(wrappedWriter, r)
		//calculate the duration
		duration := time.Since(start)

		//Log the request

		fmt.Printf("Method: %s, URL: %s, Status: %d, Duration:%v\n ", r.Method, r.URL, wrappedWriter.status, duration.String())
		fmt.Println("Sent resposne in ResponseTime")

	})
}

type responseWritter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWritter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

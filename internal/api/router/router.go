package router

import (
	"goapi/internal/api/handlers"
	"net/http"
)

func Router() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Route root"))
	})

	mux.HandleFunc("/teachers", handlers.TeachersHandlers)

	mux.HandleFunc("/students", handlers.StudentsHandlers)

	return mux
}

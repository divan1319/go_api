package handlers

import "net/http"

func StudentsHandlers(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("students root"))
}

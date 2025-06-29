package utils

import "net/http"

type Middleware func(http.Handler) http.Handler

/*
applyMiddlewares is a function that applies a list of middlewares to a handler
*/
func ApplyMiddlewares(handler http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}

	return handler
}

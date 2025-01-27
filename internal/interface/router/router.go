package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router interface {
	GET(uri string, f func(w http.ResponseWriter, r *http.Request), middlewares ...mux.MiddlewareFunc)
	POST(uri string, f func(w http.ResponseWriter, r *http.Request), middlewares ...mux.MiddlewareFunc)
	PUT(uri string, f func(w http.ResponseWriter, r *http.Request), middlewares ...mux.MiddlewareFunc)
	DELETE(uri string, f func(w http.ResponseWriter, r *http.Request), middlewares ...mux.MiddlewareFunc)
	SERVE(port string)
	Mux() *mux.Router
}

package router

import (
	"net/http"
	"product-service/package/logging"
	"product-service/package/middleware"

	"github.com/gorilla/mux"
)

var (
	muxDispatcher = mux.NewRouter()
)

type muxRouter struct{}

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request), middlewares ...mux.MiddlewareFunc) {
	if middlewares != nil {
		HandleWithMiddleware(uri, "GET", f, middlewares...)
	} else {
		muxDispatcher.HandleFunc(uri, f).Methods("GET")
	}

}
func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request), middlewares ...mux.MiddlewareFunc) {
	if middlewares != nil {
		HandleWithMiddleware(uri, "POST", f, middlewares...)
	} else {
		muxDispatcher.HandleFunc(uri, f).Methods("POST")
	}
}

func (*muxRouter) PUT(uri string, f func(w http.ResponseWriter, r *http.Request), middlewares ...mux.MiddlewareFunc) {
	if middlewares != nil {
		HandleWithMiddleware(uri, "PUT", f, middlewares...)
	} else {
		muxDispatcher.HandleFunc(uri, f).Methods("PUT")
	}
}

func (*muxRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request), middlewares ...mux.MiddlewareFunc) {
	if middlewares != nil {
		HandleWithMiddleware(uri, "DELETE", f, middlewares...)
	} else {
		muxDispatcher.HandleFunc(uri, f).Methods("DELETE")
	}
}
func (*muxRouter) SERVE(port string) {
	logging.Log.Infof("Http server listen in port %s", port)
	muxDispatcher.Use(middleware.LoggingMiddleware)
	http.ListenAndServe(port, muxDispatcher)
}

func HandleWithMiddleware(uri string, method string, f func(w http.ResponseWriter, r *http.Request), middlewares ...mux.MiddlewareFunc) {
	subRouter := muxDispatcher.PathPrefix(uri).Subrouter()
	subRouter.Use(middlewares...)
	subRouter.HandleFunc("", f).Methods(method)
}

func (m *muxRouter) Mux() *mux.Router {
	return muxDispatcher
}

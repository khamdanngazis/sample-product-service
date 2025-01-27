package handler

import (
	"net/http"
	"product-service/package/middleware"
)

type Pinghandler struct {
}

func NewPinghandler() *Pinghandler {
	return &Pinghandler{}
}

func (h *Pinghandler) Ping(w http.ResponseWriter, r *http.Request) {

	middleware.WriteResponse(w, http.StatusOK, "", "Pong")
}

package routes

import (
	"github.com/go-chi/chi"
	"github.com/googollee/go-socket.io"
)

// CreateRouterWS ...
func CreateRouterWS(r *chi.Router, server *socketio.Server) {
	(*r).Handle("/ws/", server)
}

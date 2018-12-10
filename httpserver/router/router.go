package router

import (
	"github.com/go-chi/chi"
	"github.com/googollee/go-socket.io"
	"github.com/wa-gtwy/httpserver/routes"
)

// InitializeRouter ...
func InitializeRouter(server *socketio.Server) *chi.Mux {
	rootRouter := chi.NewRouter()
	rootRouter.Route("/v1", func(r chi.Router) {
		routes.CreateRouterWS(&r, server)
	})
	return rootRouter
}

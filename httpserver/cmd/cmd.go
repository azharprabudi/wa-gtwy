package cmd

import (
	"net/http"

	"github.com/googollee/go-socket.io"

	"github.com/wa-gtwy/httpserver/router"
)

// StartServer ...
func StartServer(server *socketio.Server) {
	r := router.InitializeRouter(server)
	http.ListenAndServe(":5000", r)
}

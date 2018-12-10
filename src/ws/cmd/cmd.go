package ws

import (
	"github.com/googollee/go-socket.io"
	"github.com/wa-gtwy/src/whatsapp/service"
	"github.com/wa-gtwy/src/ws/listener"
)

// InitializeWS ...
func InitializeWS(ws waservice.WhatsappServiceInterface) (*socketio.Server, error) {
	server, err := socketio.NewServer(nil)
	if err != nil {
		return nil, err
	}

	wslistener.CreateListenerWS(server, ws)
	return server, nil
}

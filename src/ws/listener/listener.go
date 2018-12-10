package wslistener

import (
	"log"

	"github.com/googollee/go-socket.io"
	"github.com/wa-gtwy/src/whatsapp/service"
)

// CreateListenerWS ...
func CreateListenerWS(server *socketio.Server, wa waservice.WhatsappServiceInterface) {

	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection")

		so.On("test", func() {
			print(1)
		})

		so.On("sendmessage", func(phoneNumber string, message string) {
			print(phoneNumber)
			print(message)
			wa.SendTextMessagePersonal("6281398420279", "test ndel")
		})
	})

	server.On("sendmessage", func(phoneNumber string, message string) {
		print(phoneNumber)
		print(message)
		wa.SendTextMessagePersonal("6281398420279", "test ndel")
	})

	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

}

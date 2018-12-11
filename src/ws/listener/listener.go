package wslistener

import (
	"encoding/json"
	"log"

	socketio "github.com/googollee/go-socket.io"
	waservice "github.com/wa-gtwy/src/whatsapp/service"
)

// CreateListenerWS ...
func CreateListenerWS(server *socketio.Server, wa waservice.WhatsappServiceInterface) {

	server.On("connection", func(so socketio.Socket) {
		print("another client has been connected with us")
		so.On("chat message", func(message string) {
			model := &struct {
				PhoneNumber string `json:"phoneNumber"`
				Message     string `json:"message"`
			}{}
			err := json.Unmarshal([]byte(message), model)
			if err != nil {
				panic(err)
			}

			wa.SendTextMessagePersonal(model.PhoneNumber, model.Message)
		})

		so.On("test", func() {
			print("helo helo")
		})

		so.On("ahah", func(test interface{}) {
			print("ih ih")
		})
	})

	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

}

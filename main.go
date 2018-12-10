package main

import (
	"github.com/wa-gtwy/httpserver/cmd"
	"github.com/wa-gtwy/src/whatsapp/cmd"
	ws "github.com/wa-gtwy/src/ws/cmd"
)

func main() {
	wa, err := wa.InitializeWA()
	if err != nil {
		panic(err)
	}

	wsServer, err := ws.InitializeWS(wa)
	if err != nil {
		panic(err)
	}

	cmd.StartServer(wsServer)
}

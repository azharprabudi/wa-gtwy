package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Baozisoftware/qrcode-terminal-go"
	whatsapp "github.com/Rhymen/go-whatsapp"
)

func main() {
	wac, err := whatsapp.NewConn(10000 * time.Second)

	if err != nil {
		panic(err)
	}

	qr := make(chan string)
	go func() {
		terminal := qrcodeTerminal.New()
		terminal.Get(<-qr).Print()
	}()

	sess, err := wac.Login(qr)
	if err != nil {
		panic(err)
	}

	<-time.After(3 * time.Second)

	msg := whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: "6281398420279@s.whatsapp.net",
		},
		Text: "Message sent by github.com/Rhymen/go-whatsapp",
	}

	err = wac.Send(msg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error sending message: %v", err)
	}

	fmt.Println("helo", sess)
}

package main

import (
	"fmt"
	"whatsapi/messages"
	"whatsapi/whats"
)

func main() {

	msg := messages.GetMessage()

	for _, message := range msg {
		fmt.Println(messages.FormatMessage(message))
		fmt.Println("------------------------")
		whats.SendWhatsapp("+554299146456", messages.FormatMessage(message))
		
	}

	
}

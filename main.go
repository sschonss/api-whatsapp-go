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
		whats.SendWhatsapp("NUMBER_PHONE", messages.FormatMessage(message))
		
	}

	
}

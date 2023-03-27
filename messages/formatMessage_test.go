package messages_test

import (
	"fmt"
	"testing"
	"whatsapi/messages"
)

func TestFormatMessage(t *testing.T) {
	t.Run("FormatMessage", func(t *testing.T) {
		msg := messages.GetMessage()
		for _, message := range msg {
			fmt.Println(messages.FormatMessage(message))
			fmt.Println("------------------------")
		}
	})
}


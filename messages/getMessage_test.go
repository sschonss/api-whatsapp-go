package messages_test

import (
	"testing"
	"whatsapi/messages"
)

func TestGetMessage(t *testing.T) {
	t.Run("GetMessage", func(t *testing.T) {
		msg := messages.GetMessage()
		if len(msg) == 0 {
			t.Error("Error getting messages")
		}
	})
}


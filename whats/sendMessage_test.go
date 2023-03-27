package whats_test

import (
	"testing"
	"whatsapi/whats"
)

func TestSendWhatsapp(t *testing.T) {
	t.Run("SendWhatsapp", func(t *testing.T) {
		whats.SendWhatsapp("+554299146456", "Teste")
	})
}
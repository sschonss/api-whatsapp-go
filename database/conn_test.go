package database_test

import (
	"testing"
	"whatsapi/database"
)

func TestConnect(t *testing.T) {
	t.Run("Connect", func(t *testing.T) {
		test := database.Connect()
		if test == nil {
			t.Error("Error connecting to database")
		}
	})
}

func TestClose(t *testing.T) {
	t.Run("Close", func(t *testing.T) {
		test := database.Connect()
		database.Close(test)
	})
}


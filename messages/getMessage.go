package messages

import (
	"fmt"
	"whatsapi/database"
)

type Message struct {
	ID int
	Start string
	End string
	Sensor string
	Content string
}

// GetMessage get message from database
func GetMessage() []Message {
	db := database.Connect()

	var messages []Message

	// make query select from database and return many rows or one row
	rows, err := db.Query("SELECT id, Al_Start_Time, Al_Norm_Time, Al_Tag, Al_Message FROM tb_alarmes WHERE Al_Active = 1")
	if err != nil {
		fmt.Println("Error query: ", err)
	}

	// loop rows
	for rows.Next() {
		var message Message
		err = rows.Scan(&message.ID, &message.Start, &message.End, &message.Sensor, &message.Content)
		if err != nil {
			fmt.Println("Error scan: ", err)
		}
		messages = append(messages, message)
	}

	return messages
	
}
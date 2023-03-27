package messages

import (
	"fmt"
)

// formatMessage format message to send to whatsapp
func FormatMessage(message Message) string {
	id_alarm := message.ID
	start := message.Start
	end := message.End
	sensor := message.Sensor
	content := message.Content

	return fmt.Sprintf("ID: %d\nStart: %s\nEnd: %s\nSensor: %s\nContent: %s", id_alarm, start, end, sensor, content)

}

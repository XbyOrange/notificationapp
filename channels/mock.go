package channels

import (
	"github.com/OrangeB4B/notificationapp/db"
	"strconv"
	"time"
	"errors"
	"github.com/OrangeB4B/notificationapp/config"
	"fmt"
)

type EventForFake struct {
	TriggeredEvent db.Event
}

//ParseTemplate : Parsing Template for SMS
func (event EventForFake) ParseTemplate() ([]db.Message, error) {
	if event.TriggeredEvent.TemplateType == "FAKE"{
		var messages []db.Message
		message := db.Message{}
		message.Recipient = "FAKE-RECIPIENT"
		message.Subject = event.TriggeredEvent.Subject
		message.Channel = event.TriggeredEvent.AccountID + "FAKE"
		message.DateCreated = time.Now()
		message.TemplateId = event.TriggeredEvent.AccountID
		message.Content = "FAKE"
		message.CustomData = event.TriggeredEvent.CustomData
		message.MessageID = strconv.Itoa(time.Now().Nanosecond())+ event.TriggeredEvent.AccountID
		messages = append(messages, message)
		fmt.Println("Template FAKE")
		return messages,nil
	}else{
		return nil, errors.New("ERROR")
	}
}

//SendMessage : Message Sending  for SMS
func (event EventForFake) SendMessage(message db.Message) db.MessageResponse {
	if message.Subject == "FAKE-SUBJECT"{
		return db.MessageResponse{Status: config.SUCCESS,}
	}
	return db.MessageResponse{Status:config.FAILED}

}



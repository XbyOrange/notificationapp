package service

import (
	"github.com/alknopfler/notificationapp/db"
	"github.com/alknopfler/notificationapp/channels"
	"github.com/alknopfler/notificationapp/logger"
	"github.com/alknopfler/notificationapp/config"
	"errors"
)

// EventProcessorForChannel : Event Processor For Channel
func EventProcessorForChannel(event db.Event) ([]db.Message, error){
		var errorSMS,errorEmail, errorFake error
		var messages,messagesSMS,messagesEmail,messagesFake []db.Message

			//The event could have different channels at the same time. Could be possible send N notitifications (email, sms, etc...) at the same time
		if channels.CheckChannel(event, db.SMS) {
			logger.Print("INFO",config.SrvName,config.ComponentName,"1","200","Processing " + event.AccountID+ " for SMS")
			smsChannel := channels.EventForSMS{event}
			messagesSMS,errorSMS=ProcessEvent(smsChannel)
			messages = append(messages,messagesSMS...)
		}
		if channels.CheckChannel(event, db.EMAIL) {
			logger.Print("INFO",config.SrvName,config.ComponentName,"1","200","Processing " + event.AccountID+ " for EMAIL")
			emailChannel := channels.EventForEmail{event}
			messagesEmail,errorEmail=ProcessEvent(emailChannel)
			messages = append(messages,messagesEmail...)
		}
		if channels.CheckChannel(event, db.FAKE) {
			fakeChannel := channels.EventForFake{event}
			messagesFake,errorFake=ProcessEvent(fakeChannel)
			messages = append(messages,messagesFake...)
		}
		if errorSMS!=nil || errorEmail!=nil || errorFake != nil {
			logger.Print("ERROR",config.SrvName,config.ComponentName,"1","500","Error sending the message")
			return nil,errors.New("Error Sending the Message")
		}
		return messages,nil
}


// ProcessEvent : Process Event is agnostic for the type of channel because are implemented by the interface elements
func ProcessEvent(eventForMessage channels.EventForMessage) ([]db.Message,error) {
	messages, err := eventForMessage.ParseTemplate()
	if err != nil {
		logger.Print("ERROR",config.SrvName,config.ComponentName,"1","500","Error parsing template Error :" + err.Error() + "", err)
		return nil,err
	} else {
		var err2 error
		for _, msg := range messages {

			response := eventForMessage.SendMessage(msg)
			if response.Status == config.FAILED {
				err2=errors.New("Send Message Failed")
				logger.Print("ERROR",config.SrvName,config.ComponentName,"1","500","Error sending message in Proccess Event")
			}
		}
		return messages,err2
	}
}
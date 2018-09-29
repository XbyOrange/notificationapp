package channels

import (
	"gopkg.in/gomail.v2"
	"github.com/alknopfler/notificationapp/config"
	"github.com/alknopfler/notificationapp/db"
	"errors"
	"github.com/alknopfler/notificationapp/templates"
	"time"
	"strconv"
	"github.com/smancke/mailck"
	"github.com/alknopfler/notificationapp/logger"
)




//EventForEmail : Email implementation for SMS
type EventForEmail struct {
	TriggeredEvent db.Event
}

//ParseTemplate : Template Parser Implementation for Email
func (event EventForEmail) ParseTemplate() ([]db.Message, error) {
	var messages []db.Message
	channelSupported := CheckChannel(event.TriggeredEvent, db.EMAIL)
	if !channelSupported {
		logger.Print("ERROR",config.SrvName,config.Channel,"1","417","Dropping event ['" + event.TriggeredEvent.AccountID+ "']. EMAIL channel not supported.")
		return messages, errors.New("EMAIL channel not supported")
	}
	logger.Print("INFO",config.SrvName,config.Channel,"1","200","Channel supported...")
	numOfRecipient := len(event.TriggeredEvent.Recipient)
	if numOfRecipient <= 0 {
		logger.Print("ERROR",config.SrvName,config.Channel,"1","405","Dropping event ['" + event.TriggeredEvent.AccountID+ "']. No recipient found.")
		return messages, errors.New("no recipients found")
	}
	emailContent, err := templates.ParseTemplateForMessage(event.TriggeredEvent, db.EMAIL, event.TriggeredEvent.Lang)
	if err!=nil{
		logger.Print("ERROR",config.SrvName,config.Channel,"1","500","Error parsing the template . ", err)
		return messages, errors.New("Error parsing the template")
	}
	//parse each mail separately because it may vary by recipient
	for _, em := range event.TriggeredEvent.Recipient {
		if validateEmail(em) {
			logger.Print("INFO",config.SrvName,config.Channel,"1","200","Email: "+ em+" format is OK...creating the message")
			dateCreated := time.Now()
			message := db.Message{}
			message.Recipient = em
			message.Subject = event.TriggeredEvent.Subject
			message.Channel = "EMAIL"
			message.DateCreated = dateCreated
			message.TemplateId = event.TriggeredEvent.TemplateType +"-EMAIL-"+ event.TriggeredEvent.Lang
			message.Content = emailContent
			message.CustomData = event.TriggeredEvent.CustomData
			message.MessageID = strconv.Itoa(dateCreated.Nanosecond()) + em + event.TriggeredEvent.AccountID
			messages = append(messages, message)
			logger.Print("INFO",config.SrvName,config.Channel,"1","200","Message created and added to the list of message to send to recipient")
		}
	}
	return messages, nil
}

//SendMessage : Messaging Sending for Email
func (event EventForEmail) SendMessage(message db.Message) db.MessageResponse {
	var smtpDialer = gomail.NewPlainDialer(
		config.EmailSMTPHost,
		config.EmailSMTPPort,
		config.EmailSMTPUser,
		config.EmailSMTPPass)

	if message.Content == "" {
		logger.Print("ERROR",config.SrvName,config.Channel,"1","400","Sending  Failed. Message body empty")
		return db.MessageResponse{Status: config.FAILED, Response: "MESSAGE EMPTY", TimeOfResponse: time.Now()}
	}
	logger.Print("INFO",config.SrvName,config.Channel,"1","200","SendMessage ready to delivery message")

	emailResponse := db.MessageResponse{}
	m := gomail.NewMessage()
	m.SetHeader("From", config.EmailSMTPFrom)
	m.SetAddressHeader("To", message.Recipient, message.Recipient)
	m.SetHeader("Subject", message.Subject)
	m.SetBody("text/html", message.Content)
	if message.FileAttached != "" {
		m.Attach(message.FileAttached)
		logger.Print("INFO",config.SrvName,config.Channel,"1","200","Attached the file to mail")
	}
	logger.Print("INFO",config.SrvName,config.Channel,"1","200","Created the New Message before sending")

	s, err := smtpDialer.Dial()

	if err != nil {
		logger.Print("ERROR",config.SrvName,config.Channel,"1","500","Error sending email " + err.Error(), err)
		emailResponse.Response = err.Error()
		emailResponse.Status = config.FAILED
		emailResponse.TimeOfResponse = time.Now()
		return emailResponse
	}
	logger.Print("INFO",config.SrvName,config.Channel,"1","200","Connected to the smtpServer successfully")

	er := gomail.Send(s, m)
	if er != nil {
		emailResponse.Response = er.Error()
		emailResponse.Status = config.FAILED
		emailResponse.TimeOfResponse = time.Now()
		logger.Print("ERROR",config.SrvName,config.Channel,"1","500","Error sending email " + er.Error(), err)
	} else {
		emailResponse.Response = "SENT"
		emailResponse.Status = config.SUCCESS
		emailResponse.TimeOfResponse = time.Now()
		logger.Print("INFO",config.SrvName,config.Channel,"1","200","Email sent to  ['" + message.Recipient + "']")
	}
	return emailResponse
}

func validateEmail(email string) bool {
	return mailck.CheckSyntax(email)
}

package channels

import (
	"github.com/OrangeB4B/notificationapp/db"
	"github.com/OrangeB4B/notificationapp/logger"
	"regexp"
	"github.com/OrangeB4B/notificationapp/config"
	"errors"
	"github.com/OrangeB4B/notificationapp/templates"
	"github.com/aws/aws-sdk-go/aws"
	"time"
	"strconv"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

//EventForSMS : SMS implementation for SMS
type EventForSMS struct {
	TriggeredEvent db.Event
}

//ParseTemplate : Parsing Template for SMS
func (event EventForSMS) ParseTemplate() ([]db.Message, error) {
	var messages []db.Message
	channelSupported := CheckChannel(event.TriggeredEvent, db.SMS)
	if !channelSupported {
		logger.Print("ERROR",config.SrvName,config.Channel,"1","417","Dropping event ['" + event.TriggeredEvent.AccountID+ "']. SMS channel not supported.")
		return messages, errors.New("SMS channel not supported")
	}
	logger.Print("INFO",config.SrvName,config.Channel,"1","200","Channel supported...")
	numOfRecipient := len(event.TriggeredEvent.Recipient)
	if numOfRecipient <= 0 {
		logger.Print("ERROR",config.SrvName,config.Channel,"1","405","Dropping event ['" + event.TriggeredEvent.AccountID+ "']. No recipient found.")
		return messages, errors.New("no recipients found")
	}
	smsContent, err := templates.ParseTemplateForMessage(event.TriggeredEvent, db.SMS, event.TriggeredEvent.Lang)
	if err!=nil{
		logger.Print("ERROR",config.SrvName,config.Channel,"1","500","Error parsing the template . ",err)
		return messages, errors.New("Error parsing the template")
	}
	//parse each sms separately because it may vary by recipient
	for _, em := range event.TriggeredEvent.Recipient {
		if validatePhone(em) {
			logger.Print("INFO",config.SrvName,config.Channel,"1","200","SMS: "+ em+" format is OK...creating the message")
			dateCreated := time.Now()
			message := db.Message{}
			message.Recipient = em
			message.Subject = event.TriggeredEvent.Subject
			message.Channel = "SMS"
			message.DateCreated = dateCreated
			message.TemplateId = event.TriggeredEvent.TemplateType +"-SMS-"+ event.TriggeredEvent.Lang
			message.Content = smsContent
			message.CustomData = event.TriggeredEvent.CustomData
			message.MessageID = strconv.Itoa(dateCreated.Nanosecond()) + em + event.TriggeredEvent.AccountID
			messages = append(messages, message)
			logger.Print("INFO",config.SrvName,config.Channel,"1","200","Message created and added to the list of message to send to recipient")
		}
	}
	return messages, nil
}

//SendMessage : Message Sending  for SMS
func (event EventForSMS) SendMessage(message db.Message) db.MessageResponse {

	sess,err := session.NewSession()
	if err!= nil{
		logger.Print("ERROR",config.SrvName,config.Channel,"1","500","Error creating the SNS session",err)
		return db.MessageResponse{Status: config.FAILED, Response: err.Error(), TimeOfResponse: time.Now()}
	}
	logger.Print("INFO",config.SrvName,config.Channel,"1","200","SNS session created successfully")

	svc := sns.New(sess)
	logger.Print("INFO",config.SrvName,config.Channel,"1","200","SNS client created successfully")

	if message.Content == "" || message.Recipient == "" {
		logger.Print("ERROR",config.SrvName,config.Channel,"1","400","Sending  Failed. Message body or recipient empty")
		return db.MessageResponse{Status: config.FAILED, Response: "MESSAGE or RECIPIENT EMPTY. Message: "+message.Content+ " - Recipient: "+message.Recipient, TimeOfResponse: time.Now()}
	}
	logger.Print("INFO",config.SrvName,config.Channel,"1","200","SendMessage ready to delivery message")

	smsResponse := db.MessageResponse{}
	params := &sns.PublishInput{
		Message: aws.String(message.Content),
		PhoneNumber: aws.String(message.Recipient),
	}
	var sender = make(map[string]*string)
	a := config.SENDER
	sender["DefaultSenderID"]=&a
	svc.SetSMSAttributes(&sns.SetSMSAttributesInput{Attributes: sender})
	resp, err := svc.Publish(params)

	if err != nil {
		smsResponse.Response = err.Error()
		smsResponse.Status = config.FAILED
		smsResponse.TimeOfResponse = time.Now()
		logger.Print("ERROR",config.SrvName,config.Channel,"1","500","Error sending SMS to  " + message.Recipient + " with error: "+ err.Error(), err)
	} else {
		smsResponse.Response = "SENT: "+resp.String()
		smsResponse.Status = config.SUCCESS
		smsResponse.TimeOfResponse = time.Now()
		logger.Print("INFO",config.SrvName,config.Channel,"1","200","SMS sent to  ['" + message.Recipient + "'] with messageID: "+resp.String())
	}
	return smsResponse
}

func validatePhone(phone string) bool {
	re := regexp.MustCompile(`^[\+|00]\d+$`) //regex to verify +3412345667 format
	return re.MatchString(phone)
}

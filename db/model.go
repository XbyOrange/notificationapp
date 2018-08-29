package db

import (
	"time"
)

const (
	SMS = "Sms"
	EMAIL = "Email"
	FAKE = "Fake"
)

// Event struct
type Event struct {
	AccountID    string            `json:"AccountId"`
	Subject      string            `json:"Subject"`
	Channel      map[string]bool   `json:"Channel"`
	Lang		 string			   `json:"Language"`
	Recipient    []string          `json:"Recipient"`
	CustomData   map[string]string `json:"CustomData"`
	TemplateType string            `json:"TemplateType"`
	DateCreated  time.Time         `json:"dateCreated"`
}

// Message struct
type Message struct {
	MessageID       string            `json:"messageId,omitempty"`
	Channel         string            `json:"channel,omitempty"`
	TemplateId      string            `json:"templateId,omitempty"`
	CustomData      map[string]string `json:"customData,omitempty"`
	Subject         string            `json:"subject,omitempty"`
	Content         string            `json:"content,omitempty"`
	Recipient       string            `json:"recipient,omitempty"`
	FileAttached    string            `json:"fileAttached,omitempty"`
	MessageResponse MessageResponse   `json:"messageResponse,omitempty"`
	DateCreated     time.Time         `json:"dateCreated,omitempty"`
}

// MessageResponse struct
type MessageResponse struct {
	Response       string    `json:"response,omitempty"`
	Status         string    `json:"status,omitempty"`
	APIStatus      string    `json:"apiStatus,omitempty"`
	TimeOfResponse time.Time `json:"timeOfResponse,omitempty"`
}


// Template struct
type Template struct{
	TemplateId	string			  `bson:"_id" json:"TemplateId"`
	Channel     string			  `bson:"Channel" json:"Channel"`
	Language	string			  `bson:"Lang" json:"Lang"`
	Content		string			  `bson:"Content" json:"Content"`
}

type TemplateDB struct{
	Template  string 		`bson:"_id"` //This will be templateId + channel + lang
	Content	  string		`bson:"Content"`
}
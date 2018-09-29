package templates

import (
	"github.com/alknopfler/notificationapp/db"
	"testing"
	"github.com/stretchr/testify/assert"
)

var fakeTempEvent = db.Event{
	AccountID:    "eventid123456",
	TemplateType: "FAKE",
	CustomData: map[string]string{
		"Name":     "Alknopfler",
		"ItemName": "Mi message",
	},
	Recipient: []string{},
	Channel: map[string]bool{
		"CHANNEL": true,
	},
}

func TestParseTemplateForMessage404Template(t *testing.T) {
	failureEvent := db.Event{TemplateType:"DUMMY_SERVICE"}
	parsed, err := ParseTemplateForMessage( failureEvent, "Sms","ES")
	assert.Error(t,err)
	assert.Equal(t,parsed,"")
}

func TestParseTemplateForMessageOKTemplate(t *testing.T) {
	parsed, err := ParseTemplateForMessage(fakeTempEvent, "CHANNEL","LANG")
	assert.NoError(t,err)
	assert.Equal(t,parsed,"ContentFAKE Alknopfler")
}




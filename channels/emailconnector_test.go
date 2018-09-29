package channels

import (
	"github.com/alknopfler/notificationapp/db"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/alknopfler/notificationapp/config"
)

var fakeEmailRecipient = "alknopfler@xxxx.com"
var fakeEmailEvent = db.Event{
	AccountID:    "eventid123456",
	TemplateType: "FAKE",
	CustomData: map[string]string{
		"Name":     "alknopfler",
		"ItemName": "Mi item configurable",
	},
	Recipient: []string{fakeEmailRecipient},
	Channel: map[string]bool{
		"CHANNEL": true,
	},
	Lang: "LANG",
}

func TestParseTemplateInvalidChannelEmail(t *testing.T) {
	fakeEmailEvent.Channel = map[string]bool{
		"Sms": true,
	}
	_, err := EventForEmail{fakeEmailEvent}.ParseTemplate()
	assert.Error(t,err)
}

func TestParseTemplateForAllMessagesEmail(t *testing.T) {
	fakeEmailEvent.Recipient = []string{
		fakeEmailRecipient,
		"alknopfler@xxxxx.com",
		"largo.alknopfler@xxxxx.com",
	}
	fakeEmailEvent.Channel = map[string]bool{
		"Email": true,
	}
	msg, err := EventForEmail{fakeEmailEvent}.ParseTemplate()
	assert.NoError(t,err)
	assert.NotNil(t,msg)
}

func TestParseTemplateAllMessagesExceptInvalidRecipientsEmail(t *testing.T) {

	fakeEmailEvent.Recipient = []string{
		fakeEmailRecipient,
		"alknopfler@xxxxxxxx.com",
		"233271234567",
	}
	fakeEmailEvent.Channel = map[string]bool{
		"Email": true,
	}
	msg, err := EventForEmail{fakeEmailEvent}.ParseTemplate()
	assert.NoError(t,err)
	assert.Equal(t,len(msg),len(fakeEmailEvent.Recipient) - 1)

}

func TestParseTemplateInvalidRecipientEmail(t *testing.T) {
	fakeEmailEvent.Recipient = []string{}
	_, err := EventForEmail{fakeEmailEvent}.ParseTemplate()
	assert.Error(t,err)
}

func TestParseTemplateEmail(t *testing.T) {
	fakeEmailEvent.Recipient = []string{
		fakeEmailRecipient,
	}
	result, err := EventForEmail{fakeEmailEvent}.ParseTemplate()
	assert.NoError(t,err)
	if result == nil || result[0].Content == "" {
		t.Errorf("Test failed. Result unexpected")
	}
}

func TestSendEmail(t *testing.T) {
	fakeEmailEvent.Recipient = []string{
		fakeEmailRecipient,
	}
	msg,err := EventForEmail{fakeEmailEvent}.ParseTemplate()
	assert.NoError(t,err)
	msg[0].FileAttached="fileAttached"
	resp:=EventForEmail{fakeEmailEvent}.SendMessage(msg[0])
	assert.Equal(t,config.FAILED,resp.Status)
}

func TestSendEmailEmptyContent(t *testing.T) {
	fakeEmailEvent.Recipient = []string{
		fakeEmailRecipient,
	}
	msg,err := EventForEmail{fakeEmailEvent}.ParseTemplate()
	assert.NoError(t,err)
	resp:=EventForEmail{fakeEmailEvent}.SendMessage(msg[0])
	assert.Equal(t,config.FAILED,resp.Status)
}
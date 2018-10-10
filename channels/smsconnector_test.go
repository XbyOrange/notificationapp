package channels

import (
	"github.com/OrangeB4B/notificationapp/db"
	"testing"
	"github.com/stretchr/testify/assert"
	"time"
	"github.com/OrangeB4B/notificationapp/config"
)

var dbEvent1 = &db.Event{
	AccountID: "fakeAccountID",
	Subject: "fakeSubject",
	Channel: map[string]bool{"Sms": true},
	Lang: "LANG",
	Recipient: []string{
		"003412345678",
		"+12345432345",
		},
	CustomData: map[string]string{"Name": "Bobby"},
	TemplateType: "FAKE",
	DateCreated: time.Now(),
}

var dbEvent2 = &db.Event{
	AccountID: "fakeAccountID",
	Subject: "fakeSubject",
	Channel: map[string]bool{"WHAT": true},
	Lang: "LANG",
	Recipient: []string{"003412345678",
		"+12345432345",
	},
	CustomData: map[string]string{"Name": "Bobby"},
	TemplateType: "FAKE",
	DateCreated: time.Now(),
}
var dbEvent3 = &db.Event{
	AccountID: "fakeAccountID",
	Subject: "fakeSubject",
	Channel: map[string]bool{"Sms": true},
	Lang: "LANG",
	Recipient: []string{},
	CustomData: map[string]string{"Name": "Bobby"},
	TemplateType: "FAKE",
	DateCreated: time.Now(),
}


func TestEventForSMS_ParseTemplateOk(t *testing.T) {
	messages, err := EventForSMS{*dbEvent1}.ParseTemplate()
	assert.NoError(t, err)
	assert.NotNil(t, messages)
	assert.Equal(t, 2, len(messages))
}


func TestEventForSMS_ParseTemplateChannelNotSupported(t *testing.T) {
	messages, err := EventForSMS{*dbEvent2}.ParseTemplate()
	assert.Nil(t, messages)
	assert.Error(t, err)
}

func TestEventForSMS_ParseTemplateNumRecipientMinusZero(t *testing.T) {
	messages, err := EventForSMS{*dbEvent3}.ParseTemplate()
	assert.Nil(t, messages)
	assert.Error(t, err)
}

func TestEventForSMS_ParseTemplateParseTemplateForMessageError(t *testing.T) {
	dbEvent1.Lang = "JP"
	messages, err := EventForSMS{*dbEvent1}.ParseTemplate()
	assert.Nil(t, messages)
	assert.Error(t, err)
	dbEvent1.Lang = "LANG"
}

func TestValidatePhoneOK(t *testing.T) {
	phones := []string{
		"0087654321",
		"+3423242341325",
		"+194379387349",
		"+0184375984957",
		"0081798798769687987",
	}

	for _, phone := range phones {
		assert.True(t, validatePhone(phone))
	}
}

func TestValidatePhoneKO(t *testing.T) {
	phones := []string{
		"3423242341325",
		"",
	}

	for _, phone := range phones {
		assert.False(t, validatePhone(phone))
	}
}

func TestEventForSMS_SendMessageOk(t *testing.T) {
	messages, err := EventForSMS{*dbEvent1}.ParseTemplate()
	assert.NoError(t, err)
	assert.NotNil(t, messages)
	resp:=EventForSMS{*dbEvent1}.SendMessage(messages[0])
	assert.Equal(t,config.FAILED,resp.Status)
}

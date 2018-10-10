package service

import (
	"testing"
	"github.com/OrangeB4B/notificationapp/db"
	"github.com/stretchr/testify/assert"
)

var (
	fakeTempEventGoodChannel = db.Event{
	TemplateType: "FAKE",
	Subject: "FAKE-SUBJECT",
	Channel: map[string]bool{
		"Fake": true,
	},
	}

	fakeTempEventBadChannel = db.Event{
		TemplateType: "FAKE",
		Channel: map[string]bool{
			"BAD": true,
		},
	}

	fakeTempEventErrorParser = db.Event{
		TemplateType: "BADTEMPLATE",
		Channel: map[string]bool{
			"Fake": true,
		},
	}
	fakeTempEventErrorSending = db.Event{
		TemplateType: "BADTEMPLATE",
		Subject: "FAKE-SUBJECT",
		Channel: map[string]bool{
			"Fake": true,
		},
	}


)

func TestEventProcessorForChannelGoodChannel(t *testing.T) {
	_,err:=EventProcessorForChannel(fakeTempEventGoodChannel)
	assert.NoError(t,err)
}

func TestEventProcessorForChannelBadChannel(t *testing.T) {
	_,err:=EventProcessorForChannel(fakeTempEventBadChannel)
	assert.NoError(t,err) //no error because it's not necessary to send the notification
}

func TestEventProcessorForChannelErrorTemplateParser(t *testing.T) {
	_,err:=EventProcessorForChannel(fakeTempEventErrorParser)
	assert.NotNil(t,err)
}

func TestEventProcessorForChannelErrorSending(t *testing.T) {
	_,err:=EventProcessorForChannel(fakeTempEventErrorSending)
	assert.NotNil(t,err)
}



func TestProcessEvent(t *testing.T) {

}
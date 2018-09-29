package api

import (
	"testing"
	"github.com/alknopfler/notificationapp/db"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
)

func TestHandleNotificationsOK(t *testing.T){
	//The rest of URL and headers,method, paths, etc... are probed by gorilla mux. Not Need test for that
	bodyFake:=db.Event{
		AccountID:    "eventid123456",
		TemplateType: "FAKE",
		CustomData: map[string]string{
			"Name":     "Alknopfler",
			"ItemName": "Mi message",
		},
		Recipient: []string{},
		Channel: map[string]bool{
			"Fake": true,
		},
		Subject: "FAKE-SUBJECT",
	}
	url := "http://localhost:8080/notifications"
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(bodyFake)

	req, _ := http.NewRequest("POST", url,body)
	w := httptest.NewRecorder()

	HandleNotifications(w,req)
	assert.Equal(t,w.Code,http.StatusCreated)
}

func TestHandleNotificationsBodyEmpty(t *testing.T){
	//The rest of URL and headers,method, paths, etc... are probed by gorilla mux. Not Need test for that

	url := "http://localhost:8080/notifications"

	req, _ := http.NewRequest("POST", url,nil)
	w := httptest.NewRecorder()

	HandleNotifications(w,req)
	assert.Equal(t,w.Code,http.StatusBadRequest)
}

func TestHandleNotificationsBodyMalformed(t *testing.T){
	//The rest of URL and headers,method, paths, etc... are probed by gorilla mux. Not Need test for that
	bodyFake:=`{
"Account":evento1",     

}`
	url := "http://localhost:8080/notifications"
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(bodyFake)

	req, _ := http.NewRequest("POST", url,body)
	w := httptest.NewRecorder()

	HandleNotifications(w,req)
	assert.Equal(t,w.Code,http.StatusBadRequest)
}

func TestHandleNotificationsNotSent(t *testing.T){
	//The rest of URL and headers,method, paths, etc... are probed by gorilla mux. Not Need test for that
	bodyFake:=db.Event{
		AccountID:    "eventid123456",
		TemplateType: "FAKE",
		CustomData: map[string]string{
			"Name":     "Alknopfler",
			"ItemName": "Mi message",
		},
		Recipient: []string{},
		Channel: map[string]bool{
			"Fake": true,
		},
		Subject: "FAKE-SUBJECTXXX",
	}
	url := "http://localhost:8080/notifications"
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(bodyFake)

	req, _ := http.NewRequest("POST", url,body)
	w := httptest.NewRecorder()

	HandleNotifications(w,req)
	assert.Equal(t,w.Code,http.StatusBadRequest)
}

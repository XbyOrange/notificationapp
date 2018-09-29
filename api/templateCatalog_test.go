package api

import (
	"testing"
	"github.com/alknopfler/notificationapp/db"
	"net/http"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
)

func TestHandleCreateTemplateOK(t *testing.T){
	//The rest of URL and headers,method, paths, etc... are probed by gorilla mux. Not Need test for that
	bodyFake:=db.Template{
		TemplateId: "Fake-template",
		Channel: "CHANNEL",
		Language: "LANG",
		Content: "CONTENIDO FAKE",
	}
	url := "http://localhost:8080/templates"
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(bodyFake)

	req, _ := http.NewRequest("POST", url,body)
	w := httptest.NewRecorder()

	HandleCreateTemplate(w,req)
	assert.Equal(t,w.Code,http.StatusCreated)
}

func TestHandleCreateTemplateBodyNil(t *testing.T){
	//The rest of URL and headers,method, paths, etc... are probed by gorilla mux. Not Need test for that
	url := "http://localhost:8080/templates"

	req, _ := http.NewRequest("POST", url,nil)
	w := httptest.NewRecorder()

	HandleCreateTemplate(w,req)
	assert.Equal(t,w.Code,http.StatusBadRequest)
}

func TestHandleCreateTemplateJsonMalformed(t *testing.T) {
	bodyFake:=`{
		"TemplateId": "CREATETEMP",
		"Channel": "Email",
		"LangXXXXX "EN",
		"Content": "<html><head></head><body> Hola {{.CustomData.Name}}, Gracias por contratar {{.CustomData.ItemName}}.<br><br>Pulse <a href={{.CustomData.ActivationLink}}>aquí</a> para activar de su cuenta </body></html>"
	}`
	url := "http://localhost:8080/templates"
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(bodyFake)

	req, _ := http.NewRequest("POST", url,body)
	w := httptest.NewRecorder()

	HandleCreateTemplate(w,req)
	assert.Equal(t,w.Code,http.StatusBadRequest)
}

func TestHandleCreateTemplateErrorInserting(t *testing.T) {
	bodyFake:=db.Template{
		TemplateId: "CREATETEMP",
		Channel: "CHANNEL",
		Language: "LANG",
		Content: "CONTENIDO FAKE",
	}
	url := "http://localhost:8080/templates"
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(bodyFake)

	req, _ := http.NewRequest("POST", url,body)
	w := httptest.NewRecorder()

	HandleCreateTemplate(w,req)
	assert.Equal(t,w.Code,http.StatusExpectationFailed)
}

func TestHandleDeleteTemplateEmptyParam(t *testing.T) {

	url := "http://localhost:8080/templates/"

	req, _ := http.NewRequest("DELETE", url,nil)
	w := httptest.NewRecorder()

	HandleDeleteTemplate(w,req)
	assert.Equal(t,w.Code,http.StatusExpectationFailed)
}

func TestHandleDeleteTemplateOK(t *testing.T) {

	url := "http://localhost:8080/templates/FAKE-CHANNEL-LANG"

	req, _ := http.NewRequest("DELETE", url,nil)
	w := httptest.NewRecorder()

	HandlerController().ServeHTTP(w, req)

	assert.Equal(t,w.Code,http.StatusOK)
}

func TestHandleDeleteTemplateNotFound(t *testing.T) {

	url := "http://localhost:8080/templates/FAKE-CHANNEL-LANGXXXXXX"

	req, _ := http.NewRequest("DELETE", url,nil)
	w := httptest.NewRecorder()

	HandlerController().ServeHTTP(w, req)

	assert.Equal(t,417,w.Code)
}


func TestHandleGetTemplate(t *testing.T) {

	url := "http://localhost:8080/templates/FAKE-CHANNEL-LANG"

	req, _ := http.NewRequest("GET", url,nil)
	w := httptest.NewRecorder()

	HandlerController().ServeHTTP(w, req)

	assert.Equal(t,w.Code,http.StatusOK)
}

func TestHandleGetTemplateContentEmpty(t *testing.T) {

	url := "http://localhost:8080/templates/FAKE-CHANNEL-LANG-Empty"

	req, _ := http.NewRequest("GET", url,nil)
	w := httptest.NewRecorder()

	HandlerController().ServeHTTP(w, req)

	assert.Equal(t,w.Code,http.StatusNotFound)
}

func TestHandleGetTemplateEmpty(t *testing.T) {

	url := "http://localhost:8080/templates/"

	req, _ := http.NewRequest("GET", url,nil)
	w := httptest.NewRecorder()

	HandlerController().ServeHTTP(w, req)

	assert.Equal(t,http.StatusNotFound,w.Code)
}

func TestHandleGetTemplateNotFound(t *testing.T) {

	url := "http://localhost:8080/templates/FAKE-CHANNEL-LANGXXXX"

	req, _ := http.NewRequest("GET", url,nil)
	w := httptest.NewRecorder()

	HandlerController().ServeHTTP(w, req)

	assert.Equal(t,http.StatusNotFound,w.Code)
}

func TestHandleGetTemplateAll(t *testing.T) {

	url := "http://localhost:8080/templates"

	req, _ := http.NewRequest("GET", url,nil)
	w := httptest.NewRecorder()

	HandlerController().ServeHTTP(w, req)

	assert.Equal(t,w.Code,http.StatusOK)
}
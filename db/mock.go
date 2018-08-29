package db

import "errors"

//Fake struct
type Fake struct{}

var mockIds = map[string]bool {
	"FAKE-CHANNEL-LANG": true,
	"FAKE-Email-LANG": true,
	"FAKE-Sms-LANG": true,
}

func (f *Fake) Insert(t TemplateDB) error {
	if t.Template=="CREATETEMP-CHANNEL-LANG"{
		return errors.New("ERROR MOCK-FAKE")
	}
	return nil
}

func (f *Fake) FindAll() ([]TemplateDB, error){
	var msg []TemplateDB
	msg=append(msg,TemplateDB{
		Template: "TemplateFAKE",
		Content: "ContentFAKE",
	})
	return msg, nil
}

func (f *Fake) Find(id string) (TemplateDB,error){
	if mockIds[id] {
		return TemplateDB{"TemplateFAKE","ContentFAKE {{.CustomData.Name}}"},nil
	}
	if id == "FAKE-CHANNEL-LANG-Empty"{
		return TemplateDB{"TemplateFAKE",""},nil
	}
	return TemplateDB{}, errors.New("ERROR MOCK FAKE")
}

func (f *Fake) Remove (id string) bool{
	return id=="FAKE-CHANNEL-LANG"
}


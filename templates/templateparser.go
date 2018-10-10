package templates

import (
	"bytes"
	"github.com/OrangeB4B/notificationapp/db"
	"text/template"
	"github.com/OrangeB4B/notificationapp/config"
	"github.com/OrangeB4B/notificationapp/logger"
	"errors"
)

//ParseTemplateForMessage : Parses Template
func ParseTemplateForMessage(event db.Event, channel string, lang string) (string, error) {
	var parse string
	temp:=getTemplateFromDB(event.TemplateType,channel,lang)
	if temp == "" {
		logger.Print("ERROR",config.SrvName,config.Template,"1","404","Template not available. Error parsing while looking for the right template")
		return "", errors.New("Error parsing while looking for the template")
	}
	t := template.New("Template")
	t, err := t.Parse(temp)
	if err != nil {
		logger.Print("ERROR",config.SrvName,config.Template,"1","417","Error parsing template. Event dropped. Reason: " + err.Error(), err)
		return parse, err
	}
	var tpl bytes.Buffer
	err = t.Execute(&tpl, event)
	if err != nil {
		logger.Print("ERROR",config.SrvName,config.Template,"1","417","Error parsing template. Event dropped. Reason: " + err.Error(), err)
		return parse, err
	}
	logger.Print("INFO",config.SrvName,config.Template,"1","0","Template: "+event.TemplateType+"-"+channel+"-"+lang+" parsed successfully")
	return tpl.String(), err
}

func getTemplateFromDB(nombre,channel,lang string) string{
	logger.Print("INFO",config.SrvName,config.Template,"1","0","Getting template: "+nombre+"-"+channel+"-"+lang+" from database")
	result,err:=db.TypeDb.Find(nombre+"-"+channel+"-"+lang)
	if err!=nil{
		logger.Print("ERROR",config.SrvName,config.Template,"1","0","Template: "+nombre+"-"+channel+"-"+lang+" Not found", err)
		return ""
	}
	return result.Content
}

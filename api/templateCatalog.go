package api

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/alknopfler/notificationapp/logger"
	"github.com/alknopfler/notificationapp/config"
	"github.com/alknopfler/notificationapp/db"
	"github.com/gorilla/mux"
)

func HandleCreateTemplate(w http.ResponseWriter, r *http.Request) {
	if r.Body!=nil {
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			logger.Print("ERROR", config.SrvName, config.Api, "1", "400", "Error reading the request body json: "+err.Error(), err)
			responseWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		var input db.Template
		err = json.Unmarshal(b, &input)
		if err != nil {
			logger.Print("ERROR", config.SrvName, config.Api, "1", "400", "Error while unmarshalling input JSON"+err.Error(), err)
			responseWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		value := db.TemplateDB{
			Template: input.TemplateId + "-" + input.Channel + "-" + input.Language,
			Content:  input.Content,
		}

		err2 := db.TypeDb.Insert(value)
		if err2 != nil {
			logger.Print("ERROR", config.SrvName, config.Api, "1", "400", "Error inserting in mongo"+err2.Error(), err2)
			responseWithError(w, http.StatusExpectationFailed, err2.Error())
			return
		}
		logger.Print("INFO", config.SrvName, config.Api, "1", "200", "Template created successfully by handlerCreateTemplate")
		responseWithJSON(w, http.StatusCreated, "Template created successfully")
		return
	}
	logger.Print("ERROR", config.SrvName, config.Api, "1", "400", "Error reading the request body json")
	responseWithError(w, http.StatusBadRequest,"Request Body is Empty")
	return
}


func HandleGetTemplate(w http.ResponseWriter, r *http.Request) {
	if mux.Vars(r)["templateId"]!= ""{
		result,err:=db.TypeDb.Find(mux.Vars(r)["templateId"])
		if err!=nil{
			logger.Print("ERROR", config.SrvName,config.Api,"1","400", "Error searching in mongo"+err.Error(), err)
			responseWithError(w,http.StatusNotFound,"")
			return
		}
		if result.Content != "" {
			logger.Print("INFO", config.SrvName,config.Api,"1","200", "Getting the template list by api")
			responseWithJSON(w, http.StatusOK, result)
			return
		}else{
			logger.Print("ERROR", config.SrvName,config.Api,"1","404", "Template not Found")
			responseWithError(w,http.StatusNotFound,"Template not found")
			return
		}
	}
	logger.Print("ERROR", config.SrvName,config.Api,"1","400", "The uri param is empty")
	responseWithError(w,http.StatusExpectationFailed,"The URI param is empty")
	return
}

func HandleGetAllTemplates(w http.ResponseWriter, r *http.Request) {
		result,err:=db.TypeDb.FindAll()
		if err!=nil{
			logger.Print("ERROR", config.SrvName,config.Api,"1","400", "Error searching in mongo"+err.Error(), err)
			responseWithError(w,http.StatusExpectationFailed,err.Error())
			return
		}
	logger.Print("INFO", config.SrvName,config.Api,"1","200", "Getting all templates by api")
	responseWithJSON(w,http.StatusOK,result)
		return
}


func HandleDeleteTemplate(w http.ResponseWriter, r *http.Request) {
	if mux.Vars(r)["templateId"]!= ""{
		result:=db.TypeDb.Remove(mux.Vars(r)["templateId"])
		if ! result {
			logger.Print("ERROR", config.SrvName,config.Api,"1","404", "Error Deleting template in mongo")
			responseWithError(w,http.StatusExpectationFailed,"Error deleting the Template ")
			return
		}
		logger.Print("INFO", config.SrvName,config.Api,"1","200", "Deleting successfully template by the api")
		responseWithJSON(w,http.StatusOK,"Successfully deleted")
		return
	}
	logger.Print("ERROR", config.SrvName,config.Api,"1","400", "The uri param is empty")
	responseWithError(w,http.StatusExpectationFailed,"The URI param is empty")
	return
}
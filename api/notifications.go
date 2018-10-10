package api

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/OrangeB4B/notificationapp/logger"
	"github.com/OrangeB4B/notificationapp/config"
	"github.com/OrangeB4B/notificationapp/db"
	"github.com/OrangeB4B/notificationapp/service"
)

func HandleNotifications(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil{
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			logger.Print("ERROR", config.SrvName, config.Api,"1", "400","Error reading the request body json: "+err.Error(), err)
			responseWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		var value db.Event
		err = json.Unmarshal(b, &value)
		if err != nil {
			logger.Print("ERROR", config.SrvName,config.Api,"1","400", "Error while unmarshalling input JSON"+err.Error(), err)
			responseWithError(w,http.StatusBadRequest,err.Error())
			return
		}
		msg,err2 := service.EventProcessorForChannel(value)
		if err2!=nil{
			logger.Print("ERROR", config.SrvName, config.Api,"1", "400","Error reading the request body json: "+err2.Error(), err2)
			responseWithError(w, http.StatusBadRequest, err2.Error())
			return
		}
		logger.Print("INFO",config.SrvName,config.Api,"1","200","Successful Operation handling the api notification")
		responseWithJSON(w,http.StatusCreated,msg)
		return
	}
	logger.Print("ERROR", config.SrvName, config.Api,"1", "400","Error reading the request body ")
	responseWithError(w, http.StatusBadRequest, "Error Body Empty")
	return
}
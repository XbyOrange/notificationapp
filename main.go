package main

import (
	"gitlab.com/backend/notifications/notificationApp/config"
	"gitlab.com/backend/notifications/notificationApp/logger"
	"os"
	"net/http"
	"gitlab.com/backend/notifications/notificationApp/api"
)


func init() {
	//Load the logger configuration at the beginning
	go logger.Init(os.Stdout, config.LogLevel)
	logger.Print("INFO",config.SrvName,"main","1","200","Initialization finished...")
}

func main() {

	err := http.ListenAndServe(config.SrvPort, api.HandlerController())
	if err != nil {
		logger.Print("ERROR",config.SrvName,"main","1","500","Failed to listen the server...",err)
	}

}

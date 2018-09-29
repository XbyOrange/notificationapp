package api

import (
	"net/http"
	"net"
	"github.com/alknopfler/notificationapp/config"
)

func handleLiveness(w http.ResponseWriter, r *http.Request) {
	if (!CheckIsUP(config.SrvPort)) {
		w.WriteHeader(http.StatusExpectationFailed)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	return
}

func handleReadiness(w http.ResponseWriter, r *http.Request) {
	if !CheckIsUP(config.MongoHost+":27017") || !CheckIsUP(config.SrvPort) {
		w.WriteHeader(http.StatusExpectationFailed)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	return
}

func CheckIsUP(connection string) bool{
	ln, err := net.Dial("tcp",connection)
	if err != nil {
		return false
	}
	err = ln.Close()
	return true
}

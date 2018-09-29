package api

import (
	"github.com/gorilla/mux"
)

//HandlerController function
func HandlerController() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/notifications", HandleNotifications).Methods("POST")
	r.HandleFunc("/templates", HandleCreateTemplate).Methods("POST")
	r.HandleFunc("/templates",HandleGetAllTemplates).Methods("GET")
	r.HandleFunc("/templates/{templateId}", HandleGetTemplate).Methods("GET")
	r.HandleFunc("/templates/{templateId}", HandleDeleteTemplate).Methods("DELETE")
	r.HandleFunc("/liveness",handleLiveness).Methods("GET")
	r.HandleFunc("/readiness",handleReadiness).Methods("GET")

	//r.HandleFunc("/templates/{templateId}", handleUpdateTemplate).Methods("PUT")

	return r
}


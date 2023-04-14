package router

import (
	"chat/controller"
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {
    router := mux.NewRouter()
	router.HandleFunc("/api/send", controller.SendMessage).Methods("POST");
	router.HandleFunc("/api/unread", controller.GetUnreadMessages).Methods("GET");
	return router
}
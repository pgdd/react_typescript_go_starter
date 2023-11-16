package routes

import (
	"gorestapi/api/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/accounts", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/accounts/{id}", controllers.GetAccount).Methods("GET")
	router.HandleFunc("/transfers", controllers.InitiateTransfer).Methods("POST")
}

package main

import (
	"fmt"
	"jwt-practice/database"
	"jwt-practice/service"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router

func CreateRouter() {
	router = mux.NewRouter()
}

func InitializeRouter() {
	router.HandleFunc("/signup", service.SignUp).Methods("POST")
	router.HandleFunc("/login", service.Login).Methods("POST")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	database.InitialMigration()
	CreateRouter()
	InitializeRouter()

}

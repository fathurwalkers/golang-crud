package main

import (
	"golang-crud/config"
	"golang-crud/controllers/homecontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	// 1. Home Page 
	http.HandleFunc("/", homecontroller.Welcome)

	log.Println("Server UP and Running on Port : 5002")
	http.ListenAndServe(":5002", nil)
}
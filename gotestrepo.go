package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func clientLoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am a Test HTTP Server")
	fmt.Println("This is a Modules Test")
	fmt.Println("Checking these changes from MAC")
}

func main() {
	// Init a new router
	URLRouter := mux.NewRouter()
	URLRouter.HandleFunc("/login", clientLoginHandler).Methods("POST")

	fmt.Println("Starting Server on Port 8030...")
	log.Fatal(http.ListenAndServe(":8030", URLRouter))

}

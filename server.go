package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	file, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	log.Println("Welcome to the Orders App!")

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Fatal(http.ListenAndServe(":3000", mux))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Orders App!\n"))
}

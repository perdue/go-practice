package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/perdue/go-practice/pkg/handlers"
)

func logPrefix() string {
	return fmt.Sprintf("%-20s", "[main]")
}

func logMsg(msg string) {
	log.SetPrefix(logPrefix())
	log.Println(msg)
}

func main() {
	file, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	logMsg("Welcome to the Orders App!")

	h, _ := handlers.New()
	logMsg(fmt.Sprintf("%#v", h))

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Fatal(http.ListenAndServe(":3000", mux))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Orders App!\n"))
}

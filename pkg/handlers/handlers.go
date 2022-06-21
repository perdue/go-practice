package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func logPrefix() string {
	return fmt.Sprintf("%-20s", "[handlers]")
}

type handler struct {
}

type Handler interface {
	Index(w http.ResponseWriter, r *http.Request)
}

func New() (Handler, error) {
	logMsg("Creating new handler")
	h := handler{}
	return &h, nil
}

// Index returns a simple welcome response for the home page
func (h *handler) Index(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, http.StatusOK, "Welcome to the Orders App!", nil)
}

func logMsg(msg string) {
	log.SetPrefix(logPrefix())
	log.Println(msg)
}

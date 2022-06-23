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
	ProductIndex(w http.ResponseWriter, r *http.Request)
	OrderShow(w http.ResponseWriter, r *http.Request)
	OrderInsert(w http.ResponseWriter, r *http.Request)
	Close(w http.ResponseWriter, r *http.Request)
	Stats(w http.ResponseWriter, r *http.Request)
	OrderReverse(w http.ResponseWriter, r *http.Request)
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

// ProductIndex displays all the products in the system
func (h *handler) ProductIndex(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, http.StatusOK, "Welcome to the Product Index", nil)
}

// OrderShow
func (h *handler) OrderShow(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, http.StatusOK, "Welcome to the Order Show", nil)
}

// OrderInsert
func (h *handler) OrderInsert(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, http.StatusOK, "Welcome to the Order Insert", nil)
}

// Close
func (h *handler) Close(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, http.StatusOK, "Welcome to the Close", nil)
}

// Stats
func (h *handler) Stats(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, http.StatusOK, "Welcome to the Stats", nil)
}

// OrderReverse
func (h *handler) OrderReverse(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, http.StatusOK, "Welcome to the Order Reverse", nil)
}

func logMsg(msg string) {
	log.SetPrefix(logPrefix())
	log.Println(msg)
}

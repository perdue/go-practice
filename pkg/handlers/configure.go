package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func ConfigureHandlers(handler Handler) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// TODO: Implement Index
	router.Methods("GET").Path("/").Handler(http.HandlerFunc(notImplemented))

	// TODO: Implement ProductIndex
	router.Methods("GET").Path("/products").Handler(http.HandlerFunc(notImplemented))

	// TODO: Implement OrderShow
	router.Methods("GET").Path("/orders/{orderId}").Handler(http.HandlerFunc(notImplemented))

	// TODO: Implement OrderInsert
	router.Methods("POST").Path("/orders").Handler(http.HandlerFunc(notImplemented))

	// TODO: Implement Close
	router.Methods("POST").Path("/close").Handler(http.HandlerFunc(notImplemented))

	// TODO: Implement Stats
	router.Methods("GET").Path("/stats").Handler(http.HandlerFunc(notImplemented))

	// TODO: Implement OrderReverse
	router.Methods("DELETE").Path("/orders/{orderId}").Handler(http.HandlerFunc(notImplemented))

	return router
}

func notImplemented(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not implemented yet\n"))
}

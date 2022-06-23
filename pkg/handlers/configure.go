package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func ConfigureHandlers(handler Handler) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("/").Handler(http.HandlerFunc(handler.Index))

	// TODO: Implement ProductIndex
	router.Methods("GET").Path("/products").Handler(http.HandlerFunc(handler.ProductIndex))

	// TODO: Implement OrderShow
	router.Methods("GET").Path("/orders/{orderId}").Handler(http.HandlerFunc(handler.OrderShow))

	// TODO: Implement OrderInsert
	router.Methods("POST").Path("/orders").Handler(http.HandlerFunc(handler.OrderInsert))

	// TODO: Implement Close
	router.Methods("POST").Path("/close").Handler(http.HandlerFunc(handler.Close))

	// TODO: Implement Stats
	router.Methods("GET").Path("/stats").Handler(http.HandlerFunc(handler.Stats))

	// TODO: Implement OrderReverse
	router.Methods("DELETE").Path("/orders/{orderId}").Handler(http.HandlerFunc(handler.OrderReverse))

	return router
}

func notImplemented(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not implemented yet\n"))
}

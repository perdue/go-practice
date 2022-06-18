package handlers

import (
	"fmt"
	"log"
)

func logPrefix() string {
	return fmt.Sprintf("%-20s", "[handlers]")
}

type handler struct {
}

type Handler interface {
}

func New() (Handler, error) {
	logMsg("Creating new handler")
	h := handler{}
	return &h, nil
}

func logMsg(msg string) {
	log.SetPrefix(logPrefix())
	log.Println(msg)
}

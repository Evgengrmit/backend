package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func checkError(w io.Writer, err error, funcName string) {
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		log.Println(fmt.Errorf("ERROR\tfunction %s: %s", funcName, err.Error()))
	}
}

// Транслирование данных  обратно клиенту и в лог
func translateError(w io.Writer, err error, funcName string) {
	_, _ = w.Write([]byte(err.Error()))
	log.Println(fmt.Errorf("ERROR\tfunction %s: %s", funcName, err.Error()))
}
func translateInfo(w http.ResponseWriter, info string, funcName string) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(info))
	log.Println(fmt.Sprintf("INFO\tfunction %s: %s", funcName, info))
}

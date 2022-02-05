package helper

import (
	"fmt"
	"io"
	"log"
)

func CheckError(w io.Writer, err error, funcName string) {
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		log.Println(fmt.Errorf("ERROR\tfunction %s: %s", funcName, err.Error()))
	}
}

// TranslateError Транслирование данных  обратно клиенту и в лог
func TranslateError(w io.Writer, err error, funcName string) {
	_, _ = w.Write([]byte(err.Error()))
	log.Println(fmt.Errorf("ERROR\tfunction %s: %s", funcName, err.Error()))
}

//func TranslateInfo(w http.ResponseWriter, info string, funcName string) {
//	w.Header().Set("Content-Type", "application/json")
//	_, _ = w.Write([]byte(info))
//	log.Println(fmt.Sprintf("INFO\tfunction %s: %s", funcName, info))
//}

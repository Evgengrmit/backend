package main

import (
	"backend/db"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	dbUri := os.Getenv("DB_URI")
	if dbUri == "" {
		log.Fatal("Provide db uri")
	}
	db.ConnectDatabase(dbUri)

	// Запуск миграции данных, т.е. создаются таблицы/наполняются инфомарцией и так далее
	//

	log.Fatal(RunServer(port))
}

package main

import (
	"backend/db"
	"backend/user"
	"fmt"
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

	row, err := db.DB.Query("select * from \"user\"")
	if err != nil {
		fmt.Println(err)
	}

	var newUser []user.User
	for row.Next() {
		u := user.User{}
		err = row.Scan(&u.ID, &u.Name, &u.Age, &u.Login, &u.Email, &u.HashedPassword)
		newUser = append(newUser, u)
	}

	for _, u := range newUser {
		fmt.Println(u)
	}

	// Запуск миграции данных, т.е. создаются таблицы/наполняются инфомарцией и так далее
	//

	log.Fatal(RunServer(port))
}

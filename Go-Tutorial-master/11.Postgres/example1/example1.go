package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"fmt"
)

func main () {
	db, err := sql.Open("postgres", "password=password host=localhost dbname=mit_workshop sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	var insert_users string = "Insert into Users(id, email) values($1, $2)"
	insert_users_prepare, err := db.Prepare(insert_users)
	if err!= nil {
		log.Fatal(err)
	}
	defer insert_users_prepare.Close()

	insert_res, err := insert_users_prepare.Exec(1, "praveen@yopmail.com")
	if err != nil || insert_res == nil{
		log.Fatal(err)
	}

	email_query, err := db.Query("SELECT email FROM users where id = 1")
	if err != nil {
		log.Fatal(err)
	}
	defer email_query.Close()

	for email_query.Next() {
		var email string

		err := email_query.Scan(&email)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Email:", email)
	}
	db.Close()
}

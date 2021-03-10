package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var db, errOpen = sql.Open("sqlite3", "./data.db")
	defer db.Close()
	if errOpen != nil {
		log.Fatal(errOpen)
	}

	//Prepare: Prepare a sentence and I can use it a lot of times
	//Close statement when it not use any more
	var statement, err = db.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	defer statement.Close()
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec() //Executes a query (no returns rows)

	statement, err = db.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
	defer statement.Close()
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec("Marco", "Diaz")

	statement, err = db.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
	defer statement.Close()
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec("Ana", "Gomez")

	//Executes a query (returns rows)
	var rows, errSelect = db.Query("SELECT id, firstname, lastname FROM people")
	defer rows.Close()
	if errSelect != nil {
		log.Fatal(errSelect)
	}

	var id int
	var firstname string
	var lastname string

	//Bucle For for nexts rows
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
	}
}

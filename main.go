package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var db, _ = sql.Open("sqlite3", "./data.db")
	//Prepare: Prepare a sentence and I can use it a lot of times
	//Close statement when it not use any more
	var statement, _ = db.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	//Executes a query (no returns rows)
	statement.Exec()
	statement, _ = db.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
	statement.Exec("Marco", "Diaz")
	statement, _ = db.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
	statement.Exec("Ana", "Gomez")

	//Close statement
	statement.Close()

	//Executes a query (returns rows)
	var rows, _ = db.Query("SELECT id, firstname, lastname FROM people")
	var id int
	var firstname string
	var lastname string

	//Bucle For for nexts rows
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
	}
}

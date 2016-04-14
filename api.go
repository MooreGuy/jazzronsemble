package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Concert struct {
	Name        string
	Description string
	Id          int
	Date        time.Time
}

func handleConcerts(w http.ResponseWriter, r *http.Request) {
	params, ok := r.URL.Query()["id"]
	if !ok {
		log.Println("Id not provided, but was required.")
		return
	}

	if len(params) < 1 {
		log.Println("Id not provided, but was required.")
		return
	}
	id, err := strconv.Atoi(params[0])
	if err != nil {
		log.Println(err.Error())
	}

	db, err := sql.Open("mysql", "root@/ronsemble")
	if err != nil {
		log.Println(err.Error())
		return
	}

	stmt, err := db.Prepare(`
		SELECT name, description, id, date
		FROM concerts
		WHERE id = ?`)
	if err != nil {
		log.Println(err.Error())
		return
	}

	rows, err := stmt.Query(strconv.Itoa(id))
	if err != nil {
		log.Println(err.Error())
	}

	rows.Next()
	var nextConcertName Concert
	err = rows.Scan(&nextConcertName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(nextConcertName)

	json, err := json.Marshal(nextConcertName)
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Fprintf(w, string(json))
}

func handleNamePuns(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "CD-Ron McCarly")
}

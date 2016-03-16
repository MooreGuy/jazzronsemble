package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

const nameInput string = "newname"

type Ronsemble struct {
	RonName   string
	Title     string
	NameInput string
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/newName", handleNewName)
	http.Handle("/images/", http.StripPrefix("/images/",
		http.FileServer(http.Dir("/srv/img"))))
	http.Handle("/css/", http.StripPrefix("/css/",
		http.FileServer(http.Dir("/srv/css"))))

	log.Fatal(http.ListenAndServe(":80", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile("index.html")
	if err != nil {
		log.Fatal(err.Error())
	}

	templ, err := template.New("jazzronsemble").Parse(string(file))
	if err != nil {
		log.Println("Failed to parse template.")
	}

	name := getName()
	data := Ronsemble{name, "Title, Yo", nameInput}

	err = templ.ExecuteTemplate(w, "jazzronsemble", data)
}

func handleNewName(w http.ResponseWriter, r *http.Request) {
	newName := r.FormValue(nameInput)

	db := connectDB()
	defer db.Close()
	createName(db, newName)

	handleIndex(w, r)
}

func getName() string {
	db := connectDB()
	defer db.Close()
	names := getNames(db)

	index := rand.Int() % len(names)
	return names[index]
}

func connectDB() *sql.DB {
	db, err := sql.Open("mysql", "ronsemble@/ronsemble")
	if err != nil {
		panic(err.Error())
	}

	return db
}

func createName(db *sql.DB, name string) {
	query, err := db.Prepare("INSERT INTO names VALUES (?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()

	// 0 value for automatically incremented nameid column
	_, err = query.Exec(name, 0)
	if err != nil {
		panic(err.Error())
	}
}

func getNames(db *sql.DB) []string {
	query, err := db.Prepare("SELECT name, nameid FROM names")
	if err != nil {
		panic(err.Error())
	}

	rows, err := query.Query()
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	names := []string{}
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&name, &id); err != nil {
			panic(err.Error())
		}

		names = append(names, name)
	}

	if err := rows.Err(); err != nil {
		panic(err.Error())
	}

	return names
}

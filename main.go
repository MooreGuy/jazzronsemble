package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

const nameInput string = "newname"

type Ronsemble struct {
	RonName   string
	Title     string
	NameInput string
}

func main() {
	connectDB()
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

	f, err := os.OpenFile("/srv/names.txt", os.O_APPEND|os.O_WRONLY, 077)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = f.WriteString(newName + "\n")
	if err != nil {
		log.Fatal(err.Error())
	}

	handleIndex(w, r)
}

func getName() string {
	f, err := os.OpenFile("/srv/names.txt", os.O_RDONLY, 0777)
	if err != nil {
		log.Println(err.Error())
	}

	rawNames := make([]byte, 1024)
	_, err = f.Read(rawNames)
	if err != nil {
		log.Println(err.Error())
	}

	names := strings.Split(string(rawNames), "\n")
	index := rand.Int() % len(names)
	return names[index]
}

func connectDB() *sql.DB {
	db, err := sql.Open("mysql", "ronsemble@/ronsemble")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	return db
}

func createName(db *sql.DB) {
	query, err := db.Prepare("INSERT INTO names VALUES (?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()

	_, err = query.Exec("Ronald Duck", 0)
	if err != nil {
		panic(err.Error())
	}
}

func getNames(db *sql.DB) *sql.Rows {
	query, err := db.Prepare("SELECT name, nameid FROM names")
	if err != nil {
		panic(err.Error())
	}

	values, err := query.Query()
	if err != nil {
		panic(err.Error())
	}

	return values
}

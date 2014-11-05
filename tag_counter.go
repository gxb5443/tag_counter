package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/coopernurse/gorp"
	_ "github.com/lib/pq"
)

type Tag struct {
	Id    int64
	Tag   string
	Count int64
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Web Server Online! Param is %s!", r.URL.Path[1:])
}

func taghandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "You got it alright!")
	} else {
		fmt.Fprintf(w, "Sorry, %s functionality not yet supported", r.Method)
	}
}

func tag2handler(w http.ResponseWriter, r *http.Request) {
	req_tag := r.URL.Path[len("/tag/"):]
	if r.Method == "GET" {
		fmt.Fprintf(w, "Looking for tag: \"%s\"", req_tag)
	} else {
		fmt.Fprintf(w, "Sorry, %s functionality not yet supported", r.Method)
	}
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func initDb(config string) *gorp.DbMap {
	//db, err := sql.Open("postgres", "user=admin password=admin dbname=tag_counting sslmode=disable")
	db, err := sql.Open("postgres", config)
	checkErr(err, "postgres.Open failed")

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	table := dbmap.AddTableWithName(Tag{}, "tags").SetKeys(true, "Id")
	table.ColMap("Tag").SetNotNull(true)
	table.ColMap("Count").SetNotNull(true)

	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Could not create tables")

	return dbmap
}

func main() {
	config, err := ioutil.ReadFile("db.conf")
	PanicIf(err)
	dbmap := initDb(string(config))
	defer dbmap.Db.Close()
	fmt.Printf("Starting Web Server...")
	http.HandleFunc("/tag/", tag2handler)
	http.HandleFunc("/tag", taghandler)
	http.HandleFunc("/", handler)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Could not start server: %s", err)
	}
}

package main

import (
	"fmt"
	"net/http"
)

type Tag struct {
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

func main() {
	fmt.Printf("Starting Web Server...")
	http.HandleFunc("/tag/", tag2handler)
	http.HandleFunc("/tag", taghandler)
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Could not start server: %s", err)
	}
}

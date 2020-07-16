package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	log.Println(r.Form)
	fmt.Fprintf(w, "Hi there, I love %s %s!", r.URL.Path[1:])
}


func main() {
	http.HandleFunc("/message", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
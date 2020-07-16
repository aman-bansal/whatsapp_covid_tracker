package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var value []byte
	err := json.NewDecoder(r.Body).Decode(&value)
	if err != nil {
		log.Println("error", err)
	}
	log.Println(value)
	fmt.Fprintf(w, "Hi there, I love %s %s!", r.URL.Path[1:], value)
}


func main() {
	http.HandleFunc("/message", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
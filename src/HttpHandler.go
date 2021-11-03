package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func HttpHandler(w http.ResponseWriter, r *http.Request) {
	errRarse := r.ParseForm()
	if errRarse != nil {
		log.Fatal("ParseForm: ", errRarse)
	}
	fmt.Println("path", r.URL.Path)
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	_, errWrite := fmt.Fprintf(w, "Hi this is a Microservices Service Registry Center.")
	if errWrite != nil {
		log.Fatal("Fprintf: ", errWrite)
	}
}

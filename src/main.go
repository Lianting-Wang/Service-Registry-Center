package main

import (
	"log"
	"net/http"
)

var HeartBeatLists HeartBeatList

func main() {
	HeartBeatLists.initialize()
	http.HandleFunc("/", HttpHandler)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

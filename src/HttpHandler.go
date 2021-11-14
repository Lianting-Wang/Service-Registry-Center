package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type HttpResponse struct {
	Service string `json:"service"`
	Code    int    `json:"code"`
	Address string `json:"address"`
}

func HttpHandler(w http.ResponseWriter, r *http.Request) {
	errRarse := r.ParseForm()
	if errRarse != nil {
		log.Fatal("ParseForm: ", errRarse)
	}
	var httpServerResponse HttpResponse
	if r.Method == "POST" {
		for k, v := range r.Form {
			var NewServer Servers
			go NewServer.initialize(k, strings.Join(v, ""))
			httpServerResponse = HttpResponse{
				k,
				0,
				strings.Join(v, "")}
			returnResponse, _ := json.Marshal(&httpServerResponse)
			_, errWrite := fmt.Fprintf(w, string(returnResponse))
			if errWrite != nil {
				log.Fatal("Fprintf: ", errWrite)
			}
			HeartBeatLists.AddServer(&NewServer)
		}
	} else if r.Method == "GET" {
		code := -1
		ServerName := strings.TrimLeft(r.URL.Path, "/")
		address := HeartBeatLists.FindServer(ServerName)
		if address != "" {
			code = 0
		}
		httpServerResponse = HttpResponse{
			ServerName,
			code,
			address}
		returnResponse, _ := json.Marshal(&httpServerResponse)
		_, errWrite := fmt.Fprintf(w, string(returnResponse))
		if errWrite != nil {
			log.Fatal("Fprintf: ", errWrite)
		}
	}
}

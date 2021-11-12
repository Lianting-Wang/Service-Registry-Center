package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type HeartBeatList struct {
	timeOfLostContact int
	waitDuration      int
	name              string
	address           string
	signal            int
}

func HeartBeatSender(h *HeartBeatList) {
	fmt.Println("Start to get...")
	resp, err := http.Get(h.address + "/HeartBeat")
	if err != nil {
		fmt.Println("Have a get err with HeartBeatSender")
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Have an err with CloseBody")
		}
	}(resp.Body)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	if string(body) == h.address {
		h.signal = 1
		return
	}
	h.signal = -1
	return
}

func (h *HeartBeatList) initialize(name string, address string) {
	h.name = name
	h.address = address
	h.timeOfLostContact = 0
	h.waitDuration = 30
	h.signal = 1
	for true {
		time.Sleep(time.Duration(h.waitDuration) * time.Second)
		fmt.Println("Start...")
		h.signal = 0
		go HeartBeatSender(h)
		time.Sleep(time.Duration(1) * time.Second)
		if h.signal == 1 {
			continue
		} else if h.signal == 0 {
			time.Sleep(time.Duration(10) * time.Second)
			if h.signal == 1 {
				continue
			} else if h.signal == 0 {
				fmt.Println("Connection timed out...")
			} else {
				fmt.Println("Connection error")
			}
		} else {
			fmt.Println("Connection error")
		}
		break
	}
	h.address = ""
}

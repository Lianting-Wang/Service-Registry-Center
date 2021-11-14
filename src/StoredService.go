package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Servers struct {
	timeOfLostContact int
	waitDuration      int
	name              string
	address           string
	signal            int
}

func HeartBeatSender(s *Servers) {
	resp, err := http.Get(s.address + "/HeartBeat")
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
	if string(body) == s.address {
		s.signal = 1
		return
	}
	s.signal = -1
	return
}

func (s *Servers) IsRunning() bool {
	return s.address != ""
}

func (s *Servers) GetName() string {
	return s.name
}

func (s *Servers) GetAddress() string {
	return s.address
}

func (s *Servers) initialize(name string, address string) {
	s.name = name
	s.address = address
	s.timeOfLostContact = 0
	s.waitDuration = 30
	s.signal = 1
	for true {
		time.Sleep(time.Duration(s.waitDuration) * time.Second)
		s.signal = 0
		go HeartBeatSender(s)
		time.Sleep(time.Duration(1) * time.Second)
		if s.signal == 1 {
			continue
		} else if s.signal == 0 {
			time.Sleep(time.Duration(10) * time.Second)
			if s.signal == 1 {
				continue
			} else if s.signal == 0 {
				fmt.Println("Connection timed out...")
			} else {
				fmt.Println("Connection error")
			}
		} else {
			fmt.Println("Connection error")
		}
		break
	}
	s.address = ""
	HeartBeatLists.RefreshServer()
}

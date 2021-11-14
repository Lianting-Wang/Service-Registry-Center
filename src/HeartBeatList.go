package main

type HeartBeatList struct {
	nums       int
	ServersMap map[int]*Servers
}

func (h *HeartBeatList) initialize() {
	h.nums = 1
	h.ServersMap = make(map[int]*Servers)
}

func (h *HeartBeatList) AddServer(NewServer *Servers) {
	h.ServersMap[h.nums] = NewServer
	h.nums++
}

func (h *HeartBeatList) FindServer(name string) string {
	for index := range h.ServersMap {
		servers := h.ServersMap[index]
		if servers.GetName() == name {
			return servers.GetAddress()
		}
	}
	return ""
}

//func (h *HeartBeatList) AllServer() {
//	for index := range h.ServersMap {
//		servers := h.ServersMap[index]
//		fmt.Println(servers.GetName())
//		fmt.Println(servers.GetAddress())
//	}
//}

func (h *HeartBeatList) RefreshServer() {
	for index := range h.ServersMap {
		servers := h.ServersMap[index]
		if !servers.IsRunning() {
			delete(h.ServersMap, index)
		}
	}
}

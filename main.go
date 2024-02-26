package main

import (
	"fmt"
	"time"
)

var serverList []Server

func main() {

	loadServersFile("servers.json")
	serverList = getServerList()

	for _, server := range serverList {
		go update(server)
	}

	server()
}

const updateInterval = 5

func update(server Server) {

	for time.Now().Minute()%updateInterval != 0 {
		return
	}

	trim(server)
	timestamp := time.Now().Format("2006-01-02 15:04")

	server.Counts[timestamp] = getServerData(server.IP).Players.Online
	fmt.Println("Updated", server.Name, "at", timestamp, "with", server.Counts[timestamp], "players")

	time.Sleep(1 * time.Minute)
	update(server)
}

func trim(server Server) {
	if len(server.Counts) > ((60 * 24 * 7) / updateInterval) {

		var firstKey string
		for key := range server.Counts {
			firstKey = key
			break
		}

		delete(server.Counts, firstKey)
		fmt.Println("Trimmed", server.Name, "at", firstKey)

	}
}

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var data []byte

type Server struct {
	Name   string `json:"name"`
	IP     string `json:"ip"`
	Counts map[string]int
}

type JSONData struct {
	Servers []Server `json:"servers"`
}

func loadServersFile(file string) {
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading config file")
		os.Exit(1)
	}
	data = content
}

func getServerList() []Server {
	var jsonData JSONData

	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		fmt.Println("Error parsing JSON")
		os.Exit(1)
	}

	servers := jsonData.Servers

	for i := range servers {
		servers[i].Counts = make(map[string]int)
	}

	return servers
}

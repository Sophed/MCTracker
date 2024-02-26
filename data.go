package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const URL = "https://mcapi.us/server/status?ip="

type ServerData struct {
	Players struct {
		Online int `json:"now"`
		Max    int `json:"max"`
	} `json:"players"`

	Server struct {
		Software string `json:"name"`
		Protocol int    `json:"protocol"`
	} `json:"server"`

	Icon string `json:"favicon"`
}

func getServerData(serverIP string) ServerData {
	response, err := http.Get(URL + serverIP)
	if err != nil {
		fmt.Println("Failed to get server data")
		os.Exit(1)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Reading response failed")
		os.Exit(1)
	}
	var serverData ServerData
	json.Unmarshal(body, &serverData)
	return serverData
}

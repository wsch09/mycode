package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type SpaceX []struct {
	CapsuleSerial      string    `json:"capsule_serial"`
	CapsuleID          string    `json:"capsule_id"`
	Status             string    `json:"status"`
	OriginalLaunch     time.Time `json:"original_launch"`
	OriginalLaunchUnix int       `json:"original_launch_unix"`
	Missions           []struct {
		Name   string `json:"name"`
		Flight int    `json:"flight"`
	} `json:"missions"`
	Landings   int    `json:"landings"`
	Type       string `json:"type"`
	Details    string `json:"details"`
	ReuseCount int    `json:"reuse_count"`
}

var (
	url = "https://api.spacexdata.com/v3/capsules"
)

func main() {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("request failed")
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error in  client call")
		return
	}

	var record SpaceX
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}
	for launchNo, launchData := range record {
		fmt.Println("Capsult Record     =\n", launchData)
		fmt.Println("Record Number      =", launchNo)
	}
}

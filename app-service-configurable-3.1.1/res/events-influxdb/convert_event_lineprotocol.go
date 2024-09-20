package main

import (
	"encoding/json"
	"flag"
	"fmt"
)

// Event represents the structure of the event data payload
type Event struct {
	APIVersion string `json:"apiVersion"`
	RequestID  string `json:"requestId"`
	Event      struct {
		APIVersion  string `json:"apiVersion"`
		ID          string `json:"id"`
		DeviceName  string `json:"deviceName"`
		ProfileName string `json:"profileName"`
		SourceName  string `json:"sourceName"`
		Origin      int64  `json:"origin"`
		Readings    []struct {
			ID           string `json:"id"`
			Origin       int64  `json:"origin"`
			DeviceName   string `json:"deviceName"`
			ResourceName string `json:"resourceName"`
			ProfileName  string `json:"profileName"`
			ValueType    string `json:"valueType"`
			Value        string `json:"value"`
		} `json:"readings"`
	} `json:"event"`
}

func main() {
	// Define the input flag
	input := flag.String("input", "", "Event data payload in JSON format")
	flag.Parse()

	if *input == "" {
		fmt.Println("Please provide the event data payload using the -input flag.")
		return
	}

	// Parse the input JSON
	var event Event
	err := json.Unmarshal([]byte(*input), &event)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Convert to line protocol format
	for _, reading := range event.Event.Readings {
		fmt.Printf("\nJson has been parsed, displaying payload in Line protocol format::\n")
		lineProtocol := fmt.Sprintf("%s,%s=%s,%s=%s %s=%s %d",
			event.Event.ProfileName,
			"deviceName", reading.DeviceName,
			"resourceName", reading.ResourceName,
			"value", reading.Value,
			event.Event.Origin)
		fmt.Println(lineProtocol)
	}
}

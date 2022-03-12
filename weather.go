package main

import (
	"fmt"
	"log"
	"net/http"
)

// Api endpoints
const (
	base = "https://api.weather.gov/"
)

// Make API call, response is returned
func getRequest(inquiry string) *http.Response {
	// Create client
	client := &http.Client{}
	// Build request
	req, err := http.NewRequest("GET", inquiry, nil)
	if err != nil {
		log.Println(err)
	}
	// Add headers to request
	req.Header.Add("User-Agent", "github.com/rBigChill")
	// Make Request
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	return resp
}

func main() {
	resp := getRequest(base)
	fmt.Println(resp.Status)
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// Make API call, response is returned
func getRequest(url string) *http.Response {
	// Create client
	client := &http.Client{}
	// Build request
	req, err := http.NewRequest("GET", url, nil)
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
	// Create url
	url := url.URL{
		Scheme: "https",
		Host:   "api.weather.gov",
	}
	resp := getRequest(url.String())
	fmt.Println(resp.Status)
}

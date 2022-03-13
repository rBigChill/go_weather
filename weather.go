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
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
	}

	// Add headers to request
	req.Header.Add("User-Agent", "github.com/rBigChill")

	// Make Request, Get Response
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
	}
	return response
}

func makeURL(endpoint string) string {

	// Create url
	url := url.URL{
		Scheme: "https",
		Host:   "api.weather.gov",
	}
	response := getRequest(url.String())
	return response
}

func main() {

	endpoint := "points"
	response := makeURL(endpoint)
	fmt.Println(response.Status)
}

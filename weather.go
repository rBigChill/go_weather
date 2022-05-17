package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// Create url
func makeURL(endpoint string) *url.URL {
	url := &url.URL{
		Scheme: "https",
		Host:   "api.weather.gov",
		Path:   endpoint,
	}
	return url
}

// Make API call, response is returned
func getRequest(url *url.URL) *http.Response {
	// Create client
	client := &http.Client{}
	// Build request
	request, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		log.Println(err)
	}
	// Add headers to request
	request.Header.Add("User-Agent", "github.com/rBigChill")
	// Make Request, Get Response
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
	}
	return response
}

// Points endpoint
func points() *http.Response {
	endpoint := "points/32.5115,-94.7964"
	url := makeURL(endpoint)
	response := getRequest(url)
	return response
}

func main() {

	response := points()
	fmt.Println(response)

}

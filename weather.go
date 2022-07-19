package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

// Base url
func baseURL(endpoint string) *url.URL {
	url := &url.URL{
		Scheme: "https",
		Host:   "api.weather.gov",
		Path:   endpoint,
	}
	return url
}

// Create client
func createClient() *http.Client {
	client := &http.Client{}
	return client
}

// Build request
func buildRequest(url *url.URL) *http.Request {
	request, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		log.Println(err)
	}
	// Add Headers
	request.Header.Add("User-Agent", "github.com/rBigChill")
	return request
}

// Make Request
func makeRequest(request *http.Request) *http.Response {
	// Create Client
	client := createClient()
	// Make Request
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
	}
	return response
}

// Make API call, response is returned
func requestType(endpoint string) *http.Response {
	// Create URL
	url := baseURL(endpoint)
	// Build request
	request := buildRequest(url)
	// Make Request
	response := makeRequest(request)
	return response
}

func main() {
	endpoint := "glossary"
	response := requestType(endpoint)
	x, _ := io.ReadAll(response.Body)
	fmt.Printf("%s", x)
}

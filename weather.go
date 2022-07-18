package main

import (
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
func buildURL(url *url.URL) *http.Request {
	request, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		log.Println(err)
	}
	// Add headers to request
	request.Header.Add("User-Agent", "github.com/rBigChill")
	return request
}

// Make API call, response is returned
func createRequest(endpoint *url.URL) *http.Response {
	// Build request
	request := buildURL(endpoint)

	// Create client
	client := createClient()
	// Make Request, Store Response
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
	}
	return response
}

func getRequest() {
}

func main() {

}

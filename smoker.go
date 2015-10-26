package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Page struct {
	Code int    `json:"code"`
	URL  string `json:"url"`
	SSL  bool   `json:"ssl"`
}

// Returns a slice of all the pages in the json urls file
func getPages(urlFile string) ([]*Page, error) {
	file, err := os.Open(urlFile)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	// Decode the json file into a slice of pointers to Page values.
	var pages []*Page
	err = json.NewDecoder(file).Decode(&pages)

	return pages, err
}

// Return an http.Response and duration of the request for a single page
func getUrl(url string, cookies string) (*http.Response, time.Duration) {
	time_start := time.Now()

	// Cannot use http.Get as this always follows redirects
	// which is undesirable
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("cookie", cookies)

	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		log.Printf("Error fetching: %v", err)
	}

	defer resp.Body.Close()

	requestTime := time.Since(time_start)
	return resp, requestTime
}

func main() {
	urlFile := flag.String("file", "urls.json", "File that contains URLs to test")
	cookies := flag.String("cookies", "", "List of cookies to send, semi-colon separated")
	flag.Parse()

	pages, err := getPages(*urlFile)
	if err != nil {
		log.Fatal("Error getting pages: %v", err)
	}

	failures := 0

	for _, page := range pages {
		response, time := getUrl(page.URL, *cookies)
		responseCode := response.StatusCode

		if responseCode != page.Code {
			fmt.Printf("%v %v %v (FAIL: Status should have been %v)\n", page.URL, responseCode, time, page.Code)
			failures++
		} else {
			fmt.Printf("%v %v %v\n", page.URL, responseCode, time)
		}
	}

	// Exit non-zero if any pages fail or respond with unexpected status code
	if failures > 0 {
		if failures == 1 {
			fmt.Printf("\nThere was %v page failure\n", failures)
		} else {
			fmt.Printf("\nThere were %v page failures\n", failures)
		}
		os.Exit(1)
	}
}

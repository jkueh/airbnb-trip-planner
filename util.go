package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func extractListingID(url string) string {
	re := regexp.MustCompile("/rooms/([0-9]+)")
	match := re.FindStringSubmatch(url)
	return match[1]
}

var baseURL = "https://www.airbnb.com.au/api/v2"
var apiKey string

func makeAPIQuery(url string, params []string) []byte {
	url = addQueryStringsToURL(baseURL+"/"+url, append(
		append([]string{"key=" + apiKey}, params...),
	))
	if debug {
		log.Println("Attempting a call to: " + url)
	}
	return getResponseBody(url)
}

func addQueryStringsToURL(url string, params []string) string {
	var buffer bytes.Buffer
	buffer.WriteString(url)
	if len(params) > 0 {
		var separator string
		for index, param := range params {
			if index == 0 {
				separator = "?"
			} else {
				separator = "&"
			}
			buffer.WriteString(separator + param)
		}
	}
	return buffer.String()
}

func getRandomUserAgent() string {
	userAgents := []string{
		fmt.Sprintf(
			"%s %s",
			"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko)",
			"Ubuntu Chromium/71.0.3578.98 Chrome/71.0.3578.98 Safari/537.36",
		),
		fmt.Sprintf(
			"%s %s",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
			"(KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36",
		),
	}
	return userAgents[rand.Intn(len(userAgents))]
}
func getNewRequest(url string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)
	errExit(err)
	req.Header.Set("X-Airbnb-API-Key", apiKey)
	req.Header.Set("User-Agent", getRandomUserAgent())
	return req
}

func getResponseBody(url string) []byte {
	retryCount := 0
	retryLimit := 15
	retrySleepMin := 500
	retrySleepMax := 2500

	client := &http.Client{}

	rand.Seed(time.Now().Unix())
	resp, err := client.Do(getNewRequest(url))
	for resp.StatusCode != 200 {
		rand.Seed(time.Now().UnixNano())
		if retryCount > retryLimit {
			break
		}
		if debug {
			log.Println("Response Code:", resp.StatusCode)
		}
		retryCount++
		sleepDur := time.Duration(
			rand.Intn(retrySleepMax-retrySleepMin)+retrySleepMin,
		) * time.Millisecond
		if debug {
			log.Println("Waiting for", sleepDur)
		}
		time.Sleep(sleepDur)
		resp, err = client.Do(getNewRequest(url))
	}

	if resp.StatusCode != 200 {
		errExit(errors.New("Unable to get listing information"))
	}

	errExit(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	errExit(err)
	return body
}
func amenityExists(amenityList []ListingAmenity, needle string) bool {
	for _, amenity := range amenityList {
		stringsMatch := strings.ToLower(amenity.Name) == strings.ToLower(needle)
		if stringsMatch && amenity.IsPresent {
			return true
		}
	}
	return false
}
func errExit(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

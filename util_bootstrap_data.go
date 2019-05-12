package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func getListingPage(ID int) string {
	listingID := strconv.Itoa(ID)
	req := getNewGETRequest("https://www.airbnb.com.au/rooms/" + listingID)
	return string(getResponseBody(*req))
}

func getListingBootstrap(ID int) (BootstrapPayload, error) {
	data := getListingPage(ID)
	re := regexp.MustCompile(
		"<script data-state=\"true\" type=\"application/json\"><!--(.*?)--></script>",
	)
	match := re.FindStringSubmatch(data)
	bootstrapPayload := BootstrapPayload{}
	if len(match) > 1 {
		json.Unmarshal([]byte(match[1]), &bootstrapPayload)
	} else {
		if verbose {
			log.Println("[ERROR] getListingBootstrap() - No Regex match.")
		}
		if debug {
			log.Println("Page payload below:")
			fmt.Println(string(data))
			log.Fatalln("DEBUG MODE EXIT")
		}
		return bootstrapPayload, errors.New("No regex match for bootstrap data")
	}
	return bootstrapPayload, nil
}

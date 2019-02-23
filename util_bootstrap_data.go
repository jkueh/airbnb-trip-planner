package main

import (
	"encoding/json"
	"log"
	"regexp"
	"strconv"
)

func getListingBootstrap(ID int) BootstrapPayload {
	listingID := strconv.Itoa(ID)
	re := regexp.MustCompile("><!--({\"bootstrapData\".*})--></script>")
	data := getResponseBody("https://www.airbnb.com.au/rooms/" + listingID)
	match := re.FindStringSubmatch(string(data))
	bootstrapPayload := BootstrapPayload{}
	if len(match) > 1 {
		json.Unmarshal([]byte(match[1]), &bootstrapPayload)
	} else {
		if debug {
			log.Fatalln("getListingBootstrap() - No Regex match.")
		} else {
			log.Fatalln(
				"An error occurred while trying to fetch further listing",
				"information.",
			)
		}
	}
	return bootstrapPayload
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	yaml "gopkg.in/yaml.v2"
)

var config Config
var verbose bool
var debug bool

var gracefulStop = make(chan os.Signal, 1)
var gracefulStopRequested = false

func init() {

	// Set up the SIGTERM and SIGINT notifiers
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		sig := <-gracefulStop
		log.Printf("Received signal %+v - Waiting for job to finish...\n", sig)
		gracefulStopRequested = true
	}()

	verboseString := strings.ToLower(os.Getenv("VERBOSE"))
	debugString := strings.ToLower(os.Getenv("DEBUG"))

	if debugString == "true" {
		verbose = true
		debug = true
	} else if verboseString == "true" {
		verbose = true
	}

	apiKey = os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatalln("API Key not specified.")
	}
	configFile := os.Getenv("TRIP_FILE")
	if configFile == "" {
		configFile = "my_trip.yml"
		log.Println("No TRIP_FILE specified. Defaulting to:", configFile)
	} else {
		if debug {
			log.Println("Using config file: " + configFile)
		}
	}
	fileContents, err := ioutil.ReadFile(configFile)
	errExit(err)
	errExit(yaml.Unmarshal(fileContents, &config))
}

func main() {

	//	Set up the array that we're going to pass to the CSV writer later
	outputData := [][]string{
		append(
			[]string{
				"Search",
				"ID",
				"Name",
				"City",
				"RoomAndPropertyType",
				"RoomType",
				"RoomTypeCategory",
				"SpaceType",
				"Link",
				"Nightly Cost",
				"Total Cost",
				"BathroomLabel",
				"Bathrooms",
				"BedLabel",
				"BedroomLabel",
				"Bedrooms",
				"Beds",
				"ApproxLat",
				"ApproxLong",
			},
			config.Amenities...,
		),
	}

	log.Println("Beginning search...")
	for _, search := range config.Searches {
		exploreTabs := getAllExploreTabs(search)
		searchListings := mergeAllListingsFromExploreTabs(exploreTabs)
		log.Println(
			"Found",
			len(searchListings),
			"listings for search:",
			search.Name,
		)
		for searchListingIndex, searchListing := range searchListings {
			if gracefulStopRequested {
				log.Println("See ya later!")
				os.Exit(0)
			}
			listingOutput := []string{search.Name}
			searchResultListing := searchListing
			listing := getBookingDetails(
				strconv.Itoa(searchResultListing.ID),
				search.Dates.CheckIn,
				search.Dates.CheckOut,
			)
			log.Printf(
				"Processing listing %3d of %3d: %s\n",
				searchListingIndex+1,
				len(searchListings),
				searchResultListing.Name,
			)

			if len(listing.PdpListingBookingDetails) != 1 {
				log.Println(
					"\tUnable to retrieve listing information from API for listing",
					searchResultListing.ID,
				)
				continue
			}
			if listing.PdpListingBookingDetails[0].Available == false {
				if verbose {
					log.Println("Not available:", searchListing.ID)
				}
				continue
			}

			listingBootstrapData := getListingBootstrap(searchListing.ID)
			homePDP := listingBootstrapData.BootstrapData.ReduxData.HomePDP
			listingData := homePDP.ListingInfo.Listing
			bookingDetails := listing.PdpListingBookingDetails[0]

			stayCost := 0.00
			for _, total := range bookingDetails.Price.PriceItems {
				stayCost += total.Total.Amount
			}

			amenityCount := len(listingData.ListingAmenities)

			// Add data to the listingOutput array
			url := addQueryStringsToURL(
				fmt.Sprintf(
					"https://www.airbnb.com.au/rooms/%s",
					listing.PdpListingBookingDetails[0].ID,
				),
				[]string{
					fmt.Sprintf("guests=%d", search.GetTotalGuests()),
					fmt.Sprintf("adults=%d", search.Guests.Adults),
					fmt.Sprintf("children=%d", search.Guests.Children),
					fmt.Sprintf("infants=%d", search.Guests.Infants),
					"check_in=" + search.Dates.CheckIn,
					"check_out=" + search.Dates.CheckOut,
				},
			)
			nightlyCost := bookingDetails.BasePriceBreakdown[0].Amount
			listingOutput = append(
				listingOutput, []string{
					fmt.Sprintf("%d", searchResultListing.ID),
					searchResultListing.Name,
					searchResultListing.City,
					searchResultListing.RoomAndPropertyType,
					searchResultListing.RoomType,
					searchResultListing.RoomTypeCategory,
					searchResultListing.SpaceType,
					url,
					fmt.Sprintf("%.2f", nightlyCost),
					fmt.Sprintf("%.2f", stayCost),
					searchResultListing.BathroomLabel,
					fmt.Sprintf("%f", searchResultListing.Bathrooms),
					searchResultListing.BedLabel,
					searchResultListing.BedroomLabel,
					fmt.Sprintf("%d", searchResultListing.Bedrooms),
					fmt.Sprintf("%d", searchResultListing.Beds),
					fmt.Sprintf("%f", searchResultListing.Lat),
					fmt.Sprintf("%f", searchResultListing.Lng),
				}...,
			)
			if verbose { // Print out primitive listing information
				log.Println("Found:", searchResultListing.ID)
				log.Println("\tNightly Cost:", nightlyCost)
				log.Println("\tTotal Cost:", stayCost)
				log.Println("\tName:", searchResultListing.Name)
				log.Println("\tCity:", searchResultListing.City)
				log.Println("\tBedrooms:", searchResultListing.Bedrooms)
			}
			for _, amenity := range config.Amenities {
				if amenityExists(listingData.ListingAmenities, amenity) {
					listingOutput = append(
						listingOutput,
						"TRUE",
					)
					if verbose {
						log.Println("\t"+amenity+":", "TRUE")
					}
				} else {
					listingOutput = append(
						listingOutput,
						"FALSE",
					)
					if verbose {
						log.Println("\t"+amenity+":", "FALSE")
					}
				}
			}
			if verbose {
				log.Println("\tURL:", url)
			}
			outputData = append(outputData, listingOutput)
			if debug && amenityCount > 0 {
				log.Printf("\tAmentities (%d):\n", amenityCount)
				for _, amenity := range listingData.ListingAmenities {
					log.Println("\t- " + amenity.Name)
				}
			}
		}
		writeLineToCSVFile(config.OutputFile, outputData)
	}
}

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

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
				"Processing listing %3d of %3d: [%d] %s\n",
				searchListingIndex+1,
				len(searchListings),
				searchResultListing.ID,
				searchResultListing.Name,
			)

			if len(listing.PdpListingBookingDetails) != 1 {
				log.Println(
					"\tUnable to retrieve listing information from API for listing",
					searchResultListing.ID,
				)
				continue
			}

			if listing.PdpListingBookingDetails[0].Available != true {
				if verbose {
					log.Println("Not available:", searchListing.ID)
				}
				continue
			}

			listingBootstrapData, err := getListingBootstrap(searchListing.ID)

			if err != nil {
				if verbose {
					log.Println(err)
				}
				log.Println(
					"Skipping listing",
					searchListing.ID,
					"due to missing bootstrap data.",
				)
				continue
			}
			// If the BootstrapData is blank, then it's likely that the property isn't
			// available for those dates.
			if listingBootstrapData.IsBlank() {
				log.Println(
					"Skipping listing",
					searchListing.ID,
					"due to missing bootstrap data.",
				)
				if verbose {
					log.Println("BootstrapData was received, but was likely empty.")
				}
				if debug {
					log.Fatalln(getListingPage(searchListing.ID))
				}
				continue
			}

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
				log.Printf("\tAmenities (%d):\n", amenityCount)
				for _, amenity := range listingData.ListingAmenities {
					log.Println("\t- " + amenity.Name)
				}
			}
		}
		writeLineToCSVFile(config.OutputFile, outputData)
	}
}

package main

import (
	"encoding/json"
	"fmt"
)

func getAllExploreTabs(search Search) []ExploreTabs {
	returnArray := []ExploreTabs{}
	itemsPerPage := 50

	//	Do the initial query
	initialSearchListings := getSearchListings(search, itemsPerPage, 0)
	returnArray = append(returnArray, initialSearchListings)
	hasNextPage := initialSearchListings.ExploreTabs[0].PaginationMetadata.HasNextPage
	currentOffset := itemsPerPage
	if hasNextPage == true {
		for hasNextPage != false {
			returnData := getSearchListings(search, itemsPerPage, currentOffset)
			returnArray = append(returnArray, returnData)
			currentOffset += itemsPerPage
			hasNextPage = returnData.ExploreTabs[0].PaginationMetadata.HasNextPage
		}
	}
	return returnArray
}

func mergeAllListingsFromExploreTabs(tabs []ExploreTabs) []ExploreTabsListing {
	var returnArray []ExploreTabsListing
	for _, exploreTab := range tabs {
		sectionLength := len(exploreTab.ExploreTabs[0].Sections)
		listings := exploreTab.ExploreTabs[0].Sections[sectionLength-1].Listings
		for _, listing := range listings {

			// A little something to prevent duplicate entries
			if len(returnArray) > 1 {
				for _, existingListing := range returnArray {
					if listing.Listing.ID == existingListing.ID {
						continue
					}
				}
			}

			// Append to the array
			returnArray = append(returnArray, listing.Listing)
		}
	}
	return returnArray
}

func getSearchListings(search Search, itemsPerPage int, offset int) ExploreTabs {
	// 50 items per grid Seems to be the artificial limit at the API end for now
	var params = []string{
		fmt.Sprintf("guests=%d", search.GetTotalGuests()),
		fmt.Sprintf("adults=%d", search.Guests.Adults),
		fmt.Sprintf("children=%d", search.Guests.Children),
		fmt.Sprintf("infants=%d", search.Guests.Infants),
		"_format=for_explore_search_web",
		"check_in=" + search.Dates.CheckIn,
		"check_out=" + search.Dates.CheckOut,
		"search_by_map=true",
		fmt.Sprintf("sw_lat=%f", search.SouthWest[0]),
		fmt.Sprintf("sw_lng=%f", search.SouthWest[1]),
		fmt.Sprintf("ne_lat=%f", search.NorthEast[0]),
		fmt.Sprintf("ne_lng=%f", search.NorthEast[1]),
		fmt.Sprintf("items_per_grid=%d", itemsPerPage),
		fmt.Sprintf("items_offset=%d", offset),
		"query=" + search.SearchQuery,
		"currency=AUD",
		"locale=en-AU",
	}
	var exploreResults ExploreTabs
	resp := makeAPIQuery("explore_tabs", params)
	errExit(json.Unmarshal(resp, &exploreResults))
	return exploreResults
}

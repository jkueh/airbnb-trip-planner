package main

import (
	"encoding/json"
)

func getBookingDetails(listingID string, checkIn string, checkOut string) BookingDetails {
	var params = []string{
		"guests=1",
		"listing_id=" + listingID,
		"_format=for_web_with_date",
		"check_in=" + checkIn,
		"check_out=" + checkOut,
		// "number_of_adults=1",
		// "number_of_children=0",
		// "number_of_infants=0",
		"currency=AUD",
		"locale=en-AU",
		"show_smart_promotion=0",
		// "_parent_request_uuid=2b672287-44ed-40b7-8d35-6dbfa2f7cfe3",
	}
	var bookingDetails BookingDetails
	resp := makeAPIQuery("pdp_listing_booking_details", params)
	errExit(json.Unmarshal(resp, &bookingDetails))
	return bookingDetails
}

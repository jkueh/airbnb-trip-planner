package main

// BootstrapPayload is the json blob that AirBnB injects into the page
type BootstrapPayload struct {
	BootstrapData struct {
		ReduxData struct {
			HomePDP struct {
				ListingInfo struct {
					Listing struct {
						ListingAmenities []ListingAmenity `json:"listing_amenities"`
					} `json:"listing"`
				} `json:"listingInfo"`
			} `json:"homePDP"`
		} `json:"reduxData"`
	} `json:"bootstrapData"`
}

// ListingAmenity contains all the goodness that is amenity information.
type ListingAmenity struct {
	Category               interface{} `json:"category"`
	Icon                   string      `json:"icon"`
	ID                     int         `json:"id"`
	IsBusinessReadyFeature bool        `json:"is_business_ready_feature"`
	IsPresent              bool        `json:"is_present"`
	IsSafetyFeature        bool        `json:"is_safety_feature"`
	Name                   string      `json:"name"`
	SelectListViewPhoto    interface{} `json:"select_list_view_photo"`
	SelectTileViewPhoto    interface{} `json:"select_tile_view_photo"`
	Tag                    string      `json:"tag"`
	Tooltip                string      `json:"tooltip"`
}

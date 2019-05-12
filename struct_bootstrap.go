package main

// BootstrapPayload is the json blob that AirBnB injects into the page
type BootstrapPayload struct {
	BootstrapData struct {
		CanonicalURL string `json:"canonical_url"`
		ReduxData    struct {
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

// IsBlank returns a true/false value if the Canonical URL is blank.
func (b BootstrapPayload) IsBlank() bool {
	amenitiesCount := len(
		b.BootstrapData.ReduxData.HomePDP.ListingInfo.Listing.ListingAmenities,
	)
	urlBlank := b.BootstrapData.CanonicalURL == ""
	return amenitiesCount <= 0 && urlBlank
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

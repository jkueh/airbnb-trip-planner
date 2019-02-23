package main

// Generated from https://mholt.github.io/json-to-go/

// ExploreTabs is the payload that comes back from the API call.
type ExploreTabs struct {
	ExploreTabs []struct {
		TabID              string `json:"tab_id"`
		TabName            string `json:"tab_name"`
		PaginationMetadata struct {
			HasNextPage     bool   `json:"has_next_page"`
			ItemsOffset     int    `json:"items_offset"`
			SectionOffset   int    `json:"section_offset"`
			SearchSessionID string `json:"search_session_id"`
		} `json:"pagination_metadata"`
		Sections []struct {
			BackendSearchID     string        `json:"backend_search_id"`
			DisplayType         string        `json:"display_type"`
			ExperimentsMetadata []interface{} `json:"experiments_metadata"`
			ResultType          string        `json:"result_type"`
			SearchSessionID     string        `json:"search_session_id"`
			SectionID           string        `json:"section_id"`
			SectionTypeUID      string        `json:"section_type_uid"`
			IsPaginated         bool          `json:"is_paginated"`
			BankaiSectionID     string        `json:"bankai_section_id"`
			Refinements         []interface{} `json:"refinements"`
			Messages            []struct {
				Style string `json:"style"`
				Title string `json:"title"`
			} `json:"messages,omitempty"`
			ReviewItems     []interface{} `json:"review_items"`
			Breadcrumbs     []interface{} `json:"breadcrumbs"`
			SectionMetadata struct {
			} `json:"section_metadata"`
			NearbyLocations []interface{} `json:"nearby_locations"`
			Title           string        `json:"title,omitempty"`
			Listings        []struct {
				Listing      ExploreTabsListing `json:"listing"`
				PricingQuote struct {
					CanInstantBook     bool    `json:"can_instant_book"`
					MonthlyPriceFactor float64 `json:"monthly_price_factor"`
					PriceString        string  `json:"price_string"`
					Rate               struct {
						Amount           float64 `json:"amount"`
						AmountFormatted  string  `json:"amount_formatted"`
						Currency         string  `json:"currency"`
						IsMicrosAccuracy bool    `json:"is_micros_accuracy"`
					} `json:"rate"`
					RateType           string `json:"rate_type"`
					RateWithServiceFee struct {
						Amount           float64 `json:"amount"`
						AmountFormatted  string  `json:"amount_formatted"`
						Currency         string  `json:"currency"`
						IsMicrosAccuracy bool    `json:"is_micros_accuracy"`
					} `json:"rate_with_service_fee"`
					WeeklyPriceFactor   float64 `json:"weekly_price_factor"`
					ShouldShowFromLabel bool    `json:"should_show_from_label"`
				} `json:"pricing_quote"`
				Verified struct {
					Enabled            bool   `json:"enabled"`
					BadgeText          string `json:"badge_text"`
					BadgeSecondaryText string `json:"badge_secondary_text"`
					KickerBadgeType    string `json:"kicker_badge_type"`
				} `json:"verified"`
				VerifiedCard bool `json:"verified_card"`
			} `json:"listings,omitempty"`
			LocalizedListingCount string `json:"localized_listing_count,omitempty"`
		} `json:"sections"`
		ExperimentsMetadata []interface{} `json:"experiments_metadata"`
		HomeTabMetadata     struct {
			UrgencyCommitment struct {
				MessageType string `json:"message_type"`
				Message     struct {
				} `json:"message"`
			} `json:"urgency_commitment"`
			GoldenTicketUrgencyCommitment struct {
				MessageType string `json:"message_type"`
				Message     struct {
				} `json:"message"`
			} `json:"golden_ticket_urgency_commitment"`
			ListingCardsUrgencyCommitmentMetadata          []interface{} `json:"listing_cards_urgency_commitment_metadata"`
			ListingCardsPriceLineUrgencyCommitmentMetadata []interface{} `json:"listing_cards_price_line_urgency_commitment_metadata"`
			Messages                                       struct {
			} `json:"messages"`
			Facets struct {
			} `json:"facets"`
			Overrides struct {
			} `json:"overrides"`
			SearchFeedItems []interface{} `json:"search_feed_items"`
			RemarketingIds  []int         `json:"remarketing_ids"`
			Location        struct {
				CanonicalLocation string `json:"canonical_location"`
				DisplayLocation   string `json:"display_location"`
			} `json:"location"`
			Breadcrumbs []struct {
				LocationName          string `json:"location_name"`
				CanonicalLocationName string `json:"canonical_location_name"`
				Type                  string `json:"type"`
			} `json:"breadcrumbs"`
			ListingsCount int `json:"listings_count"`
			Search        struct {
				NativeCurrency           string `json:"native_currency"`
				PriceRangeMaxNative      int    `json:"price_range_max_native"`
				PriceRangeMinNative      int    `json:"price_range_min_native"`
				PriceType                string `json:"price_type"`
				SearchID                 string `json:"search_id"`
				MobileSessionID          string `json:"mobile_session_id"`
				IsBusinessTravelVerified bool   `json:"is_business_travel_verified"`
				BusinessTravelReadyData  struct {
					FilterCriteria struct {
					} `json:"filter_criteria"`
					ShowBtrUpsell bool `json:"show_btr_upsell"`
				} `json:"business_travel_ready_data"`
				Limit int `json:"limit"`
			} `json:"search"`
			Geography struct {
				Accuracy          int     `json:"accuracy"`
				City              string  `json:"city"`
				CountryCode       string  `json:"country_code"`
				Country           string  `json:"country"`
				Lat               float64 `json:"lat"`
				Lng               float64 `json:"lng"`
				PlaceID           string  `json:"place_id"`
				Precision         string  `json:"precision"`
				ResultType        string  `json:"result_type"`
				State             string  `json:"state"`
				StateShort        string  `json:"state_short"`
				ExtendedPrecision string  `json:"extended_precision"`
				FullAddress       string  `json:"full_address"`
				Market            string  `json:"market"`
				Province          string  `json:"province"`
				Zip               string  `json:"zip"`
			} `json:"geography"`
			PriceHistogram struct {
				AveragePrice int   `json:"average_price"`
				Histogram    []int `json:"histogram"`
			} `json:"price_histogram"`
			Filters struct {
				Sections []struct {
					Items []struct {
						Selected bool   `json:"selected"`
						Type     string `json:"type"`
						Title    string `json:"title"`
						Params   []struct {
							Key             string `json:"key"`
							Delete          bool   `json:"delete"`
							InvisibleToUser bool   `json:"invisible_to_user"`
							ValueType       string `json:"value_type"`
						} `json:"params"`
					} `json:"items"`
					Selected                  bool          `json:"selected"`
					Title                     string        `json:"title,omitempty"`
					FilterBarTitle            string        `json:"filter_bar_title,omitempty"`
					FilterSectionID           string        `json:"filter_section_id"`
					BarItems                  []interface{} `json:"bar_items"`
					CollapseState             string        `json:"collapse_state,omitempty"`
					Subsections               []interface{} `json:"subsections"`
					ExperimentsMetadata       []interface{} `json:"experiments_metadata"`
					CollapseSubtitle          string        `json:"collapse_subtitle,omitempty"`
					CollapseSelectedSubtitles []interface{} `json:"collapse_selected_subtitles,omitempty"`
					ExpandText                string        `json:"expand_text,omitempty"`
					CollapseText              string        `json:"collapse_text,omitempty"`
				} `json:"sections"`
				MoreFiltersButton struct {
					Text     string `json:"text"`
					Verified bool   `json:"verified"`
					BgColor  string `json:"bg_color"`
				} `json:"more_filters_button"`
				FilterBarOrdering struct {
					Default []string `json:"default"`
					Small   []string `json:"small"`
					Medium  []string `json:"medium"`
				} `json:"filter_bar_ordering"`
				MoreFiltersOrdering struct {
					Default []string `json:"default"`
					Small   []string `json:"small"`
					Medium  []string `json:"medium"`
				} `json:"more_filters_ordering"`
				MoreFiltersCounts struct {
					Default int `json:"default"`
					Small   int `json:"small"`
					Medium  int `json:"medium"`
				} `json:"more_filters_counts"`
				FilterBarCounts struct {
					Default []int `json:"default"`
					Small   []int `json:"small"`
					Medium  []int `json:"medium"`
				} `json:"filter_bar_counts"`
				ShowLeftFiltersPanel        bool `json:"show_left_filters_panel"`
				EligibleForLeftFiltersPanel bool `json:"eligible_for_left_filters_panel"`
			} `json:"filters"`
		} `json:"home_tab_metadata"`
	} `json:"explore_tabs"`
	Metadata struct {
		CurrentTabID             string        `json:"current_tab_id"`
		RequestUUID              string        `json:"request_uuid"`
		SearchID                 string        `json:"search_id"`
		Query                    string        `json:"query"`
		QueryText                string        `json:"query_text"`
		CurrentRefinementPaths   []string      `json:"current_refinement_paths"`
		FederatedSearchID        string        `json:"federated_search_id"`
		FederatedSearchSessionID string        `json:"federated_search_session_id"`
		Suggestions              []interface{} `json:"suggestions"`
		HasMap                   bool          `json:"has_map"`
		ShowAsHint               bool          `json:"show_as_hint"`
		MarqueeMode              string        `json:"marquee_mode"`
		SatoriVersion            string        `json:"satori_version"`
		GtmExperiments           []interface{} `json:"gtm_experiments"`
		SavedSearchResponses     []interface{} `json:"saved_search_responses"`
		MapToggle                bool          `json:"map_toggle"`
		PriceDisplayStrategy     string        `json:"price_display_strategy"`
		SatoriConfig             struct {
			Market      string `json:"market"`
			StateCode   string `json:"state_code"`
			CountryCode string `json:"country_code"`
			RegionID    int    `json:"region_id"`
			Version     string `json:"version"`
		} `json:"satori_config"`
		PageMetadata struct {
			PageTitle     string `json:"page_title"`
			RenderType    string `json:"render_type"`
			CanonicalURL  string `json:"canonical_url"`
			LocationQuery bool   `json:"location_query"`
		} `json:"page_metadata"`
		ExploreMapMetadata struct {
			Pins []interface{} `json:"pins"`
			Type string        `json:"type"`
		} `json:"explore_map_metadata"`
		Geography struct {
			Accuracy          int     `json:"accuracy"`
			City              string  `json:"city"`
			CountryCode       string  `json:"country_code"`
			Country           string  `json:"country"`
			Lat               float64 `json:"lat"`
			Lng               float64 `json:"lng"`
			PlaceID           string  `json:"place_id"`
			Precision         string  `json:"precision"`
			ResultType        string  `json:"result_type"`
			State             string  `json:"state"`
			StateShort        string  `json:"state_short"`
			ExtendedPrecision string  `json:"extended_precision"`
			FullAddress       string  `json:"full_address"`
			Market            string  `json:"market"`
			Province          string  `json:"province"`
			Zip               string  `json:"zip"`
		} `json:"geography"`
		RefinementDisplayText string `json:"refinement_display_text"`
	} `json:"metadata"`
}

// ExploreTabsListing is the listing information we care about
type ExploreTabsListing struct {
	Badges                []interface{} `json:"badges"`
	BathroomLabel         string        `json:"bathroom_label"`
	Bathrooms             float64       `json:"bathrooms"`
	BedLabel              string        `json:"bed_label"`
	BedroomLabel          string        `json:"bedroom_label"`
	Bedrooms              int           `json:"bedrooms"`
	Beds                  int           `json:"beds"`
	City                  string        `json:"city"`
	GuestLabel            string        `json:"guest_label"`
	HostLanguages         []interface{} `json:"host_languages"`
	HostThumbnailURLSmall string        `json:"host_thumbnail_url_small"`
	HostThumbnailURL      string        `json:"host_thumbnail_url"`
	ID                    int           `json:"id"`
	IsBusinessTravelReady bool          `json:"is_business_travel_ready"`
	IsFullyRefundable     bool          `json:"is_fully_refundable"`
	IsHostHighlyRated     bool          `json:"is_host_highly_rated"`
	IsNewListing          bool          `json:"is_new_listing"`
	IsSuperhost           bool          `json:"is_superhost"`
	KickerContent         struct {
		Messages  []string `json:"messages"`
		TextColor string   `json:"text_color"`
	} `json:"kicker_content"`
	Lat            float64  `json:"lat"`
	Lng            float64  `json:"lng"`
	LocalizedCity  string   `json:"localized_city"`
	Name           string   `json:"name"`
	PersonCapacity int      `json:"person_capacity"`
	PictureCount   int      `json:"picture_count"`
	PictureURL     string   `json:"picture_url"`
	PictureUrls    []string `json:"picture_urls"`
	Picture        struct {
		ID                     int    `json:"id"`
		DominantSaturatedColor string `json:"dominant_saturated_color"`
		LargeRo                string `json:"large_ro"`
		Picture                string `json:"picture"`
		PreviewEncodedPng      string `json:"preview_encoded_png"`
		SaturatedA11YDarkColor string `json:"saturated_a11y_dark_color"`
		ScrimColor             string `json:"scrim_color"`
	} `json:"picture"`
	PreviewAmenities    string  `json:"preview_amenities"`
	PreviewEncodedPng   string  `json:"preview_encoded_png"`
	PropertyTypeID      int     `json:"property_type_id"`
	ReviewsCount        int     `json:"reviews_count"`
	RoomAndPropertyType string  `json:"room_and_property_type"`
	RoomTypeCategory    string  `json:"room_type_category"`
	RoomType            string  `json:"room_type"`
	ScrimColor          string  `json:"scrim_color"`
	ShowStructuredName  bool    `json:"show_structured_name"`
	SpaceType           string  `json:"space_type"`
	StarRating          float64 `json:"star_rating"`
	TierID              int     `json:"tier_id"`
	User                struct {
		FirstName     string `json:"first_name"`
		HasProfilePic bool   `json:"has_profile_pic"`
		ID            int    `json:"id"`
		IsSuperhost   bool   `json:"is_superhost"`
		PictureURL    string `json:"picture_url"`
		SmartName     string `json:"smart_name"`
		ThumbnailURL  string `json:"thumbnail_url"`
	} `json:"user"`
	WideKickerContent struct {
		Messages  []string `json:"messages"`
		TextColor string   `json:"text_color"`
	} `json:"wide_kicker_content"`
	NeighborhoodOverview string        `json:"neighborhood_overview"`
	PublicAddress        string        `json:"public_address"`
	SeoReviews           []interface{} `json:"seo_reviews"`
	Space                string        `json:"space"`
	Summary              string        `json:"summary"`
	AmenityIds           []int         `json:"amenity_ids"`
	PreviewAmenityNames  []string      `json:"preview_amenity_names"`
	Reviews              []interface{} `json:"reviews"`
	StarRatingColor      string        `json:"star_rating_color"`
	PreviewTags          []interface{} `json:"preview_tags"`
	AvgRating            float64       `json:"avg_rating"`
}

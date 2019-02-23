package main

// BookingDetails is the listing information that comes back when we query it.
type BookingDetails struct {
	PdpListingBookingDetails []struct {
		Available                bool `json:"available"`
		AverageRateWithoutTaxUsd int  `json:"average_rate_without_tax_usd"`
		BasePriceBreakdown       []struct {
			Amount           float64 `json:"amount"`
			AmountMicros     float64 `json:"amount_micros"`
			AmountFormatted  string  `json:"amount_formatted"`
			IsMicrosAccuracy bool    `json:"is_micros_accuracy"`
			Currency         string  `json:"currency"`
		} `json:"base_price_breakdown"`
		CanInstantBook       bool `json:"can_instant_book"`
		CancellationPolicies []struct {
			CancellationSectionPlacement    interface{} `json:"cancellation_section_placement"`
			CleaningFee                     interface{} `json:"cleaning_fee"`
			FreeCancellationUpsellBody      interface{} `json:"free_cancellation_upsell_body"`
			FreeCancellationUpsellPlacement interface{} `json:"free_cancellation_upsell_placement"`
			FreeCancellationUpsellTitle     interface{} `json:"free_cancellation_upsell_title"`
			Link                            string      `json:"link"`
			LinkText                        string      `json:"link_text"`
			ServiceFee                      interface{} `json:"service_fee"`
			SubtitleRootForPlus             interface{} `json:"subtitle_root_for_plus"`
			SubtitleShort                   interface{} `json:"subtitle_short"`
			Subtitles                       []string    `json:"subtitles"`
			Timeline                        interface{} `json:"timeline"`
			Title                           string      `json:"title"`
			Milestones                      []struct {
				Titles                   []string    `json:"titles"`
				Subtitles                []string    `json:"subtitles"`
				Type                     string      `json:"type"`
				Color                    string      `json:"color"`
				TimelineLengthPercentage float64     `json:"timeline_length_percentage"`
				TimelineWidthPercentage  interface{} `json:"timeline_width_percentage"`
				IsDated                  interface{} `json:"is_dated"`
			} `json:"milestones"`
			Disclaimer                                string      `json:"disclaimer"`
			FreeCancellationUpsellTitleMobile         interface{} `json:"free_cancellation_upsell_title_mobile"`
			FreeCancellationUpsellMessageBookItModule interface{} `json:"free_cancellation_upsell_message_book_it_module"`
			LocalizedCancellationPolicyName           string      `json:"localized_cancellation_policy_name"`
			CancellationPolicyLabel                   string      `json:"cancellation_policy_label"`
			CancellationPolicyPriceType               interface{} `json:"cancellation_policy_price_type"`
			CancellationPolicyPriceFactor             float64     `json:"cancellation_policy_price_factor"`
			CancellationPolicyID                      int         `json:"cancellation_policy_id"`
			BookItModuleTooltip                       string      `json:"book_it_module_tooltip"`
		} `json:"cancellation_policies"`
		CancellationPolicyLabel  string `json:"cancellation_policy_label"`
		CheckIn                  string `json:"check_in"`
		CheckOut                 string `json:"check_out"`
		CleaningFeeAsGuest       int    `json:"cleaning_fee_as_guest"`
		DepositUpsellMessageData struct {
			UpsellEligible bool   `json:"upsell_eligible"`
			Message        string `json:"message"`
		} `json:"deposit_upsell_message_data"`
		DiscountData  string `json:"discount_data"`
		ExtraGuestFee struct {
			Amount           float64 `json:"amount"`
			AmountMicros     float64 `json:"amount_micros"`
			AmountFormatted  string  `json:"amount_formatted"`
			IsMicrosAccuracy bool    `json:"is_micros_accuracy"`
			Currency         string  `json:"currency"`
		} `json:"extra_guest_fee"`
		GuestDetails struct {
			NumberOfAdults       int    `json:"number_of_adults"`
			NumberOfChildren     int    `json:"number_of_children"`
			NumberOfInfants      int    `json:"number_of_infants"`
			LocalizedDescription string `json:"localized_description"`
		} `json:"guest_details"`
		Guests                          int    `json:"guests"`
		ID                              string `json:"id"`
		LocalizedCancellationPolicyName string `json:"localized_cancellation_policy_name"`
		LocalizedUnavailabilityMessage  string `json:"localized_unavailability_message"`
		Nights                          int    `json:"nights"`
		P3CancellationSection           struct {
			CancellationSectionPlacement    interface{} `json:"cancellation_section_placement"`
			CleaningFee                     interface{} `json:"cleaning_fee"`
			FreeCancellationUpsellBody      interface{} `json:"free_cancellation_upsell_body"`
			FreeCancellationUpsellPlacement interface{} `json:"free_cancellation_upsell_placement"`
			FreeCancellationUpsellTitle     interface{} `json:"free_cancellation_upsell_title"`
			Link                            string      `json:"link"`
			LinkText                        string      `json:"link_text"`
			ServiceFee                      interface{} `json:"service_fee"`
			SubtitleRootForPlus             interface{} `json:"subtitle_root_for_plus"`
			SubtitleShort                   interface{} `json:"subtitle_short"`
			Subtitles                       []string    `json:"subtitles"`
			Timeline                        interface{} `json:"timeline"`
			Title                           string      `json:"title"`
			Milestones                      []struct {
				Titles                   []string    `json:"titles"`
				Subtitles                []string    `json:"subtitles"`
				Type                     string      `json:"type"`
				Color                    string      `json:"color"`
				TimelineLengthPercentage float64     `json:"timeline_length_percentage"`
				TimelineWidthPercentage  interface{} `json:"timeline_width_percentage"`
				IsDated                  interface{} `json:"is_dated"`
			} `json:"milestones"`
			Disclaimer                                string      `json:"disclaimer"`
			FreeCancellationUpsellTitleMobile         interface{} `json:"free_cancellation_upsell_title_mobile"`
			FreeCancellationUpsellMessageBookItModule interface{} `json:"free_cancellation_upsell_message_book_it_module"`
			LocalizedCancellationPolicyName           string      `json:"localized_cancellation_policy_name"`
			CancellationPolicyLabel                   string      `json:"cancellation_policy_label"`
			CancellationPolicyPriceType               interface{} `json:"cancellation_policy_price_type"`
			CancellationPolicyPriceFactor             float64     `json:"cancellation_policy_price_factor"`
			CancellationPolicyID                      int         `json:"cancellation_policy_id"`
			BookItModuleTooltip                       string      `json:"book_it_module_tooltip"`
		} `json:"p3_cancellation_section"`
		P3DisplayRate struct {
			Amount           float64 `json:"amount"`
			AmountMicros     float64 `json:"amount_micros"`
			AmountFormatted  string  `json:"amount_formatted"`
			IsMicrosAccuracy bool    `json:"is_micros_accuracy"`
			Currency         string  `json:"currency"`
		} `json:"p3_display_rate"`
		P3MessageData struct {
			MessageType string `json:"message_type"`
			Message     struct {
				Icon              string      `json:"icon"`
				Headline          string      `json:"headline"`
				Body              string      `json:"body"`
				ContextualMessage interface{} `json:"contextualMessage"`
			} `json:"message"`
			Action struct {
				ConfirmationMessageType interface{} `json:"confirmation_message_type"`
				Label                   interface{} `json:"label"`
				Payload                 struct {
					CheckIn  interface{} `json:"check_in"`
					CheckOut interface{} `json:"check_out"`
				} `json:"payload"`
				Type interface{} `json:"type"`
			} `json:"action"`
		} `json:"p3_message_data"`
		P3PercentageRecommended interface{} `json:"p3_percentage_recommended"`
		Price                   struct {
			DiscountData struct {
				TieredPricingDiscountData  interface{} `json:"tiered_pricing_discount_data"`
				ChinaDiscountPromotionData interface{} `json:"china_discount_promotion_data"`
				PricingDiscountData        interface{} `json:"pricing_discount_data"`
			} `json:"discount_data"`
			PriceItems []struct {
				DiscountData         interface{} `json:"discount_data"`
				LocalizedExplanation interface{} `json:"localized_explanation"`
				LocalizedTitle       string      `json:"localized_title"`
				Total                struct {
					Amount           float64 `json:"amount"`
					AmountMicros     float64 `json:"amount_micros"`
					AmountFormatted  string  `json:"amount_formatted"`
					IsMicrosAccuracy bool    `json:"is_micros_accuracy"`
					Currency         string  `json:"currency"`
				} `json:"total"`
				Type string `json:"type"`
			} `json:"price_items"`
			Total struct {
				Amount           float64 `json:"amount"`
				AmountMicros     float64 `json:"amount_micros"`
				AmountFormatted  string  `json:"amount_formatted"`
				IsMicrosAccuracy bool    `json:"is_micros_accuracy"`
				Currency         string  `json:"currency"`
			} `json:"total"`
		} `json:"price"`
		PriceDisclaimer         string `json:"price_disclaimer"`
		PricingQuoteRequestUUID string `json:"pricing_quote_request_uuid"`
		PrivacySettings         struct {
			ShouldHidePii                      bool   `json:"should_hide_pii"`
			AddressHiddenExplanationTranslated string `json:"address_hidden_explanation_translated"`
		} `json:"privacy_settings"`
		RateType           string `json:"rate_type"`
		RateWithServiceFee struct {
			Amount           float64 `json:"amount"`
			AmountMicros     float64 `json:"amount_micros"`
			AmountFormatted  string  `json:"amount_formatted"`
			IsMicrosAccuracy bool    `json:"is_micros_accuracy"`
			Currency         string  `json:"currency"`
		} `json:"rate_with_service_fee"`
		SecondaryDisplayRateData interface{} `json:"secondary_display_rate_data"`
		SecurityDepositDetails   struct {
			FormattedPrice             interface{} `json:"formatted_price"`
			LocalizedAuthorizationTime interface{} `json:"localized_authorization_time"`
			SecurityDeposit            interface{} `json:"security_deposit"`
			AdditionalInfoTitle        interface{} `json:"additional_info_title"`
			AdditionalInfoDetails      interface{} `json:"additional_info_details"`
		} `json:"security_deposit_details"`
		ShouldShowFromLabel bool `json:"should_show_from_label"`
		TaxAmountUsd        int  `json:"tax_amount_usd"`
	} `json:"pdp_listing_booking_details"`
	Metadata struct {
	} `json:"metadata"`
}

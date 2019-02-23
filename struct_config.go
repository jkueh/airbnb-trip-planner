package main

// Config represents the entire YAML file
type Config struct {
	OutputFile    string `yaml:"output_file"`
	FirstDayDates struct {
		CheckIn  string `yaml:"check_in"`
		CheckOut string `yaml:"check_out"`
	} `yaml:"first_day_dates"`
	Searches  []Search  `yaml:"searches"`
	Listings  []Listing `yaml:"listings"`
	Amenities []string  `yaml:"amenities"`
}

// Search is a single instance of a search, with all its parameters
type Search struct {
	Name  string `yaml:"name"`
	Dates struct {
		CheckIn  string `yaml:"check_in"`
		CheckOut string `yaml:"check_out"`
	} `yaml:"dates"`
	SearchQuery string    `yaml:"search_query"`
	SouthWest   []float64 `yaml:"sw"`
	NorthEast   []float64 `yaml:"ne"`
	Guests      struct {
		Adults   int `yaml:"adults"`
		Children int `yaml:"children"`
		Infants  int `yaml:"infants"`
	} `yaml:"guests"`
}

// GetTotalGuests returns the total number of guests
func (search Search) GetTotalGuests() int {
	guestCount := 0
	guestCount += search.Guests.Adults
	guestCount += search.Guests.Children
	guestCount += search.Guests.Infants
	return guestCount
}

// Listing is the array of listings we're looking at.
type Listing struct {
	URL   string
	Dates struct {
		CheckIn  string `yaml:"check_in"`
		CheckOut string `yaml:"check_out"`
	}
}

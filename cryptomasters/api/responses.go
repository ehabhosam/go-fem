package api

type GetRatesResponse struct {
	// you can only use the fields you need
	// structs can contain metadata (optional)
	// objects can read the metadata and use it
	// json package can see this metadata and use it
	// to convert the json to the struct you want
	Timestamp             string  `json:"timestamp"`
	Low                   string  `json:"low"`
	High                  string  `json:"high"`
	Last                  string  `json:"last"`
	Volume                string  `json:"volume"`
	Volume30D             string  `json:"volume30d"`
	Bid                   float64 `json:"bid"`
	Ask                   float64 `json:"ask"`
	PriceChange           string  `json:"priceChange"`
	PriceChangePercentage string  `json:"priceChangePercentage"`
	Pair                  string  `json:"pair"`
}

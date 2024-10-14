package api

import (
	"cryptomasters-project/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func getConversionApiUrl(currency string) string {
	return fmt.Sprintf(apiUrl, currency)
}

func GetRate(currency string) (*model.Rate, error) {
	if len(currency) != 3 {
		return nil, fmt.Errorf("currency code must be 3 characters")
	}
	// get rates from cex api
	// and return it as a Rate struct
	upCurrency := strings.ToUpper(currency)
	url := getConversionApiUrl(upCurrency)
	// go is synchronous, so it will wait for the response
	// no async, await here :)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusOK {
		// response.body is a readable stream
		// we make sure to read it all
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		// now we have to parse the body
		var response GetRatesResponse
		// this will parse the json and store it in the struct
		// you will get an error if the json is not valid (not same structure as the struct)
		err = json.Unmarshal(bodyBytes, &response)
		if err != nil {
			return nil, err
		}
		cryptoRate := model.Rate{
			Currency: upCurrency,
			Price:    model.Price(response.Bid),
		}
		return &cryptoRate, nil
	}
	return nil, fmt.Errorf("status code: %d", resp.StatusCode)
}

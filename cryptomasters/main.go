package main

import (
	"cryptomasters-project/api"
	"fmt"
	"sync"
)

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)
	if err != nil {
		fmt.Println("Something went wrong trying to get the rate", err)
	} else {
		fmt.Println(rate.String()) // you did not have to do the pointer receiver
	}
}

func main() {
	// lazy way to do it
	// go getCurrencyData("BTC")
	// go getCurrencyData("ETH")
	// go getCurrencyData("XRP")
	// time.Sleep(time.Second * 3)
	wait_groups()
}

func wait_groups() {
	currencies := []string{"BTC", "ETH", "XRP", "LTC", "BCH", "ADA", "XLM", "EOS", "NEO", "TRX"}
	var wg sync.WaitGroup
	for _, currency := range currencies {
		wg.Add(1) // as u r pushing to a stack
		// use closures to pass the wg
		// currency should be passed thru a function
		// to copy its value otherwise only last value will be passed
		go func(currency string) {
			getCurrencyData(currency)
			wg.Done() // as u r popping from a stack
		}(currency)
	}
	wg.Wait() // wait for all the goroutines to finish (when stack is empty)
}

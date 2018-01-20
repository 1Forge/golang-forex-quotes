/*
 * This library is provided without warranty under the MIT license
 * Created by Jacob Davis <jacob@1forge.com>
 */

package main

import (
	Forex "../golang-forex-quotes"
	"fmt"
)

var api_key = "YOUR_API_KEY"

func main() {

	//Get the list of available symbols
	symbol_list := Forex.GetSymbols(api_key)
	fmt.Println(symbol_list)

	//Get quotes for specified symbols
	symbols := []string {"EURUSD", "AUDJPY", "GBPCHF"}
	quotes := Forex.GetQuotes(symbols, api_key)
	for _, quote := range quotes {
		fmt.Println(quote.Symbol)
		fmt.Println(quote.Bid)
		fmt.Println(quote.Ask)
		fmt.Println(quote.Price)
		fmt.Println(quote.Timestamp)
	}

	//Convert from one currency to another
	conversion_result := Forex.Convert("EUR", "USD", 100, api_key)
	fmt.Println(conversion_result.Value)
	fmt.Println(conversion_result.Text)
	fmt.Println(conversion_result.Timestamp)

	//Get the market status
	if Forex.MarketIsOpen(api_key) {
		fmt.Println("Market is open")
	} else {
		fmt.Println("Market is closed")
	}
}
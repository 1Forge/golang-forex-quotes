/*
 * This library is provided without warranty under the MIT license
 * Created by Jacob Davis <jacob@1forge.com>
 */

package main

import (
	Forex "github.com/1Forge/golang-forex-quotes"
	"fmt"
	"log"
)

func main() {
	//Initialize the client
	client := Forex.NewClient(apiKey)

	//Get the list of available symbols
	symbolList, e := client.GetSymbols()

	if e != nil {
		log.Fatal(e)
	}

	fmt.Println(symbolList)

	//Get quotes for specified symbols
	symbols := []string{"EURUSD", "AUDJPY", "GBPCHF"}
	quotes, e := client.GetQuotes(symbols)

	if e != nil {
		log.Fatal(e)
	}

	for _, quote := range quotes {
		fmt.Println(quote.Symbol)
		fmt.Println(quote.Bid)
		fmt.Println(quote.Ask)
		fmt.Println(quote.Price)
		fmt.Println(quote.Timestamp)
	}

	//Convert from one currency to another
	conversionResult, e := client.Convert("EUR", "USD", 100)

	if e != nil {
		log.Fatal(e)
	}

	fmt.Println(conversionResult.Value)
	fmt.Println(conversionResult.Text)
	fmt.Println(conversionResult.Timestamp)

	//Get the market status
	marketIsOpen, e := client.MarketIsOpen()

	if e != nil {
		log.Fatal(e)
	}

	if marketIsOpen {
		fmt.Println("Market is open")
	} else {
		fmt.Println("Market is closed")
	}

	//Check quota
	quota, e := client.GetQuota()

	if e != nil {
		log.Fatal(e)
	}

	fmt.Println("Quota used", quota.QuotaUsed)
	fmt.Println("Quota limit", quota.QuotaLimit)
	fmt.Println("Quota remaining", quota.QuotaRemaining)
	fmt.Println("Hours until reset", quota.HoursUntilReset)
}

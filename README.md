# golang-forex-quotes

golang-forex-quotes is a Golang library for fetching realtime forex quotes.
This library is a work-in-progress as I learn Golang.
Any contributions or issues opened are greatly appreciated.
Please see examples in the [/examples] folder.

# Table of Contents
- [Requirements](#requirements)
- [Usage](#usage)
    - [List of Symbols available](#get-the-list-of-available-symbols)
    - [Get quotes for specific symbols](#get-quotes-for-specified-symbols)
    - [Convert from one currency to another](#convert-from-one-currency-to-another)
    - [Check if the market is open](#check-if-the-market-is-open)
- [Support / Contact](#support-and-contact)
- [License / Terms](#license-and-terms)

## Requirements
* An API key which you can obtain for free at http://1forge.com/forex-data-api
* Golang

## Usage

### Get the list of available symbols:
```go
	symbol_list := Forex.GetSymbols(api_key)
	fmt.Println(symbol_list)
```

### Get quotes for specified symbols:
```go
	symbols := []string {"EURUSD", "AUDJPY", "GBPCHF"}
	quotes := Forex.GetQuotes(symbols, api_key)
	for _, quote := range quotes {
		fmt.Println(quote.Symbol)
		fmt.Println(quote.Bid)
		fmt.Println(quote.Ask)
		fmt.Println(quote.Price)
		fmt.Println(quote.Timestamp)
	}
```

### Convert from one currency to another:
```go
	conversion_result := Forex.Convert("EUR", "USD", 100, api_key)
	fmt.Println(conversion_result.Value)
	fmt.Println(conversion_result.Text)
	fmt.Println(conversion_result.Timestamp)
```

### Check if the market is open:
```go
	if Forex.MarketIsOpen(api_key) {
		fmt.Println("Market is open")
	} else {
		fmt.Println("Market is closed")
	}
```

## Support and Contact
Please contact me at contact@1forge.com if you have any questions or requests.

## License and Terms
This library is provided without warranty under the MIT license.
# golang-forex-quotes

golang-forex-quotes is a Golang library for fetching realtime forex quotes.
Any contributions or issues opened are greatly appreciated.
Please see examples in the [/examples](https://github.com/1Forge/golang-forex-quotes/tree/master/examples) folder.

# Table of Contents
- [Requirements](#requirements)
- [Known Issues](#known-issues)
- [Installation](#installation)
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

## Known Issues
Please see the list of known issues here: [Issues](https://github.com/1Forge/golang-forex-quotes/issues)

## Installation

`go get github.com/1Forge/golang-forex-quotes`

## Usage

### Initialize the client
```go
import (
	Forex "github.com/1Forge/golang-forex-quotes"
	"fmt"
)

client := Forex.NewClient("YOUR_API_KEY")
```

### Get the list of available symbols:
```go
symbolList, e := client.GetSymbols()

if e != nil {
    log.Fatal(e)
}

fmt.Println(symbolList)

```

### Get quotes for specified symbols:
```go
symbols := []string {"EURUSD", "AUDJPY", "GBPCHF"}
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
```

### Convert from one currency to another:
```go
conversionResult, e := client.Convert("EUR", "USD", 100)

if e != nil {
    log.Fatal(e)
}

fmt.Println(conversionResult.Value)
fmt.Println(conversionResult.Text)
fmt.Println(conversionResult.Timestamp)
```

### Check if the market is open:
```go
marketIsOpen, e := client.MarketIsOpen()

if e != nil {
    log.Fatal(e)
}

if marketIsOpen {
    fmt.Println("Market is open")
} else {
    fmt.Println("Market is closed")
}
```

### Quota used
```go
quota, e := client.GetQuota()

if e != nil {
    log.Fatal(e)
}

fmt.Println("Quota used", quota.QuotaUsed)
fmt.Println("Quota limit", quota.QuotaLimit)
fmt.Println("Quota remaining", quota.QuotaRemaining)
fmt.Println("Hours until reset", quota.HoursUntilReset)
```

## Support and Contact
Please contact me at contact@1forge.com if you have any questions or requests.

## License and Terms
This library is provided without warranty under the MIT license.

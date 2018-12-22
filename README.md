# golang-forex-quotes

golang-forex-quotes is a Golang library for fetching realtime forex quotes.
Any contributions or issues opened are greatly appreciated.
Please see examples in the [/examples](https://github.com/1Forge/golang-forex-quotes/tree/master/examples) folder.

# Table of Contents
- [Requirements](#requirements)
- [Known Issues](#known-issues)
- [Installation](#installation)
- [Usage](#usage)
- [Support / Contact](#support-and-contact)
- [License / Terms](#license-and-terms)

## Requirements
* An API key which you can obtain at http://1forge.com/forex-data-api

## Known Issues
Please see the list of known issues here: [Issues](https://github.com/1Forge/golang-forex-quotes/issues)

## Installation

`go get github.com/1Forge/golang-forex-quotes`

## Usage

### Initialize the client
```go
import (
	Forex "github.com/1Forge/golang-forex-quotes"
)
```

### WebSocket API
```go
func main() {
    client := Forex.CreateClient("YOUR_API_KEY")

	symbols := []string{"BTCJPY", "AUDJPY", "GBPCHF"}

	// Specify the update handler
	client.OnUpdate(func(q Forge.Quote) {
		fmt.Println(q)
	})

	// Specify the message handler
	client.OnMessage(func(m string) {
		fmt.Println(m)
	})

	// Specify the disconnection handler
	client.OnDisconnection(func() {
		fmt.Println("Disconnected")
	})

	// Specify the login success handler
	client.OnLoginSuccess(func() {
		fmt.Println("Successfully logged in")

		// Subscribe to some symbols
		client.SubscribeTo(symbols)

		// Subscribe to all symbols
		client.SubscribeToAll()
	})

	// Specify the connection handler
	client.OnConnection(func() {
		fmt.Println("Connected")
	})

	// Connect to the socket server
	client.Connect()

	// Wait 25 seconds
	time.Sleep(25 * time.Second)

	// Unsubscribe from some symbols
	client.UnsubscribeFrom(symbols)

	// Unsubscribe from all symbols
	client.UnsubscribeFromAll()

	// Disconnect
    client.Disconnect()
}
```

### RESTful API

```go
func main() {
    client := Forex.CreateClient("YOUR_API_KEY")

    // Get the list of symbols
	symbols, e := client.GetSymbols()

	if e != nil {
		log.Fatal(e)
	}

	// Gets quotes
	quotes, e := client.GetQuotes(symbols)

	if e != nil {
		log.Fatal(e)
	}

	fmt.Println(quotes)

	// Convert currencies
	conversion, e := client.Convert("EUR", "USD", 100)
	if e != nil {
		log.Fatal(e)
	}

	fmt.Println(conversion.Value)
	fmt.Println(conversion.Text)
	fmt.Println(conversion.Timestamp)

	// Get the market status
	marketStatus, e := client.GetMarketStatus()

	if e != nil {
		log.Fatal(e)
	}

	fmt.Println("Is the market open?", marketStatus.MarketIsOpen)

	// Get current quota
	quota, e := client.GetQuota()

	if e != nil {
		log.Fatal(e)
	}

	fmt.Println("Quota used", quota.QuotaUsed)
	fmt.Println("Quota limit", quota.QuotaLimit)
	fmt.Println("Quota remaining", quota.QuotaRemaining)
    fmt.Println("Hours until reset", quota.HoursUntilReset)
}
```

## Support and Contact
Please contact me at contact@1forge.com if you have any questions or requests.

## License and Terms
This library is provided without warranty under the MIT license.

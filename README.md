## Currency API Go SDK

[![Release](https://img.shields.io/github/release/moatsystems/currensees.svg)](https://github.com/moatsystems/currensees/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/moatsystems/currensees)](https://goreportcard.com/report/github.com/moatsystems/currensees)
[![Go Reference](https://pkg.go.dev/badge/github.com/moatsystems/currensees.svg)](https://pkg.go.dev/github.com/moatsystems/currensees)
[![License](https://img.shields.io/github/license/moatsystems/currensees)](/LICENSE)

The SDK provides convenient access to the [Currency API](https://moatsystems.com/currency-api/) for applications written in the [Go](https://go.dev/) Programming Language.

### Usage Example

```go
// main.go
package main

import (
	"fmt"
	"currensees"
)

func main() {
	// Authentication
	username := "your_username"
	password := "your_password"

	authResponse := currensees.Authenticate(username, password)
	fmt.Println("Authentication Response:")
	fmt.Println(authResponse)

	// Currency Conversion
	date := "2023_04_02"
	baseCurrency := "GBP"
	targetCurrency := "CAD"
	amount := "500"

	convertResponse := currensees.Convert(date, baseCurrency, targetCurrency, amount)
	fmt.Println("Currency Conversion Response:")
	fmt.Println(convertResponse)

	// Convert All Currencies
	baseCurrencyAll := "GBP"
	amountAll := "120"
	dateAll := "2023_04_02"

	convertAllResponse := currensees.ConvertAll(baseCurrencyAll, amountAll, dateAll)
	fmt.Println("Currency Conversion for All Currencies Response:")
	fmt.Println(convertAllResponse)
}
```

### Setting up Currency API Account

Subscribe [here](https://moatsystems.com/currency-api/) for a user account.

### Using the Currency API

You can read the [API documentation](https://docs.currensees.com/) to understand what’s possible with the Currency API. If you need further assistance, don’t hesitate to [contact us](https://moatsystems.com/contact/).

### License

This project is licensed under the [BSD 3-Clause License](./LICENSE).

### Copyright

(c) 2023 [Moat Systems Limited](https://moatsystems.com).
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

	day := "02"
	month := "04"
	year := "2023"
	date := "2023_04_02"
	from_date := "2023_04_03"
	to_date := "2023_04_07"

	authResponse := currensees.Authenticate(username, password)
	fmt.Println("Authentication Response:")
	fmt.Println(authResponse)

	// Currency Conversion
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

	// Retrieve All Currencies
	response := currensees.GetCurrencies(username, day, month, year)
	fmt.Println(response)

	// Retrieve a currency by UUID
	uuid := "your_uuid"
	response := currensees.GetCurrencyByUUID(uuid, username, day, month, year)
	fmt.Println(response)

	// Retrieve All Historical Currencies
	response := currensees.GetHistorical(username, date, day, month, year)
	fmt.Println(response)

	// Retrieve a historical currency by UUID
	uuid := "your_uuid"
	response := currensees.GetHistoricalByUUID(uuid, username, day, month, year, date)
	fmt.Println(response)

	// Retrieve the daily average for a particular date
	response := currensees.DailyAverage(date)
	fmt.Println(response)

	// Retrieve the weekly average for a date range
	response := currensees.WeeklyAverage(from_date, to_date)
	fmt.Println(response)

	// Get all margins and spreads for a specific date
	allMarginsSpreads := currensees.GetMarginsSpreads(username, day, month, year)
	fmt.Println("All Margins and Spreads:", allMarginsSpreads)

	// Get margins and spreads by UUID
	uuid := "your_uuid"
	marginsSpreadsByUUID := currensees.GetMarginsSpreadsByUUID(uuid, username, day, month, year)
	fmt.Println("Margins and Spreads by UUID:", marginsSpreadsByUUID)
}
```

### Setting up Currency API Account

Subscribe [here](https://moatsystems.com/currency-api/) for a user account.

### Using the Currency API

You can read the [API documentation](https://docs.currensees.com/) to understand what's possible with the Currency API. If you need further assistance, don't hesitate to [contact us](https://moatsystems.com/contact/).

### License

This project is licensed under the [BSD 3-Clause License](./LICENSE).

### Copyright

(c) 2023 [Moat Systems Limited](https://moatsystems.com).
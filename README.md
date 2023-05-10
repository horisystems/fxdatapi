## Currency API Go SDK

[![Release](https://img.shields.io/github/release/moatsystems/fxdatapi.svg)](https://github.com/moatsystems/fxdatapi/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/moatsystems/fxdatapi)](https://goreportcard.com/report/github.com/moatsystems/fxdatapi)
[![Go Reference](https://pkg.go.dev/badge/github.com/moatsystems/fxdatapi.svg)](https://pkg.go.dev/github.com/moatsystems/fxdatapi)
[![License](https://img.shields.io/github/license/moatsystems/fxdatapi)](/LICENSE)

The SDK provides convenient access to the [Currency API](https://moatsystems.com/currency-api/) for applications written in the [Go](https://go.dev/) Programming Language.

### Usage Example

```go
// main.go
package main

import (
	"fmt"
	"fxdatapi"
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

	authResponse := fxdatapi.Authenticate(username, password)
	fmt.Println("Authentication Response:")
	fmt.Println(authResponse)

	// Currency Conversion
	baseCurrency := "GBP"
	targetCurrency := "CAD"
	amount := "500"

	convertResponse := fxdatapi.Convert(username, date, baseCurrency, targetCurrency, amount)
	fmt.Println("Currency Conversion Response:")
	fmt.Println(convertResponse)

	// Convert All Currencies
	baseCurrencyAll := "GBP"
	amountAll := "120"
	dateAll := "2023_04_02"

	convertAllResponse := fxdatapi.ConvertAll(username, baseCurrencyAll, amountAll, dateAll)
	fmt.Println("Currency Conversion for All Currencies Response:")
	fmt.Println(convertAllResponse)

	// Retrieve All Currencies
	response := fxdatapi.GetCurrencies(username, day, month, year)
	fmt.Println(response)

	// Retrieve a currency by UUID
	uuid := "currencies_uuid"
	response := fxdatapi.GetCurrencyByUUID(uuid, username, day, month, year)
	fmt.Println(response)

	// Retrieve All Historical Currencies
	response := fxdatapi.GetHistorical(username, date, day, month, year)
	fmt.Println(response)

	// Retrieve a historical currency by UUID
	uuid := "historical_uuid"
	response := fxdatapi.GetHistoricalByUUID(uuid, username, day, month, year, date)
	fmt.Println(response)

	// Retrieve the daily average for a particular date
	response := fxdatapi.DailyAverage(date)
	fmt.Println(response)

	// Retrieve the weekly average for a date range
	response := fxdatapi.WeeklyAverage(from_date, to_date)
	fmt.Println(response)

	// Retrieve the monthly average for a specific month and year
	response := fxdatapi.MonthlyAverage(year, month)
	fmt.Println(response)

	// Get all margins and spreads for a specific date
	allMarginsSpreads := fxdatapi.GetMarginsSpreads(username, day, month, year)
	fmt.Println("All Margins and Spreads:", allMarginsSpreads)

	// Get margins and spreads by UUID
	uuid := "margins_spreads_uuid"
	marginsSpreadsByUUID := fxdatapi.GetMarginsSpreadsByUUID(uuid, username, day, month, year)
	fmt.Println("Margins and Spreads by UUID:", marginsSpreadsByUUID)

	// Get all performances
	allPerformances := fxdatapi.GetPerformances(username)
	fmt.Println("All Performances:")
	fmt.Println(allPerformances)

	// Get performance by UUID
	performanceUUID := "performance_uuid"
	performanceByUUID := fxdatapi.GetPerformanceByUUID(performanceUUID, username)
	fmt.Println("Performance by UUID:")
	fmt.Println(performanceByUUID)

	// Get all signals
	allSignals := fxdatapi.GetSignals(username)
	fmt.Println("All Signals:")
	fmt.Println(allSignals)

	// Get signal by UUID
	signalUUID := "signal_uuid"
	signalByUUID := fxdatapi.GetSignalByUUID(signalUUID, username)
	fmt.Println("Signal by UUID:")
	fmt.Println(signalByUUID)
}
```

### Setting up Currency API Account

Subscribe [here](https://moatsystems.com/currency-api/) for a user account.

### Using the Currency API

You can read the [API documentation](https://docs.fxdatapi.com/) to understand what's possible with the Currency API. If you need further assistance, don't hesitate to [contact us](https://moatsystems.com/contact/).

### License

This project is licensed under the [BSD 3-Clause License](./LICENSE).

### Copyright

(c) 2023 [Moat Systems Limited](https://moatsystems.com).
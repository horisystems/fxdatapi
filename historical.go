// historical.go
package currensees

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetHistorical(username, date, day, month, year string) string {
	url := fmt.Sprintf("https://currensees.com/v1/historical?username=%s&date=%s&day=%s&month=%s&year=%s", username, date, day, month, year)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(body)
}

func GetHistoricalByUUID(uuid, username, day, month, year, date string) string {
	url := fmt.Sprintf("https://currensees.com/v1/historical/%s?username=%s&day=%s&month=%s&year=%s&date_string=%s", uuid, username, day, month, year, date)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(body)
}

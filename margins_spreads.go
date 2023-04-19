// margins_spreads.go
package currensees

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetMarginsSpreads(username, day, month, year string) string {
	url := fmt.Sprintf("https://currensees.com/v1/margins_spreads?username=%s&day=%s&month=%s&year=%s", username, day, month, year)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Add("Content-Type", "application/json")
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

func GetMarginsSpreadsByUUID(uuid, username, day, month, year string) string {
	url := fmt.Sprintf("https://currensees.com/v1/margins_spreads/%s?username=%s&day=%s&month=%s&year=%s", uuid, username, day, month, year)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Add("Content-Type", "application/json")
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

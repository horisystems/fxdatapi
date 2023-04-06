// convert_all.go
package currensees

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func ConvertAll(baseCurrency, amount, date string) string {
	url := "https://currensees.com/v1/convert_all"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`{
  "base_currency": "%s",
  "amount": %s,
  "date": "%s"
}`, baseCurrency, amount, date))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

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

package httpjson

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetJSON Returns a JSON Struct from a url using GET
func GetJSON(url string, obj interface{}) error {
	client := new(http.Client)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "quantocustaobitcoin.com.br Block Checker Bot")

	res, err := client.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("Status not 200\n%s", body)
	}

	var quote byte = '"'
	_, isString := obj.(*string)
	if isString && body[0] != quote {
		body = append([]byte{quote}, body...)
		body = append(body, quote)
	}

	return json.Unmarshal(body, &obj)
}

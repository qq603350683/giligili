package util

import (
	"io/ioutil"
	"log"
	"net/http"
)

func CURL(method string, url string, data interface{}, headers interface{}) string {
	var err error
	var resp *http.Response

	switch method {
	case "GET":
		resp, err = http.Get(url)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("CURL fail, StatusCode is: %d",  resp.StatusCode)

		return ""
	}

	if err != nil {
		return ""
	}

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	str := string(all)

	return str
}

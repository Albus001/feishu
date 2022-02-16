package common

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func SendPostRequest(url string, headers map[string]string, body []byte) (resp []byte, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(string(body)))
	if err != nil {
		return resp, err
	}
	for key, header := range headers {
		req.Header.Set(key, header)
	}
	res, err := client.Do(req)
	defer res.Body.Close()
	resp, err = ioutil.ReadAll(res.Body)
	return resp, nil
}
func SendGetRequest(url string, headers map[string]string) (resp []byte, err error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return resp, err
	}
	for key, header := range headers {
		req.Header.Set(key, header)
	}
	res, err := client.Do(req)
	defer res.Body.Close()
	resp, err = ioutil.ReadAll(res.Body)
	return resp, nil
}

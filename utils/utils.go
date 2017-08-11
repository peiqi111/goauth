package utils

import (
	crand "crypto/rand"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func GetState(length int) string {
	randomBytes := make([]byte, length)
	_, err := crand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", randomBytes)
}

func GetUrlData(inputUrl string, method string) ([]byte, error) {
	parseUrl, _ := url.Parse(inputUrl)
	req := http.Request{
		URL:        parseUrl,
		Method:     method,
		ProtoMajor: 1,
		ProtoMinor: 1,
		Close:      true,
	}

	// Add the header params
	header := make(http.Header)
	header.Set("Accept", "application/json")
	header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header = header
	resp, err := http.DefaultClient.Do(&req)
	if err != nil {
		return []byte(""), err
	}
	// Get the response body
	raw, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return []byte(""), err
	}
	return raw, nil
}

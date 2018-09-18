package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func HTTPGetData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("StatusCode %v", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}

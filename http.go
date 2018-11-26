package utils

import (
	"fmt"
	"io"
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

func HttpDo(client *http.Client, request *http.Request, filters ...HTTPResponseFilter) (io.ReadCloser, error) {
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	for _, filter := range filters {
		err = filter(resp)
		if err != nil {
			break
		}
	}
	if err != nil {
		resp.Body.Close()
		return nil, err
	}
	return resp.Body, nil
}

func HttpDoAndClose(client *http.Client, request *http.Request, filters ...HTTPResponseFilter) ([]byte, error) {
	body, err := HttpDo(client, request, filters...)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(body)
}

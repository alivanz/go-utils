package utils

import (
	"errors"
	"mime"
	"net/http"
	"strings"
)

var (
	Non200          = errors.New("return code !200")
	ContentMismatch = errors.New("Content-Type mismatch")
)

const (
	octetStream = "application/octet-stream"
)

type HTTPResponseFilter func(*http.Response) error

func MustHTTP200(resp *http.Response) error {
	if resp.StatusCode != 200 {
		return Non200
	}
	return nil
}
func MustHTTPContentType(s string) HTTPResponseFilter {
	return func(r *http.Response) error {
		contentType := r.Header.Get("Content-type")
		if s == octetStream && contentType == "" {
			return nil
		}
		for _, v := range strings.Split(contentType, ",") {
			t, _, err := mime.ParseMediaType(v)
			if err != nil {
				break
			}
			if t == s {
				return nil
			}
		}
		return ContentMismatch
	}
}
func MustHTTPOctetStream() HTTPResponseFilter {
	return MustHTTPContentType(octetStream)
}

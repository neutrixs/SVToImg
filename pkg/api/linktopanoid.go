package api

import (
	"errors"
	"net/http"
	"strings"
)

var InvalidSVURLError error = errors.New("Invalid Streetview URL!")
var InvalidURLError error = errors.New("Invalid URL!")

func ShortlinkToPanoid(url string) (string, error){
	hasPrefix := strings.HasPrefix(url, "https://")

	if !hasPrefix {
		return "", InvalidURLError
	}

	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	finalURL := resp.Request.URL.String()

	var splitted []string

	splitted = strings.Split(finalURL, "/data=!3m6!1e1!3m4!1s")

	if len(splitted) < 2 {
		return "", InvalidSVURLError
	}

	panoid := strings.Split(splitted[1], "!")[0]

	return panoid, nil
}
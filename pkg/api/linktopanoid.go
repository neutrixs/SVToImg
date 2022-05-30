package api

import (
	"errors"
	"net/http"
	"regexp"
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

	finalURLHasPrefix := strings.HasPrefix(finalURL, "https://www.google.com/maps/")

	if !finalURLHasPrefix {
		return "", InvalidSVURLError
	}

	var splitted []string

	// wtf?
	// when you first click the blue line, it's 3m7 and 3m5, but once you move, it becomes 3m6 and 3m4
	regex := regexp.MustCompile(`\/data=!3m[0-9]!1e1!3m[0-9]!1s`)

	splitted = regex.Split(finalURL, -1)

	if len(splitted) < 2 {
		return "", InvalidSVURLError
	}

	panoid := strings.Split(splitted[1], "!")[0]

	return panoid, nil
}
package api

import (
	"errors"
	"fmt"
	"net/http"
)

var PanoidDoesNotExistError = errors.New("panoid does not exist!")

func getURL(panoid string, x, y, zoom int) string {
	const URL string = "https://streetviewpixels-pa.googleapis.com/v1/tile?cb_client=maps_sv.tactile&panoid=%v&x=%d&y=%d&zoom=%d&nbt=1&fover=2"
	
	return fmt.Sprintf(URL, panoid, x, y, zoom)
}

// returns 0, error if it returned an error.
// for now, both gen 2 and gen 3 will return 2, nil (because they have the same x length)
func GetGeneration(panoid string) (int, error) {
	var resp *http.Response
	var err error

	// detect if panoid actually exists
	resp, err = http.Get(getURL(panoid, 0, 0, 0))

	if err != nil {
		return 0, err
	}

	if resp.StatusCode == 400 {
		return 0, PanoidDoesNotExistError
	}

	// detects gen 1
	resp, err = http.Get(getURL(panoid, 0, 0, 4))

	if err != nil {
		return 0, err
	}

	if resp.StatusCode == 400 {
		return 1, nil
	}

	// detects gen 2 OR 3 (THEY HAVE THE EXACT SAME RESOLUTION)
	resp, err = http.Get(getURL(panoid, 26, 0, 5))

	if err != nil {
		return 0, err
	}

	if resp.StatusCode == 400 {
		return 2, nil
	}

	return 4, nil
}
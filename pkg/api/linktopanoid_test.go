package api

import "testing"

func TestShortlinkToPanoid(t *testing.T) {
	var url string
	var wantPanoid string
	var wantErr error
	var gotPanoid string
	var gotErr error

	url = "https://goo.gl/maps/v3V1nGaqSHuxHpF17"
	wantPanoid = "jpmaYUuc5-QXvTD29Kz54A"
	wantErr = nil

	gotPanoid, gotErr = ShortlinkToPanoid(url)

	if gotPanoid != wantPanoid || gotErr != wantErr {
		t.Fatalf(`ShortLinkToPanoid("%v") = %v, %v, want match for %v, %v`, url, gotPanoid, gotErr, wantPanoid, wantErr)
	}

	url = "random"
	wantPanoid = ""
	wantErr = InvalidURLError

	gotPanoid, gotErr = ShortlinkToPanoid(url)

	if gotPanoid != wantPanoid || gotErr != wantErr {
		t.Fatalf(`ShortLinkToPanoid("%v") = %v, %v, want match for %v, %v`, url, gotPanoid, gotErr, wantPanoid, wantErr)
	}

	url = "https://google.com/maps"
	wantPanoid = ""
	wantErr = InvalidSVURLError

	gotPanoid, gotErr = ShortlinkToPanoid(url)

	if gotPanoid != wantPanoid || gotErr != wantErr {
		t.Fatalf(`ShortLinkToPanoid("%v") = %v, %v, want match for %v, %v`, url, gotPanoid, gotErr, wantPanoid, wantErr)
	}
}
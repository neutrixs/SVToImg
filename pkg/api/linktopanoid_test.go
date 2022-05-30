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

	url = "https://www.google.com/maps/@-6.256829,106.9520042,3a,75y,277.25h,77.76t/data=!3m6!1e1!3m4!1scP2u46sSj0aUXmyfFCYRWA!2e0!7i16384!8i8192"
	wantPanoid = "cP2u46sSj0aUXmyfFCYRWA"
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

	url = "https://example.com/data=!3m6!1e1!3m4!1stest"
	wantPanoid = ""
	wantErr = InvalidSVURLError

	gotPanoid, gotErr = ShortlinkToPanoid(url)

	if gotPanoid != wantPanoid || gotErr != wantErr {
		t.Fatalf(`ShortLinkToPanoid("%v") = %v, %v, want match for %v, %v`, url, gotPanoid, gotErr, wantPanoid, wantErr)
	}
}
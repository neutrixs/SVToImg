package photometa

import (
	"testing"
)

func TestGetURL(t *testing.T) {
	const panoid = "V0XUIqgxMQYLmklWop_B6Q"
	const wantURL = "https://www.google.com/maps/photometa/v1?authuser=0&hl=id&gl=id&pb=!1m4!1smaps_sv.tactile!11m2!2m1!1b1!2m2!1sid!2sid!3m3!1m2!1e2!2sV0XUIqgxMQYLmklWop_B6Q!4m57!1e1!1e2!1e3!1e4!1e5!1e6!1e8!1e12!2m1!1e1!4m1!1i48!5m1!1e1!5m1!1e2!6m1!1e1!6m1!1e2!9m36!1m3!1e2!2b1!3e2!1m3!1e2!2b0!3e3!1m3!1e3!2b1!3e2!1m3!1e3!2b0!3e3!1m3!1e8!2b0!3e3!1m3!1e1!2b0!3e3!1m3!1e4!2b0!3e3!1m3!1e10!2b1!3e2!1m3!1e10!2b0!3e3"

	gotURL := GetURL(panoid)

	if wantURL != gotURL {
		t.Fatalf(`GetURL("%v") = %v, want match for %v`, panoid, gotURL, wantURL)
	}
}
package api

import "testing"

func TestShortlinkToPanoid(t *testing.T) {
	test := func(url string, wantPanoid string, wantErr error) {
		gotPanoid, gotErr := ShortlinkToPanoid(url)

		if gotPanoid != wantPanoid || gotErr != wantErr {
			t.Fatalf(`ShortLinkToPanoid("%v") = %v, %v, want match for %v, %v`, url, gotPanoid, gotErr, wantPanoid, wantErr)
		}
	}

	test("https://goo.gl/maps/v3V1nGaqSHuxHpF17", "jpmaYUuc5-QXvTD29Kz54A", nil)
	test("https://goo.gl/maps/mQac1a9wFkt3QRHo7", "JSTF9BICkKl76cYXG5pQvQ", nil)
	test("https://goo.gl/maps/Fgcz4tr1kPf86Uko6", "Zfh23J2FjPaNcohLhAaWrQ", nil)
	test("random", "", InvalidURLError)
	test("https://google.com/maps", "", InvalidSVURLError)
	test("https://example.com/data=!3m6!1e1!3m4!1stest", "", InvalidSVURLError)

	test(
		"https://www.google.com/maps/@-6.256829,106.9520042,3a,75y,277.25h,77.76t/data=!3m6!1e1!3m4!1scP2u46sSj0aUXmyfFCYRWA!2e0!7i16384!8i8192",
		"cP2u46sSj0aUXmyfFCYRWA",
		nil,
	)
}
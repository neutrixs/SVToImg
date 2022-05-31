package api

import (
	"testing"

	"github.com/corona10/goimagehash"
)

func TestGetImages(t *testing.T) {
	testExpectErr := func(panoid string) {
		wantImg := EmptyImage
		wantErr := PanoidDoesNotExistError

		gotImg, gotErr := GetImages(panoid)

		if wantImg != gotImg || wantErr != gotErr {
			t.Fatalf(`GetImages("%v") = %v, %v, want match for %v, %v`, panoid, gotImg, gotErr, wantImg, wantErr)
		}
	}

	testHash := func(panoid string, wantHash uint64, wantErr error) {
		img, gotErr := GetImages(panoid)

		hashResult, _ := goimagehash.AverageHash(img)
		gotHash := hashResult.GetHash()

		if gotHash != wantHash {
			t.Fatalf(`GetImages("%v") has a hash of %d, want hash of %d`, panoid, gotHash, wantHash)
		}

		if gotErr != wantErr {
			t.Fatalf(`GetImages("%v") has an err of %v, want err of %v`, panoid, gotErr, wantErr)
		}
	}

	testExpectErr("random")

	testHash("cqlr_mv0kHsWFT0EoOj60Q", uint64(18446743030032498688), nil)
	testHash("4vsCZ9XiObL2UXMbNzUXOg", uint64(18446744052234715136), nil)
	testHash("YNuEAlZ0fJ0cQrfkvJtpHA", uint64(18446554768730947584), nil)
}
package api

import (
	"errors"
	"fmt"
	"testing"
)

func TestGetGeneration(t *testing.T) {
	var panoid string
	var got int
	var gotErr error
	var want int
	var wantErr error

	panoid = "invalidID"
	want, wantErr = 0, errors.New("panoid does not exist!")
	got, gotErr = GetGeneration(panoid)

	fmt.Println(gotErr)

	if got != want || gotErr.Error() != wantErr.Error() {
		t.Fatalf(`getGeneration("%v") = %v, %v, want match for %v, %v`, panoid, got, gotErr, 0, wantErr)
	}

	// gen 1 panoid
	panoid = "_drdwKa3KNQxUR6O7ZUCmQ"
	got, gotErr = GetGeneration(panoid)
	want, wantErr = 1, nil

	if got != want || gotErr != wantErr {
		t.Fatalf(`getGeneration("%v") = %v, %v, want match for %v, %v`, panoid, got, gotErr, want, wantErr)
	}

	// gen 2 panoid
	panoid = "hLPja2vrvr5g-6gi3f8-FQ"
	got, gotErr = GetGeneration(panoid)
	want, wantErr = 2, nil

	if got != want || gotErr != wantErr {
		t.Fatalf(`getGeneration("%v") = %v, %v, want match for %v, %v`, panoid, got, gotErr, want, wantErr)
	}

	// gen 3 panoid
	// temporarily, this should return 2 (because they have the same x length)
	panoid = "ga9K9_YYEwxb4p5ApmQSkA"
	got, gotErr = GetGeneration(panoid)
	want, wantErr = 2, nil

	if got != want || gotErr != wantErr {
		t.Fatalf(`getGeneration("%v") = %v, %v, want match for %v, %v`, panoid, got, gotErr, want, wantErr)
	}

	// gen 4 panoid
	panoid = "9ApA6xOofv6Dq-BrQrdGbQ"
	got, gotErr = GetGeneration(panoid)
	want, wantErr = 4, nil

	if got != want || gotErr != wantErr {
		t.Fatalf(`getGeneration("%v") = %v, %v, want match for %v, %v`, panoid, got, gotErr, want, wantErr)
	}
}
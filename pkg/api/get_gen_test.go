package api

import (
	"errors"
	"testing"
)

func TestGetGeneration(t *testing.T) {
	var got int
	var gotErr error
	var want int
	var wantErr error

	got, gotErr = GetGeneration("invalidID")
	want, wantErr = 0, errors.New("panoid does not exist!")

	if got != want || gotErr == wantErr {
		t.Fatalf(`getGeneration("invalidID") = %v, %v, want match for %v, %v`, got, gotErr, 0, wantErr)
	}

	// gen 1 panoid
	got, gotErr = GetGeneration("_drdwKa3KNQxUR6O7ZUCmQ")
	want, wantErr = 1, nil

	if got != want || gotErr != wantErr {
		t.Fatalf(`getGeneration("_drdwKa3KNQxUR6O7ZUCmQ") = %v, %v, want match for %v, %v`, got, gotErr, want, wantErr)
	}

	// gen 2 panoid
	got, gotErr = GetGeneration("hLPja2vrvr5g-6gi3f8-FQ")
	want, wantErr = 2, nil

	if got != want || gotErr != wantErr {
		t.Fatalf(`getGeneration("hLPja2vrvr5g-6gi3f8-FQ") = %v, %v, want match for %v, %v`, got, gotErr, want, wantErr)
	}

	// gen 3 panoid
	// temporarily, this should return 2 (because they have the same x length)
	got, gotErr = GetGeneration("ga9K9_YYEwxb4p5ApmQSkA")
	want, wantErr = 2, nil

	if got != want || gotErr != wantErr {
		t.Fatalf(`getGeneration("ga9K9_YYEwxb4p5ApmQSkA") = %v, %v, want match for %v, %v`, got, gotErr, want, wantErr)
	}

	// gen 4 panoid
	got, gotErr = GetGeneration("9ApA6xOofv6Dq-BrQrdGbQ")
	want, wantErr = 4, nil

	if got != want || gotErr != wantErr {
		t.Fatalf(`getGeneration("9ApA6xOofv6Dq-BrQrdGbQ") = %v, %v, want match for %v, %v`, got, gotErr, want, wantErr)
	}
}
package api

import (
	"testing"
)

func TestGetGeneration(t *testing.T) {
	var got int
	var gotErr error

	got, gotErr = GetGeneration("invalidID")

	if got != 0 || gotErr == nil {
		t.Fatalf(`getGeneration("invalidID") = %v, %v, want match for %v, %v`, got, gotErr, 0, "panoid does not exist!")
	}

	// gen 1 panoid
	got, gotErr = GetGeneration("_drdwKa3KNQxUR6O7ZUCmQ")

	if got != 1 || gotErr != nil {
		t.Fatalf(`getGeneration("_drdwKa3KNQxUR6O7ZUCmQ") = %v, %v, want match for %v, %v`, got, gotErr, 1, nil)
	}

	// gen 2 panoid
	got, gotErr = GetGeneration("hLPja2vrvr5g-6gi3f8-FQ")

	if got != 2 || gotErr != nil {
		t.Fatalf(`getGeneration("hLPja2vrvr5g-6gi3f8-FQ") = %v, %v, want match for %v, %v`, got, gotErr, 2, nil)
	}

	// gen 3 panoid
	// temporarily, this should return 2 (because they have the same x length)
	got, gotErr = GetGeneration("ga9K9_YYEwxb4p5ApmQSkA")

	if got != 2 || gotErr != nil {
		t.Fatalf(`getGeneration("ga9K9_YYEwxb4p5ApmQSkA") = %v, %v, want match for %v, %v`, got, gotErr, 2, nil)
	}

	// gen 4 panoid
	got, gotErr = GetGeneration("9ApA6xOofv6Dq-BrQrdGbQ")

	if got != 4 || gotErr != nil {
		t.Fatalf(`getGeneration("9ApA6xOofv6Dq-BrQrdGbQ") = %v, %v, want match for %v, %v`, got, gotErr, 4, nil)
	}
}
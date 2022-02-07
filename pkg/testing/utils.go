package apptest

import (
	"testing"
)

// AssertNoError fails if err is not nil
func AssertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("expected no error but got %v", err)
	}
}

// AssertError fails if err is not instance of error
func AssertError(t *testing.T, got error, want error) {
	t.Helper()
	if got != want {
		t.Fatalf("expected to have ErrEmailInvalid error but didn't got one %v", got)
	}
}

// AssertEmptyString fails if string is not empty
func AssertEmptyString(t *testing.T, str string) {
	t.Helper()
	if str != "" {
		t.Fatalf("expected string to be empty but got %v", str)
	}
}

// AssertNil fails if got is not nil
func AssertNil(t *testing.T, got interface{}) {
	t.Helper()
	if got != nil {
		t.Fatalf("expected nil but got %v", got)
	}
}

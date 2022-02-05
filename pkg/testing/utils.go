package apptest

import "testing"

// AssertNoError throws error if err is not nil
func AssertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("expected no error but got %v", err)
	}
}

// AssertError throws error if err is not instance of error
func AssertError(t *testing.T, got error, want error) {
	t.Helper()
	if got != want {
		t.Fatalf("expected to have ErrEmailInvalid error but didn't got one %v", got)
	}
}

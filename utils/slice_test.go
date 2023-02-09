package utils

import (
	"testing"
)

func TestCheckInStringArray(t *testing.T) {
	inStringArray := CheckInStringArray("hello", &[]string{"hello", "world"})
	if !inStringArray {
		t.Fatal("failed to find 'hello' in []string{'hello', 'world')")
	}

	inStringArray = CheckInStringArray("not-exist", &[]string{"hello", "world"})
	if inStringArray {
		t.Fatal("'not-exist' in []string{'hello', 'world')")
	}
}

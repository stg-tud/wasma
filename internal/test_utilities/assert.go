package test_utilities

import (
	"fmt"
	"testing"
)

func AssertEqualM(t *testing.T, a interface{}, b interface{}, message string) {
	if a != b {
		t.Fatal(fmt.Sprintf("%v %v != %v", message, a, b))
	}
}

func AssertEqual(t *testing.T, a interface{}, b interface{}) {
	AssertEqualM(t, a, b, "")
}

func AssertTrueM(t *testing.T, value bool, message string) {
	if !value {
		t.Fatal(fmt.Sprintf("%v %v != %v", message, value, true))
	}
}

func AssertTrue(t *testing.T, value bool) {
	AssertTrueM(t, value, "")
}

func AssertFalseM(t *testing.T, value bool, message string) {
	if value {
		t.Fatal(fmt.Sprintf("%v %v != %v", message, value, false))
	}
}

func AssertFalse(t *testing.T, value bool) {
	AssertFalseM(t, value, "")
}

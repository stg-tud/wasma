package test_utilities

import "testing"

func CompareByte(expected byte, actual byte, t *testing.T) {
	if expected != actual {
		t.Errorf("expected value: %v != actual value: %v", expected, actual)
	}
}

func CompareBytes(expected []byte, actual []byte, t *testing.T) {
	if len(expected) != len(actual) {
		t.Errorf("expected value: %v != actual value: %v", expected, actual)
	}

	for index, _ := range expected {
		if expected[index] != actual[index] {
			t.Errorf("expected value: %v != actual value: %v", expected[index], actual[index])
		}
	}
}

func CompareString(expected string, actual string, t *testing.T) {
	if expected != actual {
		t.Errorf("expected value: %v != actual value: %v", expected, actual)
	}
}

func CompareStrings(expected []string, actual []string, t *testing.T) {
	if len(expected) != len(actual) {
		t.Errorf("expected value: %v != actual value: %v", expected, actual)
	}

	for i, _ := range expected {
		if expected[i] != actual[i] {
			t.Errorf("expected value: %v != actual value: %v", expected, actual)
		}
	}
}

func CompareInt(expected int, actual int, t *testing.T) {
	if expected != actual {
		t.Errorf("expected value: %v != actual value: %v", expected, actual)
	}
}

func CompareUInt32(expected uint32, actual uint32, t *testing.T) {
	if expected != actual {
		t.Errorf("expected value: %v != actual value: %v", expected, actual)
	}
}

func CompareUInt32s(expected []uint32, actual []uint32, t *testing.T) {
	if len(expected) != len(actual) {
		t.Errorf("expected value: %v != actual value: %v", expected, actual)
	}

	for i, _ := range expected {
		if expected[i] != actual[i] {
			t.Errorf("expected value: %v != actual value: %v", expected, actual)
		}
	}
}

func CompareUInt64(expected uint64, actual uint64, t *testing.T) {
	if expected != actual {
		t.Errorf("expected value: %v != actual value: %v", expected, actual)
	}
}

func CompareInt32(expected int32, actual int32, t *testing.T) {
	if expected != actual {
		t.Errorf("expected value: %v != actual value: %v", expected, actual)
	}
}

func CompareInt64(expected int64, actual int64, t *testing.T) {
	if expected != actual {
		t.Errorf("expected value: %v != actual value: %v", expected, actual)
	}
}

func CompareFloat32(expected float32, actual float32, t *testing.T) {
	if expected != actual {
		t.Errorf("expected value: %v != actual value: %v", expected, actual)
	}
}

func CompareFloat64(expected float64, actual float64, t *testing.T) {
	if expected != actual {
		t.Errorf("expected value: %v != actual value: %v", expected, actual)
	}
}

func CompareErrorMessage(expected error, actual error, t *testing.T) {
	if actual.Error() != expected.Error() {
		t.Errorf("expected value: %v != actual value: %v", expected.Error(), actual.Error())
	}
}

func ErrorNil(actual error, t *testing.T) {
	if actual != nil {
		t.Errorf("expected nil but got: %s", actual.Error())
	}
}

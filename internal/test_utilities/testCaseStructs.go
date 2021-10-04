package test_utilities

import "io"

type TestCaseError struct {
	Reader io.Reader
	Err    error
}

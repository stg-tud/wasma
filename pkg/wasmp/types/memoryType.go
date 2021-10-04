package types

import (
	"errors"
	"io"
	"strings"
)

func NewMemoryType(reader io.Reader) (*Limit, error) {
	limit, err := NewLimit(reader)

	if err == nil {
		return limit, err
	} else {
		return limit, errors.New(strings.ReplaceAll(err.Error(),"Error while reading Limit.", "Error while reading memory type."))
	}
}
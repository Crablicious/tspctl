package client

import (
	"fmt"
)

type OutputType string

const (
	OutputTypeTable = "table"
	OutputTypeJSON  = "json"
)

var OutputTypes = []string{OutputTypeTable, OutputTypeJSON}

type UnsupportedError struct {
	unsupported OutputType
}

func NewUnsupportedError(unsupported OutputType) error {
	return &UnsupportedError{unsupported}
}

func (e *UnsupportedError) Error() string {
	return fmt.Sprintf("unsupported output type \"%s\"", e.unsupported)
}

func contains(v string, a []string) bool {
	for _, i := range a {
		if i == v {
			return true
		}
	}
	return false
}

func (o *OutputType) String() string {
	return string(*o)
}

func (o *OutputType) Set(v string) error {
	if !contains(v, OutputTypes) {
		return fmt.Errorf("must be one of %v", OutputTypes)
	}
	*o = OutputType(v)
	return nil
}

func (o *OutputType) Type() string {
	return "type"
}

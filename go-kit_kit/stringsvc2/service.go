package main

import (
	"errors"
	"strings"
)

/* ----------------------------- Business Logic ----------------------------- */
// In Go kit, we model a service as an interface
// StringService provides operations on strings.
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

type stringService struct{}

func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}

// ErrEmpty is returned when input string is empty.
var ErrEmpty = errors.New("empty string")

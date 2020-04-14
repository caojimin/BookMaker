package errors

import "strings"

func New(texts ...string) error {
	return &errorString{strings.Join(texts, " ")}
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return "BookMaker: " + e.s
}

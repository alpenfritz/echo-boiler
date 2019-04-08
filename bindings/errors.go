package bindings

import (
	"strings"
)

// RequestErrors - Struct for RequestErrors
type RequestErrors struct {
	errs []error
}

// Append - Append new error to RequestErrors struct
func (re *RequestErrors) Append(err error) {
	re.errs = append(re.errs, err)
}

// Len - Return number of errors in RequestErrors struct
func (re *RequestErrors) Len() int {
	return len(re.errs)
}

// Error - Return string of all errors in RequestErrors struct
func (re *RequestErrors) Error() string {
	errstr := []string{}
	for _, e := range re.errs {
		errstr = append(errstr, e.Error())
	}
	return strings.Join(errstr, ", ")
}

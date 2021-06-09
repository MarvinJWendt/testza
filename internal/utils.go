package internal

import (
	"testing"
)

type testingT interface {
	Error(args ...interface{})
}

type helper interface {
	Helper()
}

// GetTest converts to *testing.T.
func GetTest(t testingT) *testing.T {
	if test, ok := t.(*testing.T); ok {
		return test
	}

	return nil
}

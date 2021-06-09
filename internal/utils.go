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

func GetTest(t testingT) *testing.T {
	if test, ok := t.(*testing.T); ok {
		return test
	}

	return nil
}

package internal

func Fail(t testingT, args ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	t.Error(args...)
}

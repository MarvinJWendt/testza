package testutil

import (
	"math/rand"
	"testing"

	"github.com/atomicgo/testutil/internal"
)

// StringsHelper contains strings test sets.
type StringsHelper struct{}

// Usernames returns a test set with usernames.
func (s StringsHelper) Usernames() []string {
	return []string{"MarvinJWendt", "Zipper1337", "n00b", "l33t"}
}

// HtmlTags returns a test set with html tags.
func (s StringsHelper) HtmlTags() []string {
	return []string{
		"<script>alert('XSS')</script>",
		"<script>",
		`<a href="https://github.com/atomicgo/testutil">link</a>`,
		`</body>`,
		`</html>`,
	}
}

// All contains all string test sets plus ten generated random StringsHelper.
func (s StringsHelper) All() (ret []string) {
	ret = append(ret, s.Usernames()...)
	ret = append(ret, s.HtmlTags()...)
	ret = append(ret, s.GenerateRandom(rand.Intn(10), 10)...)

	return
}

// Limit limits a test set in size.
func (s StringsHelper) Limit(testSet []string, max int) []string {
	if len(testSet) <= max {
		return testSet
	}

	return testSet[:max]
}

// GenerateRandom returns random StringsHelper in a test set.
func (s StringsHelper) GenerateRandom(length, count int) (result []string) {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	for i := 0; i < count; i++ {
		str := make([]rune, length)

		for i := range str {
			str[i] = letters[rand.Intn(len(letters))]
		}
		result = append(result, string(str))
	}
	return
}

// RunTests runs tests with a specific test set.
func (s StringsHelper) RunTests(t testingT, testSet []string, testFunc func(t *testing.T, index int, str string)) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	test := internal.GetTest(t)
	if test == nil {
		t.Error("Cannot run sub tests for test that is not a type of *testing.T")
		return
	}

	for i, str := range testSet {
		test.Run(str, func(t *testing.T) {
			t.Helper()

			testFunc(t, i, str)
		})
	}
}

// Modify returns a modified version of a test set.
func (s StringsHelper) Modify(inputSlice []string, f func(index int, value string) string) (ret []string) {
	for i, str := range inputSlice {
		ret = append(ret, f(i, str))
	}

	return
}

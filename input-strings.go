package testutil

import (
	"math/rand"
	"testing"
	"time"
	"unsafe"

	"github.com/atomicgo/testutil/internal"
)

type strings struct{}

// Usernames returns a test set with usernames.
func (s strings) Usernames() []string {
	return []string{"MarvinJWendt", "Zipper1337", "n00b", "l33t"}
}

// HtmlTags returns a test set with html tags.
func (s strings) HtmlTags() []string {
	return []string{
		"<script>alert('XSS')</script>",
		"<script>",
		`<a href="https://github.com/atomicgo/testutil">link</a>`,
		`</body>`,
		`</html>`,
	}
}

// All contains all string test sets plus ten generated random strings.
func (s strings) All() (ret []string) {
	ret = append(ret, s.Usernames()...)
	ret = append(ret, s.HtmlTags()...)
	ret = append(ret, s.GenerateRandom(rand.Intn(10), 10)...)

	return
}

// Limit limits a test set in size.
func (s strings) Limit(testSet []string, max int) []string {
	if len(testSet) <= max {
		return testSet
	}

	return testSet[:max]
}

// GenerateRandom returns random strings in a test set.
func (s strings) GenerateRandom(lenght, count int) (result []string) {
	const (
		letters       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		letterIdxBits = 6
		letterIdxMask = 1<<letterIdxBits - 1
		letterIdxMax  = 63 / letterIdxBits
	)

	for i := 0; i < count; i++ {
		b := make([]byte, lenght)
		var src = rand.NewSource(time.Now().UnixNano() + int64(i))
		for i, cache, remain := lenght-1, src.Int63(), letterIdxMax; i >= 0; {
			if remain == 0 {
				cache, remain = src.Int63(), letterIdxMax
			}
			if idx := int(cache & letterIdxMask); idx < len(letters) {
				b[i] = letters[idx]
				i--
			}
			cache >>= letterIdxBits
			remain--
		}
		result = append(result, *(*string)(unsafe.Pointer(&b)))
	}

	return
}

// RunTests runs tests with a specific test set.
func (s strings) RunTests(t TestingT, testSet []string, testFunc func(t *testing.T, index int, str string)) {
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
func (s strings) Modify(inputSlice []string, f func(index int, value string) string) (ret []string) {
	for i, str := range inputSlice {
		ret = append(ret, f(i, str))
	}

	return
}

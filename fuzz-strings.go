package testza

import (
	"math/rand"
)

// FuzzStringEmpty returns a test set with a single empty string.
func FuzzStringEmpty() []string {
	return []string{""}
}

// FuzzStringLong returns a test set with long random strings.
// Returns:
// [0]: Random string (length: 25)
// [1]: Random string (length: 50)
// [2]: Random string (length: 100)
// [3]: Random string (length: 1,000)
// [4]: Random string (length: 100,000)
func FuzzStringLong() (testSet []string) {
	testSet = append(testSet, FuzzStringGenerateRandom(1, 25)...)
	testSet = append(testSet, FuzzStringGenerateRandom(1, 50)...)
	testSet = append(testSet, FuzzStringGenerateRandom(1, 100)...)
	testSet = append(testSet, FuzzStringGenerateRandom(1, 1_000)...)
	testSet = append(testSet, FuzzStringGenerateRandom(1, 100_000)...)

	return
}

// FuzzStringNumeric returns a test set with strings that are numeric.
// The highest number in here is "9223372036854775807", which is equal to the maxmim int64.
func FuzzStringNumeric() []string {
	positiveNumbers := []string{"0", "1", "2", "3", "100", "1.1", "1337", "13.37", "0.000000000001", "9223372036854775807"}
	negativeNumbers := FuzzUtilModifySet(positiveNumbers, func(index int, value string) string { return "-" + value })
	return append(positiveNumbers, negativeNumbers...)
}

// FuzzStringUsernames returns a test set with usernames.
func FuzzStringUsernames() []string {
	return []string{"MarvinJWendt", "Zipper1337", "n00b", "l33t", "j0rgan", "test", "test123", "TEST", "test_", "TEST_"}
}

// FuzzStringEmailAddresses returns a test set with valid email addresses.
// The addresses may look like they are invalid, but they are all conform to RFC 2822 and could be used.
// You can use this test set to test your email validation process.
func FuzzStringEmailAddresses() []string {
	return []string{
		"hello@world.com",
		"hello+world@example.com",
		"hello.world@example.com",
		"a@a.xyz",
		"test@127.0.0.1",
		"test@[127.0.0.1]",
		"1@example.com",
		"_____@example.com",
		"test@subdomain.domain.xyz",
		`valid.”email\ address@example.com`,
		`first.last@iana.org`,
		`1234567890123456789012345678901234567890123456789012345678901234@iana.org`,
		`"first\"last"@iana.org`,
		`"first@last"@iana.org`,
		`"first\\last"@iana.org`,
		`x@x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x2`,
		`1234567890123456789012345678901234567890123456789012345678@12345678901234567890123456789012345678901234567890123456789.12345678901234567890123456789012345678901234567890123456789.123456789012345678901234567890123456789012345678901234567890123.iana.org`,
		`first.last@[12.34.56.78]`,
		`first.last@[IPv6:::12.34.56.78]`,
		`first.last@[IPv6:::b3:b4]`,
		`first.last@[IPv6:::]`,
		`first.last@[IPv6:1111:2222:3333::4444:12.34.56.78]`,
		`"first\last"@iana.org`,
		`user+mailbox@iana.org`,
		`customer/department@iana.org `,
		`customer/department=shipping@iana.org`,
		`"Doug \"Ace\" L."@iana.org`,
		`+1~1+@iana.org`,
		`{_test_}@iana.org`,
		`"[[ test ]]"@iana.org`,
		`"test&#13;&#10; blah"@iana.org`,
		`(foo)cal(bar)@(baz)iamcal.com(quux)`,
		`cal(woo(yay)hoopla)@iamcal.com`,
		`cal(foo\@bar)@iamcal.com`,
		`cal(foo\)bar)@iamcal.com`,
		`first(Welcome to&#13;&#10; the ("wonderful" (!)) world&#13;&#10; of email)@iana.org`,
		`pete(his account)@silly.test(his host)`,
		`c@(Chris's host.)public.example`,
	}
}

// FuzzStringHtmlTags returns a test set with different html tags.
//
// Example:
//   - <script>
//   - <script>alert('XSS')</script>
//   - <a href="https://github.com/MarvinJWendt/testza">link</a>
func FuzzStringHtmlTags() []string {
	return []string{
		"<script>alert('XSS')</script>",
		"<script>",
		`<a href="https://github.com/MarvinJWendt/testza">link</a>`,
		`</body>`,
		`</html>`,
	}
}

// FuzzStringFull contains all string test sets plus ten generated random strings.
// This test set is huge and should only be used if you want to make sure that no string, at all, can crash a process.
func FuzzStringFull() (ret []string) {
	ret = append(ret, FuzzStringUsernames()...)
	ret = append(ret, FuzzStringHtmlTags()...)
	ret = append(ret, FuzzStringEmailAddresses()...)
	ret = append(ret, FuzzStringEmpty()...)
	ret = append(ret, FuzzStringNumeric()...)
	ret = append(ret, FuzzStringLong()...)

	for i := 0; i < 10; i++ {
		ret = append(ret, FuzzStringGenerateRandom(1, i)...)
	}

	return
}

// FuzzStringGenerateRandom returns random strings in a test set.
func FuzzStringGenerateRandom(count, length int) (result []string) {
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

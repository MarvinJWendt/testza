package testza

import (
	"io"
	"os"
)

func MockStdinString(t testRunner, f func() error) (output string, err error) {
	r, w, err := os.Pipe()
	if err != nil {
		return "", err
	}

	_, err = w.Write([]byte("Hello World\n"))
	if err != nil {
		return "", err
	}
	err = w.Close()
	if err != nil {
		return "", err
	}

	stdin := os.Stdin
	os.Stdin = r

	output, err = CaptureStdout(func(w io.Writer) error {
		return f()
	})
	if err != nil {
		return "", err
	}

	os.Stdin = stdin

	return
}

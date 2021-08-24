package testza

import (
	"fmt"
	"io"
	"os"
)

func MockStdinString(t testRunner, f func() error) (output string, err error) {
	r, w, err := os.Pipe()
	if err != nil {
		return "", fmt.Errorf("could not mock stdin: %w", err)
	}

	_, err = w.Write([]byte("Hello World\n"))
	if err != nil {
		return "", fmt.Errorf("could not mock stdin: %w", err)
	}
	err = w.Close()
	if err != nil {
		return "", fmt.Errorf("could not mock stdin: %w", err)
	}

	stdin := os.Stdin
	os.Stdin = r

	output, err = CaptureStdout(func(w io.Writer) error {
		return f()
	})
	if err != nil {
		return "", fmt.Errorf("could not mock stdin: %w", err)
	}

	os.Stdin = stdin

	return
}

package testza

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// CaptureHelper contains methods to capture terminal output.
type CaptureHelper struct{}

// CaptureStdout captures everything written to stdout from a specific function.
// You can use this method in tests, to validate that your functions writes a string to the terminal.
func (h *CaptureHelper) Stdout(capture func(w io.Writer) error) (string, error) {
	originalStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		return "", fmt.Errorf("could not capture stdout: %w", err)
	}
	os.Stdout = w

	err = capture(w)
	if err != nil {
		return "", fmt.Errorf("error inside capture function while capturing stdout: %w", err)
	}

	err = w.Close()
	if err != nil {
		return "", fmt.Errorf("could not capture stdout: %w", err)
	}
	out, err := ioutil.ReadAll(r)
	if err != nil {
		return "", fmt.Errorf("could not capture stdout: %w", err)
	}
	os.Stdout = originalStdout
	err = r.Close()
	if err != nil {
		return "", fmt.Errorf("could not capture stdout: %w", err)
	}

	return string(out), nil
}

// CaptureStderr captures everything written to stderr from a specific function.
// You can use this method in tests, to validate that your functions writes a string to the terminal.
func (h *CaptureHelper) Stderr(capture func(w io.Writer) error) (string, error) {
	originalStderr := os.Stderr
	r, w, err := os.Pipe()
	if err != nil {
		return "", fmt.Errorf("could not capture stderr: %w", err)
	}
	os.Stderr = w

	err = capture(w)
	if err != nil {
		return "", fmt.Errorf("error inside capture function while capturing stderr: %w", err)
	}

	err = w.Close()
	if err != nil {
		return "", fmt.Errorf("could not capture stderr: %w", err)
	}
	out, err := ioutil.ReadAll(r)
	if err != nil {
		return "", fmt.Errorf("could not capture stderr: %w", err)
	}
	os.Stderr = originalStderr
	err = r.Close()
	if err != nil {
		return "", fmt.Errorf("could not capture stderr: %w", err)
	}

	return string(out), nil
}

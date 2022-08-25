package testza

import (
	"fmt"
	"io"
	"os"
)

// CaptureStdout captures everything written to stdout from a specific function.
// You can use this method in tests, to validate that your functions writes a string to the terminal.
//
// Example:
//
//	stdout, err := testza.CaptureStdout(func(w io.Writer) error {
//		fmt.Println("Hello, World!")
//		return nil
//	})
//
//	testza.AssertNoError(t, err)
//	testza.AssertEqual(t, "Hello, World!", stdout)
func CaptureStdout(capture func(w io.Writer) error) (string, error) {
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
	out, err := io.ReadAll(r)
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
//
// Example:
//
//	stderr, err := testza.CaptureStderr(func(w io.Writer) error {
//		_, err := fmt.Fprint(os.Stderr, "Hello, World!")
//		testza.AssertNoError(t, err)
//		return nil
//	})
//
//	testza.AssertNoError(t, err)
//	testza.AssertEqual(t, "Hello, World!", stderr)
func CaptureStderr(capture func(w io.Writer) error) (string, error) {
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
	out, err := io.ReadAll(r)
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

// CaptureStdoutAndStderr captures everything written to stdout and stderr from a specific function.
// You can use this method in tests, to validate that your functions writes a string to the terminal.
//
// Example:
//
//	stdout, stderr, err := testza.CaptureStdoutAndStderr(func(stdoutWriter, stderrWriter io.Writer) error {
//		fmt.Fprint(os.Stdout, "Hello")
//		fmt.Fprint(os.Stderr, "World")
//		return nil
//	})
//
//	testza.AssertNoError(t, err)
//	testza.AssertEqual(t, "Hello", stdout)
//	testza.AssertEqual(t, "World", stderr)
func CaptureStdoutAndStderr(capture func(stdoutWriter, stderrWriter io.Writer) error) (stdout, stderr string, err error) {
	originalStdout := os.Stdout
	originalStderr := os.Stderr

	stdoutR, stdoutW, err := os.Pipe()
	if err != nil {
		return "", "", fmt.Errorf("could not capture stdout or stderr: %w", err)
	}

	stderrR, stderrW, err := os.Pipe()
	if err != nil {
		return "", "", fmt.Errorf("could not capture stdout or stderr: %w", err)
	}

	os.Stdout = stdoutW
	os.Stderr = stderrW

	err = capture(stdoutW, stderrW)
	if err != nil {
		return "", "", fmt.Errorf("error inside capture function while capturing stdout or stderr: %w", err)
	}

	err = stdoutW.Close()
	if err != nil {
		return "", "", fmt.Errorf("could not capture stdout or stderr: %w", err)
	}

	err = stderrW.Close()
	if err != nil {
		return "", "", fmt.Errorf("could not capture stdout or stderr: %w", err)
	}

	stdoutOut, err := io.ReadAll(stdoutR)
	if err != nil {
		return "", "", fmt.Errorf("could not capture stdout or stderr: %w", err)
	}

	stderrOut, err := io.ReadAll(stderrR)
	if err != nil {
		return "", "", fmt.Errorf("could not capture stdout or stderr: %w", err)
	}

	os.Stdout = originalStdout
	os.Stderr = originalStderr

	err = stdoutR.Close()
	if err != nil {
		return "", "", fmt.Errorf("could not capture stdout or stderr: %w", err)
	}

	err = stderrR.Close()
	if err != nil {
		return "", "", fmt.Errorf("could not capture stdout or stderr: %w", err)
	}

	return string(stdoutOut), string(stderrOut), nil
}

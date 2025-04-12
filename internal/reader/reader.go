/**
 * Package reader - The package contains the API to interact with the various input
 * objects for the "message" command (and others).
 */
package reader

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// The `Reader` interface defines the ability to read a string of input from a source.
// The said source of the input can be either the STDIN or a file.
type Reader interface {
	Read() (string, error)
}

// `stdinReader` implements `Reader` by reading a single line from STDIN
type stdinReader struct {
	source *os.File
}

// The `NewStdinReader()` constructor creates a new instance of `stdinReader` to read
// from STDIN (using `os.Stdin`)
func NewStdinReader() *stdinReader {
	return &stdinReader{
		source: os.Stdin,
	}
}

// The `Read()` method of the `stdinReader` struct reads a line of input from STDIN.
// Thereafter, it checks whether the input is piped in (i.e., not from a TTY). If the
// input is valid, a string is returned else throws an error.
func (s *stdinReader) Read() (string, error) {
	stat, err := s.source.Stat()
	if err != nil {
		return "", fmt.Errorf("failed to stat stdin: %w", err)
	}

	// Check if STDIN is being piped (i.e., not from a TTY)
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(s.source)
		if scanner.Scan() {
			return scanner.Text(), nil
		}
		if err := scanner.Err(); err != nil {
			return "", fmt.Errorf("error scanning stdin: %w", err)
		}
		return "", fmt.Errorf("no input received from stdin")
	}

	return "", errors.New("no input received from STDIN")
}

// The `fileReader` struct implements the `reader` interface to read from a file.
type fileReader struct {
	path string
}

// The `NewFileReader()` constructor creates an instance of the `fileReader` struct with
// a given path. It returns an error if the file does not exist or is not readable.
func NewFileReader(path string) (*fileReader, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("could not resolve absolute path: %w", err)
	}
	return &fileReader{path: absPath}, nil
}

// The `Read()` method of the `fileReader` struct reads the contents of the file and
// returns it as a string. If the reading fails then an error is raised.
func (f *fileReader) Read() (string, error) {
	data, err := os.ReadFile(f.path)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %w", f.path, err)
	}
	return string(data), nil
}

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

/**
 * readStdin - Reads input from STDIN or fallback to a commit message file.
 *
 * This function first checks if there is content being piped into `STDIN`. If content
 * is found in `STDIN`, it reads and returns the first line. If `STDIN` is connected to
 * a terminal (i.e., no content is piped in), it falls back to reading the commit
 * message from the `.git/COMMIT_EDITMSG` file. If the file is found, it reads and
 * returns the first line from the commit message file. If neither of these conditions
 * are met, it returns an error.
 *
 * Parameters:
 *  - None
 *
 * Returns:
 *  - string: The content read from `STDIN` or the `.git/COMMIT_EDITMSG` file. If no
 *            content is available, it returns an empty string.
 *  - error: Returns an error if no content is found in either `STDIN` or the commit
 *           message file, or if there is an issue opening the commit message file.
 */
func ReadStdin() (string, error) {
	stat, err := os.Stdin.Stat()

	// Check if the input is from `STDIN`, else fallback to reading from the local
	// `COMMIT_EDITMSG` file
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		if scanner := bufio.NewScanner(os.Stdin); scanner.Scan() {
			return scanner.Text(), nil
		}
	} else {
		commitMessagePath := filepath.Join(".", ".git", "COMMIT_EDITMSG")
		file, _ := os.Open(commitMessagePath)
		if scanner := bufio.NewScanner(file); scanner.Scan() {
			return scanner.Text(), nil
		}
	}

	return "", fmt.Errorf("%v", err)
}

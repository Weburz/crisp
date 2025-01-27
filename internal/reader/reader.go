/**
 * Package reader - The package contains the API to interact with the various input
 * objects for the "message" command (and others).
 */
package reader

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

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

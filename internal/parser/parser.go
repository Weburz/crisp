/**
 * Package parser - The package containing the parsing logic for a git-commit message.
 */
package parser

import (
	"fmt"
	"regexp"
)

/**
 * Message - A struct to represent a git-commit Message.
 *
 * Fields:
 *   Type:        A string representing the type of the Message.
 *   Scope:       A string indicating the scope or category of the Message.
 *   Description: A brief description or title of the Message.
 *   Body:        The main content or body of the Message.
 *   Footer:      Additional information or footer text to accompany the Message.
 */
type Message struct {
	Type        string // Type of the message (e.g., "info", "error")
	Scope       string // Scope of the message (e.g., "global", "user")
	Description string // A brief description or title of the message
	Body        string // The main content or body of the message
	Footer      string // Additional footer information
}

/**
  * ParseMessage: Accept a string and parser it into a struct for validation.
  *
  * Parameters:
*   message (string): The Git commit message to parser.
  *
  * Returns:
  *   None
*/
func ParseMessage(input string) (Message, error) {
	re := regexp.MustCompile(
		`^(?P<Type>\w+)(?:\((?P<Scope>\w+)\))?: \s*(?P<Description>.+)$`,
	)

	matches := re.FindStringSubmatch(input)

	if matches == nil {
		return Message{}, fmt.Errorf("invalid commit message format")
	}

	return Message{
		Type:        matches[1],
		Scope:       matches[2],
		Description: matches[3],
	}, nil
}

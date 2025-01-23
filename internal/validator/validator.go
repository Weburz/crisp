/**
 * Package validator - Package containing the validation logic for a git-commit message.
 */
package validator

import (
	"fmt"
	"slices"
	"strings"
	"unicode"

	"github.com/Weburz/crisp/internal/parser"
)

/**
 * checkMessageType - Checks if the provided string is a valid commit message type.
 *
 * This function checks if the given commit message type (e.g., "build", "ci", "docs",
 * etc.) exists in a predefined list of valid message types and its casing. If the
 * message type is invalid, it returns an error; otherwise, it returns nil indicating
 * the type is valid.
 *
 * Parameters:
 *  - t (*string): A pointer to a string that represents the commit message type to
 *                 check. The value should be one of the predefined types such as
 *                 "build", "ci", "docs", etc.
 *
 * Returns:
 *  - error: Returns an error if the provided message type is not in the predefined
 *           list. If the message type is valid, it returns nil indicating no error.
 */
func checkMessageType(t *string) error {
	// See the list of valid message types in the documentation below:
	// https://github.com/angular/angular/blob/22b96b9/CONTRIBUTING.md#type
	types := []string{
		"build",
		"ci",
		"docs",
		"feat",
		"fix",
		"perf",
		"refactor",
		"style",
		"test",
		"chore",
	}

	// Commit message types should be in lower casing, hence check for it as well as the
	// fact that it is an accepted in type in accordance to established conventions
	if *t != strings.ToLower(*t) {
		if !slices.Contains(types, strings.ToLower(*t)) {
			return fmt.Errorf("invalid commit message type: %s", *t)
		}

		return fmt.Errorf(
			"invalid commit message casing, \"%s\" should be \"%s\"",
			*t,
			strings.ToLower(*t),
		)
	}

	// Return no error if the validation was a success
	return nil
}

/**
 * checkMessageScope - Validates the casing of a commit message scope.
 *
 * This function checks if the provided commit message scope is in lowercase. If the
 * scope is provided and not in lowercase, it returns an error with the correct casing
 * suggestion. If no scope is provided, no validation occurs.
 *
 * Parameters:
 *  - s (*string): A pointer to the commit message scope. It can be an empty string or
 *    a valid scope (e.g., "feat", "fix").
 *
 * Returns:
 *  - error: Returns an error if the scope is provided and not in lowercase.
 *    Returns nil if the scope is valid or empty.
 */
func checkMessageScope(s *string) error {
	if *s != "" {
		if *s != strings.ToLower(*s) {
			return fmt.Errorf(
				"invalid commit message scope casing, \"%s\" should be \"%s\"",
				*s,
				strings.ToLower(*s),
			)
		}
	}

	return nil
}

/**
 * checkMessageScope - Validates the casing of a commit message scope.
 *
 * This function checks if the provided commit message scope is lowercase. If the scope
 * is provided and not in lowercase, it returns an error with a suggestion to correct
 * the casing. If no scope is provided, no validation is performed.
 *
 * Parameters:
 *  - s (*string): A pointer to the commit message scope. Can be an empty string or a
 *    valid scope (e.g., "feat", "fix").
 *
 * Returns:
 *  - error: Returns an error if the scope is provided and not lowercase; nil otherwise.
 */
func checkMessageSubject(s *string) error {
	if unicode.IsUpper(rune((*s)[0])) || rune((*s)[len((*s))-1]) == '.' {
		return fmt.Errorf(
			"commit message subject should be lowercased & not end with a period(.)",
		)
	}

	return nil
}

/**
 * ValidateMessage: Validate a "git-commit" message.
 *
 * Parameters:
 *   message (parser.Message): The struct to represent a "git-commit" message.
 *
 * Returns:
 *   A "success" message signifying successful validation or an error message
 *   signifying a failed validation logic.
 */
func ValidateMessage(message *parser.Message) (string, error) {
	// Validate the commit message type and return an appropriate message else an error
	if err := checkMessageType(&message.Type); err != nil {
		return "", fmt.Errorf("%s", err)
	}

	// Validate the commit message scope
	if err := checkMessageScope(&message.Scope); err != nil {
		return "", fmt.Errorf("%s", err)
	}

	// Validate the commit message subject
	if err := checkMessageSubject(&message.Description); err != nil {
		return "", fmt.Errorf("%s", err)
	}

	// Return a success message if the commit message validation was successful
	return "valid commit message", nil
}

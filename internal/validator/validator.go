// The package `validator` provides the validation logic for Git commit messages
// following the Conventional Commits specifications
package validator

import (
	"errors"
	"fmt"
	"slices"
	"strings"
	"unicode"

	"github.com/Weburz/crisp/internal/parser"
)

type validator struct{}

// The NewValidator() constructor creates and returns an instance of the validator
// struct
func NewValidator() *validator {
	return &validator{}
}

// The isValidType() method validates the type of the commit message.
//
// It checks whether the type is one of the allowed Conventional Commit types and
// ensures its is written in lowercase. If the validation fails, then throw an error.
//
// Reference: https://github.com/angular/angular/blob/22b96b9/CONTRIBUTING.md#type
func (v *validator) isValidType(s string) error {
	validTypes := []string{
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

	normalized := strings.ToLower(s)
	if !slices.Contains(validTypes, normalized) {
		return fmt.Errorf("invalid commit message type: %s", s)
	}

	if s != normalized {
		return fmt.Errorf(
			"invalid commit message casing, \"%s\" should be \"%s\"",
			s,
			normalized,
		)
	}

	return nil
}

// The isValidScope method validates the scope of the commit message.
//
// It ensure the scope is lowercased (if provided) else silently pass since the message
// scope is optional.
func (v *validator) isValidScope(s string) error {
	if s != "" {
		if s != strings.ToLower(s) {
			return fmt.Errorf(
				"invalid commit message scope casing, \"%s\" should be \"%s\"",
				s,
				strings.ToLower(s),
			)
		}
	}

	return nil
}

// isValidSubject() validates the subject of the commit message.
//
// The subject must start with a lowercase letter and must not end with a period.
// Additionally, the subject is compulsory and it will throw an error if not provided.
func (v *validator) isValidSubject(s string) error {
	if len(s) == 0 {
		return errors.New("commit message subject is empty")
	}

	if unicode.IsUpper(rune((s)[0])) || rune((s)[len((s))-1]) == '.' {
		return fmt.Errorf(
			"commit message subject should be lowercased & not end with a period(.)",
		)
	}

	return nil
}

// ValidateMessage() validates a Conventional Commit message.
//
// It checks the type, scope and subject of the committ message for validation errors.
// If any of the validation fails, then throw an error or retturn a message signifying
// a successfull validation.
func ValidateMessage(s *parser.CommitMessage) (string, error) {
	v := NewValidator()

	// Validate the commit message type and return an appropriate message else an error
	if err := v.isValidType(s.Type); err != nil {
		return "", fmt.Errorf("%w", err)
	}

	// Validate the commit message scope
	if err := v.isValidScope(s.Scope); err != nil {
		return "", fmt.Errorf("%s", err)
	}

	// Validate the commit message subject
	if err := v.isValidSubject(s.Description); err != nil {
		return "", fmt.Errorf("%s", err)
	}

	// Return a success message if the commit message validation was successful
	return "valid commit message", nil
}

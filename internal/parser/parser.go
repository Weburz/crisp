// The `parser` package provides the functionality to parse Git commit messages in
// accordance to the Conventional Commits specifications.
package parser

import (
	"errors"
	"regexp"
	"strings"
)

// The `Parser` interface defines the method to parse a commit message and return
// structured data or throw an error if parsing fails.
type Parser interface {
	Parse(message string) (*CommitMessage, error)
}

// The `CommitMessage` struct represents a parsed Git commit message following the
// Conventional Commits specifications and it includes:
//   - Type: The type of change (e.g., feat, fix, chore, and so on).
//   - Scope: The optional scope of change (e.g., parser, auth).
//   - Description: A short summary of the change.
//   - Body: The optional detailed description of the change.
//   - Footer: The optional metadata such as a breaking changes or issues closed info.
type CommitMessage struct {
	Type, Scope, Description, Body, Footer string
	IsBreaking                             bool
}

// The `parserImpl` is the internal implementation of the `Parser` interface.
type parserImpl struct{}

// The `NewParser` function creates and returns an implementation of the `Parser`
// implementation.
func NewParser() Parser {
	return &parserImpl{}
}

// The `Parse` takes a raw commit message string and parses it into a `CommitMessage`
// struct. The message should follow the Conventional Commits format, e.g.:
//
//	feat(parser): add support for new syntax
//
//	This adds support for a new @foo tag in the parser.
//
//	BREAKING CHANGE: parsing of legacy tags is no longer supported.
//
// The method returns an error if the header is malformed or missing.
func (p *parserImpl) Parse(message string) (*CommitMessage, error) {
	lines := strings.Split(message, "\n")
	if len(lines) == 0 || strings.TrimSpace(lines[0]) == "" {
		return nil, errors.New("no commit message was passed")
	}

	headerPattern := regexp.MustCompile(`^([a-zA-Z]+)(?:\(([^)]+)\))?(!)?:\s*(.+)$`)
	matches := headerPattern.FindStringSubmatch(lines[0])
	if matches == nil {
		return nil, errors.New("failed to parse the commit header")
	}

	parsed := &CommitMessage{
		Type:        matches[1],
		Scope:       matches[2],
		IsBreaking:  matches[3] == "!",
		Description: matches[4],
	}

	if len(lines) > 1 {
		bodyAndFooter := strings.Join(lines[1:], "\n")
		parts := strings.SplitN(bodyAndFooter, "\n\n", 2)
		parsed.Body = strings.TrimSpace(parts[0])
		if len(parts) > 1 {
			parsed.Footer = strings.TrimSpace(parts[1])
		}
	}

	return parsed, nil
}

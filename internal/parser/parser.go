// Package parser provides functionality to parse Git commit messages
// according to the Conventional Commits specification.
package parser

import (
	"fmt"
	"regexp"
	"strings"
)

// CommitMessage represents a structured Git commit message.
// The struct fields corresponds to components in the Conventional Commits specification
type CommitMessage struct {
	Type, Scope, Description, Body string
	Footers                        map[string]string
}

// parseHeader extracts the type, scope and description from the commit message header.
// It returns an error if the header does not conform to the Conventional Commits
// format.
//
// Example:
//
//	Input: "feat(parser): add support for new syntax"
//	Output: "feat", "parser", "add support for new syntax", nil
func parseHeader(header string) (string, string, string, error) {
	// Create the regex pattern to parse the header
	re := regexp.MustCompile(
		`^(?P<Type>\w+)(?:\((?P<Scope>[^\)]+)\))?: (?P<Description>.+)$`,
	)

	// Parse the header into it sections (or throw an error on parsing failure)
	match := re.FindStringSubmatch(header)
	if match == nil {
		return "", "", "", fmt.Errorf(
			"error: invalid commit message: %q\n\n"+
				"info: acceptable message structure is:\n"+
				"<TYPE>(<SCOPE>): <DESCRIPTION>\n\n"+
				"refer to the Conventional Commits specifications for guidance:\n"+
				"https://www.conventionalcommits.org",
			header,
		)
	}

	// Get the list of individual sections of the commit message header and the content
	groupNames := re.SubexpNames()
	var typ, scope, desc string
	for idx, name := range groupNames {
		switch name {
		case "Type":
			typ = match[idx]
		case "Scope":
			scope = match[idx]
		case "Description":
			desc = match[idx]
		}
	}

	// Return the header content (without an error)
	return typ, scope, desc, nil

}

// isKnownFooter checks whether a given key is a recognised Conventional Commits footer.
// Returns true if the key matches one of the known footers.
func isKnownFooter(key string) bool {
	switch key {
	case "BREAKING CHANGE", "Closes", "Fixes", "Refs":
		return true
	default:
		return false
	}
}

// tryParseFooter attempts to parse a single line into a known footer key-value pair.
// Returns the key, value and true if the line is a valid footer, otherwise returns
// false.
func tryParseFooter(line string) (string, string, bool) {
	// Split the commit message footer into two parts at the ":" (colon) seperator
	parts := strings.SplitN(line, ":", 2)
	if len(parts) != 2 {
		return "", "", false
	}

	key := strings.TrimSpace(parts[0])
	val := strings.TrimSpace(parts[1])

	// Check if the key-value pair of the footer are recognised according to the
	// Conventional Commits specification
	if isKnownFooter(key) {
		return key, val, true
	}

	return "", "", false
}

// parseBodyAndFooter splits the commit message lines into body and footers.
// Returns a cleaned body text and a map of parsed footers.
func parseBodyAndFooter(lines []string) (string, map[string]string) {
	bodyLines := []string{}        // The body content initially set to an empty string
	footers := map[string]string{} // The footers initially set to an empty map
	inFooter := false              // Boolean flag to check body/footer parsing state

	// Loop through the lines, trimming all whitespaces and try to parse the body/footer
	for _, line := range lines {
		line = strings.TrimSpace(line)

		// On an empty line, if it is not a footer, continue looping through the content
		// and parsing it
		if line == "" {
			if !inFooter {
				bodyLines = append(bodyLines, "")
			}
			continue
		}

		// Parse the footer content and construct the "footers" map
		if key, val, ok := tryParseFooter(line); ok {
			footers[key] = val
			inFooter = true
			continue
		}

		// If the parsing logic is outside the footer section then append the body
		// strings to the list of the body content
		if !inFooter {
			bodyLines = append(bodyLines, line)
		}
	}

	// Construct the commit message body from the list of the body content parsed above
	body := strings.TrimSpace(strings.Join(bodyLines, "\n"))

	return body, footers
}

// ParseCommitMessage parses a commit message string into its components, including
// header, body and footers.
//
// The commit message is expected to the follow the Conventional Commits format::
//
//	feat(parser): add support for new syntax
//
//	This adds support for a new @foo tag in the parser.
//
//	BREAKING CHANGE: parsing of legacy tags is no longer supported.
//
// Returns a populated CommitMessage struct or an error if parsing fails
func ParseCommitMessage(message string) (*CommitMessage, error) {
	// Split the commit message file content for further parsing (and processing)
	lines := strings.Split(message, "\n")

	// Check if the commit message file content is empty, if so throw an error
	if len(lines) == 0 || strings.TrimSpace(lines[0]) == "" {
		return nil, fmt.Errorf("no commit message was passed")
	}

	// Parse the commit message header into its individual sections (or throw an error
	// on parsing failure)
	typ, scope, desc, err := parseHeader(lines[0])
	if err != nil {
		return nil, err
	}

	// Parse the body and footer contents of the commit message
	body, footers := parseBodyAndFooter(lines[1:])

	// Return an instantiated struct for further processing and validation if no errors
	// were raised earlier
	return &CommitMessage{
		Type:        typ,
		Scope:       scope,
		Description: desc,
		Body:        body,
		Footers:     footers,
	}, nil
}

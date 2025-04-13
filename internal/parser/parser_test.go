package parser

import (
	"reflect"
	"testing"
)

func TestParse_ValidMessages(t *testing.T) {
	parser := NewParser()

	tests := []struct {
		name     string
		input    string
		expected *CommitMessage
	}{
		{
			name:  "Header only",
			input: "feat(auth): add login feature",
			expected: &CommitMessage{
				Type:        "feat",
				Scope:       "auth",
				Description: "add login feature",
			},
		},
		{
			name: "Header and Body",
			input: `fix(parser): handle empty lines

This fixes a bug where the parser crashed when encountering empty lines.`,
			expected: &CommitMessage{
				Type:        "fix",
				Scope:       "parser",
				Description: "handle empty lines",
				Body:        "This fixes a bug where the parser crashed when encountering empty lines.",
			},
		},
		{
			name: "Header, Body and Footer",
			input: `feat(core): add config file support

Allows reading from config.json and merging it with CLI args.

BREAKING CHANGE: config CLI flags have changed.`,
			expected: &CommitMessage{
				Type:        "feat",
				Scope:       "core",
				Description: "add config file support",
				Body:        "Allows reading from config.json and merging it with CLI args.",
				Footer:      "BREAKING CHANGE: config CLI flags have changed.",
			},
		},
		{
			name: "Header with no scope",
			input: `chore: update dependencies

Minor version bumps for all packages.`,
			expected: &CommitMessage{
				Type:        "chore",
				Scope:       "",
				Description: "update dependencies",
				Body:        "Minor version bumps for all packages.",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parser.Parse(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %+v, got %+v", tt.expected, result)
			}
		})
	}
}

func TestParse_InvalidMessages(t *testing.T) {
	parser := NewParser()

	tests := []struct {
		name  string
		input string
	}{
		{"Empty message", ""},
		{"Whitespace only", "   \n  \n"},
		{"Missing colon", "feat(auth) add login feature"},
		{"No type prefix", "add login feature"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := parser.Parse(tt.input)
			if err == nil {
				t.Errorf("expected an error for input: %q", tt.input)
			}
		})
	}
}

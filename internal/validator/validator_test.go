package validator

import (
	"testing"

	"github.com/Weburz/crisp/internal/parser"
)

func TestIsValidType(t *testing.T) {
	v := NewValidator()

	tests := []struct {
		input       string
		expectError bool
	}{
		{"feat", false},
		{"FIX", true},
		{"unknown", true},
		{"docs", false},
		{"ReFacTor", true},
	}

	for _, tt := range tests {
		err := v.isValidType(tt.input)
		if tt.expectError && err == nil {
			t.Errorf("expected error for input %q, but got nil", tt.input)
		}
		if !tt.expectError && err != nil {
			t.Errorf("did not expect error for input %q, but got: %v", tt.input, err)
		}
	}
}

func TestIsValidScope(t *testing.T) {
	v := NewValidator()

	tests := []struct {
		scope       string
		expectError bool
	}{
		{"", false},
		{"parser", false},
		{"Parser", true},
		{"Auth", true},
	}

	for _, tt := range tests {
		err := v.isValidScope(tt.scope)
		if tt.expectError && err == nil {
			t.Errorf("expected error for scope %q, but got nil", tt.scope)
		}
		if !tt.expectError && err != nil {
			t.Errorf("did not expect error for scope %q, but got: %v", tt.scope, err)
		}
	}
}

func TestIsValidSubject(t *testing.T) {
	v := NewValidator()

	tests := []struct {
		description string
		expectError bool
	}{
		{"add new feature", false},
		{"Add new feature", true},
		{"fix user login.", true},
		{"", true},
	}

	for _, tt := range tests {
		err := v.isValidSubject(tt.description)
		if tt.expectError && err == nil {
			t.Errorf("expected error for description %q, but got nil", tt.description)
		}
		if !tt.expectError && err != nil {
			t.Errorf(
				"did not expect error for description %q, but got: %v",
				tt.description,
				err,
			)
		}
	}
}

func TestValidateMessage(t *testing.T) {
	tests := []struct {
		name        string
		msg         *parser.CommitMessage
		expectError bool
	}{
		{
			name: "valid message",
			msg: &parser.CommitMessage{
				Type:        "feat",
				Scope:       "parser",
				Description: "add parsing support",
			},
			expectError: false,
		},
		{
			name: "invalid type",
			msg: &parser.CommitMessage{
				Type:        "Feat",
				Scope:       "parser",
				Description: "add parsing support",
			},
			expectError: true,
		},
		{
			name: "invalid scope",
			msg: &parser.CommitMessage{
				Type:        "fix",
				Scope:       "Parser",
				Description: "fix the bug",
			},
			expectError: true,
		},
		{
			name: "invalid description",
			msg: &parser.CommitMessage{
				Type:        "fix",
				Scope:       "auth",
				Description: "Fix login issue.",
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ValidateMessage(tt.msg)
			if tt.expectError && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tt.expectError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

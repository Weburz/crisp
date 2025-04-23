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

func TestIsValidLength(t *testing.T) {
	v := NewValidator()

	tests := []struct {
		name        string
		input       string
		wantErr     bool
		expectedErr string
	}{
		{
			name:    "valid length message under 50 characters",
			input:   "feat(auth): add OAuth2 login flow",
			wantErr: false,
		},
		{
			name:    "valid length message exactly 50 characters",
			input:   "fix(parser): handle nil pointers in error check",
			wantErr: false,
		},
		{
			name: "invalid length message over 50 characters",
			input: "feat(user): this message definitely exceeds the fifty " +
				"character limit",
			wantErr:     true,
			expectedErr: "commit message is 69 long but expected length should be 50",
		},
		{
			name:    "empty message is valid",
			input:   "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := v.isValidLength(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error, got nil")
				} else if err.Error() != tt.expectedErr {
					t.Errorf(
						"unexpected error message:\ngot:  %s\nwant: %s",
						err.Error(),
						tt.expectedErr,
					)
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got: %v", err)
				}
			}
		})
	}
}

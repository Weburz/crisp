package parser

import (
	"testing"
)

func TestParseMessage(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantType string
		wantScope string
		wantDesc string
		wantErr  bool
	}{
		{
			name:     "Valid commit with scope",
			input:    "feat(auth): add login feature",
			wantType: "feat",
			wantScope: "auth",
			wantDesc: "add login feature",
			wantErr:  false,
		},
		{
			name:     "Valid commit without scope",
			input:    "fix: resolve bug",
			wantType: "fix",
			wantScope: "",
			wantDesc: "resolve bug",
			wantErr:  false,
		},
		{
			name:     "Invalid format",
			input:    "invalid commit message",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseMessage(tt.input)

			// Check if error matches expectation
			if (err != nil) != tt.wantErr {
				t.Errorf("Test Failed: expected error=%v, got error=%v", tt.wantErr, err)
			}

			// Validate output if no error
			if err == nil {
				if result.Type != tt.wantType {
					t.Errorf("Test Failed: expected Type='%s', got Type='%s'", tt.wantType, result.Type)
				}
				if result.Scope != tt.wantScope {
					t.Errorf("Test Failed: expected Scope='%s', got Scope='%s'", tt.wantScope, result.Scope)
				}
				if result.Description != tt.wantDesc {
					t.Errorf("Test Failed: expected Description='%s', got Description='%s'", tt.wantDesc, result.Description)
				}
			}
		})
	}
}

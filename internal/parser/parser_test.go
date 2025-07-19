package parser

import (
	"reflect"
	"testing"
)

func TestParseHeader(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedType  string
		expectedScope string
		expectedDesc  string
		expectError   bool
	}{
		{
			name:          "valid header with scope",
			input:         "feat(parser): add new feature",
			expectedType:  "feat",
			expectedScope: "parser",
			expectedDesc:  "add new feature",
		},
		{
			name:         "valid header without scope",
			input:        "fix: correct typo",
			expectedType: "fix",
			expectedDesc: "correct typo",
		},
		{
			name:        "invalid header format",
			input:       "invalid header line",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			typ, scope, desc, err := parseHeader(tt.input)
			if tt.expectError && err == nil {
				t.Fatalf("expected error but got none")
			}
			if !tt.expectError {
				if typ != tt.expectedType || scope != tt.expectedScope ||
					desc != tt.expectedDesc {
					t.Errorf(
						"expected (%s, %s, %s), got (%s, %s, %s)",
						tt.expectedType,
						tt.expectedScope,
						tt.expectedDesc,
						typ,
						scope,
						desc,
					)
				}
			}
		})
	}
}

func TestTryParseFooter(t *testing.T) {
	tests := []struct {
		input       string
		expectKey   string
		expectValue string
		expectOK    bool
	}{
		{
			"BREAKING CHANGE: update API behavior",
			"BREAKING CHANGE",
			"update API behavior",
			true,
		},
		{"Fixes: #123", "Fixes", "#123", true},
		{"Random line", "", "", false},
		{"NotAFooter - no colon", "", "", false},
	}

	for _, tt := range tests {
		key, val, ok := tryParseFooter(tt.input)
		if ok != tt.expectOK || key != tt.expectKey || val != tt.expectValue {
			t.Errorf(
				"for input %q: expected (%q, %q, %v), got (%q, %q, %v)",
				tt.input,
				tt.expectKey,
				tt.expectValue,
				tt.expectOK,
				key,
				val,
				ok,
			)
		}
	}
}

func TestParseBodyAndFooter(t *testing.T) {
	lines := []string{
		"",
		"This is a detailed explanation.",
		"BREAKING CHANGE: legacy tags are not supported",
		"Closes: #42",
	}

	expectedBody := "This is a detailed explanation."
	expectedFooters := map[string]string{
		"BREAKING CHANGE": "legacy tags are not supported",
		"Closes":          "#42",
	}

	body, footers := parseBodyAndFooter(lines)

	if body != expectedBody {
		t.Errorf("expected body %q, got %q", expectedBody, body)
	}

	if !reflect.DeepEqual(footers, expectedFooters) {
		t.Errorf("expected footers %+v, got %+v", expectedFooters, footers)
	}
}

func TestParseCommitMessage(t *testing.T) {
	message := `feat(auth): add OAuth login

Implements login with OAuth 2.0 to support third-party auth.

BREAKING CHANGE: existing login method removed
Fixes: #101`

	expected := &CommitMessage{
		Type:        "feat",
		Scope:       "auth",
		Description: "add OAuth login",
		Body:        "Implements login with OAuth 2.0 to support third-party auth.",
		Footers: map[string]string{
			"BREAKING CHANGE": "existing login method removed",
			"Fixes":           "#101",
		},
	}

	got, err := ParseCommitMessage(message)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got.Type != expected.Type || got.Scope != expected.Scope ||
		got.Description != expected.Description {
		t.Errorf("unexpected header fields: got %+v, want %+v", got, expected)
	}

	if got.Body != expected.Body {
		t.Errorf("unexpected body: got %q, want %q", got.Body, expected.Body)
	}

	if !reflect.DeepEqual(got.Footers, expected.Footers) {
		t.Errorf("unexpected footers: got %+v, want %+v", got.Footers, expected.Footers)
	}
}

func TestParseCommitMessage_InvalidHeader(t *testing.T) {
	message := `This is not a conventional commit

Just some message with no proper format`

	_, err := ParseCommitMessage(message)
	if err == nil {
		t.Fatal("expected error for invalid commit message, got none")
	}
}

func TestIsKnownFooter(t *testing.T) {
	if isKnownFooter("Reviewed-by") {
		t.Error("expected Reviewed-by to be unknown footer")
	}
	if !isKnownFooter("BREAKING CHANGE") {
		t.Error("expected BREAKING CHANGE to be known footer")
	}
}

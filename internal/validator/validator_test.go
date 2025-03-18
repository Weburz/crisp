package validator

import (
	"fmt"
	"testing"

	"github.com/Weburz/crisp/internal/parser"
)

// TestCheckMessageType tests the checkMessageType function.
func TestCheckMessageType(t *testing.T) {
	testCases := []struct {
		name          string
		messageType   string
		expectedError string
	}{
		{
			name:          "Valid Type - feat",
			messageType:   "feat",
			expectedError: "",
		},
		{
			name:          "Invalid Type - Uppercase",
			messageType:   "Feat",
			expectedError: `invalid commit message casing, "Feat" should be "feat"`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := checkMessageType(&tc.messageType)
			if tc.expectedError == "" {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
			} else {
				if err == nil || err.Error() != tc.expectedError {
					t.Errorf("expected error %q, got %q", tc.expectedError, err)
				}
			}
		})
	}
}

// TestCheckMessageScope tests the checkMessageScope function.
func TestCheckMessageScope(t *testing.T) {
	testCases := []struct {
		name          string
		scope         string
		expectedError string
	}{
		{
			name:          "Valid Scope - Lowercase",
			scope:         "user",
			expectedError: "",
		},
		{
			name:          "Valid Scope - Empty",
			scope:         "",
			expectedError: "",
		},
		{
			name:          "Invalid Scope - Uppercase",
			scope:         "User",
			expectedError: `invalid commit message scope casing, "User" should be "user"`, // Expected error message
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := checkMessageScope(&tc.scope)
			if tc.expectedError == "" {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
			} else {
				if err == nil || err.Error() != tc.expectedError {
					t.Errorf("expected error %q, got %q", tc.expectedError, err)
				}
			}
		})
	}
}

// TestCheckMessageSubject tests the checkMessageSubject function.
func TestCheckMessageSubject(t *testing.T) {
	testCases := []struct {
		name          string
		subject       string
		expectedError string
	}{
		{
			name:          "Valid Subject - Lowercase",
			subject:       "add a new feature",
			expectedError: "",
		},
		{
			name:          "Invalid Subject - Uppercase",
			subject:       "Add a new feature",
			expectedError: "commit message subject should be lowercased & not end with a period(.)",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := checkMessageSubject(&tc.subject)
			if tc.expectedError == "" {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
			} else {
				if err == nil || err.Error() != tc.expectedError {
					t.Errorf("expected error %q, got %q", tc.expectedError, err)
				}
			}
		})
	}
}

// TestCheckMessageLength tests the checkMessageLength function.
func TestCheckMessageLength(t *testing.T) {
	testCases := []struct {
		name          string
		message       string
		expectedError string
	}{
		{
			name:          "Valid Length",
			message:       "feat: add a new feature",
			expectedError: "",
		},
		{
			name:    "Invalid Length",
			message: "Add a new feature that makes the application run faster and more efficiently",
			expectedError: fmt.Sprintf("commit message exceeds 50 characters, current length: %d",
				len("Add a new feature that makes the application run faster and more efficiently")),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := checkMessageLength(&tc.message)
			if tc.expectedError == "" {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
			} else {
				if err == nil || err.Error() != tc.expectedError {
					t.Errorf("expected error %q, got %q", tc.expectedError, err)
				}
			}
		})
	}
}

// TestValidateMessage tests the ValidateMessage function.
func TestValidateMessage(t *testing.T) {
	testCases := []struct {
		name            string
		message         parser.Message
		expectedError   string
		expectedMessage string
	}{
		{
			name: "Valid Message",
			message: parser.Message{
				Type:        "feat",
				Scope:       "user",
				Description: "add a new feature",
			},
			expectedError:   "",
			expectedMessage: "valid commit message",
		},
		{
			name: "Invalid Message - Invalid Type",
			message: parser.Message{
				Type:        "Invalid",
				Scope:       "user",
				Description: "add a new feature",
			},
			expectedError:   `invalid commit message type: Invalid`,
			expectedMessage: "",
		},
		{
			name: "Invalid Message - Invalid Scope",
			message: parser.Message{
				Type:        "feat",
				Scope:       "User",
				Description: "add a new feature",
			},
			expectedError:   `invalid commit message scope casing, "User" should be "user"`,
			expectedMessage: "",
		},
		{
			name: "Invalid Message - Invalid Subject",
			message: parser.Message{
				Type:        "feat",
				Scope:       "user",
				Description: "Add a new feature",
			},
			expectedError:   "commit message subject should be lowercased & not end with a period(.)",
			expectedMessage: "",
		},
		{
			name: "Invalid Message - Exceeds Length",
			message: parser.Message{
				Type:        "feat",
				Scope:       "user",
				Description: "add a new feature that makes the application run faster and more efficiently", // More than 50 chars
			},
			expectedError: fmt.Sprintf("commit message exceeds 50 characters, current length: %d",
				len("Add a new feature that makes the application run faster and more efficiently")),
			expectedMessage: "",
		},
		{
			name: "Valid Message - Empty Scope",
			message: parser.Message{
				Type:        "feat",
				Scope:       "",
				Description: "add a new feature",
			},
			expectedError:   "",
			expectedMessage: "valid commit message",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			messageResult, err := ValidateMessage(&tc.message)

			if tc.expectedError == "" && err != nil {
				t.Errorf("expected no error, got %v", err)
			} else if err != nil && err.Error() != tc.expectedError {
				t.Errorf("expected error %q, got %q", tc.expectedError, err.Error())
			}

			if messageResult != tc.expectedMessage {
				t.Errorf("expected message %q, got %q", tc.expectedMessage, messageResult)
			}
		})
	}
}

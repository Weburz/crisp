/**
 * Package validator - Package containing the validation logic for a git-commit message.
 */
package validator

import (
	"github.com/Weburz/crisp/internal/parser"
)

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
func ValidateMessage(message *parser.Message) string {
	// Validate the message type based on the following pointers in this document:
	// https://github.com/angular/angular/blob/22b96b9/CONTRIBUTING.md#type
	return message.Type
}

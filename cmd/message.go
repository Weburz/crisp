package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/Weburz/crisp/internal/parser"
	"github.com/Weburz/crisp/internal/reader"
	"github.com/Weburz/crisp/internal/validator"
)

var shortUsage = `Lint a Git commit message.`

var longUsage = `Lint a Git commit message.

Use this command to lint Git commit messages according to the Conventional
Commit v1.0.0 specifications. To learn more about the specifications, refer to
its the documentations here - https://www.conventionalcommits.org.
`

var messageCmd = &cobra.Command{
	Use:     "message",
	Aliases: []string{"msg"},
	Short:   shortUsage,
	Long:    longUsage,
	Example: "crisp message \"chore: fix an annoying bug\"",
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		useStdin, _ := cmd.Flags().GetBool("stdin")

		var message string

		// Read from STDIN if possible else read from the flag passed to the command
		if useStdin {
			message, _ = reader.ReadStdin()
		} else {
			message = args[0]
		}

		// Convert the git-commit message into a Go struct for further validation
		msg, err := parser.ParseMessage(message)

		// Raise an error and exit with non-zero status if the parsing logic failed
		if err != nil {
			cmd.PrintErrf("error: %s\n", err)
			os.Exit(1)
		}

		// Validate the "git-commit" message
		if status, err := validator.ValidateMessage(&msg); err != nil {
			cmd.PrintErrf("error: %s\n", err)
			os.Exit(1)
		} else {
			fmt.Println(status)
		}
	},
}

func init() {
	// Add the "--stdin" flag to the message command
	messageCmd.Flags().
		BoolP("stdin", "s", false, "Read message from STDIN instead of arguments")

	// Add the "message" command to the root command
	rootCmd.AddCommand(messageCmd)
}

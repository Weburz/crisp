package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/Weburz/crisp/internal/parser"
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

		// If --stdin is set, read from stdin (piped input)
		if useStdin {
			// Read input from stdin
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				message = scanner.Text()
			}
			if err := scanner.Err(); err != nil {
				fmt.Println("Error reading from stdin:", err)
				return
			}
		} else {
			// If stdin isn't used, take the first argument as the message
			message = args[0]
		}

		// Convert the git-commit message into a Go struct for further validation
		msg, err := parser.ParseMessage(message)

		if err != nil {
			cmd.PrintErrf("error: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("Type: %s\n", msg.Type)
		fmt.Printf("Scope: %s\n", msg.Scope)
		fmt.Printf("Description: %s\n", msg.Description)
	},
}

func init() {
	// Add the "--stdin" flag to the message command
	messageCmd.Flags().
		BoolP("stdin", "s", false, "Read message from STDIN instead of arguments")

	// Add the "message" command to the root command
	rootCmd.AddCommand(messageCmd)
}

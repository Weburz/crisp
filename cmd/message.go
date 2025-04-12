package cmd

import (
	"fmt"
	"os"
	"path/filepath"

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
		var err error

		if useStdin {
			r := reader.NewStdinReader()
			message, err = r.Read()
			if err != nil {
				cmd.PrintErrf(
					"warning: failed to read stdin: %s, reading .git/COMMIT_EDITMSG\n",
					err,
				)

				commitMsgFile := filepath.Join(".", ".git", "COMMIT_EDITMSG")
				r, err := reader.NewFileReader(commitMsgFile)
				if err != nil {
					cmd.PrintErrf("error reading commit message file: %s\n", err)
					os.Exit(1)
				}
				message, err = r.Read()
				if err != nil {
					cmd.PrintErrf("error reading commit message file: %s\n", err)
					os.Exit(1)
				}
			}
		} else {
			if len(args) == 0 {
				cmd.PrintErrln("error: no commit message provided")
				os.Exit(1)
			}
			message = args[0]
		}

		msg, err := parser.ParseMessage(message)
		if err != nil {
			cmd.PrintErrf("error: %s\n", err)
			os.Exit(1)
		}

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

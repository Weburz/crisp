package cmd

import (
	"github.com/spf13/cobra"
)

var shortUsage = `
Lint a Git commit message.
`

var longUsage = `
Lint a Git commit message.

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
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rootCmd.Println(args[0])
	},
}

func init() {
	rootCmd.AddCommand(messageCmd)
}

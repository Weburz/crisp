package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var messageCmd = &cobra.Command{
	Use:     "message",
	Short:   "The commit message to lint",
	Example: "crisp message \"chore: fix an annoying bug\"",

	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}

		return fmt.Errorf("invalid colour specified: %s", args[0])
	},

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}

func init() {
	rootCmd.AddCommand(messageCmd)
}

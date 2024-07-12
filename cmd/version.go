package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the Crisp version number.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Crisp linter v0.0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

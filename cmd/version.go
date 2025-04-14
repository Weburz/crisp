package cmd

import (
	"github.com/Weburz/crisp/internal/version"
	"github.com/spf13/cobra"
)

var shortUsageHelp = "Print the version information."

var longUsageHelp = `Prints a detailed version information of the application including
version, commit hash, the build date and the version of Go used to develop it.`

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: shortUsageHelp,
	Long:  longUsageHelp,
	Run: func(cmd *cobra.Command, args []string) {
		v := version.GetVersionInfo()

		cmd.Println("### Crisp Build Information ###")
		cmd.Printf("Version: \t%s\n", v.Version)
		cmd.Printf("Git Version: \t%s\n", v.GitVersion)
		cmd.Printf("Git Commit: \t%s\n", v.GitCommit)
		cmd.Printf("Build Date: \t%s\n", v.BuildDate)
		cmd.Printf("Go Version: \t%s\n", v.GoVersion)
		cmd.Printf("Compiler: \t%s\n", v.Compiler)
		cmd.Printf("Platform: \t%s\n", v.Platform)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cliVersion = "dev"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show CLI version information",
	Long:  "Display the version information for leaseweb-cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("leaseweb-cli version %s\n", cliVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

package cmd

import (
	"fmt"
	"runtime/debug"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Long:  "Display the version and build information for leaseweb-cli",
	Run: func(cmd *cobra.Command, args []string) {
		printVersion()
	},
}

func printVersion() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Println("Version information not available")
		return
	}

	version := info.Main.Version
	if version == "(devel)" || version == "" {
		version = "development"
	}

	fmt.Printf("leaseweb-cli version %s\n", version)

	fmt.Printf("Go version: %s\n", info.GoVersion)

	for _, setting := range info.Settings {
		if setting.Key == "vcs.revision" {
			fmt.Printf("Git commit: %s\n", setting.Value[:8])
		}
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var apiKey string
var alertConfigPath string
var notifierKind string

var rootCmd = &cobra.Command{
	Use:           "leaseweb-cli",
	SilenceUsage:  true,
	SilenceErrors: true,
	Short:         "A CLI tool to interact with Leaseweb API",
	Long:          "leaseweb-cli allows you to manage Leaseweb servers via the API.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		skipAPIKeyCommands := []string{"version"}

		for _, skipCmd := range skipAPIKeyCommands {
			if cmd.Name() == skipCmd {
				return nil
			}
		}

		if apiKey == "" {
			apiKey = os.Getenv("LEASEWEB_API_KEY")
		}

		if apiKey == "" {
			return fmt.Errorf("API key is required. Set LEASEWEB_API_KEY or use --api-key flag")
		}

		InitLeasewebClient(apiKey)
		return nil
	},
}

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&apiKey, "api-key", "", "Leaseweb API key (optional, overrides LEASEWEB_API_KEY)")

	rootCmd.PersistentFlags().StringVar(
		&alertConfigPath,
		"alert-config",
		"alert-rules.yaml",
		"path to alert rules config",
	)

	rootCmd.PersistentFlags().StringVar(
		&notifierKind,
		"notifier",
		"mattermost",
		"notifier kind: mattermost|telegram|email",
	)
}

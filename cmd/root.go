package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var apiKey string

var rootCmd = &cobra.Command{
	Use:   "leaseweb-cli",
	Short: "A CLI tool to interact with Leaseweb API",
	Long:  "leaseweb-cli allows you to manage Leaseweb servers via the API.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if apiKey == "" {
			apiKey = os.Getenv("LEASEWEB_API_KEY")
		}

		if apiKey == "" {
			fmt.Println("Error: API key is required. Set LEASEWEB_API_KEY or use --api-key flag.")
			os.Exit(1)
		}

		InitLeasewebClient(apiKey)
	},
}

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&apiKey, "api-key", "", "Leaseweb API key (optional, overrides LEASEWEB_API_KEY)")
}

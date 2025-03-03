package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"syscall"

	dedicatedserver "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver/v2"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var (
	apiKeyPath     string = "/.lsw"
	ctx            context.Context
        args           []string
	leasewebClient *dedicatedserver.APIClient
)

type Client struct {
	DedicatedserverAPI dedicatedserver.DedicatedserverAPI
}

func InitLeasewebClient(apiKey string) Client {
	cfg := dedicatedserver.NewConfiguration()

        cfg.AddDefaultHeader("X-LSW-Auth", apiKey)

        dedicatedserverAPI := dedicatedserver.NewAPIClient(cfg)

        return Client{
		DedicatedserverAPI: dedicatedserverAPI.DedicatedserverAPI,
	}
}

func Login() {
	apiKey := readFile(apiKeyPath)
	if apiKey == "" {
		fmt.Println("No API key found. Please log in using `login` command.")
		os.Exit(1)
	}
	InitLeasewebClient(apiKey)
}

func Logout() {
	writeFile(apiKeyPath, "")
	fmt.Println("Logged out successfully!")
}

func getHomeDir() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error: Unable to get home directory.")
		os.Exit(1)
	}
	return dir
}

func writeFile(path, content string) {
	fullPath := getHomeDir() + path
	err := os.WriteFile(fullPath, []byte(content), 0600) // Secure file permissions
	if err != nil {
		fmt.Println("Error writing to file:", err)
		os.Exit(1)
	}
}

func readFile(path string) string {
	fullPath := getHomeDir() + path
	apiKey, err := os.ReadFile(fullPath)
	if err != nil {
		fmt.Println("Error reading API key file:", err)
		os.Exit(1)
	}
	return string(apiKey)
}

func isFileExists(path string) bool {
	fullPath := getHomeDir() + path
	_, err := os.Stat(fullPath)
	return err == nil || !errors.Is(err, os.ErrNotExist)
}

func init() {
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(logoutCmd)
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to the Leaseweb account",
	Long:  "Log in to the Leaseweb account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Enter API key: ")
		apiKey, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Println("\nError reading API key")
			os.Exit(1)
		}
		fmt.Println("\nAuthenticating...")

		// Store API key securely
		writeFile(apiKeyPath, string(apiKey))

		// Try logging in
		InitLeasewebClient(string(apiKey))
		fmt.Println("Logged in successfully!")
	},
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out from the Leaseweb account",
	Long:  "Log out from the Leaseweb account",
	Run: func(cmd *cobra.Command, args []string) {
		Logout()
	},
}

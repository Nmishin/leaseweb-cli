package cmd

import (
	"context"
	"fmt"
	"os"

	dedicatedserver "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
	"github.com/spf13/cobra"
)

func init() {
	dedicatedServerCmd.AddCommand(dedicatedServerlistCmd)
	dedicatedServerCmd.AddCommand(dedicatedServerGetCmd)
	dedicatedServerCmd.AddCommand(dedicatedServerHardwareGetCmd)
	dedicatedServerCmd.AddCommand(dedicatedServerPowerOnCmd)
	dedicatedServerCmd.AddCommand(dedicatedServerPowerOffCmd)
	dedicatedServerCmd.AddCommand(dedicatedServerCredsGetCmd)
	dedicatedServerCmd.AddCommand(dedicatedServerIpGetCmd)
	rootCmd.AddCommand(dedicatedServerCmd)
}

var dedicatedServerCmd = &cobra.Command{
	Use:   "dedicated-server",
	Short: "Get information about your Dedicated Servers and manage them",
	Long:  "Get information about your Dedicated Servers and manage them",
}

var dedicatedServerlistCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieve the list of Dedicated Servers",
	Long:  "Retrieve the list of Dedicated Servers",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		_, r, err := leasewebClient.DedicatedserverAPI.GetServerList(ctx).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetServerList``: %v\n", err)
		}

		prettyPrintResponse(r)
	},
}

var dedicatedServerGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve details of Dedicated Server",
	Long:  "Retrieve details of Dedicated Server",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		server, r, err := leasewebClient.DedicatedserverAPI.GetServer(ctx, args[0]).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `DedicatedserverAPI.GetServer``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}

		printResponse(server)
	},
}

var dedicatedServerHardwareGetCmd = &cobra.Command{
	Use:   "get-hardware",
	Short: "Retrieve details of Dedicated Server",
	Long:  "Retrieve details of Dedicated Server",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		_, r, err := leasewebClient.DedicatedserverAPI.GetServerHardware(ctx, args[0]).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling `DedicatedserverAPI.GetServerHardware`: %v\n", err)
		}

		prettyPrintResponse(r) // print HTTP response here, because API returns an error
	},
}

var dedicatedServerCredsGetCmd = &cobra.Command{
	Use:   "get-creds <serverId> <type> <username>",
	Short: "Retrieve Dedicated Server credentials",
	Long: `Retrieve credentials for a Dedicated Server by providing:

- serverId: The ID of the dedicated server.
- type: The credential type (e.g., "OPERATING_SYSTEM", "REMOTE_MANAGEMENT").
- username: The username associated with the credential.`,
	Example: `  # Get IPMI credentials for a OS
  leaseweb get-creds 12345 OPERATING_SYSTEM root

  # Get OS credentials for a remote management
  leaseweb get-creds 12345 REMOTE_MANAGEMENT admin`,
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		serverID := args[0]
		credType := dedicatedserver.CredentialType(args[1])
		username := args[2]

		ctx := context.Background()
		_, r, err := leasewebClient.DedicatedserverAPI.
			GetServerCredential(ctx, serverID, credType, username).
			Execute()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling `DedicatedserverAPI.GetServerCredential`: %v\n", err)
			os.Exit(1)
		}

		prettyPrintResponse(r)
	},
}

var dedicatedServerPowerOnCmd = &cobra.Command{
	Use:   "power-on",
	Short: "Power-on a Dedicated Server",
	Long:  "Power-on a Dedicated Server",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		_, err := leasewebClient.DedicatedserverAPI.PowerServerOn(ctx, args[0]).Execute()
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

var dedicatedServerPowerOffCmd = &cobra.Command{
	Use:   "power-off",
	Short: "Power-off a Dedicated Server",
	Long:  "Power-off a Dedicated Server",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		_, err := leasewebClient.DedicatedserverAPI.PowerServerOff(ctx, args[0]).Execute()
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

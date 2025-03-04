package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	dedicatedServerCmd.AddCommand(dedicatedServerlistCmd)
	dedicatedServerCmd.AddCommand(dedicatedServerGetCmd)
	dedicatedServerCmd.AddCommand(dedicatedServerHardwareGetCmd)
	dedicatedServerCmd.AddCommand(dedicatedServerPowerOnCmd)
	dedicatedServerCmd.AddCommand(dedicatedServerPowerOffCmd)
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
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Fprintln(os.Stderr, "Error: Missing server ID argument")
			os.Exit(1)
		}

		ctx := context.Background()
		_, r, err := leasewebClient.DedicatedserverAPI.GetServerHardware(ctx, args[0]).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling `DedicatedserverAPI.GetServerHardware`: %v\n", err)
		}

		prettyPrintResponse(r) // print HTTP response here, because API returns error
	},
}

var dedicatedServerPowerOnCmd = &cobra.Command{
	Use:   "power-on",
	Short: "Power-on a Dedicated Server",
	Long:  "Power-on a Dedicated Server",
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
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		_, err := leasewebClient.DedicatedserverAPI.PowerServerOff(ctx, args[0]).Execute()
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

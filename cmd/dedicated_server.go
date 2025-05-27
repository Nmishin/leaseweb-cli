package cmd

import (
	"context"
	"fmt"
	"time"

	dedicatedserver "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver/v2"
	"github.com/spf13/cobra"
)

func init() {
	registerDedicatedServerCommands()
	registerDedicatedServerListFlags()
}

func registerDedicatedServerCommands() {
	dedicatedServerCmd.AddCommand(dedicatedServerlistCmd)
	dedicatedServerCmd.AddCommand(dedicatedServerGetOS)
	dedicatedServerCmd.AddCommand(dedicatedServerGetCmd)
	dedicatedServerCmd.AddCommand(dedicatedServerHardwareGetCmd)
	dedicatedServerCmd.AddCommand(dedicatedServerPowerOnCmd)
	dedicatedServerCmd.AddCommand(dedicatedServerPowerOffCmd)
	dedicatedServerCmd.AddCommand(dedicatedServerCredsGetCmd)
	dedicatedServerCmd.AddCommand(dedicatedServerContractRenewalCmd)
	dedicatedServerCmd.AddCommand(dedicatedServerPowerCycleCmd)

	rootCmd.AddCommand(dedicatedServerCmd)
}

func registerDedicatedServerListFlags() {
	dedicatedServerlistCmd.Flags().Int32Var(&serverLimit, "limit", 0, "Maximum number of servers to retrieve (default unlimited)")
	dedicatedServerlistCmd.Flags().Int32Var(&serverOffset, "offset", 0, "Return results starting from the given offset")
	dedicatedServerGetOS.Flags().Int32Var(&osLimit, "limit", 0, "Maximum number of servers to retrieve (default unlimited)")
	dedicatedServerGetOS.Flags().Int32Var(&osOffset, "offset", 0, "Return results starting from the given offset")

	// Add filters for dedicated server list
	dedicatedServerlistCmd.Flags().StringVar(&reference, "reference", "", "Filter by reference")
	dedicatedServerlistCmd.Flags().StringVar(&ip, "ip", "", "Filter by IP address")
	dedicatedServerlistCmd.Flags().StringVar(&macAddress, "mac", "", "Filter by MAC address")
	dedicatedServerlistCmd.Flags().StringVar(&site, "site", "", "Filter by site")
	dedicatedServerlistCmd.Flags().StringVar(&privateRackId, "private-rack-id", "", "Filter by rack ID")
	dedicatedServerlistCmd.Flags().StringVar(&privateNetworkCapable, "private-network-capable", "", "Filter for private network capable servers")
	dedicatedServerlistCmd.Flags().StringVar(&privateNetworkEnabled, "private-network-enabled", "", "Filter for private network enabled servers")
}

var dedicatedServerCmd = &cobra.Command{
	Use:   "dedicated-server",
	Short: "Manage dedicated servers",
}

var dedicatedServerlistCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieve the list of servers",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		allServers := []dedicatedserver.Server{}

		currentOffset := serverOffset
		apiMaxLimit := int32(50)

		fetchAll := serverLimit == 0

		for {
			batchLimit := apiMaxLimit
			if !fetchAll {
				if remaining := serverLimit - int32(len(allServers)); remaining < apiMaxLimit && remaining > 0 {
					batchLimit = remaining
				}
			}

			req := leasewebClient.DedicatedserverAPI.GetServerList(ctx).
				Limit(batchLimit).
				Offset(currentOffset)

			if reference != "" {
				req = req.Reference(reference)
			}
			if ip != "" {
				req = req.Ip(ip)
			}
			if macAddress != "" {
				req = req.MacAddress(macAddress)
			}
			if site != "" {
				req = req.Site(site)
			}
			if privateRackId != "" {
				req = req.PrivateRackId(privateRackId)
			}
			if privateNetworkCapable != "" {
				req = req.PrivateNetworkCapable(privateNetworkCapable)
			}
			if privateNetworkEnabled != "" {
				req = req.PrivateNetworkEnabled(privateNetworkEnabled)
			}

			serverResponse, _, err := req.Execute()
			if err != nil {
				return fmt.Errorf("retrieving the list of servers: %w", err)
			}

			allServers = append(allServers, serverResponse.Servers...)

			if len(serverResponse.Servers) < int(batchLimit) {
				break
			}

			if !fetchAll && int32(len(allServers)) >= serverLimit {
				break
			}

			currentOffset += batchLimit
		}

		printResponse(allServers)
		return nil
	},
}

var dedicatedServerGetOS = &cobra.Command{
	Use:   "list-os",
	Short: "Retrieve the list of available operating systems",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		allOSes := []dedicatedserver.OperatingSystem{}

		currentOffset := osOffset
		apiMaxLimit := int32(50)

		fetchAll := osLimit == 0

		for {
			batchLimit := apiMaxLimit
			if !fetchAll {
				if remaining := serverLimit - int32(len(allOSes)); remaining < apiMaxLimit && remaining > 0 {
					batchLimit = remaining
				}
			}

			req := leasewebClient.DedicatedserverAPI.GetOperatingSystemList(ctx).
				Limit(batchLimit).
				Offset(currentOffset)

			osResponse, _, err := req.Execute()
			if err != nil {
				return fmt.Errorf("retrieving the list of OSes: %w", err)
			}

			allOSes = append(allOSes, osResponse.OperatingSystems...)

			if len(osResponse.OperatingSystems) < int(batchLimit) {
				break
			}

			if !fetchAll && int32(len(allOSes)) >= serverLimit {
				break
			}

			currentOffset += batchLimit
		}

		printResponse(allOSes)
		return nil
	},
}

var dedicatedServerGetCmd = &cobra.Command{
	Use:     "get",
	Short:   "Retrieve details of the server by ID",
	Example: "leaseweb-cli dedicated-server get 12345",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		server, _, err := leasewebClient.DedicatedserverAPI.GetServer(ctx, args[0]).Execute()
		if err != nil {
			return fmt.Errorf("retrieving details of the server: %w", err)
		}

		printResponse(server)
		return nil
	},
}

var dedicatedServerHardwareGetCmd = &cobra.Command{
	Use:     "get-hardware",
	Short:   "Retrieve hardware details of the server by ID",
	Example: "leaseweb-cli dedicated-server get-hardware 12345",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		_, r, err := leasewebClient.DedicatedserverAPI.GetHardware(ctx, args[0]).Execute()
		if err != nil {
			return fmt.Errorf("retrieving hardware details of the server: %w", err)
		}

		prettyPrintResponse(r)
		return nil
	},
}

var dedicatedServerContractRenewalCmd = &cobra.Command{
	Use:   "get-contract-renewal",
	Short: "Retrieve next contract renewal date in milliseconds since epoch by server ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		server, _, err := leasewebClient.DedicatedserverAPI.GetServer(ctx, args[0]).Execute()
		if err != nil {
			return fmt.Errorf("retrieving next contract renewal date: %w", err)
		}

		if server.Contract == nil || server.Contract.StartsAt == nil {
			return fmt.Errorf("contract start date is missing in API response")
		}
		if server.Contract.ContractTerm == nil {
			return fmt.Errorf("contract term is missing in API response")
		}

		startDate := *server.Contract.StartsAt
		contractTerm := int(*server.Contract.ContractTerm)

		now := time.Now()

		renewalDate := startDate
		for !renewalDate.After(now) {
			renewalDate = renewalDate.AddDate(0, contractTerm, 0)
		}

		fmt.Println(renewalDate.UnixMilli())
		return nil
	},
}

var dedicatedServerCredsGetCmd = &cobra.Command{
	Use:   "get-creds <serverId> <type> <username>",
	Short: "Retrieve the server credentials",
	Long: `Retrieve credentials for the server by providing:

- serverId: The ID of the dedicated server.
- type: The credential type (e.g., "OPERATING_SYSTEM", "REMOTE_MANAGEMENT").
- username: The username associated with the credential.`,
	Example: `  # Get credentials for a OS
  leaseweb-cli get-creds 12345 OPERATING_SYSTEM root

  # Get credentials for a remote management
  leaseweb-cli get-creds 12345 REMOTE_MANAGEMENT admin`,
	Args: cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		serverID := args[0]
		credType := dedicatedserver.CredentialType(args[1])
		username := args[2]

		ctx := context.Background()
		_, r, err := leasewebClient.DedicatedserverAPI.
			GetCredential(ctx, serverID, credType, username).
			Execute()

		if err != nil {
			return fmt.Errorf("retrieving the server credentials: %w", err)
		}

		prettyPrintResponse(r)
		return nil
	},
}

var dedicatedServerPowerOnCmd = &cobra.Command{
	Use:   "power-on",
	Short: "Power on the server",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		_, err := leasewebClient.DedicatedserverAPI.PowerOn(ctx, args[0]).Execute()
		if err != nil {
			return fmt.Errorf("power on the server:  %w", err)
		}
		return nil
	},
}

var dedicatedServerPowerOffCmd = &cobra.Command{
	Use:   "power-off",
	Short: "Power off the server",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		_, err := leasewebClient.DedicatedserverAPI.PowerOff(ctx, args[0]).Execute()
		if err != nil {
			return fmt.Errorf("power off the server:  %w", err)
		}
		return nil
	},
}

var dedicatedServerPowerCycleCmd = &cobra.Command{
	Use:   "reboot",
	Short: "Power cycle the server",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		_, err := leasewebClient.DedicatedserverAPI.PowerCycle(ctx, args[0]).Execute()
		if err != nil {
			return fmt.Errorf("power cycle the server:  %w", err)
		}
		return nil
	},
}

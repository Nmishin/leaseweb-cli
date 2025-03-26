package cmd

import (
	"context"
	"fmt"
	"os"

	dedicatedserver "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver/v2"
	"github.com/spf13/cobra"
)

func init() {
	dedicatedServerCmd.AddCommand(dedicatedServerIpGetCmd)
	dedicatedServerCmd.AddCommand(dedicatedServerIpListCmd)
	dedicatedServerIpListCmd.Flags().StringVar(&networkType, "network-type", "", "Filter by network type")
	dedicatedServerIpListCmd.Flags().StringVar(&version, "version", "", "Filter by IP version")
	dedicatedServerIpListCmd.Flags().StringVar(&nullRouted, "null-routed", "", "Filter by null-routed status")
	dedicatedServerIpListCmd.Flags().StringVar(&ips, "ips", "", "Filter by specific IPs")
	dedicatedServerIpListCmd.Flags().Int32Var(&limit, "limit", 20, "Limit the number of results")
	dedicatedServerIpListCmd.Flags().Int32Var(&offset, "offset", 0, "Offset for pagination")
}

var (
	networkType string
	version     string
	nullRouted  string
	ips         string
	limit       int32
	offset      int32
)

var dedicatedServerIpListCmd = &cobra.Command{
	Use:   "get-ips <serverId>",
	Short: "List the server IPs",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		serverID := args[0]
		ctx := context.Background()

		req := leasewebClient.DedicatedserverAPI.GetIpList(ctx, serverID)
		if networkType != "" {
			req = req.NetworkType(dedicatedserver.NetworkType(networkType))
		}
		if version != "" {
			req = req.Version(version)
		}
		if nullRouted != "" {
			req = req.NullRouted(nullRouted)
		}
		if ips != "" {
			req = req.Ips(ips)
		}
		if limit > 0 {
			req = req.Limit(limit)
		}
		if offset > 0 {
			req = req.Offset(offset)
		}

		server, _, err := req.Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		printResponse(server)
	},
}

var dedicatedServerIpGetCmd = &cobra.Command{
	Use:   "get-ip <serverId> <ip>",
	Short: "Describe the server IP",
	Long:  "Describe the server IP by server ID and IP address",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		serverID := args[0]
		ip := args[1]
		ctx := context.Background()

		req := leasewebClient.DedicatedserverAPI.GetIp(ctx, serverID, ip)
		server, _, err := req.Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		printResponse(server)
	},
}

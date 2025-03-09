package cmd

import (
	"context"
	"fmt"
	"os"

	dedicatedserver "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
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
	Short: "List Dedicated Server IPs",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		serverID := args[0]
		ctx := context.Background()

		req := leasewebClient.DedicatedserverAPI.GetServerIpList(ctx, serverID)
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

		_, r, err := req.Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		prettyPrintResponse(r)
	},
}

var dedicatedServerIpGetCmd = &cobra.Command{
	Use:   "get-ip <serverId> <ip>",
	Short: "Describe Dedicated Server IP",
	Long:  "Describe Dedicated Server IP by server ID and IP address",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		serverID := args[0]
		ip := args[1]
		ctx := context.Background()

		req := leasewebClient.DedicatedserverAPI.GetServerIp(ctx, serverID, ip)
		_, r, err := req.Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		prettyPrintResponse(r)
	},
}

package cmd

import (
	"fmt"

	"github.com/cheynewallace/tabby"
	dedicatedserver "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver/v2"
	"github.com/spf13/cobra"
)

func init() {
	dedicatedServerCmd.AddCommand(dedicatedServerlistCmd)
	dedicatedServerCmd.AddCommand(dedicatedServerGetCmd)
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
		client := dedicatedserver.NewAPIClient(dedicatedserver.NewConfiguration())
		request := client.DedicatedserverAPI
		server, _, err := request.GetServerList(ctx).Execute()
		if err == nil {
			t := tabby.New()
			t.AddHeader("#", "Id", "Asset id", "Rack", "Site", "Suite", "Unit", "Rack type")
			for i, server := range server.Servers {
				t.AddLine(i+1, server.Id, server.AssetId, server.Location.Rack, server.Location.Site, server.Location.Suite, server.Location.Unit, server.Rack.Type)
			}
			t.Print()
		} else {
			fmt.Println(err)
		}
	},
}

var dedicatedServerGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve details of Dedicated Server",
	Long:  "Retrieve details of Dedicated Server",
	Run: func(cmd *cobra.Command, args []string) {
		client := dedicatedserver.NewAPIClient(dedicatedserver.NewConfiguration())
		request := client.DedicatedserverAPI
		server, _, err := request.GetServer(ctx, args[0]).Execute()
		if err != nil {
			fmt.Println(err)
			return
		}

		t := tabby.New()
		t.AddLine("Id:", server.Id)
		t.AddLine("Asset Id:", server.AssetId)
		t.AddLine("Location (Rack):", server.Location.Rack)
		t.AddLine("Location (Site):", server.Location.Site)
		t.AddLine("Location (Suite):", server.Location.Suite)
		t.AddLine("Location (Unit):", server.Location.Unit)
		t.AddLine("Rack Type:", server.Rack.Type)
		t.Print()
	},
}

var dedicatedServerPowerOnCmd = &cobra.Command{
	Use:   "power-on",
	Short: "Power-on a Dedicated Server",
	Long:  "Power-on a Dedicated Server",
	Run: func(cmd *cobra.Command, args []string) {
		client := dedicatedserver.NewAPIClient(dedicatedserver.NewConfiguration())
		request := client.DedicatedserverAPI
		_, err := request.PowerOn(ctx, args[0]).Execute()
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
		client := dedicatedserver.NewAPIClient(dedicatedserver.NewConfiguration())
		request := client.DedicatedserverAPI
		_, err := request.PowerOff(ctx, args[0]).Execute()
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

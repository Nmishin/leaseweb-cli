package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	dedicatedserver "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver"
)

var (
	leasewebClient Client
)

type Client struct {
	DedicatedserverAPI dedicatedserver.DedicatedserverAPI
}

func InitLeasewebClient(apiKey string) {
	cfg := dedicatedserver.NewConfiguration()

	cfg.AddDefaultHeader("X-LSW-Auth", apiKey)

	leasewebClient = Client{
		DedicatedserverAPI: dedicatedserver.NewAPIClient(cfg).DedicatedserverAPI,
	}
}

func Login() {
	apiKey := os.Getenv("LEASEWEB_API_KEY")
	if apiKey == "" {
		fmt.Println("Error: LEASEWEB_API_KEY environment variable is not set.")
		os.Exit(1)
	}
	InitLeasewebClient(apiKey)
}

func printResponse(resp interface{}) {
	jsonData, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error marshalling response: %v\n", err)
		return
	}
	fmt.Fprintf(os.Stdout, "Response from `DedicatedserverAPI.GetServer`:\n%s\n", jsonData)
}

func prettyPrintResponse(r *http.Response) {
	if r == nil {
		fmt.Println("No response received")
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	defer r.Body.Close()

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, body, "", "    ")
	if err != nil {
		fmt.Println("Error formatting JSON:", err)
		return
	}

	fmt.Println(prettyJSON.String())
}

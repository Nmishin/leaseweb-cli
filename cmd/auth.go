package cmd

import (
	dedicatedserver "github.com/leaseweb/leaseweb-go-sdk/dedicatedserver/v2"
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

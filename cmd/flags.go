package cmd

var (
	// Filters for dedicated server listing
	serverLimit           int32
	serverOffset          int32
	reference             string
	ip                    string
	macAddress            string
	site                  string
	privateRackId         string
	privateNetworkCapable string
	privateNetworkEnabled string

	// Filters for dedicated server IP listing
	ipLimit     int32
	ipOffset    int32
	networkType string
	version     string
	nullRouted  string
	ips         string
)

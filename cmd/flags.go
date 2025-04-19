package cmd

var (
	// Pagination flags (used by multiple commands)
	limit    int32
	offset   int32
        fetchAll bool

	// Filters for dedicated server listing
	reference             string
	ip                    string
	macAddress            string
	site                  string
	privateRackId         string
	privateNetworkCapable string
	privateNetworkEnabled string

	// Filters for dedicated server IP listing
	networkType string
	version     string
	nullRouted  string
	ips         string
)

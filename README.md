# Leaseweb CLI

[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/Nmishin/leaseweb-cli)](https://github.com/Nmishin/leaseweb-cli/releases/latest)

![GitHub all releases](https://img.shields.io/github/downloads/Nmishin/leaseweb-cli/total?label=GitHub%20Total%20Downloads)

### About

This tool interacts with Leaseweb's API, allowing users to retrieve details about dedicated servers. It includes commands like `get` to fetch server details and `get-hardware` to retrieve hardware information.

### Installation
Download the binary from the [Releases page](https://github.com/Nmishin/leaseweb-cli/releases).

### Build
```bash
git clone git@github.com:Nmishin/leaseweb-cli.git
cd leaseweb-cli
go build -o leaseweb
```

### Generate your API Key
You can generate your API key at the [Customer Portal](https://secure.leaseweb.com/)

### Authentication
For authentication need to export API Key from previous step, or set it as `--api-key` flag.
```bash
export LEASEWEB_API_KEY=<>
```

### Usage

```bash
$ leaseweb-cli dedicated-server -h

Manage dedicated servers

Usage:
  leaseweb-cli dedicated-server [command]

Available Commands:
  get                  Retrieve details of the server by ID
  get-contract-renewal Retrieve next contract renewal date in milliseconds since epoch by server ID
  get-creds            Retrieve the server credentials
  get-hardware         Retrieve hardware details of the server by ID
  get-ip               Describe the server IP
  get-ips              List the server IPs
  list                 Retrieve the list of servers
  power-off            Power off the server
  power-on             Power on the server
  reboot               Power cycle the server

Flags:
  -h, --help   help for dedicated-server

Global Flags:
      --api-key string   Leaseweb API key (optional, overrides LEASEWEB_API_KEY)

Use "leaseweb-cli dedicated-server [command] --help" for more information about a command.
```

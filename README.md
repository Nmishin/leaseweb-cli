# Leaseweb CLI

[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/Nmishin/leaseweb-cli)](https://github.com/Nmishin/leaseweb-cli/releases/latest)
![GitHub all releases](https://img.shields.io/github/downloads/Nmishin/leaseweb-cli/total?label=GitHub%20Total%20Downloads)

## About

`leaseweb-cli` is unofficial Leaseweb command line tool.

## Installation

### macOS

`leaseweb-cli` is available via [Homebrew](https://brew.sh/), and as a downloadable binary from [releases page](https://github.com/Nmishin/leaseweb-cli/releases/latest).

#### Homebrew

```bash
brew tap nmishin/tap
brew install leaseweb-cli
```

### Linux

`leaseweb-cli` is available via [Homebrew](https://brew.sh/), and as a downloadable binary from [releases page](https://github.com/Nmishin/leaseweb-cli/releases/latest).

#### Homebrew

```bash
brew tap nmishin/tap
brew install leaseweb-cli
```

## Usage

### Generate your API Key
You can generate your API key at the [Customer Portal](https://secure.leaseweb.com/)

### Authentication
For authentication need to export API Key from previous step, or set it as `--api-key` flag.
```bash
export LEASEWEB_API_KEY=<>
```

### Supported commands for dedicated-server
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

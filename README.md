# Leaseweb CLI

This tool interacts with Leaseweb's API, allowing users to retrieve details about dedicated servers. It includes commands like `get` to fetch server details and `get-hardware` to retrieve hardware information.

### Install
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
For authentication need to export API Key from previous step
```bash
export LEASEWEB_API_KEY=<>
```

---

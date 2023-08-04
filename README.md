<div align="center">
<h3>
    <code>addr</code>
</h3>
<br/>
Look up route origin information from the command-line
<br/>
<br/>
    <a href="https://github.com/thatmattlove/addr/actions/workflows/test.yml">
        <img alt="GitHub Workflow Status" src="https://img.shields.io/github/actions/workflow/status/thatmattlove/addr/test.yml?style=for-the-badge">
    </a>
    <a href="https://app.codecov.io/gh/thatmattlove/addr">
        <img alt="Codecov" src="https://img.shields.io/codecov/c/github/thatmattlove/addr">
    </a>
    <a href="https://github.com/thatmattlove/addr/releases">
        <img alt="GitHub release (latest SemVer)" src="https://img.shields.io/github/v/release/thatmattlove/addr?label=version&style=for-the-badge">
    </a>

</div>
<div align="center">
    <br>
    <code>addr</code> gets its information from the wonderful <a href="https://bgp.tools" target="_blank">bgp.tools</a>.
</div>

## Installation

### macOS

#### Homebrew

```console
brew tap thatmattlove/addr
brew install addr
```

### Linux

#### Debian/Ubuntu (APT)

```console
echo "deb [trusted=yes] https://repo.fury.io/thatmattlove/ /" > /etc/apt/sources.list.d/thatmattlove.fury.list
sudo apt update
sudo apt install addr
```

#### RHEL/CentOS (YUM)

```console
echo -e "[fury-thatmattlove]\nname=thatmattlove\nbaseurl=https://repo.fury.io/thatmattlove/\nenabled=1\ngpgcheck=0" > /etc/yum.repos.d/thatmattlove.fury.repo
sudo yum update
sudo yum install addr
```

### Windows

*TODO* In the meantime, download from [releases](https://github.com/thatmattlove/addr/releases)

## Usage

```console
❯ ./addr --help
addr is a tool to look up IP & ASN ownership and routing information.

Usage:
  addr [flags]
  addr [command]

Available Commands:
  asn         Look up an ASN
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  ip          Look up an IP address or prefix

Flags:
  -h, --help      help for addr
  -v, --version   version for addr

Use "addr [command] --help" for more information about a command.
```

### ASN

![](https://github.com/thatmattlove/addr/blob/main/screenshot1.png?raw=true)

### IP Address/Prefix

![](https://github.com/thatmattlove/addr/blob/main/screenshot2.png?raw=true)

![GitHub](https://img.shields.io/github/license/thatmattlove/addr?style=for-the-badge&color=black)
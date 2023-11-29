# `gox`

`gox` is a tool to manage binary golang packages.

## Features

- 🔥 Compile and manage binary packages easily
- 🧹 Single static binary, no dependencies
- 🚀 Blazingly fast
- Without polluting your home directory with Go directories
- No need of adding GOBIN to your path

## Functionality

### Install

Compiles the binary in a temporary directory, without polluting your `GOPATH`. 
Installs a simple binary to `~/.local/bin`. No fuss with GOBIN.

![install](./gifs/install.gif)

### Remove

Removes the package instantly. Any dependencies were never installed, so easy cleanup.

![remove](./gifs/remove.gif)

### Upgrade

Upgrades any given package from it's source, does everything in isolation.

![upgrade](./gifs/upgrade.gif)

### List

Lists all the installed commands along with their installed location.

![list](./gifs/upgrade.gif)

## Installation

Use the installation script or download from the releases page. 

## Getting Help

You can invoke `gox help` for help about anything. For more information about subcommands run `gox help [subcommand]`.

## Contributing

You are more than welcome to raise issues and PR in the repository

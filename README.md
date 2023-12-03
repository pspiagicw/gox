# `gox`

`gox` is a tool to manage binary golang packages.

## Features

- 📦️ Compile and manage binary packages easily
- 🔥 Single static binary, no dependencies
- ⚡️ Blazingly fast
- 🧹 Doesn't polluting your home directory or `GOPATH`

## Inspiration

Whenever you install a golang binary package (`lazygit/lazydocker/gh/micro` etc), it installs all it's dependencies into `GOPATH`.
It installs the binary into `GOBIN` which should be included in your `PATH` variable.

When you want to update the package, how to update it ?

When you want to remove it, you can remove the binary with a simple `rm` command, but what about the dependencies ?

Analogous to `pipx` and `npx`, `gox` installs the binary in isolation and symlinks it into `~/.local/bin`.
You can remove it easily and the dependencies were never in your way!

Also, it can manage your personal projects, so you can easily install and update binaries while development.

## Functionality

### Install

- Compiles the binary in a isolated environment
- Automatically installs it in the `~/.local/bin` directory.
- Provide `.` as a argument to install the project you are inside!
- Or provide a path on your system to manage the binary of your projects.

![install](./gifs/install.gif)

### Remove

- Removes the package instantly. 
- Any dependencies were never installed, so easy cleanup.

![remove](./gifs/remove.gif)

### Update

- Updates any given package from it's source.

![update](./gifs/update.gif)

### List

- Lists all the installed commands along with their installed location.

![list](./gifs/list.gif)

## Installation

You can download the binary from the [release](https://github.com/pspiagicw/gox/releases) section.
You can bootstrap `gox` from itself!

```sh
/download/location/of/gox install github.com/pspiagicw/gox@latest
```


## Getting Help

You can invoke `gox help` for help about anything. 

For more information about any subcommand run `gox help [subcommand]`.

## Tools

There are plenty of tools that are best enjoyed with `gox`. `gox` makes the installation and management of these tools a breeze.

- [gopls](https://golang.org/x/tools/gopls) (The Golang Language Server) `gox install golang.org/x/tools/gopls@latest`
- [lazygit](https://github.com/jesseduffield/lazygit) (TUI interface for git) `gox install github.com/jesseduffield/lazygit@latest`
- [lazydocker](https://github.com/jesseduffield/lazydocker) (TUI interface for docker) `gox install github.com/jesseduffield/lazydocker@latest`
- [circumflex](https://github.com/bensadeh/circumflex) (TUI interface for hackernews) `gox install github.com/bensadeh/circumflex@latest`

Shameless Plug!

- [qemantra](https://github.com/pspiagicw/qemantra) (CLI interface to QEMU) `gox install github.com/pspiagicw/qemantra@latest`
- [groom](https://github.com/pspiagicw/groom) (Build tool written in Golang) `gox install github.com/pspiagicw/groom@latest`

These are just examples to present the case of why this tool might be useful in your workflow.

## Similar projects

- [gup](https://github.com/nao1215/gup)

## Contributing

You are more than welcome to raise issues and PR in the repository


# tuner1

![Product demo](/assets/demo.gif)

A basic guitar tuner TUI with customizable templates

[![build](https://github.com/lxsavage/tuner1/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/lxsavage/tuner1/actions/workflows/build.yml)
![GitHub Release](https://img.shields.io/github/v/release/lxsavage/tuner1)

## Installation

### Scripts

#### Install

If this install script does not support your platform/architecture, the program
will have to be [manually built and installed](#manual-buildinstall).

MacOS and Linux

```sh
curl -sSL https://raw.githubusercontent.com/lxsavage/tuner1/refs/heads/main/scripts/install.sh | bash
```

Windows

```powershell
irm "https://raw.githubusercontent.com/lxsavage/tuner1/refs/heads/main/scripts/install.ps1" | iex
```

#### Uninstall

MacOS and Linux

```sh
curl -sSL https://raw.githubusercontent.com/lxsavage/tuner1/refs/heads/main/scripts/uninstall.sh | bash
```

Windows

```powershell
irm "https://raw.githubusercontent.com/lxsavage/tuner1/refs/heads/main/scripts/uninstall.ps1" | iex
```

### Manual Build/Install

In order to build this project, the Golang CLI needs to be installed and on
path. For more information on how to do this, check the
[Golang install guide](https://go.dev/doc/install).

**Note for Linux systems: the build depends on the ALSA dev library (e.g.
`libasound2-dev` on Ubuntu), which will also need to be installed through your
respective package manager.**

To build and install, use `make install`.

The program adds itself under `/usr/local/bin/` and creates a template config
under `~/.config/tuner1/standards.txt`.

Upgrading from a previous version is as simple as pulling the latest changes,
then running `make upgrade`.

Uninstallation is just `make uninstall`.

**Note: by default, this will be installed under `/usr/local/bin`. This can be
changed by adjusting the makefile `INSTALL_DIR` variable to the intended path
before running any of these make commands.**

## Usage

For most use cases, these commands will be sufficient, but additional features
can be shown with `tuner1 -h`.

- `tuner1 -ls`: List templates available
- `tuner1 -tuning +<template name>`: Launch using a template by name

## Editing templates

To change templates, edit standards.txt to add/remove K:V pairs for
templates in the format:

```plain
<template name>:<csv of scientific-notation note names from low to high>
```

This file is by default located at:

- MacOS: `~/Library/Application Support/tuner1/standards.txt`
- Linux: `~/.config/tuner1/standards.txt`
- Windows: `%APPDATA%\tuner1\standards.txt`

Afterwards, call the template with `go run . -tuning +<template name>`

---

It is also possible to test a template csv by manually calling the TUI with
the template:

```sh
tuner1 -tuning "<csv of scientific-notation note names from low to high>"
```

For example, the following is equivalent to the `+e-standard` template:

```sh
tuner1 -tuning "E2,A2,D3,G3,B3,E4"
```

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md)

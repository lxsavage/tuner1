# tuner1

![Product demo](/assets/demo.gif)

A basic guitar tuner TUI with customizable templates

[![build](https://github.com/lxsavage/tuner1/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/lxsavage/tuner1/actions/workflows/build.yml)
![GitHub Release](https://img.shields.io/github/v/release/lxsavage/tuner1)

## Installation

### Scripts

#### Install (MacOS and Linux x64)

```sh
curl -sSL https://raw.githubusercontent.com/lxsavage/tuner1/refs/heads/main/scripts/install.sh | bash
```

If this install script does not support your platform, the program will have
to be [manually built and installed](#manual-buildinstall). If you think your
platform should be included, create an issue for it and I will consider adding
it to the next version. Note that a Windows build is currently being worked on
and will be included in the next major version.

#### Uninstall

```sh
curl -sSL https://raw.githubusercontent.com/lxsavage/tuner1/refs/heads/main/scripts/uninstall.sh | bash
```

### Manual Build/Install

In order to build this project, the go CLI needs to be installed and on path.
For more information on how to do this, check the
[go install guide](https://go.dev/doc/install)

**Note for Linux systems: the build depends on `libasound2-dev`,
which will also need to be installed through your respective package
manager.**

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

- `tuner1 -ls`: List templates available
- `tuner1 -tuning +<template name>`: Launch using a template by name
- `tuner1 -tuning "<csv of scientific-notation note names from low to high>"`:
  Launch using a manually-defined CSV tuning

These can also be viewed by invoking `tuner1 -h`

## Editing templates

To change templates, edit config/standards.txt to add/remove K:V pairs for
templates in the format:

```plain
<template name>:<csv of scientific-notation note names from low to high>
```

Afterwards, call the template with `go run . -tuning +<template name>`

It is also possible to test a template csv by manually calling the TUI with
the template:

```bash
tuner1 -tuning "<csv of scientific-notation note names from low to high>"
```

## Running dev environment

1. `go get .`
2. `go run . -standards config/standards.txt -ls` (to see tuning templates)
3. `go run . -standards config/standards.txt -tuning +template_name`

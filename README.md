# tuner1

![Product demo](/assets/demo.gif)

A basic guitar tuner TUI with customizable templates

[![build](https://github.com/lxsavage/tuner1/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/lxsavage/tuner1/actions/workflows/build.yml)
![GitHub Release](https://img.shields.io/github/v/release/lxsavage/tuner1)

## Installation

### Scripts

#### Install

> [!NOTE]
> If this install script does not support your platform/architecture, the program
> will have to be [manually built and installed](#manual-buildinstall).

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

> [!IMPORTANT]
> For Linux systems: the build depends on the ALSA dev library (e.g.
> `libasound2-dev` on Ubuntu), which will also need to be installed through your
> respective package manager.

To build and install, use `make install`.

> [!NOTE]
> By default, this will be installed under `/usr/local/bin`. This can be changed
> by adjusting the makefile `INSTALL_DIR` variable to the intended path before
> running any of these make commands.

Upgrading from a previous version is as simple as pulling the latest changes,
then running `make upgrade`.

Uninstallation is just `make uninstall`.

## Usage

> [!TIP]
> Additional features can be shown with `tuner1 -h`

- `tuner1 -ls`: List templates available
- `tuner1 -tuning +<template name>`: Launch using a template by name

## Editing templates

To change templates, use `tuner1 -edit-standards` to open the standards file in
your default editor. These templates are in the format:

```plain
<template name>:<csv of scientific-notation note names from low to high>
```

There should be at most one template per-line, with empty lines ignored. After
you are done editing, call the template with `tuner1 -tuning +<template name>`.

> [!TIP]
>  The configuration file location is dependent on your OS:
>
> - MacOS: `~/Library/Application Support/tuner1/standards.txt`
> - Linux: `~/.config/tuner1/standards.txt`
> - Windows: `%APPDATA%\tuner1\standards.txt`


> [!TIP]
> It is also possible to test a template csv by manually calling the TUI with
> the template:
>
> ```sh
> tuner1 -tuning "<csv of scientific-notation note names from low to high>"
> ```
>
> For example, the following is equivalent to the `+e-standard` template:
>
> ```sh
> tuner1 -tuning "E2,A2,D3,G3,B3,E4"
> ```

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md)

# Contributing

Patches and bug reports are welcome with a few guidelines.

## Submitting a Bug Report

Bug reports will go through the issues tab of the repo. Make sure to follow the
template in place to ensure that all of the necessary information is sent to be
able to fix the bug. If possible,
[submitting a pull request](#submitting-a-pull-request) with a fix would be
helpful, but if not, I'll take a look at the issue and fix it as soon as
possible.

## Submitting a Pull Request

Before submitting a PR, ensure that the following are true about your change,
otherwise it will be rejected:

1. Created a separate branch off of the `main` branch for changes
2. Ensure that the feature/change is working correctly without any undocumented
   environmental changes (see [Running Dev Environment](#running-dev-environment))
3. Created any new tests if applicable for the new functionality
4. Ran the full test suite to ensure there were no regressions (see [Testing](#testing))

After all of these are true, then submit a PR.

## Running Dev Environment

In order to build/run this project, the Golang CLI needs to be installed and on
your path. For more information on how to do this, check the
[Golang install guide](https://go.dev/doc/install) if you don't already have
this installed.

Outside of the typical setup for running a Go project, you need to manually
specify the path of the `standards.txt` file that is used, otherwise it will try
to look in the default configuration directory based on your OS.

The main commands you will need for getting this up and running are:

```sh
# Initial module retrieval
go get .

# Testing templates
go run . -standards config/standards.txt -ls
go run . -standards config/standards.txt -tuning +template_name
```

## Testing

Tests are set up with the standard Go test runner, which can be run with

```sh
go test .
```

after initially building and running the project. This should result in an
"ok" message for the `lxsavage/tuner1` module.

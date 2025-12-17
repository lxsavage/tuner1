# tuner1

A basic guitar tuner TUI with customizable templates

## Running

There are not binary builds set up yet for this, so it has to be manually run
for now. There are plans to get this set up soon.

1. `go get .`
2. `go run . -ls` (to see tuning templates)
3. `go run . -tuning +template_name`

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
go run . -tuning "<csv of scientific-notation note names from low to high>"
```

# Packer Plugin Gridscale

[![Automatic Releaser](https://github.com/gridscale/packer-plugin-gridscale/actions/workflows/release.yml/badge.svg?branch=master)](https://github.com/gridscale/packer-plugin-gridscale/actions/workflows/release.yml)
[![GoDoc](https://godoc.org/github.com/gridscale/packer-plugin-gridscale?status.svg)](https://godoc.org/github.com/gridscale/packer-plugin-gridscale)
[![GitHub latest release](https://img.shields.io/github/release/gridscale/packer-plugin-gridscale.svg)](https://github.com/gridscale/packer-plugin-gridscale/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/gridscale/packer-plugin-gridscale)](https://goreportcard.com/report/github.com/gridscale/packer-plugin-gridscale)

The [gridscale](https://www.gridscale.io/) packer plugin can be used with HashiCorp [Packer](https://www.packer.io)
to create custom templates in [gridscale](https://www.gridscale.io/). Docs: [docs](docs).

## Maintainers

This provider plugin is maintained by the Packer team at [gridscale](https://www.gridscale.io/).

## Requirements

- [Packer](https://www.packer.io/intro/getting-started/install.html) ≥ v1.7.0
- [Go](https://golang.org/doc/install) ≥ 1.17 (to build the provider plugin)

## Installation

### Automatic installation:
Packer (≥ v1.7.0) supports a new `packer init` command allowing
automatic installation of Packer plugins. [`Plugin installation`](https://www.packer.io/docs/plugins#installing-plugins).

To install this plugin, copy and paste this code into your Packer configuration .
Then, run [`packer init`](https://www.packer.io/docs/commands/init).

```hcl
packer {
  required_plugins {
    gridscale = {
      version = ">= 0.0.1"
      source  = "github.com/gridscale/gridscale"
    }
  }
}
```

#### Manual installation
[`Plugin installation`](https://www.packer.io/docs/plugins#installing-plugins)

## Building/Developing the Provider

Build:

    $ git clone git@github.com:gridscale/packer-plugin-gridscale.git
    $ cd packer-plugin-gridscale
    $ make build

For dev:

    $ make dev


## Running Acceptance Tests
Requirements:
- Environment varialbe `GRIDSCALE_UUID` has to be set.
- Environment varialbe `GRIDSCALE_TOKEN` has to be set.

In `packer-plugin-gridscale` directory, run:
```
make acctest
```

This will run the acceptance test for `packer-plugin-gridscale`.

## Examples:

## Releasing the Provider:
- Update version in `version/version.go`
- Do `make generate` to regenerate docs and go file `config.hcl2spec.go` (if there are updates in `config.go`).
- Add entries to `CHANGELOG.md`

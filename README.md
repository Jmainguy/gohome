# gohome
[![Go Report Card](https://goreportcard.com/badge/github.com/Jmainguy/gohome)](https://goreportcard.com/badge/github.com/Jmainguy/gohome)
[![Release](https://img.shields.io/github/release/Jmainguy/gohome.svg?style=flat-square)](https://github.com/Jmainguy/gohome/releases/latest)
[![Coverage Status](https://coveralls.io/repos/github/Jmainguy/gohome/badge.svg?branch=main)](https://coveralls.io/github/Jmainguy/gohome?branch=main)

Dynamic DNS using nsd and ssh

## Usage
Run against a bind file.

```/bin/bash
gohome /tmp/soh.re
```

## PreBuilt Binaries
Grab Binaries from [The Releases Page](https://github.com/Jmainguy/gohome/releases)

## Install

### Homebrew

```/bin/bash
brew install Jmainguy/tap/gohome
```

## Build
```/bin/bash
export GO111MODULE=on
go build
```

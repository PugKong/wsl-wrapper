# WSL Wrapper

[![build](https://github.com/PugKong/wsl-wrapper/actions/workflows/ci.yml/badge.svg)](https://github.com/PugKong/wsl-wrapper/actions/workflows/ci.yml)
[![License: WTFPL](https://img.shields.io/badge/License-WTFPL-brightgreen.svg)](http://www.wtfpl.net/about/)
[![Coverage Status](https://coveralls.io/repos/github/PugKong/wsl-wrapper/badge.svg?branch=master)](https://coveralls.io/github/PugKong/wsl-wrapper?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/PugKong/wsl-wrapper)](https://goreportcard.com/report/github.com/PugKong/wsl-wrapper)
[![LoC](https://tokei.rs/b1/github/PugKong/wsl-wrapper)](https://github.com/PugKong/wsl-wrapper)
[![Release](https://img.shields.io/github/release/PugKong/wsl-wrapper.svg?style=flat-square)](https://github.com/PugKong/wsl-wrapper/releases/latest)

A tiny tool for running commands from your WSL distribution as a standalone executable. At the moment supports only the
default distro.

## Usage

Download the binary from the [Releases](https://github.com/PugKong/wsl-wrapper/releases) page and rename it to the
command you want to wrap. For example, if you want to wrap git, then rename the file to `git.exe`.

Then use it as normal Windows app, i.e. `.\git.exe add .\.github\workflows\ci.yml` will fire
`wsl git add ./.github/workflows/ci.yml` command with redirected input and output, the log file `git.log` will be placed
next to the binary.

## Why?

Now JetBrains Idea needs (or I just don't know how to set up it properly) some `docker.exe` to work with docker
installed inside WSL. That's why the wrapper is here. Hope one day I will exterminate the repo.

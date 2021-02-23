# gotest [![Build Status](https://travis-ci.com/mfroes/gotest.svg?token=onas63qBL39kCorPgzps&branch=main)](https://travis-ci.com/mfroes/gotest)

Simple http server written in Go

## Overview

This is a boilerplate written in Go for http endpoints.

* Dockerised service written in Go.
* CI Pipeline which triggers on push to master branch

## Architecture

Continuous Integration Architecture

![arch](docs/arch.png)

## Processes

List of the processes used.

| Requirement       | Tech or Process                                              |
| ----------------- | ------------------------------------------------------------ |
| Development Model | [Acceptence Test Driven Development](https://en.wikipedia.org/wiki/Acceptance_test%E2%80%93driven_development) |
| Branching Model   | [GitHub Flow](https://githubflow.github.io/) |
| Versioning        | [Semantic Versioning](https://semver.org/)                   |

## Technologies

List of the technologies used.

| Requirement                 | Tech or Process                                              |
| --------------------------- | ------------------------------------------------------------ |
| Language                    | [Go](https://golang.org/doc/) |
| Dependency Management | [go.mod](https://golang.org/doc/modules/gomod-ref) |
| Linters                     | [go lint](https://github.com/golang/lint) |
| Formatter                   | [gofmt](https://golang.org/cmd/gofmt/) |
| Testing Framework           | [go test](https://golang.org/doc/code#Testing) |
| SCM                         | [Github](https://github.com/) |
| Packaging                   | [Docker](https://www.docker.com/) |
| Pre-Commit Pipeline         | [pre-commit](https://pre-commit.com/) |
| CI Pipeline                 | [travis CI](https://travis-ci.org) |
# goperset


Golang Apache Superset API Package

This project structure is based on: [gocloak](https://github.com/Nerzal/gocloak)

For Questions either raise an issue, or come to the [gopher-slack](https://invite.slack.golangbridge.org/) into the channel [#goperset](https://gophers.slack.com/app_redirect?channel=goperset)

__SOON__: Benchmarks can be found [here](https://justmike1.github.io/goperset/dev/bench/)

## Contribution

(WIP) <https://github.com/justmike1/goperset/wiki/Contribute>

## Changelog

For release notes please consult the specific releases [here](https://github.com/justmike1/goperset/releases)


## Usage

### Installation

```shell
go get github.com/justmike1/goperset@latest
```

### Importing

```go
 import "github.com/justmike1/goperset@latest"
```

### Create New User

```go
 client := goperset.NewClient("https://mycool.superset.instance")
 ctx := context.Background()
```
 
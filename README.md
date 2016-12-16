# pubip

A simple library for getting your public IP address by several services. It's
inspired by [go-ipify](https://github.com/rdegges/go-ipify).

[![GitHub License](https://img.shields.io/badge/license-Unlicense-blue.svg)](https://raw.githubusercontent.com/chyeh/pubip/master/UNLICENSE)
[![GoDoc](https://godoc.org/github.com/chyeh/pubip?status.svg)](https://godoc.org/github.com/chyeh/pubip)
[![Build Status](https://travis-ci.org/chyeh/pubip.svg?branch=master)](https://travis-ci.org/chyeh/pubip)


## Introduction

In short, It validates the results from several services and returns the IP
address a valid one is found. If you have ever tried to deploy services in
China, you would understand what are the [fallacies of distributed computing](fallacies of distributed computing)
very much. Based on the assumption that the services your program depends on are
not always available, it's better to have more backups. That's why it's better
to get your public IP address by several services. This library gets you the
public IP address from several APIs that I found.


## Installation

To install `pubip`, simply run:

```console
$ go get -u github.com/chyeh/pubip
```

This will install the latest version of the library automatically.


## Usage

Here's a simple example:

```go
package main

import (
    "fmt"
    "github.com/chyeh/pubip"
)

func main() {
    ip, err := pubip.Get()
    if err != nil {
        fmt.Println("Couldn't get my IP address:", err)
    } else {
        fmt.Println("My IP address is:", ip)
    }
}
```

For more details, please take a look at the [GoDoc](https://godoc.org/github.com/chyeh/pubip)

## Error handling

It returns error when the followings happen:

- It fails to get at least 3 results from the services
- The results from different service are not identical


## Contributing

Just send me a PR or open an issue. It will be nice if tests are included in
your PR.

# TCP server using abstract socket

## Overview

This example shows how to use TCP server that listens on an abstract socket.

## Run the proxy

Run the proxy with the command.

```sh
go run ./main.go
```

## Test

### Run upstream HTTP server

[forio](https://github.com/fortio/fortio) is used as a upstream HTTP server here.

Upstream server:

```sh
fortio server -http-port 9000
```

### Send requests

Requests to the abstract socket using curl.

```sh
curl --abstract-unix-socket 'example' http://localhost:8000/debug
```

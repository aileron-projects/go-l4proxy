# TCP server using path socket

## Overview

This example shows how to use TCP server that listens on a path socket.

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

Requests to the path socket using curl.

```sh
curl --unix-socket '/var/run/example.sock' http://localhost:8000/debug
```

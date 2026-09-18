# TCP server with ip whitelist

## Overview

This example shows how to use TCP server with ip whitelist.

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

Requests from `127.0.0.1` will be rejected.

```sh
curl --interface 127.0.0.1 localhost:8000/debug
```

Requests from `127.0.0.2`, `127.0.0.3` will be allowed.

```sh
curl --interface 127.0.0.2 localhost:8000/debug
```

```sh
curl --interface 127.0.0.3 localhost:8000/debug
```

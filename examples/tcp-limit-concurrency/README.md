# TCP server example with limit listener

## Overview

This example shows how to limit the number of concurrent TCP connections.

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

Send requests concurrently from fortio client to the proxy server.

```sh
fortio load -t 60s -qps 100 -c 15 http://localhost:8000?delay=3s
```

- `-t` duration
- `-qps rate`
- `-c connections`

### Check concurrenrcy

Use `lsof` to check concurrency on linux.

```sh
lsof -i:8000
```

Note that the linux command "netstat" or "ss" like below do not show the correct number of connections.
See also <https://github.com/golang/go/issues/36212#issuecomment-567838193>.

- netstat -uant | grep ESTABLISHED | grep 8080 | wc
- ss -o state established "( dport = :8080 )" -np | wc

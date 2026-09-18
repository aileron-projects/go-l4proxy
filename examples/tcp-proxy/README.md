# TCP simple proxy example

## Overview

This example shows how to simply proxy tcp.

## Run the proxy

Run the proxy with the command.

```sh
go run ./main.go
```

## Test

### Use `netperf`

Upstream server:

```sh
netserver -d -D -p 9000
```

Client:

```sh
netperf -d -H localhost -p 8000 -t TCP_STREAM
```

### Use `iperf`

Upstream server:

```sh
iperf --server --port 9000
```

Client:

```sh
iperf -c localhost --port 8000
```

### Use `iperf3`

Upstream server:

```sh
iperf3 --server --port 9000
```

Client:

```sh
iperf3 -c localhost --port 8000
```

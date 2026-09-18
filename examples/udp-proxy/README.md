# UDP simple proxy example

## Overview

This example shows how to simply proxy udp.

## Run the proxy

Run the proxy with the command.

```sh
go run ./main.go
```

## Test

### Use `socat`

Upstream server:

```sh
socat - UDP-LISTEN:9000
```

Client:

```sh
echo -n 'hello' | socat - UDP:127.0.0.1:8000
```

or interactively:

```sh
socat - UDP:127.0.0.1:8000
```

### Use `nc`

Upstream server:

```sh
nc -u -l 9000
```

Client:

```sh
echo -n "hello" | nc -u 127.0.0.1 8000
```

or interactively:

```sh
nc -u 127.0.0.1 8000
```

### Use `iperf`

Upstream server:

```sh
iperf --server --port 9000 --udp
```

Client:

```sh
iperf -c localhost --port 8000 --udp
```

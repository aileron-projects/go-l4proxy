# UDP simple server example

## Overview

This example shows how to simply run a UDP server.

## Run the server

Run the UDP server with the command.

```sh
go run ./main.go
```

## Test

### Use `nc`

Send an udp packet to the server.

```sh
echo -n 'hello' | nc -u 127.0.0.1 8000
```

or interactively:

```sh
nc -u 127.0.0.1 8000
```

### Use `socat`

Send an udp packet to the server.

```sh
echo -n 'hello' | socat - UDP:127.0.0.1:8000
```

or interactively:

```sh
socat - UDP:127.0.0.1:8000
```

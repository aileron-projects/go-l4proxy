# TCP simple server example

## Overview

This example shows how to simply run a TCP server.

## Run the server

Run the TCP server with the command.

```sh
go run ./main.go
```

## Test

### Use `nc`

Send tcp messages to the server.

```sh
echo -n 'hello' | nc 127.0.0.1 8000
```

or interactively:

```sh
nc 127.0.0.1 8000
```

### Use `socat`

Send tcp messages to the server.

```sh
echo -n 'hello' | socat - TCP:127.0.0.1:8000
```

or interactively:

```sh
socat - TCP:127.0.0.1:8000
```

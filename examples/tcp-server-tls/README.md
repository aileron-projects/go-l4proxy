# TCP TLS server

## Overview

This example shows how to run a TLS server.

## Run the server

Run the server with the command.

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

### Send TLS requests

TLS requests will be reach to the fortio server.

```sh
curl --cert cert.pem --key key.pem -k https://localhost:8000/debug
```

Non-TLS will be rejected.

```sh
curl http://localhost:8000/debug
```

## Test cert files

- `cert.pem` is a self-signed server certification.
- `key.pem` is a signing key for the cert.

Files are generated with [openssl](https://docs.openssl.org/master/man1/openssl-req/).

```bash
openssl req -newkey rsa:4096 -nodes -keyout key.pem -x509 -days 36500 -out cert.pem -addext 'subjectAltName = DNS:localhost,DNS:*.sock,IP:127.0.0.1' -subj '/CN=127.0.0.1'
```

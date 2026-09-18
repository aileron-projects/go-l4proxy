<!-- markdownlint-disable MD033 MD041 -->

<div align="center">

[![Release](https://img.shields.io/github/v/release/aileron-projects/go-l4proxy?sort=semver)](https://github.com/aileron-projects/go-l4proxy/releases)
[![Reference](https://pkg.go.dev/badge/github.com/aileron-projects/go-l4proxy.svg)](https://pkg.go.dev/github.com/aileron-projects/go-l4proxy)
[![DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/aileron-projects/go-l4proxy)
[![Test](https://github.com/aileron-projects/go-l4proxy/actions/workflows/test.yaml/badge.svg)](https://github.com/aileron-projects/go-l4proxy/actions/workflows/test.yaml)

[![Insights](https://badgen.net/badge/Insights/open%2Fsource%2Finsights/cyan)](https://deps.dev/go/github.com%2Faileron-projects%2Fgo-l4proxy)
[![Insights](https://badgen.net/badge/Insights/OSS%2FInsight/orange)](https://ossinsight.io/analyze/aileron-projects/go-l4proxy)

</div>

# go-l4proxy

**Layer4, TCP and UDP, proxy servers and library for Go.**

## Features

- TCP server
- TCP proxy
- UDP server
- UDP proxy
- High performance

## Usages

### TCP Server and proxy

```go
svr := &tcp.Server{
    Addr:    ":8000",
    Handler: tcp.NewProxy("localhost:9000"),
}

if err := svr.ListenAndServe(); err != nil && err != tcp.ErrServerClosed {
    panic(err)
}
```

### UDP Server and proxy

```go
svr := &udp.Server{
    Addr:    ":8000",
    Handler: udp.NewProxy("localhost:9000"),
}

if err := svr.ListenAndServe(); err != nil && err != udp.ErrServerClosed {
    panic(err)
}
```

## Docs & Examples

- GoDoc: <https://pkg.go.dev/github.com/aileron-projects/go-l4proxy>
- Examples:
  - TCP proxy: [examples/tcp-proxy](./examples/tcp-proxy)
  - TCP proxy with upstream TLS: [examples/tcp-proxy-tls](./examples/tcp-proxy-tls)
  - TCP server: [examples/tcp-server-simple](./examples/tcp-server-simple)
  - TCP server graceful shutdown: [examples/tcp-server-runner](./examples/tcp-server-runner)
  - TCP TLS server: [examples/tcp-server-tls](./examples/tcp-server-tls)
  - TCP server with ip whitelist: [examples/tcp-server-whitelist](./examples/tcp-server-whitelist)
  - TCP server with max connections: [examples/tcp-limit-concurrency](./examples/tcp-limit-concurrency)
  - TCP server listens on a unix abstract socket: [examples/tcp-socket-abstract](./examples/tcp-socket-abstract)
  - TCP server listens on a unix path socket: [examples/tcp-socket-path](./examples/tcp-socket-path)
  - UDP proxy: [examples/udp-proxy](./examples/udp-proxy)
  - UDP server: [examples/udp-server-simple](./examples/udp-server-simple)
  - UDP server graceful shutdown: [examples/udp-server-runner](./examples/udp-server-runner)
  - UDP server with ip whitelist: [examples/udp-server-whitelist](./examples/udp-server-whitelist)
  - UDP server listens on a unix abstract socket: [examples/udp-socket-abstract](./examples/udp-socket-abstract)
  - UDP server listens on a unix path socket: [examples/udp-socket-path](./examples/udp-socket-path)

## References

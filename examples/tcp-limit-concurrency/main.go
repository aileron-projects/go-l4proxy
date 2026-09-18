package main

import (
	"context"
	"log"
	"net"

	"github.com/aileron-projects/go-l4proxy/tcp"
	"golang.org/x/net/netutil"
)

func main() {
	svr := &tcp.Server{
		Addr:    "", // This is not used when we call [Server.Serve].
		Handler: tcp.NewProxy("localhost:9000"),
	}

	var lc net.ListenConfig
	ln, _ := lc.Listen(context.Background(), "tcp", ":8000") // Create a new TCP listener.
	ln = netutil.LimitListener(ln, 10)                       // Limit concurrency.

	log.Println("starting tcp server at " + ln.Addr().String())
	if err := svr.Serve(ln); err != nil && err != tcp.ErrServerClosed {
		panic(err)
	}
}

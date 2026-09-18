package main

import (
	"context"
	"log"
	"net"

	"github.com/aileron-projects/go-ipfilter"
	"github.com/aileron-projects/go-l4proxy/tcp"
)

func main() {
	svr := &tcp.Server{
		Addr:    "",
		Handler: tcp.NewProxy("localhost:9000"),
	}

	var lc net.ListenConfig
	ln, _ := lc.Listen(context.Background(), "tcp", ":8000")
	ln, _ = ipfilter.WhitelistListener(ln, "127.0.0.2", "127.0.0.3") // Whitelist
	// ln, _ = ipfilter.BlacklistListener(ln, "127.0.0.1") // Blacklist

	log.Println("starting tcp server at " + ln.Addr().String())
	if err := svr.Serve(ln); err != nil && err != tcp.ErrServerClosed {
		panic(err)
	}
}

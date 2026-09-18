package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/aileron-projects/go-ipfilter"
	"github.com/aileron-projects/go-l4proxy/udp"
)

func main() {
	svr := &udp.Server{
		Addr:    "",
		Handler: udp.HandlerFunc(handleConn),
	}

	var lc net.ListenConfig
	p, err := lc.ListenPacket(context.Background(), "udp", ":8080")
	if err != nil {
		panic(err)
	}
	wl := ipfilter.NewWhitelist()
	_ = wl.Allow("127.0.0.1/32", "::1/128")
	pc := &WhitelistPacketConn{PacketConn: p, wl: wl}

	log.Println("starting udp server at " + pc.LocalAddr().String())
	if err := svr.Serve(pc); err != nil && err != udp.ErrServerClosed {
		panic(err)
	}
}

// handleConn reads and prints UDP packets received from the conn.
func handleConn(ctx context.Context, conn udp.Conn) {
	buf := make([]byte, 1<<10)
	for {
		n, err := conn.Read(buf)
		fmt.Println(string(buf[:n]))
		if err != nil && err != io.EOF {
			if errors.Is(err, net.ErrClosed) {
				return
			}
			panic(err)
		}
	}
}

type WhitelistPacketConn struct {
	net.PacketConn
	wl *ipfilter.Whitelist
}

func (c *WhitelistPacketConn) ReadFrom(p []byte) (n int, addr net.Addr, err error) {
	n, addr, err = c.PacketConn.ReadFrom(p)
	if err == nil {
		host, _, err := net.SplitHostPort(addr.String())
		if err != nil {
			host = addr.String() // Fallback
		}
		if !c.wl.Allowed(host) {
			return n, addr, udp.ErrSkipHandler // Return udp.ErrSkipHandler.
		}
	}
	return n, addr, err
}

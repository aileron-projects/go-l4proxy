package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/aileron-projects/go-l4proxy/udp"
)

func main() {
	svr := &udp.Server{
		Addr:    "unixgram://@example",
		Handler: udp.HandlerFunc(handleConn),
	}
	log.Println("starting udp server at " + svr.Addr)
	if err := svr.ListenAndServe(); err != nil && err != udp.ErrServerClosed {
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

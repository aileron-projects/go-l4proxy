package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/aileron-projects/go-l4proxy/tcp"
)

func main() {
	svr := &tcp.Server{
		Addr:    ":8000",
		Handler: tcp.HandlerFunc(handleConn),
	}
	log.Println("starting tcp server at " + svr.Addr)
	if err := svr.ListenAndServe(); err != nil && err != tcp.ErrServerClosed {
		panic(err)
	}
}

// handleConn reads and prints TCP data received from the conn.
func handleConn(ctx context.Context, conn net.Conn) {
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

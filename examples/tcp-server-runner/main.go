package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/aileron-projects/go-l4proxy"
	"github.com/aileron-projects/go-l4proxy/tcp"
)

func main() {
	proxy := tcp.NewProxy("localhost:9000")
	svr := &tcp.Server{
		Addr:    ":8000",
		Handler: proxy,
	}

	runner := &l4proxy.ServerRunner{
		Serve:           svr.ListenAndServe,
		Shutdown:        svr.Shutdown,
		Close:           svr.Close,
		ShutdownTimeout: 10 * time.Second,
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	log.Println("starting tcp server at " + svr.Addr)
	if err := runner.Run(ctx); err != nil {
		panic(err)
	}
}

package main

import (
	"log"

	"github.com/aileron-projects/go-l4proxy/tcp"
)

func main() {
	svr := &tcp.Server{
		Addr:    ":8000",
		Handler: tcp.NewProxy("localhost:9000"),
	}

	log.Println("starting tcp proxy server at " + svr.Addr)
	if err := svr.ListenAndServe(); err != nil && err != tcp.ErrServerClosed {
		panic(err)
	}
}

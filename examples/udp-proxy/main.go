package main

import (
	"log"

	"github.com/aileron-projects/go-l4proxy/udp"
)

func main() {
	svr := &udp.Server{
		Addr:    ":8000",
		Handler: udp.NewProxy("localhost:9000"),
	}

	log.Println("starting udp proxy server at " + svr.Addr)
	if err := svr.ListenAndServe(); err != nil {
		panic(err)
	}
}

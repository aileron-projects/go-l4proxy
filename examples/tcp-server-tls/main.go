package main

import (
	"log"

	"github.com/aileron-projects/go-l4proxy/tcp"
)

func main() {
	svr := &tcp.Server{
		Addr:      ":8000",
		Handler:   tcp.NewProxy("localhost:9000"),
		TLSConfig: nil,
	}

	log.Println("starting tcp server at " + svr.Addr)
	if err := svr.ListenAndServeTLS("cert.pem", "key.pem"); err != nil && err != tcp.ErrServerClosed {
		panic(err)
	}
}

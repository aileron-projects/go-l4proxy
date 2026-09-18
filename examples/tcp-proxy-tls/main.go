package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"log"
	"net"
	"os"

	"github.com/aileron-projects/go-l4proxy/tcp"
)

func main() {
	pem, _ := os.ReadFile("cert.pem")
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(pem)

	tlsDialer := &tls.Dialer{
		Config: &tls.Config{RootCAs: pool},
	}
	svr := &tcp.Server{
		Addr: ":8000",
		Handler: &tcp.Proxy{
			Dial: func(ctx context.Context, dc net.Conn) (uc net.Conn, err error) {
				return tlsDialer.DialContext(context.Background(), "tcp", "localhost:9000")
			},
		},
	}

	log.Println("starting tcp proxy server at " + svr.Addr)
	if err := svr.ListenAndServe(); err != nil {
		panic(err)
	}
}

package main

import (
	"context"
	"log"

	"github.com/aliamerj/aircup/relay/internal"
	"github.com/aliamerj/aircup/relay/relay"
	"github.com/quic-go/quic-go"
)

func main() {
	tlsConfig, err := internal.GenerateTLSConfig()
	if err != nil {
		panic(err)
	}

	listener, err := quic.ListenAddr(
		":4242",
		tlsConfig,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	r := relay.New(listener)
	if err := r.Run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

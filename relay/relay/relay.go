package relay

import (
	"context"
	"log/slog"

	"github.com/quic-go/quic-go"
)

type Relay struct {
	listener *quic.Listener
	registry *Registry
}

func New(listener *quic.Listener) *Relay {
	return &Relay{
		listener: listener,
		registry: NewRegistry(),
	}
}

func (r *Relay) Run(ctx context.Context) error {
	slog.Info("relay listening", "addr", ":4242")

	for {
		conn, err := r.listener.Accept(ctx)
		if err != nil {
			return err
		}

		go r.handleConn(conn)
	}
}

func (r *Relay) handleConn(conn *quic.Conn) {
	slog.Info(
		"connection accepted",
		"addr",
		conn.RemoteAddr(),
	)

	for {
		stream, err := conn.AcceptStream(context.Background())
		if err != nil {
			slog.Error(err.Error())
			return
		}

		go r.handleStream(conn, stream)
	}
}

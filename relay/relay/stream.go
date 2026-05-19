package relay

import (
	"log/slog"

	"github.com/aliamerj/aircup/shared/protocol"
	"github.com/quic-go/quic-go"
)

func (r *Relay) handleStream(
	conn *quic.Conn,
	stream *quic.Stream,
) {

	hello, err := protocol.DecodeMessage[protocol.Hello](
		stream,
		protocol.TypeHello,
	)
	if err != nil {
		return
	}

	session := &Session{
		Conn:  conn,
		Hello: *hello,
	}

	r.registry.Add(hello.NodeID, session)

	slog.Info(
		"node connected",
		"id",
		hello.NodeID,
		"name",
		hello.Name,
	)

	<-conn.Context().Done()

	r.registry.Remove(hello.NodeID)

	slog.Info(
		"node disconnected",
		"id",
		hello.NodeID,
	)
}

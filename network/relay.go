package network

import (
	"context"
	"crypto/tls"
	"log/slog"

	"github.com/aliamerj/aircup/config"
	"github.com/aliamerj/aircup/shared/protocol"
	"github.com/quic-go/quic-go"
)

type RelayClient struct {
	Conn *quic.Conn
}

func ConnectRelay(ctx context.Context, addr string) (*RelayClient, error) {
	conn, err := quic.DialAddr(ctx, addr, &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"aircup-relay"},
	}, nil)
	if err != nil {
		return nil, err
	}
	slog.Info("connected to relay", "addr", addr)

	return &RelayClient{
		Conn: conn,
	}, nil
}

func (r *RelayClient) SendHello(ctx context.Context, cfg *config.Config) error {
	stream, err := r.Conn.OpenStreamSync(ctx)
	if err != nil {
		return err
	}

	data, err := protocol.EncodeMessage(
		protocol.TypeHello,
		protocol.Hello{
			NodeID: cfg.NodeID,
			Name:   cfg.Name,
		},
	)
	if err != nil {
		return err
	}
	_, errWrite := stream.Write(data)

	return errWrite

}

func (r *RelayClient) Close() {
	if err := r.Conn.CloseWithError(0, ""); err != nil {
		slog.Error(err.Error())
	}
}

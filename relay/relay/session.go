package relay

import (
	"github.com/aliamerj/aircup/shared/protocol"
	"github.com/quic-go/quic-go"
)

type Session struct {
	Conn  *quic.Conn
	Hello protocol.Hello
}

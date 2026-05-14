package network

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/grandcat/zeroconf"
)

type Advertiser struct {
	server *zeroconf.Server
}

func StartMDNS(addr string) (*Advertiser, error) {
	port, err := strconv.Atoi(strings.Split(addr, ":")[1])
	if err != nil {
		return nil, err
	}

	s, err := zeroconf.Register(
		"aircup",
		"_aircup._tcp",
		"local.",
		port,
		[]string{
			"version=1",
			"protocol=http",
		},
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("mdns register: %w", err)
	}

	return &Advertiser{
		server: s,
	}, nil
}

func (a *Advertiser) Close() {
	if a.server != nil {
		a.server.Shutdown()
	}
}

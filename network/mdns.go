package network

import (
	"context"
	"fmt"
	"net"
	"strconv"

	"github.com/grandcat/zeroconf"
)

const Instance = "Aircup"
const Service = "_aircup._tcp"

type Advertiser struct {
	server *zeroconf.Server
}

func StartMDNS(addr string) (*Advertiser, error) {
	_, portStr, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, err
	}

	server, err := zeroconf.Register(
		Instance,
		Service,
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
		server: server,
	}, nil
}

func (a *Advertiser) Close() {
	if a.server != nil {
		a.server.Shutdown()
	}
}

func ShareURL(ctx context.Context) (string, error) {
	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		return "", err
	}
	entry := make(chan *zeroconf.ServiceEntry)
	err = resolver.Browse(ctx, Service, "local.", entry)
	if err != nil {
		return "", err
	}

	for {
		select {
		case e := <-entry:
			if e == nil {
				continue
			}

			if e.Instance != Instance {
				continue
			}

			if len(e.AddrIPv4) == 0 {
				continue
			}

			return fmt.Sprintf(
				"http://%s:%d",
				e.AddrIPv4[0].String(),
				e.Port,
			), nil

		case <-ctx.Done():
			return "", ctx.Err()
		}
	}
}

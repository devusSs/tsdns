package main

import (
	"context"
	"net"
)

const (
	ts3SRVService = "ts3"
	ts3SRVProto   = "udp"
)

func lookupTS3SRV(ctx context.Context, host string, proto string) ([]*net.SRV, error) {
	if ctx == nil {
		return nil, errNilContext
	}

	if host == "" {
		return nil, errEmptyHost
	}

	return lookupSRV(ctx, ts3SRVService, proto, host)
}

const (
	tsDNSSRVService = "tsdns"
	tsDNSSRVProto   = "tcp"
)

func lookupTSDNSSRV(ctx context.Context, host string, proto string) ([]*net.SRV, error) {
	if ctx == nil {
		return nil, errNilContext
	}

	if host == "" {
		return nil, errEmptyHost
	}

	return lookupSRV(ctx, tsDNSSRVService, proto, host)
}

func lookupSRV(ctx context.Context, service string, proto string, host string) ([]*net.SRV, error) {
	_, addrs, err := resolver.LookupSRV(ctx, service, proto, host)
	return addrs, err
}

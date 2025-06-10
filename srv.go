package main

import (
	"context"
	"net"
)

const ts3SRVService = "ts3"

var ts3SRVProto = "udp"

func lookupTS3SRV(ctx context.Context, host string) ([]*net.SRV, error) {
	if ctx == nil {
		return nil, errNilContext
	}

	if host == "" {
		return nil, errEmptyHost
	}

	return lookupSRV(ctx, ts3SRVService, ts3SRVProto, host)
}

const tsDNSSRVService = "tsdns"

var tsDNSSRVProto = "tcp"

func lookupTSDNSSRV(ctx context.Context, host string) ([]*net.SRV, error) {
	if ctx == nil {
		return nil, errNilContext
	}

	if host == "" {
		return nil, errEmptyHost
	}

	return lookupSRV(ctx, tsDNSSRVService, tsDNSSRVProto, host)
}

func lookupSRV(ctx context.Context, service string, proto string, host string) ([]*net.SRV, error) {
	_, addrs, err := resolver.LookupSRV(ctx, service, proto, host)
	return addrs, err
}

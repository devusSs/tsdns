package main

import (
	"context"
)

func lookupDNS(ctx context.Context, host string) ([]string, error) {
	if ctx == nil {
		return nil, errNilContext
	}

	if host == "" {
		return nil, errEmptyHost
	}

	return resolver.LookupHost(ctx, host)
}

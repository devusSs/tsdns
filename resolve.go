package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"strings"
)

func resolve(host, service, proto string) error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	switch service {
	case "dns":
		return resolveDNS(ctx, host, service)

	case "ts3":
		return resolveTS3(ctx, host, service, proto)

	case "tsdns":
		return resolveTSDNS(ctx, host, service, proto)

	case "nick":
		return resolveNick(ctx, host, service, proto)

	default:
		return fmt.Errorf("unsupported service: %s", service)
	}
}

func resolveDNS(ctx context.Context, host string, service string) error {
	slog.Info("dns lookup started",
		slog.String("host", host),
		slog.String("service", service))

	addrs, err := lookupDNS(ctx, host)
	if err != nil {
		return fmt.Errorf("dns lookup failed: %w", err)
	}

	if len(addrs) == 0 {
		return fmt.Errorf("no addresses found for host: %s", host)
	}

	slog.Debug("dns lookup successful",
		slog.String("host", host),
		slog.Any("addresses", addrs),
	)

	slog.Info("dns lookup result",
		slog.String("host", host),
		slog.Any("addr", addrs[0]),
	)

	return nil
}

//nolint:dupl // This is not a duplicate, stop annoying me.
func resolveTS3(ctx context.Context, host, service, proto string) error {
	if proto == "" {
		proto = ts3SRVProto
		slog.Debug("using default protocol for ts3 srv lookup", slog.String("protocol", proto))
	}

	slog.Info("ts3 srv lookup started",
		slog.String("host", host),
		slog.String("service", service),
		slog.String("protocol", proto))

	addrs, err := lookupTS3SRV(ctx, host, proto)
	if err != nil {
		return fmt.Errorf("ts3 srv lookup failed: %w", err)
	}

	if len(addrs) == 0 {
		return fmt.Errorf("no SRV records found for host: %s", host)
	}

	slog.Debug("ts3 srv lookup successful",
		slog.String("host", host),
		slog.Any("srv_records", addrs),
	)

	slog.Info(
		"ts3 srv lookup result",
		slog.String("host", host),
		slog.String(
			"resolved",
			fmt.Sprintf("%s:%d", strings.TrimSuffix(addrs[0].Target, "."), addrs[0].Port),
		),
	)

	return nil
}

//nolint:dupl // This is not a duplicate, stop annoying me.
func resolveTSDNS(ctx context.Context, host, service, proto string) error {
	if proto == "" {
		proto = tsDNSSRVProto
		slog.Debug("using default protocol for tsdns srv lookup", slog.String("protocol", proto))
	}

	slog.Info("tsdns srv lookup started",
		slog.String("host", host),
		slog.String("service", service),
		slog.String("protocol", proto))

	addrs, err := lookupTSDNSSRV(ctx, host, proto)
	if err != nil {
		return fmt.Errorf("tsdns srv lookup failed: %w", err)
	}

	if len(addrs) == 0 {
		return fmt.Errorf("no srv records found for host: %s", host)
	}

	slog.Debug("tsdns srv lookup successful",
		slog.String("host", host),
		slog.Any("srv_records", addrs),
	)

	slog.Info(
		"tsdns srv lookup result",
		slog.String("host", host),
		slog.String(
			"resolved",
			fmt.Sprintf("%s:%d", strings.TrimSuffix(addrs[0].Target, "."), addrs[0].Port),
		),
	)

	return nil
}

func resolveNick(ctx context.Context, host string, service string, proto string) error {
	if proto == "" {
		proto = ts3SRVProto
		slog.Debug("using default protocol for nick srv lookup", slog.String("protocol", proto))
	}

	slog.Info("nick lookup started",
		slog.String("nick", host),
		slog.String("service", service),
	)

	addr, err := lookupNick(ctx, host)
	if err != nil {
		return fmt.Errorf("nick lookup failed: %w", err)
	}

	if addr == "" {
		return fmt.Errorf("no address found for nick: %s", host)
	}

	slog.Debug("nick lookup successful",
		slog.String("nick", host),
		slog.String("address", addr),
	)

	addrs, err := lookupTS3SRV(ctx, addr, proto)
	if err != nil {
		return fmt.Errorf("ts3 srv lookup for nick failed: %w", err)
	}

	if len(addrs) == 0 {
		return fmt.Errorf("no srv records found for nick: %s", host)
	}

	slog.Debug("ts3 srv lookup for nick successful",
		slog.String("nick", host),
		slog.Any("srv_records", addrs),
	)

	slog.Info(
		"nick lookup result",
		slog.String("nick", host),
		slog.String(
			"resolved",
			fmt.Sprintf("%s:%d", strings.TrimSuffix(addrs[0].Target, "."), addrs[0].Port),
		),
	)

	return nil
}

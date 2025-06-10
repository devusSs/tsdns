package main

import (
	"flag"
	"log/slog"
	"net"
	"os"
	"strings"
	"time"
)

var (
	start    = time.Now()
	resolver = &net.Resolver{}
)

func main() {
	hostFlag := flag.String("host", "", "the host to resolve to an ip and port")
	serviceFlag := flag.String(
		"service",
		"ts3",
		"the service to use to resolve the host (dns, ts3, tsdns or nick)",
	)
	protoFlag := flag.String(
		"protocol",
		"",
		"the protocol to use for SRV (ts3, tsdns) lookup (udp or tcp)",
	)
	flag.Parse()

	setupLog()

	slog.Debug(
		"app start",
		slog.String("build", getBuild().String()),
		slog.String("build_json", getBuild().json()),
	)

	slog.Debug("parsed flags",
		slog.String("host", *hostFlag),
		slog.String("service", *serviceFlag),
		slog.String("protocol", *protoFlag),
	)

	host := strings.ToLower(strings.TrimSpace(*hostFlag))
	service := strings.ToLower(strings.TrimSpace(*serviceFlag))
	proto := strings.ToLower(strings.TrimSpace(*protoFlag))

	slog.Debug("normalized flags",
		slog.String("host", host),
		slog.String("service", service),
		slog.String("protocol", proto),
	)

	err := checkHost(host)
	if err != nil {
		slog.Error("invalid host", slog.Any("err", err))
		os.Exit(1)
	}

	slog.Debug("host check passed", slog.String("host", host))

	err = checkService(service)
	if err != nil {
		slog.Error("invalid service", slog.Any("err", err))
		os.Exit(1)
	}

	slog.Debug("service check passed", slog.String("service", service))

	err = checkProto(proto)
	if err != nil {
		slog.Error("invalid protocol", slog.Any("err", err))
		os.Exit(1)
	}

	slog.Debug("protocol check passed", slog.String("protocol", proto))

	err = resolve(host, service, proto)
	if err != nil {
		slog.Error("lookup failed", slog.Any("err", err))
		os.Exit(1)
	}

	slog.Debug("app exit", slog.Duration("took", time.Since(start)))
}

func setupLog() {
	level := os.Getenv("TSDNS_LOG_LEVEL")
	if level == "" {
		level = "info"
	}

	var logLevel slog.Level
	switch strings.ToLower(level) {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}

	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel})
	slog.SetDefault(slog.New(handler))
}

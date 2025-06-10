package main

import (
	"errors"
)

var (
	errNilContext   = errors.New("context cannot be nil")
	errEmptyHost    = errors.New("host cannot be empty")
	errEmptyService = errors.New("service cannot be empty")
)

func checkHost(host string) error {
	if host == "" {
		return errEmptyHost
	}

	return nil
}

func checkService(service string) error {
	if service == "" {
		return errEmptyService
	}

	switch service {
	case "dns", "ts3", "tsdns", "nick":
		return nil
	default:
		return errors.New("unsupported service: " + service)
	}
}

func checkProto(proto string) error {
	if proto == "" {
		return nil
	}

	switch proto {
	case "udp", "tcp":
		return nil
	default:
		return errors.New("unsupported protocol: " + proto)
	}
}

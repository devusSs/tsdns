package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

const nickResolveURL = "https://named.myteamspeak.com/lookup?name="

func lookupNick(ctx context.Context, nick string) (string, error) {
	if ctx == nil {
		return "", errNilContext
	}

	if nick == "" {
		return "", errEmptyHost
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, nickResolveURL+nick, nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if len(b) == 0 {
		return "", fmt.Errorf("no data returned for nick: %s", nick)
	}

	return string(b), nil
}

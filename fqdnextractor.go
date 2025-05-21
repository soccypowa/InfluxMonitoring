package main

import (
	"fmt"
	"net"
	"net/url"
	"strings"
)

func extractFQDN(rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("failed parsing URL: %v", err)
	}

	// handeling missing schemes
	if u.Host == "" {
		u, err = url.Parse("https://" + rawURL)
		if err != nil {
			return "", fmt.Errorf("failed parsing URL after adding sheme: %v", err)
		}
	}

	// split host:port if present
	host := u.Host
	if strings.Contains(host, ":") {
		host, _, err = net.SplitHostPort(host)
		if err != nil {
			return "", fmt.Errorf("failed to split host and port: %v", err)
		}
	}

	// Remove trailing dots
	host = strings.Trim(host, ".")

	return host, nil
}

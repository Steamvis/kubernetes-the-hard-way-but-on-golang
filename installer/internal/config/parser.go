package config

import (
	"errors"
	"net"
)

func parseIpFromString(ipStr string) (net.IP, error) {
	ip := net.ParseIP(ipStr)

	if ip != nil {
		return ip, nil
	}

	return nil, errors.New("invalid IP address")
}

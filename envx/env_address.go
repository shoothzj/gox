package envx

import (
	"github.com/shoothzj/gox/netx"
	"net"
	"os"
	"strconv"
	"strings"
)

// GetEnvAddressList reads a comma-separated list of addresses from an environment variable
// and returns a slice of Address structs. If the environment variable is not set or
// contains invalid addresses, it returns an empty slice.
func GetEnvAddressList(key string) []netx.Address {
	aux := os.Getenv(key)
	if aux == "" {
		return []netx.Address{}
	}

	addressStrings := strings.Split(aux, ",")
	var addresses []netx.Address
	for _, addressStr := range addressStrings {
		address := parseAddress(addressStr)
		if address != nil {
			addresses = append(addresses, *address)
		}
	}

	return addresses
}

// GetEnvAddress reads a single address from an environment variable
// and returns an Address struct. If the environment variable is not set or
// contains an invalid address, it returns nil.
func GetEnvAddress(key string) *netx.Address {
	aux := os.Getenv(key)
	if aux == "" {
		return nil
	}

	return parseAddress(aux)
}

// parseAddress parses a string in the format "host:port" into an Address struct.
// If the format is invalid or the port is not a valid integer, it returns nil.
func parseAddress(addressStr string) *netx.Address {
	host, portStr, err := net.SplitHostPort(addressStr)
	if err != nil {
		return nil
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil
	}

	return &netx.Address{Host: host, Port: port}
}

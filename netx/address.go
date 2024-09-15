package netx

import (
	"crypto/tls"
	"net"
	"strconv"
)

type Address struct {
	// Host domain name or ipv4, ipv6 address
	Host string
	// Port service port
	Port int
}

func (addr Address) Addr() string {
	return net.JoinHostPort(addr.Host, strconv.Itoa(addr.Port))
}

func Dial(addr Address, tlsConfig *tls.Config) (net.Conn, error) {
	if tlsConfig == nil {
		return net.Dial("tcp", addr.Addr())
	} else {
		return tls.Dial("tcp", addr.Addr(), tlsConfig)
	}
}

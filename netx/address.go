package netx

import (
	"net"
	"strconv"
)

type Address struct {
	Host string
	Port int
}

func (addr Address) Addr() string {
	return net.JoinHostPort(addr.Host, strconv.Itoa(addr.Port))
}

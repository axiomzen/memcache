package memcache

import (
	"net"
)

type Addr struct {
	net.Addr
	s string
	n string
	u string
	p string
}

func (a *Addr) String() string {
	return a.s
}

func NewAddr(addr net.Addr, username, password string) *Addr {
	return &Addr{
		Addr: addr,
		s:    addr.String(),
		n:    addr.Network(),
		u:    username,
		p:    password,
	}
}

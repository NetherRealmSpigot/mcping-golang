package protocols

import (
	"context"
	"net"
)

func Resolve(host string) ([]net.IP, error) {
	return net.DefaultResolver.LookupIP(context.Background(), "ip4", host)
}

func ResolveFromSRV(host string) ([]net.IP, uint16) {
	_, srvs, _ := net.LookupSRV("minecraft", "tcp", host)
	if len(srvs) > 0 {
		for _, srv := range srvs {
			srvips, err := Resolve(srv.Target)
			if err != nil {
				continue
			}
			return srvips, srv.Port
		}
	}
	return nil, 0
}

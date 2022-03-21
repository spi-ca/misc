package networking

import (
	"context"
	"github.com/pkg/errors"
	"log"
	"math/rand"
	"net"
	"strconv"
)

// RoundrobinDialer is a wrapper for the DialContext function.
type (
	RoundrobinDialer struct {
		Dialer       *net.Dialer
		FallbackHost net.TCPAddr
	}
)

// DialContext is a connector method with Fail-Over approach.
func (d *RoundrobinDialer) DialContext(ctx context.Context, network string, hosts ...string) (net.Conn, error) {
	list := rand.Perm(len(hosts))
	for _, idx := range list {
		host := hosts[idx]
		for _, resolvedIP := range d.resolveHost(ctx, network, host) {
			var (
				conn net.Conn
				err  error
			)
			if conn, err = d.Dialer.DialContext(ctx, resolvedIP.Network(), resolvedIP.AddrPort().String()); err != nil {
				log.Printf("failed to connected to %s => %s : %s", host, resolvedIP.String(), err)
			} else {
				return conn, nil
			}
		}
	}

	if !d.FallbackHost.IP.IsUnspecified() {
		log.Printf("attempting to connect fallback address(%s)\n", d.FallbackHost.String())
		return d.Dialer.DialContext(ctx, d.FallbackHost.Network(), d.FallbackHost.AddrPort().String())
	}
	return nil, errors.New("name resolve failure")
}

// resolveHost is actually resolve network addresses from a given hostname.
func (d *RoundrobinDialer) resolveHost(ctx context.Context, network, connectAddr string) (addrs []net.TCPAddr) {
	log.Printf("attempting to connect %s", connectAddr)
	var (
		host        string
		portString  string
		resolvedIPs []net.IP
		tcpPort     int
		err         error
	)
	if host, portString, err = net.SplitHostPort(connectAddr); err != nil {
		log.Printf("invalid connection string format: %s", err)
		return
	} else if port, _ := strconv.ParseInt(portString, 10, 32); port < 0 || port > 65535 {
		log.Printf("invalid port format : %s", portString)
		return
	} else {
		tcpPort = int(port)
	}

	if resolvedIPs, err = d.Dialer.Resolver.LookupIP(ctx, network, host); err != nil {
		log.Printf("cannot resolve host %s : %s", err)
		return
	}

	addrs = make([]net.TCPAddr, 0, len(resolvedIPs))
	for _, resolvedIP := range resolvedIPs {
		addrs = append(addrs, net.TCPAddr{IP: resolvedIP, Port: tcpPort})
	}
	return
}

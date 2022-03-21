package networking

import "net"

// LocalIP returns the non loopback local IP of the host
func LocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

// ResolveIp is resolve a network address from a given hostname with default resolver.
func ResolveIp(name string) (net.IP, error) {
	if addrs, err := net.ResolveIPAddr("ip4", name); err != nil {
		return nil, err
	} else {
		return addrs.IP, nil
	}
}

func ExtractIp(remoteAddr net.Addr) string {
	if addr, ok := remoteAddr.(*net.TCPAddr); ok {
		return addr.IP.String()
	} else {
		return remoteAddr.String()
	}
}

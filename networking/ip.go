package networking

import (
	"net"
)

// Bigger than we need, not too big to worry about overflow
const big = 0x100

// Parse IPv4 address (d.d.d.d).
func ParseIPv4(s []byte) net.IP {
	var p [net.IPv4len]byte
	for i := 0; i < net.IPv4len; i++ {
		if len(s) == 0 {
			// Missing octets.
			return nil
		}

		if i > 0 {
			if s[0] != '.' {
				return nil
			}
			s = s[1:]
		}

		n, c, ok := dtoi(s)
		if !ok {
			return nil
		}
		if c > 1 && s[0] == '0' {
			// Reject non-zero components with leading zeroes.
			return nil
		}
		s = s[c:]
		p[i] = byte(n)
	}
	if len(s) != 0 {
		return nil
	}
	return net.IPv4(p[0], p[1], p[2], p[3])
}

// Decimal to integer.
// Returns number, characters consumed, success.
func dtoi(s []byte) (n int, i int, ok bool) {
	n = 0
	for i = 0; i < len(s) && '0' <= s[i] && s[i] <= '9'; i++ {
		n = n*10 + int(s[i]-'0')
		if n >= big {
			return big, i, false
		}
	}
	if i == 0 {
		return 0, 0, false
	}
	return n, i, true
}

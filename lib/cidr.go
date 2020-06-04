package lib

import (
	"net"
)

func ParseCIDRBlock(s string) (*net.IPNet, error) {
	_, c, err := net.ParseCIDR(s)
	return c, err
}

func LastIp(c *net.IPNet) net.IP {
	lastIp := make(net.IP, len(c.IP))

	ones, bits := c.Mask.Size()
	highMask := net.CIDRMask(ones, bits)
	for i, b := range c.IP {
		lastIp[i] = b | (^highMask[i])
	}

	return lastIp
}

func SubnetMask(m net.IPMask) string {
	return net.IP(m).String()
}

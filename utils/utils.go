package utils

import (
	"fmt"
	"log"
	"net"
	"strconv"
)

// number class helpers

func Atoi(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		i = 0
	}
	return i
}

func PosAtoi(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil || i < 0 {
		i = 0
	}
	return i
}

// network address class helpers

func IsValidHost(addr string) bool {
	return net.ParseIP(addr) != nil
}

func IsValidPort(port int) bool {
	return port > 0 && port < 65536
}

func GetAddress(host string, port int) *net.UDPAddr {
	str := fmt.Sprintf("%s:%d", host, port)
	addr, err := net.ResolveUDPAddr("udp", str)
	if err != nil {
		log.Fatalf("[CRIT] RESOLVE: Cannot resolve host address %s: %v.", str, err)
	}
	return addr
}

func GetNetwork(address, netmask string) string {
	addr := net.ParseIP(address)
	mask := net.ParseIP(netmask)
	for i := 0; i < len(addr); i++ {
		addr[i] &= mask[i]
	}
	return addr.String()
}

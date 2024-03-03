package main

import (
	"fmt"
	"net"
	"os"
	"runtime"

	"github.com/denisbrodbeck/machineid"
)

type machine struct {
	name string
	id   string
	os   string
	ipv4 string
	ipv6 string
}

// MachineInterface represents an interface for a machine.
type IMachine interface {
	GetName() string
	GetID() string
	GetOS() string
	GetIPV4() string
	GetIPV6() string
}

// NewMachine returns a new machine object.
func Machine() (IMachine, error) {
	ipv4, ipv6 := getIp()
	name := getName()
	uniqid, err := machineid.ID()
	if err != nil {
		return nil, fmt.Errorf("error get id: %v", err)
	}
	os := runtime.GOOS

	return &machine{
		ipv4: ipv4,
		ipv6: ipv6,
		name: name,
		id:   uniqid,
		os:   os,
	}, nil
}

// GetIPV4 returns the IPv4 address of the machine.
func (m *machine) GetIPV4() string {
	return m.ipv4
}

// GetIPV6 returns the IPv6 address of the machine.
func (m *machine) GetIPV6() string {
	return m.ipv6
}

// GetName returns the name of the machine.
func (m *machine) GetName() string {
	return m.name
}

// GetID returns the ID of the machine.
func (m *machine) GetID() string {
	return m.id
}

// GetOS returns the operating system of the machine.
func (m *machine) GetOS() string {
	return m.os
}

// getIp returns the IPv4 and IPv6 addresses of the machine.
func getIp() (string, string) {
	_defv4, _defv6 := "127.0.0.1", "::1"

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return _defv4, _defv6
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				_defv4 = ipnet.IP.String()
			} else if ipnet.IP.To16() != nil {
				_defv6 = ipnet.IP.String()
			}
		}
	}

	return _defv4, _defv6
}

// getName returns the name of the machine.
func getName() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "unknown"
	}
	return hostname
}

package model

import "net"

type Entity struct {
	ID              string
	Name            string
	ObtainedFromIP  net.IP
	ObtainedFromMAC net.HardwareAddr
}

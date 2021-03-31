package network

import (
	"github.com/vishvananda/netlink"
	"net"
)

type Driver interface {
	Name() string
	Create(sub, name string) (*Network, error)
	Delete(new *Network) error

	Connect(net *Network, endpoint *Endpoint) error
	Disconnect(net *Network, endpoint *Endpoint) error
}

type Endpoint struct {
	ID      string
	Driver  netlink.Veth
	IP      net.IP
	Mac     net.HardwareAddr
	Ports   []string
	NetWore *Network
}

var (
	drivers = map[string]Driver{}
)

func init() {
	var bridgeDriver = BridgeNetworkDriver{}
	drivers[bridgeDriver.Name()] = &bridgeDriver
}

package network

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/vishvananda/netlink"
	"net"
	"os"
	"path"
)

type Net struct {
	Name, Driver string
	IpRange      *net.IPNet
}

func (n *Net) dump(dumpPath string) error {
	if _, err := os.Stat(dumpPath); err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(dumpPath, 0644)
		} else {
			return err
		}
	}

	nwPath := path.Join(dumpPath, n.Name)
	nwFile, err := os.OpenFile(nwPath, os.O_TRUNC | os.O_WRONLY | os.O_CREATE, 0644)
	if err != nil {
		logrus.Errorf("error：", err)
		return err
	}
	defer nwFile.Close()

	nwJson, err := json.Marshal(n)
	if err != nil {
		logrus.Errorf("error：", err)
		return err
	}

	_, err = nwFile.Write(nwJson)
	if err != nil {
		logrus.Errorf("error：", err)
		return err
	}
	return nil
}

func (n *Net) remove(dumpPath string) error {
	if _, err := os.Stat(path.Join(dumpPath, n.Name)); err != nil {
		if os.IsNotExist(err) {
			return nil
		} else {
			return err
		}
	} else {
		return os.Remove(path.Join(dumpPath, n.Name))
	}
}

func (nw *Net) load(dumpPath string) error {
	nwConfigFile, err := os.Open(dumpPath)
	defer nwConfigFile.Close()
	if err != nil {
		return err
	}
	nwJson := make([]byte, 2000)
	n, err := nwConfigFile.Read(nwJson)
	if err != nil {
		return err
	}

	err = json.Unmarshal(nwJson[:n], nw)
	if err != nil {
		logrus.Errorf("Error load nw info", err)
		return err
	}
	return nil
}


type Endpoint struct {
	ID      string
	Driver  netlink.Veth
	IP      net.IP
	Mac     net.HardwareAddr
	Ports   []string
	NetWore *Net
}

var (
	defaultNetworkPath = "/var/run/mydocker/network/network/"
	drivers = map[string]Driver{}
	ipAllocator = &IPAM{
		path: IPAMDefaultAllocatorPath,
	}
)

func CreateNetwork(driver, subnet, name string) error {
	_, cidr, _ := net.ParseCIDR(subnet)
	ip, err := ipAllocator.Allocate(cidr)
	if err != nil {
		return err
	}
	cidr.IP = ip
	network, err := drivers[driver].Create(cidr.String(), name)
	if err != nil {
		return err
	}
	return network.dump(defaultNetworkPath)
}

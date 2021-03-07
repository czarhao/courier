package network

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net"
	"os"
	"path"
)

type Network struct {
	Name, Driver string
	IpRange      *net.IPNet
}

var (
	defaultNetworkPath = "/root/.network/"
	ipAllocator = NewIPAM(IPAMDefaultAllocatorPath)
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

func (n *Network) dump(dumpPath string) error {
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
		logrus.Errorf("error：%v", err)
		return err
	}
	defer nwFile.Close()

	nwJson, err := json.Marshal(n)
	if err != nil {
		logrus.Errorf("error：%v", err)
		return err
	}

	_, err = nwFile.Write(nwJson)
	if err != nil {
		logrus.Errorf("error：%v", err)
		return err
	}
	return nil
}

func (n *Network) remove(dumpPath string) error {
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

func (nw *Network) load(dumpPath string) error {
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
		logrus.Errorf("Error load nw info, err: %v", err)
		return err
	}
	return nil
}

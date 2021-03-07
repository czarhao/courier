package network

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"os"
	"path"
	"strings"
)

const (
	IPAMDefaultAllocatorPath = "/root/.subnet.json"
)

type IPAM interface {
	Allocate(subnet *net.IPNet) (net.IP, error)
	Release(subnet *net.IPNet, ipaddr *net.IP) error
}

type ipam struct {
	path    string
	subnets map[string]string
}

func NewIPAM(path string) *ipam {
	if path == "" {
		path = IPAMDefaultAllocatorPath
	}
	return &ipam{
		path:    path,
		subnets: map[string]string{},
	}
}

func (i *ipam) load() error {
	if _, err := os.Stat(i.path); err != nil {
		if os.IsNotExist(err) {
			return nil
		} else {
			return err
		}
	}
	subnetCfg, err := os.Open(i.path)
	if err != nil {
		return err
	}
	defer subnetCfg.Close()
	config, err := ioutil.ReadAll(subnetCfg)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(config, &i.subnets); err != nil {
		return err
	}
	return nil
}

func (i *ipam) dump() error {
	dir, _ := path.Split(i.path)
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) && os.MkdirAll(i.path, 0644) != nil {
			return err
		}
	}
	configFile, err := os.OpenFile(i.path, os.O_TRUNC | os.O_WRONLY | os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer configFile.Close()
	configJson, err := json.Marshal(i.subnets)
	if err != nil {
		return err
	}
	_, err = configFile.Write(configJson)
	return err
}

func (i *ipam) Allocate(subnet *net.IPNet) (net.IP, error) {
	if err := i.load(); err != nil {
		return nil, err
	}
	_, subnet, _ = net.ParseCIDR(subnet.String())
	one, size := subnet.Mask.Size()

	if _, exist := i.subnets[subnet.String()]; !exist {
		i.subnets[subnet.String()] = strings.Repeat("0", 1 << uint8(size - one))
	}

	var ip net.IP

	for c := range i.subnets[subnet.String()] {
		if i.subnets[subnet.String()][c] == '0' {
			ipalloc := []byte(i.subnets[subnet.String()])
			ipalloc[c] = '1'
			i.subnets[subnet.String()] = string(ipalloc)
			ip = subnet.IP
			for t := uint(4); t > 0; t -=1{
				[]byte(ip)[4-t] += uint8(c >> ((t - 1) * 8))
			}
			ip[3] += 1
			break
		}
	}

	return ip, i.dump()
}

func (i *ipam) Release(subnet *net.IPNet, ipaddr *net.IP) error {
	i.subnets = map[string]string{}
	_, subnet, _ = net.ParseCIDR(subnet.String())

	if err := i.load(); err != nil {
		return err
	}
	var (
		c = 0
		releaseIP = ipaddr.To4()
	)

	releaseIP[3]-=1
	for t := uint(4); t > 0; t-=1 {
		c += int(releaseIP[t-1] - subnet.IP[t-1]) << ((4-t) * 8)
	}

	ipalloc := []byte(i.subnets[subnet.String()])
	ipalloc[c] = '0'
	i.subnets[subnet.String()] = string(ipalloc)

	return i.dump()
}
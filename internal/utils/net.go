package utils

import (
	"errors"
	"net"
)

var (
	errInterfaceNotFound = errors.New("net: provided interface does not exist")
)

func Net_MacAdrr(ifName string) ([]byte, error) {
	ifas, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, ifa := range ifas {
		if ifa.Name == ifName {
			return []byte(ifa.HardwareAddr), nil
		}
	}
	return nil, errInterfaceNotFound
}

package gobtphelper

import (
	"strings"
)

var GlobalServices []*Service

func GetServiceAddress(serviceName string) (addresses []string) {
	for _, v := range GlobalServices {
		if strings.HasPrefix(v.ServiceName, serviceName+"@") {
			addresses = append(addresses, v.Address)
		}
	}
	return addresses
}

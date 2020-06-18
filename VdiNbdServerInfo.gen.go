// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w VdiNbdServerInfo.gen.go

package go_xen_client

import (
	"reflect"
	"strconv"

	"github.com/nilshell/xmlrpc"
)

//VdiNbdServerInfo: Details for connecting to a VDI using the Network Block Device protocol
type VdiNbdServerInfo struct {
	Exportname string // The exportname to request over NBD. This holds details including an authentication token, so it must be protected appropriately. Clients should regard the exportname as an opaque string or token.
	Address    string // An address on which the server can be reached; this can be IPv4, IPv6, or a DNS name.
	Port       int    // The TCP port
	Cert       string // The TLS certificate of the server
	Subject    string // For convenience, this redundant field holds a DNS (hostname) subject of the certificate. This can be a wildcard, but only for a certificate that has a wildcard subject and no concrete hostname subjects.

}

func FromVdiNbdServerInfoToXml(vdi_nbd_server_info *VdiNbdServerInfo) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["exportname"] =

		vdi_nbd_server_info.Exportname

	result["address"] =

		vdi_nbd_server_info.Address

	result["port"] =

		strconv.Itoa(vdi_nbd_server_info.Port)

	result["cert"] =

		vdi_nbd_server_info.Cert

	result["subject"] =

		vdi_nbd_server_info.Subject

	return result
}

func ToVdiNbdServerInfo(obj interface{}) (resultObj *VdiNbdServerInfo) {

	objValue := reflect.ValueOf(obj)
	resultObj = &VdiNbdServerInfo{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "exportname":
			if v, ok := keyValue.(string); ok {
				resultObj.Exportname = v
			}
		case "address":
			if v, ok := keyValue.(string); ok {
				resultObj.Address = v
			}
		case "port":
			if v, ok := keyValue.(int); ok {
				resultObj.Port = v
			}
		case "cert":
			if v, ok := keyValue.(string); ok {
				resultObj.Cert = v
			}
		case "subject":
			if v, ok := keyValue.(string); ok {
				resultObj.Subject = v
			}
		}
	}

	return resultObj
}

// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w PVSServer.gen.go

package go_xen_client

import (
	"reflect"
	"strconv"

	"github.com/nilshell/xmlrpc"
)

//PVSServer: individual machine serving provisioning (block) data
type PVSServer struct {
	Uuid      string   // Unique identifier/object reference
	Addresses []string // IPv4 addresses of this server
	FirstPort int      // First UDP port accepted by this server
	LastPort  int      // Last UDP port accepted by this server
	Site      string   // PVS site this server is part of

}

func FromPVSServerToXml(PVS_server *PVSServer) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] = PVS_server.Uuid

	result["addresses"] = PVS_server.Addresses

	result["first_port"] = strconv.Itoa(PVS_server.FirstPort)

	result["last_port"] = strconv.Itoa(PVS_server.LastPort)

	result["site"] = PVS_server.Site

	return result
}

func ToPVSServer(obj interface{}) (resultObj *PVSServer) {

	objValue := reflect.ValueOf(obj)
	resultObj = &PVSServer{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.Uuid = v
			}
		case "addresses":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.Addresses = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.Addresses[i] = v
					}
				}
			}
		case "first_port":
			if v, ok := keyValue.(int); ok {
				resultObj.FirstPort = v
			}
		case "last_port":
			if v, ok := keyValue.(int); ok {
				resultObj.LastPort = v
			}
		case "site":
			if v, ok := keyValue.(string); ok {
				resultObj.Site = v
			}
		}
	}

	return resultObj
}

/* GetAllRecords: Return a map of PVS_server references to PVS_server records for all PVS_servers known to the system. */
func (client *XenClient) PVSServerGetAllRecords() (result map[string]PVSServer, err error) {
	obj, err := client.APICall("PVS_server.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]PVSServer{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToPVSServer(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the PVS_servers known to the system. */
func (client *XenClient) PVSServerGetAll() (result []string, err error) {
	obj, err := client.APICall("PVS_server.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* Forget: forget a PVS server */
func (client *XenClient) PVSServerForget(self string) (err error) {
	_, err = client.APICall("PVS_server.forget", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Introduce: introduce new PVS server */
func (client *XenClient) PVSServerIntroduce(addresses []string, first_port int, last_port int, site string) (result string, err error) {
	obj, err := client.APICall("PVS_server.introduce", addresses, first_port, last_port, site)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetSite: Get the site field of the given PVS_server. */
func (client *XenClient) PVSServerGetSite(self string) (result string, err error) {
	obj, err := client.APICall("PVS_server.get_site", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetLastPort: Get the last_port field of the given PVS_server. */
func (client *XenClient) PVSServerGetLastPort(self string) (result int, err error) {
	obj, err := client.APICall("PVS_server.get_last_port", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetFirstPort: Get the first_port field of the given PVS_server. */
func (client *XenClient) PVSServerGetFirstPort(self string) (result int, err error) {
	obj, err := client.APICall("PVS_server.get_first_port", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetAddresses: Get the addresses field of the given PVS_server. */
func (client *XenClient) PVSServerGetAddresses(self string) (result []string, err error) {
	obj, err := client.APICall("PVS_server.get_addresses", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetUuid: Get the uuid field of the given PVS_server. */
func (client *XenClient) PVSServerGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("PVS_server.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the PVS_server instance with the specified UUID. */
func (client *XenClient) PVSServerGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("PVS_server.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given PVS_server. */
func (client *XenClient) PVSServerGetRecord(self string) (result PVSServer, err error) {
	obj, err := client.APICall("PVS_server.get_record", self)
	if err != nil {
		return
	}

	result = *ToPVSServer(obj.(interface{}))

	return
}

// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w Tunnel.gen.go

package go_xen_client

import (
	"reflect"

	"github.com/nilshell/xmlrpc"
)

//Tunnel: A tunnel for network traffic
type Tunnel struct {
	Uuid         string            // Unique identifier/object reference
	AccessPIF    string            // The interface through which the tunnel is accessed
	TransportPIF string            // The interface used by the tunnel
	Status       map[string]string // Status information about the tunnel
	OtherConfig  map[string]string // Additional configuration

}

func FromTunnelToXml(tunnel *Tunnel) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] = tunnel.Uuid

	result["access_PIF"] = tunnel.AccessPIF

	result["transport_PIF"] = tunnel.TransportPIF

	status := make(xmlrpc.Struct)
	for key, value := range tunnel.Status {
		status[key] = value
	}
	result["status"] = status

	other_config := make(xmlrpc.Struct)
	for key, value := range tunnel.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	return result
}

func ToTunnel(obj interface{}) (resultObj *Tunnel) {

	objValue := reflect.ValueOf(obj)
	resultObj = &Tunnel{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.Uuid = v
			}
		case "access_PIF":
			if v, ok := keyValue.(string); ok {
				resultObj.AccessPIF = v
			}
		case "transport_PIF":
			if v, ok := keyValue.(string); ok {
				resultObj.TransportPIF = v
			}
		case "status":

			resultObj.Status = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.Status[mapKeyName] = v
				} else {
					resultObj.Status[mapKeyName] = ""
				}
			}
		case "other_config":

			resultObj.OtherConfig = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.OtherConfig[mapKeyName] = v
				} else {
					resultObj.OtherConfig[mapKeyName] = ""
				}
			}
		}
	}

	return resultObj
}

/* GetAllRecords: Return a map of tunnel references to tunnel records for all tunnels known to the system. */
func (client *XenClient) TunnelGetAllRecords() (result map[string]Tunnel, err error) {
	obj, err := client.APICall("tunnel.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]Tunnel{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToTunnel(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the tunnels known to the system. */
func (client *XenClient) TunnelGetAll() (result []string, err error) {
	obj, err := client.APICall("tunnel.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* Destroy: Destroy a tunnel */
func (client *XenClient) TunnelDestroy(self string) (err error) {
	_, err = client.APICall("tunnel.destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Create: Create a tunnel */
func (client *XenClient) TunnelCreate(transport_PIF string, network string) (result string, err error) {
	obj, err := client.APICall("tunnel.create", transport_PIF, network)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given tunnel.  If the key is not in that Map, then do nothing. */
func (client *XenClient) TunnelRemoveFromOtherConfig(self string, key string) (err error) {
	_, err = client.APICall("tunnel.remove_from_other_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToOtherConfig: Add the given key-value pair to the other_config field of the given tunnel. */
func (client *XenClient) TunnelAddToOtherConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("tunnel.add_to_other_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOtherConfig: Set the other_config field of the given tunnel. */
func (client *XenClient) TunnelSetOtherConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("tunnel.set_other_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveFromStatus: Remove the given key and its corresponding value from the status field of the given tunnel.  If the key is not in that Map, then do nothing. */
func (client *XenClient) TunnelRemoveFromStatus(self string, key string) (err error) {
	_, err = client.APICall("tunnel.remove_from_status", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToStatus: Add the given key-value pair to the status field of the given tunnel. */
func (client *XenClient) TunnelAddToStatus(self string, key string, value string) (err error) {
	_, err = client.APICall("tunnel.add_to_status", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetStatus: Set the status field of the given tunnel. */
func (client *XenClient) TunnelSetStatus(self string, value map[string]string) (err error) {
	_, err = client.APICall("tunnel.set_status", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetOtherConfig: Get the other_config field of the given tunnel. */
func (client *XenClient) TunnelGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("tunnel.get_other_config", self)
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]string{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		result[key.String()] = obj.String()
	}

	return
}

/* GetStatus: Get the status field of the given tunnel. */
func (client *XenClient) TunnelGetStatus(self string) (result map[string]string, err error) {
	obj, err := client.APICall("tunnel.get_status", self)
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]string{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		result[key.String()] = obj.String()
	}

	return
}

/* GetTransportPIF: Get the transport_PIF field of the given tunnel. */
func (client *XenClient) TunnelGetTransportPIF(self string) (result string, err error) {
	obj, err := client.APICall("tunnel.get_transport_PIF", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetAccessPIF: Get the access_PIF field of the given tunnel. */
func (client *XenClient) TunnelGetAccessPIF(self string) (result string, err error) {
	obj, err := client.APICall("tunnel.get_access_PIF", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given tunnel. */
func (client *XenClient) TunnelGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("tunnel.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the tunnel instance with the specified UUID. */
func (client *XenClient) TunnelGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("tunnel.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given tunnel. */
func (client *XenClient) TunnelGetRecord(self string) (result Tunnel, err error) {
	obj, err := client.APICall("tunnel.get_record", self)
	if err != nil {
		return
	}

	result = *ToTunnel(obj.(interface{}))

	return
}

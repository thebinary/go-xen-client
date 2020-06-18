// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w PBD.gen.go

package go_xen_client

import (
	"reflect"

	"github.com/nilshell/xmlrpc"
)

//PBD: The physical block devices through which hosts access SRs
type PBD struct {
	Uuid              string            // Unique identifier/object reference
	Host              string            // physical machine on which the pbd is available
	SR                string            // the storage repository that the pbd realises
	DeviceConfig      map[string]string // a config string to string map that is provided to the host's SR-backend-driver
	CurrentlyAttached bool              // is the SR currently attached on this host?
	OtherConfig       map[string]string // additional configuration

}

func FromPBDToXml(PBD *PBD) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] =

		PBD.Uuid

	result["host"] =

		PBD.Host

	result["SR"] =

		PBD.SR

	device_config := make(xmlrpc.Struct)
	for key, value := range PBD.DeviceConfig {
		device_config[key] = value
	}
	result["device_config"] = device_config

	result["currently_attached"] =

		PBD.CurrentlyAttached

	other_config := make(xmlrpc.Struct)
	for key, value := range PBD.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	return result
}

func ToPBD(obj interface{}) (resultObj *PBD) {

	objValue := reflect.ValueOf(obj)
	resultObj = &PBD{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.Uuid = v
			}
		case "host":
			if v, ok := keyValue.(string); ok {
				resultObj.Host = v
			}
		case "SR":
			if v, ok := keyValue.(string); ok {
				resultObj.SR = v
			}
		case "device_config":

			resultObj.DeviceConfig = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.DeviceConfig[mapKeyName] = v
				} else {
					resultObj.DeviceConfig[mapKeyName] = ""
				}
			}
		case "currently_attached":
			if v, ok := keyValue.(bool); ok {
				resultObj.CurrentlyAttached = v
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

/* GetAllRecords: Return a map of PBD references to PBD records for all PBDs known to the system. */
func (client *XenClient) PBDGetAllRecords() (result map[string]PBD, err error) {
	obj, err := client.APICall("PBD.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]PBD{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToPBD(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the PBDs known to the system. */
func (client *XenClient) PBDGetAll() (result []string, err error) {
	obj, err := client.APICall("PBD.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* SetDeviceConfig: Sets the PBD's device_config field */
func (client *XenClient) PBDSetDeviceConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("PBD.set_device_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Unplug: Deactivate the specified PBD, causing the referenced SR to be detached and nolonger scanned */
func (client *XenClient) PBDUnplug(self string) (err error) {
	_, err = client.APICall("PBD.unplug", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Plug: Activate the specified PBD, causing the referenced SR to be attached and scanned */
func (client *XenClient) PBDPlug(self string) (err error) {
	_, err = client.APICall("PBD.plug", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given PBD.  If the key is not in that Map, then do nothing. */
func (client *XenClient) PBDRemoveFromOtherConfig(self string, key string) (err error) {
	_, err = client.APICall("PBD.remove_from_other_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToOtherConfig: Add the given key-value pair to the other_config field of the given PBD. */
func (client *XenClient) PBDAddToOtherConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("PBD.add_to_other_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOtherConfig: Set the other_config field of the given PBD. */
func (client *XenClient) PBDSetOtherConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("PBD.set_other_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetOtherConfig: Get the other_config field of the given PBD. */
func (client *XenClient) PBDGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("PBD.get_other_config", self)
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

/* GetCurrentlyAttached: Get the currently_attached field of the given PBD. */
func (client *XenClient) PBDGetCurrentlyAttached(self string) (result bool, err error) {
	obj, err := client.APICall("PBD.get_currently_attached", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetDeviceConfig: Get the device_config field of the given PBD. */
func (client *XenClient) PBDGetDeviceConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("PBD.get_device_config", self)
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

/* GetSR: Get the SR field of the given PBD. */
func (client *XenClient) PBDGetSR(self string) (result string, err error) {
	obj, err := client.APICall("PBD.get_SR", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetHost: Get the host field of the given PBD. */
func (client *XenClient) PBDGetHost(self string) (result string, err error) {
	obj, err := client.APICall("PBD.get_host", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given PBD. */
func (client *XenClient) PBDGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("PBD.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* Destroy: Destroy the specified PBD instance. */
func (client *XenClient) PBDDestroy(self string) (err error) {
	_, err = client.APICall("PBD.destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Create: Create a new PBD instance, and return its handle.
The constructor args are: host*, SR*, device_config*, other_config (* = non-optional). */
func (client *XenClient) PBDCreate(args PBD) (result string, err error) {
	obj, err := client.APICall("PBD.create", FromPBDToXml(&args))
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the PBD instance with the specified UUID. */
func (client *XenClient) PBDGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("PBD.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given PBD. */
func (client *XenClient) PBDGetRecord(self string) (result PBD, err error) {
	obj, err := client.APICall("PBD.get_record", self)
	if err != nil {
		return
	}

	result = *ToPBD(obj.(interface{}))

	return
}

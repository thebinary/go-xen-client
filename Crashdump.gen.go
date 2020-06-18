// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w Crashdump.gen.go

package go_xen_client

import (
	"reflect"

	"github.com/nilshell/xmlrpc"
)

//Crashdump: A VM crashdump
type Crashdump struct {
	Uuid        string            // Unique identifier/object reference
	VM          string            // the virtual machine
	VDI         string            // the virtual disk
	OtherConfig map[string]string // additional configuration

}

func FromCrashdumpToXml(crashdump *Crashdump) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] = crashdump.Uuid

	result["VM"] = crashdump.VM

	result["VDI"] = crashdump.VDI

	other_config := make(xmlrpc.Struct)
	for key, value := range crashdump.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	return result
}

func ToCrashdump(obj interface{}) (resultObj *Crashdump) {

	objValue := reflect.ValueOf(obj)
	resultObj = &Crashdump{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.Uuid = v
			}
		case "VM":
			if v, ok := keyValue.(string); ok {
				resultObj.VM = v
			}
		case "VDI":
			if v, ok := keyValue.(string); ok {
				resultObj.VDI = v
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

/* GetAllRecords: Return a map of crashdump references to crashdump records for all crashdumps known to the system. */
func (client *XenClient) CrashdumpGetAllRecords() (result map[string]Crashdump, err error) {
	obj, err := client.APICall("crashdump.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]Crashdump{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToCrashdump(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the crashdumps known to the system. */
func (client *XenClient) CrashdumpGetAll() (result []string, err error) {
	obj, err := client.APICall("crashdump.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* Destroy: Destroy the specified crashdump */
func (client *XenClient) CrashdumpDestroy(self string) (err error) {
	_, err = client.APICall("crashdump.destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given crashdump.  If the key is not in that Map, then do nothing. */
func (client *XenClient) CrashdumpRemoveFromOtherConfig(self string, key string) (err error) {
	_, err = client.APICall("crashdump.remove_from_other_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToOtherConfig: Add the given key-value pair to the other_config field of the given crashdump. */
func (client *XenClient) CrashdumpAddToOtherConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("crashdump.add_to_other_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOtherConfig: Set the other_config field of the given crashdump. */
func (client *XenClient) CrashdumpSetOtherConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("crashdump.set_other_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetOtherConfig: Get the other_config field of the given crashdump. */
func (client *XenClient) CrashdumpGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("crashdump.get_other_config", self)
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

/* GetVDI: Get the VDI field of the given crashdump. */
func (client *XenClient) CrashdumpGetVDI(self string) (result string, err error) {
	obj, err := client.APICall("crashdump.get_VDI", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetVM: Get the VM field of the given crashdump. */
func (client *XenClient) CrashdumpGetVM(self string) (result string, err error) {
	obj, err := client.APICall("crashdump.get_VM", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given crashdump. */
func (client *XenClient) CrashdumpGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("crashdump.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the crashdump instance with the specified UUID. */
func (client *XenClient) CrashdumpGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("crashdump.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given crashdump. */
func (client *XenClient) CrashdumpGetRecord(self string) (result Crashdump, err error) {
	obj, err := client.APICall("crashdump.get_record", self)
	if err != nil {
		return
	}

	result = *ToCrashdump(obj.(interface{}))

	return
}

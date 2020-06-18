// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w VUSB.gen.go

package go_xen_client

import (
	"reflect"

	"github.com/nilshell/xmlrpc"
)

//VUSB: Describes the vusb device
type VUSB struct {
	Uuid              string                    // Unique identifier/object reference
	AllowedOperations []VusbOperations          // list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	CurrentOperations map[string]VusbOperations // links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	VM                string                    // VM that owns the VUSB
	USBGroup          string                    // USB group used by the VUSB
	OtherConfig       map[string]string         // Additional configuration
	CurrentlyAttached bool                      // is the device currently attached

}

func FromVUSBToXml(VUSB *VUSB) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] =

		VUSB.Uuid

	result["allowed_operations"] =

		VUSB.AllowedOperations

	current_operations := make(xmlrpc.Struct)
	for key, value := range VUSB.CurrentOperations {
		current_operations[key] = value
	}
	result["current_operations"] = current_operations

	result["VM"] =

		VUSB.VM

	result["USB_group"] =

		VUSB.USBGroup

	other_config := make(xmlrpc.Struct)
	for key, value := range VUSB.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	result["currently_attached"] =

		VUSB.CurrentlyAttached

	return result
}

func ToVUSB(obj interface{}) (resultObj *VUSB) {

	objValue := reflect.ValueOf(obj)
	resultObj = &VUSB{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.Uuid = v
			}
		case "allowed_operations":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.AllowedOperations = make([]VusbOperations, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(VusbOperations); ok {
						resultObj.AllowedOperations[i] = v
					}
				}
			}
		case "current_operations":

			resultObj.CurrentOperations = map[string]VusbOperations{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.CurrentOperations[mapKeyName] = ToVusbOperations(v)
				} else {
					resultObj.CurrentOperations[mapKeyName] = 0
				}
			}
		case "VM":
			if v, ok := keyValue.(string); ok {
				resultObj.VM = v
			}
		case "USB_group":
			if v, ok := keyValue.(string); ok {
				resultObj.USBGroup = v
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
		case "currently_attached":
			if v, ok := keyValue.(bool); ok {
				resultObj.CurrentlyAttached = v
			}
		}
	}

	return resultObj
}

/* GetAllRecords: Return a map of VUSB references to VUSB records for all VUSBs known to the system. */
func (client *XenClient) VUSBGetAllRecords() (result map[string]VUSB, err error) {
	obj, err := client.APICall("VUSB.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]VUSB{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToVUSB(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the VUSBs known to the system. */
func (client *XenClient) VUSBGetAll() (result []string, err error) {
	obj, err := client.APICall("VUSB.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* Destroy: Removes a VUSB record from the database */
func (client *XenClient) VUSBDestroy(self string) (err error) {
	_, err = client.APICall("VUSB.destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Unplug: Unplug the vusb device from the vm. */
func (client *XenClient) VUSBUnplug(self string) (err error) {
	_, err = client.APICall("VUSB.unplug", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Create: Create a new VUSB record in the database only */
func (client *XenClient) VUSBCreate(VM string, USB_group string, other_config map[string]string) (result string, err error) {
	obj, err := client.APICall("VUSB.create", VM, USB_group, other_config)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given VUSB.  If the key is not in that Map, then do nothing. */
func (client *XenClient) VUSBRemoveFromOtherConfig(self string, key string) (err error) {
	_, err = client.APICall("VUSB.remove_from_other_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToOtherConfig: Add the given key-value pair to the other_config field of the given VUSB. */
func (client *XenClient) VUSBAddToOtherConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("VUSB.add_to_other_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOtherConfig: Set the other_config field of the given VUSB. */
func (client *XenClient) VUSBSetOtherConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("VUSB.set_other_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetCurrentlyAttached: Get the currently_attached field of the given VUSB. */
func (client *XenClient) VUSBGetCurrentlyAttached(self string) (result bool, err error) {
	obj, err := client.APICall("VUSB.get_currently_attached", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetOtherConfig: Get the other_config field of the given VUSB. */
func (client *XenClient) VUSBGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VUSB.get_other_config", self)
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

/* GetUSBGroup: Get the USB_group field of the given VUSB. */
func (client *XenClient) VUSBGetUSBGroup(self string) (result string, err error) {
	obj, err := client.APICall("VUSB.get_USB_group", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetVM: Get the VM field of the given VUSB. */
func (client *XenClient) VUSBGetVM(self string) (result string, err error) {
	obj, err := client.APICall("VUSB.get_VM", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetCurrentOperations: Get the current_operations field of the given VUSB. */
func (client *XenClient) VUSBGetCurrentOperations(self string) (result map[string]VusbOperations, err error) {
	obj, err := client.APICall("VUSB.get_current_operations", self)
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]VusbOperations{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToVusbOperations(obj.String())
		result[key.String()] = mapObj
	}

	return
}

/* GetAllowedOperations: Get the allowed_operations field of the given VUSB. */
func (client *XenClient) VUSBGetAllowedOperations(self string) (result []VusbOperations, err error) {
	obj, err := client.APICall("VUSB.get_allowed_operations", self)
	if err != nil {
		return
	}

	result = make([]VusbOperations, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = ToVusbOperations(value.(string))
	}

	return
}

/* GetUuid: Get the uuid field of the given VUSB. */
func (client *XenClient) VUSBGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("VUSB.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the VUSB instance with the specified UUID. */
func (client *XenClient) VUSBGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("VUSB.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given VUSB. */
func (client *XenClient) VUSBGetRecord(self string) (result VUSB, err error) {
	obj, err := client.APICall("VUSB.get_record", self)
	if err != nil {
		return
	}

	result = *ToVUSB(obj.(interface{}))

	return
}

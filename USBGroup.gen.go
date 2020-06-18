// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w USBGroup.gen.go

package go_xen_client

import (
	"reflect"

	"github.com/nilshell/xmlrpc"
)

//USBGroup: A group of compatible USBs across the resource pool
type USBGroup struct {
	Uuid            string            // Unique identifier/object reference
	NameLabel       string            // a human-readable name
	NameDescription string            // a notes field containing human-readable description
	PUSBs           []string          // List of PUSBs in the group
	VUSBs           []string          // List of VUSBs using the group
	OtherConfig     map[string]string // Additional configuration

}

func FromUSBGroupToXml(USB_group *USBGroup) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] =

		USB_group.Uuid

	result["name_label"] =

		USB_group.NameLabel

	result["name_description"] =

		USB_group.NameDescription

	result["PUSBs"] =

		USB_group.PUSBs

	result["VUSBs"] =

		USB_group.VUSBs

	other_config := make(xmlrpc.Struct)
	for key, value := range USB_group.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	return result
}

func ToUSBGroup(obj interface{}) (resultObj *USBGroup) {

	objValue := reflect.ValueOf(obj)
	resultObj = &USBGroup{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.Uuid = v
			}
		case "name_label":
			if v, ok := keyValue.(string); ok {
				resultObj.NameLabel = v
			}
		case "name_description":
			if v, ok := keyValue.(string); ok {
				resultObj.NameDescription = v
			}
		case "PUSBs":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.PUSBs = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.PUSBs[i] = v
					}
				}
			}
		case "VUSBs":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.VUSBs = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.VUSBs[i] = v
					}
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

/* GetAllRecords: Return a map of USB_group references to USB_group records for all USB_groups known to the system. */
func (client *XenClient) USBGroupGetAllRecords() (result map[string]USBGroup, err error) {
	obj, err := client.APICall("USB_group.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]USBGroup{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToUSBGroup(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the USB_groups known to the system. */
func (client *XenClient) USBGroupGetAll() (result []string, err error) {
	obj, err := client.APICall("USB_group.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* Destroy:  */
func (client *XenClient) USBGroupDestroy(self string) (err error) {
	_, err = client.APICall("USB_group.destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Create:  */
func (client *XenClient) USBGroupCreate(name_label string, name_description string, other_config map[string]string) (result string, err error) {
	obj, err := client.APICall("USB_group.create", name_label, name_description, other_config)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given USB_group.  If the key is not in that Map, then do nothing. */
func (client *XenClient) USBGroupRemoveFromOtherConfig(self string, key string) (err error) {
	_, err = client.APICall("USB_group.remove_from_other_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToOtherConfig: Add the given key-value pair to the other_config field of the given USB_group. */
func (client *XenClient) USBGroupAddToOtherConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("USB_group.add_to_other_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOtherConfig: Set the other_config field of the given USB_group. */
func (client *XenClient) USBGroupSetOtherConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("USB_group.set_other_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetNameDescription: Set the name/description field of the given USB_group. */
func (client *XenClient) USBGroupSetNameDescription(self string, value string) (err error) {
	_, err = client.APICall("USB_group.set_name_description", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetNameLabel: Set the name/label field of the given USB_group. */
func (client *XenClient) USBGroupSetNameLabel(self string, value string) (err error) {
	_, err = client.APICall("USB_group.set_name_label", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetOtherConfig: Get the other_config field of the given USB_group. */
func (client *XenClient) USBGroupGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("USB_group.get_other_config", self)
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

/* GetVUSBs: Get the VUSBs field of the given USB_group. */
func (client *XenClient) USBGroupGetVUSBs(self string) (result []string, err error) {
	obj, err := client.APICall("USB_group.get_VUSBs", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetPUSBs: Get the PUSBs field of the given USB_group. */
func (client *XenClient) USBGroupGetPUSBs(self string) (result []string, err error) {
	obj, err := client.APICall("USB_group.get_PUSBs", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetNameDescription: Get the name/description field of the given USB_group. */
func (client *XenClient) USBGroupGetNameDescription(self string) (result string, err error) {
	obj, err := client.APICall("USB_group.get_name_description", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetNameLabel: Get the name/label field of the given USB_group. */
func (client *XenClient) USBGroupGetNameLabel(self string) (result string, err error) {
	obj, err := client.APICall("USB_group.get_name_label", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given USB_group. */
func (client *XenClient) USBGroupGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("USB_group.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByNameLabel: Get all the USB_group instances with the given label. */
func (client *XenClient) USBGroupGetByNameLabel(label string) (result []string, err error) {
	obj, err := client.APICall("USB_group.get_by_name_label", label)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetByUuid: Get a reference to the USB_group instance with the specified UUID. */
func (client *XenClient) USBGroupGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("USB_group.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given USB_group. */
func (client *XenClient) USBGroupGetRecord(self string) (result USBGroup, err error) {
	obj, err := client.APICall("USB_group.get_record", self)
	if err != nil {
		return
	}

	result = *ToUSBGroup(obj.(interface{}))

	return
}

// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w VLAN.gen.go

package go_xen_client

import (
	"reflect"
	"strconv"

	"github.com/nilshell/xmlrpc"
)

//VLAN: A VLAN mux/demux
type VLAN struct {
	Uuid        string            // Unique identifier/object reference
	TaggedPIF   string            // interface on which traffic is tagged
	UntaggedPIF string            // interface on which traffic is untagged
	Tag         int               // VLAN tag in use
	OtherConfig map[string]string // additional configuration

}

func FromVLANToXml(VLAN *VLAN) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] = VLAN.Uuid

	result["tagged_PIF"] = VLAN.TaggedPIF

	result["untagged_PIF"] = VLAN.UntaggedPIF

	result["tag"] = strconv.Itoa(VLAN.Tag)

	other_config := make(xmlrpc.Struct)
	for key, value := range VLAN.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	return result
}

func ToVLAN(obj interface{}) (resultObj *VLAN) {

	objValue := reflect.ValueOf(obj)
	resultObj = &VLAN{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.Uuid = v
			}
		case "tagged_PIF":
			if v, ok := keyValue.(string); ok {
				resultObj.TaggedPIF = v
			}
		case "untagged_PIF":
			if v, ok := keyValue.(string); ok {
				resultObj.UntaggedPIF = v
			}
		case "tag":
			if v, ok := keyValue.(int); ok {
				resultObj.Tag = v
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

/* GetAllRecords: Return a map of VLAN references to VLAN records for all VLANs known to the system. */
func (client *XenClient) VLANGetAllRecords() (result map[string]VLAN, err error) {
	obj, err := client.APICall("VLAN.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]VLAN{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToVLAN(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the VLANs known to the system. */
func (client *XenClient) VLANGetAll() (result []string, err error) {
	obj, err := client.APICall("VLAN.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* Destroy: Destroy a VLAN mux/demuxer */
func (client *XenClient) VLANDestroy(self string) (err error) {
	_, err = client.APICall("VLAN.destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Create: Create a VLAN mux/demuxer */
func (client *XenClient) VLANCreate(tagged_PIF string, tag int, network string) (result string, err error) {
	obj, err := client.APICall("VLAN.create", tagged_PIF, tag, network)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given VLAN.  If the key is not in that Map, then do nothing. */
func (client *XenClient) VLANRemoveFromOtherConfig(self string, key string) (err error) {
	_, err = client.APICall("VLAN.remove_from_other_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToOtherConfig: Add the given key-value pair to the other_config field of the given VLAN. */
func (client *XenClient) VLANAddToOtherConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("VLAN.add_to_other_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOtherConfig: Set the other_config field of the given VLAN. */
func (client *XenClient) VLANSetOtherConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("VLAN.set_other_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetOtherConfig: Get the other_config field of the given VLAN. */
func (client *XenClient) VLANGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VLAN.get_other_config", self)
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

/* GetTag: Get the tag field of the given VLAN. */
func (client *XenClient) VLANGetTag(self string) (result int, err error) {
	obj, err := client.APICall("VLAN.get_tag", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetUntaggedPIF: Get the untagged_PIF field of the given VLAN. */
func (client *XenClient) VLANGetUntaggedPIF(self string) (result string, err error) {
	obj, err := client.APICall("VLAN.get_untagged_PIF", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetTaggedPIF: Get the tagged_PIF field of the given VLAN. */
func (client *XenClient) VLANGetTaggedPIF(self string) (result string, err error) {
	obj, err := client.APICall("VLAN.get_tagged_PIF", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given VLAN. */
func (client *XenClient) VLANGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("VLAN.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the VLAN instance with the specified UUID. */
func (client *XenClient) VLANGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("VLAN.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given VLAN. */
func (client *XenClient) VLANGetRecord(self string) (result VLAN, err error) {
	obj, err := client.APICall("VLAN.get_record", self)
	if err != nil {
		return
	}

	result = *ToVLAN(obj.(interface{}))

	return
}

// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w HostPatch.gen.go

package go_xen_client

import (
	"reflect"
	"strconv"
	"time"

	"github.com/nilshell/xmlrpc"
)

//HostPatch: Represents a patch stored on a server
type HostPatch struct {
	Uuid             string            // Unique identifier/object reference
	NameLabel        string            // a human-readable name
	NameDescription  string            // a notes field containing human-readable description
	Version          string            // Patch version number
	Host             string            // Host the patch relates to
	Applied          bool              // True if the patch has been applied
	TimestampApplied time.Time         // Time the patch was applied
	Size             int               // Size of the patch
	PoolPatch        string            // The patch applied
	OtherConfig      map[string]string // additional configuration

}

func FromHostPatchToXml(host_patch *HostPatch) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] =

		host_patch.Uuid

	result["name_label"] =

		host_patch.NameLabel

	result["name_description"] =

		host_patch.NameDescription

	result["version"] =

		host_patch.Version

	result["host"] =

		host_patch.Host

	result["applied"] =

		host_patch.Applied

	result["timestamp_applied"] =

		host_patch.TimestampApplied

	result["size"] =

		strconv.Itoa(host_patch.Size)

	result["pool_patch"] =

		host_patch.PoolPatch

	other_config := make(xmlrpc.Struct)
	for key, value := range host_patch.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	return result
}

func ToHostPatch(obj interface{}) (resultObj *HostPatch) {

	objValue := reflect.ValueOf(obj)
	resultObj = &HostPatch{}

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
		case "version":
			if v, ok := keyValue.(string); ok {
				resultObj.Version = v
			}
		case "host":
			if v, ok := keyValue.(string); ok {
				resultObj.Host = v
			}
		case "applied":
			if v, ok := keyValue.(bool); ok {
				resultObj.Applied = v
			}
		case "timestamp_applied":
			if v, ok := keyValue.(time.Time); ok {
				resultObj.TimestampApplied = v
			}
		case "size":
			if v, ok := keyValue.(int); ok {
				resultObj.Size = v
			}
		case "pool_patch":
			if v, ok := keyValue.(string); ok {
				resultObj.PoolPatch = v
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

/* GetAllRecords: Return a map of host_patch references to host_patch records for all host_patchs known to the system. */
func (client *XenClient) HostPatchGetAllRecords() (result map[string]HostPatch, err error) {
	obj, err := client.APICall("host_patch.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]HostPatch{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToHostPatch(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the host_patchs known to the system. */
func (client *XenClient) HostPatchGetAll() (result []string, err error) {
	obj, err := client.APICall("host_patch.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* Apply: Apply the selected patch and return its output */
func (client *XenClient) HostPatchApply(self string) (result string, err error) {
	obj, err := client.APICall("host_patch.apply", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* Destroy: Destroy the specified host patch, removing it from the disk. This does NOT reverse the patch */
func (client *XenClient) HostPatchDestroy(self string) (err error) {
	_, err = client.APICall("host_patch.destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given host_patch.  If the key is not in that Map, then do nothing. */
func (client *XenClient) HostPatchRemoveFromOtherConfig(self string, key string) (err error) {
	_, err = client.APICall("host_patch.remove_from_other_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToOtherConfig: Add the given key-value pair to the other_config field of the given host_patch. */
func (client *XenClient) HostPatchAddToOtherConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("host_patch.add_to_other_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOtherConfig: Set the other_config field of the given host_patch. */
func (client *XenClient) HostPatchSetOtherConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("host_patch.set_other_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetOtherConfig: Get the other_config field of the given host_patch. */
func (client *XenClient) HostPatchGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("host_patch.get_other_config", self)
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

/* GetPoolPatch: Get the pool_patch field of the given host_patch. */
func (client *XenClient) HostPatchGetPoolPatch(self string) (result string, err error) {
	obj, err := client.APICall("host_patch.get_pool_patch", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetSize: Get the size field of the given host_patch. */
func (client *XenClient) HostPatchGetSize(self string) (result int, err error) {
	obj, err := client.APICall("host_patch.get_size", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetTimestampApplied: Get the timestamp_applied field of the given host_patch. */
func (client *XenClient) HostPatchGetTimestampApplied(self string) (result time.Time, err error) {
	obj, err := client.APICall("host_patch.get_timestamp_applied", self)
	if err != nil {
		return
	}
	result = obj.(time.Time)
	return
}

/* GetApplied: Get the applied field of the given host_patch. */
func (client *XenClient) HostPatchGetApplied(self string) (result bool, err error) {
	obj, err := client.APICall("host_patch.get_applied", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetHost: Get the host field of the given host_patch. */
func (client *XenClient) HostPatchGetHost(self string) (result string, err error) {
	obj, err := client.APICall("host_patch.get_host", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetVersion: Get the version field of the given host_patch. */
func (client *XenClient) HostPatchGetVersion(self string) (result string, err error) {
	obj, err := client.APICall("host_patch.get_version", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetNameDescription: Get the name/description field of the given host_patch. */
func (client *XenClient) HostPatchGetNameDescription(self string) (result string, err error) {
	obj, err := client.APICall("host_patch.get_name_description", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetNameLabel: Get the name/label field of the given host_patch. */
func (client *XenClient) HostPatchGetNameLabel(self string) (result string, err error) {
	obj, err := client.APICall("host_patch.get_name_label", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given host_patch. */
func (client *XenClient) HostPatchGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("host_patch.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByNameLabel: Get all the host_patch instances with the given label. */
func (client *XenClient) HostPatchGetByNameLabel(label string) (result []string, err error) {
	obj, err := client.APICall("host_patch.get_by_name_label", label)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetByUuid: Get a reference to the host_patch instance with the specified UUID. */
func (client *XenClient) HostPatchGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("host_patch.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given host_patch. */
func (client *XenClient) HostPatchGetRecord(self string) (result HostPatch, err error) {
	obj, err := client.APICall("host_patch.get_record", self)
	if err != nil {
		return
	}

	result = *ToHostPatch(obj.(interface{}))

	return
}

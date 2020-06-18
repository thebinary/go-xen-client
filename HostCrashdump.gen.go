// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w HostCrashdump.gen.go

package go_xen_client

import (
	"reflect"
	"strconv"
	"time"

	"github.com/nilshell/xmlrpc"
)

//HostCrashdump: Represents a host crash dump
type HostCrashdump struct {
	Uuid        string            // Unique identifier/object reference
	Host        string            // Host the crashdump relates to
	Timestamp   time.Time         // Time the crash happened
	Size        int               // Size of the crashdump
	OtherConfig map[string]string // additional configuration

}

func FromHostCrashdumpToXml(host_crashdump *HostCrashdump) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] =

		host_crashdump.Uuid

	result["host"] =

		host_crashdump.Host

	result["timestamp"] =

		host_crashdump.Timestamp

	result["size"] =

		strconv.Itoa(host_crashdump.Size)

	other_config := make(xmlrpc.Struct)
	for key, value := range host_crashdump.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	return result
}

func ToHostCrashdump(obj interface{}) (resultObj *HostCrashdump) {

	objValue := reflect.ValueOf(obj)
	resultObj = &HostCrashdump{}

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
		case "timestamp":
			if v, ok := keyValue.(time.Time); ok {
				resultObj.Timestamp = v
			}
		case "size":
			if v, ok := keyValue.(int); ok {
				resultObj.Size = v
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

/* GetAllRecords: Return a map of host_crashdump references to host_crashdump records for all host_crashdumps known to the system. */
func (client *XenClient) HostCrashdumpGetAllRecords() (result map[string]HostCrashdump, err error) {
	obj, err := client.APICall("host_crashdump.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]HostCrashdump{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToHostCrashdump(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the host_crashdumps known to the system. */
func (client *XenClient) HostCrashdumpGetAll() (result []string, err error) {
	obj, err := client.APICall("host_crashdump.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* Upload: Upload the specified host crash dump to a specified URL */
func (client *XenClient) HostCrashdumpUpload(self string, url string, options map[string]string) (err error) {
	_, err = client.APICall("host_crashdump.upload", self, url, options)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Destroy: Destroy specified host crash dump, removing it from the disk. */
func (client *XenClient) HostCrashdumpDestroy(self string) (err error) {
	_, err = client.APICall("host_crashdump.destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given host_crashdump.  If the key is not in that Map, then do nothing. */
func (client *XenClient) HostCrashdumpRemoveFromOtherConfig(self string, key string) (err error) {
	_, err = client.APICall("host_crashdump.remove_from_other_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToOtherConfig: Add the given key-value pair to the other_config field of the given host_crashdump. */
func (client *XenClient) HostCrashdumpAddToOtherConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("host_crashdump.add_to_other_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOtherConfig: Set the other_config field of the given host_crashdump. */
func (client *XenClient) HostCrashdumpSetOtherConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("host_crashdump.set_other_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetOtherConfig: Get the other_config field of the given host_crashdump. */
func (client *XenClient) HostCrashdumpGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("host_crashdump.get_other_config", self)
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

/* GetSize: Get the size field of the given host_crashdump. */
func (client *XenClient) HostCrashdumpGetSize(self string) (result int, err error) {
	obj, err := client.APICall("host_crashdump.get_size", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetTimestamp: Get the timestamp field of the given host_crashdump. */
func (client *XenClient) HostCrashdumpGetTimestamp(self string) (result time.Time, err error) {
	obj, err := client.APICall("host_crashdump.get_timestamp", self)
	if err != nil {
		return
	}
	result = obj.(time.Time)
	return
}

/* GetHost: Get the host field of the given host_crashdump. */
func (client *XenClient) HostCrashdumpGetHost(self string) (result string, err error) {
	obj, err := client.APICall("host_crashdump.get_host", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given host_crashdump. */
func (client *XenClient) HostCrashdumpGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("host_crashdump.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the host_crashdump instance with the specified UUID. */
func (client *XenClient) HostCrashdumpGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("host_crashdump.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given host_crashdump. */
func (client *XenClient) HostCrashdumpGetRecord(self string) (result HostCrashdump, err error) {
	obj, err := client.APICall("host_crashdump.get_record", self)
	if err != nil {
		return
	}

	result = *ToHostCrashdump(obj.(interface{}))

	return
}

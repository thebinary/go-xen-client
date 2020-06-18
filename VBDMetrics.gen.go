// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w VBDMetrics.gen.go

package go_xen_client

import (
	"reflect"
	"time"

	"github.com/nilshell/xmlrpc"
)

//VBDMetrics: The metrics associated with a virtual block device
type VBDMetrics struct {
	Uuid        string            // Unique identifier/object reference
	IoReadKbs   float32           // Read bandwidth (KiB/s)
	IoWriteKbs  float32           // Write bandwidth (KiB/s)
	LastUpdated time.Time         // Time at which this information was last updated
	OtherConfig map[string]string // additional configuration

}

func FromVBDMetricsToXml(VBD_metrics *VBDMetrics) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] =

		VBD_metrics.Uuid

	result["io_read_kbs"] =

		VBD_metrics.IoReadKbs

	result["io_write_kbs"] =

		VBD_metrics.IoWriteKbs

	result["last_updated"] =

		VBD_metrics.LastUpdated

	other_config := make(xmlrpc.Struct)
	for key, value := range VBD_metrics.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	return result
}

func ToVBDMetrics(obj interface{}) (resultObj *VBDMetrics) {

	objValue := reflect.ValueOf(obj)
	resultObj = &VBDMetrics{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.Uuid = v
			}
		case "io_read_kbs":
			if v, ok := keyValue.(float32); ok {
				resultObj.IoReadKbs = v
			}
		case "io_write_kbs":
			if v, ok := keyValue.(float32); ok {
				resultObj.IoWriteKbs = v
			}
		case "last_updated":
			if v, ok := keyValue.(time.Time); ok {
				resultObj.LastUpdated = v
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

/* GetAllRecords: Return a map of VBD_metrics references to VBD_metrics records for all VBD_metrics instances known to the system. */
func (client *XenClient) VBDMetricsGetAllRecords() (result map[string]VBDMetrics, err error) {
	obj, err := client.APICall("VBD_metrics.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]VBDMetrics{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToVBDMetrics(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the VBD_metrics instances known to the system. */
func (client *XenClient) VBDMetricsGetAll() (result []string, err error) {
	obj, err := client.APICall("VBD_metrics.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given VBD_metrics.  If the key is not in that Map, then do nothing. */
func (client *XenClient) VBDMetricsRemoveFromOtherConfig(self string, key string) (err error) {
	_, err = client.APICall("VBD_metrics.remove_from_other_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToOtherConfig: Add the given key-value pair to the other_config field of the given VBD_metrics. */
func (client *XenClient) VBDMetricsAddToOtherConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("VBD_metrics.add_to_other_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOtherConfig: Set the other_config field of the given VBD_metrics. */
func (client *XenClient) VBDMetricsSetOtherConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("VBD_metrics.set_other_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetOtherConfig: Get the other_config field of the given VBD_metrics. */
func (client *XenClient) VBDMetricsGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VBD_metrics.get_other_config", self)
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

/* GetLastUpdated: Get the last_updated field of the given VBD_metrics. */
func (client *XenClient) VBDMetricsGetLastUpdated(self string) (result time.Time, err error) {
	obj, err := client.APICall("VBD_metrics.get_last_updated", self)
	if err != nil {
		return
	}
	result = obj.(time.Time)
	return
}

/* GetIoWriteKbs: Get the io/write_kbs field of the given VBD_metrics. */
func (client *XenClient) VBDMetricsGetIoWriteKbs(self string) (result float32, err error) {
	obj, err := client.APICall("VBD_metrics.get_io_write_kbs", self)
	if err != nil {
		return
	}
	result = obj.(float32)
	return
}

/* GetIoReadKbs: Get the io/read_kbs field of the given VBD_metrics. */
func (client *XenClient) VBDMetricsGetIoReadKbs(self string) (result float32, err error) {
	obj, err := client.APICall("VBD_metrics.get_io_read_kbs", self)
	if err != nil {
		return
	}
	result = obj.(float32)
	return
}

/* GetUuid: Get the uuid field of the given VBD_metrics. */
func (client *XenClient) VBDMetricsGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("VBD_metrics.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the VBD_metrics instance with the specified UUID. */
func (client *XenClient) VBDMetricsGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("VBD_metrics.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given VBD_metrics. */
func (client *XenClient) VBDMetricsGetRecord(self string) (result VBDMetrics, err error) {
	obj, err := client.APICall("VBD_metrics.get_record", self)
	if err != nil {
		return
	}

	result = *ToVBDMetrics(obj.(interface{}))

	return
}

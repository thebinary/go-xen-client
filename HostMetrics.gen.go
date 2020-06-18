// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w HostMetrics.gen.go

package go_xen_client

import (
	"reflect"
	"strconv"
	"time"

	"github.com/nilshell/xmlrpc"
)

//HostMetrics: The metrics associated with a host
type HostMetrics struct {
	Uuid        string            // Unique identifier/object reference
	MemoryTotal int               // Total host memory (bytes)
	MemoryFree  int               // Free host memory (bytes)
	Live        bool              // Pool master thinks this host is live
	LastUpdated time.Time         // Time at which this information was last updated
	OtherConfig map[string]string // additional configuration

}

func FromHostMetricsToXml(host_metrics *HostMetrics) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] = host_metrics.Uuid

	result["memory_total"] = strconv.Itoa(host_metrics.MemoryTotal)

	result["memory_free"] = strconv.Itoa(host_metrics.MemoryFree)

	result["live"] = host_metrics.Live

	result["last_updated"] = host_metrics.LastUpdated

	other_config := make(xmlrpc.Struct)
	for key, value := range host_metrics.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	return result
}

func ToHostMetrics(obj interface{}) (resultObj *HostMetrics) {

	objValue := reflect.ValueOf(obj)
	resultObj = &HostMetrics{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.Uuid = v
			}
		case "memory_total":
			if v, ok := keyValue.(int); ok {
				resultObj.MemoryTotal = v
			}
		case "memory_free":
			if v, ok := keyValue.(int); ok {
				resultObj.MemoryFree = v
			}
		case "live":
			if v, ok := keyValue.(bool); ok {
				resultObj.Live = v
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

/* GetAllRecords: Return a map of host_metrics references to host_metrics records for all host_metrics instances known to the system. */
func (client *XenClient) HostMetricsGetAllRecords() (result map[string]HostMetrics, err error) {
	obj, err := client.APICall("host_metrics.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]HostMetrics{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToHostMetrics(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the host_metrics instances known to the system. */
func (client *XenClient) HostMetricsGetAll() (result []string, err error) {
	obj, err := client.APICall("host_metrics.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given host_metrics.  If the key is not in that Map, then do nothing. */
func (client *XenClient) HostMetricsRemoveFromOtherConfig(self string, key string) (err error) {
	_, err = client.APICall("host_metrics.remove_from_other_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToOtherConfig: Add the given key-value pair to the other_config field of the given host_metrics. */
func (client *XenClient) HostMetricsAddToOtherConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("host_metrics.add_to_other_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOtherConfig: Set the other_config field of the given host_metrics. */
func (client *XenClient) HostMetricsSetOtherConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("host_metrics.set_other_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetOtherConfig: Get the other_config field of the given host_metrics. */
func (client *XenClient) HostMetricsGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("host_metrics.get_other_config", self)
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

/* GetLastUpdated: Get the last_updated field of the given host_metrics. */
func (client *XenClient) HostMetricsGetLastUpdated(self string) (result time.Time, err error) {
	obj, err := client.APICall("host_metrics.get_last_updated", self)
	if err != nil {
		return
	}
	result = obj.(time.Time)
	return
}

/* GetLive: Get the live field of the given host_metrics. */
func (client *XenClient) HostMetricsGetLive(self string) (result bool, err error) {
	obj, err := client.APICall("host_metrics.get_live", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetMemoryFree: Get the memory/free field of the given host_metrics. */
func (client *XenClient) HostMetricsGetMemoryFree(self string) (result int, err error) {
	obj, err := client.APICall("host_metrics.get_memory_free", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetMemoryTotal: Get the memory/total field of the given host_metrics. */
func (client *XenClient) HostMetricsGetMemoryTotal(self string) (result int, err error) {
	obj, err := client.APICall("host_metrics.get_memory_total", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetUuid: Get the uuid field of the given host_metrics. */
func (client *XenClient) HostMetricsGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("host_metrics.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the host_metrics instance with the specified UUID. */
func (client *XenClient) HostMetricsGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("host_metrics.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given host_metrics. */
func (client *XenClient) HostMetricsGetRecord(self string) (result HostMetrics, err error) {
	obj, err := client.APICall("host_metrics.get_record", self)
	if err != nil {
		return
	}

	result = *ToHostMetrics(obj.(interface{}))

	return
}

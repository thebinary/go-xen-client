// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w HostCpu.gen.go

package go_xen_client

import (
	"reflect"
	"strconv"

	"github.com/nilshell/xmlrpc"
)

//HostCpu: A physical CPU
type HostCpu struct {
	Uuid        string            // Unique identifier/object reference
	Host        string            // the host the CPU is in
	Number      int               // the number of the physical CPU within the host
	Vendor      string            // the vendor of the physical CPU
	Speed       int               // the speed of the physical CPU
	Modelname   string            // the model name of the physical CPU
	Family      int               // the family (number) of the physical CPU
	Model       int               // the model number of the physical CPU
	Stepping    string            // the stepping of the physical CPU
	Flags       string            // the flags of the physical CPU (a decoded version of the features field)
	Features    string            // the physical CPU feature bitmap
	Utilisation float32           // the current CPU utilisation
	OtherConfig map[string]string // additional configuration

}

func FromHostCpuToXml(host_cpu *HostCpu) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] =

		host_cpu.Uuid

	result["host"] =

		host_cpu.Host

	result["number"] =

		strconv.Itoa(host_cpu.Number)

	result["vendor"] =

		host_cpu.Vendor

	result["speed"] =

		strconv.Itoa(host_cpu.Speed)

	result["modelname"] =

		host_cpu.Modelname

	result["family"] =

		strconv.Itoa(host_cpu.Family)

	result["model"] =

		strconv.Itoa(host_cpu.Model)

	result["stepping"] =

		host_cpu.Stepping

	result["flags"] =

		host_cpu.Flags

	result["features"] =

		host_cpu.Features

	result["utilisation"] =

		host_cpu.Utilisation

	other_config := make(xmlrpc.Struct)
	for key, value := range host_cpu.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	return result
}

func ToHostCpu(obj interface{}) (resultObj *HostCpu) {

	objValue := reflect.ValueOf(obj)
	resultObj = &HostCpu{}

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
		case "number":
			if v, ok := keyValue.(int); ok {
				resultObj.Number = v
			}
		case "vendor":
			if v, ok := keyValue.(string); ok {
				resultObj.Vendor = v
			}
		case "speed":
			if v, ok := keyValue.(int); ok {
				resultObj.Speed = v
			}
		case "modelname":
			if v, ok := keyValue.(string); ok {
				resultObj.Modelname = v
			}
		case "family":
			if v, ok := keyValue.(int); ok {
				resultObj.Family = v
			}
		case "model":
			if v, ok := keyValue.(int); ok {
				resultObj.Model = v
			}
		case "stepping":
			if v, ok := keyValue.(string); ok {
				resultObj.Stepping = v
			}
		case "flags":
			if v, ok := keyValue.(string); ok {
				resultObj.Flags = v
			}
		case "features":
			if v, ok := keyValue.(string); ok {
				resultObj.Features = v
			}
		case "utilisation":
			if v, ok := keyValue.(float32); ok {
				resultObj.Utilisation = v
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

/* GetAllRecords: Return a map of host_cpu references to host_cpu records for all host_cpus known to the system. */
func (client *XenClient) HostCpuGetAllRecords() (result map[string]HostCpu, err error) {
	obj, err := client.APICall("host_cpu.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]HostCpu{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToHostCpu(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the host_cpus known to the system. */
func (client *XenClient) HostCpuGetAll() (result []string, err error) {
	obj, err := client.APICall("host_cpu.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given host_cpu.  If the key is not in that Map, then do nothing. */
func (client *XenClient) HostCpuRemoveFromOtherConfig(self string, key string) (err error) {
	_, err = client.APICall("host_cpu.remove_from_other_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToOtherConfig: Add the given key-value pair to the other_config field of the given host_cpu. */
func (client *XenClient) HostCpuAddToOtherConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("host_cpu.add_to_other_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOtherConfig: Set the other_config field of the given host_cpu. */
func (client *XenClient) HostCpuSetOtherConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("host_cpu.set_other_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetOtherConfig: Get the other_config field of the given host_cpu. */
func (client *XenClient) HostCpuGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("host_cpu.get_other_config", self)
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

/* GetUtilisation: Get the utilisation field of the given host_cpu. */
func (client *XenClient) HostCpuGetUtilisation(self string) (result float32, err error) {
	obj, err := client.APICall("host_cpu.get_utilisation", self)
	if err != nil {
		return
	}
	result = obj.(float32)
	return
}

/* GetFeatures: Get the features field of the given host_cpu. */
func (client *XenClient) HostCpuGetFeatures(self string) (result string, err error) {
	obj, err := client.APICall("host_cpu.get_features", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetFlags: Get the flags field of the given host_cpu. */
func (client *XenClient) HostCpuGetFlags(self string) (result string, err error) {
	obj, err := client.APICall("host_cpu.get_flags", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetStepping: Get the stepping field of the given host_cpu. */
func (client *XenClient) HostCpuGetStepping(self string) (result string, err error) {
	obj, err := client.APICall("host_cpu.get_stepping", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetModel: Get the model field of the given host_cpu. */
func (client *XenClient) HostCpuGetModel(self string) (result int, err error) {
	obj, err := client.APICall("host_cpu.get_model", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetFamily: Get the family field of the given host_cpu. */
func (client *XenClient) HostCpuGetFamily(self string) (result int, err error) {
	obj, err := client.APICall("host_cpu.get_family", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetModelname: Get the modelname field of the given host_cpu. */
func (client *XenClient) HostCpuGetModelname(self string) (result string, err error) {
	obj, err := client.APICall("host_cpu.get_modelname", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetSpeed: Get the speed field of the given host_cpu. */
func (client *XenClient) HostCpuGetSpeed(self string) (result int, err error) {
	obj, err := client.APICall("host_cpu.get_speed", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetVendor: Get the vendor field of the given host_cpu. */
func (client *XenClient) HostCpuGetVendor(self string) (result string, err error) {
	obj, err := client.APICall("host_cpu.get_vendor", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetNumber: Get the number field of the given host_cpu. */
func (client *XenClient) HostCpuGetNumber(self string) (result int, err error) {
	obj, err := client.APICall("host_cpu.get_number", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetHost: Get the host field of the given host_cpu. */
func (client *XenClient) HostCpuGetHost(self string) (result string, err error) {
	obj, err := client.APICall("host_cpu.get_host", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given host_cpu. */
func (client *XenClient) HostCpuGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("host_cpu.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the host_cpu instance with the specified UUID. */
func (client *XenClient) HostCpuGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("host_cpu.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given host_cpu. */
func (client *XenClient) HostCpuGetRecord(self string) (result HostCpu, err error) {
	obj, err := client.APICall("host_cpu.get_record", self)
	if err != nil {
		return
	}

	result = *ToHostCpu(obj.(interface{}))

	return
}

// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w VMMetrics.gen.go

package go_xen_client

import (
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/nilshell/xmlrpc"
)

//VMMetrics: The metrics associated with a VM
type VMMetrics struct {
	Uuid              string            // Unique identifier/object reference
	MemoryActual      int               // Guest's actual memory (bytes)
	VCPUsNumber       int               // Current number of VCPUs
	VCPUsUtilisation  map[int]float32   // Utilisation for all of guest's current VCPUs
	VCPUsCPU          map[int]int       // VCPU to PCPU map
	VCPUsParams       map[string]string // The live equivalent to VM.VCPUs_params
	VCPUsFlags        map[int][]string  // CPU flags (blocked,online,running)
	State             []string          // The state of the guest, eg blocked, dying etc
	StartTime         time.Time         // Time at which this VM was last booted
	InstallTime       time.Time         // Time at which the VM was installed
	LastUpdated       time.Time         // Time at which this information was last updated
	OtherConfig       map[string]string // additional configuration
	Hvm               bool              // hardware virtual machine
	NestedVirt        bool              // VM supports nested virtualisation
	Nomigrate         bool              // VM is immobile and can't migrate between hosts
	CurrentDomainType DomainType        // The current domain type of the VM (for running,suspended, or paused VMs). The last-known domain type for halted VMs.

}

func FromVMMetricsToXml(VM_metrics *VMMetrics) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] = VM_metrics.Uuid

	result["memory_actual"] = strconv.Itoa(VM_metrics.MemoryActual)

	result["VCPUs_number"] = strconv.Itoa(VM_metrics.VCPUsNumber)

	VCPUs_utilisation := make(xmlrpc.Struct)
	for key, value := range VM_metrics.VCPUsUtilisation {
		VCPUs_utilisation[strconv.Itoa(key)] = value
	}
	result["VCPUs_utilisation"] = VCPUs_utilisation

	VCPUs_CPU := make(xmlrpc.Struct)
	for key, value := range VM_metrics.VCPUsCPU {
		VCPUs_CPU[strconv.Itoa(key)] = value
	}
	result["VCPUs_CPU"] = VCPUs_CPU

	VCPUs_params := make(xmlrpc.Struct)
	for key, value := range VM_metrics.VCPUsParams {
		VCPUs_params[key] = value
	}
	result["VCPUs_params"] = VCPUs_params

	VCPUs_flags := make(xmlrpc.Struct)
	for key, value := range VM_metrics.VCPUsFlags {
		VCPUs_flags[strconv.Itoa(key)] = value
	}
	result["VCPUs_flags"] = VCPUs_flags

	result["state"] = VM_metrics.State

	result["start_time"] = VM_metrics.StartTime

	result["install_time"] = VM_metrics.InstallTime

	result["last_updated"] = VM_metrics.LastUpdated

	other_config := make(xmlrpc.Struct)
	for key, value := range VM_metrics.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	result["hvm"] = VM_metrics.Hvm

	result["nested_virt"] = VM_metrics.NestedVirt

	result["nomigrate"] = VM_metrics.Nomigrate

	result["current_domain_type"] = VM_metrics.CurrentDomainType.String()

	return result
}

func ToVMMetrics(obj interface{}) (resultObj *VMMetrics) {

	objValue := reflect.ValueOf(obj)
	resultObj = &VMMetrics{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.Uuid = v
			}
		case "memory_actual":
			if v, ok := keyValue.(int); ok {
				resultObj.MemoryActual = v
			}
		case "VCPUs_number":
			if v, ok := keyValue.(int); ok {
				resultObj.VCPUsNumber = v
			}
		case "VCPUs_utilisation":

		case "VCPUs_CPU":

		case "VCPUs_params":

			resultObj.VCPUsParams = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.VCPUsParams[mapKeyName] = v
				} else {
					resultObj.VCPUsParams[mapKeyName] = ""
				}
			}
		case "VCPUs_flags":

		case "state":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.State = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.State[i] = v
					}
				}
			}
		case "start_time":
			if v, ok := keyValue.(time.Time); ok {
				resultObj.StartTime = v
			}
		case "install_time":
			if v, ok := keyValue.(time.Time); ok {
				resultObj.InstallTime = v
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
		case "hvm":
			if v, ok := keyValue.(bool); ok {
				resultObj.Hvm = v
			}
		case "nested_virt":
			if v, ok := keyValue.(bool); ok {
				resultObj.NestedVirt = v
			}
		case "nomigrate":
			if v, ok := keyValue.(bool); ok {
				resultObj.Nomigrate = v
			}
		case "current_domain_type":
			if v, ok := keyValue.(DomainType); ok {
				resultObj.CurrentDomainType = v
			}
		}
	}

	return resultObj
}

/* GetAllRecords: Return a map of VM_metrics references to VM_metrics records for all VM_metrics instances known to the system. */
func (client *XenClient) VMMetricsGetAllRecords() (result map[string]VMMetrics, err error) {
	obj, err := client.APICall("VM_metrics.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]VMMetrics{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToVMMetrics(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the VM_metrics instances known to the system. */
func (client *XenClient) VMMetricsGetAll() (result []string, err error) {
	obj, err := client.APICall("VM_metrics.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given VM_metrics.  If the key is not in that Map, then do nothing. */
func (client *XenClient) VMMetricsRemoveFromOtherConfig(self string, key string) (err error) {
	_, err = client.APICall("VM_metrics.remove_from_other_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToOtherConfig: Add the given key-value pair to the other_config field of the given VM_metrics. */
func (client *XenClient) VMMetricsAddToOtherConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("VM_metrics.add_to_other_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOtherConfig: Set the other_config field of the given VM_metrics. */
func (client *XenClient) VMMetricsSetOtherConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("VM_metrics.set_other_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetCurrentDomainType: Get the current_domain_type field of the given VM_metrics. */
func (client *XenClient) VMMetricsGetCurrentDomainType(self string) (result DomainType, err error) {
	obj, err := client.APICall("VM_metrics.get_current_domain_type", self)
	if err != nil {
		return
	}

	result = ToDomainType(obj.(string))

	return
}

/* GetNomigrate: Get the nomigrate field of the given VM_metrics. */
func (client *XenClient) VMMetricsGetNomigrate(self string) (result bool, err error) {
	obj, err := client.APICall("VM_metrics.get_nomigrate", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetNestedVirt: Get the nested_virt field of the given VM_metrics. */
func (client *XenClient) VMMetricsGetNestedVirt(self string) (result bool, err error) {
	obj, err := client.APICall("VM_metrics.get_nested_virt", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetHvm: Get the hvm field of the given VM_metrics. */
func (client *XenClient) VMMetricsGetHvm(self string) (result bool, err error) {
	obj, err := client.APICall("VM_metrics.get_hvm", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetOtherConfig: Get the other_config field of the given VM_metrics. */
func (client *XenClient) VMMetricsGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VM_metrics.get_other_config", self)
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

/* GetLastUpdated: Get the last_updated field of the given VM_metrics. */
func (client *XenClient) VMMetricsGetLastUpdated(self string) (result time.Time, err error) {
	obj, err := client.APICall("VM_metrics.get_last_updated", self)
	if err != nil {
		return
	}
	result = obj.(time.Time)
	return
}

/* GetInstallTime: Get the install_time field of the given VM_metrics. */
func (client *XenClient) VMMetricsGetInstallTime(self string) (result time.Time, err error) {
	obj, err := client.APICall("VM_metrics.get_install_time", self)
	if err != nil {
		return
	}
	result = obj.(time.Time)
	return
}

/* GetStartTime: Get the start_time field of the given VM_metrics. */
func (client *XenClient) VMMetricsGetStartTime(self string) (result time.Time, err error) {
	obj, err := client.APICall("VM_metrics.get_start_time", self)
	if err != nil {
		return
	}
	result = obj.(time.Time)
	return
}

/* GetState: Get the state field of the given VM_metrics. */
func (client *XenClient) VMMetricsGetState(self string) (result []string, err error) {
	obj, err := client.APICall("VM_metrics.get_state", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetVCPUsFlags: Get the VCPUs/flags field of the given VM_metrics. */
func (client *XenClient) VMMetricsGetVCPUsFlags(self string) (result map[int][]string, err error) {
	obj, err := client.APICall("VM_metrics.get_VCPUs_flags", self)
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[int][]string{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		interimObj := obj.Interface().([]interface{})
		interimResult := make([]string, len(interimObj))
		for i, interimValue := range interimObj {
			interimResult[i] = interimValue.(string)
		}
		result[int(key.Int())] = interimResult
	}

	return
}

/* GetVCPUsParams: Get the VCPUs/params field of the given VM_metrics. */
func (client *XenClient) VMMetricsGetVCPUsParams(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VM_metrics.get_VCPUs_params", self)
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

/* GetVCPUsCPU: Get the VCPUs/CPU field of the given VM_metrics. */
func (client *XenClient) VMMetricsGetVCPUsCPU(self string) (result map[int]int, err error) {
	obj, err := client.APICall("VM_metrics.get_VCPUs_CPU", self)
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[int]int{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		result[int(key.Int())] = int(obj.Int())
	}

	return
}

/* GetVCPUsUtilisation: Get the VCPUs/utilisation field of the given VM_metrics. */
func (client *XenClient) VMMetricsGetVCPUsUtilisation(self string) (result map[int]float32, err error) {
	obj, err := client.APICall("VM_metrics.get_VCPUs_utilisation", self)
	if err != nil {
		return
	}
	//result conversion not implemented yet
	log.Printf("%+v", obj)
	return
}

/* GetVCPUsNumber: Get the VCPUs/number field of the given VM_metrics. */
func (client *XenClient) VMMetricsGetVCPUsNumber(self string) (result int, err error) {
	obj, err := client.APICall("VM_metrics.get_VCPUs_number", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetMemoryActual: Get the memory/actual field of the given VM_metrics. */
func (client *XenClient) VMMetricsGetMemoryActual(self string) (result int, err error) {
	obj, err := client.APICall("VM_metrics.get_memory_actual", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetUuid: Get the uuid field of the given VM_metrics. */
func (client *XenClient) VMMetricsGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("VM_metrics.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the VM_metrics instance with the specified UUID. */
func (client *XenClient) VMMetricsGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("VM_metrics.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given VM_metrics. */
func (client *XenClient) VMMetricsGetRecord(self string) (result VMMetrics, err error) {
	obj, err := client.APICall("VM_metrics.get_record", self)
	if err != nil {
		return
	}

	result = *ToVMMetrics(obj.(interface{}))

	return
}

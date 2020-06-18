// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w VMGuestMetrics.gen.go

package go_xen_client

import (
	"reflect"
	"time"

	"github.com/nilshell/xmlrpc"
)

//VMGuestMetrics: The metrics reported by the guest (as opposed to inferred from outside)
type VMGuestMetrics struct {
	Uuid              string            // Unique identifier/object reference
	OsVersion         map[string]string // version of the OS
	PVDriversVersion  map[string]string // version of the PV drivers
	PVDriversUpToDate bool              // Logically equivalent to PV_drivers_detected
	Memory            map[string]string // This field exists but has no data. Use the memory and memory_internal_free RRD data-sources instead.
	Disks             map[string]string // This field exists but has no data.
	Networks          map[string]string // network configuration
	Other             map[string]string // anything else
	LastUpdated       time.Time         // Time at which this information was last updated
	OtherConfig       map[string]string // additional configuration
	Live              bool              // True if the guest is sending heartbeat messages via the guest agent
	CanUseHotplugVbd  TristateType      // The guest's statement of whether it supports VBD hotplug, i.e. whether it is capable of responding immediately to instantiation of a new VBD by bringing online a new PV block device. If the guest states that it is not capable, then the VBD plug and unplug operations will not be allowed while the guest is running.
	CanUseHotplugVif  TristateType      // The guest's statement of whether it supports VIF hotplug, i.e. whether it is capable of responding immediately to instantiation of a new VIF by bringing online a new PV network device. If the guest states that it is not capable, then the VIF plug and unplug operations will not be allowed while the guest is running.
	PVDriversDetected bool              // At least one of the guest's devices has successfully connected to the backend.

}

func FromVMGuestMetricsToXml(VM_guest_metrics *VMGuestMetrics) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] =

		VM_guest_metrics.Uuid

	os_version := make(xmlrpc.Struct)
	for key, value := range VM_guest_metrics.OsVersion {
		os_version[key] = value
	}
	result["os_version"] = os_version

	PV_drivers_version := make(xmlrpc.Struct)
	for key, value := range VM_guest_metrics.PVDriversVersion {
		PV_drivers_version[key] = value
	}
	result["PV_drivers_version"] = PV_drivers_version

	result["PV_drivers_up_to_date"] =

		VM_guest_metrics.PVDriversUpToDate

	memory := make(xmlrpc.Struct)
	for key, value := range VM_guest_metrics.Memory {
		memory[key] = value
	}
	result["memory"] = memory

	disks := make(xmlrpc.Struct)
	for key, value := range VM_guest_metrics.Disks {
		disks[key] = value
	}
	result["disks"] = disks

	networks := make(xmlrpc.Struct)
	for key, value := range VM_guest_metrics.Networks {
		networks[key] = value
	}
	result["networks"] = networks

	other := make(xmlrpc.Struct)
	for key, value := range VM_guest_metrics.Other {
		other[key] = value
	}
	result["other"] = other

	result["last_updated"] =

		VM_guest_metrics.LastUpdated

	other_config := make(xmlrpc.Struct)
	for key, value := range VM_guest_metrics.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	result["live"] =

		VM_guest_metrics.Live

	result["can_use_hotplug_vbd"] =

		VM_guest_metrics.CanUseHotplugVbd.String()

	result["can_use_hotplug_vif"] =

		VM_guest_metrics.CanUseHotplugVif.String()

	result["PV_drivers_detected"] =

		VM_guest_metrics.PVDriversDetected

	return result
}

func ToVMGuestMetrics(obj interface{}) (resultObj *VMGuestMetrics) {

	objValue := reflect.ValueOf(obj)
	resultObj = &VMGuestMetrics{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.Uuid = v
			}
		case "os_version":

			resultObj.OsVersion = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.OsVersion[mapKeyName] = v
				} else {
					resultObj.OsVersion[mapKeyName] = ""
				}
			}
		case "PV_drivers_version":

			resultObj.PVDriversVersion = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.PVDriversVersion[mapKeyName] = v
				} else {
					resultObj.PVDriversVersion[mapKeyName] = ""
				}
			}
		case "PV_drivers_up_to_date":
			if v, ok := keyValue.(bool); ok {
				resultObj.PVDriversUpToDate = v
			}
		case "memory":

			resultObj.Memory = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.Memory[mapKeyName] = v
				} else {
					resultObj.Memory[mapKeyName] = ""
				}
			}
		case "disks":

			resultObj.Disks = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.Disks[mapKeyName] = v
				} else {
					resultObj.Disks[mapKeyName] = ""
				}
			}
		case "networks":

			resultObj.Networks = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.Networks[mapKeyName] = v
				} else {
					resultObj.Networks[mapKeyName] = ""
				}
			}
		case "other":

			resultObj.Other = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.Other[mapKeyName] = v
				} else {
					resultObj.Other[mapKeyName] = ""
				}
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
		case "live":
			if v, ok := keyValue.(bool); ok {
				resultObj.Live = v
			}
		case "can_use_hotplug_vbd":
			if v, ok := keyValue.(TristateType); ok {
				resultObj.CanUseHotplugVbd = v
			}
		case "can_use_hotplug_vif":
			if v, ok := keyValue.(TristateType); ok {
				resultObj.CanUseHotplugVif = v
			}
		case "PV_drivers_detected":
			if v, ok := keyValue.(bool); ok {
				resultObj.PVDriversDetected = v
			}
		}
	}

	return resultObj
}

/* GetAllRecords: Return a map of VM_guest_metrics references to VM_guest_metrics records for all VM_guest_metrics instances known to the system. */
func (client *XenClient) VMGuestMetricsGetAllRecords() (result map[string]VMGuestMetrics, err error) {
	obj, err := client.APICall("VM_guest_metrics.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]VMGuestMetrics{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToVMGuestMetrics(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the VM_guest_metrics instances known to the system. */
func (client *XenClient) VMGuestMetricsGetAll() (result []string, err error) {
	obj, err := client.APICall("VM_guest_metrics.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given VM_guest_metrics.  If the key is not in that Map, then do nothing. */
func (client *XenClient) VMGuestMetricsRemoveFromOtherConfig(self string, key string) (err error) {
	_, err = client.APICall("VM_guest_metrics.remove_from_other_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToOtherConfig: Add the given key-value pair to the other_config field of the given VM_guest_metrics. */
func (client *XenClient) VMGuestMetricsAddToOtherConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("VM_guest_metrics.add_to_other_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOtherConfig: Set the other_config field of the given VM_guest_metrics. */
func (client *XenClient) VMGuestMetricsSetOtherConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("VM_guest_metrics.set_other_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetPVDriversDetected: Get the PV_drivers_detected field of the given VM_guest_metrics. */
func (client *XenClient) VMGuestMetricsGetPVDriversDetected(self string) (result bool, err error) {
	obj, err := client.APICall("VM_guest_metrics.get_PV_drivers_detected", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetCanUseHotplugVif: Get the can_use_hotplug_vif field of the given VM_guest_metrics. */
func (client *XenClient) VMGuestMetricsGetCanUseHotplugVif(self string) (result TristateType, err error) {
	obj, err := client.APICall("VM_guest_metrics.get_can_use_hotplug_vif", self)
	if err != nil {
		return
	}

	result = ToTristateType(obj.(string))

	return
}

/* GetCanUseHotplugVbd: Get the can_use_hotplug_vbd field of the given VM_guest_metrics. */
func (client *XenClient) VMGuestMetricsGetCanUseHotplugVbd(self string) (result TristateType, err error) {
	obj, err := client.APICall("VM_guest_metrics.get_can_use_hotplug_vbd", self)
	if err != nil {
		return
	}

	result = ToTristateType(obj.(string))

	return
}

/* GetLive: Get the live field of the given VM_guest_metrics. */
func (client *XenClient) VMGuestMetricsGetLive(self string) (result bool, err error) {
	obj, err := client.APICall("VM_guest_metrics.get_live", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetOtherConfig: Get the other_config field of the given VM_guest_metrics. */
func (client *XenClient) VMGuestMetricsGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VM_guest_metrics.get_other_config", self)
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

/* GetLastUpdated: Get the last_updated field of the given VM_guest_metrics. */
func (client *XenClient) VMGuestMetricsGetLastUpdated(self string) (result time.Time, err error) {
	obj, err := client.APICall("VM_guest_metrics.get_last_updated", self)
	if err != nil {
		return
	}
	result = obj.(time.Time)
	return
}

/* GetOther: Get the other field of the given VM_guest_metrics. */
func (client *XenClient) VMGuestMetricsGetOther(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VM_guest_metrics.get_other", self)
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

/* GetNetworks: Get the networks field of the given VM_guest_metrics. */
func (client *XenClient) VMGuestMetricsGetNetworks(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VM_guest_metrics.get_networks", self)
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

/* GetDisks: Get the disks field of the given VM_guest_metrics. */
func (client *XenClient) VMGuestMetricsGetDisks(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VM_guest_metrics.get_disks", self)
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

/* GetMemory: Get the memory field of the given VM_guest_metrics. */
func (client *XenClient) VMGuestMetricsGetMemory(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VM_guest_metrics.get_memory", self)
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

/* GetPVDriversUpToDate: Get the PV_drivers_up_to_date field of the given VM_guest_metrics. */
func (client *XenClient) VMGuestMetricsGetPVDriversUpToDate(self string) (result bool, err error) {
	obj, err := client.APICall("VM_guest_metrics.get_PV_drivers_up_to_date", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetPVDriversVersion: Get the PV_drivers_version field of the given VM_guest_metrics. */
func (client *XenClient) VMGuestMetricsGetPVDriversVersion(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VM_guest_metrics.get_PV_drivers_version", self)
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

/* GetOsVersion: Get the os_version field of the given VM_guest_metrics. */
func (client *XenClient) VMGuestMetricsGetOsVersion(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VM_guest_metrics.get_os_version", self)
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

/* GetUuid: Get the uuid field of the given VM_guest_metrics. */
func (client *XenClient) VMGuestMetricsGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("VM_guest_metrics.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the VM_guest_metrics instance with the specified UUID. */
func (client *XenClient) VMGuestMetricsGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("VM_guest_metrics.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given VM_guest_metrics. */
func (client *XenClient) VMGuestMetricsGetRecord(self string) (result VMGuestMetrics, err error) {
	obj, err := client.APICall("VM_guest_metrics.get_record", self)
	if err != nil {
		return
	}

	result = *ToVMGuestMetrics(obj.(interface{}))

	return
}

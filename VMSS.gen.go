// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w VMSS.gen.go

package go_xen_client

import (
	"reflect"
	"strconv"
	"time"

	"github.com/nilshell/xmlrpc"
)

//VMSS: VM Snapshot Schedule
type VMSS struct {
	Uuid              string            // Unique identifier/object reference
	NameLabel         string            // a human-readable name
	NameDescription   string            // a notes field containing human-readable description
	Enabled           bool              // enable or disable this snapshot schedule
	Type              VmssType          // type of the snapshot schedule
	RetainedSnapshots int               // maximum number of snapshots that should be stored at any time
	Frequency         VmssFrequency     // frequency of taking snapshot from snapshot schedule
	Schedule          map[string]string // schedule of the snapshot containing 'hour', 'min', 'days'. Date/time-related information is in Local Timezone
	LastRunTime       time.Time         // time of the last snapshot
	VMs               []string          // all VMs attached to this snapshot schedule

}

func FromVMSSToXml(VMSS *VMSS) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] = VMSS.Uuid

	result["name_label"] = VMSS.NameLabel

	result["name_description"] = VMSS.NameDescription

	result["enabled"] = VMSS.Enabled

	result["type"] = VMSS.Type.String()

	result["retained_snapshots"] = strconv.Itoa(VMSS.RetainedSnapshots)

	result["frequency"] = VMSS.Frequency.String()

	schedule := make(xmlrpc.Struct)
	for key, value := range VMSS.Schedule {
		schedule[key] = value
	}
	result["schedule"] = schedule

	result["last_run_time"] = VMSS.LastRunTime

	result["VMs"] = VMSS.VMs

	return result
}

func ToVMSS(obj interface{}) (resultObj *VMSS) {

	objValue := reflect.ValueOf(obj)
	resultObj = &VMSS{}

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
		case "enabled":
			if v, ok := keyValue.(bool); ok {
				resultObj.Enabled = v
			}
		case "type":
			if v, ok := keyValue.(VmssType); ok {
				resultObj.Type = v
			}
		case "retained_snapshots":
			if v, ok := keyValue.(int); ok {
				resultObj.RetainedSnapshots = v
			}
		case "frequency":
			if v, ok := keyValue.(VmssFrequency); ok {
				resultObj.Frequency = v
			}
		case "schedule":

			resultObj.Schedule = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.Schedule[mapKeyName] = v
				} else {
					resultObj.Schedule[mapKeyName] = ""
				}
			}
		case "last_run_time":
			if v, ok := keyValue.(time.Time); ok {
				resultObj.LastRunTime = v
			}
		case "VMs":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.VMs = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.VMs[i] = v
					}
				}
			}
		}
	}

	return resultObj
}

/* GetAllRecords: Return a map of VMSS references to VMSS records for all VMSSs known to the system. */
func (client *XenClient) VMSSGetAllRecords() (result map[string]VMSS, err error) {
	obj, err := client.APICall("VMSS.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]VMSS{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToVMSS(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the VMSSs known to the system. */
func (client *XenClient) VMSSGetAll() (result []string, err error) {
	obj, err := client.APICall("VMSS.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* SetType:  */
func (client *XenClient) VMSSSetType(self string, value VmssType) (err error) {
	_, err = client.APICall("VMSS.set_type", self, value.String())
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetLastRunTime:  */
func (client *XenClient) VMSSSetLastRunTime(self string, value time.Time) (err error) {
	_, err = client.APICall("VMSS.set_last_run_time", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveFromSchedule:  */
func (client *XenClient) VMSSRemoveFromSchedule(self string, key string) (err error) {
	_, err = client.APICall("VMSS.remove_from_schedule", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToSchedule:  */
func (client *XenClient) VMSSAddToSchedule(self string, key string, value string) (err error) {
	_, err = client.APICall("VMSS.add_to_schedule", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetSchedule:  */
func (client *XenClient) VMSSSetSchedule(self string, value map[string]string) (err error) {
	_, err = client.APICall("VMSS.set_schedule", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetFrequency: Set the value of the frequency field */
func (client *XenClient) VMSSSetFrequency(self string, value VmssFrequency) (err error) {
	_, err = client.APICall("VMSS.set_frequency", self, value.String())
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetRetainedSnapshots:  */
func (client *XenClient) VMSSSetRetainedSnapshots(self string, value int) (err error) {
	_, err = client.APICall("VMSS.set_retained_snapshots", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SnapshotNow: This call executes the snapshot schedule immediately */
func (client *XenClient) VMSSSnapshotNow(vmss string) (result string, err error) {
	obj, err := client.APICall("VMSS.snapshot_now", vmss)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* SetEnabled: Set the enabled field of the given VMSS. */
func (client *XenClient) VMSSSetEnabled(self string, value bool) (err error) {
	_, err = client.APICall("VMSS.set_enabled", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetNameDescription: Set the name/description field of the given VMSS. */
func (client *XenClient) VMSSSetNameDescription(self string, value string) (err error) {
	_, err = client.APICall("VMSS.set_name_description", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetNameLabel: Set the name/label field of the given VMSS. */
func (client *XenClient) VMSSSetNameLabel(self string, value string) (err error) {
	_, err = client.APICall("VMSS.set_name_label", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetVMs: Get the VMs field of the given VMSS. */
func (client *XenClient) VMSSGetVMs(self string) (result []string, err error) {
	obj, err := client.APICall("VMSS.get_VMs", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetLastRunTime: Get the last_run_time field of the given VMSS. */
func (client *XenClient) VMSSGetLastRunTime(self string) (result time.Time, err error) {
	obj, err := client.APICall("VMSS.get_last_run_time", self)
	if err != nil {
		return
	}
	result = obj.(time.Time)
	return
}

/* GetSchedule: Get the schedule field of the given VMSS. */
func (client *XenClient) VMSSGetSchedule(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VMSS.get_schedule", self)
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

/* GetFrequency: Get the frequency field of the given VMSS. */
func (client *XenClient) VMSSGetFrequency(self string) (result VmssFrequency, err error) {
	obj, err := client.APICall("VMSS.get_frequency", self)
	if err != nil {
		return
	}

	result = ToVmssFrequency(obj.(string))

	return
}

/* GetRetainedSnapshots: Get the retained_snapshots field of the given VMSS. */
func (client *XenClient) VMSSGetRetainedSnapshots(self string) (result int, err error) {
	obj, err := client.APICall("VMSS.get_retained_snapshots", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetType: Get the type field of the given VMSS. */
func (client *XenClient) VMSSGetType(self string) (result VmssType, err error) {
	obj, err := client.APICall("VMSS.get_type", self)
	if err != nil {
		return
	}

	result = ToVmssType(obj.(string))

	return
}

/* GetEnabled: Get the enabled field of the given VMSS. */
func (client *XenClient) VMSSGetEnabled(self string) (result bool, err error) {
	obj, err := client.APICall("VMSS.get_enabled", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetNameDescription: Get the name/description field of the given VMSS. */
func (client *XenClient) VMSSGetNameDescription(self string) (result string, err error) {
	obj, err := client.APICall("VMSS.get_name_description", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetNameLabel: Get the name/label field of the given VMSS. */
func (client *XenClient) VMSSGetNameLabel(self string) (result string, err error) {
	obj, err := client.APICall("VMSS.get_name_label", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given VMSS. */
func (client *XenClient) VMSSGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("VMSS.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByNameLabel: Get all the VMSS instances with the given label. */
func (client *XenClient) VMSSGetByNameLabel(label string) (result []string, err error) {
	obj, err := client.APICall("VMSS.get_by_name_label", label)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* Destroy: Destroy the specified VMSS instance. */
func (client *XenClient) VMSSDestroy(self string) (err error) {
	_, err = client.APICall("VMSS.destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Create: Create a new VMSS instance, and return its handle.
The constructor args are: name_label, name_description, enabled, type*, retained_snapshots, frequency*, schedule (* = non-optional). */
func (client *XenClient) VMSSCreate(args VMSS) (result string, err error) {
	obj, err := client.APICall("VMSS.create", FromVMSSToXml(&args))
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the VMSS instance with the specified UUID. */
func (client *XenClient) VMSSGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("VMSS.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given VMSS. */
func (client *XenClient) VMSSGetRecord(self string) (result VMSS, err error) {
	obj, err := client.APICall("VMSS.get_record", self)
	if err != nil {
		return
	}

	result = *ToVMSS(obj.(interface{}))

	return
}

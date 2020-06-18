// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w VBD.gen.go

package go_xen_client

import (
	"reflect"
	"strconv"

	"github.com/nilshell/xmlrpc"
)

//VBD: A virtual block device
type VBD struct {
	Uuid                   string                   // Unique identifier/object reference
	AllowedOperations      []VbdOperations          // list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	CurrentOperations      map[string]VbdOperations // links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	VM                     string                   // the virtual machine
	VDI                    string                   // the virtual disk
	Device                 string                   // device seen by the guest e.g. hda1
	Userdevice             string                   // user-friendly device name e.g. 0,1,2,etc.
	Bootable               bool                     // true if this VBD is bootable
	Mode                   VbdMode                  // the mode the VBD should be mounted with
	Type                   VbdType                  // how the VBD will appear to the guest (e.g. disk or CD)
	Unpluggable            bool                     // true if this VBD will support hot-unplug
	StorageLock            bool                     // true if a storage level lock was acquired
	Empty                  bool                     // if true this represents an empty drive
	OtherConfig            map[string]string        // additional configuration
	CurrentlyAttached      bool                     // is the device currently attached (erased on reboot)
	StatusCode             int                      // error/success code associated with last attach-operation (erased on reboot)
	StatusDetail           string                   // error/success information associated with last attach-operation status (erased on reboot)
	RuntimeProperties      map[string]string        // Device runtime properties
	QosAlgorithmType       string                   // QoS algorithm to use
	QosAlgorithmParams     map[string]string        // parameters for chosen QoS algorithm
	QosSupportedAlgorithms []string                 // supported QoS algorithms for this VBD
	Metrics                string                   // metrics associated with this VBD

}

func FromVBDToXml(VBD *VBD) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] =

		VBD.Uuid

	result["allowed_operations"] =

		VBD.AllowedOperations

	current_operations := make(xmlrpc.Struct)
	for key, value := range VBD.CurrentOperations {
		current_operations[key] = value
	}
	result["current_operations"] = current_operations

	result["VM"] =

		VBD.VM

	result["VDI"] =

		VBD.VDI

	result["device"] =

		VBD.Device

	result["userdevice"] =

		VBD.Userdevice

	result["bootable"] =

		VBD.Bootable

	result["mode"] =

		VBD.Mode.String()

	result["type"] =

		VBD.Type.String()

	result["unpluggable"] =

		VBD.Unpluggable

	result["storage_lock"] =

		VBD.StorageLock

	result["empty"] =

		VBD.Empty

	other_config := make(xmlrpc.Struct)
	for key, value := range VBD.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	result["currently_attached"] =

		VBD.CurrentlyAttached

	result["status_code"] =

		strconv.Itoa(VBD.StatusCode)

	result["status_detail"] =

		VBD.StatusDetail

	runtime_properties := make(xmlrpc.Struct)
	for key, value := range VBD.RuntimeProperties {
		runtime_properties[key] = value
	}
	result["runtime_properties"] = runtime_properties

	result["qos_algorithm_type"] =

		VBD.QosAlgorithmType

	qos_algorithm_params := make(xmlrpc.Struct)
	for key, value := range VBD.QosAlgorithmParams {
		qos_algorithm_params[key] = value
	}
	result["qos_algorithm_params"] = qos_algorithm_params

	result["qos_supported_algorithms"] =

		VBD.QosSupportedAlgorithms

	result["metrics"] =

		VBD.Metrics

	return result
}

func ToVBD(obj interface{}) (resultObj *VBD) {

	objValue := reflect.ValueOf(obj)
	resultObj = &VBD{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.Uuid = v
			}
		case "allowed_operations":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.AllowedOperations = make([]VbdOperations, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(VbdOperations); ok {
						resultObj.AllowedOperations[i] = v
					}
				}
			}
		case "current_operations":

			resultObj.CurrentOperations = map[string]VbdOperations{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.CurrentOperations[mapKeyName] = ToVbdOperations(v)
				} else {
					resultObj.CurrentOperations[mapKeyName] = 0
				}
			}
		case "VM":
			if v, ok := keyValue.(string); ok {
				resultObj.VM = v
			}
		case "VDI":
			if v, ok := keyValue.(string); ok {
				resultObj.VDI = v
			}
		case "device":
			if v, ok := keyValue.(string); ok {
				resultObj.Device = v
			}
		case "userdevice":
			if v, ok := keyValue.(string); ok {
				resultObj.Userdevice = v
			}
		case "bootable":
			if v, ok := keyValue.(bool); ok {
				resultObj.Bootable = v
			}
		case "mode":
			if v, ok := keyValue.(VbdMode); ok {
				resultObj.Mode = v
			}
		case "type":
			if v, ok := keyValue.(VbdType); ok {
				resultObj.Type = v
			}
		case "unpluggable":
			if v, ok := keyValue.(bool); ok {
				resultObj.Unpluggable = v
			}
		case "storage_lock":
			if v, ok := keyValue.(bool); ok {
				resultObj.StorageLock = v
			}
		case "empty":
			if v, ok := keyValue.(bool); ok {
				resultObj.Empty = v
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
		case "currently_attached":
			if v, ok := keyValue.(bool); ok {
				resultObj.CurrentlyAttached = v
			}
		case "status_code":
			if v, ok := keyValue.(int); ok {
				resultObj.StatusCode = v
			}
		case "status_detail":
			if v, ok := keyValue.(string); ok {
				resultObj.StatusDetail = v
			}
		case "runtime_properties":

			resultObj.RuntimeProperties = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.RuntimeProperties[mapKeyName] = v
				} else {
					resultObj.RuntimeProperties[mapKeyName] = ""
				}
			}
		case "qos_algorithm_type":
			if v, ok := keyValue.(string); ok {
				resultObj.QosAlgorithmType = v
			}
		case "qos_algorithm_params":

			resultObj.QosAlgorithmParams = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.QosAlgorithmParams[mapKeyName] = v
				} else {
					resultObj.QosAlgorithmParams[mapKeyName] = ""
				}
			}
		case "qos_supported_algorithms":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.QosSupportedAlgorithms = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.QosSupportedAlgorithms[i] = v
					}
				}
			}
		case "metrics":
			if v, ok := keyValue.(string); ok {
				resultObj.Metrics = v
			}
		}
	}

	return resultObj
}

/* GetAllRecords: Return a map of VBD references to VBD records for all VBDs known to the system. */
func (client *XenClient) VBDGetAllRecords() (result map[string]VBD, err error) {
	obj, err := client.APICall("VBD.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]VBD{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToVBD(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the VBDs known to the system. */
func (client *XenClient) VBDGetAll() (result []string, err error) {
	obj, err := client.APICall("VBD.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* SetMode: Sets the mode of the VBD. The power_state of the VM must be halted. */
func (client *XenClient) VBDSetMode(self string, value VbdMode) (err error) {
	_, err = client.APICall("VBD.set_mode", self, value.String())
	if err != nil {
		return
	}
	// no return result
	return
}

/* AssertAttachable: Throws an error if this VBD could not be attached to this VM if the VM were running. Intended for debugging. */
func (client *XenClient) VBDAssertAttachable(self string) (err error) {
	_, err = client.APICall("VBD.assert_attachable", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* UnplugForce: Forcibly unplug the specified VBD */
func (client *XenClient) VBDUnplugForce(self string) (err error) {
	_, err = client.APICall("VBD.unplug_force", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Unplug: Hot-unplug the specified VBD, dynamically unattaching it from the running VM */
func (client *XenClient) VBDUnplug(self string) (err error) {
	_, err = client.APICall("VBD.unplug", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Plug: Hotplug the specified VBD, dynamically attaching it to the running VM */
func (client *XenClient) VBDPlug(self string) (err error) {
	_, err = client.APICall("VBD.plug", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Insert: Insert new media into the device */
func (client *XenClient) VBDInsert(vbd string, vdi string) (err error) {
	_, err = client.APICall("VBD.insert", vbd, vdi)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Eject: Remove the media from the device and leave it empty */
func (client *XenClient) VBDEject(vbd string) (err error) {
	_, err = client.APICall("VBD.eject", vbd)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveFromQosAlgorithmParams: Remove the given key and its corresponding value from the qos/algorithm_params field of the given VBD.  If the key is not in that Map, then do nothing. */
func (client *XenClient) VBDRemoveFromQosAlgorithmParams(self string, key string) (err error) {
	_, err = client.APICall("VBD.remove_from_qos_algorithm_params", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToQosAlgorithmParams: Add the given key-value pair to the qos/algorithm_params field of the given VBD. */
func (client *XenClient) VBDAddToQosAlgorithmParams(self string, key string, value string) (err error) {
	_, err = client.APICall("VBD.add_to_qos_algorithm_params", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetQosAlgorithmParams: Set the qos/algorithm_params field of the given VBD. */
func (client *XenClient) VBDSetQosAlgorithmParams(self string, value map[string]string) (err error) {
	_, err = client.APICall("VBD.set_qos_algorithm_params", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetQosAlgorithmType: Set the qos/algorithm_type field of the given VBD. */
func (client *XenClient) VBDSetQosAlgorithmType(self string, value string) (err error) {
	_, err = client.APICall("VBD.set_qos_algorithm_type", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given VBD.  If the key is not in that Map, then do nothing. */
func (client *XenClient) VBDRemoveFromOtherConfig(self string, key string) (err error) {
	_, err = client.APICall("VBD.remove_from_other_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToOtherConfig: Add the given key-value pair to the other_config field of the given VBD. */
func (client *XenClient) VBDAddToOtherConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("VBD.add_to_other_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOtherConfig: Set the other_config field of the given VBD. */
func (client *XenClient) VBDSetOtherConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("VBD.set_other_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetUnpluggable: Set the unpluggable field of the given VBD. */
func (client *XenClient) VBDSetUnpluggable(self string, value bool) (err error) {
	_, err = client.APICall("VBD.set_unpluggable", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetType: Set the type field of the given VBD. */
func (client *XenClient) VBDSetType(self string, value VbdType) (err error) {
	_, err = client.APICall("VBD.set_type", self, value.String())
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetBootable: Set the bootable field of the given VBD. */
func (client *XenClient) VBDSetBootable(self string, value bool) (err error) {
	_, err = client.APICall("VBD.set_bootable", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetUserdevice: Set the userdevice field of the given VBD. */
func (client *XenClient) VBDSetUserdevice(self string, value string) (err error) {
	_, err = client.APICall("VBD.set_userdevice", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetMetrics: Get the metrics field of the given VBD. */
func (client *XenClient) VBDGetMetrics(self string) (result string, err error) {
	obj, err := client.APICall("VBD.get_metrics", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetQosSupportedAlgorithms: Get the qos/supported_algorithms field of the given VBD. */
func (client *XenClient) VBDGetQosSupportedAlgorithms(self string) (result []string, err error) {
	obj, err := client.APICall("VBD.get_qos_supported_algorithms", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetQosAlgorithmParams: Get the qos/algorithm_params field of the given VBD. */
func (client *XenClient) VBDGetQosAlgorithmParams(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VBD.get_qos_algorithm_params", self)
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

/* GetQosAlgorithmType: Get the qos/algorithm_type field of the given VBD. */
func (client *XenClient) VBDGetQosAlgorithmType(self string) (result string, err error) {
	obj, err := client.APICall("VBD.get_qos_algorithm_type", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRuntimeProperties: Get the runtime_properties field of the given VBD. */
func (client *XenClient) VBDGetRuntimeProperties(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VBD.get_runtime_properties", self)
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

/* GetStatusDetail: Get the status_detail field of the given VBD. */
func (client *XenClient) VBDGetStatusDetail(self string) (result string, err error) {
	obj, err := client.APICall("VBD.get_status_detail", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetStatusCode: Get the status_code field of the given VBD. */
func (client *XenClient) VBDGetStatusCode(self string) (result int, err error) {
	obj, err := client.APICall("VBD.get_status_code", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetCurrentlyAttached: Get the currently_attached field of the given VBD. */
func (client *XenClient) VBDGetCurrentlyAttached(self string) (result bool, err error) {
	obj, err := client.APICall("VBD.get_currently_attached", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetOtherConfig: Get the other_config field of the given VBD. */
func (client *XenClient) VBDGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VBD.get_other_config", self)
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

/* GetEmpty: Get the empty field of the given VBD. */
func (client *XenClient) VBDGetEmpty(self string) (result bool, err error) {
	obj, err := client.APICall("VBD.get_empty", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetStorageLock: Get the storage_lock field of the given VBD. */
func (client *XenClient) VBDGetStorageLock(self string) (result bool, err error) {
	obj, err := client.APICall("VBD.get_storage_lock", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetUnpluggable: Get the unpluggable field of the given VBD. */
func (client *XenClient) VBDGetUnpluggable(self string) (result bool, err error) {
	obj, err := client.APICall("VBD.get_unpluggable", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetType: Get the type field of the given VBD. */
func (client *XenClient) VBDGetType(self string) (result VbdType, err error) {
	obj, err := client.APICall("VBD.get_type", self)
	if err != nil {
		return
	}

	result = ToVbdType(obj.(string))

	return
}

/* GetMode: Get the mode field of the given VBD. */
func (client *XenClient) VBDGetMode(self string) (result VbdMode, err error) {
	obj, err := client.APICall("VBD.get_mode", self)
	if err != nil {
		return
	}

	result = ToVbdMode(obj.(string))

	return
}

/* GetBootable: Get the bootable field of the given VBD. */
func (client *XenClient) VBDGetBootable(self string) (result bool, err error) {
	obj, err := client.APICall("VBD.get_bootable", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetUserdevice: Get the userdevice field of the given VBD. */
func (client *XenClient) VBDGetUserdevice(self string) (result string, err error) {
	obj, err := client.APICall("VBD.get_userdevice", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetDevice: Get the device field of the given VBD. */
func (client *XenClient) VBDGetDevice(self string) (result string, err error) {
	obj, err := client.APICall("VBD.get_device", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetVDI: Get the VDI field of the given VBD. */
func (client *XenClient) VBDGetVDI(self string) (result string, err error) {
	obj, err := client.APICall("VBD.get_VDI", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetVM: Get the VM field of the given VBD. */
func (client *XenClient) VBDGetVM(self string) (result string, err error) {
	obj, err := client.APICall("VBD.get_VM", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetCurrentOperations: Get the current_operations field of the given VBD. */
func (client *XenClient) VBDGetCurrentOperations(self string) (result map[string]VbdOperations, err error) {
	obj, err := client.APICall("VBD.get_current_operations", self)
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]VbdOperations{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToVbdOperations(obj.String())
		result[key.String()] = mapObj
	}

	return
}

/* GetAllowedOperations: Get the allowed_operations field of the given VBD. */
func (client *XenClient) VBDGetAllowedOperations(self string) (result []VbdOperations, err error) {
	obj, err := client.APICall("VBD.get_allowed_operations", self)
	if err != nil {
		return
	}

	result = make([]VbdOperations, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = ToVbdOperations(value.(string))
	}

	return
}

/* GetUuid: Get the uuid field of the given VBD. */
func (client *XenClient) VBDGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("VBD.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* Destroy: Destroy the specified VBD instance. */
func (client *XenClient) VBDDestroy(self string) (err error) {
	_, err = client.APICall("VBD.destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Create: Create a new VBD instance, and return its handle.
The constructor args are: VM*, VDI*, userdevice*, bootable*, mode*, type*, unpluggable, empty*, other_config*, qos_algorithm_type*, qos_algorithm_params* (* = non-optional). */
func (client *XenClient) VBDCreate(args VBD) (result string, err error) {
	obj, err := client.APICall("VBD.create", FromVBDToXml(&args))
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the VBD instance with the specified UUID. */
func (client *XenClient) VBDGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("VBD.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given VBD. */
func (client *XenClient) VBDGetRecord(self string) (result VBD, err error) {
	obj, err := client.APICall("VBD.get_record", self)
	if err != nil {
		return
	}

	result = *ToVBD(obj.(interface{}))

	return
}

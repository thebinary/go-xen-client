// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w VGPU.gen.go

package go_xen_client

import (
	"reflect"

	"github.com/nilshell/xmlrpc"
)

//VGPU: A virtual GPU (vGPU)
type VGPU struct {
	Uuid                    string            // Unique identifier/object reference
	VM                      string            // VM that owns the vGPU
	GPUGroup                string            // GPU group used by the vGPU
	Device                  string            // Guest PCI slot (a value of 0 means auto-assign to first empty slot, valid slot is in range of [11,31] for multi-VGPU purpose)
	CurrentlyAttached       bool              // Reflects whether the virtual device is currently connected to a physical device
	OtherConfig             map[string]string // Additional configuration
	Type                    string            // Preset type for this VGPU
	ResidentOn              string            // The PGPU on which this VGPU is running
	ScheduledToBeResidentOn string            // The PGPU on which this VGPU is scheduled to run
	CompatibilityMetadata   map[string]string // VGPU metadata to determine whether a VGPU can migrate between two PGPUs
	ExtraArgs               string            // Extra arguments for vGPU and passed to demu

}

func FromVGPUToXml(VGPU *VGPU) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] =

		VGPU.Uuid

	result["VM"] =

		VGPU.VM

	result["GPU_group"] =

		VGPU.GPUGroup

	result["device"] =

		VGPU.Device

	result["currently_attached"] =

		VGPU.CurrentlyAttached

	other_config := make(xmlrpc.Struct)
	for key, value := range VGPU.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	result["type"] =

		VGPU.Type

	result["resident_on"] =

		VGPU.ResidentOn

	result["scheduled_to_be_resident_on"] =

		VGPU.ScheduledToBeResidentOn

	compatibility_metadata := make(xmlrpc.Struct)
	for key, value := range VGPU.CompatibilityMetadata {
		compatibility_metadata[key] = value
	}
	result["compatibility_metadata"] = compatibility_metadata

	result["extra_args"] =

		VGPU.ExtraArgs

	return result
}

func ToVGPU(obj interface{}) (resultObj *VGPU) {

	objValue := reflect.ValueOf(obj)
	resultObj = &VGPU{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.Uuid = v
			}
		case "VM":
			if v, ok := keyValue.(string); ok {
				resultObj.VM = v
			}
		case "GPU_group":
			if v, ok := keyValue.(string); ok {
				resultObj.GPUGroup = v
			}
		case "device":
			if v, ok := keyValue.(string); ok {
				resultObj.Device = v
			}
		case "currently_attached":
			if v, ok := keyValue.(bool); ok {
				resultObj.CurrentlyAttached = v
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
		case "type":
			if v, ok := keyValue.(string); ok {
				resultObj.Type = v
			}
		case "resident_on":
			if v, ok := keyValue.(string); ok {
				resultObj.ResidentOn = v
			}
		case "scheduled_to_be_resident_on":
			if v, ok := keyValue.(string); ok {
				resultObj.ScheduledToBeResidentOn = v
			}
		case "compatibility_metadata":

			resultObj.CompatibilityMetadata = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.CompatibilityMetadata[mapKeyName] = v
				} else {
					resultObj.CompatibilityMetadata[mapKeyName] = ""
				}
			}
		case "extra_args":
			if v, ok := keyValue.(string); ok {
				resultObj.ExtraArgs = v
			}
		}
	}

	return resultObj
}

/* GetAllRecords: Return a map of VGPU references to VGPU records for all VGPUs known to the system. */
func (client *XenClient) VGPUGetAllRecords() (result map[string]VGPU, err error) {
	obj, err := client.APICall("VGPU.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]VGPU{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToVGPU(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the VGPUs known to the system. */
func (client *XenClient) VGPUGetAll() (result []string, err error) {
	obj, err := client.APICall("VGPU.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* Destroy:  */
func (client *XenClient) VGPUDestroy(self string) (err error) {
	_, err = client.APICall("VGPU.destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Create:  */
func (client *XenClient) VGPUCreate(VM string, GPU_group string, device string, other_config map[string]string, xtype string) (result string, err error) {
	obj, err := client.APICall("VGPU.create", VM, GPU_group, device, other_config, xtype)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* SetExtraArgs: Set the extra_args field of the given VGPU. */
func (client *XenClient) VGPUSetExtraArgs(self string, value string) (err error) {
	_, err = client.APICall("VGPU.set_extra_args", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given VGPU.  If the key is not in that Map, then do nothing. */
func (client *XenClient) VGPURemoveFromOtherConfig(self string, key string) (err error) {
	_, err = client.APICall("VGPU.remove_from_other_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToOtherConfig: Add the given key-value pair to the other_config field of the given VGPU. */
func (client *XenClient) VGPUAddToOtherConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("VGPU.add_to_other_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOtherConfig: Set the other_config field of the given VGPU. */
func (client *XenClient) VGPUSetOtherConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("VGPU.set_other_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetExtraArgs: Get the extra_args field of the given VGPU. */
func (client *XenClient) VGPUGetExtraArgs(self string) (result string, err error) {
	obj, err := client.APICall("VGPU.get_extra_args", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetCompatibilityMetadata: Get the compatibility_metadata field of the given VGPU. */
func (client *XenClient) VGPUGetCompatibilityMetadata(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VGPU.get_compatibility_metadata", self)
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

/* GetScheduledToBeResidentOn: Get the scheduled_to_be_resident_on field of the given VGPU. */
func (client *XenClient) VGPUGetScheduledToBeResidentOn(self string) (result string, err error) {
	obj, err := client.APICall("VGPU.get_scheduled_to_be_resident_on", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetResidentOn: Get the resident_on field of the given VGPU. */
func (client *XenClient) VGPUGetResidentOn(self string) (result string, err error) {
	obj, err := client.APICall("VGPU.get_resident_on", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetType: Get the type field of the given VGPU. */
func (client *XenClient) VGPUGetType(self string) (result string, err error) {
	obj, err := client.APICall("VGPU.get_type", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetOtherConfig: Get the other_config field of the given VGPU. */
func (client *XenClient) VGPUGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VGPU.get_other_config", self)
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

/* GetCurrentlyAttached: Get the currently_attached field of the given VGPU. */
func (client *XenClient) VGPUGetCurrentlyAttached(self string) (result bool, err error) {
	obj, err := client.APICall("VGPU.get_currently_attached", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetDevice: Get the device field of the given VGPU. */
func (client *XenClient) VGPUGetDevice(self string) (result string, err error) {
	obj, err := client.APICall("VGPU.get_device", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetGPUGroup: Get the GPU_group field of the given VGPU. */
func (client *XenClient) VGPUGetGPUGroup(self string) (result string, err error) {
	obj, err := client.APICall("VGPU.get_GPU_group", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetVM: Get the VM field of the given VGPU. */
func (client *XenClient) VGPUGetVM(self string) (result string, err error) {
	obj, err := client.APICall("VGPU.get_VM", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given VGPU. */
func (client *XenClient) VGPUGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("VGPU.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the VGPU instance with the specified UUID. */
func (client *XenClient) VGPUGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("VGPU.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given VGPU. */
func (client *XenClient) VGPUGetRecord(self string) (result VGPU, err error) {
	obj, err := client.APICall("VGPU.get_record", self)
	if err != nil {
		return
	}

	result = *ToVGPU(obj.(interface{}))

	return
}

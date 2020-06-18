// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w GPUGroup.gen.go

package go_xen_client

import (
	"reflect"

	"github.com/nilshell/xmlrpc"
)

//GPUGroup: A group of compatible GPUs across the resource pool
type GPUGroup struct {
	Uuid                string              // Unique identifier/object reference
	NameLabel           string              // a human-readable name
	NameDescription     string              // a notes field containing human-readable description
	PGPUs               []string            // List of pGPUs in the group
	VGPUs               []string            // List of vGPUs using the group
	GPUTypes            []string            // List of GPU types (vendor+device ID) that can be in this group
	OtherConfig         map[string]string   // Additional configuration
	AllocationAlgorithm AllocationAlgorithm // Current allocation of vGPUs to pGPUs for this group
	SupportedVGPUTypes  []string            // vGPU types supported on at least one of the pGPUs in this group
	EnabledVGPUTypes    []string            // vGPU types supported on at least one of the pGPUs in this group

}

func FromGPUGroupToXml(GPU_group *GPUGroup) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] = GPU_group.Uuid

	result["name_label"] = GPU_group.NameLabel

	result["name_description"] = GPU_group.NameDescription

	result["PGPUs"] = GPU_group.PGPUs

	result["VGPUs"] = GPU_group.VGPUs

	result["GPU_types"] = GPU_group.GPUTypes

	other_config := make(xmlrpc.Struct)
	for key, value := range GPU_group.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	result["allocation_algorithm"] = GPU_group.AllocationAlgorithm.String()

	result["supported_VGPU_types"] = GPU_group.SupportedVGPUTypes

	result["enabled_VGPU_types"] = GPU_group.EnabledVGPUTypes

	return result
}

func ToGPUGroup(obj interface{}) (resultObj *GPUGroup) {

	objValue := reflect.ValueOf(obj)
	resultObj = &GPUGroup{}

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
		case "PGPUs":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.PGPUs = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.PGPUs[i] = v
					}
				}
			}
		case "VGPUs":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.VGPUs = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.VGPUs[i] = v
					}
				}
			}
		case "GPU_types":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.GPUTypes = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.GPUTypes[i] = v
					}
				}
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
		case "allocation_algorithm":
			if v, ok := keyValue.(AllocationAlgorithm); ok {
				resultObj.AllocationAlgorithm = v
			}
		case "supported_VGPU_types":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.SupportedVGPUTypes = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.SupportedVGPUTypes[i] = v
					}
				}
			}
		case "enabled_VGPU_types":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.EnabledVGPUTypes = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.EnabledVGPUTypes[i] = v
					}
				}
			}
		}
	}

	return resultObj
}

/* GetAllRecords: Return a map of GPU_group references to GPU_group records for all GPU_groups known to the system. */
func (client *XenClient) GPUGroupGetAllRecords() (result map[string]GPUGroup, err error) {
	obj, err := client.APICall("GPU_group.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]GPUGroup{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToGPUGroup(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the GPU_groups known to the system. */
func (client *XenClient) GPUGroupGetAll() (result []string, err error) {
	obj, err := client.APICall("GPU_group.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetRemainingCapacity:  */
func (client *XenClient) GPUGroupGetRemainingCapacity(self string, vgpu_type string) (result int, err error) {
	obj, err := client.APICall("GPU_group.get_remaining_capacity", self, vgpu_type)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* Destroy:  */
func (client *XenClient) GPUGroupDestroy(self string) (err error) {
	_, err = client.APICall("GPU_group.destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Create:  */
func (client *XenClient) GPUGroupCreate(name_label string, name_description string, other_config map[string]string) (result string, err error) {
	obj, err := client.APICall("GPU_group.create", name_label, name_description, other_config)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* SetAllocationAlgorithm: Set the allocation_algorithm field of the given GPU_group. */
func (client *XenClient) GPUGroupSetAllocationAlgorithm(self string, value AllocationAlgorithm) (err error) {
	_, err = client.APICall("GPU_group.set_allocation_algorithm", self, value.String())
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given GPU_group.  If the key is not in that Map, then do nothing. */
func (client *XenClient) GPUGroupRemoveFromOtherConfig(self string, key string) (err error) {
	_, err = client.APICall("GPU_group.remove_from_other_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToOtherConfig: Add the given key-value pair to the other_config field of the given GPU_group. */
func (client *XenClient) GPUGroupAddToOtherConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("GPU_group.add_to_other_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOtherConfig: Set the other_config field of the given GPU_group. */
func (client *XenClient) GPUGroupSetOtherConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("GPU_group.set_other_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetNameDescription: Set the name/description field of the given GPU_group. */
func (client *XenClient) GPUGroupSetNameDescription(self string, value string) (err error) {
	_, err = client.APICall("GPU_group.set_name_description", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetNameLabel: Set the name/label field of the given GPU_group. */
func (client *XenClient) GPUGroupSetNameLabel(self string, value string) (err error) {
	_, err = client.APICall("GPU_group.set_name_label", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetEnabledVGPUTypes: Get the enabled_VGPU_types field of the given GPU_group. */
func (client *XenClient) GPUGroupGetEnabledVGPUTypes(self string) (result []string, err error) {
	obj, err := client.APICall("GPU_group.get_enabled_VGPU_types", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetSupportedVGPUTypes: Get the supported_VGPU_types field of the given GPU_group. */
func (client *XenClient) GPUGroupGetSupportedVGPUTypes(self string) (result []string, err error) {
	obj, err := client.APICall("GPU_group.get_supported_VGPU_types", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetAllocationAlgorithm: Get the allocation_algorithm field of the given GPU_group. */
func (client *XenClient) GPUGroupGetAllocationAlgorithm(self string) (result AllocationAlgorithm, err error) {
	obj, err := client.APICall("GPU_group.get_allocation_algorithm", self)
	if err != nil {
		return
	}

	result = ToAllocationAlgorithm(obj.(string))

	return
}

/* GetOtherConfig: Get the other_config field of the given GPU_group. */
func (client *XenClient) GPUGroupGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("GPU_group.get_other_config", self)
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

/* GetGPUTypes: Get the GPU_types field of the given GPU_group. */
func (client *XenClient) GPUGroupGetGPUTypes(self string) (result []string, err error) {
	obj, err := client.APICall("GPU_group.get_GPU_types", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetVGPUs: Get the VGPUs field of the given GPU_group. */
func (client *XenClient) GPUGroupGetVGPUs(self string) (result []string, err error) {
	obj, err := client.APICall("GPU_group.get_VGPUs", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetPGPUs: Get the PGPUs field of the given GPU_group. */
func (client *XenClient) GPUGroupGetPGPUs(self string) (result []string, err error) {
	obj, err := client.APICall("GPU_group.get_PGPUs", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetNameDescription: Get the name/description field of the given GPU_group. */
func (client *XenClient) GPUGroupGetNameDescription(self string) (result string, err error) {
	obj, err := client.APICall("GPU_group.get_name_description", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetNameLabel: Get the name/label field of the given GPU_group. */
func (client *XenClient) GPUGroupGetNameLabel(self string) (result string, err error) {
	obj, err := client.APICall("GPU_group.get_name_label", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given GPU_group. */
func (client *XenClient) GPUGroupGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("GPU_group.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByNameLabel: Get all the GPU_group instances with the given label. */
func (client *XenClient) GPUGroupGetByNameLabel(label string) (result []string, err error) {
	obj, err := client.APICall("GPU_group.get_by_name_label", label)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetByUuid: Get a reference to the GPU_group instance with the specified UUID. */
func (client *XenClient) GPUGroupGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("GPU_group.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given GPU_group. */
func (client *XenClient) GPUGroupGetRecord(self string) (result GPUGroup, err error) {
	obj, err := client.APICall("GPU_group.get_record", self)
	if err != nil {
		return
	}

	result = *ToGPUGroup(obj.(interface{}))

	return
}

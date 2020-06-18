// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w PGPU.gen.go

package go_xen_client

import (
	"reflect"

	"github.com/nilshell/xmlrpc"
)

//PGPU: A physical GPU (pGPU)
type PGPU struct {
	Uuid                       string            // Unique identifier/object reference
	PCI                        string            // Link to underlying PCI device
	GPUGroup                   string            // GPU group the pGPU is contained in
	Host                       string            // Host that owns the GPU
	OtherConfig                map[string]string // Additional configuration
	SupportedVGPUTypes         []string          // List of VGPU types supported by the underlying hardware
	EnabledVGPUTypes           []string          // List of VGPU types which have been enabled for this PGPU
	ResidentVGPUs              []string          // List of VGPUs running on this PGPU
	SupportedVGPUMaxCapacities map[string]int    // A map relating each VGPU type supported on this GPU to the maximum number of VGPUs of that type which can run simultaneously on this GPU
	Dom0Access                 PgpuDom0Access    // The accessibility of this device from dom0
	IsSystemDisplayDevice      bool              // Is this device the system display device
	CompatibilityMetadata      map[string]string // PGPU metadata to determine whether a VGPU can migrate between two PGPUs

}

func FromPGPUToXml(PGPU *PGPU) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] = PGPU.Uuid

	result["PCI"] = PGPU.PCI

	result["GPU_group"] = PGPU.GPUGroup

	result["host"] = PGPU.Host

	other_config := make(xmlrpc.Struct)
	for key, value := range PGPU.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	result["supported_VGPU_types"] = PGPU.SupportedVGPUTypes

	result["enabled_VGPU_types"] = PGPU.EnabledVGPUTypes

	result["resident_VGPUs"] = PGPU.ResidentVGPUs

	supported_VGPU_max_capacities := make(xmlrpc.Struct)
	for key, value := range PGPU.SupportedVGPUMaxCapacities {
		supported_VGPU_max_capacities[key] = value
	}
	result["supported_VGPU_max_capacities"] = supported_VGPU_max_capacities

	result["dom0_access"] = PGPU.Dom0Access.String()

	result["is_system_display_device"] = PGPU.IsSystemDisplayDevice

	compatibility_metadata := make(xmlrpc.Struct)
	for key, value := range PGPU.CompatibilityMetadata {
		compatibility_metadata[key] = value
	}
	result["compatibility_metadata"] = compatibility_metadata

	return result
}

func ToPGPU(obj interface{}) (resultObj *PGPU) {

	objValue := reflect.ValueOf(obj)
	resultObj = &PGPU{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.Uuid = v
			}
		case "PCI":
			if v, ok := keyValue.(string); ok {
				resultObj.PCI = v
			}
		case "GPU_group":
			if v, ok := keyValue.(string); ok {
				resultObj.GPUGroup = v
			}
		case "host":
			if v, ok := keyValue.(string); ok {
				resultObj.Host = v
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
		case "resident_VGPUs":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.ResidentVGPUs = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.ResidentVGPUs[i] = v
					}
				}
			}
		case "supported_VGPU_max_capacities":

			resultObj.SupportedVGPUMaxCapacities = map[string]int{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(int); ok {
					resultObj.SupportedVGPUMaxCapacities[mapKeyName] = v
				} else {
					resultObj.SupportedVGPUMaxCapacities[mapKeyName] = 0
				}
			}
		case "dom0_access":
			if v, ok := keyValue.(PgpuDom0Access); ok {
				resultObj.Dom0Access = v
			}
		case "is_system_display_device":
			if v, ok := keyValue.(bool); ok {
				resultObj.IsSystemDisplayDevice = v
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
		}
	}

	return resultObj
}

/* GetAllRecords: Return a map of PGPU references to PGPU records for all PGPUs known to the system. */
func (client *XenClient) PGPUGetAllRecords() (result map[string]PGPU, err error) {
	obj, err := client.APICall("PGPU.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]PGPU{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToPGPU(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the PGPUs known to the system. */
func (client *XenClient) PGPUGetAll() (result []string, err error) {
	obj, err := client.APICall("PGPU.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* DisableDom0Access:  */
func (client *XenClient) PGPUDisableDom0Access(self string) (result PgpuDom0Access, err error) {
	obj, err := client.APICall("PGPU.disable_dom0_access", self)
	if err != nil {
		return
	}

	result = ToPgpuDom0Access(obj.(string))

	return
}

/* EnableDom0Access:  */
func (client *XenClient) PGPUEnableDom0Access(self string) (result PgpuDom0Access, err error) {
	obj, err := client.APICall("PGPU.enable_dom0_access", self)
	if err != nil {
		return
	}

	result = ToPgpuDom0Access(obj.(string))

	return
}

/* GetRemainingCapacity:  */
func (client *XenClient) PGPUGetRemainingCapacity(self string, vgpu_type string) (result int, err error) {
	obj, err := client.APICall("PGPU.get_remaining_capacity", self, vgpu_type)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* SetGPUGroup:  */
func (client *XenClient) PGPUSetGPUGroup(self string, value string) (err error) {
	_, err = client.APICall("PGPU.set_GPU_group", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetEnabledVGPUTypes:  */
func (client *XenClient) PGPUSetEnabledVGPUTypes(self string, value []string) (err error) {
	_, err = client.APICall("PGPU.set_enabled_VGPU_types", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveEnabledVGPUTypes:  */
func (client *XenClient) PGPURemoveEnabledVGPUTypes(self string, value string) (err error) {
	_, err = client.APICall("PGPU.remove_enabled_VGPU_types", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddEnabledVGPUTypes:  */
func (client *XenClient) PGPUAddEnabledVGPUTypes(self string, value string) (err error) {
	_, err = client.APICall("PGPU.add_enabled_VGPU_types", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given PGPU.  If the key is not in that Map, then do nothing. */
func (client *XenClient) PGPURemoveFromOtherConfig(self string, key string) (err error) {
	_, err = client.APICall("PGPU.remove_from_other_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToOtherConfig: Add the given key-value pair to the other_config field of the given PGPU. */
func (client *XenClient) PGPUAddToOtherConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("PGPU.add_to_other_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOtherConfig: Set the other_config field of the given PGPU. */
func (client *XenClient) PGPUSetOtherConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("PGPU.set_other_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetCompatibilityMetadata: Get the compatibility_metadata field of the given PGPU. */
func (client *XenClient) PGPUGetCompatibilityMetadata(self string) (result map[string]string, err error) {
	obj, err := client.APICall("PGPU.get_compatibility_metadata", self)
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

/* GetIsSystemDisplayDevice: Get the is_system_display_device field of the given PGPU. */
func (client *XenClient) PGPUGetIsSystemDisplayDevice(self string) (result bool, err error) {
	obj, err := client.APICall("PGPU.get_is_system_display_device", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetDom0Access: Get the dom0_access field of the given PGPU. */
func (client *XenClient) PGPUGetDom0Access(self string) (result PgpuDom0Access, err error) {
	obj, err := client.APICall("PGPU.get_dom0_access", self)
	if err != nil {
		return
	}

	result = ToPgpuDom0Access(obj.(string))

	return
}

/* GetSupportedVGPUMaxCapacities: Get the supported_VGPU_max_capacities field of the given PGPU. */
func (client *XenClient) PGPUGetSupportedVGPUMaxCapacities(self string) (result map[string]int, err error) {
	obj, err := client.APICall("PGPU.get_supported_VGPU_max_capacities", self)
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]int{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		result[key.String()] = int(obj.Int())
	}

	return
}

/* GetResidentVGPUs: Get the resident_VGPUs field of the given PGPU. */
func (client *XenClient) PGPUGetResidentVGPUs(self string) (result []string, err error) {
	obj, err := client.APICall("PGPU.get_resident_VGPUs", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetEnabledVGPUTypes: Get the enabled_VGPU_types field of the given PGPU. */
func (client *XenClient) PGPUGetEnabledVGPUTypes(self string) (result []string, err error) {
	obj, err := client.APICall("PGPU.get_enabled_VGPU_types", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetSupportedVGPUTypes: Get the supported_VGPU_types field of the given PGPU. */
func (client *XenClient) PGPUGetSupportedVGPUTypes(self string) (result []string, err error) {
	obj, err := client.APICall("PGPU.get_supported_VGPU_types", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetOtherConfig: Get the other_config field of the given PGPU. */
func (client *XenClient) PGPUGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("PGPU.get_other_config", self)
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

/* GetHost: Get the host field of the given PGPU. */
func (client *XenClient) PGPUGetHost(self string) (result string, err error) {
	obj, err := client.APICall("PGPU.get_host", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetGPUGroup: Get the GPU_group field of the given PGPU. */
func (client *XenClient) PGPUGetGPUGroup(self string) (result string, err error) {
	obj, err := client.APICall("PGPU.get_GPU_group", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetPCI: Get the PCI field of the given PGPU. */
func (client *XenClient) PGPUGetPCI(self string) (result string, err error) {
	obj, err := client.APICall("PGPU.get_PCI", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given PGPU. */
func (client *XenClient) PGPUGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("PGPU.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the PGPU instance with the specified UUID. */
func (client *XenClient) PGPUGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("PGPU.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given PGPU. */
func (client *XenClient) PGPUGetRecord(self string) (result PGPU, err error) {
	obj, err := client.APICall("PGPU.get_record", self)
	if err != nil {
		return
	}

	result = *ToPGPU(obj.(interface{}))

	return
}

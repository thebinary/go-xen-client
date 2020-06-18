// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w VGPUType.gen.go

package go_xen_client

import (
	"reflect"
	"strconv"

	"github.com/nilshell/xmlrpc"
)

//VGPUType: A type of virtual GPU
type VGPUType struct {
	Uuid                 string                 // Unique identifier/object reference
	VendorName           string                 // Name of VGPU vendor
	ModelName            string                 // Model name associated with the VGPU type
	FramebufferSize      int                    // Framebuffer size of the VGPU type, in bytes
	MaxHeads             int                    // Maximum number of displays supported by the VGPU type
	MaxResolutionX       int                    // Maximum resolution (width) supported by the VGPU type
	MaxResolutionY       int                    // Maximum resolution (height) supported by the VGPU type
	SupportedOnPGPUs     []string               // List of PGPUs that support this VGPU type
	EnabledOnPGPUs       []string               // List of PGPUs that have this VGPU type enabled
	VGPUs                []string               // List of VGPUs of this type
	SupportedOnGPUGroups []string               // List of GPU groups in which at least one PGPU supports this VGPU type
	EnabledOnGPUGroups   []string               // List of GPU groups in which at least one have this VGPU type enabled
	Implementation       VgpuTypeImplementation // The internal implementation of this VGPU type
	Identifier           string                 // Key used to identify VGPU types and avoid creating duplicates - this field is used internally and not intended for interpretation by API clients
	Experimental         bool                   // Indicates whether VGPUs of this type should be considered experimental
	CompatibleTypesInVm  []string               // List of VGPU types which are compatible in one VM

}

func FromVGPUTypeToXml(VGPU_type *VGPUType) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] = VGPU_type.Uuid

	result["vendor_name"] = VGPU_type.VendorName

	result["model_name"] = VGPU_type.ModelName

	result["framebuffer_size"] = strconv.Itoa(VGPU_type.FramebufferSize)

	result["max_heads"] = strconv.Itoa(VGPU_type.MaxHeads)

	result["max_resolution_x"] = strconv.Itoa(VGPU_type.MaxResolutionX)

	result["max_resolution_y"] = strconv.Itoa(VGPU_type.MaxResolutionY)

	result["supported_on_PGPUs"] = VGPU_type.SupportedOnPGPUs

	result["enabled_on_PGPUs"] = VGPU_type.EnabledOnPGPUs

	result["VGPUs"] = VGPU_type.VGPUs

	result["supported_on_GPU_groups"] = VGPU_type.SupportedOnGPUGroups

	result["enabled_on_GPU_groups"] = VGPU_type.EnabledOnGPUGroups

	result["implementation"] = VGPU_type.Implementation.String()

	result["identifier"] = VGPU_type.Identifier

	result["experimental"] = VGPU_type.Experimental

	result["compatible_types_in_vm"] = VGPU_type.CompatibleTypesInVm

	return result
}

func ToVGPUType(obj interface{}) (resultObj *VGPUType) {

	objValue := reflect.ValueOf(obj)
	resultObj = &VGPUType{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.Uuid = v
			}
		case "vendor_name":
			if v, ok := keyValue.(string); ok {
				resultObj.VendorName = v
			}
		case "model_name":
			if v, ok := keyValue.(string); ok {
				resultObj.ModelName = v
			}
		case "framebuffer_size":
			if v, ok := keyValue.(int); ok {
				resultObj.FramebufferSize = v
			}
		case "max_heads":
			if v, ok := keyValue.(int); ok {
				resultObj.MaxHeads = v
			}
		case "max_resolution_x":
			if v, ok := keyValue.(int); ok {
				resultObj.MaxResolutionX = v
			}
		case "max_resolution_y":
			if v, ok := keyValue.(int); ok {
				resultObj.MaxResolutionY = v
			}
		case "supported_on_PGPUs":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.SupportedOnPGPUs = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.SupportedOnPGPUs[i] = v
					}
				}
			}
		case "enabled_on_PGPUs":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.EnabledOnPGPUs = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.EnabledOnPGPUs[i] = v
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
		case "supported_on_GPU_groups":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.SupportedOnGPUGroups = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.SupportedOnGPUGroups[i] = v
					}
				}
			}
		case "enabled_on_GPU_groups":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.EnabledOnGPUGroups = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.EnabledOnGPUGroups[i] = v
					}
				}
			}
		case "implementation":
			if v, ok := keyValue.(VgpuTypeImplementation); ok {
				resultObj.Implementation = v
			}
		case "identifier":
			if v, ok := keyValue.(string); ok {
				resultObj.Identifier = v
			}
		case "experimental":
			if v, ok := keyValue.(bool); ok {
				resultObj.Experimental = v
			}
		case "compatible_types_in_vm":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.CompatibleTypesInVm = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.CompatibleTypesInVm[i] = v
					}
				}
			}
		}
	}

	return resultObj
}

/* GetAllRecords: Return a map of VGPU_type references to VGPU_type records for all VGPU_types known to the system. */
func (client *XenClient) VGPUTypeGetAllRecords() (result map[string]VGPUType, err error) {
	obj, err := client.APICall("VGPU_type.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]VGPUType{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToVGPUType(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the VGPU_types known to the system. */
func (client *XenClient) VGPUTypeGetAll() (result []string, err error) {
	obj, err := client.APICall("VGPU_type.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetCompatibleTypesInVm: Get the compatible_types_in_vm field of the given VGPU_type. */
func (client *XenClient) VGPUTypeGetCompatibleTypesInVm(self string) (result []string, err error) {
	obj, err := client.APICall("VGPU_type.get_compatible_types_in_vm", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetExperimental: Get the experimental field of the given VGPU_type. */
func (client *XenClient) VGPUTypeGetExperimental(self string) (result bool, err error) {
	obj, err := client.APICall("VGPU_type.get_experimental", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetIdentifier: Get the identifier field of the given VGPU_type. */
func (client *XenClient) VGPUTypeGetIdentifier(self string) (result string, err error) {
	obj, err := client.APICall("VGPU_type.get_identifier", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetImplementation: Get the implementation field of the given VGPU_type. */
func (client *XenClient) VGPUTypeGetImplementation(self string) (result VgpuTypeImplementation, err error) {
	obj, err := client.APICall("VGPU_type.get_implementation", self)
	if err != nil {
		return
	}

	result = ToVgpuTypeImplementation(obj.(string))

	return
}

/* GetEnabledOnGPUGroups: Get the enabled_on_GPU_groups field of the given VGPU_type. */
func (client *XenClient) VGPUTypeGetEnabledOnGPUGroups(self string) (result []string, err error) {
	obj, err := client.APICall("VGPU_type.get_enabled_on_GPU_groups", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetSupportedOnGPUGroups: Get the supported_on_GPU_groups field of the given VGPU_type. */
func (client *XenClient) VGPUTypeGetSupportedOnGPUGroups(self string) (result []string, err error) {
	obj, err := client.APICall("VGPU_type.get_supported_on_GPU_groups", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetVGPUs: Get the VGPUs field of the given VGPU_type. */
func (client *XenClient) VGPUTypeGetVGPUs(self string) (result []string, err error) {
	obj, err := client.APICall("VGPU_type.get_VGPUs", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetEnabledOnPGPUs: Get the enabled_on_PGPUs field of the given VGPU_type. */
func (client *XenClient) VGPUTypeGetEnabledOnPGPUs(self string) (result []string, err error) {
	obj, err := client.APICall("VGPU_type.get_enabled_on_PGPUs", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetSupportedOnPGPUs: Get the supported_on_PGPUs field of the given VGPU_type. */
func (client *XenClient) VGPUTypeGetSupportedOnPGPUs(self string) (result []string, err error) {
	obj, err := client.APICall("VGPU_type.get_supported_on_PGPUs", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetMaxResolutionY: Get the max_resolution_y field of the given VGPU_type. */
func (client *XenClient) VGPUTypeGetMaxResolutionY(self string) (result int, err error) {
	obj, err := client.APICall("VGPU_type.get_max_resolution_y", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetMaxResolutionX: Get the max_resolution_x field of the given VGPU_type. */
func (client *XenClient) VGPUTypeGetMaxResolutionX(self string) (result int, err error) {
	obj, err := client.APICall("VGPU_type.get_max_resolution_x", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetMaxHeads: Get the max_heads field of the given VGPU_type. */
func (client *XenClient) VGPUTypeGetMaxHeads(self string) (result int, err error) {
	obj, err := client.APICall("VGPU_type.get_max_heads", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetFramebufferSize: Get the framebuffer_size field of the given VGPU_type. */
func (client *XenClient) VGPUTypeGetFramebufferSize(self string) (result int, err error) {
	obj, err := client.APICall("VGPU_type.get_framebuffer_size", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetModelName: Get the model_name field of the given VGPU_type. */
func (client *XenClient) VGPUTypeGetModelName(self string) (result string, err error) {
	obj, err := client.APICall("VGPU_type.get_model_name", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetVendorName: Get the vendor_name field of the given VGPU_type. */
func (client *XenClient) VGPUTypeGetVendorName(self string) (result string, err error) {
	obj, err := client.APICall("VGPU_type.get_vendor_name", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given VGPU_type. */
func (client *XenClient) VGPUTypeGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("VGPU_type.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the VGPU_type instance with the specified UUID. */
func (client *XenClient) VGPUTypeGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("VGPU_type.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given VGPU_type. */
func (client *XenClient) VGPUTypeGetRecord(self string) (result VGPUType, err error) {
	obj, err := client.APICall("VGPU_type.get_record", self)
	if err != nil {
		return
	}

	result = *ToVGPUType(obj.(interface{}))

	return
}

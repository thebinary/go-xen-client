// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w PCI.gen.go

package go_xen_client

import (
	"reflect"

	"github.com/nilshell/xmlrpc"
)

//PCI: A PCI device
type PCI struct {
	Uuid                string            // Unique identifier/object reference
	ClassName           string            // PCI class name
	VendorName          string            // Vendor name
	DeviceName          string            // Device name
	Host                string            // Physical machine that owns the PCI device
	PciId               string            // PCI ID of the physical device
	Dependencies        []string          // List of dependent PCI devices
	OtherConfig         map[string]string // Additional configuration
	SubsystemVendorName string            // Subsystem vendor name
	SubsystemDeviceName string            // Subsystem device name
	DriverName          string            // Driver name

}

func FromPCIToXml(PCI *PCI) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] =

		PCI.Uuid

	result["class_name"] =

		PCI.ClassName

	result["vendor_name"] =

		PCI.VendorName

	result["device_name"] =

		PCI.DeviceName

	result["host"] =

		PCI.Host

	result["pci_id"] =

		PCI.PciId

	result["dependencies"] =

		PCI.Dependencies

	other_config := make(xmlrpc.Struct)
	for key, value := range PCI.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	result["subsystem_vendor_name"] =

		PCI.SubsystemVendorName

	result["subsystem_device_name"] =

		PCI.SubsystemDeviceName

	result["driver_name"] =

		PCI.DriverName

	return result
}

func ToPCI(obj interface{}) (resultObj *PCI) {

	objValue := reflect.ValueOf(obj)
	resultObj = &PCI{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.Uuid = v
			}
		case "class_name":
			if v, ok := keyValue.(string); ok {
				resultObj.ClassName = v
			}
		case "vendor_name":
			if v, ok := keyValue.(string); ok {
				resultObj.VendorName = v
			}
		case "device_name":
			if v, ok := keyValue.(string); ok {
				resultObj.DeviceName = v
			}
		case "host":
			if v, ok := keyValue.(string); ok {
				resultObj.Host = v
			}
		case "pci_id":
			if v, ok := keyValue.(string); ok {
				resultObj.PciId = v
			}
		case "dependencies":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.Dependencies = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.Dependencies[i] = v
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
		case "subsystem_vendor_name":
			if v, ok := keyValue.(string); ok {
				resultObj.SubsystemVendorName = v
			}
		case "subsystem_device_name":
			if v, ok := keyValue.(string); ok {
				resultObj.SubsystemDeviceName = v
			}
		case "driver_name":
			if v, ok := keyValue.(string); ok {
				resultObj.DriverName = v
			}
		}
	}

	return resultObj
}

/* GetAllRecords: Return a map of PCI references to PCI records for all PCIs known to the system. */
func (client *XenClient) PCIGetAllRecords() (result map[string]PCI, err error) {
	obj, err := client.APICall("PCI.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]PCI{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToPCI(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the PCIs known to the system. */
func (client *XenClient) PCIGetAll() (result []string, err error) {
	obj, err := client.APICall("PCI.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given PCI.  If the key is not in that Map, then do nothing. */
func (client *XenClient) PCIRemoveFromOtherConfig(self string, key string) (err error) {
	_, err = client.APICall("PCI.remove_from_other_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToOtherConfig: Add the given key-value pair to the other_config field of the given PCI. */
func (client *XenClient) PCIAddToOtherConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("PCI.add_to_other_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOtherConfig: Set the other_config field of the given PCI. */
func (client *XenClient) PCISetOtherConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("PCI.set_other_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetDriverName: Get the driver_name field of the given PCI. */
func (client *XenClient) PCIGetDriverName(self string) (result string, err error) {
	obj, err := client.APICall("PCI.get_driver_name", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetSubsystemDeviceName: Get the subsystem_device_name field of the given PCI. */
func (client *XenClient) PCIGetSubsystemDeviceName(self string) (result string, err error) {
	obj, err := client.APICall("PCI.get_subsystem_device_name", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetSubsystemVendorName: Get the subsystem_vendor_name field of the given PCI. */
func (client *XenClient) PCIGetSubsystemVendorName(self string) (result string, err error) {
	obj, err := client.APICall("PCI.get_subsystem_vendor_name", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetOtherConfig: Get the other_config field of the given PCI. */
func (client *XenClient) PCIGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("PCI.get_other_config", self)
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

/* GetDependencies: Get the dependencies field of the given PCI. */
func (client *XenClient) PCIGetDependencies(self string) (result []string, err error) {
	obj, err := client.APICall("PCI.get_dependencies", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetPciId: Get the pci_id field of the given PCI. */
func (client *XenClient) PCIGetPciId(self string) (result string, err error) {
	obj, err := client.APICall("PCI.get_pci_id", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetHost: Get the host field of the given PCI. */
func (client *XenClient) PCIGetHost(self string) (result string, err error) {
	obj, err := client.APICall("PCI.get_host", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetDeviceName: Get the device_name field of the given PCI. */
func (client *XenClient) PCIGetDeviceName(self string) (result string, err error) {
	obj, err := client.APICall("PCI.get_device_name", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetVendorName: Get the vendor_name field of the given PCI. */
func (client *XenClient) PCIGetVendorName(self string) (result string, err error) {
	obj, err := client.APICall("PCI.get_vendor_name", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetClassName: Get the class_name field of the given PCI. */
func (client *XenClient) PCIGetClassName(self string) (result string, err error) {
	obj, err := client.APICall("PCI.get_class_name", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given PCI. */
func (client *XenClient) PCIGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("PCI.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the PCI instance with the specified UUID. */
func (client *XenClient) PCIGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("PCI.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given PCI. */
func (client *XenClient) PCIGetRecord(self string) (result PCI, err error) {
	obj, err := client.APICall("PCI.get_record", self)
	if err != nil {
		return
	}

	result = *ToPCI(obj.(interface{}))

	return
}

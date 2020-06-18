// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w NetworkSriov.gen.go

package go_xen_client

import (
	"reflect"

	"github.com/nilshell/xmlrpc"
)

//NetworkSriov: network-sriov which connects logical pif and physical pif
type NetworkSriov struct {
	Uuid              string                 // Unique identifier/object reference
	PhysicalPIF       string                 // The PIF that has SR-IOV enabled
	LogicalPIF        string                 // The logical PIF to connect to the SR-IOV network after enable SR-IOV on the physical PIF
	RequiresReboot    bool                   // Indicates whether the host need to be rebooted before SR-IOV is enabled on the physical PIF
	ConfigurationMode SriovConfigurationMode // The mode for configure network sriov

}

func FromNetworkSriovToXml(network_sriov *NetworkSriov) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] = network_sriov.Uuid

	result["physical_PIF"] = network_sriov.PhysicalPIF

	result["logical_PIF"] = network_sriov.LogicalPIF

	result["requires_reboot"] = network_sriov.RequiresReboot

	result["configuration_mode"] = network_sriov.ConfigurationMode.String()

	return result
}

func ToNetworkSriov(obj interface{}) (resultObj *NetworkSriov) {

	objValue := reflect.ValueOf(obj)
	resultObj = &NetworkSriov{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.Uuid = v
			}
		case "physical_PIF":
			if v, ok := keyValue.(string); ok {
				resultObj.PhysicalPIF = v
			}
		case "logical_PIF":
			if v, ok := keyValue.(string); ok {
				resultObj.LogicalPIF = v
			}
		case "requires_reboot":
			if v, ok := keyValue.(bool); ok {
				resultObj.RequiresReboot = v
			}
		case "configuration_mode":
			if v, ok := keyValue.(SriovConfigurationMode); ok {
				resultObj.ConfigurationMode = v
			}
		}
	}

	return resultObj
}

/* GetAllRecords: Return a map of network_sriov references to network_sriov records for all network_sriovs known to the system. */
func (client *XenClient) NetworkSriovGetAllRecords() (result map[string]NetworkSriov, err error) {
	obj, err := client.APICall("network_sriov.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]NetworkSriov{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToNetworkSriov(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the network_sriovs known to the system. */
func (client *XenClient) NetworkSriovGetAll() (result []string, err error) {
	obj, err := client.APICall("network_sriov.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetRemainingCapacity: Get the number of free SR-IOV VFs on the associated PIF */
func (client *XenClient) NetworkSriovGetRemainingCapacity(self string) (result int, err error) {
	obj, err := client.APICall("network_sriov.get_remaining_capacity", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* Destroy: Disable SR-IOV on the specific PIF. It will destroy the network-sriov and the logical PIF accordingly. */
func (client *XenClient) NetworkSriovDestroy(self string) (err error) {
	_, err = client.APICall("network_sriov.destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Create: Enable SR-IOV on the specific PIF. It will create a network-sriov based on the specific PIF and automatically create a logical PIF to connect the specific network. */
func (client *XenClient) NetworkSriovCreate(pif string, network string) (result string, err error) {
	obj, err := client.APICall("network_sriov.create", pif, network)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetConfigurationMode: Get the configuration_mode field of the given network_sriov. */
func (client *XenClient) NetworkSriovGetConfigurationMode(self string) (result SriovConfigurationMode, err error) {
	obj, err := client.APICall("network_sriov.get_configuration_mode", self)
	if err != nil {
		return
	}

	result = ToSriovConfigurationMode(obj.(string))

	return
}

/* GetRequiresReboot: Get the requires_reboot field of the given network_sriov. */
func (client *XenClient) NetworkSriovGetRequiresReboot(self string) (result bool, err error) {
	obj, err := client.APICall("network_sriov.get_requires_reboot", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetLogicalPIF: Get the logical_PIF field of the given network_sriov. */
func (client *XenClient) NetworkSriovGetLogicalPIF(self string) (result string, err error) {
	obj, err := client.APICall("network_sriov.get_logical_PIF", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetPhysicalPIF: Get the physical_PIF field of the given network_sriov. */
func (client *XenClient) NetworkSriovGetPhysicalPIF(self string) (result string, err error) {
	obj, err := client.APICall("network_sriov.get_physical_PIF", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given network_sriov. */
func (client *XenClient) NetworkSriovGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("network_sriov.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the network_sriov instance with the specified UUID. */
func (client *XenClient) NetworkSriovGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("network_sriov.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given network_sriov. */
func (client *XenClient) NetworkSriovGetRecord(self string) (result NetworkSriov, err error) {
	obj, err := client.APICall("network_sriov.get_record", self)
	if err != nil {
		return
	}

	result = *ToNetworkSriov(obj.(interface{}))

	return
}

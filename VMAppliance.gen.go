// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w VMAppliance.gen.go

package go_xen_client

import (
	"reflect"

	"github.com/nilshell/xmlrpc"
)

//VMAppliance: VM appliance
type VMAppliance struct {
	Uuid              string                          // Unique identifier/object reference
	NameLabel         string                          // a human-readable name
	NameDescription   string                          // a notes field containing human-readable description
	AllowedOperations []VmApplianceOperation          // list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	CurrentOperations map[string]VmApplianceOperation // links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	VMs               []string                        // all VMs in this appliance

}

func FromVMApplianceToXml(VM_appliance *VMAppliance) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] =

		VM_appliance.Uuid

	result["name_label"] =

		VM_appliance.NameLabel

	result["name_description"] =

		VM_appliance.NameDescription

	result["allowed_operations"] =

		VM_appliance.AllowedOperations

	current_operations := make(xmlrpc.Struct)
	for key, value := range VM_appliance.CurrentOperations {
		current_operations[key] = value
	}
	result["current_operations"] = current_operations

	result["VMs"] =

		VM_appliance.VMs

	return result
}

func ToVMAppliance(obj interface{}) (resultObj *VMAppliance) {

	objValue := reflect.ValueOf(obj)
	resultObj = &VMAppliance{}

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
		case "allowed_operations":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.AllowedOperations = make([]VmApplianceOperation, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(VmApplianceOperation); ok {
						resultObj.AllowedOperations[i] = v
					}
				}
			}
		case "current_operations":

			resultObj.CurrentOperations = map[string]VmApplianceOperation{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.CurrentOperations[mapKeyName] = ToVmApplianceOperation(v)
				} else {
					resultObj.CurrentOperations[mapKeyName] = 0
				}
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

/* GetAllRecords: Return a map of VM_appliance references to VM_appliance records for all VM_appliances known to the system. */
func (client *XenClient) VMApplianceGetAllRecords() (result map[string]VMAppliance, err error) {
	obj, err := client.APICall("VM_appliance.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]VMAppliance{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToVMAppliance(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the VM_appliances known to the system. */
func (client *XenClient) VMApplianceGetAll() (result []string, err error) {
	obj, err := client.APICall("VM_appliance.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* Recover: Recover the VM appliance */
func (client *XenClient) VMApplianceRecover(self string, session_to string, force bool) (err error) {
	_, err = client.APICall("VM_appliance.recover", self, session_to, force)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetSRsRequiredForRecovery: Get the list of SRs required by the VM appliance to recover. */
func (client *XenClient) VMApplianceGetSRsRequiredForRecovery(self string, session_to string) (result []string, err error) {
	obj, err := client.APICall("VM_appliance.get_SRs_required_for_recovery", self, session_to)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* AssertCanBeRecovered: Assert whether all SRs required to recover this VM appliance are available. */
func (client *XenClient) VMApplianceAssertCanBeRecovered(self string, session_to string) (err error) {
	_, err = client.APICall("VM_appliance.assert_can_be_recovered", self, session_to)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Shutdown: For each VM in the appliance, try to shut it down cleanly. If this fails, perform a hard shutdown of the VM. */
func (client *XenClient) VMApplianceShutdown(self string) (err error) {
	_, err = client.APICall("VM_appliance.shutdown", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* HardShutdown: Perform a hard shutdown of all the VMs in the appliance */
func (client *XenClient) VMApplianceHardShutdown(self string) (err error) {
	_, err = client.APICall("VM_appliance.hard_shutdown", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* CleanShutdown: Perform a clean shutdown of all the VMs in the appliance */
func (client *XenClient) VMApplianceCleanShutdown(self string) (err error) {
	_, err = client.APICall("VM_appliance.clean_shutdown", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Start: Start all VMs in the appliance */
func (client *XenClient) VMApplianceStart(self string, paused bool) (err error) {
	_, err = client.APICall("VM_appliance.start", self, paused)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetNameDescription: Set the name/description field of the given VM_appliance. */
func (client *XenClient) VMApplianceSetNameDescription(self string, value string) (err error) {
	_, err = client.APICall("VM_appliance.set_name_description", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetNameLabel: Set the name/label field of the given VM_appliance. */
func (client *XenClient) VMApplianceSetNameLabel(self string, value string) (err error) {
	_, err = client.APICall("VM_appliance.set_name_label", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetVMs: Get the VMs field of the given VM_appliance. */
func (client *XenClient) VMApplianceGetVMs(self string) (result []string, err error) {
	obj, err := client.APICall("VM_appliance.get_VMs", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetCurrentOperations: Get the current_operations field of the given VM_appliance. */
func (client *XenClient) VMApplianceGetCurrentOperations(self string) (result map[string]VmApplianceOperation, err error) {
	obj, err := client.APICall("VM_appliance.get_current_operations", self)
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]VmApplianceOperation{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToVmApplianceOperation(obj.String())
		result[key.String()] = mapObj
	}

	return
}

/* GetAllowedOperations: Get the allowed_operations field of the given VM_appliance. */
func (client *XenClient) VMApplianceGetAllowedOperations(self string) (result []VmApplianceOperation, err error) {
	obj, err := client.APICall("VM_appliance.get_allowed_operations", self)
	if err != nil {
		return
	}

	result = make([]VmApplianceOperation, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = ToVmApplianceOperation(value.(string))
	}

	return
}

/* GetNameDescription: Get the name/description field of the given VM_appliance. */
func (client *XenClient) VMApplianceGetNameDescription(self string) (result string, err error) {
	obj, err := client.APICall("VM_appliance.get_name_description", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetNameLabel: Get the name/label field of the given VM_appliance. */
func (client *XenClient) VMApplianceGetNameLabel(self string) (result string, err error) {
	obj, err := client.APICall("VM_appliance.get_name_label", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given VM_appliance. */
func (client *XenClient) VMApplianceGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("VM_appliance.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByNameLabel: Get all the VM_appliance instances with the given label. */
func (client *XenClient) VMApplianceGetByNameLabel(label string) (result []string, err error) {
	obj, err := client.APICall("VM_appliance.get_by_name_label", label)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* Destroy: Destroy the specified VM_appliance instance. */
func (client *XenClient) VMApplianceDestroy(self string) (err error) {
	_, err = client.APICall("VM_appliance.destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Create: Create a new VM_appliance instance, and return its handle.
The constructor args are: name_label, name_description (* = non-optional). */
func (client *XenClient) VMApplianceCreate(args VMAppliance) (result string, err error) {
	obj, err := client.APICall("VM_appliance.create", FromVMApplianceToXml(&args))
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the VM_appliance instance with the specified UUID. */
func (client *XenClient) VMApplianceGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("VM_appliance.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given VM_appliance. */
func (client *XenClient) VMApplianceGetRecord(self string) (result VMAppliance, err error) {
	obj, err := client.APICall("VM_appliance.get_record", self)
	if err != nil {
		return
	}

	result = *ToVMAppliance(obj.(interface{}))

	return
}

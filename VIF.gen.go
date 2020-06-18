// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w VIF.gen.go

package go_xen_client

import (
	"reflect"
	"strconv"

	"github.com/nilshell/xmlrpc"
)

//VIF: A virtual network interface
type VIF struct {
	Uuid                   string                   // Unique identifier/object reference
	AllowedOperations      []VifOperations          // list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	CurrentOperations      map[string]VifOperations // links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	Device                 string                   // order in which VIF backends are created by xapi
	Network                string                   // virtual network to which this vif is connected
	VM                     string                   // virtual machine to which this vif is connected
	MAC                    string                   // ethernet MAC address of virtual interface, as exposed to guest
	MTU                    int                      // MTU in octets
	OtherConfig            map[string]string        // additional configuration
	CurrentlyAttached      bool                     // is the device currently attached (erased on reboot)
	StatusCode             int                      // error/success code associated with last attach-operation (erased on reboot)
	StatusDetail           string                   // error/success information associated with last attach-operation status (erased on reboot)
	RuntimeProperties      map[string]string        // Device runtime properties
	QosAlgorithmType       string                   // QoS algorithm to use
	QosAlgorithmParams     map[string]string        // parameters for chosen QoS algorithm
	QosSupportedAlgorithms []string                 // supported QoS algorithms for this VIF
	Metrics                string                   // metrics associated with this VIF
	MACAutogenerated       bool                     // true if the MAC was autogenerated; false indicates it was set manually
	LockingMode            VifLockingMode           // current locking mode of the VIF
	Ipv4Allowed            []string                 // A list of IPv4 addresses which can be used to filter traffic passing through this VIF
	Ipv6Allowed            []string                 // A list of IPv6 addresses which can be used to filter traffic passing through this VIF
	Ipv4ConfigurationMode  VifIpv4ConfigurationMode // Determines whether IPv4 addresses are configured on the VIF
	Ipv4Addresses          []string                 // IPv4 addresses in CIDR format
	Ipv4Gateway            string                   // IPv4 gateway (the empty string means that no gateway is set)
	Ipv6ConfigurationMode  VifIpv6ConfigurationMode // Determines whether IPv6 addresses are configured on the VIF
	Ipv6Addresses          []string                 // IPv6 addresses in CIDR format
	Ipv6Gateway            string                   // IPv6 gateway (the empty string means that no gateway is set)

}

func FromVIFToXml(VIF *VIF) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] =

		VIF.Uuid

	result["allowed_operations"] =

		VIF.AllowedOperations

	current_operations := make(xmlrpc.Struct)
	for key, value := range VIF.CurrentOperations {
		current_operations[key] = value
	}
	result["current_operations"] = current_operations

	result["device"] =

		VIF.Device

	result["network"] =

		VIF.Network

	result["VM"] =

		VIF.VM

	result["MAC"] =

		VIF.MAC

	result["MTU"] =

		strconv.Itoa(VIF.MTU)

	other_config := make(xmlrpc.Struct)
	for key, value := range VIF.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	result["currently_attached"] =

		VIF.CurrentlyAttached

	result["status_code"] =

		strconv.Itoa(VIF.StatusCode)

	result["status_detail"] =

		VIF.StatusDetail

	runtime_properties := make(xmlrpc.Struct)
	for key, value := range VIF.RuntimeProperties {
		runtime_properties[key] = value
	}
	result["runtime_properties"] = runtime_properties

	result["qos_algorithm_type"] =

		VIF.QosAlgorithmType

	qos_algorithm_params := make(xmlrpc.Struct)
	for key, value := range VIF.QosAlgorithmParams {
		qos_algorithm_params[key] = value
	}
	result["qos_algorithm_params"] = qos_algorithm_params

	result["qos_supported_algorithms"] =

		VIF.QosSupportedAlgorithms

	result["metrics"] =

		VIF.Metrics

	result["MAC_autogenerated"] =

		VIF.MACAutogenerated

	result["locking_mode"] =

		VIF.LockingMode.String()

	result["ipv4_allowed"] =

		VIF.Ipv4Allowed

	result["ipv6_allowed"] =

		VIF.Ipv6Allowed

	result["ipv4_configuration_mode"] =

		VIF.Ipv4ConfigurationMode.String()

	result["ipv4_addresses"] =

		VIF.Ipv4Addresses

	result["ipv4_gateway"] =

		VIF.Ipv4Gateway

	result["ipv6_configuration_mode"] =

		VIF.Ipv6ConfigurationMode.String()

	result["ipv6_addresses"] =

		VIF.Ipv6Addresses

	result["ipv6_gateway"] =

		VIF.Ipv6Gateway

	return result
}

func ToVIF(obj interface{}) (resultObj *VIF) {

	objValue := reflect.ValueOf(obj)
	resultObj = &VIF{}

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
				resultObj.AllowedOperations = make([]VifOperations, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(VifOperations); ok {
						resultObj.AllowedOperations[i] = v
					}
				}
			}
		case "current_operations":

			resultObj.CurrentOperations = map[string]VifOperations{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.CurrentOperations[mapKeyName] = ToVifOperations(v)
				} else {
					resultObj.CurrentOperations[mapKeyName] = 0
				}
			}
		case "device":
			if v, ok := keyValue.(string); ok {
				resultObj.Device = v
			}
		case "network":
			if v, ok := keyValue.(string); ok {
				resultObj.Network = v
			}
		case "VM":
			if v, ok := keyValue.(string); ok {
				resultObj.VM = v
			}
		case "MAC":
			if v, ok := keyValue.(string); ok {
				resultObj.MAC = v
			}
		case "MTU":
			if v, ok := keyValue.(int); ok {
				resultObj.MTU = v
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
		case "MAC_autogenerated":
			if v, ok := keyValue.(bool); ok {
				resultObj.MACAutogenerated = v
			}
		case "locking_mode":
			if v, ok := keyValue.(VifLockingMode); ok {
				resultObj.LockingMode = v
			}
		case "ipv4_allowed":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.Ipv4Allowed = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.Ipv4Allowed[i] = v
					}
				}
			}
		case "ipv6_allowed":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.Ipv6Allowed = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.Ipv6Allowed[i] = v
					}
				}
			}
		case "ipv4_configuration_mode":
			if v, ok := keyValue.(VifIpv4ConfigurationMode); ok {
				resultObj.Ipv4ConfigurationMode = v
			}
		case "ipv4_addresses":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.Ipv4Addresses = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.Ipv4Addresses[i] = v
					}
				}
			}
		case "ipv4_gateway":
			if v, ok := keyValue.(string); ok {
				resultObj.Ipv4Gateway = v
			}
		case "ipv6_configuration_mode":
			if v, ok := keyValue.(VifIpv6ConfigurationMode); ok {
				resultObj.Ipv6ConfigurationMode = v
			}
		case "ipv6_addresses":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.Ipv6Addresses = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.Ipv6Addresses[i] = v
					}
				}
			}
		case "ipv6_gateway":
			if v, ok := keyValue.(string); ok {
				resultObj.Ipv6Gateway = v
			}
		}
	}

	return resultObj
}

/* GetAllRecords: Return a map of VIF references to VIF records for all VIFs known to the system. */
func (client *XenClient) VIFGetAllRecords() (result map[string]VIF, err error) {
	obj, err := client.APICall("VIF.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]VIF{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToVIF(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the VIFs known to the system. */
func (client *XenClient) VIFGetAll() (result []string, err error) {
	obj, err := client.APICall("VIF.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* ConfigureIpv6: Configure IPv6 settings for this virtual interface */
func (client *XenClient) VIFConfigureIpv6(self string, mode VifIpv6ConfigurationMode, address string, gateway string) (err error) {
	_, err = client.APICall("VIF.configure_ipv6", self, mode.String(), address, gateway)
	if err != nil {
		return
	}
	// no return result
	return
}

/* ConfigureIpv4: Configure IPv4 settings for this virtual interface */
func (client *XenClient) VIFConfigureIpv4(self string, mode VifIpv4ConfigurationMode, address string, gateway string) (err error) {
	_, err = client.APICall("VIF.configure_ipv4", self, mode.String(), address, gateway)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveIpv6Allowed: Removes an IPv6 address from this VIF */
func (client *XenClient) VIFRemoveIpv6Allowed(self string, value string) (err error) {
	_, err = client.APICall("VIF.remove_ipv6_allowed", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddIpv6Allowed: Associates an IPv6 address with this VIF */
func (client *XenClient) VIFAddIpv6Allowed(self string, value string) (err error) {
	_, err = client.APICall("VIF.add_ipv6_allowed", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetIpv6Allowed: Set the IPv6 addresses to which traffic on this VIF can be restricted */
func (client *XenClient) VIFSetIpv6Allowed(self string, value []string) (err error) {
	_, err = client.APICall("VIF.set_ipv6_allowed", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveIpv4Allowed: Removes an IPv4 address from this VIF */
func (client *XenClient) VIFRemoveIpv4Allowed(self string, value string) (err error) {
	_, err = client.APICall("VIF.remove_ipv4_allowed", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddIpv4Allowed: Associates an IPv4 address with this VIF */
func (client *XenClient) VIFAddIpv4Allowed(self string, value string) (err error) {
	_, err = client.APICall("VIF.add_ipv4_allowed", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetIpv4Allowed: Set the IPv4 addresses to which traffic on this VIF can be restricted */
func (client *XenClient) VIFSetIpv4Allowed(self string, value []string) (err error) {
	_, err = client.APICall("VIF.set_ipv4_allowed", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetLockingMode: Set the locking mode for this VIF */
func (client *XenClient) VIFSetLockingMode(self string, value VifLockingMode) (err error) {
	_, err = client.APICall("VIF.set_locking_mode", self, value.String())
	if err != nil {
		return
	}
	// no return result
	return
}

/* Move: Move the specified VIF to the specified network, even while the VM is running */
func (client *XenClient) VIFMove(self string, network string) (err error) {
	_, err = client.APICall("VIF.move", self, network)
	if err != nil {
		return
	}
	// no return result
	return
}

/* UnplugForce: Forcibly unplug the specified VIF */
func (client *XenClient) VIFUnplugForce(self string) (err error) {
	_, err = client.APICall("VIF.unplug_force", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Unplug: Hot-unplug the specified VIF, dynamically unattaching it from the running VM */
func (client *XenClient) VIFUnplug(self string) (err error) {
	_, err = client.APICall("VIF.unplug", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Plug: Hotplug the specified VIF, dynamically attaching it to the running VM */
func (client *XenClient) VIFPlug(self string) (err error) {
	_, err = client.APICall("VIF.plug", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveFromQosAlgorithmParams: Remove the given key and its corresponding value from the qos/algorithm_params field of the given VIF.  If the key is not in that Map, then do nothing. */
func (client *XenClient) VIFRemoveFromQosAlgorithmParams(self string, key string) (err error) {
	_, err = client.APICall("VIF.remove_from_qos_algorithm_params", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToQosAlgorithmParams: Add the given key-value pair to the qos/algorithm_params field of the given VIF. */
func (client *XenClient) VIFAddToQosAlgorithmParams(self string, key string, value string) (err error) {
	_, err = client.APICall("VIF.add_to_qos_algorithm_params", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetQosAlgorithmParams: Set the qos/algorithm_params field of the given VIF. */
func (client *XenClient) VIFSetQosAlgorithmParams(self string, value map[string]string) (err error) {
	_, err = client.APICall("VIF.set_qos_algorithm_params", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetQosAlgorithmType: Set the qos/algorithm_type field of the given VIF. */
func (client *XenClient) VIFSetQosAlgorithmType(self string, value string) (err error) {
	_, err = client.APICall("VIF.set_qos_algorithm_type", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given VIF.  If the key is not in that Map, then do nothing. */
func (client *XenClient) VIFRemoveFromOtherConfig(self string, key string) (err error) {
	_, err = client.APICall("VIF.remove_from_other_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToOtherConfig: Add the given key-value pair to the other_config field of the given VIF. */
func (client *XenClient) VIFAddToOtherConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("VIF.add_to_other_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOtherConfig: Set the other_config field of the given VIF. */
func (client *XenClient) VIFSetOtherConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("VIF.set_other_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetIpv6Gateway: Get the ipv6_gateway field of the given VIF. */
func (client *XenClient) VIFGetIpv6Gateway(self string) (result string, err error) {
	obj, err := client.APICall("VIF.get_ipv6_gateway", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetIpv6Addresses: Get the ipv6_addresses field of the given VIF. */
func (client *XenClient) VIFGetIpv6Addresses(self string) (result []string, err error) {
	obj, err := client.APICall("VIF.get_ipv6_addresses", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetIpv6ConfigurationMode: Get the ipv6_configuration_mode field of the given VIF. */
func (client *XenClient) VIFGetIpv6ConfigurationMode(self string) (result VifIpv6ConfigurationMode, err error) {
	obj, err := client.APICall("VIF.get_ipv6_configuration_mode", self)
	if err != nil {
		return
	}

	result = ToVifIpv6ConfigurationMode(obj.(string))

	return
}

/* GetIpv4Gateway: Get the ipv4_gateway field of the given VIF. */
func (client *XenClient) VIFGetIpv4Gateway(self string) (result string, err error) {
	obj, err := client.APICall("VIF.get_ipv4_gateway", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetIpv4Addresses: Get the ipv4_addresses field of the given VIF. */
func (client *XenClient) VIFGetIpv4Addresses(self string) (result []string, err error) {
	obj, err := client.APICall("VIF.get_ipv4_addresses", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetIpv4ConfigurationMode: Get the ipv4_configuration_mode field of the given VIF. */
func (client *XenClient) VIFGetIpv4ConfigurationMode(self string) (result VifIpv4ConfigurationMode, err error) {
	obj, err := client.APICall("VIF.get_ipv4_configuration_mode", self)
	if err != nil {
		return
	}

	result = ToVifIpv4ConfigurationMode(obj.(string))

	return
}

/* GetIpv6Allowed: Get the ipv6_allowed field of the given VIF. */
func (client *XenClient) VIFGetIpv6Allowed(self string) (result []string, err error) {
	obj, err := client.APICall("VIF.get_ipv6_allowed", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetIpv4Allowed: Get the ipv4_allowed field of the given VIF. */
func (client *XenClient) VIFGetIpv4Allowed(self string) (result []string, err error) {
	obj, err := client.APICall("VIF.get_ipv4_allowed", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetLockingMode: Get the locking_mode field of the given VIF. */
func (client *XenClient) VIFGetLockingMode(self string) (result VifLockingMode, err error) {
	obj, err := client.APICall("VIF.get_locking_mode", self)
	if err != nil {
		return
	}

	result = ToVifLockingMode(obj.(string))

	return
}

/* GetMACAutogenerated: Get the MAC_autogenerated field of the given VIF. */
func (client *XenClient) VIFGetMACAutogenerated(self string) (result bool, err error) {
	obj, err := client.APICall("VIF.get_MAC_autogenerated", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetMetrics: Get the metrics field of the given VIF. */
func (client *XenClient) VIFGetMetrics(self string) (result string, err error) {
	obj, err := client.APICall("VIF.get_metrics", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetQosSupportedAlgorithms: Get the qos/supported_algorithms field of the given VIF. */
func (client *XenClient) VIFGetQosSupportedAlgorithms(self string) (result []string, err error) {
	obj, err := client.APICall("VIF.get_qos_supported_algorithms", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetQosAlgorithmParams: Get the qos/algorithm_params field of the given VIF. */
func (client *XenClient) VIFGetQosAlgorithmParams(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VIF.get_qos_algorithm_params", self)
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

/* GetQosAlgorithmType: Get the qos/algorithm_type field of the given VIF. */
func (client *XenClient) VIFGetQosAlgorithmType(self string) (result string, err error) {
	obj, err := client.APICall("VIF.get_qos_algorithm_type", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRuntimeProperties: Get the runtime_properties field of the given VIF. */
func (client *XenClient) VIFGetRuntimeProperties(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VIF.get_runtime_properties", self)
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

/* GetStatusDetail: Get the status_detail field of the given VIF. */
func (client *XenClient) VIFGetStatusDetail(self string) (result string, err error) {
	obj, err := client.APICall("VIF.get_status_detail", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetStatusCode: Get the status_code field of the given VIF. */
func (client *XenClient) VIFGetStatusCode(self string) (result int, err error) {
	obj, err := client.APICall("VIF.get_status_code", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetCurrentlyAttached: Get the currently_attached field of the given VIF. */
func (client *XenClient) VIFGetCurrentlyAttached(self string) (result bool, err error) {
	obj, err := client.APICall("VIF.get_currently_attached", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetOtherConfig: Get the other_config field of the given VIF. */
func (client *XenClient) VIFGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VIF.get_other_config", self)
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

/* GetMTU: Get the MTU field of the given VIF. */
func (client *XenClient) VIFGetMTU(self string) (result int, err error) {
	obj, err := client.APICall("VIF.get_MTU", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetMAC: Get the MAC field of the given VIF. */
func (client *XenClient) VIFGetMAC(self string) (result string, err error) {
	obj, err := client.APICall("VIF.get_MAC", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetVM: Get the VM field of the given VIF. */
func (client *XenClient) VIFGetVM(self string) (result string, err error) {
	obj, err := client.APICall("VIF.get_VM", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetNetwork: Get the network field of the given VIF. */
func (client *XenClient) VIFGetNetwork(self string) (result string, err error) {
	obj, err := client.APICall("VIF.get_network", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetDevice: Get the device field of the given VIF. */
func (client *XenClient) VIFGetDevice(self string) (result string, err error) {
	obj, err := client.APICall("VIF.get_device", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetCurrentOperations: Get the current_operations field of the given VIF. */
func (client *XenClient) VIFGetCurrentOperations(self string) (result map[string]VifOperations, err error) {
	obj, err := client.APICall("VIF.get_current_operations", self)
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]VifOperations{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToVifOperations(obj.String())
		result[key.String()] = mapObj
	}

	return
}

/* GetAllowedOperations: Get the allowed_operations field of the given VIF. */
func (client *XenClient) VIFGetAllowedOperations(self string) (result []VifOperations, err error) {
	obj, err := client.APICall("VIF.get_allowed_operations", self)
	if err != nil {
		return
	}

	result = make([]VifOperations, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = ToVifOperations(value.(string))
	}

	return
}

/* GetUuid: Get the uuid field of the given VIF. */
func (client *XenClient) VIFGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("VIF.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* Destroy: Destroy the specified VIF instance. */
func (client *XenClient) VIFDestroy(self string) (err error) {
	_, err = client.APICall("VIF.destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Create: Create a new VIF instance, and return its handle.
The constructor args are: device*, network*, VM*, MAC*, MTU*, other_config*, qos_algorithm_type*, qos_algorithm_params*, locking_mode, ipv4_allowed, ipv6_allowed (* = non-optional). */
func (client *XenClient) VIFCreate(args VIF) (result string, err error) {
	obj, err := client.APICall("VIF.create", FromVIFToXml(&args))
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the VIF instance with the specified UUID. */
func (client *XenClient) VIFGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("VIF.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given VIF. */
func (client *XenClient) VIFGetRecord(self string) (result VIF, err error) {
	obj, err := client.APICall("VIF.get_record", self)
	if err != nil {
		return
	}

	result = *ToVIF(obj.(interface{}))

	return
}

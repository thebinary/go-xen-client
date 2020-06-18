// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w Network.gen.go

package go_xen_client

import (
	"reflect"
	"strconv"

	"github.com/nilshell/xmlrpc"
)

//Network: A virtual network
type Network struct {
	Uuid               string                       // Unique identifier/object reference
	NameLabel          string                       // a human-readable name
	NameDescription    string                       // a notes field containing human-readable description
	AllowedOperations  []NetworkOperations          // list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	CurrentOperations  map[string]NetworkOperations // links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	VIFs               []string                     // list of connected vifs
	PIFs               []string                     // list of connected pifs
	MTU                int                          // MTU in octets
	OtherConfig        map[string]string            // additional configuration
	Bridge             string                       // name of the bridge corresponding to this network on the local host
	Managed            bool                         // true if the bridge is managed by xapi
	Blobs              map[string]string            // Binary blobs associated with this network
	Tags               []string                     // user-specified tags for categorization purposes
	DefaultLockingMode NetworkDefaultLockingMode    // The network will use this value to determine the behaviour of all VIFs where locking_mode = default
	AssignedIps        map[string]string            // The IP addresses assigned to VIFs on networks that have active xapi-managed DHCP
	Purpose            []NetworkPurpose             // Set of purposes for which the server will use this network

}

func FromNetworkToXml(network *Network) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] =

		network.Uuid

	result["name_label"] =

		network.NameLabel

	result["name_description"] =

		network.NameDescription

	result["allowed_operations"] =

		network.AllowedOperations

	current_operations := make(xmlrpc.Struct)
	for key, value := range network.CurrentOperations {
		current_operations[key] = value
	}
	result["current_operations"] = current_operations

	result["VIFs"] =

		network.VIFs

	result["PIFs"] =

		network.PIFs

	result["MTU"] =

		strconv.Itoa(network.MTU)

	other_config := make(xmlrpc.Struct)
	for key, value := range network.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	result["bridge"] =

		network.Bridge

	result["managed"] =

		network.Managed

	blobs := make(xmlrpc.Struct)
	for key, value := range network.Blobs {
		blobs[key] = value
	}
	result["blobs"] = blobs

	result["tags"] =

		network.Tags

	result["default_locking_mode"] =

		network.DefaultLockingMode.String()

	assigned_ips := make(xmlrpc.Struct)
	for key, value := range network.AssignedIps {
		assigned_ips[key] = value
	}
	result["assigned_ips"] = assigned_ips

	result["purpose"] =

		network.Purpose

	return result
}

func ToNetwork(obj interface{}) (resultObj *Network) {

	objValue := reflect.ValueOf(obj)
	resultObj = &Network{}

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
				resultObj.AllowedOperations = make([]NetworkOperations, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(NetworkOperations); ok {
						resultObj.AllowedOperations[i] = v
					}
				}
			}
		case "current_operations":

			resultObj.CurrentOperations = map[string]NetworkOperations{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.CurrentOperations[mapKeyName] = ToNetworkOperations(v)
				} else {
					resultObj.CurrentOperations[mapKeyName] = 0
				}
			}
		case "VIFs":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.VIFs = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.VIFs[i] = v
					}
				}
			}
		case "PIFs":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.PIFs = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.PIFs[i] = v
					}
				}
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
		case "bridge":
			if v, ok := keyValue.(string); ok {
				resultObj.Bridge = v
			}
		case "managed":
			if v, ok := keyValue.(bool); ok {
				resultObj.Managed = v
			}
		case "blobs":

			resultObj.Blobs = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.Blobs[mapKeyName] = v
				} else {
					resultObj.Blobs[mapKeyName] = ""
				}
			}
		case "tags":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.Tags = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.Tags[i] = v
					}
				}
			}
		case "default_locking_mode":
			if v, ok := keyValue.(NetworkDefaultLockingMode); ok {
				resultObj.DefaultLockingMode = v
			}
		case "assigned_ips":

			resultObj.AssignedIps = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.AssignedIps[mapKeyName] = v
				} else {
					resultObj.AssignedIps[mapKeyName] = ""
				}
			}
		case "purpose":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.Purpose = make([]NetworkPurpose, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(NetworkPurpose); ok {
						resultObj.Purpose[i] = v
					}
				}
			}
		}
	}

	return resultObj
}

/* GetAllRecords: Return a map of network references to network records for all networks known to the system. */
func (client *XenClient) NetworkGetAllRecords() (result map[string]Network, err error) {
	obj, err := client.APICall("network.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]Network{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToNetwork(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the networks known to the system. */
func (client *XenClient) NetworkGetAll() (result []string, err error) {
	obj, err := client.APICall("network.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* RemovePurpose: Remove a purpose from a network (if present) */
func (client *XenClient) NetworkRemovePurpose(self string, value NetworkPurpose) (err error) {
	_, err = client.APICall("network.remove_purpose", self, value.String())
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddPurpose: Give a network a new purpose (if not present already) */
func (client *XenClient) NetworkAddPurpose(self string, value NetworkPurpose) (err error) {
	_, err = client.APICall("network.add_purpose", self, value.String())
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetDefaultLockingMode: Set the default locking mode for VIFs attached to this network */
func (client *XenClient) NetworkSetDefaultLockingMode(network string, value NetworkDefaultLockingMode) (err error) {
	_, err = client.APICall("network.set_default_locking_mode", network, value.String())
	if err != nil {
		return
	}
	// no return result
	return
}

/* CreateNewBlob: Create a placeholder for a named binary blob of data that is associated with this pool */
func (client *XenClient) NetworkCreateNewBlob(network string, name string, mime_type string, public bool) (result string, err error) {
	obj, err := client.APICall("network.create_new_blob", network, name, mime_type, public)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* RemoveTags: Remove the given value from the tags field of the given network.  If the value is not in that Set, then do nothing. */
func (client *XenClient) NetworkRemoveTags(self string, value string) (err error) {
	_, err = client.APICall("network.remove_tags", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddTags: Add the given value to the tags field of the given network.  If the value is already in that Set, then do nothing. */
func (client *XenClient) NetworkAddTags(self string, value string) (err error) {
	_, err = client.APICall("network.add_tags", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetTags: Set the tags field of the given network. */
func (client *XenClient) NetworkSetTags(self string, value []string) (err error) {
	_, err = client.APICall("network.set_tags", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given network.  If the key is not in that Map, then do nothing. */
func (client *XenClient) NetworkRemoveFromOtherConfig(self string, key string) (err error) {
	_, err = client.APICall("network.remove_from_other_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToOtherConfig: Add the given key-value pair to the other_config field of the given network. */
func (client *XenClient) NetworkAddToOtherConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("network.add_to_other_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOtherConfig: Set the other_config field of the given network. */
func (client *XenClient) NetworkSetOtherConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("network.set_other_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetMTU: Set the MTU field of the given network. */
func (client *XenClient) NetworkSetMTU(self string, value int) (err error) {
	_, err = client.APICall("network.set_MTU", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetNameDescription: Set the name/description field of the given network. */
func (client *XenClient) NetworkSetNameDescription(self string, value string) (err error) {
	_, err = client.APICall("network.set_name_description", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetNameLabel: Set the name/label field of the given network. */
func (client *XenClient) NetworkSetNameLabel(self string, value string) (err error) {
	_, err = client.APICall("network.set_name_label", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetPurpose: Get the purpose field of the given network. */
func (client *XenClient) NetworkGetPurpose(self string) (result []NetworkPurpose, err error) {
	obj, err := client.APICall("network.get_purpose", self)
	if err != nil {
		return
	}

	result = make([]NetworkPurpose, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = ToNetworkPurpose(value.(string))
	}

	return
}

/* GetAssignedIps: Get the assigned_ips field of the given network. */
func (client *XenClient) NetworkGetAssignedIps(self string) (result map[string]string, err error) {
	obj, err := client.APICall("network.get_assigned_ips", self)
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

/* GetDefaultLockingMode: Get the default_locking_mode field of the given network. */
func (client *XenClient) NetworkGetDefaultLockingMode(self string) (result NetworkDefaultLockingMode, err error) {
	obj, err := client.APICall("network.get_default_locking_mode", self)
	if err != nil {
		return
	}

	result = ToNetworkDefaultLockingMode(obj.(string))

	return
}

/* GetTags: Get the tags field of the given network. */
func (client *XenClient) NetworkGetTags(self string) (result []string, err error) {
	obj, err := client.APICall("network.get_tags", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetBlobs: Get the blobs field of the given network. */
func (client *XenClient) NetworkGetBlobs(self string) (result map[string]string, err error) {
	obj, err := client.APICall("network.get_blobs", self)
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

/* GetManaged: Get the managed field of the given network. */
func (client *XenClient) NetworkGetManaged(self string) (result bool, err error) {
	obj, err := client.APICall("network.get_managed", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetBridge: Get the bridge field of the given network. */
func (client *XenClient) NetworkGetBridge(self string) (result string, err error) {
	obj, err := client.APICall("network.get_bridge", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetOtherConfig: Get the other_config field of the given network. */
func (client *XenClient) NetworkGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("network.get_other_config", self)
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

/* GetMTU: Get the MTU field of the given network. */
func (client *XenClient) NetworkGetMTU(self string) (result int, err error) {
	obj, err := client.APICall("network.get_MTU", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetPIFs: Get the PIFs field of the given network. */
func (client *XenClient) NetworkGetPIFs(self string) (result []string, err error) {
	obj, err := client.APICall("network.get_PIFs", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetVIFs: Get the VIFs field of the given network. */
func (client *XenClient) NetworkGetVIFs(self string) (result []string, err error) {
	obj, err := client.APICall("network.get_VIFs", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetCurrentOperations: Get the current_operations field of the given network. */
func (client *XenClient) NetworkGetCurrentOperations(self string) (result map[string]NetworkOperations, err error) {
	obj, err := client.APICall("network.get_current_operations", self)
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]NetworkOperations{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToNetworkOperations(obj.String())
		result[key.String()] = mapObj
	}

	return
}

/* GetAllowedOperations: Get the allowed_operations field of the given network. */
func (client *XenClient) NetworkGetAllowedOperations(self string) (result []NetworkOperations, err error) {
	obj, err := client.APICall("network.get_allowed_operations", self)
	if err != nil {
		return
	}

	result = make([]NetworkOperations, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = ToNetworkOperations(value.(string))
	}

	return
}

/* GetNameDescription: Get the name/description field of the given network. */
func (client *XenClient) NetworkGetNameDescription(self string) (result string, err error) {
	obj, err := client.APICall("network.get_name_description", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetNameLabel: Get the name/label field of the given network. */
func (client *XenClient) NetworkGetNameLabel(self string) (result string, err error) {
	obj, err := client.APICall("network.get_name_label", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given network. */
func (client *XenClient) NetworkGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("network.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByNameLabel: Get all the network instances with the given label. */
func (client *XenClient) NetworkGetByNameLabel(label string) (result []string, err error) {
	obj, err := client.APICall("network.get_by_name_label", label)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* Destroy: Destroy the specified network instance. */
func (client *XenClient) NetworkDestroy(self string) (err error) {
	_, err = client.APICall("network.destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Create: Create a new network instance, and return its handle.
The constructor args are: name_label, name_description, MTU, other_config*, bridge, managed, tags (* = non-optional). */
func (client *XenClient) NetworkCreate(args Network) (result string, err error) {
	obj, err := client.APICall("network.create", FromNetworkToXml(&args))
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the network instance with the specified UUID. */
func (client *XenClient) NetworkGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("network.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given network. */
func (client *XenClient) NetworkGetRecord(self string) (result Network, err error) {
	obj, err := client.APICall("network.get_record", self)
	if err != nil {
		return
	}

	result = *ToNetwork(obj.(interface{}))

	return
}

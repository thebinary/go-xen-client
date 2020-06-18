// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w ClusterHost.gen.go

package go_xen_client

import (
	"reflect"

	"github.com/nilshell/xmlrpc"
)

//ClusterHost: Cluster member metadata
type ClusterHost struct {
	Uuid              string                          // Unique identifier/object reference
	Cluster           string                          // Reference to the Cluster object
	Host              string                          // Reference to the Host object
	Enabled           bool                            // Whether the cluster host believes that clustering should be enabled on this host
	PIF               string                          // Reference to the PIF object
	Joined            bool                            // Whether the cluster host has joined the cluster
	AllowedOperations []ClusterHostOperation          // list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	CurrentOperations map[string]ClusterHostOperation // links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	OtherConfig       map[string]string               // Additional configuration

}

func FromClusterHostToXml(Cluster_host *ClusterHost) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] = Cluster_host.Uuid

	result["cluster"] = Cluster_host.Cluster

	result["host"] = Cluster_host.Host

	result["enabled"] = Cluster_host.Enabled

	result["PIF"] = Cluster_host.PIF

	result["joined"] = Cluster_host.Joined

	result["allowed_operations"] = Cluster_host.AllowedOperations

	current_operations := make(xmlrpc.Struct)
	for key, value := range Cluster_host.CurrentOperations {
		current_operations[key] = value
	}
	result["current_operations"] = current_operations

	other_config := make(xmlrpc.Struct)
	for key, value := range Cluster_host.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	return result
}

func ToClusterHost(obj interface{}) (resultObj *ClusterHost) {

	objValue := reflect.ValueOf(obj)
	resultObj = &ClusterHost{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.Uuid = v
			}
		case "cluster":
			if v, ok := keyValue.(string); ok {
				resultObj.Cluster = v
			}
		case "host":
			if v, ok := keyValue.(string); ok {
				resultObj.Host = v
			}
		case "enabled":
			if v, ok := keyValue.(bool); ok {
				resultObj.Enabled = v
			}
		case "PIF":
			if v, ok := keyValue.(string); ok {
				resultObj.PIF = v
			}
		case "joined":
			if v, ok := keyValue.(bool); ok {
				resultObj.Joined = v
			}
		case "allowed_operations":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.AllowedOperations = make([]ClusterHostOperation, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(ClusterHostOperation); ok {
						resultObj.AllowedOperations[i] = v
					}
				}
			}
		case "current_operations":

			resultObj.CurrentOperations = map[string]ClusterHostOperation{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.CurrentOperations[mapKeyName] = ToClusterHostOperation(v)
				} else {
					resultObj.CurrentOperations[mapKeyName] = 0
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
		}
	}

	return resultObj
}

/* GetAllRecords: Return a map of Cluster_host references to Cluster_host records for all Cluster_hosts known to the system. */
func (client *XenClient) ClusterHostGetAllRecords() (result map[string]ClusterHost, err error) {
	obj, err := client.APICall("Cluster_host.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]ClusterHost{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToClusterHost(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the Cluster_hosts known to the system. */
func (client *XenClient) ClusterHostGetAll() (result []string, err error) {
	obj, err := client.APICall("Cluster_host.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* Disable: Disable cluster membership for an enabled cluster host. */
func (client *XenClient) ClusterHostDisable(self string) (err error) {
	_, err = client.APICall("Cluster_host.disable", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* ForceDestroy: Remove a host from an existing cluster forcefully. */
func (client *XenClient) ClusterHostForceDestroy(self string) (err error) {
	_, err = client.APICall("Cluster_host.force_destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Enable: Enable cluster membership for a disabled cluster host. */
func (client *XenClient) ClusterHostEnable(self string) (err error) {
	_, err = client.APICall("Cluster_host.enable", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Destroy: Remove a host from an existing cluster. */
func (client *XenClient) ClusterHostDestroy(self string) (err error) {
	_, err = client.APICall("Cluster_host.destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Create: Add a new host to an existing cluster. */
func (client *XenClient) ClusterHostCreate(cluster string, host string, pif string) (result string, err error) {
	obj, err := client.APICall("Cluster_host.create", cluster, host, pif)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetOtherConfig: Get the other_config field of the given Cluster_host. */
func (client *XenClient) ClusterHostGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("Cluster_host.get_other_config", self)
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

/* GetCurrentOperations: Get the current_operations field of the given Cluster_host. */
func (client *XenClient) ClusterHostGetCurrentOperations(self string) (result map[string]ClusterHostOperation, err error) {
	obj, err := client.APICall("Cluster_host.get_current_operations", self)
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]ClusterHostOperation{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToClusterHostOperation(obj.String())
		result[key.String()] = mapObj
	}

	return
}

/* GetAllowedOperations: Get the allowed_operations field of the given Cluster_host. */
func (client *XenClient) ClusterHostGetAllowedOperations(self string) (result []ClusterHostOperation, err error) {
	obj, err := client.APICall("Cluster_host.get_allowed_operations", self)
	if err != nil {
		return
	}

	result = make([]ClusterHostOperation, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = ToClusterHostOperation(value.(string))
	}

	return
}

/* GetJoined: Get the joined field of the given Cluster_host. */
func (client *XenClient) ClusterHostGetJoined(self string) (result bool, err error) {
	obj, err := client.APICall("Cluster_host.get_joined", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetPIF: Get the PIF field of the given Cluster_host. */
func (client *XenClient) ClusterHostGetPIF(self string) (result string, err error) {
	obj, err := client.APICall("Cluster_host.get_PIF", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetEnabled: Get the enabled field of the given Cluster_host. */
func (client *XenClient) ClusterHostGetEnabled(self string) (result bool, err error) {
	obj, err := client.APICall("Cluster_host.get_enabled", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetHost: Get the host field of the given Cluster_host. */
func (client *XenClient) ClusterHostGetHost(self string) (result string, err error) {
	obj, err := client.APICall("Cluster_host.get_host", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetCluster: Get the cluster field of the given Cluster_host. */
func (client *XenClient) ClusterHostGetCluster(self string) (result string, err error) {
	obj, err := client.APICall("Cluster_host.get_cluster", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given Cluster_host. */
func (client *XenClient) ClusterHostGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("Cluster_host.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the Cluster_host instance with the specified UUID. */
func (client *XenClient) ClusterHostGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("Cluster_host.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given Cluster_host. */
func (client *XenClient) ClusterHostGetRecord(self string) (result ClusterHost, err error) {
	obj, err := client.APICall("Cluster_host.get_record", self)
	if err != nil {
		return
	}

	result = *ToClusterHost(obj.(interface{}))

	return
}

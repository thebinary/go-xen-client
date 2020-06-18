// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w Cluster.gen.go

package go_xen_client

import (
	"reflect"

	"github.com/nilshell/xmlrpc"
)

//Cluster: Cluster-wide Cluster metadata
type Cluster struct {
	Uuid                    string                      // Unique identifier/object reference
	ClusterHosts            []string                    // A list of the cluster_host objects associated with the Cluster
	PendingForget           []string                    // Internal field used by Host.destroy to store the IP of cluster members marked as permanently dead but not yet removed
	ClusterToken            string                      // The secret key used by xapi-clusterd when it talks to itself on other hosts
	ClusterStack            string                      // Simply the string 'corosync'. No other cluster stacks are currently supported
	AllowedOperations       []ClusterOperation          // list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	CurrentOperations       map[string]ClusterOperation // links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	PoolAutoJoin            bool                        // True if automatically joining new pool members to the cluster. This will be `true` in the first release
	TokenTimeout            float32                     // The corosync token timeout in seconds
	TokenTimeoutCoefficient float32                     // The corosync token timeout coefficient in seconds
	ClusterConfig           map[string]string           // Contains read-only settings for the cluster, such as timeouts and other options. It can only be set at cluster create time
	OtherConfig             map[string]string           // Additional configuration

}

func FromClusterToXml(Cluster *Cluster) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] =

		Cluster.Uuid

	result["cluster_hosts"] =

		Cluster.ClusterHosts

	result["pending_forget"] =

		Cluster.PendingForget

	result["cluster_token"] =

		Cluster.ClusterToken

	result["cluster_stack"] =

		Cluster.ClusterStack

	result["allowed_operations"] =

		Cluster.AllowedOperations

	current_operations := make(xmlrpc.Struct)
	for key, value := range Cluster.CurrentOperations {
		current_operations[key] = value
	}
	result["current_operations"] = current_operations

	result["pool_auto_join"] =

		Cluster.PoolAutoJoin

	result["token_timeout"] =

		Cluster.TokenTimeout

	result["token_timeout_coefficient"] =

		Cluster.TokenTimeoutCoefficient

	cluster_config := make(xmlrpc.Struct)
	for key, value := range Cluster.ClusterConfig {
		cluster_config[key] = value
	}
	result["cluster_config"] = cluster_config

	other_config := make(xmlrpc.Struct)
	for key, value := range Cluster.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	return result
}

func ToCluster(obj interface{}) (resultObj *Cluster) {

	objValue := reflect.ValueOf(obj)
	resultObj = &Cluster{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.Uuid = v
			}
		case "cluster_hosts":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.ClusterHosts = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.ClusterHosts[i] = v
					}
				}
			}
		case "pending_forget":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.PendingForget = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.PendingForget[i] = v
					}
				}
			}
		case "cluster_token":
			if v, ok := keyValue.(string); ok {
				resultObj.ClusterToken = v
			}
		case "cluster_stack":
			if v, ok := keyValue.(string); ok {
				resultObj.ClusterStack = v
			}
		case "allowed_operations":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.AllowedOperations = make([]ClusterOperation, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(ClusterOperation); ok {
						resultObj.AllowedOperations[i] = v
					}
				}
			}
		case "current_operations":

			resultObj.CurrentOperations = map[string]ClusterOperation{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.CurrentOperations[mapKeyName] = ToClusterOperation(v)
				} else {
					resultObj.CurrentOperations[mapKeyName] = 0
				}
			}
		case "pool_auto_join":
			if v, ok := keyValue.(bool); ok {
				resultObj.PoolAutoJoin = v
			}
		case "token_timeout":
			if v, ok := keyValue.(float32); ok {
				resultObj.TokenTimeout = v
			}
		case "token_timeout_coefficient":
			if v, ok := keyValue.(float32); ok {
				resultObj.TokenTimeoutCoefficient = v
			}
		case "cluster_config":

			resultObj.ClusterConfig = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.ClusterConfig[mapKeyName] = v
				} else {
					resultObj.ClusterConfig[mapKeyName] = ""
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

/* GetAllRecords: Return a map of Cluster references to Cluster records for all Clusters known to the system. */
func (client *XenClient) ClusterGetAllRecords() (result map[string]Cluster, err error) {
	obj, err := client.APICall("Cluster.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]Cluster{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToCluster(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the Clusters known to the system. */
func (client *XenClient) ClusterGetAll() (result []string, err error) {
	obj, err := client.APICall("Cluster.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* PoolResync: Resynchronise the cluster_host objects across the pool. Creates them where they need creating and then plugs them */
func (client *XenClient) ClusterPoolResync(self string) (err error) {
	_, err = client.APICall("Cluster.pool_resync", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* PoolDestroy: Attempt to destroy the Cluster_host objects for all hosts in the pool and then destroy the Cluster. */
func (client *XenClient) ClusterPoolDestroy(self string) (err error) {
	_, err = client.APICall("Cluster.pool_destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* PoolForceDestroy: Attempt to force destroy the Cluster_host objects, and then destroy the Cluster. */
func (client *XenClient) ClusterPoolForceDestroy(self string) (err error) {
	_, err = client.APICall("Cluster.pool_force_destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* PoolCreate: Attempt to create a Cluster from the entire pool */
func (client *XenClient) ClusterPoolCreate(network string, cluster_stack string, token_timeout float32, token_timeout_coefficient float32) (result string, err error) {
	obj, err := client.APICall("Cluster.pool_create", network, cluster_stack, token_timeout, token_timeout_coefficient)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetNetwork: Returns the network used by the cluster for inter-host communication, i.e. the network shared by all cluster host PIFs */
func (client *XenClient) ClusterGetNetwork(self string) (result string, err error) {
	obj, err := client.APICall("Cluster.get_network", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* Destroy: Destroys a Cluster object and the one remaining Cluster_host member */
func (client *XenClient) ClusterDestroy(self string) (err error) {
	_, err = client.APICall("Cluster.destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Create: Creates a Cluster object and one Cluster_host object as its first member */
func (client *XenClient) ClusterCreate(PIF string, cluster_stack string, pool_auto_join bool, token_timeout float32, token_timeout_coefficient float32) (result string, err error) {
	obj, err := client.APICall("Cluster.create", PIF, cluster_stack, pool_auto_join, token_timeout, token_timeout_coefficient)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given Cluster.  If the key is not in that Map, then do nothing. */
func (client *XenClient) ClusterRemoveFromOtherConfig(self string, key string) (err error) {
	_, err = client.APICall("Cluster.remove_from_other_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToOtherConfig: Add the given key-value pair to the other_config field of the given Cluster. */
func (client *XenClient) ClusterAddToOtherConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("Cluster.add_to_other_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOtherConfig: Set the other_config field of the given Cluster. */
func (client *XenClient) ClusterSetOtherConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("Cluster.set_other_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetOtherConfig: Get the other_config field of the given Cluster. */
func (client *XenClient) ClusterGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("Cluster.get_other_config", self)
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

/* GetClusterConfig: Get the cluster_config field of the given Cluster. */
func (client *XenClient) ClusterGetClusterConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("Cluster.get_cluster_config", self)
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

/* GetTokenTimeoutCoefficient: Get the token_timeout_coefficient field of the given Cluster. */
func (client *XenClient) ClusterGetTokenTimeoutCoefficient(self string) (result float32, err error) {
	obj, err := client.APICall("Cluster.get_token_timeout_coefficient", self)
	if err != nil {
		return
	}
	result = obj.(float32)
	return
}

/* GetTokenTimeout: Get the token_timeout field of the given Cluster. */
func (client *XenClient) ClusterGetTokenTimeout(self string) (result float32, err error) {
	obj, err := client.APICall("Cluster.get_token_timeout", self)
	if err != nil {
		return
	}
	result = obj.(float32)
	return
}

/* GetPoolAutoJoin: Get the pool_auto_join field of the given Cluster. */
func (client *XenClient) ClusterGetPoolAutoJoin(self string) (result bool, err error) {
	obj, err := client.APICall("Cluster.get_pool_auto_join", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetCurrentOperations: Get the current_operations field of the given Cluster. */
func (client *XenClient) ClusterGetCurrentOperations(self string) (result map[string]ClusterOperation, err error) {
	obj, err := client.APICall("Cluster.get_current_operations", self)
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]ClusterOperation{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToClusterOperation(obj.String())
		result[key.String()] = mapObj
	}

	return
}

/* GetAllowedOperations: Get the allowed_operations field of the given Cluster. */
func (client *XenClient) ClusterGetAllowedOperations(self string) (result []ClusterOperation, err error) {
	obj, err := client.APICall("Cluster.get_allowed_operations", self)
	if err != nil {
		return
	}

	result = make([]ClusterOperation, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = ToClusterOperation(value.(string))
	}

	return
}

/* GetClusterStack: Get the cluster_stack field of the given Cluster. */
func (client *XenClient) ClusterGetClusterStack(self string) (result string, err error) {
	obj, err := client.APICall("Cluster.get_cluster_stack", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetClusterToken: Get the cluster_token field of the given Cluster. */
func (client *XenClient) ClusterGetClusterToken(self string) (result string, err error) {
	obj, err := client.APICall("Cluster.get_cluster_token", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetPendingForget: Get the pending_forget field of the given Cluster. */
func (client *XenClient) ClusterGetPendingForget(self string) (result []string, err error) {
	obj, err := client.APICall("Cluster.get_pending_forget", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetClusterHosts: Get the cluster_hosts field of the given Cluster. */
func (client *XenClient) ClusterGetClusterHosts(self string) (result []string, err error) {
	obj, err := client.APICall("Cluster.get_cluster_hosts", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetUuid: Get the uuid field of the given Cluster. */
func (client *XenClient) ClusterGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("Cluster.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the Cluster instance with the specified UUID. */
func (client *XenClient) ClusterGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("Cluster.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given Cluster. */
func (client *XenClient) ClusterGetRecord(self string) (result Cluster, err error) {
	obj, err := client.APICall("Cluster.get_record", self)
	if err != nil {
		return
	}

	result = *ToCluster(obj.(interface{}))

	return
}

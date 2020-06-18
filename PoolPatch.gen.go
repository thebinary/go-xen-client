// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w PoolPatch.gen.go

package go_xen_client

import (
	"reflect"
	"strconv"

	"github.com/nilshell/xmlrpc"
)

//PoolPatch: Pool-wide patches
type PoolPatch struct {
	Uuid               string               // Unique identifier/object reference
	NameLabel          string               // a human-readable name
	NameDescription    string               // a notes field containing human-readable description
	Version            string               // Patch version number
	Size               int                  // Size of the patch
	PoolApplied        bool                 // This patch should be applied across the entire pool
	HostPatches        []string             // This hosts this patch is applied to.
	AfterApplyGuidance []AfterApplyGuidance // What the client should do after this patch has been applied.
	PoolUpdate         string               // A reference to the associated pool_update object
	OtherConfig        map[string]string    // additional configuration

}

func FromPoolPatchToXml(pool_patch *PoolPatch) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] =

		pool_patch.Uuid

	result["name_label"] =

		pool_patch.NameLabel

	result["name_description"] =

		pool_patch.NameDescription

	result["version"] =

		pool_patch.Version

	result["size"] =

		strconv.Itoa(pool_patch.Size)

	result["pool_applied"] =

		pool_patch.PoolApplied

	result["host_patches"] =

		pool_patch.HostPatches

	result["after_apply_guidance"] =

		pool_patch.AfterApplyGuidance

	result["pool_update"] =

		pool_patch.PoolUpdate

	other_config := make(xmlrpc.Struct)
	for key, value := range pool_patch.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	return result
}

func ToPoolPatch(obj interface{}) (resultObj *PoolPatch) {

	objValue := reflect.ValueOf(obj)
	resultObj = &PoolPatch{}

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
		case "version":
			if v, ok := keyValue.(string); ok {
				resultObj.Version = v
			}
		case "size":
			if v, ok := keyValue.(int); ok {
				resultObj.Size = v
			}
		case "pool_applied":
			if v, ok := keyValue.(bool); ok {
				resultObj.PoolApplied = v
			}
		case "host_patches":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.HostPatches = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.HostPatches[i] = v
					}
				}
			}
		case "after_apply_guidance":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.AfterApplyGuidance = make([]AfterApplyGuidance, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(AfterApplyGuidance); ok {
						resultObj.AfterApplyGuidance[i] = v
					}
				}
			}
		case "pool_update":
			if v, ok := keyValue.(string); ok {
				resultObj.PoolUpdate = v
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

/* GetAllRecords: Return a map of pool_patch references to pool_patch records for all pool_patchs known to the system. */
func (client *XenClient) PoolPatchGetAllRecords() (result map[string]PoolPatch, err error) {
	obj, err := client.APICall("pool_patch.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]PoolPatch{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToPoolPatch(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the pool_patchs known to the system. */
func (client *XenClient) PoolPatchGetAll() (result []string, err error) {
	obj, err := client.APICall("pool_patch.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* CleanOnHost: Removes the patch's files from the specified host */
func (client *XenClient) PoolPatchCleanOnHost(self string, host string) (err error) {
	_, err = client.APICall("pool_patch.clean_on_host", self, host)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Destroy: Removes the patch's files from all hosts in the pool, and removes the database entries.  Only works on unapplied patches. */
func (client *XenClient) PoolPatchDestroy(self string) (err error) {
	_, err = client.APICall("pool_patch.destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* PoolClean: Removes the patch's files from all hosts in the pool, but does not remove the database entries */
func (client *XenClient) PoolPatchPoolClean(self string) (err error) {
	_, err = client.APICall("pool_patch.pool_clean", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Clean: Removes the patch's files from the server */
func (client *XenClient) PoolPatchClean(self string) (err error) {
	_, err = client.APICall("pool_patch.clean", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Precheck: Execute the precheck stage of the selected patch on a host and return its output */
func (client *XenClient) PoolPatchPrecheck(self string, host string) (result string, err error) {
	obj, err := client.APICall("pool_patch.precheck", self, host)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* PoolApply: Apply the selected patch to all hosts in the pool and return a map of host_ref -> patch output */
func (client *XenClient) PoolPatchPoolApply(self string) (err error) {
	_, err = client.APICall("pool_patch.pool_apply", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Apply: Apply the selected patch to a host and return its output */
func (client *XenClient) PoolPatchApply(self string, host string) (result string, err error) {
	obj, err := client.APICall("pool_patch.apply", self, host)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given pool_patch.  If the key is not in that Map, then do nothing. */
func (client *XenClient) PoolPatchRemoveFromOtherConfig(self string, key string) (err error) {
	_, err = client.APICall("pool_patch.remove_from_other_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToOtherConfig: Add the given key-value pair to the other_config field of the given pool_patch. */
func (client *XenClient) PoolPatchAddToOtherConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("pool_patch.add_to_other_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOtherConfig: Set the other_config field of the given pool_patch. */
func (client *XenClient) PoolPatchSetOtherConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("pool_patch.set_other_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetOtherConfig: Get the other_config field of the given pool_patch. */
func (client *XenClient) PoolPatchGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("pool_patch.get_other_config", self)
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

/* GetPoolUpdate: Get the pool_update field of the given pool_patch. */
func (client *XenClient) PoolPatchGetPoolUpdate(self string) (result string, err error) {
	obj, err := client.APICall("pool_patch.get_pool_update", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetAfterApplyGuidance: Get the after_apply_guidance field of the given pool_patch. */
func (client *XenClient) PoolPatchGetAfterApplyGuidance(self string) (result []AfterApplyGuidance, err error) {
	obj, err := client.APICall("pool_patch.get_after_apply_guidance", self)
	if err != nil {
		return
	}

	result = make([]AfterApplyGuidance, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = ToAfterApplyGuidance(value.(string))
	}

	return
}

/* GetHostPatches: Get the host_patches field of the given pool_patch. */
func (client *XenClient) PoolPatchGetHostPatches(self string) (result []string, err error) {
	obj, err := client.APICall("pool_patch.get_host_patches", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetPoolApplied: Get the pool_applied field of the given pool_patch. */
func (client *XenClient) PoolPatchGetPoolApplied(self string) (result bool, err error) {
	obj, err := client.APICall("pool_patch.get_pool_applied", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetSize: Get the size field of the given pool_patch. */
func (client *XenClient) PoolPatchGetSize(self string) (result int, err error) {
	obj, err := client.APICall("pool_patch.get_size", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetVersion: Get the version field of the given pool_patch. */
func (client *XenClient) PoolPatchGetVersion(self string) (result string, err error) {
	obj, err := client.APICall("pool_patch.get_version", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetNameDescription: Get the name/description field of the given pool_patch. */
func (client *XenClient) PoolPatchGetNameDescription(self string) (result string, err error) {
	obj, err := client.APICall("pool_patch.get_name_description", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetNameLabel: Get the name/label field of the given pool_patch. */
func (client *XenClient) PoolPatchGetNameLabel(self string) (result string, err error) {
	obj, err := client.APICall("pool_patch.get_name_label", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given pool_patch. */
func (client *XenClient) PoolPatchGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("pool_patch.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByNameLabel: Get all the pool_patch instances with the given label. */
func (client *XenClient) PoolPatchGetByNameLabel(label string) (result []string, err error) {
	obj, err := client.APICall("pool_patch.get_by_name_label", label)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetByUuid: Get a reference to the pool_patch instance with the specified UUID. */
func (client *XenClient) PoolPatchGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("pool_patch.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given pool_patch. */
func (client *XenClient) PoolPatchGetRecord(self string) (result PoolPatch, err error) {
	obj, err := client.APICall("pool_patch.get_record", self)
	if err != nil {
		return
	}

	result = *ToPoolPatch(obj.(interface{}))

	return
}

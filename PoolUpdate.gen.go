// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w PoolUpdate.gen.go

package go_xen_client

import (
	"reflect"
	"strconv"

	"github.com/nilshell/xmlrpc"
)

//PoolUpdate: Pool-wide updates to the host software
type PoolUpdate struct {
	Uuid               string                     // Unique identifier/object reference
	NameLabel          string                     // a human-readable name
	NameDescription    string                     // a notes field containing human-readable description
	Version            string                     // Update version number
	InstallationSize   int                        // Size of the update in bytes
	Key                string                     // GPG key of the update
	AfterApplyGuidance []UpdateAfterApplyGuidance // What the client should do after this update has been applied.
	Vdi                string                     // VDI the update was uploaded to
	Hosts              []string                   // The hosts that have applied this update.
	OtherConfig        map[string]string          // additional configuration
	EnforceHomogeneity bool                       // Flag - if true, all hosts in a pool must apply this update

}

func FromPoolUpdateToXml(pool_update *PoolUpdate) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] = pool_update.Uuid

	result["name_label"] = pool_update.NameLabel

	result["name_description"] = pool_update.NameDescription

	result["version"] = pool_update.Version

	result["installation_size"] = strconv.Itoa(pool_update.InstallationSize)

	result["key"] = pool_update.Key

	result["after_apply_guidance"] = pool_update.AfterApplyGuidance

	result["vdi"] = pool_update.Vdi

	result["hosts"] = pool_update.Hosts

	other_config := make(xmlrpc.Struct)
	for key, value := range pool_update.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	result["enforce_homogeneity"] = pool_update.EnforceHomogeneity

	return result
}

func ToPoolUpdate(obj interface{}) (resultObj *PoolUpdate) {

	objValue := reflect.ValueOf(obj)
	resultObj = &PoolUpdate{}

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
		case "installation_size":
			if v, ok := keyValue.(int); ok {
				resultObj.InstallationSize = v
			}
		case "key":
			if v, ok := keyValue.(string); ok {
				resultObj.Key = v
			}
		case "after_apply_guidance":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.AfterApplyGuidance = make([]UpdateAfterApplyGuidance, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(UpdateAfterApplyGuidance); ok {
						resultObj.AfterApplyGuidance[i] = v
					}
				}
			}
		case "vdi":
			if v, ok := keyValue.(string); ok {
				resultObj.Vdi = v
			}
		case "hosts":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.Hosts = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.Hosts[i] = v
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
		case "enforce_homogeneity":
			if v, ok := keyValue.(bool); ok {
				resultObj.EnforceHomogeneity = v
			}
		}
	}

	return resultObj
}

/* GetAllRecords: Return a map of pool_update references to pool_update records for all pool_updates known to the system. */
func (client *XenClient) PoolUpdateGetAllRecords() (result map[string]PoolUpdate, err error) {
	obj, err := client.APICall("pool_update.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]PoolUpdate{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToPoolUpdate(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the pool_updates known to the system. */
func (client *XenClient) PoolUpdateGetAll() (result []string, err error) {
	obj, err := client.APICall("pool_update.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* Destroy: Removes the database entry. Only works on unapplied update. */
func (client *XenClient) PoolUpdateDestroy(self string) (err error) {
	_, err = client.APICall("pool_update.destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* PoolClean: Removes the update's files from all hosts in the pool, but does not revert the update */
func (client *XenClient) PoolUpdatePoolClean(self string) (err error) {
	_, err = client.APICall("pool_update.pool_clean", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* PoolApply: Apply the selected update to all hosts in the pool */
func (client *XenClient) PoolUpdatePoolApply(self string) (err error) {
	_, err = client.APICall("pool_update.pool_apply", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Apply: Apply the selected update to a host */
func (client *XenClient) PoolUpdateApply(self string, host string) (err error) {
	_, err = client.APICall("pool_update.apply", self, host)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Precheck: Execute the precheck stage of the selected update on a host */
func (client *XenClient) PoolUpdatePrecheck(self string, host string) (result LivepatchStatus, err error) {
	obj, err := client.APICall("pool_update.precheck", self, host)
	if err != nil {
		return
	}

	result = ToLivepatchStatus(obj.(string))

	return
}

/* Introduce: Introduce update VDI */
func (client *XenClient) PoolUpdateIntroduce(vdi string) (result string, err error) {
	obj, err := client.APICall("pool_update.introduce", vdi)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given pool_update.  If the key is not in that Map, then do nothing. */
func (client *XenClient) PoolUpdateRemoveFromOtherConfig(self string, key string) (err error) {
	_, err = client.APICall("pool_update.remove_from_other_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToOtherConfig: Add the given key-value pair to the other_config field of the given pool_update. */
func (client *XenClient) PoolUpdateAddToOtherConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("pool_update.add_to_other_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOtherConfig: Set the other_config field of the given pool_update. */
func (client *XenClient) PoolUpdateSetOtherConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("pool_update.set_other_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetEnforceHomogeneity: Get the enforce_homogeneity field of the given pool_update. */
func (client *XenClient) PoolUpdateGetEnforceHomogeneity(self string) (result bool, err error) {
	obj, err := client.APICall("pool_update.get_enforce_homogeneity", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetOtherConfig: Get the other_config field of the given pool_update. */
func (client *XenClient) PoolUpdateGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("pool_update.get_other_config", self)
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

/* GetHosts: Get the hosts field of the given pool_update. */
func (client *XenClient) PoolUpdateGetHosts(self string) (result []string, err error) {
	obj, err := client.APICall("pool_update.get_hosts", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetVdi: Get the vdi field of the given pool_update. */
func (client *XenClient) PoolUpdateGetVdi(self string) (result string, err error) {
	obj, err := client.APICall("pool_update.get_vdi", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetAfterApplyGuidance: Get the after_apply_guidance field of the given pool_update. */
func (client *XenClient) PoolUpdateGetAfterApplyGuidance(self string) (result []UpdateAfterApplyGuidance, err error) {
	obj, err := client.APICall("pool_update.get_after_apply_guidance", self)
	if err != nil {
		return
	}

	result = make([]UpdateAfterApplyGuidance, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = ToUpdateAfterApplyGuidance(value.(string))
	}

	return
}

/* GetKey: Get the key field of the given pool_update. */
func (client *XenClient) PoolUpdateGetKey(self string) (result string, err error) {
	obj, err := client.APICall("pool_update.get_key", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetInstallationSize: Get the installation_size field of the given pool_update. */
func (client *XenClient) PoolUpdateGetInstallationSize(self string) (result int, err error) {
	obj, err := client.APICall("pool_update.get_installation_size", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetVersion: Get the version field of the given pool_update. */
func (client *XenClient) PoolUpdateGetVersion(self string) (result string, err error) {
	obj, err := client.APICall("pool_update.get_version", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetNameDescription: Get the name/description field of the given pool_update. */
func (client *XenClient) PoolUpdateGetNameDescription(self string) (result string, err error) {
	obj, err := client.APICall("pool_update.get_name_description", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetNameLabel: Get the name/label field of the given pool_update. */
func (client *XenClient) PoolUpdateGetNameLabel(self string) (result string, err error) {
	obj, err := client.APICall("pool_update.get_name_label", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given pool_update. */
func (client *XenClient) PoolUpdateGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("pool_update.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByNameLabel: Get all the pool_update instances with the given label. */
func (client *XenClient) PoolUpdateGetByNameLabel(label string) (result []string, err error) {
	obj, err := client.APICall("pool_update.get_by_name_label", label)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetByUuid: Get a reference to the pool_update instance with the specified UUID. */
func (client *XenClient) PoolUpdateGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("pool_update.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given pool_update. */
func (client *XenClient) PoolUpdateGetRecord(self string) (result PoolUpdate, err error) {
	obj, err := client.APICall("pool_update.get_record", self)
	if err != nil {
		return
	}

	result = *ToPoolUpdate(obj.(interface{}))

	return
}

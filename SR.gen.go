// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w SR.gen.go

package go_xen_client

import (
	"reflect"
	"strconv"

	"github.com/nilshell/xmlrpc"
)

//SR: A storage repository
type SR struct {
	Uuid                string                       // Unique identifier/object reference
	NameLabel           string                       // a human-readable name
	NameDescription     string                       // a notes field containing human-readable description
	AllowedOperations   []StorageOperations          // list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	CurrentOperations   map[string]StorageOperations // links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	VDIs                []string                     // all virtual disks known to this storage repository
	PBDs                []string                     // describes how particular hosts can see this storage repository
	VirtualAllocation   int                          // sum of virtual_sizes of all VDIs in this storage repository (in bytes)
	PhysicalUtilisation int                          // physical space currently utilised on this storage repository (in bytes). Note that for sparse disk formats, physical_utilisation may be less than virtual_allocation
	PhysicalSize        int                          // total physical size of the repository (in bytes)
	Type                string                       // type of the storage repository
	ContentType         string                       // the type of the SR's content, if required (e.g. ISOs)
	Shared              bool                         // true if this SR is (capable of being) shared between multiple hosts
	OtherConfig         map[string]string            // additional configuration
	Tags                []string                     // user-specified tags for categorization purposes
	SmConfig            map[string]string            // SM dependent data
	Blobs               map[string]string            // Binary blobs associated with this SR
	LocalCacheEnabled   bool                         // True if this SR is assigned to be the local cache for its host
	IntroducedBy        string                       // The disaster recovery task which introduced this SR
	Clustered           bool                         // True if the SR is using aggregated local storage
	IsToolsSr           bool                         // True if this is the SR that contains the Tools ISO VDIs

}

func FromSRToXml(SR *SR) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] =

		SR.Uuid

	result["name_label"] =

		SR.NameLabel

	result["name_description"] =

		SR.NameDescription

	result["allowed_operations"] =

		SR.AllowedOperations

	current_operations := make(xmlrpc.Struct)
	for key, value := range SR.CurrentOperations {
		current_operations[key] = value
	}
	result["current_operations"] = current_operations

	result["VDIs"] =

		SR.VDIs

	result["PBDs"] =

		SR.PBDs

	result["virtual_allocation"] =

		strconv.Itoa(SR.VirtualAllocation)

	result["physical_utilisation"] =

		strconv.Itoa(SR.PhysicalUtilisation)

	result["physical_size"] =

		strconv.Itoa(SR.PhysicalSize)

	result["type"] =

		SR.Type

	result["content_type"] =

		SR.ContentType

	result["shared"] =

		SR.Shared

	other_config := make(xmlrpc.Struct)
	for key, value := range SR.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	result["tags"] =

		SR.Tags

	sm_config := make(xmlrpc.Struct)
	for key, value := range SR.SmConfig {
		sm_config[key] = value
	}
	result["sm_config"] = sm_config

	blobs := make(xmlrpc.Struct)
	for key, value := range SR.Blobs {
		blobs[key] = value
	}
	result["blobs"] = blobs

	result["local_cache_enabled"] =

		SR.LocalCacheEnabled

	result["introduced_by"] =

		SR.IntroducedBy

	result["clustered"] =

		SR.Clustered

	result["is_tools_sr"] =

		SR.IsToolsSr

	return result
}

func ToSR(obj interface{}) (resultObj *SR) {

	objValue := reflect.ValueOf(obj)
	resultObj = &SR{}

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
				resultObj.AllowedOperations = make([]StorageOperations, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(StorageOperations); ok {
						resultObj.AllowedOperations[i] = v
					}
				}
			}
		case "current_operations":

			resultObj.CurrentOperations = map[string]StorageOperations{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.CurrentOperations[mapKeyName] = ToStorageOperations(v)
				} else {
					resultObj.CurrentOperations[mapKeyName] = 0
				}
			}
		case "VDIs":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.VDIs = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.VDIs[i] = v
					}
				}
			}
		case "PBDs":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.PBDs = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.PBDs[i] = v
					}
				}
			}
		case "virtual_allocation":
			if v, ok := keyValue.(int); ok {
				resultObj.VirtualAllocation = v
			}
		case "physical_utilisation":
			if v, ok := keyValue.(int); ok {
				resultObj.PhysicalUtilisation = v
			}
		case "physical_size":
			if v, ok := keyValue.(int); ok {
				resultObj.PhysicalSize = v
			}
		case "type":
			if v, ok := keyValue.(string); ok {
				resultObj.Type = v
			}
		case "content_type":
			if v, ok := keyValue.(string); ok {
				resultObj.ContentType = v
			}
		case "shared":
			if v, ok := keyValue.(bool); ok {
				resultObj.Shared = v
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
		case "tags":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.Tags = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.Tags[i] = v
					}
				}
			}
		case "sm_config":

			resultObj.SmConfig = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.SmConfig[mapKeyName] = v
				} else {
					resultObj.SmConfig[mapKeyName] = ""
				}
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
		case "local_cache_enabled":
			if v, ok := keyValue.(bool); ok {
				resultObj.LocalCacheEnabled = v
			}
		case "introduced_by":
			if v, ok := keyValue.(string); ok {
				resultObj.IntroducedBy = v
			}
		case "clustered":
			if v, ok := keyValue.(bool); ok {
				resultObj.Clustered = v
			}
		case "is_tools_sr":
			if v, ok := keyValue.(bool); ok {
				resultObj.IsToolsSr = v
			}
		}
	}

	return resultObj
}

/* GetAllRecords: Return a map of SR references to SR records for all SRs known to the system. */
func (client *XenClient) SRGetAllRecords() (result map[string]SR, err error) {
	obj, err := client.APICall("SR.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]SR{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToSR(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the SRs known to the system. */
func (client *XenClient) SRGetAll() (result []string, err error) {
	obj, err := client.APICall("SR.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* ForgetDataSourceArchives: Forget the recorded statistics related to the specified data source */
func (client *XenClient) SRForgetDataSourceArchives(sr string, data_source string) (err error) {
	_, err = client.APICall("SR.forget_data_source_archives", sr, data_source)
	if err != nil {
		return
	}
	// no return result
	return
}

/* QueryDataSource: Query the latest value of the specified data source */
func (client *XenClient) SRQueryDataSource(sr string, data_source string) (result float32, err error) {
	obj, err := client.APICall("SR.query_data_source", sr, data_source)
	if err != nil {
		return
	}
	result = obj.(float32)
	return
}

/* RecordDataSource: Start recording the specified data source */
func (client *XenClient) SRRecordDataSource(sr string, data_source string) (err error) {
	_, err = client.APICall("SR.record_data_source", sr, data_source)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetDataSources:  */
func (client *XenClient) SRGetDataSources(sr string) (result []DataSource, err error) {
	obj, err := client.APICall("SR.get_data_sources", sr)
	if err != nil {
		return
	}

	result = make([]DataSource, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = *ToDataSource(value)
	}

	return
}

/* DisableDatabaseReplication:  */
func (client *XenClient) SRDisableDatabaseReplication(sr string) (err error) {
	_, err = client.APICall("SR.disable_database_replication", sr)
	if err != nil {
		return
	}
	// no return result
	return
}

/* EnableDatabaseReplication:  */
func (client *XenClient) SREnableDatabaseReplication(sr string) (err error) {
	_, err = client.APICall("SR.enable_database_replication", sr)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AssertSupportsDatabaseReplication: Returns successfully if the given SR supports database replication. Otherwise returns an error to explain why not. */
func (client *XenClient) SRAssertSupportsDatabaseReplication(sr string) (err error) {
	_, err = client.APICall("SR.assert_supports_database_replication", sr)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AssertCanHostHaStatefile: Returns successfully if the given SR can host an HA statefile. Otherwise returns an error to explain why not */
func (client *XenClient) SRAssertCanHostHaStatefile(sr string) (err error) {
	_, err = client.APICall("SR.assert_can_host_ha_statefile", sr)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetPhysicalSize: Sets the SR's physical_size field */
func (client *XenClient) SRSetPhysicalSize(self string, value int) (err error) {
	_, err = client.APICall("SR.set_physical_size", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* CreateNewBlob: Create a placeholder for a named binary blob of data that is associated with this SR */
func (client *XenClient) SRCreateNewBlob(sr string, name string, mime_type string, public bool) (result string, err error) {
	obj, err := client.APICall("SR.create_new_blob", sr, name, mime_type, public)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* SetNameDescription: Set the name description of the SR */
func (client *XenClient) SRSetNameDescription(sr string, value string) (err error) {
	_, err = client.APICall("SR.set_name_description", sr, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetNameLabel: Set the name label of the SR */
func (client *XenClient) SRSetNameLabel(sr string, value string) (err error) {
	_, err = client.APICall("SR.set_name_label", sr, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetShared: Sets the shared flag on the SR */
func (client *XenClient) SRSetShared(sr string, value bool) (err error) {
	_, err = client.APICall("SR.set_shared", sr, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* ProbeExt: Perform a backend-specific scan, using the given device_config.  If the device_config is complete, then this will return a list of the SRs present of this type on the device, if any.  If the device_config is partial, then a backend-specific scan will be performed, returning results that will guide the user in improving the device_config. */
func (client *XenClient) SRProbeExt(host string, device_config map[string]string, xtype string, sm_config map[string]string) (result []ProbeResult, err error) {
	obj, err := client.APICall("SR.probe_ext", host, device_config, xtype, sm_config)
	if err != nil {
		return
	}

	result = make([]ProbeResult, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = *ToProbeResult(value)
	}

	return
}

/* Probe: Perform a backend-specific scan, using the given device_config.  If the device_config is complete, then this will return a list of the SRs present of this type on the device, if any.  If the device_config is partial, then a backend-specific scan will be performed, returning results that will guide the user in improving the device_config. */
func (client *XenClient) SRProbe(host string, device_config map[string]string, xtype string, sm_config map[string]string) (result string, err error) {
	obj, err := client.APICall("SR.probe", host, device_config, xtype, sm_config)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* Scan: Refreshes the list of VDIs associated with an SR */
func (client *XenClient) SRScan(sr string) (err error) {
	_, err = client.APICall("SR.scan", sr)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetSupportedTypes: Return a set of all the SR types supported by the system */
func (client *XenClient) SRGetSupportedTypes() (result []string, err error) {
	obj, err := client.APICall("SR.get_supported_types")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* Update: Refresh the fields on the SR object */
func (client *XenClient) SRUpdate(sr string) (err error) {
	_, err = client.APICall("SR.update", sr)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Forget: Removing specified SR-record from database, without attempting to remove SR from disk */
func (client *XenClient) SRForget(sr string) (err error) {
	_, err = client.APICall("SR.forget", sr)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Destroy: Destroy specified SR, removing SR-record from database and remove SR from disk. (In order to affect this operation the appropriate device_config is read from the specified SR's PBD on current host) */
func (client *XenClient) SRDestroy(sr string) (err error) {
	_, err = client.APICall("SR.destroy", sr)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Make: Create a new Storage Repository on disk. This call is deprecated: use SR.create instead. */
func (client *XenClient) SRMake(host string, device_config map[string]string, physical_size int, name_label string, name_description string, xtype string, content_type string, sm_config map[string]string) (result string, err error) {
	obj, err := client.APICall("SR.make", host, device_config, physical_size, name_label, name_description, xtype, content_type, sm_config)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* Introduce: Introduce a new Storage Repository into the managed system */
func (client *XenClient) SRIntroduce(uuid string, name_label string, name_description string, xtype string, content_type string, shared bool, sm_config map[string]string) (result string, err error) {
	obj, err := client.APICall("SR.introduce", uuid, name_label, name_description, xtype, content_type, shared, sm_config)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* Create: Create a new Storage Repository and introduce it into the managed system, creating both SR record and PBD record to attach it to current host (with specified device_config parameters) */
func (client *XenClient) SRCreate(host string, device_config map[string]string, physical_size int, name_label string, name_description string, xtype string, content_type string, shared bool, sm_config map[string]string) (result string, err error) {
	obj, err := client.APICall("SR.create", host, device_config, physical_size, name_label, name_description, xtype, content_type, shared, sm_config)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* RemoveFromSmConfig: Remove the given key and its corresponding value from the sm_config field of the given SR.  If the key is not in that Map, then do nothing. */
func (client *XenClient) SRRemoveFromSmConfig(self string, key string) (err error) {
	_, err = client.APICall("SR.remove_from_sm_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToSmConfig: Add the given key-value pair to the sm_config field of the given SR. */
func (client *XenClient) SRAddToSmConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("SR.add_to_sm_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetSmConfig: Set the sm_config field of the given SR. */
func (client *XenClient) SRSetSmConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("SR.set_sm_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveTags: Remove the given value from the tags field of the given SR.  If the value is not in that Set, then do nothing. */
func (client *XenClient) SRRemoveTags(self string, value string) (err error) {
	_, err = client.APICall("SR.remove_tags", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddTags: Add the given value to the tags field of the given SR.  If the value is already in that Set, then do nothing. */
func (client *XenClient) SRAddTags(self string, value string) (err error) {
	_, err = client.APICall("SR.add_tags", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetTags: Set the tags field of the given SR. */
func (client *XenClient) SRSetTags(self string, value []string) (err error) {
	_, err = client.APICall("SR.set_tags", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given SR.  If the key is not in that Map, then do nothing. */
func (client *XenClient) SRRemoveFromOtherConfig(self string, key string) (err error) {
	_, err = client.APICall("SR.remove_from_other_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToOtherConfig: Add the given key-value pair to the other_config field of the given SR. */
func (client *XenClient) SRAddToOtherConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("SR.add_to_other_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOtherConfig: Set the other_config field of the given SR. */
func (client *XenClient) SRSetOtherConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("SR.set_other_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetIsToolsSr: Get the is_tools_sr field of the given SR. */
func (client *XenClient) SRGetIsToolsSr(self string) (result bool, err error) {
	obj, err := client.APICall("SR.get_is_tools_sr", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetClustered: Get the clustered field of the given SR. */
func (client *XenClient) SRGetClustered(self string) (result bool, err error) {
	obj, err := client.APICall("SR.get_clustered", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetIntroducedBy: Get the introduced_by field of the given SR. */
func (client *XenClient) SRGetIntroducedBy(self string) (result string, err error) {
	obj, err := client.APICall("SR.get_introduced_by", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetLocalCacheEnabled: Get the local_cache_enabled field of the given SR. */
func (client *XenClient) SRGetLocalCacheEnabled(self string) (result bool, err error) {
	obj, err := client.APICall("SR.get_local_cache_enabled", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetBlobs: Get the blobs field of the given SR. */
func (client *XenClient) SRGetBlobs(self string) (result map[string]string, err error) {
	obj, err := client.APICall("SR.get_blobs", self)
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

/* GetSmConfig: Get the sm_config field of the given SR. */
func (client *XenClient) SRGetSmConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("SR.get_sm_config", self)
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

/* GetTags: Get the tags field of the given SR. */
func (client *XenClient) SRGetTags(self string) (result []string, err error) {
	obj, err := client.APICall("SR.get_tags", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetOtherConfig: Get the other_config field of the given SR. */
func (client *XenClient) SRGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("SR.get_other_config", self)
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

/* GetShared: Get the shared field of the given SR. */
func (client *XenClient) SRGetShared(self string) (result bool, err error) {
	obj, err := client.APICall("SR.get_shared", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetContentType: Get the content_type field of the given SR. */
func (client *XenClient) SRGetContentType(self string) (result string, err error) {
	obj, err := client.APICall("SR.get_content_type", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetType: Get the type field of the given SR. */
func (client *XenClient) SRGetType(self string) (result string, err error) {
	obj, err := client.APICall("SR.get_type", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetPhysicalSize: Get the physical_size field of the given SR. */
func (client *XenClient) SRGetPhysicalSize(self string) (result int, err error) {
	obj, err := client.APICall("SR.get_physical_size", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetPhysicalUtilisation: Get the physical_utilisation field of the given SR. */
func (client *XenClient) SRGetPhysicalUtilisation(self string) (result int, err error) {
	obj, err := client.APICall("SR.get_physical_utilisation", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetVirtualAllocation: Get the virtual_allocation field of the given SR. */
func (client *XenClient) SRGetVirtualAllocation(self string) (result int, err error) {
	obj, err := client.APICall("SR.get_virtual_allocation", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetPBDs: Get the PBDs field of the given SR. */
func (client *XenClient) SRGetPBDs(self string) (result []string, err error) {
	obj, err := client.APICall("SR.get_PBDs", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetVDIs: Get the VDIs field of the given SR. */
func (client *XenClient) SRGetVDIs(self string) (result []string, err error) {
	obj, err := client.APICall("SR.get_VDIs", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetCurrentOperations: Get the current_operations field of the given SR. */
func (client *XenClient) SRGetCurrentOperations(self string) (result map[string]StorageOperations, err error) {
	obj, err := client.APICall("SR.get_current_operations", self)
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]StorageOperations{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToStorageOperations(obj.String())
		result[key.String()] = mapObj
	}

	return
}

/* GetAllowedOperations: Get the allowed_operations field of the given SR. */
func (client *XenClient) SRGetAllowedOperations(self string) (result []StorageOperations, err error) {
	obj, err := client.APICall("SR.get_allowed_operations", self)
	if err != nil {
		return
	}

	result = make([]StorageOperations, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = ToStorageOperations(value.(string))
	}

	return
}

/* GetNameDescription: Get the name/description field of the given SR. */
func (client *XenClient) SRGetNameDescription(self string) (result string, err error) {
	obj, err := client.APICall("SR.get_name_description", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetNameLabel: Get the name/label field of the given SR. */
func (client *XenClient) SRGetNameLabel(self string) (result string, err error) {
	obj, err := client.APICall("SR.get_name_label", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given SR. */
func (client *XenClient) SRGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("SR.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByNameLabel: Get all the SR instances with the given label. */
func (client *XenClient) SRGetByNameLabel(label string) (result []string, err error) {
	obj, err := client.APICall("SR.get_by_name_label", label)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetByUuid: Get a reference to the SR instance with the specified UUID. */
func (client *XenClient) SRGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("SR.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given SR. */
func (client *XenClient) SRGetRecord(self string) (result SR, err error) {
	obj, err := client.APICall("SR.get_record", self)
	if err != nil {
		return
	}

	result = *ToSR(obj.(interface{}))

	return
}

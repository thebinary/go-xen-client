// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w VDI.gen.go

package go_xen_client

import (
	"reflect"
	"strconv"
	"time"

	"github.com/nilshell/xmlrpc"
)

//VDI: A virtual disk image
type VDI struct {
	Uuid                string                   // Unique identifier/object reference
	NameLabel           string                   // a human-readable name
	NameDescription     string                   // a notes field containing human-readable description
	AllowedOperations   []VdiOperations          // list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	CurrentOperations   map[string]VdiOperations // links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	SR                  string                   // storage repository in which the VDI resides
	VBDs                []string                 // list of vbds that refer to this disk
	CrashDumps          []string                 // list of crash dumps that refer to this disk
	VirtualSize         int                      // size of disk as presented to the guest (in bytes). Note that, depending on storage backend type, requested size may not be respected exactly
	PhysicalUtilisation int                      // amount of physical space that the disk image is currently taking up on the storage repository (in bytes)
	Type                VdiType                  // type of the VDI
	Sharable            bool                     // true if this disk may be shared
	ReadOnly            bool                     // true if this disk may ONLY be mounted read-only
	OtherConfig         map[string]string        // additional configuration
	StorageLock         bool                     // true if this disk is locked at the storage level
	Location            string                   // location information
	Managed             bool                     //
	Missing             bool                     // true if SR scan operation reported this VDI as not present on disk
	Parent              string                   // This field is always null. Deprecated
	XenstoreData        map[string]string        // data to be inserted into the xenstore tree (/local/domain/0/backend/vbd/<domid>/<device-id>/sm-data) after the VDI is attached. This is generally set by the SM backends on vdi_attach.
	SmConfig            map[string]string        // SM dependent data
	IsASnapshot         bool                     // true if this is a snapshot.
	SnapshotOf          string                   // Ref pointing to the VDI this snapshot is of.
	Snapshots           []string                 // List pointing to all the VDIs snapshots.
	SnapshotTime        time.Time                // Date/time when this snapshot was created.
	Tags                []string                 // user-specified tags for categorization purposes
	AllowCaching        bool                     // true if this VDI is to be cached in the local cache SR
	OnBoot              OnBoot                   // The behaviour of this VDI on a VM boot
	MetadataOfPool      string                   // The pool whose metadata is contained in this VDI
	MetadataLatest      bool                     // Whether this VDI contains the latest known accessible metadata for the pool
	IsToolsIso          bool                     // Whether this VDI is a Tools ISO
	CbtEnabled          bool                     // True if changed blocks are tracked for this VDI

}

func FromVDIToXml(VDI *VDI) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] = VDI.Uuid

	result["name_label"] = VDI.NameLabel

	result["name_description"] = VDI.NameDescription

	result["allowed_operations"] = VDI.AllowedOperations

	current_operations := make(xmlrpc.Struct)
	for key, value := range VDI.CurrentOperations {
		current_operations[key] = value
	}
	result["current_operations"] = current_operations

	result["SR"] = VDI.SR

	result["VBDs"] = VDI.VBDs

	result["crash_dumps"] = VDI.CrashDumps

	result["virtual_size"] = strconv.Itoa(VDI.VirtualSize)

	result["physical_utilisation"] = strconv.Itoa(VDI.PhysicalUtilisation)

	result["type"] = VDI.Type.String()

	result["sharable"] = VDI.Sharable

	result["read_only"] = VDI.ReadOnly

	other_config := make(xmlrpc.Struct)
	for key, value := range VDI.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	result["storage_lock"] = VDI.StorageLock

	result["location"] = VDI.Location

	result["managed"] = VDI.Managed

	result["missing"] = VDI.Missing

	result["parent"] = VDI.Parent

	xenstore_data := make(xmlrpc.Struct)
	for key, value := range VDI.XenstoreData {
		xenstore_data[key] = value
	}
	result["xenstore_data"] = xenstore_data

	sm_config := make(xmlrpc.Struct)
	for key, value := range VDI.SmConfig {
		sm_config[key] = value
	}
	result["sm_config"] = sm_config

	result["is_a_snapshot"] = VDI.IsASnapshot

	result["snapshot_of"] = VDI.SnapshotOf

	result["snapshots"] = VDI.Snapshots

	result["snapshot_time"] = VDI.SnapshotTime

	result["tags"] = VDI.Tags

	result["allow_caching"] = VDI.AllowCaching

	result["on_boot"] = VDI.OnBoot.String()

	result["metadata_of_pool"] = VDI.MetadataOfPool

	result["metadata_latest"] = VDI.MetadataLatest

	result["is_tools_iso"] = VDI.IsToolsIso

	result["cbt_enabled"] = VDI.CbtEnabled

	return result
}

func ToVDI(obj interface{}) (resultObj *VDI) {

	objValue := reflect.ValueOf(obj)
	resultObj = &VDI{}

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
				resultObj.AllowedOperations = make([]VdiOperations, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(VdiOperations); ok {
						resultObj.AllowedOperations[i] = v
					}
				}
			}
		case "current_operations":

			resultObj.CurrentOperations = map[string]VdiOperations{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.CurrentOperations[mapKeyName] = ToVdiOperations(v)
				} else {
					resultObj.CurrentOperations[mapKeyName] = 0
				}
			}
		case "SR":
			if v, ok := keyValue.(string); ok {
				resultObj.SR = v
			}
		case "VBDs":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.VBDs = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.VBDs[i] = v
					}
				}
			}
		case "crash_dumps":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.CrashDumps = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.CrashDumps[i] = v
					}
				}
			}
		case "virtual_size":
			if v, ok := keyValue.(int); ok {
				resultObj.VirtualSize = v
			}
		case "physical_utilisation":
			if v, ok := keyValue.(int); ok {
				resultObj.PhysicalUtilisation = v
			}
		case "type":
			if v, ok := keyValue.(VdiType); ok {
				resultObj.Type = v
			}
		case "sharable":
			if v, ok := keyValue.(bool); ok {
				resultObj.Sharable = v
			}
		case "read_only":
			if v, ok := keyValue.(bool); ok {
				resultObj.ReadOnly = v
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
		case "storage_lock":
			if v, ok := keyValue.(bool); ok {
				resultObj.StorageLock = v
			}
		case "location":
			if v, ok := keyValue.(string); ok {
				resultObj.Location = v
			}
		case "managed":
			if v, ok := keyValue.(bool); ok {
				resultObj.Managed = v
			}
		case "missing":
			if v, ok := keyValue.(bool); ok {
				resultObj.Missing = v
			}
		case "parent":
			if v, ok := keyValue.(string); ok {
				resultObj.Parent = v
			}
		case "xenstore_data":

			resultObj.XenstoreData = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.XenstoreData[mapKeyName] = v
				} else {
					resultObj.XenstoreData[mapKeyName] = ""
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
		case "is_a_snapshot":
			if v, ok := keyValue.(bool); ok {
				resultObj.IsASnapshot = v
			}
		case "snapshot_of":
			if v, ok := keyValue.(string); ok {
				resultObj.SnapshotOf = v
			}
		case "snapshots":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.Snapshots = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.Snapshots[i] = v
					}
				}
			}
		case "snapshot_time":
			if v, ok := keyValue.(time.Time); ok {
				resultObj.SnapshotTime = v
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
		case "allow_caching":
			if v, ok := keyValue.(bool); ok {
				resultObj.AllowCaching = v
			}
		case "on_boot":
			if v, ok := keyValue.(OnBoot); ok {
				resultObj.OnBoot = v
			}
		case "metadata_of_pool":
			if v, ok := keyValue.(string); ok {
				resultObj.MetadataOfPool = v
			}
		case "metadata_latest":
			if v, ok := keyValue.(bool); ok {
				resultObj.MetadataLatest = v
			}
		case "is_tools_iso":
			if v, ok := keyValue.(bool); ok {
				resultObj.IsToolsIso = v
			}
		case "cbt_enabled":
			if v, ok := keyValue.(bool); ok {
				resultObj.CbtEnabled = v
			}
		}
	}

	return resultObj
}

/* GetAllRecords: Return a map of VDI references to VDI records for all VDIs known to the system. */
func (client *XenClient) VDIGetAllRecords() (result map[string]VDI, err error) {
	obj, err := client.APICall("VDI.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]VDI{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToVDI(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the VDIs known to the system. */
func (client *XenClient) VDIGetAll() (result []string, err error) {
	obj, err := client.APICall("VDI.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetNbdInfo: Get details specifying how to access this VDI via a Network Block Device server. For each of a set of NBD server addresses on which the VDI is available, the return value set contains a vdi_nbd_server_info object that contains an exportname to request once the NBD connection is established, and connection details for the address. An empty list is returned if there is no network that has a PIF on a host with access to the relevant SR, or if no such network has been assigned an NBD-related purpose in its purpose field. To access the given VDI, any of the vdi_nbd_server_info objects can be used to make a connection to a server, and then the VDI will be available by requesting the exportname. */
func (client *XenClient) VDIGetNbdInfo(self string) (result []VdiNbdServerInfo, err error) {
	obj, err := client.APICall("VDI.get_nbd_info", self)
	if err != nil {
		return
	}

	result = make([]VdiNbdServerInfo, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = *ToVdiNbdServerInfo(value)
	}

	return
}

/* ListChangedBlocks: Compare two VDIs in 64k block increments and report which blocks differ. This operation is not allowed when vdi_to is attached to a VM. */
func (client *XenClient) VDIListChangedBlocks(vdi_from string, vdi_to string) (result string, err error) {
	obj, err := client.APICall("VDI.list_changed_blocks", vdi_from, vdi_to)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* DataDestroy: Delete the data of the snapshot VDI, but keep its changed block tracking metadata. When successful, this call changes the type of the VDI to cbt_metadata. This operation is idempotent: calling it on a VDI of type cbt_metadata results in a no-op, and no error will be thrown. */
func (client *XenClient) VDIDataDestroy(self string) (err error) {
	_, err = client.APICall("VDI.data_destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* DisableCbt: Disable changed block tracking for the VDI. This call is only allowed on VDIs that support enabling CBT. It is an idempotent operation - disabling CBT for a VDI for which CBT is not enabled results in a no-op, and no error will be thrown. */
func (client *XenClient) VDIDisableCbt(self string) (err error) {
	_, err = client.APICall("VDI.disable_cbt", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* EnableCbt: Enable changed block tracking for the VDI. This call is idempotent - enabling CBT for a VDI for which CBT is already enabled results in a no-op, and no error will be thrown. */
func (client *XenClient) VDIEnableCbt(self string) (err error) {
	_, err = client.APICall("VDI.enable_cbt", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* PoolMigrate: Migrate a VDI, which may be attached to a running guest, to a different SR. The destination SR must be visible to the guest. */
func (client *XenClient) VDIPoolMigrate(vdi string, sr string, options map[string]string) (result string, err error) {
	obj, err := client.APICall("VDI.pool_migrate", vdi, sr, options)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* ReadDatabasePoolUuid: Check the VDI cache for the pool UUID of the database on this VDI. */
func (client *XenClient) VDIReadDatabasePoolUuid(self string) (result string, err error) {
	obj, err := client.APICall("VDI.read_database_pool_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* OpenDatabase: Load the metadata found on the supplied VDI and return a session reference which can be used in API calls to query its contents. */
func (client *XenClient) VDIOpenDatabase(self string) (result string, err error) {
	obj, err := client.APICall("VDI.open_database", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* SetAllowCaching: Set the value of the allow_caching parameter. This value can only be changed when the VDI is not attached to a running VM. The caching behaviour is only affected by this flag for VHD-based VDIs that have one parent and no child VHDs. Moreover, caching only takes place when the host running the VM containing this VDI has a nominated SR for local caching. */
func (client *XenClient) VDISetAllowCaching(self string, value bool) (err error) {
	_, err = client.APICall("VDI.set_allow_caching", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOnBoot: Set the value of the on_boot parameter. This value can only be changed when the VDI is not attached to a running VM. */
func (client *XenClient) VDISetOnBoot(self string, value OnBoot) (err error) {
	_, err = client.APICall("VDI.set_on_boot", self, value.String())
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetNameDescription: Set the name description of the VDI. This can only happen when its SR is currently attached. */
func (client *XenClient) VDISetNameDescription(self string, value string) (err error) {
	_, err = client.APICall("VDI.set_name_description", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetNameLabel: Set the name label of the VDI. This can only happen when then its SR is currently attached. */
func (client *XenClient) VDISetNameLabel(self string, value string) (err error) {
	_, err = client.APICall("VDI.set_name_label", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetReadOnly: Sets the VDI's read_only field */
func (client *XenClient) VDISetReadOnly(self string, value bool) (err error) {
	_, err = client.APICall("VDI.set_read_only", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetSharable: Sets the VDI's sharable field */
func (client *XenClient) VDISetSharable(self string, value bool) (err error) {
	_, err = client.APICall("VDI.set_sharable", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Forget: Removes a VDI record from the database */
func (client *XenClient) VDIForget(vdi string) (err error) {
	_, err = client.APICall("VDI.forget", vdi)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Copy: Copy either a full VDI or the block differences between two VDIs into either a fresh VDI or an existing VDI. */
func (client *XenClient) VDICopy(vdi string, sr string, base_vdi string, into_vdi string) (result string, err error) {
	obj, err := client.APICall("VDI.copy", vdi, sr, base_vdi, into_vdi)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* Update: Ask the storage backend to refresh the fields in the VDI object */
func (client *XenClient) VDIUpdate(vdi string) (err error) {
	_, err = client.APICall("VDI.update", vdi)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Introduce: Create a new VDI record in the database only */
func (client *XenClient) VDIIntroduce(uuid string, name_label string, name_description string, SR string, xtype VdiType, sharable bool, read_only bool, other_config map[string]string, location string, xenstore_data map[string]string, sm_config map[string]string, managed bool, virtual_size int, physical_utilisation int, metadata_of_pool string, is_a_snapshot bool, snapshot_time time.Time, snapshot_of string) (result string, err error) {
	obj, err := client.APICall("VDI.introduce", uuid, name_label, name_description, SR, xtype, sharable, read_only, other_config, location, xenstore_data, sm_config, managed, virtual_size, physical_utilisation, metadata_of_pool, is_a_snapshot, snapshot_time, snapshot_of)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* ResizeOnline: Resize the VDI which may or may not be attached to running guests. */
func (client *XenClient) VDIResizeOnline(vdi string, size int) (err error) {
	_, err = client.APICall("VDI.resize_online", vdi, size)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Resize: Resize the VDI. */
func (client *XenClient) VDIResize(vdi string, size int) (err error) {
	_, err = client.APICall("VDI.resize", vdi, size)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Clone: Take an exact copy of the VDI and return a reference to the new disk. If any driver_params are specified then these are passed through to the storage-specific substrate driver that implements the clone operation. NB the clone lives in the same Storage Repository as its parent. */
func (client *XenClient) VDIClone(vdi string, driver_params map[string]string) (result string, err error) {
	obj, err := client.APICall("VDI.clone", vdi, driver_params)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* Snapshot: Take a read-only snapshot of the VDI, returning a reference to the snapshot. If any driver_params are specified then these are passed through to the storage-specific substrate driver that takes the snapshot. NB the snapshot lives in the same Storage Repository as its parent. */
func (client *XenClient) VDISnapshot(vdi string, driver_params map[string]string) (result string, err error) {
	obj, err := client.APICall("VDI.snapshot", vdi, driver_params)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* RemoveTags: Remove the given value from the tags field of the given VDI.  If the value is not in that Set, then do nothing. */
func (client *XenClient) VDIRemoveTags(self string, value string) (err error) {
	_, err = client.APICall("VDI.remove_tags", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddTags: Add the given value to the tags field of the given VDI.  If the value is already in that Set, then do nothing. */
func (client *XenClient) VDIAddTags(self string, value string) (err error) {
	_, err = client.APICall("VDI.add_tags", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetTags: Set the tags field of the given VDI. */
func (client *XenClient) VDISetTags(self string, value []string) (err error) {
	_, err = client.APICall("VDI.set_tags", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveFromSmConfig: Remove the given key and its corresponding value from the sm_config field of the given VDI.  If the key is not in that Map, then do nothing. */
func (client *XenClient) VDIRemoveFromSmConfig(self string, key string) (err error) {
	_, err = client.APICall("VDI.remove_from_sm_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToSmConfig: Add the given key-value pair to the sm_config field of the given VDI. */
func (client *XenClient) VDIAddToSmConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("VDI.add_to_sm_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetSmConfig: Set the sm_config field of the given VDI. */
func (client *XenClient) VDISetSmConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("VDI.set_sm_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveFromXenstoreData: Remove the given key and its corresponding value from the xenstore_data field of the given VDI.  If the key is not in that Map, then do nothing. */
func (client *XenClient) VDIRemoveFromXenstoreData(self string, key string) (err error) {
	_, err = client.APICall("VDI.remove_from_xenstore_data", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToXenstoreData: Add the given key-value pair to the xenstore_data field of the given VDI. */
func (client *XenClient) VDIAddToXenstoreData(self string, key string, value string) (err error) {
	_, err = client.APICall("VDI.add_to_xenstore_data", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetXenstoreData: Set the xenstore_data field of the given VDI. */
func (client *XenClient) VDISetXenstoreData(self string, value map[string]string) (err error) {
	_, err = client.APICall("VDI.set_xenstore_data", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given VDI.  If the key is not in that Map, then do nothing. */
func (client *XenClient) VDIRemoveFromOtherConfig(self string, key string) (err error) {
	_, err = client.APICall("VDI.remove_from_other_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToOtherConfig: Add the given key-value pair to the other_config field of the given VDI. */
func (client *XenClient) VDIAddToOtherConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("VDI.add_to_other_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOtherConfig: Set the other_config field of the given VDI. */
func (client *XenClient) VDISetOtherConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("VDI.set_other_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetCbtEnabled: Get the cbt_enabled field of the given VDI. */
func (client *XenClient) VDIGetCbtEnabled(self string) (result bool, err error) {
	obj, err := client.APICall("VDI.get_cbt_enabled", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetIsToolsIso: Get the is_tools_iso field of the given VDI. */
func (client *XenClient) VDIGetIsToolsIso(self string) (result bool, err error) {
	obj, err := client.APICall("VDI.get_is_tools_iso", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetMetadataLatest: Get the metadata_latest field of the given VDI. */
func (client *XenClient) VDIGetMetadataLatest(self string) (result bool, err error) {
	obj, err := client.APICall("VDI.get_metadata_latest", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetMetadataOfPool: Get the metadata_of_pool field of the given VDI. */
func (client *XenClient) VDIGetMetadataOfPool(self string) (result string, err error) {
	obj, err := client.APICall("VDI.get_metadata_of_pool", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetOnBoot: Get the on_boot field of the given VDI. */
func (client *XenClient) VDIGetOnBoot(self string) (result OnBoot, err error) {
	obj, err := client.APICall("VDI.get_on_boot", self)
	if err != nil {
		return
	}

	result = ToOnBoot(obj.(string))

	return
}

/* GetAllowCaching: Get the allow_caching field of the given VDI. */
func (client *XenClient) VDIGetAllowCaching(self string) (result bool, err error) {
	obj, err := client.APICall("VDI.get_allow_caching", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetTags: Get the tags field of the given VDI. */
func (client *XenClient) VDIGetTags(self string) (result []string, err error) {
	obj, err := client.APICall("VDI.get_tags", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetSnapshotTime: Get the snapshot_time field of the given VDI. */
func (client *XenClient) VDIGetSnapshotTime(self string) (result time.Time, err error) {
	obj, err := client.APICall("VDI.get_snapshot_time", self)
	if err != nil {
		return
	}
	result = obj.(time.Time)
	return
}

/* GetSnapshots: Get the snapshots field of the given VDI. */
func (client *XenClient) VDIGetSnapshots(self string) (result []string, err error) {
	obj, err := client.APICall("VDI.get_snapshots", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetSnapshotOf: Get the snapshot_of field of the given VDI. */
func (client *XenClient) VDIGetSnapshotOf(self string) (result string, err error) {
	obj, err := client.APICall("VDI.get_snapshot_of", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetIsASnapshot: Get the is_a_snapshot field of the given VDI. */
func (client *XenClient) VDIGetIsASnapshot(self string) (result bool, err error) {
	obj, err := client.APICall("VDI.get_is_a_snapshot", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetSmConfig: Get the sm_config field of the given VDI. */
func (client *XenClient) VDIGetSmConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VDI.get_sm_config", self)
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

/* GetXenstoreData: Get the xenstore_data field of the given VDI. */
func (client *XenClient) VDIGetXenstoreData(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VDI.get_xenstore_data", self)
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

/* GetParent: Get the parent field of the given VDI. */
func (client *XenClient) VDIGetParent(self string) (result string, err error) {
	obj, err := client.APICall("VDI.get_parent", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetMissing: Get the missing field of the given VDI. */
func (client *XenClient) VDIGetMissing(self string) (result bool, err error) {
	obj, err := client.APICall("VDI.get_missing", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetManaged: Get the managed field of the given VDI. */
func (client *XenClient) VDIGetManaged(self string) (result bool, err error) {
	obj, err := client.APICall("VDI.get_managed", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetLocation: Get the location field of the given VDI. */
func (client *XenClient) VDIGetLocation(self string) (result string, err error) {
	obj, err := client.APICall("VDI.get_location", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetStorageLock: Get the storage_lock field of the given VDI. */
func (client *XenClient) VDIGetStorageLock(self string) (result bool, err error) {
	obj, err := client.APICall("VDI.get_storage_lock", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetOtherConfig: Get the other_config field of the given VDI. */
func (client *XenClient) VDIGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VDI.get_other_config", self)
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

/* GetReadOnly: Get the read_only field of the given VDI. */
func (client *XenClient) VDIGetReadOnly(self string) (result bool, err error) {
	obj, err := client.APICall("VDI.get_read_only", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetSharable: Get the sharable field of the given VDI. */
func (client *XenClient) VDIGetSharable(self string) (result bool, err error) {
	obj, err := client.APICall("VDI.get_sharable", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetType: Get the type field of the given VDI. */
func (client *XenClient) VDIGetType(self string) (result VdiType, err error) {
	obj, err := client.APICall("VDI.get_type", self)
	if err != nil {
		return
	}

	result = ToVdiType(obj.(string))

	return
}

/* GetPhysicalUtilisation: Get the physical_utilisation field of the given VDI. */
func (client *XenClient) VDIGetPhysicalUtilisation(self string) (result int, err error) {
	obj, err := client.APICall("VDI.get_physical_utilisation", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetVirtualSize: Get the virtual_size field of the given VDI. */
func (client *XenClient) VDIGetVirtualSize(self string) (result int, err error) {
	obj, err := client.APICall("VDI.get_virtual_size", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetCrashDumps: Get the crash_dumps field of the given VDI. */
func (client *XenClient) VDIGetCrashDumps(self string) (result []string, err error) {
	obj, err := client.APICall("VDI.get_crash_dumps", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetVBDs: Get the VBDs field of the given VDI. */
func (client *XenClient) VDIGetVBDs(self string) (result []string, err error) {
	obj, err := client.APICall("VDI.get_VBDs", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetSR: Get the SR field of the given VDI. */
func (client *XenClient) VDIGetSR(self string) (result string, err error) {
	obj, err := client.APICall("VDI.get_SR", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetCurrentOperations: Get the current_operations field of the given VDI. */
func (client *XenClient) VDIGetCurrentOperations(self string) (result map[string]VdiOperations, err error) {
	obj, err := client.APICall("VDI.get_current_operations", self)
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]VdiOperations{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToVdiOperations(obj.String())
		result[key.String()] = mapObj
	}

	return
}

/* GetAllowedOperations: Get the allowed_operations field of the given VDI. */
func (client *XenClient) VDIGetAllowedOperations(self string) (result []VdiOperations, err error) {
	obj, err := client.APICall("VDI.get_allowed_operations", self)
	if err != nil {
		return
	}

	result = make([]VdiOperations, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = ToVdiOperations(value.(string))
	}

	return
}

/* GetNameDescription: Get the name/description field of the given VDI. */
func (client *XenClient) VDIGetNameDescription(self string) (result string, err error) {
	obj, err := client.APICall("VDI.get_name_description", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetNameLabel: Get the name/label field of the given VDI. */
func (client *XenClient) VDIGetNameLabel(self string) (result string, err error) {
	obj, err := client.APICall("VDI.get_name_label", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given VDI. */
func (client *XenClient) VDIGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("VDI.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByNameLabel: Get all the VDI instances with the given label. */
func (client *XenClient) VDIGetByNameLabel(label string) (result []string, err error) {
	obj, err := client.APICall("VDI.get_by_name_label", label)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* Destroy: Destroy the specified VDI instance. */
func (client *XenClient) VDIDestroy(self string) (err error) {
	_, err = client.APICall("VDI.destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Create: Create a new VDI instance, and return its handle.
The constructor args are: name_label, name_description, SR*, virtual_size*, type*, sharable*, read_only*, other_config*, xenstore_data, sm_config, tags (* = non-optional). */
func (client *XenClient) VDICreate(args VDI) (result string, err error) {
	obj, err := client.APICall("VDI.create", FromVDIToXml(&args))
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the VDI instance with the specified UUID. */
func (client *XenClient) VDIGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("VDI.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given VDI. */
func (client *XenClient) VDIGetRecord(self string) (result VDI, err error) {
	obj, err := client.APICall("VDI.get_record", self)
	if err != nil {
		return
	}

	result = *ToVDI(obj.(interface{}))

	return
}

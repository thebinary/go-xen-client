// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w PVSCacheStorage.gen.go

package go_xen_client

import (
	"reflect"
	"strconv"

	"github.com/nilshell/xmlrpc"
)

//PVSCacheStorage: Describes the storage that is available to a PVS site for caching purposes
type PVSCacheStorage struct {
	Uuid string // Unique identifier/object reference
	Host string // The host on which this object defines PVS cache storage
	SR   string // SR providing storage for the PVS cache
	Site string // The PVS_site for which this object defines the storage
	Size int    // The size of the cache VDI (in bytes)
	VDI  string // The VDI used for caching

}

func FromPVSCacheStorageToXml(PVS_cache_storage *PVSCacheStorage) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] = PVS_cache_storage.Uuid

	result["host"] = PVS_cache_storage.Host

	result["SR"] = PVS_cache_storage.SR

	result["site"] = PVS_cache_storage.Site

	result["size"] = strconv.Itoa(PVS_cache_storage.Size)

	result["VDI"] = PVS_cache_storage.VDI

	return result
}

func ToPVSCacheStorage(obj interface{}) (resultObj *PVSCacheStorage) {

	objValue := reflect.ValueOf(obj)
	resultObj = &PVSCacheStorage{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.Uuid = v
			}
		case "host":
			if v, ok := keyValue.(string); ok {
				resultObj.Host = v
			}
		case "SR":
			if v, ok := keyValue.(string); ok {
				resultObj.SR = v
			}
		case "site":
			if v, ok := keyValue.(string); ok {
				resultObj.Site = v
			}
		case "size":
			if v, ok := keyValue.(int); ok {
				resultObj.Size = v
			}
		case "VDI":
			if v, ok := keyValue.(string); ok {
				resultObj.VDI = v
			}
		}
	}

	return resultObj
}

/* GetAllRecords: Return a map of PVS_cache_storage references to PVS_cache_storage records for all PVS_cache_storages known to the system. */
func (client *XenClient) PVSCacheStorageGetAllRecords() (result map[string]PVSCacheStorage, err error) {
	obj, err := client.APICall("PVS_cache_storage.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]PVSCacheStorage{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToPVSCacheStorage(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the PVS_cache_storages known to the system. */
func (client *XenClient) PVSCacheStorageGetAll() (result []string, err error) {
	obj, err := client.APICall("PVS_cache_storage.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetVDI: Get the VDI field of the given PVS_cache_storage. */
func (client *XenClient) PVSCacheStorageGetVDI(self string) (result string, err error) {
	obj, err := client.APICall("PVS_cache_storage.get_VDI", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetSize: Get the size field of the given PVS_cache_storage. */
func (client *XenClient) PVSCacheStorageGetSize(self string) (result int, err error) {
	obj, err := client.APICall("PVS_cache_storage.get_size", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetSite: Get the site field of the given PVS_cache_storage. */
func (client *XenClient) PVSCacheStorageGetSite(self string) (result string, err error) {
	obj, err := client.APICall("PVS_cache_storage.get_site", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetSR: Get the SR field of the given PVS_cache_storage. */
func (client *XenClient) PVSCacheStorageGetSR(self string) (result string, err error) {
	obj, err := client.APICall("PVS_cache_storage.get_SR", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetHost: Get the host field of the given PVS_cache_storage. */
func (client *XenClient) PVSCacheStorageGetHost(self string) (result string, err error) {
	obj, err := client.APICall("PVS_cache_storage.get_host", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given PVS_cache_storage. */
func (client *XenClient) PVSCacheStorageGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("PVS_cache_storage.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* Destroy: Destroy the specified PVS_cache_storage instance. */
func (client *XenClient) PVSCacheStorageDestroy(self string) (err error) {
	_, err = client.APICall("PVS_cache_storage.destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Create: Create a new PVS_cache_storage instance, and return its handle.
The constructor args are: host, SR, site, size (* = non-optional). */
func (client *XenClient) PVSCacheStorageCreate(args PVSCacheStorage) (result string, err error) {
	obj, err := client.APICall("PVS_cache_storage.create", FromPVSCacheStorageToXml(&args))
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the PVS_cache_storage instance with the specified UUID. */
func (client *XenClient) PVSCacheStorageGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("PVS_cache_storage.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given PVS_cache_storage. */
func (client *XenClient) PVSCacheStorageGetRecord(self string) (result PVSCacheStorage, err error) {
	obj, err := client.APICall("PVS_cache_storage.get_record", self)
	if err != nil {
		return
	}

	result = *ToPVSCacheStorage(obj.(interface{}))

	return
}

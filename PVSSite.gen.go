// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w PVSSite.gen.go

package go_xen_client

import (
	"reflect"

	"github.com/nilshell/xmlrpc"
)

//PVSSite: machines serving blocks of data for provisioning VMs
type PVSSite struct {
	Uuid            string   // Unique identifier/object reference
	NameLabel       string   // a human-readable name
	NameDescription string   // a notes field containing human-readable description
	PVSUuid         string   // Unique identifier of the PVS site, as configured in PVS
	CacheStorage    []string // The SR used by PVS proxy for the cache
	Servers         []string // The set of PVS servers in the site
	Proxies         []string // The set of proxies associated with the site

}

func FromPVSSiteToXml(PVS_site *PVSSite) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] = PVS_site.Uuid

	result["name_label"] = PVS_site.NameLabel

	result["name_description"] = PVS_site.NameDescription

	result["PVS_uuid"] = PVS_site.PVSUuid

	result["cache_storage"] = PVS_site.CacheStorage

	result["servers"] = PVS_site.Servers

	result["proxies"] = PVS_site.Proxies

	return result
}

func ToPVSSite(obj interface{}) (resultObj *PVSSite) {

	objValue := reflect.ValueOf(obj)
	resultObj = &PVSSite{}

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
		case "PVS_uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.PVSUuid = v
			}
		case "cache_storage":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.CacheStorage = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.CacheStorage[i] = v
					}
				}
			}
		case "servers":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.Servers = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.Servers[i] = v
					}
				}
			}
		case "proxies":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.Proxies = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.Proxies[i] = v
					}
				}
			}
		}
	}

	return resultObj
}

/* GetAllRecords: Return a map of PVS_site references to PVS_site records for all PVS_sites known to the system. */
func (client *XenClient) PVSSiteGetAllRecords() (result map[string]PVSSite, err error) {
	obj, err := client.APICall("PVS_site.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]PVSSite{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToPVSSite(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the PVS_sites known to the system. */
func (client *XenClient) PVSSiteGetAll() (result []string, err error) {
	obj, err := client.APICall("PVS_site.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* SetPVSUuid: Update the PVS UUID of the PVS site */
func (client *XenClient) PVSSiteSetPVSUuid(self string, value string) (err error) {
	_, err = client.APICall("PVS_site.set_PVS_uuid", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Forget: Remove a site's meta data */
func (client *XenClient) PVSSiteForget(self string) (err error) {
	_, err = client.APICall("PVS_site.forget", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Introduce: Introduce new PVS site */
func (client *XenClient) PVSSiteIntroduce(name_label string, name_description string, PVS_uuid string) (result string, err error) {
	obj, err := client.APICall("PVS_site.introduce", name_label, name_description, PVS_uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* SetNameDescription: Set the name/description field of the given PVS_site. */
func (client *XenClient) PVSSiteSetNameDescription(self string, value string) (err error) {
	_, err = client.APICall("PVS_site.set_name_description", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetNameLabel: Set the name/label field of the given PVS_site. */
func (client *XenClient) PVSSiteSetNameLabel(self string, value string) (err error) {
	_, err = client.APICall("PVS_site.set_name_label", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetProxies: Get the proxies field of the given PVS_site. */
func (client *XenClient) PVSSiteGetProxies(self string) (result []string, err error) {
	obj, err := client.APICall("PVS_site.get_proxies", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetServers: Get the servers field of the given PVS_site. */
func (client *XenClient) PVSSiteGetServers(self string) (result []string, err error) {
	obj, err := client.APICall("PVS_site.get_servers", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetCacheStorage: Get the cache_storage field of the given PVS_site. */
func (client *XenClient) PVSSiteGetCacheStorage(self string) (result []string, err error) {
	obj, err := client.APICall("PVS_site.get_cache_storage", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetPVSUuid: Get the PVS_uuid field of the given PVS_site. */
func (client *XenClient) PVSSiteGetPVSUuid(self string) (result string, err error) {
	obj, err := client.APICall("PVS_site.get_PVS_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetNameDescription: Get the name/description field of the given PVS_site. */
func (client *XenClient) PVSSiteGetNameDescription(self string) (result string, err error) {
	obj, err := client.APICall("PVS_site.get_name_description", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetNameLabel: Get the name/label field of the given PVS_site. */
func (client *XenClient) PVSSiteGetNameLabel(self string) (result string, err error) {
	obj, err := client.APICall("PVS_site.get_name_label", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given PVS_site. */
func (client *XenClient) PVSSiteGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("PVS_site.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByNameLabel: Get all the PVS_site instances with the given label. */
func (client *XenClient) PVSSiteGetByNameLabel(label string) (result []string, err error) {
	obj, err := client.APICall("PVS_site.get_by_name_label", label)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetByUuid: Get a reference to the PVS_site instance with the specified UUID. */
func (client *XenClient) PVSSiteGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("PVS_site.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given PVS_site. */
func (client *XenClient) PVSSiteGetRecord(self string) (result PVSSite, err error) {
	obj, err := client.APICall("PVS_site.get_record", self)
	if err != nil {
		return
	}

	result = *ToPVSSite(obj.(interface{}))

	return
}

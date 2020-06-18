// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w PVSProxy.gen.go

package go_xen_client

import (
	"reflect"

	"github.com/nilshell/xmlrpc"
)

//PVSProxy: a proxy connects a VM/VIF with a PVS site
type PVSProxy struct {
	Uuid              string         // Unique identifier/object reference
	Site              string         // PVS site this proxy is part of
	VIF               string         // VIF of the VM using the proxy
	CurrentlyAttached bool           // true = VM is currently proxied
	Status            PvsProxyStatus // The run-time status of the proxy

}

func FromPVSProxyToXml(PVS_proxy *PVSProxy) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] = PVS_proxy.Uuid

	result["site"] = PVS_proxy.Site

	result["VIF"] = PVS_proxy.VIF

	result["currently_attached"] = PVS_proxy.CurrentlyAttached

	result["status"] = PVS_proxy.Status.String()

	return result
}

func ToPVSProxy(obj interface{}) (resultObj *PVSProxy) {

	objValue := reflect.ValueOf(obj)
	resultObj = &PVSProxy{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.Uuid = v
			}
		case "site":
			if v, ok := keyValue.(string); ok {
				resultObj.Site = v
			}
		case "VIF":
			if v, ok := keyValue.(string); ok {
				resultObj.VIF = v
			}
		case "currently_attached":
			if v, ok := keyValue.(bool); ok {
				resultObj.CurrentlyAttached = v
			}
		case "status":
			if v, ok := keyValue.(PvsProxyStatus); ok {
				resultObj.Status = v
			}
		}
	}

	return resultObj
}

/* GetAllRecords: Return a map of PVS_proxy references to PVS_proxy records for all PVS_proxys known to the system. */
func (client *XenClient) PVSProxyGetAllRecords() (result map[string]PVSProxy, err error) {
	obj, err := client.APICall("PVS_proxy.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]PVSProxy{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToPVSProxy(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the PVS_proxys known to the system. */
func (client *XenClient) PVSProxyGetAll() (result []string, err error) {
	obj, err := client.APICall("PVS_proxy.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* Destroy: remove (or switch off) a PVS proxy for this VM */
func (client *XenClient) PVSProxyDestroy(self string) (err error) {
	_, err = client.APICall("PVS_proxy.destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Create: Configure a VM/VIF to use a PVS proxy */
func (client *XenClient) PVSProxyCreate(site string, VIF string) (result string, err error) {
	obj, err := client.APICall("PVS_proxy.create", site, VIF)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetStatus: Get the status field of the given PVS_proxy. */
func (client *XenClient) PVSProxyGetStatus(self string) (result PvsProxyStatus, err error) {
	obj, err := client.APICall("PVS_proxy.get_status", self)
	if err != nil {
		return
	}

	result = ToPvsProxyStatus(obj.(string))

	return
}

/* GetCurrentlyAttached: Get the currently_attached field of the given PVS_proxy. */
func (client *XenClient) PVSProxyGetCurrentlyAttached(self string) (result bool, err error) {
	obj, err := client.APICall("PVS_proxy.get_currently_attached", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetVIF: Get the VIF field of the given PVS_proxy. */
func (client *XenClient) PVSProxyGetVIF(self string) (result string, err error) {
	obj, err := client.APICall("PVS_proxy.get_VIF", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetSite: Get the site field of the given PVS_proxy. */
func (client *XenClient) PVSProxyGetSite(self string) (result string, err error) {
	obj, err := client.APICall("PVS_proxy.get_site", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given PVS_proxy. */
func (client *XenClient) PVSProxyGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("PVS_proxy.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the PVS_proxy instance with the specified UUID. */
func (client *XenClient) PVSProxyGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("PVS_proxy.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given PVS_proxy. */
func (client *XenClient) PVSProxyGetRecord(self string) (result PVSProxy, err error) {
	obj, err := client.APICall("PVS_proxy.get_record", self)
	if err != nil {
		return
	}

	result = *ToPVSProxy(obj.(interface{}))

	return
}

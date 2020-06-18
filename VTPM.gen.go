// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w VTPM.gen.go

package go_xen_client

import (
	"reflect"

	"github.com/nilshell/xmlrpc"
)

//VTPM: A virtual TPM device
type VTPM struct {
	Uuid    string // Unique identifier/object reference
	VM      string // the virtual machine
	Backend string // the domain where the backend is located

}

func FromVTPMToXml(VTPM *VTPM) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] = VTPM.Uuid

	result["VM"] = VTPM.VM

	result["backend"] = VTPM.Backend

	return result
}

func ToVTPM(obj interface{}) (resultObj *VTPM) {

	objValue := reflect.ValueOf(obj)
	resultObj = &VTPM{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.Uuid = v
			}
		case "VM":
			if v, ok := keyValue.(string); ok {
				resultObj.VM = v
			}
		case "backend":
			if v, ok := keyValue.(string); ok {
				resultObj.Backend = v
			}
		}
	}

	return resultObj
}

/* GetBackend: Get the backend field of the given VTPM. */
func (client *XenClient) VTPMGetBackend(self string) (result string, err error) {
	obj, err := client.APICall("VTPM.get_backend", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetVM: Get the VM field of the given VTPM. */
func (client *XenClient) VTPMGetVM(self string) (result string, err error) {
	obj, err := client.APICall("VTPM.get_VM", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given VTPM. */
func (client *XenClient) VTPMGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("VTPM.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* Destroy: Destroy the specified VTPM instance. */
func (client *XenClient) VTPMDestroy(self string) (err error) {
	_, err = client.APICall("VTPM.destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Create: Create a new VTPM instance, and return its handle.
The constructor args are: VM*, backend* (* = non-optional). */
func (client *XenClient) VTPMCreate(args VTPM) (result string, err error) {
	obj, err := client.APICall("VTPM.create", FromVTPMToXml(&args))
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the VTPM instance with the specified UUID. */
func (client *XenClient) VTPMGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("VTPM.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given VTPM. */
func (client *XenClient) VTPMGetRecord(self string) (result VTPM, err error) {
	obj, err := client.APICall("VTPM.get_record", self)
	if err != nil {
		return
	}

	result = *ToVTPM(obj.(interface{}))

	return
}

// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w LVHD.gen.go

package go_xen_client

import (
	"reflect"

	"github.com/nilshell/xmlrpc"
)

//LVHD: LVHD SR specific operations
type LVHD struct {
	Uuid string // Unique identifier/object reference

}

func FromLVHDToXml(LVHD *LVHD) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] =

		LVHD.Uuid

	return result
}

func ToLVHD(obj interface{}) (resultObj *LVHD) {

	objValue := reflect.ValueOf(obj)
	resultObj = &LVHD{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.Uuid = v
			}
		}
	}

	return resultObj
}

/* EnableThinProvisioning: Upgrades an LVHD SR to enable thin-provisioning. Future VDIs created in this SR will be thinly-provisioned, although existing VDIs will be left alone. Note that the SR must be attached to the SRmaster for upgrade to work. */
func (client *XenClient) LVHDEnableThinProvisioning(host string, SR string, initial_allocation int, allocation_quantum int) (result string, err error) {
	obj, err := client.APICall("LVHD.enable_thin_provisioning", host, SR, initial_allocation, allocation_quantum)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given LVHD. */
func (client *XenClient) LVHDGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("LVHD.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the LVHD instance with the specified UUID. */
func (client *XenClient) LVHDGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("LVHD.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given LVHD. */
func (client *XenClient) LVHDGetRecord(self string) (result LVHD, err error) {
	obj, err := client.APICall("LVHD.get_record", self)
	if err != nil {
		return
	}

	result = *ToLVHD(obj.(interface{}))

	return
}

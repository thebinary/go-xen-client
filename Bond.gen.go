// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w Bond.gen.go

package go_xen_client

import (
	"reflect"
	"strconv"

	"github.com/nilshell/xmlrpc"
)

//Bond:
type Bond struct {
	Uuid         string            // Unique identifier/object reference
	Master       string            // The bonded interface
	Slaves       []string          // The interfaces which are part of this bond
	OtherConfig  map[string]string // additional configuration
	PrimarySlave string            // The PIF of which the IP configuration and MAC were copied to the bond, and which will receive all configuration/VLANs/VIFs on the bond if the bond is destroyed
	Mode         BondMode          // The algorithm used to distribute traffic among the bonded NICs
	Properties   map[string]string // Additional configuration properties specific to the bond mode.
	LinksUp      int               // Number of links up in this bond

}

func FromBondToXml(Bond *Bond) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] =

		Bond.Uuid

	result["master"] =

		Bond.Master

	result["slaves"] =

		Bond.Slaves

	other_config := make(xmlrpc.Struct)
	for key, value := range Bond.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	result["primary_slave"] =

		Bond.PrimarySlave

	result["mode"] =

		Bond.Mode.String()

	properties := make(xmlrpc.Struct)
	for key, value := range Bond.Properties {
		properties[key] = value
	}
	result["properties"] = properties

	result["links_up"] =

		strconv.Itoa(Bond.LinksUp)

	return result
}

func ToBond(obj interface{}) (resultObj *Bond) {

	objValue := reflect.ValueOf(obj)
	resultObj = &Bond{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.Uuid = v
			}
		case "master":
			if v, ok := keyValue.(string); ok {
				resultObj.Master = v
			}
		case "slaves":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.Slaves = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.Slaves[i] = v
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
		case "primary_slave":
			if v, ok := keyValue.(string); ok {
				resultObj.PrimarySlave = v
			}
		case "mode":
			if v, ok := keyValue.(BondMode); ok {
				resultObj.Mode = v
			}
		case "properties":

			resultObj.Properties = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.Properties[mapKeyName] = v
				} else {
					resultObj.Properties[mapKeyName] = ""
				}
			}
		case "links_up":
			if v, ok := keyValue.(int); ok {
				resultObj.LinksUp = v
			}
		}
	}

	return resultObj
}

/* GetAllRecords: Return a map of Bond references to Bond records for all Bonds known to the system. */
func (client *XenClient) BondGetAllRecords() (result map[string]Bond, err error) {
	obj, err := client.APICall("Bond.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]Bond{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToBond(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the Bonds known to the system. */
func (client *XenClient) BondGetAll() (result []string, err error) {
	obj, err := client.APICall("Bond.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* SetProperty: Set the value of a property of the bond */
func (client *XenClient) BondSetProperty(self string, name string, value string) (err error) {
	_, err = client.APICall("Bond.set_property", self, name, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetMode: Change the bond mode */
func (client *XenClient) BondSetMode(self string, value BondMode) (err error) {
	_, err = client.APICall("Bond.set_mode", self, value.String())
	if err != nil {
		return
	}
	// no return result
	return
}

/* Destroy: Destroy an interface bond */
func (client *XenClient) BondDestroy(self string) (err error) {
	_, err = client.APICall("Bond.destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Create: Create an interface bond */
func (client *XenClient) BondCreate(network string, members []string, MAC string, mode BondMode, properties map[string]string) (result string, err error) {
	obj, err := client.APICall("Bond.create", network, members, MAC, mode.String(), properties)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given Bond.  If the key is not in that Map, then do nothing. */
func (client *XenClient) BondRemoveFromOtherConfig(self string, key string) (err error) {
	_, err = client.APICall("Bond.remove_from_other_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToOtherConfig: Add the given key-value pair to the other_config field of the given Bond. */
func (client *XenClient) BondAddToOtherConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("Bond.add_to_other_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOtherConfig: Set the other_config field of the given Bond. */
func (client *XenClient) BondSetOtherConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("Bond.set_other_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetLinksUp: Get the links_up field of the given Bond. */
func (client *XenClient) BondGetLinksUp(self string) (result int, err error) {
	obj, err := client.APICall("Bond.get_links_up", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetProperties: Get the properties field of the given Bond. */
func (client *XenClient) BondGetProperties(self string) (result map[string]string, err error) {
	obj, err := client.APICall("Bond.get_properties", self)
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

/* GetMode: Get the mode field of the given Bond. */
func (client *XenClient) BondGetMode(self string) (result BondMode, err error) {
	obj, err := client.APICall("Bond.get_mode", self)
	if err != nil {
		return
	}

	result = ToBondMode(obj.(string))

	return
}

/* GetPrimarySlave: Get the primary_slave field of the given Bond. */
func (client *XenClient) BondGetPrimarySlave(self string) (result string, err error) {
	obj, err := client.APICall("Bond.get_primary_slave", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetOtherConfig: Get the other_config field of the given Bond. */
func (client *XenClient) BondGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("Bond.get_other_config", self)
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

/* GetSlaves: Get the slaves field of the given Bond. */
func (client *XenClient) BondGetSlaves(self string) (result []string, err error) {
	obj, err := client.APICall("Bond.get_slaves", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetMaster: Get the master field of the given Bond. */
func (client *XenClient) BondGetMaster(self string) (result string, err error) {
	obj, err := client.APICall("Bond.get_master", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given Bond. */
func (client *XenClient) BondGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("Bond.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the Bond instance with the specified UUID. */
func (client *XenClient) BondGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("Bond.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given Bond. */
func (client *XenClient) BondGetRecord(self string) (result Bond, err error) {
	obj, err := client.APICall("Bond.get_record", self)
	if err != nil {
		return
	}

	result = *ToBond(obj.(interface{}))

	return
}

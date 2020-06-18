// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w Console.gen.go

package go_xen_client

import (
	"reflect"

	"github.com/nilshell/xmlrpc"
)

//Console: A console
type Console struct {
	Uuid        string            // Unique identifier/object reference
	Protocol    ConsoleProtocol   // the protocol used by this console
	Location    string            // URI for the console service
	VM          string            // VM to which this console is attached
	OtherConfig map[string]string // additional configuration

}

func FromConsoleToXml(console *Console) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] =

		console.Uuid

	result["protocol"] =

		console.Protocol.String()

	result["location"] =

		console.Location

	result["VM"] =

		console.VM

	other_config := make(xmlrpc.Struct)
	for key, value := range console.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	return result
}

func ToConsole(obj interface{}) (resultObj *Console) {

	objValue := reflect.ValueOf(obj)
	resultObj = &Console{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.Uuid = v
			}
		case "protocol":
			if v, ok := keyValue.(ConsoleProtocol); ok {
				resultObj.Protocol = v
			}
		case "location":
			if v, ok := keyValue.(string); ok {
				resultObj.Location = v
			}
		case "VM":
			if v, ok := keyValue.(string); ok {
				resultObj.VM = v
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

/* GetAllRecords: Return a map of console references to console records for all consoles known to the system. */
func (client *XenClient) ConsoleGetAllRecords() (result map[string]Console, err error) {
	obj, err := client.APICall("console.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]Console{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToConsole(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the consoles known to the system. */
func (client *XenClient) ConsoleGetAll() (result []string, err error) {
	obj, err := client.APICall("console.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given console.  If the key is not in that Map, then do nothing. */
func (client *XenClient) ConsoleRemoveFromOtherConfig(self string, key string) (err error) {
	_, err = client.APICall("console.remove_from_other_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToOtherConfig: Add the given key-value pair to the other_config field of the given console. */
func (client *XenClient) ConsoleAddToOtherConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("console.add_to_other_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOtherConfig: Set the other_config field of the given console. */
func (client *XenClient) ConsoleSetOtherConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("console.set_other_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetOtherConfig: Get the other_config field of the given console. */
func (client *XenClient) ConsoleGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("console.get_other_config", self)
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

/* GetVM: Get the VM field of the given console. */
func (client *XenClient) ConsoleGetVM(self string) (result string, err error) {
	obj, err := client.APICall("console.get_VM", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetLocation: Get the location field of the given console. */
func (client *XenClient) ConsoleGetLocation(self string) (result string, err error) {
	obj, err := client.APICall("console.get_location", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetProtocol: Get the protocol field of the given console. */
func (client *XenClient) ConsoleGetProtocol(self string) (result ConsoleProtocol, err error) {
	obj, err := client.APICall("console.get_protocol", self)
	if err != nil {
		return
	}

	result = ToConsoleProtocol(obj.(string))

	return
}

/* GetUuid: Get the uuid field of the given console. */
func (client *XenClient) ConsoleGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("console.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* Destroy: Destroy the specified console instance. */
func (client *XenClient) ConsoleDestroy(self string) (err error) {
	_, err = client.APICall("console.destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Create: Create a new console instance, and return its handle.
The constructor args are: other_config* (* = non-optional). */
func (client *XenClient) ConsoleCreate(args Console) (result string, err error) {
	obj, err := client.APICall("console.create", FromConsoleToXml(&args))
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the console instance with the specified UUID. */
func (client *XenClient) ConsoleGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("console.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given console. */
func (client *XenClient) ConsoleGetRecord(self string) (result Console, err error) {
	obj, err := client.APICall("console.get_record", self)
	if err != nil {
		return
	}

	result = *ToConsole(obj.(interface{}))

	return
}

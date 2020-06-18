// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w SDNController.gen.go

package go_xen_client

import (
	"reflect"
	"strconv"

	"github.com/nilshell/xmlrpc"
)

//SDNController: Describes the SDN controller that is to connect with the pool
type SDNController struct {
	Uuid     string                // Unique identifier/object reference
	Protocol SdnControllerProtocol // Protocol to connect with SDN controller
	Address  string                // IP address of the controller
	Port     int                   // TCP port of the controller

}

func FromSDNControllerToXml(SDN_controller *SDNController) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] =

		SDN_controller.Uuid

	result["protocol"] =

		SDN_controller.Protocol.String()

	result["address"] =

		SDN_controller.Address

	result["port"] =

		strconv.Itoa(SDN_controller.Port)

	return result
}

func ToSDNController(obj interface{}) (resultObj *SDNController) {

	objValue := reflect.ValueOf(obj)
	resultObj = &SDNController{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.Uuid = v
			}
		case "protocol":
			if v, ok := keyValue.(SdnControllerProtocol); ok {
				resultObj.Protocol = v
			}
		case "address":
			if v, ok := keyValue.(string); ok {
				resultObj.Address = v
			}
		case "port":
			if v, ok := keyValue.(int); ok {
				resultObj.Port = v
			}
		}
	}

	return resultObj
}

/* GetAllRecords: Return a map of SDN_controller references to SDN_controller records for all SDN_controllers known to the system. */
func (client *XenClient) SDNControllerGetAllRecords() (result map[string]SDNController, err error) {
	obj, err := client.APICall("SDN_controller.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]SDNController{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToSDNController(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the SDN_controllers known to the system. */
func (client *XenClient) SDNControllerGetAll() (result []string, err error) {
	obj, err := client.APICall("SDN_controller.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* Forget: Remove the OVS manager of the pool and destroy the db record. */
func (client *XenClient) SDNControllerForget(self string) (err error) {
	_, err = client.APICall("SDN_controller.forget", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Introduce: Introduce an SDN controller to the pool. */
func (client *XenClient) SDNControllerIntroduce(protocol SdnControllerProtocol, address string, port int) (result string, err error) {
	obj, err := client.APICall("SDN_controller.introduce", protocol.String(), address, port)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetPort: Get the port field of the given SDN_controller. */
func (client *XenClient) SDNControllerGetPort(self string) (result int, err error) {
	obj, err := client.APICall("SDN_controller.get_port", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetAddress: Get the address field of the given SDN_controller. */
func (client *XenClient) SDNControllerGetAddress(self string) (result string, err error) {
	obj, err := client.APICall("SDN_controller.get_address", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetProtocol: Get the protocol field of the given SDN_controller. */
func (client *XenClient) SDNControllerGetProtocol(self string) (result SdnControllerProtocol, err error) {
	obj, err := client.APICall("SDN_controller.get_protocol", self)
	if err != nil {
		return
	}

	result = ToSdnControllerProtocol(obj.(string))

	return
}

/* GetUuid: Get the uuid field of the given SDN_controller. */
func (client *XenClient) SDNControllerGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("SDN_controller.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the SDN_controller instance with the specified UUID. */
func (client *XenClient) SDNControllerGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("SDN_controller.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given SDN_controller. */
func (client *XenClient) SDNControllerGetRecord(self string) (result SDNController, err error) {
	obj, err := client.APICall("SDN_controller.get_record", self)
	if err != nil {
		return
	}

	result = *ToSDNController(obj.(interface{}))

	return
}

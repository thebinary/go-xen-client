// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w PUSB.gen.go

package go_xen_client

import (
	"reflect"

	"github.com/nilshell/xmlrpc"
)

//PUSB: A physical USB device
type PUSB struct {
	Uuid               string            // Unique identifier/object reference
	USBGroup           string            // USB group the PUSB is contained in
	Host               string            // Physical machine that owns the USB device
	Path               string            // port path of USB device
	VendorId           string            // vendor id of the USB device
	VendorDesc         string            // vendor description of the USB device
	ProductId          string            // product id of the USB device
	ProductDesc        string            // product description of the USB device
	Serial             string            // serial of the USB device
	Version            string            // USB device version
	Description        string            // USB device description
	PassthroughEnabled bool              // enabled for passthrough
	OtherConfig        map[string]string // additional configuration

}

func FromPUSBToXml(PUSB *PUSB) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] = PUSB.Uuid

	result["USB_group"] = PUSB.USBGroup

	result["host"] = PUSB.Host

	result["path"] = PUSB.Path

	result["vendor_id"] = PUSB.VendorId

	result["vendor_desc"] = PUSB.VendorDesc

	result["product_id"] = PUSB.ProductId

	result["product_desc"] = PUSB.ProductDesc

	result["serial"] = PUSB.Serial

	result["version"] = PUSB.Version

	result["description"] = PUSB.Description

	result["passthrough_enabled"] = PUSB.PassthroughEnabled

	other_config := make(xmlrpc.Struct)
	for key, value := range PUSB.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	return result
}

func ToPUSB(obj interface{}) (resultObj *PUSB) {

	objValue := reflect.ValueOf(obj)
	resultObj = &PUSB{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.Uuid = v
			}
		case "USB_group":
			if v, ok := keyValue.(string); ok {
				resultObj.USBGroup = v
			}
		case "host":
			if v, ok := keyValue.(string); ok {
				resultObj.Host = v
			}
		case "path":
			if v, ok := keyValue.(string); ok {
				resultObj.Path = v
			}
		case "vendor_id":
			if v, ok := keyValue.(string); ok {
				resultObj.VendorId = v
			}
		case "vendor_desc":
			if v, ok := keyValue.(string); ok {
				resultObj.VendorDesc = v
			}
		case "product_id":
			if v, ok := keyValue.(string); ok {
				resultObj.ProductId = v
			}
		case "product_desc":
			if v, ok := keyValue.(string); ok {
				resultObj.ProductDesc = v
			}
		case "serial":
			if v, ok := keyValue.(string); ok {
				resultObj.Serial = v
			}
		case "version":
			if v, ok := keyValue.(string); ok {
				resultObj.Version = v
			}
		case "description":
			if v, ok := keyValue.(string); ok {
				resultObj.Description = v
			}
		case "passthrough_enabled":
			if v, ok := keyValue.(bool); ok {
				resultObj.PassthroughEnabled = v
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

/* GetAllRecords: Return a map of PUSB references to PUSB records for all PUSBs known to the system. */
func (client *XenClient) PUSBGetAllRecords() (result map[string]PUSB, err error) {
	obj, err := client.APICall("PUSB.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]PUSB{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToPUSB(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the PUSBs known to the system. */
func (client *XenClient) PUSBGetAll() (result []string, err error) {
	obj, err := client.APICall("PUSB.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* SetPassthroughEnabled:  */
func (client *XenClient) PUSBSetPassthroughEnabled(self string, value bool) (err error) {
	_, err = client.APICall("PUSB.set_passthrough_enabled", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Scan:  */
func (client *XenClient) PUSBScan(host string) (err error) {
	_, err = client.APICall("PUSB.scan", host)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given PUSB.  If the key is not in that Map, then do nothing. */
func (client *XenClient) PUSBRemoveFromOtherConfig(self string, key string) (err error) {
	_, err = client.APICall("PUSB.remove_from_other_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToOtherConfig: Add the given key-value pair to the other_config field of the given PUSB. */
func (client *XenClient) PUSBAddToOtherConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("PUSB.add_to_other_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOtherConfig: Set the other_config field of the given PUSB. */
func (client *XenClient) PUSBSetOtherConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("PUSB.set_other_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetOtherConfig: Get the other_config field of the given PUSB. */
func (client *XenClient) PUSBGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("PUSB.get_other_config", self)
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

/* GetPassthroughEnabled: Get the passthrough_enabled field of the given PUSB. */
func (client *XenClient) PUSBGetPassthroughEnabled(self string) (result bool, err error) {
	obj, err := client.APICall("PUSB.get_passthrough_enabled", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetDescription: Get the description field of the given PUSB. */
func (client *XenClient) PUSBGetDescription(self string) (result string, err error) {
	obj, err := client.APICall("PUSB.get_description", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetVersion: Get the version field of the given PUSB. */
func (client *XenClient) PUSBGetVersion(self string) (result string, err error) {
	obj, err := client.APICall("PUSB.get_version", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetSerial: Get the serial field of the given PUSB. */
func (client *XenClient) PUSBGetSerial(self string) (result string, err error) {
	obj, err := client.APICall("PUSB.get_serial", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetProductDesc: Get the product_desc field of the given PUSB. */
func (client *XenClient) PUSBGetProductDesc(self string) (result string, err error) {
	obj, err := client.APICall("PUSB.get_product_desc", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetProductId: Get the product_id field of the given PUSB. */
func (client *XenClient) PUSBGetProductId(self string) (result string, err error) {
	obj, err := client.APICall("PUSB.get_product_id", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetVendorDesc: Get the vendor_desc field of the given PUSB. */
func (client *XenClient) PUSBGetVendorDesc(self string) (result string, err error) {
	obj, err := client.APICall("PUSB.get_vendor_desc", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetVendorId: Get the vendor_id field of the given PUSB. */
func (client *XenClient) PUSBGetVendorId(self string) (result string, err error) {
	obj, err := client.APICall("PUSB.get_vendor_id", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetPath: Get the path field of the given PUSB. */
func (client *XenClient) PUSBGetPath(self string) (result string, err error) {
	obj, err := client.APICall("PUSB.get_path", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetHost: Get the host field of the given PUSB. */
func (client *XenClient) PUSBGetHost(self string) (result string, err error) {
	obj, err := client.APICall("PUSB.get_host", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUSBGroup: Get the USB_group field of the given PUSB. */
func (client *XenClient) PUSBGetUSBGroup(self string) (result string, err error) {
	obj, err := client.APICall("PUSB.get_USB_group", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given PUSB. */
func (client *XenClient) PUSBGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("PUSB.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the PUSB instance with the specified UUID. */
func (client *XenClient) PUSBGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("PUSB.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given PUSB. */
func (client *XenClient) PUSBGetRecord(self string) (result PUSB, err error) {
	obj, err := client.APICall("PUSB.get_record", self)
	if err != nil {
		return
	}

	result = *ToPUSB(obj.(interface{}))

	return
}

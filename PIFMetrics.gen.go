// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w PIFMetrics.gen.go

package go_xen_client

import (
	"reflect"
	"strconv"
	"time"

	"github.com/nilshell/xmlrpc"
)

//PIFMetrics: The metrics associated with a physical network interface
type PIFMetrics struct {
	Uuid        string            // Unique identifier/object reference
	IoReadKbs   float32           // Read bandwidth (KiB/s)
	IoWriteKbs  float32           // Write bandwidth (KiB/s)
	Carrier     bool              // Report if the PIF got a carrier or not
	VendorId    string            // Report vendor ID
	VendorName  string            // Report vendor name
	DeviceId    string            // Report device ID
	DeviceName  string            // Report device name
	Speed       int               // Speed of the link (if available)
	Duplex      bool              // Full duplex capability of the link (if available)
	PciBusPath  string            // PCI bus path of the pif (if available)
	LastUpdated time.Time         // Time at which this information was last updated
	OtherConfig map[string]string // additional configuration

}

func FromPIFMetricsToXml(PIF_metrics *PIFMetrics) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] = PIF_metrics.Uuid

	result["io_read_kbs"] = PIF_metrics.IoReadKbs

	result["io_write_kbs"] = PIF_metrics.IoWriteKbs

	result["carrier"] = PIF_metrics.Carrier

	result["vendor_id"] = PIF_metrics.VendorId

	result["vendor_name"] = PIF_metrics.VendorName

	result["device_id"] = PIF_metrics.DeviceId

	result["device_name"] = PIF_metrics.DeviceName

	result["speed"] = strconv.Itoa(PIF_metrics.Speed)

	result["duplex"] = PIF_metrics.Duplex

	result["pci_bus_path"] = PIF_metrics.PciBusPath

	result["last_updated"] = PIF_metrics.LastUpdated

	other_config := make(xmlrpc.Struct)
	for key, value := range PIF_metrics.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	return result
}

func ToPIFMetrics(obj interface{}) (resultObj *PIFMetrics) {

	objValue := reflect.ValueOf(obj)
	resultObj = &PIFMetrics{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.Uuid = v
			}
		case "io_read_kbs":
			if v, ok := keyValue.(float32); ok {
				resultObj.IoReadKbs = v
			}
		case "io_write_kbs":
			if v, ok := keyValue.(float32); ok {
				resultObj.IoWriteKbs = v
			}
		case "carrier":
			if v, ok := keyValue.(bool); ok {
				resultObj.Carrier = v
			}
		case "vendor_id":
			if v, ok := keyValue.(string); ok {
				resultObj.VendorId = v
			}
		case "vendor_name":
			if v, ok := keyValue.(string); ok {
				resultObj.VendorName = v
			}
		case "device_id":
			if v, ok := keyValue.(string); ok {
				resultObj.DeviceId = v
			}
		case "device_name":
			if v, ok := keyValue.(string); ok {
				resultObj.DeviceName = v
			}
		case "speed":
			if v, ok := keyValue.(int); ok {
				resultObj.Speed = v
			}
		case "duplex":
			if v, ok := keyValue.(bool); ok {
				resultObj.Duplex = v
			}
		case "pci_bus_path":
			if v, ok := keyValue.(string); ok {
				resultObj.PciBusPath = v
			}
		case "last_updated":
			if v, ok := keyValue.(time.Time); ok {
				resultObj.LastUpdated = v
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

/* GetAllRecords: Return a map of PIF_metrics references to PIF_metrics records for all PIF_metrics instances known to the system. */
func (client *XenClient) PIFMetricsGetAllRecords() (result map[string]PIFMetrics, err error) {
	obj, err := client.APICall("PIF_metrics.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]PIFMetrics{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToPIFMetrics(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the PIF_metrics instances known to the system. */
func (client *XenClient) PIFMetricsGetAll() (result []string, err error) {
	obj, err := client.APICall("PIF_metrics.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given PIF_metrics.  If the key is not in that Map, then do nothing. */
func (client *XenClient) PIFMetricsRemoveFromOtherConfig(self string, key string) (err error) {
	_, err = client.APICall("PIF_metrics.remove_from_other_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToOtherConfig: Add the given key-value pair to the other_config field of the given PIF_metrics. */
func (client *XenClient) PIFMetricsAddToOtherConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("PIF_metrics.add_to_other_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOtherConfig: Set the other_config field of the given PIF_metrics. */
func (client *XenClient) PIFMetricsSetOtherConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("PIF_metrics.set_other_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetOtherConfig: Get the other_config field of the given PIF_metrics. */
func (client *XenClient) PIFMetricsGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("PIF_metrics.get_other_config", self)
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

/* GetLastUpdated: Get the last_updated field of the given PIF_metrics. */
func (client *XenClient) PIFMetricsGetLastUpdated(self string) (result time.Time, err error) {
	obj, err := client.APICall("PIF_metrics.get_last_updated", self)
	if err != nil {
		return
	}
	result = obj.(time.Time)
	return
}

/* GetPciBusPath: Get the pci_bus_path field of the given PIF_metrics. */
func (client *XenClient) PIFMetricsGetPciBusPath(self string) (result string, err error) {
	obj, err := client.APICall("PIF_metrics.get_pci_bus_path", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetDuplex: Get the duplex field of the given PIF_metrics. */
func (client *XenClient) PIFMetricsGetDuplex(self string) (result bool, err error) {
	obj, err := client.APICall("PIF_metrics.get_duplex", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetSpeed: Get the speed field of the given PIF_metrics. */
func (client *XenClient) PIFMetricsGetSpeed(self string) (result int, err error) {
	obj, err := client.APICall("PIF_metrics.get_speed", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetDeviceName: Get the device_name field of the given PIF_metrics. */
func (client *XenClient) PIFMetricsGetDeviceName(self string) (result string, err error) {
	obj, err := client.APICall("PIF_metrics.get_device_name", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetDeviceId: Get the device_id field of the given PIF_metrics. */
func (client *XenClient) PIFMetricsGetDeviceId(self string) (result string, err error) {
	obj, err := client.APICall("PIF_metrics.get_device_id", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetVendorName: Get the vendor_name field of the given PIF_metrics. */
func (client *XenClient) PIFMetricsGetVendorName(self string) (result string, err error) {
	obj, err := client.APICall("PIF_metrics.get_vendor_name", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetVendorId: Get the vendor_id field of the given PIF_metrics. */
func (client *XenClient) PIFMetricsGetVendorId(self string) (result string, err error) {
	obj, err := client.APICall("PIF_metrics.get_vendor_id", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetCarrier: Get the carrier field of the given PIF_metrics. */
func (client *XenClient) PIFMetricsGetCarrier(self string) (result bool, err error) {
	obj, err := client.APICall("PIF_metrics.get_carrier", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetIoWriteKbs: Get the io/write_kbs field of the given PIF_metrics. */
func (client *XenClient) PIFMetricsGetIoWriteKbs(self string) (result float32, err error) {
	obj, err := client.APICall("PIF_metrics.get_io_write_kbs", self)
	if err != nil {
		return
	}
	result = obj.(float32)
	return
}

/* GetIoReadKbs: Get the io/read_kbs field of the given PIF_metrics. */
func (client *XenClient) PIFMetricsGetIoReadKbs(self string) (result float32, err error) {
	obj, err := client.APICall("PIF_metrics.get_io_read_kbs", self)
	if err != nil {
		return
	}
	result = obj.(float32)
	return
}

/* GetUuid: Get the uuid field of the given PIF_metrics. */
func (client *XenClient) PIFMetricsGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("PIF_metrics.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the PIF_metrics instance with the specified UUID. */
func (client *XenClient) PIFMetricsGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("PIF_metrics.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given PIF_metrics. */
func (client *XenClient) PIFMetricsGetRecord(self string) (result PIFMetrics, err error) {
	obj, err := client.APICall("PIF_metrics.get_record", self)
	if err != nil {
		return
	}

	result = *ToPIFMetrics(obj.(interface{}))

	return
}

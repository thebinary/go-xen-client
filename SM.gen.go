// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w SM.gen.go

package go_xen_client

import (
	"reflect"

	"github.com/nilshell/xmlrpc"
)

//SM: A storage manager plugin
type SM struct {
	Uuid                 string            // Unique identifier/object reference
	NameLabel            string            // a human-readable name
	NameDescription      string            // a notes field containing human-readable description
	Type                 string            // SR.type
	Vendor               string            // Vendor who created this plugin
	Copyright            string            // Entity which owns the copyright of this plugin
	Version              string            // Version of the plugin
	RequiredApiVersion   string            // Minimum SM API version required on the server
	Configuration        map[string]string // names and descriptions of device config keys
	Capabilities         []string          // capabilities of the SM plugin
	Features             map[string]int    // capabilities of the SM plugin, with capability version numbers
	OtherConfig          map[string]string // additional configuration
	DriverFilename       string            // filename of the storage driver
	RequiredClusterStack []string          // The storage plugin requires that one of these cluster stacks is configured and running.

}

func FromSMToXml(SM *SM) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] = SM.Uuid

	result["name_label"] = SM.NameLabel

	result["name_description"] = SM.NameDescription

	result["type"] = SM.Type

	result["vendor"] = SM.Vendor

	result["copyright"] = SM.Copyright

	result["version"] = SM.Version

	result["required_api_version"] = SM.RequiredApiVersion

	configuration := make(xmlrpc.Struct)
	for key, value := range SM.Configuration {
		configuration[key] = value
	}
	result["configuration"] = configuration

	result["capabilities"] = SM.Capabilities

	features := make(xmlrpc.Struct)
	for key, value := range SM.Features {
		features[key] = value
	}
	result["features"] = features

	other_config := make(xmlrpc.Struct)
	for key, value := range SM.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	result["driver_filename"] = SM.DriverFilename

	result["required_cluster_stack"] = SM.RequiredClusterStack

	return result
}

func ToSM(obj interface{}) (resultObj *SM) {

	objValue := reflect.ValueOf(obj)
	resultObj = &SM{}

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
		case "type":
			if v, ok := keyValue.(string); ok {
				resultObj.Type = v
			}
		case "vendor":
			if v, ok := keyValue.(string); ok {
				resultObj.Vendor = v
			}
		case "copyright":
			if v, ok := keyValue.(string); ok {
				resultObj.Copyright = v
			}
		case "version":
			if v, ok := keyValue.(string); ok {
				resultObj.Version = v
			}
		case "required_api_version":
			if v, ok := keyValue.(string); ok {
				resultObj.RequiredApiVersion = v
			}
		case "configuration":

			resultObj.Configuration = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.Configuration[mapKeyName] = v
				} else {
					resultObj.Configuration[mapKeyName] = ""
				}
			}
		case "capabilities":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.Capabilities = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.Capabilities[i] = v
					}
				}
			}
		case "features":

			resultObj.Features = map[string]int{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(int); ok {
					resultObj.Features[mapKeyName] = v
				} else {
					resultObj.Features[mapKeyName] = 0
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
		case "driver_filename":
			if v, ok := keyValue.(string); ok {
				resultObj.DriverFilename = v
			}
		case "required_cluster_stack":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.RequiredClusterStack = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.RequiredClusterStack[i] = v
					}
				}
			}
		}
	}

	return resultObj
}

/* GetAllRecords: Return a map of SM references to SM records for all SMs known to the system. */
func (client *XenClient) SMGetAllRecords() (result map[string]SM, err error) {
	obj, err := client.APICall("SM.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]SM{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToSM(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the SMs known to the system. */
func (client *XenClient) SMGetAll() (result []string, err error) {
	obj, err := client.APICall("SM.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given SM.  If the key is not in that Map, then do nothing. */
func (client *XenClient) SMRemoveFromOtherConfig(self string, key string) (err error) {
	_, err = client.APICall("SM.remove_from_other_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToOtherConfig: Add the given key-value pair to the other_config field of the given SM. */
func (client *XenClient) SMAddToOtherConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("SM.add_to_other_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOtherConfig: Set the other_config field of the given SM. */
func (client *XenClient) SMSetOtherConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("SM.set_other_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetRequiredClusterStack: Get the required_cluster_stack field of the given SM. */
func (client *XenClient) SMGetRequiredClusterStack(self string) (result []string, err error) {
	obj, err := client.APICall("SM.get_required_cluster_stack", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetDriverFilename: Get the driver_filename field of the given SM. */
func (client *XenClient) SMGetDriverFilename(self string) (result string, err error) {
	obj, err := client.APICall("SM.get_driver_filename", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetOtherConfig: Get the other_config field of the given SM. */
func (client *XenClient) SMGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("SM.get_other_config", self)
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

/* GetFeatures: Get the features field of the given SM. */
func (client *XenClient) SMGetFeatures(self string) (result map[string]int, err error) {
	obj, err := client.APICall("SM.get_features", self)
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]int{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		result[key.String()] = int(obj.Int())
	}

	return
}

/* GetCapabilities: Get the capabilities field of the given SM. */
func (client *XenClient) SMGetCapabilities(self string) (result []string, err error) {
	obj, err := client.APICall("SM.get_capabilities", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetConfiguration: Get the configuration field of the given SM. */
func (client *XenClient) SMGetConfiguration(self string) (result map[string]string, err error) {
	obj, err := client.APICall("SM.get_configuration", self)
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

/* GetRequiredApiVersion: Get the required_api_version field of the given SM. */
func (client *XenClient) SMGetRequiredApiVersion(self string) (result string, err error) {
	obj, err := client.APICall("SM.get_required_api_version", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetVersion: Get the version field of the given SM. */
func (client *XenClient) SMGetVersion(self string) (result string, err error) {
	obj, err := client.APICall("SM.get_version", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetCopyright: Get the copyright field of the given SM. */
func (client *XenClient) SMGetCopyright(self string) (result string, err error) {
	obj, err := client.APICall("SM.get_copyright", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetVendor: Get the vendor field of the given SM. */
func (client *XenClient) SMGetVendor(self string) (result string, err error) {
	obj, err := client.APICall("SM.get_vendor", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetType: Get the type field of the given SM. */
func (client *XenClient) SMGetType(self string) (result string, err error) {
	obj, err := client.APICall("SM.get_type", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetNameDescription: Get the name/description field of the given SM. */
func (client *XenClient) SMGetNameDescription(self string) (result string, err error) {
	obj, err := client.APICall("SM.get_name_description", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetNameLabel: Get the name/label field of the given SM. */
func (client *XenClient) SMGetNameLabel(self string) (result string, err error) {
	obj, err := client.APICall("SM.get_name_label", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given SM. */
func (client *XenClient) SMGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("SM.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByNameLabel: Get all the SM instances with the given label. */
func (client *XenClient) SMGetByNameLabel(label string) (result []string, err error) {
	obj, err := client.APICall("SM.get_by_name_label", label)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetByUuid: Get a reference to the SM instance with the specified UUID. */
func (client *XenClient) SMGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("SM.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given SM. */
func (client *XenClient) SMGetRecord(self string) (result SM, err error) {
	obj, err := client.APICall("SM.get_record", self)
	if err != nil {
		return
	}

	result = *ToSM(obj.(interface{}))

	return
}

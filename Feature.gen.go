// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w Feature.gen.go

package go_xen_client

import (
	"reflect"

	"github.com/nilshell/xmlrpc"
)

//Feature: A new piece of functionality
type Feature struct {
	Uuid            string // Unique identifier/object reference
	NameLabel       string // a human-readable name
	NameDescription string // a notes field containing human-readable description
	Enabled         bool   // Indicates whether the feature is enabled
	Experimental    bool   // Indicates whether the feature is experimental (as opposed to stable and fully supported)
	Version         string // The version of this feature
	Host            string // The host where this feature is available

}

func FromFeatureToXml(Feature *Feature) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] = Feature.Uuid

	result["name_label"] = Feature.NameLabel

	result["name_description"] = Feature.NameDescription

	result["enabled"] = Feature.Enabled

	result["experimental"] = Feature.Experimental

	result["version"] = Feature.Version

	result["host"] = Feature.Host

	return result
}

func ToFeature(obj interface{}) (resultObj *Feature) {

	objValue := reflect.ValueOf(obj)
	resultObj = &Feature{}

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
		case "enabled":
			if v, ok := keyValue.(bool); ok {
				resultObj.Enabled = v
			}
		case "experimental":
			if v, ok := keyValue.(bool); ok {
				resultObj.Experimental = v
			}
		case "version":
			if v, ok := keyValue.(string); ok {
				resultObj.Version = v
			}
		case "host":
			if v, ok := keyValue.(string); ok {
				resultObj.Host = v
			}
		}
	}

	return resultObj
}

/* GetAllRecords: Return a map of Feature references to Feature records for all Features known to the system. */
func (client *XenClient) FeatureGetAllRecords() (result map[string]Feature, err error) {
	obj, err := client.APICall("Feature.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]Feature{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToFeature(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the Features known to the system. */
func (client *XenClient) FeatureGetAll() (result []string, err error) {
	obj, err := client.APICall("Feature.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetHost: Get the host field of the given Feature. */
func (client *XenClient) FeatureGetHost(self string) (result string, err error) {
	obj, err := client.APICall("Feature.get_host", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetVersion: Get the version field of the given Feature. */
func (client *XenClient) FeatureGetVersion(self string) (result string, err error) {
	obj, err := client.APICall("Feature.get_version", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetExperimental: Get the experimental field of the given Feature. */
func (client *XenClient) FeatureGetExperimental(self string) (result bool, err error) {
	obj, err := client.APICall("Feature.get_experimental", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetEnabled: Get the enabled field of the given Feature. */
func (client *XenClient) FeatureGetEnabled(self string) (result bool, err error) {
	obj, err := client.APICall("Feature.get_enabled", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetNameDescription: Get the name/description field of the given Feature. */
func (client *XenClient) FeatureGetNameDescription(self string) (result string, err error) {
	obj, err := client.APICall("Feature.get_name_description", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetNameLabel: Get the name/label field of the given Feature. */
func (client *XenClient) FeatureGetNameLabel(self string) (result string, err error) {
	obj, err := client.APICall("Feature.get_name_label", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given Feature. */
func (client *XenClient) FeatureGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("Feature.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByNameLabel: Get all the Feature instances with the given label. */
func (client *XenClient) FeatureGetByNameLabel(label string) (result []string, err error) {
	obj, err := client.APICall("Feature.get_by_name_label", label)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetByUuid: Get a reference to the Feature instance with the specified UUID. */
func (client *XenClient) FeatureGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("Feature.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given Feature. */
func (client *XenClient) FeatureGetRecord(self string) (result Feature, err error) {
	obj, err := client.APICall("Feature.get_record", self)
	if err != nil {
		return
	}

	result = *ToFeature(obj.(interface{}))

	return
}

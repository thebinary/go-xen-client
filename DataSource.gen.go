// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w DataSource.gen.go

package go_xen_client

import (
	"reflect"

	"github.com/nilshell/xmlrpc"
)

//DataSource: Data sources for logging in RRDs
type DataSource struct {
	NameLabel       string  // a human-readable name
	NameDescription string  // a notes field containing human-readable description
	Enabled         bool    // true if the data source is being logged
	Standard        bool    // true if the data source is enabled by default. Non-default data sources cannot be disabled
	Units           string  // the units of the value
	Min             float32 // the minimum value of the data source
	Max             float32 // the maximum value of the data source
	Value           float32 // current value of the data source

}

func FromDataSourceToXml(data_source *DataSource) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["name_label"] = data_source.NameLabel

	result["name_description"] = data_source.NameDescription

	result["enabled"] = data_source.Enabled

	result["standard"] = data_source.Standard

	result["units"] = data_source.Units

	result["min"] = data_source.Min

	result["max"] = data_source.Max

	result["value"] = data_source.Value

	return result
}

func ToDataSource(obj interface{}) (resultObj *DataSource) {

	objValue := reflect.ValueOf(obj)
	resultObj = &DataSource{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
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
		case "standard":
			if v, ok := keyValue.(bool); ok {
				resultObj.Standard = v
			}
		case "units":
			if v, ok := keyValue.(string); ok {
				resultObj.Units = v
			}
		case "min":
			if v, ok := keyValue.(float32); ok {
				resultObj.Min = v
			}
		case "max":
			if v, ok := keyValue.(float32); ok {
				resultObj.Max = v
			}
		case "value":
			if v, ok := keyValue.(float32); ok {
				resultObj.Value = v
			}
		}
	}

	return resultObj
}

// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w SrStat.gen.go

package go_xen_client

import (
	"reflect"
	"strconv"

	"github.com/nilshell/xmlrpc"
)

//SrStat: A set of high-level properties associated with an SR.
type SrStat struct {
	Uuid            interface{} // Uuid that uniquely identifies this SR, if one is available.
	NameLabel       string      // Short, human-readable label for the SR.
	NameDescription string      // Longer, human-readable description of the SR. Descriptions are generally only displayed by clients when the user is examining SRs in detail.
	FreeSpace       int         // Number of bytes free on the backing storage (in bytes)
	TotalSpace      int         // Total physical size of the backing storage (in bytes)
	Clustered       bool        // Indicates whether the SR uses clustered local storage.
	Health          SrHealth    // The health status of the SR.

}

func FromSrStatToXml(sr_stat *SrStat) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] =

		sr_stat.Uuid

	result["name_label"] =

		sr_stat.NameLabel

	result["name_description"] =

		sr_stat.NameDescription

	result["free_space"] =

		strconv.Itoa(sr_stat.FreeSpace)

	result["total_space"] =

		strconv.Itoa(sr_stat.TotalSpace)

	result["clustered"] =

		sr_stat.Clustered

	result["health"] =

		sr_stat.Health.String()

	return result
}

func ToSrStat(obj interface{}) (resultObj *SrStat) {

	objValue := reflect.ValueOf(obj)
	resultObj = &SrStat{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "uuid":
			if v, ok := keyValue.(interface{}); ok {
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
		case "free_space":
			if v, ok := keyValue.(int); ok {
				resultObj.FreeSpace = v
			}
		case "total_space":
			if v, ok := keyValue.(int); ok {
				resultObj.TotalSpace = v
			}
		case "clustered":
			if v, ok := keyValue.(bool); ok {
				resultObj.Clustered = v
			}
		case "health":
			if v, ok := keyValue.(SrHealth); ok {
				resultObj.Health = v
			}
		}
	}

	return resultObj
}

// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w ProbeResult.gen.go

package go_xen_client

import (
	"reflect"

	"github.com/nilshell/xmlrpc"
)

//ProbeResult: A set of properties that describe one result element of SR.probe. Result elements and properties can change dynamically based on changes to the the SR.probe input-parameters or the target.
type ProbeResult struct {
	Configuration map[string]string // Plugin-specific configuration which describes where and how to locate the storage repository. This may include the physical block device name, a remote NFS server and path or an RBD storage pool.
	Complete      bool              // True if this configuration is complete and can be used to call SR.create. False if it requires further iterative calls to SR.probe, to potentially narrow down on a configuration that can be used.
	Sr            interface{}       // Existing SR found for this configuration
	ExtraInfo     map[string]string // Additional plugin-specific information about this configuration, that might be of use for an API user. This can for example include the LUN or the WWPN.

}

func FromProbeResultToXml(probe_result *ProbeResult) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	configuration := make(xmlrpc.Struct)
	for key, value := range probe_result.Configuration {
		configuration[key] = value
	}
	result["configuration"] = configuration

	result["complete"] = probe_result.Complete

	result["sr"] = probe_result.Sr

	extra_info := make(xmlrpc.Struct)
	for key, value := range probe_result.ExtraInfo {
		extra_info[key] = value
	}
	result["extra_info"] = extra_info

	return result
}

func ToProbeResult(obj interface{}) (resultObj *ProbeResult) {

	objValue := reflect.ValueOf(obj)
	resultObj = &ProbeResult{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
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
		case "complete":
			if v, ok := keyValue.(bool); ok {
				resultObj.Complete = v
			}
		case "sr":
			if v, ok := keyValue.(interface{}); ok {
				resultObj.Sr = v
			}
		case "extra_info":

			resultObj.ExtraInfo = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.ExtraInfo[mapKeyName] = v
				} else {
					resultObj.ExtraInfo[mapKeyName] = ""
				}
			}
		}
	}

	return resultObj
}

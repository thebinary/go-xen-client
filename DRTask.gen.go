// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w DRTask.gen.go

package go_xen_client

import (
	"reflect"

	"github.com/nilshell/xmlrpc"
)

//DRTask: DR task
type DRTask struct {
	Uuid          string   // Unique identifier/object reference
	IntroducedSRs []string // All SRs introduced by this appliance

}

func FromDRTaskToXml(DR_task *DRTask) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] = DR_task.Uuid

	result["introduced_SRs"] = DR_task.IntroducedSRs

	return result
}

func ToDRTask(obj interface{}) (resultObj *DRTask) {

	objValue := reflect.ValueOf(obj)
	resultObj = &DRTask{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.Uuid = v
			}
		case "introduced_SRs":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.IntroducedSRs = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.IntroducedSRs[i] = v
					}
				}
			}
		}
	}

	return resultObj
}

/* GetAllRecords: Return a map of DR_task references to DR_task records for all DR_tasks known to the system. */
func (client *XenClient) DRTaskGetAllRecords() (result map[string]DRTask, err error) {
	obj, err := client.APICall("DR_task.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]DRTask{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToDRTask(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the DR_tasks known to the system. */
func (client *XenClient) DRTaskGetAll() (result []string, err error) {
	obj, err := client.APICall("DR_task.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* Destroy: Destroy the disaster recovery task, detaching and forgetting any SRs introduced which are no longer required */
func (client *XenClient) DRTaskDestroy(self string) (err error) {
	_, err = client.APICall("DR_task.destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Create: Create a disaster recovery task which will query the supplied list of devices */
func (client *XenClient) DRTaskCreate(xtype string, device_config map[string]string, whitelist []string) (result string, err error) {
	obj, err := client.APICall("DR_task.create", xtype, device_config, whitelist)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetIntroducedSRs: Get the introduced_SRs field of the given DR_task. */
func (client *XenClient) DRTaskGetIntroducedSRs(self string) (result []string, err error) {
	obj, err := client.APICall("DR_task.get_introduced_SRs", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetUuid: Get the uuid field of the given DR_task. */
func (client *XenClient) DRTaskGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("DR_task.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the DR_task instance with the specified UUID. */
func (client *XenClient) DRTaskGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("DR_task.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given DR_task. */
func (client *XenClient) DRTaskGetRecord(self string) (result DRTask, err error) {
	obj, err := client.APICall("DR_task.get_record", self)
	if err != nil {
		return
	}

	result = *ToDRTask(obj.(interface{}))

	return
}

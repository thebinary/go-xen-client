// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w Task.gen.go

package go_xen_client

import (
	"reflect"
	"time"

	"github.com/nilshell/xmlrpc"
)

//Task: A long-running asynchronous task
type Task struct {
	Uuid              string                           // Unique identifier/object reference
	NameLabel         string                           // a human-readable name
	NameDescription   string                           // a notes field containing human-readable description
	AllowedOperations []TaskAllowedOperations          // list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	CurrentOperations map[string]TaskAllowedOperations // links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	Created           time.Time                        // Time task was created
	Finished          time.Time                        // Time task finished (i.e. succeeded or failed). If task-status is pending, then the value of this field has no meaning
	Status            TaskStatusType                   // current status of the task
	ResidentOn        string                           // the host on which the task is running
	Progress          float32                          // This field contains the estimated fraction of the task which is complete. This field should not be used to determine whether the task is complete - for this the status field of the task should be used.
	Type              string                           // if the task has completed successfully, this field contains the type of the encoded result (i.e. name of the class whose reference is in the result field). Undefined otherwise.
	Result            string                           // if the task has completed successfully, this field contains the result value (either Void or an object reference). Undefined otherwise.
	ErrorInfo         []string                         // if the task has failed, this field contains the set of associated error strings. Undefined otherwise.
	OtherConfig       map[string]string                // additional configuration
	SubtaskOf         string                           // Ref pointing to the task this is a substask of.
	Subtasks          []string                         // List pointing to all the substasks.
	Backtrace         string                           // Function call trace for debugging.

}

func FromTaskToXml(task *Task) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] =

		task.Uuid

	result["name_label"] =

		task.NameLabel

	result["name_description"] =

		task.NameDescription

	result["allowed_operations"] =

		task.AllowedOperations

	current_operations := make(xmlrpc.Struct)
	for key, value := range task.CurrentOperations {
		current_operations[key] = value
	}
	result["current_operations"] = current_operations

	result["created"] =

		task.Created

	result["finished"] =

		task.Finished

	result["status"] =

		task.Status.String()

	result["resident_on"] =

		task.ResidentOn

	result["progress"] =

		task.Progress

	result["type"] =

		task.Type

	result["result"] =

		task.Result

	result["error_info"] =

		task.ErrorInfo

	other_config := make(xmlrpc.Struct)
	for key, value := range task.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	result["subtask_of"] =

		task.SubtaskOf

	result["subtasks"] =

		task.Subtasks

	result["backtrace"] =

		task.Backtrace

	return result
}

func ToTask(obj interface{}) (resultObj *Task) {

	objValue := reflect.ValueOf(obj)
	resultObj = &Task{}

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
		case "allowed_operations":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.AllowedOperations = make([]TaskAllowedOperations, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(TaskAllowedOperations); ok {
						resultObj.AllowedOperations[i] = v
					}
				}
			}
		case "current_operations":

			resultObj.CurrentOperations = map[string]TaskAllowedOperations{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.CurrentOperations[mapKeyName] = ToTaskAllowedOperations(v)
				} else {
					resultObj.CurrentOperations[mapKeyName] = 0
				}
			}
		case "created":
			if v, ok := keyValue.(time.Time); ok {
				resultObj.Created = v
			}
		case "finished":
			if v, ok := keyValue.(time.Time); ok {
				resultObj.Finished = v
			}
		case "status":
			if v, ok := keyValue.(TaskStatusType); ok {
				resultObj.Status = v
			}
		case "resident_on":
			if v, ok := keyValue.(string); ok {
				resultObj.ResidentOn = v
			}
		case "progress":
			if v, ok := keyValue.(float32); ok {
				resultObj.Progress = v
			}
		case "type":
			if v, ok := keyValue.(string); ok {
				resultObj.Type = v
			}
		case "result":
			if v, ok := keyValue.(string); ok {
				resultObj.Result = v
			}
		case "error_info":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.ErrorInfo = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.ErrorInfo[i] = v
					}
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
		case "subtask_of":
			if v, ok := keyValue.(string); ok {
				resultObj.SubtaskOf = v
			}
		case "subtasks":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.Subtasks = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.Subtasks[i] = v
					}
				}
			}
		case "backtrace":
			if v, ok := keyValue.(string); ok {
				resultObj.Backtrace = v
			}
		}
	}

	return resultObj
}

/* GetAllRecords: Return a map of task references to task records for all tasks known to the system. */
func (client *XenClient) TaskGetAllRecords() (result map[string]Task, err error) {
	obj, err := client.APICall("task.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]Task{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToTask(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the tasks known to the system. */
func (client *XenClient) TaskGetAll() (result []string, err error) {
	obj, err := client.APICall("task.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* SetStatus: Set the task status */
func (client *XenClient) TaskSetStatus(self string, value TaskStatusType) (err error) {
	_, err = client.APICall("task.set_status", self, value.String())
	if err != nil {
		return
	}
	// no return result
	return
}

/* Cancel: Request that a task be cancelled. Note that a task may fail to be cancelled and may complete or fail normally and note that, even when a task does cancel, it might take an arbitrary amount of time. */
func (client *XenClient) TaskCancel(task string) (err error) {
	_, err = client.APICall("task.cancel", task)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Destroy: Destroy the task object */
func (client *XenClient) TaskDestroy(self string) (err error) {
	_, err = client.APICall("task.destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Create: Create a new task object which must be manually destroyed. */
func (client *XenClient) TaskCreate(label string, description string) (result string, err error) {
	obj, err := client.APICall("task.create", label, description)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given task.  If the key is not in that Map, then do nothing. */
func (client *XenClient) TaskRemoveFromOtherConfig(self string, key string) (err error) {
	_, err = client.APICall("task.remove_from_other_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToOtherConfig: Add the given key-value pair to the other_config field of the given task. */
func (client *XenClient) TaskAddToOtherConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("task.add_to_other_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOtherConfig: Set the other_config field of the given task. */
func (client *XenClient) TaskSetOtherConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("task.set_other_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetBacktrace: Get the backtrace field of the given task. */
func (client *XenClient) TaskGetBacktrace(self string) (result string, err error) {
	obj, err := client.APICall("task.get_backtrace", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetSubtasks: Get the subtasks field of the given task. */
func (client *XenClient) TaskGetSubtasks(self string) (result []string, err error) {
	obj, err := client.APICall("task.get_subtasks", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetSubtaskOf: Get the subtask_of field of the given task. */
func (client *XenClient) TaskGetSubtaskOf(self string) (result string, err error) {
	obj, err := client.APICall("task.get_subtask_of", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetOtherConfig: Get the other_config field of the given task. */
func (client *XenClient) TaskGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("task.get_other_config", self)
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

/* GetErrorInfo: Get the error_info field of the given task. */
func (client *XenClient) TaskGetErrorInfo(self string) (result []string, err error) {
	obj, err := client.APICall("task.get_error_info", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetResult: Get the result field of the given task. */
func (client *XenClient) TaskGetResult(self string) (result string, err error) {
	obj, err := client.APICall("task.get_result", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetType: Get the type field of the given task. */
func (client *XenClient) TaskGetType(self string) (result string, err error) {
	obj, err := client.APICall("task.get_type", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetProgress: Get the progress field of the given task. */
func (client *XenClient) TaskGetProgress(self string) (result float32, err error) {
	obj, err := client.APICall("task.get_progress", self)
	if err != nil {
		return
	}
	result = obj.(float32)
	return
}

/* GetResidentOn: Get the resident_on field of the given task. */
func (client *XenClient) TaskGetResidentOn(self string) (result string, err error) {
	obj, err := client.APICall("task.get_resident_on", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetStatus: Get the status field of the given task. */
func (client *XenClient) TaskGetStatus(self string) (result TaskStatusType, err error) {
	obj, err := client.APICall("task.get_status", self)
	if err != nil {
		return
	}

	result = ToTaskStatusType(obj.(string))

	return
}

/* GetFinished: Get the finished field of the given task. */
func (client *XenClient) TaskGetFinished(self string) (result time.Time, err error) {
	obj, err := client.APICall("task.get_finished", self)
	if err != nil {
		return
	}
	result = obj.(time.Time)
	return
}

/* GetCreated: Get the created field of the given task. */
func (client *XenClient) TaskGetCreated(self string) (result time.Time, err error) {
	obj, err := client.APICall("task.get_created", self)
	if err != nil {
		return
	}
	result = obj.(time.Time)
	return
}

/* GetCurrentOperations: Get the current_operations field of the given task. */
func (client *XenClient) TaskGetCurrentOperations(self string) (result map[string]TaskAllowedOperations, err error) {
	obj, err := client.APICall("task.get_current_operations", self)
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]TaskAllowedOperations{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToTaskAllowedOperations(obj.String())
		result[key.String()] = mapObj
	}

	return
}

/* GetAllowedOperations: Get the allowed_operations field of the given task. */
func (client *XenClient) TaskGetAllowedOperations(self string) (result []TaskAllowedOperations, err error) {
	obj, err := client.APICall("task.get_allowed_operations", self)
	if err != nil {
		return
	}

	result = make([]TaskAllowedOperations, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = ToTaskAllowedOperations(value.(string))
	}

	return
}

/* GetNameDescription: Get the name/description field of the given task. */
func (client *XenClient) TaskGetNameDescription(self string) (result string, err error) {
	obj, err := client.APICall("task.get_name_description", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetNameLabel: Get the name/label field of the given task. */
func (client *XenClient) TaskGetNameLabel(self string) (result string, err error) {
	obj, err := client.APICall("task.get_name_label", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given task. */
func (client *XenClient) TaskGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("task.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByNameLabel: Get all the task instances with the given label. */
func (client *XenClient) TaskGetByNameLabel(label string) (result []string, err error) {
	obj, err := client.APICall("task.get_by_name_label", label)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetByUuid: Get a reference to the task instance with the specified UUID. */
func (client *XenClient) TaskGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("task.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given task. */
func (client *XenClient) TaskGetRecord(self string) (result Task, err error) {
	obj, err := client.APICall("task.get_record", self)
	if err != nil {
		return
	}

	result = *ToTask(obj.(interface{}))

	return
}

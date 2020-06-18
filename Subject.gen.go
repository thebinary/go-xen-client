// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w Subject.gen.go

package go_xen_client

import (
	"reflect"

	"github.com/nilshell/xmlrpc"
)

//Subject: A user or group that can log in xapi
type Subject struct {
	Uuid              string            // Unique identifier/object reference
	SubjectIdentifier string            // the subject identifier, unique in the external directory service
	OtherConfig       map[string]string // additional configuration
	Roles             []string          // the roles associated with this subject

}

func FromSubjectToXml(subject *Subject) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] =

		subject.Uuid

	result["subject_identifier"] =

		subject.SubjectIdentifier

	other_config := make(xmlrpc.Struct)
	for key, value := range subject.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	result["roles"] =

		subject.Roles

	return result
}

func ToSubject(obj interface{}) (resultObj *Subject) {

	objValue := reflect.ValueOf(obj)
	resultObj = &Subject{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.Uuid = v
			}
		case "subject_identifier":
			if v, ok := keyValue.(string); ok {
				resultObj.SubjectIdentifier = v
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
		case "roles":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.Roles = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.Roles[i] = v
					}
				}
			}
		}
	}

	return resultObj
}

/* GetAllRecords: Return a map of subject references to subject records for all subjects known to the system. */
func (client *XenClient) SubjectGetAllRecords() (result map[string]Subject, err error) {
	obj, err := client.APICall("subject.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]Subject{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToSubject(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the subjects known to the system. */
func (client *XenClient) SubjectGetAll() (result []string, err error) {
	obj, err := client.APICall("subject.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetPermissionsNameLabel: This call returns a list of permission names given a subject */
func (client *XenClient) SubjectGetPermissionsNameLabel(self string) (result []string, err error) {
	obj, err := client.APICall("subject.get_permissions_name_label", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* RemoveFromRoles: This call removes a role from a subject */
func (client *XenClient) SubjectRemoveFromRoles(self string, role string) (err error) {
	_, err = client.APICall("subject.remove_from_roles", self, role)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToRoles: This call adds a new role to a subject */
func (client *XenClient) SubjectAddToRoles(self string, role string) (err error) {
	_, err = client.APICall("subject.add_to_roles", self, role)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetRoles: Get the roles field of the given subject. */
func (client *XenClient) SubjectGetRoles(self string) (result []string, err error) {
	obj, err := client.APICall("subject.get_roles", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetOtherConfig: Get the other_config field of the given subject. */
func (client *XenClient) SubjectGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("subject.get_other_config", self)
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

/* GetSubjectIdentifier: Get the subject_identifier field of the given subject. */
func (client *XenClient) SubjectGetSubjectIdentifier(self string) (result string, err error) {
	obj, err := client.APICall("subject.get_subject_identifier", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given subject. */
func (client *XenClient) SubjectGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("subject.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* Destroy: Destroy the specified subject instance. */
func (client *XenClient) SubjectDestroy(self string) (err error) {
	_, err = client.APICall("subject.destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Create: Create a new subject instance, and return its handle.
The constructor args are: subject_identifier, other_config (* = non-optional). */
func (client *XenClient) SubjectCreate(args Subject) (result string, err error) {
	obj, err := client.APICall("subject.create", FromSubjectToXml(&args))
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the subject instance with the specified UUID. */
func (client *XenClient) SubjectGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("subject.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given subject. */
func (client *XenClient) SubjectGetRecord(self string) (result Subject, err error) {
	obj, err := client.APICall("subject.get_record", self)
	if err != nil {
		return
	}

	result = *ToSubject(obj.(interface{}))

	return
}

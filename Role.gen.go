// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w Role.gen.go

package go_xen_client

import (
	"reflect"

	"github.com/nilshell/xmlrpc"
)

//Role: A set of permissions associated with a subject
type Role struct {
	Uuid            string   // Unique identifier/object reference
	NameLabel       string   // a short user-friendly name for the role
	NameDescription string   // what this role is for
	Subroles        []string // a list of pointers to other roles or permissions

}

func FromRoleToXml(role *Role) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] = role.Uuid

	result["name_label"] = role.NameLabel

	result["name_description"] = role.NameDescription

	result["subroles"] = role.Subroles

	return result
}

func ToRole(obj interface{}) (resultObj *Role) {

	objValue := reflect.ValueOf(obj)
	resultObj = &Role{}

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
		case "subroles":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.Subroles = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.Subroles[i] = v
					}
				}
			}
		}
	}

	return resultObj
}

/* GetAllRecords: Return a map of role references to role records for all roles known to the system. */
func (client *XenClient) RoleGetAllRecords() (result map[string]Role, err error) {
	obj, err := client.APICall("role.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]Role{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToRole(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the roles known to the system. */
func (client *XenClient) RoleGetAll() (result []string, err error) {
	obj, err := client.APICall("role.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetByPermissionNameLabel: This call returns a list of roles given a permission name */
func (client *XenClient) RoleGetByPermissionNameLabel(label string) (result []string, err error) {
	obj, err := client.APICall("role.get_by_permission_name_label", label)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetByPermission: This call returns a list of roles given a permission */
func (client *XenClient) RoleGetByPermission(permission string) (result []string, err error) {
	obj, err := client.APICall("role.get_by_permission", permission)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetPermissionsNameLabel: This call returns a list of permission names given a role */
func (client *XenClient) RoleGetPermissionsNameLabel(self string) (result []string, err error) {
	obj, err := client.APICall("role.get_permissions_name_label", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetPermissions: This call returns a list of permissions given a role */
func (client *XenClient) RoleGetPermissions(self string) (result []string, err error) {
	obj, err := client.APICall("role.get_permissions", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetSubroles: Get the subroles field of the given role. */
func (client *XenClient) RoleGetSubroles(self string) (result []string, err error) {
	obj, err := client.APICall("role.get_subroles", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetNameDescription: Get the name/description field of the given role. */
func (client *XenClient) RoleGetNameDescription(self string) (result string, err error) {
	obj, err := client.APICall("role.get_name_description", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetNameLabel: Get the name/label field of the given role. */
func (client *XenClient) RoleGetNameLabel(self string) (result string, err error) {
	obj, err := client.APICall("role.get_name_label", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given role. */
func (client *XenClient) RoleGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("role.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByNameLabel: Get all the role instances with the given label. */
func (client *XenClient) RoleGetByNameLabel(label string) (result []string, err error) {
	obj, err := client.APICall("role.get_by_name_label", label)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetByUuid: Get a reference to the role instance with the specified UUID. */
func (client *XenClient) RoleGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("role.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given role. */
func (client *XenClient) RoleGetRecord(self string) (result Role, err error) {
	obj, err := client.APICall("role.get_record", self)
	if err != nil {
		return
	}

	result = *ToRole(obj.(interface{}))

	return
}

// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w User.gen.go

package go_xen_client

import (
	"reflect"

	"github.com/nilshell/xmlrpc"
)

//User: A user of the system
type User struct {
	Uuid        string            // Unique identifier/object reference
	ShortName   string            // short name (e.g. userid)
	Fullname    string            // full name
	OtherConfig map[string]string // additional configuration

}

func FromUserToXml(user *User) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] =

		user.Uuid

	result["short_name"] =

		user.ShortName

	result["fullname"] =

		user.Fullname

	other_config := make(xmlrpc.Struct)
	for key, value := range user.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	return result
}

func ToUser(obj interface{}) (resultObj *User) {

	objValue := reflect.ValueOf(obj)
	resultObj = &User{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.Uuid = v
			}
		case "short_name":
			if v, ok := keyValue.(string); ok {
				resultObj.ShortName = v
			}
		case "fullname":
			if v, ok := keyValue.(string); ok {
				resultObj.Fullname = v
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

/* RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given user.  If the key is not in that Map, then do nothing. */
func (client *XenClient) UserRemoveFromOtherConfig(self string, key string) (err error) {
	_, err = client.APICall("user.remove_from_other_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToOtherConfig: Add the given key-value pair to the other_config field of the given user. */
func (client *XenClient) UserAddToOtherConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("user.add_to_other_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOtherConfig: Set the other_config field of the given user. */
func (client *XenClient) UserSetOtherConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("user.set_other_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetFullname: Set the fullname field of the given user. */
func (client *XenClient) UserSetFullname(self string, value string) (err error) {
	_, err = client.APICall("user.set_fullname", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetOtherConfig: Get the other_config field of the given user. */
func (client *XenClient) UserGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("user.get_other_config", self)
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

/* GetFullname: Get the fullname field of the given user. */
func (client *XenClient) UserGetFullname(self string) (result string, err error) {
	obj, err := client.APICall("user.get_fullname", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetShortName: Get the short_name field of the given user. */
func (client *XenClient) UserGetShortName(self string) (result string, err error) {
	obj, err := client.APICall("user.get_short_name", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given user. */
func (client *XenClient) UserGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("user.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* Destroy: Destroy the specified user instance. */
func (client *XenClient) UserDestroy(self string) (err error) {
	_, err = client.APICall("user.destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Create: Create a new user instance, and return its handle.
The constructor args are: short_name*, fullname*, other_config (* = non-optional). */
func (client *XenClient) UserCreate(args User) (result string, err error) {
	obj, err := client.APICall("user.create", FromUserToXml(&args))
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the user instance with the specified UUID. */
func (client *XenClient) UserGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("user.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given user. */
func (client *XenClient) UserGetRecord(self string) (result User, err error) {
	obj, err := client.APICall("user.get_record", self)
	if err != nil {
		return
	}

	result = *ToUser(obj.(interface{}))

	return
}

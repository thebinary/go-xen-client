// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w Secret.gen.go

package go_xen_client

import (
	"reflect"

	"github.com/nilshell/xmlrpc"
)

//Secret: A secret
type Secret struct {
	Uuid        string            // Unique identifier/object reference
	Value       string            // the secret
	OtherConfig map[string]string // other_config

}

func FromSecretToXml(secret *Secret) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] =

		secret.Uuid

	result["value"] =

		secret.Value

	other_config := make(xmlrpc.Struct)
	for key, value := range secret.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	return result
}

func ToSecret(obj interface{}) (resultObj *Secret) {

	objValue := reflect.ValueOf(obj)
	resultObj = &Secret{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.Uuid = v
			}
		case "value":
			if v, ok := keyValue.(string); ok {
				resultObj.Value = v
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

/* GetAllRecords: Return a map of secret references to secret records for all secrets known to the system. */
func (client *XenClient) SecretGetAllRecords() (result map[string]Secret, err error) {
	obj, err := client.APICall("secret.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]Secret{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToSecret(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the secrets known to the system. */
func (client *XenClient) SecretGetAll() (result []string, err error) {
	obj, err := client.APICall("secret.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given secret.  If the key is not in that Map, then do nothing. */
func (client *XenClient) SecretRemoveFromOtherConfig(self string, key string) (err error) {
	_, err = client.APICall("secret.remove_from_other_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToOtherConfig: Add the given key-value pair to the other_config field of the given secret. */
func (client *XenClient) SecretAddToOtherConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("secret.add_to_other_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOtherConfig: Set the other_config field of the given secret. */
func (client *XenClient) SecretSetOtherConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("secret.set_other_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetValue: Set the value field of the given secret. */
func (client *XenClient) SecretSetValue(self string, value string) (err error) {
	_, err = client.APICall("secret.set_value", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetOtherConfig: Get the other_config field of the given secret. */
func (client *XenClient) SecretGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("secret.get_other_config", self)
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

/* GetValue: Get the value field of the given secret. */
func (client *XenClient) SecretGetValue(self string) (result string, err error) {
	obj, err := client.APICall("secret.get_value", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given secret. */
func (client *XenClient) SecretGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("secret.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* Destroy: Destroy the specified secret instance. */
func (client *XenClient) SecretDestroy(self string) (err error) {
	_, err = client.APICall("secret.destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Create: Create a new secret instance, and return its handle.
The constructor args are: value*, other_config (* = non-optional). */
func (client *XenClient) SecretCreate(args Secret) (result string, err error) {
	obj, err := client.APICall("secret.create", FromSecretToXml(&args))
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the secret instance with the specified UUID. */
func (client *XenClient) SecretGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("secret.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given secret. */
func (client *XenClient) SecretGetRecord(self string) (result Secret, err error) {
	obj, err := client.APICall("secret.get_record", self)
	if err != nil {
		return
	}

	result = *ToSecret(obj.(interface{}))

	return
}

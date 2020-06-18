// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w Message.gen.go

package go_xen_client

import (
	"reflect"
	"strconv"
	"time"

	"github.com/nilshell/xmlrpc"
)

//Message: An message for the attention of the administrator
type Message struct {
	Uuid      string    // Unique identifier/object reference
	Name      string    // The name of the message
	Priority  int       // The message priority, 0 being low priority
	Cls       Cls       // The class of the object this message is associated with
	ObjUuid   string    // The uuid of the object this message is associated with
	Timestamp time.Time // The time at which the message was created
	Body      string    // The body of the message

}

func FromMessageToXml(message *Message) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] = message.Uuid

	result["name"] = message.Name

	result["priority"] = strconv.Itoa(message.Priority)

	result["cls"] = message.Cls.String()

	result["obj_uuid"] = message.ObjUuid

	result["timestamp"] = message.Timestamp

	result["body"] = message.Body

	return result
}

func ToMessage(obj interface{}) (resultObj *Message) {

	objValue := reflect.ValueOf(obj)
	resultObj = &Message{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.Uuid = v
			}
		case "name":
			if v, ok := keyValue.(string); ok {
				resultObj.Name = v
			}
		case "priority":
			if v, ok := keyValue.(int); ok {
				resultObj.Priority = v
			}
		case "cls":
			if v, ok := keyValue.(Cls); ok {
				resultObj.Cls = v
			}
		case "obj_uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.ObjUuid = v
			}
		case "timestamp":
			if v, ok := keyValue.(time.Time); ok {
				resultObj.Timestamp = v
			}
		case "body":
			if v, ok := keyValue.(string); ok {
				resultObj.Body = v
			}
		}
	}

	return resultObj
}

/* GetAllRecordsWhere:  */
func (client *XenClient) MessageGetAllRecordsWhere(expr string) (result map[string]Message, err error) {
	obj, err := client.APICall("message.get_all_records_where", expr)
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]Message{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToMessage(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAllRecords:  */
func (client *XenClient) MessageGetAllRecords() (result map[string]Message, err error) {
	obj, err := client.APICall("message.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]Message{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToMessage(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetByUuid:  */
func (client *XenClient) MessageGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("message.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord:  */
func (client *XenClient) MessageGetRecord(self string) (result Message, err error) {
	obj, err := client.APICall("message.get_record", self)
	if err != nil {
		return
	}

	result = *ToMessage(obj.(interface{}))

	return
}

/* GetSince:  */
func (client *XenClient) MessageGetSince(since time.Time) (result map[string]Message, err error) {
	obj, err := client.APICall("message.get_since", since)
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]Message{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToMessage(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll:  */
func (client *XenClient) MessageGetAll() (result []string, err error) {
	obj, err := client.APICall("message.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* Get:  */
func (client *XenClient) MessageGet(cls Cls, obj_uuid string, since time.Time) (result map[string]Message, err error) {
	obj, err := client.APICall("message.get", cls.String(), obj_uuid, since)
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]Message{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToMessage(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* Destroy:  */
func (client *XenClient) MessageDestroy(self string) (err error) {
	_, err = client.APICall("message.destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Create:  */
func (client *XenClient) MessageCreate(name string, priority int, cls Cls, obj_uuid string, body string) (result string, err error) {
	obj, err := client.APICall("message.create", name, priority, cls.String(), obj_uuid, body)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

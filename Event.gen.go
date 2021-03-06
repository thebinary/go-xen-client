// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w Event.gen.go

package go_xen_client

import (
	"reflect"
	"strconv"
	"time"

	"github.com/nilshell/xmlrpc"
)

//Event: Asynchronous event registration and handling
type Event struct {
	Snapshot  interface{}    // The record of the database object that was added, changed or deleted
	Id        int            // An ID, monotonically increasing, and local to the current session
	Timestamp time.Time      // The time at which the event occurred
	Class     string         // The name of the class of the object that changed
	Operation EventOperation // The operation that was performed
	Ref       string         // A reference to the object that changed
	ObjUuid   string         // The uuid of the object that changed

}

func FromEventToXml(event *Event) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["snapshot"] = event.Snapshot

	result["id"] = strconv.Itoa(event.Id)

	result["timestamp"] = event.Timestamp

	result["class"] = event.Class

	result["operation"] = event.Operation.String()

	result["ref"] = event.Ref

	result["obj_uuid"] = event.ObjUuid

	return result
}

func ToEvent(obj interface{}) (resultObj *Event) {

	objValue := reflect.ValueOf(obj)
	resultObj = &Event{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "snapshot":
			if v, ok := keyValue.(interface{}); ok {
				resultObj.Snapshot = v
			}
		case "id":
			if v, ok := keyValue.(int); ok {
				resultObj.Id = v
			}
		case "timestamp":
			if v, ok := keyValue.(time.Time); ok {
				resultObj.Timestamp = v
			}
		case "class":
			if v, ok := keyValue.(string); ok {
				resultObj.Class = v
			}
		case "operation":
			if v, ok := keyValue.(EventOperation); ok {
				resultObj.Operation = v
			}
		case "ref":
			if v, ok := keyValue.(string); ok {
				resultObj.Ref = v
			}
		case "obj_uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.ObjUuid = v
			}
		}
	}

	return resultObj
}

/* Inject: Injects an artificial event on the given object and returns the corresponding ID in the form of a token, which can be used as a point of reference for database events. For example, to check whether an object has reached the right state before attempting an operation, one can inject an artificial event on the object and wait until the token returned by consecutive event.from calls is lexicographically greater than the one returned by event.inject. */
func (client *XenClient) EventInject(class string, ref string) (result string, err error) {
	obj, err := client.APICall("event.inject", class, ref)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetCurrentId: Return the ID of the next event to be generated by the system */
func (client *XenClient) EventGetCurrentId() (result int, err error) {
	obj, err := client.APICall("event.get_current_id")
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* From: Blocking call which returns a new token and a (possibly empty) batch of events. The returned token can be used in subsequent calls to this function. */
func (client *XenClient) EventFrom(classes []string, token string, timeout float32) (result interface{}, err error) {
	obj, err := client.APICall("event.from", classes, token, timeout)
	if err != nil {
		return
	}
	result = obj
	return
}

/* Next: Blocking call which returns a (possibly empty) batch of events. This method is only recommended for legacy use. New development should use event.from which supercedes this method. */
func (client *XenClient) EventNext() (result []Event, err error) {
	obj, err := client.APICall("event.next")
	if err != nil {
		return
	}

	result = make([]Event, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = *ToEvent(value)
	}

	return
}

/* Unregister: Removes this session's registration with the event system for a set of given classes. This method is only recommended for legacy use in conjunction with event.next. */
func (client *XenClient) EventUnregister(classes []string) (err error) {
	_, err = client.APICall("event.unregister", classes)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Register: Registers this session with the event system for a set of given classes. This method is only recommended for legacy use in conjunction with event.next. */
func (client *XenClient) EventRegister(classes []string) (err error) {
	_, err = client.APICall("event.register", classes)
	if err != nil {
		return
	}
	// no return result
	return
}

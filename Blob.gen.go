// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w Blob.gen.go

package go_xen_client

import (
	"reflect"
	"strconv"
	"time"

	"github.com/nilshell/xmlrpc"
)

//Blob: A placeholder for a binary blob
type Blob struct {
	Uuid            string    // Unique identifier/object reference
	NameLabel       string    // a human-readable name
	NameDescription string    // a notes field containing human-readable description
	Size            int       // Size of the binary data, in bytes
	Public          bool      // True if the blob is publicly accessible
	LastUpdated     time.Time // Time at which the data in the blob was last updated
	MimeType        string    // The mime type associated with this object. Defaults to 'application/octet-stream' if the empty string is supplied

}

func FromBlobToXml(blob *Blob) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] =

		blob.Uuid

	result["name_label"] =

		blob.NameLabel

	result["name_description"] =

		blob.NameDescription

	result["size"] =

		strconv.Itoa(blob.Size)

	result["public"] =

		blob.Public

	result["last_updated"] =

		blob.LastUpdated

	result["mime_type"] =

		blob.MimeType

	return result
}

func ToBlob(obj interface{}) (resultObj *Blob) {

	objValue := reflect.ValueOf(obj)
	resultObj = &Blob{}

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
		case "size":
			if v, ok := keyValue.(int); ok {
				resultObj.Size = v
			}
		case "public":
			if v, ok := keyValue.(bool); ok {
				resultObj.Public = v
			}
		case "last_updated":
			if v, ok := keyValue.(time.Time); ok {
				resultObj.LastUpdated = v
			}
		case "mime_type":
			if v, ok := keyValue.(string); ok {
				resultObj.MimeType = v
			}
		}
	}

	return resultObj
}

/* GetAllRecords: Return a map of blob references to blob records for all blobs known to the system. */
func (client *XenClient) BlobGetAllRecords() (result map[string]Blob, err error) {
	obj, err := client.APICall("blob.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]Blob{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToBlob(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the blobs known to the system. */
func (client *XenClient) BlobGetAll() (result []string, err error) {
	obj, err := client.APICall("blob.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* Destroy:  */
func (client *XenClient) BlobDestroy(self string) (err error) {
	_, err = client.APICall("blob.destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Create: Create a placeholder for a binary blob */
func (client *XenClient) BlobCreate(mime_type string, public bool) (result string, err error) {
	obj, err := client.APICall("blob.create", mime_type, public)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* SetPublic: Set the public field of the given blob. */
func (client *XenClient) BlobSetPublic(self string, value bool) (err error) {
	_, err = client.APICall("blob.set_public", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetNameDescription: Set the name/description field of the given blob. */
func (client *XenClient) BlobSetNameDescription(self string, value string) (err error) {
	_, err = client.APICall("blob.set_name_description", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetNameLabel: Set the name/label field of the given blob. */
func (client *XenClient) BlobSetNameLabel(self string, value string) (err error) {
	_, err = client.APICall("blob.set_name_label", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetMimeType: Get the mime_type field of the given blob. */
func (client *XenClient) BlobGetMimeType(self string) (result string, err error) {
	obj, err := client.APICall("blob.get_mime_type", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetLastUpdated: Get the last_updated field of the given blob. */
func (client *XenClient) BlobGetLastUpdated(self string) (result time.Time, err error) {
	obj, err := client.APICall("blob.get_last_updated", self)
	if err != nil {
		return
	}
	result = obj.(time.Time)
	return
}

/* GetPublic: Get the public field of the given blob. */
func (client *XenClient) BlobGetPublic(self string) (result bool, err error) {
	obj, err := client.APICall("blob.get_public", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetSize: Get the size field of the given blob. */
func (client *XenClient) BlobGetSize(self string) (result int, err error) {
	obj, err := client.APICall("blob.get_size", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetNameDescription: Get the name/description field of the given blob. */
func (client *XenClient) BlobGetNameDescription(self string) (result string, err error) {
	obj, err := client.APICall("blob.get_name_description", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetNameLabel: Get the name/label field of the given blob. */
func (client *XenClient) BlobGetNameLabel(self string) (result string, err error) {
	obj, err := client.APICall("blob.get_name_label", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given blob. */
func (client *XenClient) BlobGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("blob.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByNameLabel: Get all the blob instances with the given label. */
func (client *XenClient) BlobGetByNameLabel(label string) (result []string, err error) {
	obj, err := client.APICall("blob.get_by_name_label", label)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetByUuid: Get a reference to the blob instance with the specified UUID. */
func (client *XenClient) BlobGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("blob.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given blob. */
func (client *XenClient) BlobGetRecord(self string) (result Blob, err error) {
	obj, err := client.APICall("blob.get_record", self)
	if err != nil {
		return
	}

	result = *ToBlob(obj.(interface{}))

	return
}

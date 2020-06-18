// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w Auth.gen.go

package go_xen_client

import (
	"reflect"

	"github.com/nilshell/xmlrpc"
)

//Auth: Management of remote authentication services
type Auth struct {
}

func FromAuthToXml(auth *Auth) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	return result
}

func ToAuth(obj interface{}) (resultObj *Auth) {
	return &Auth{}
}

/* GetGroupMembership: This calls queries the external directory service to obtain the transitively-closed set of groups that the the subject_identifier is member of. */
func (client *XenClient) AuthGetGroupMembership(subject_identifier string) (result []string, err error) {
	obj, err := client.APICall("auth.get_group_membership", subject_identifier)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetSubjectInformationFromIdentifier: This call queries the external directory service to obtain the user information (e.g. username, organization etc) from the specified subject_identifier */
func (client *XenClient) AuthGetSubjectInformationFromIdentifier(subject_identifier string) (result map[string]string, err error) {
	obj, err := client.APICall("auth.get_subject_information_from_identifier", subject_identifier)
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

/* GetSubjectIdentifier: This call queries the external directory service to obtain the subject_identifier as a string from the human-readable subject_name */
func (client *XenClient) AuthGetSubjectIdentifier(subject_name string) (result string, err error) {
	obj, err := client.APICall("auth.get_subject_identifier", subject_name)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

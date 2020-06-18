// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w Session.gen.go

package go_xen_client

import (
	"reflect"
	"time"

	"github.com/nilshell/xmlrpc"
)

//Session: A session
type Session struct {
	Uuid             string            // Unique identifier/object reference
	ThisHost         string            // Currently connected host
	ThisUser         string            // Currently connected user
	LastActive       time.Time         // Timestamp for last time session was active
	Pool             bool              // True if this session relates to a intra-pool login, false otherwise
	OtherConfig      map[string]string // additional configuration
	IsLocalSuperuser bool              // true iff this session was created using local superuser credentials
	Subject          string            // references the subject instance that created the session. If a session instance has is_local_superuser set, then the value of this field is undefined.
	ValidationTime   time.Time         // time when session was last validated
	AuthUserSid      string            // the subject identifier of the user that was externally authenticated. If a session instance has is_local_superuser set, then the value of this field is undefined.
	AuthUserName     string            // the subject name of the user that was externally authenticated. If a session instance has is_local_superuser set, then the value of this field is undefined.
	RbacPermissions  []string          // list with all RBAC permissions for this session
	Tasks            []string          // list of tasks created using the current session
	Parent           string            // references the parent session that created this session
	Originator       string            // a key string provided by a API user to distinguish itself from other users sharing the same login name

}

func FromSessionToXml(session *Session) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] = session.Uuid

	result["this_host"] = session.ThisHost

	result["this_user"] = session.ThisUser

	result["last_active"] = session.LastActive

	result["pool"] = session.Pool

	other_config := make(xmlrpc.Struct)
	for key, value := range session.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	result["is_local_superuser"] = session.IsLocalSuperuser

	result["subject"] = session.Subject

	result["validation_time"] = session.ValidationTime

	result["auth_user_sid"] = session.AuthUserSid

	result["auth_user_name"] = session.AuthUserName

	result["rbac_permissions"] = session.RbacPermissions

	result["tasks"] = session.Tasks

	result["parent"] = session.Parent

	result["originator"] = session.Originator

	return result
}

func ToSession(obj interface{}) (resultObj *Session) {

	objValue := reflect.ValueOf(obj)
	resultObj = &Session{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.Uuid = v
			}
		case "this_host":
			if v, ok := keyValue.(string); ok {
				resultObj.ThisHost = v
			}
		case "this_user":
			if v, ok := keyValue.(string); ok {
				resultObj.ThisUser = v
			}
		case "last_active":
			if v, ok := keyValue.(time.Time); ok {
				resultObj.LastActive = v
			}
		case "pool":
			if v, ok := keyValue.(bool); ok {
				resultObj.Pool = v
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
		case "is_local_superuser":
			if v, ok := keyValue.(bool); ok {
				resultObj.IsLocalSuperuser = v
			}
		case "subject":
			if v, ok := keyValue.(string); ok {
				resultObj.Subject = v
			}
		case "validation_time":
			if v, ok := keyValue.(time.Time); ok {
				resultObj.ValidationTime = v
			}
		case "auth_user_sid":
			if v, ok := keyValue.(string); ok {
				resultObj.AuthUserSid = v
			}
		case "auth_user_name":
			if v, ok := keyValue.(string); ok {
				resultObj.AuthUserName = v
			}
		case "rbac_permissions":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.RbacPermissions = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.RbacPermissions[i] = v
					}
				}
			}
		case "tasks":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.Tasks = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.Tasks[i] = v
					}
				}
			}
		case "parent":
			if v, ok := keyValue.(string); ok {
				resultObj.Parent = v
			}
		case "originator":
			if v, ok := keyValue.(string); ok {
				resultObj.Originator = v
			}
		}
	}

	return resultObj
}

/* LogoutSubjectIdentifier: Log out all sessions associated to a user subject-identifier, except the session associated with the context calling this function */
func (client *XenClient) SessionLogoutSubjectIdentifier(subject_identifier string) (err error) {
	_, err = client.APICall("session.logout_subject_identifier", subject_identifier)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetAllSubjectIdentifiers: Return a list of all the user subject-identifiers of all existing sessions */
func (client *XenClient) SessionGetAllSubjectIdentifiers() (result []string, err error) {
	obj, err := client.APICall("session.get_all_subject_identifiers")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* LocalLogout: Log out of local session. */
func (client *XenClient) SessionLocalLogout() (err error) {
	_, err = client.APICall("session.local_logout")
	if err != nil {
		return
	}
	// no return result
	return
}

/* CreateFromDbFile:  */
func (client *XenClient) SessionCreateFromDbFile(filename string) (result string, err error) {
	obj, err := client.APICall("session.create_from_db_file", filename)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* SlaveLocalLoginWithPassword: Authenticate locally against a slave in emergency mode. Note the resulting sessions are only good for use on this host. */
func (client *XenClient) SessionSlaveLocalLoginWithPassword(uname string, pwd string) (result string, err error) {
	obj, err := client.APICall("session.slave_local_login_with_password", uname, pwd)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* ChangePassword: Change the account password; if your session is authenticated with root priviledges then the old_pwd is validated and the new_pwd is set regardless */
func (client *XenClient) SessionChangePassword(old_pwd string, new_pwd string) (err error) {
	_, err = client.APICall("session.change_password", old_pwd, new_pwd)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Logout: Log out of a session */
func (client *XenClient) SessionLogout() (err error) {
	_, err = client.APICall("session.logout")
	if err != nil {
		return
	}
	// no return result
	return
}

/* LoginWithPassword: Attempt to authenticate the user, returning a session reference if successful */
func (client *XenClient) SessionLoginWithPassword(uname string, pwd string, version string, originator string) (result string, err error) {
	obj, err := client.APICall("session.login_with_password", uname, pwd, version, originator)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given session.  If the key is not in that Map, then do nothing. */
func (client *XenClient) SessionRemoveFromOtherConfig(self string, key string) (err error) {
	_, err = client.APICall("session.remove_from_other_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToOtherConfig: Add the given key-value pair to the other_config field of the given session. */
func (client *XenClient) SessionAddToOtherConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("session.add_to_other_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOtherConfig: Set the other_config field of the given session. */
func (client *XenClient) SessionSetOtherConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("session.set_other_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetOriginator: Get the originator field of the given session. */
func (client *XenClient) SessionGetOriginator(self string) (result string, err error) {
	obj, err := client.APICall("session.get_originator", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetParent: Get the parent field of the given session. */
func (client *XenClient) SessionGetParent(self string) (result string, err error) {
	obj, err := client.APICall("session.get_parent", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetTasks: Get the tasks field of the given session. */
func (client *XenClient) SessionGetTasks(self string) (result []string, err error) {
	obj, err := client.APICall("session.get_tasks", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetRbacPermissions: Get the rbac_permissions field of the given session. */
func (client *XenClient) SessionGetRbacPermissions(self string) (result []string, err error) {
	obj, err := client.APICall("session.get_rbac_permissions", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetAuthUserName: Get the auth_user_name field of the given session. */
func (client *XenClient) SessionGetAuthUserName(self string) (result string, err error) {
	obj, err := client.APICall("session.get_auth_user_name", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetAuthUserSid: Get the auth_user_sid field of the given session. */
func (client *XenClient) SessionGetAuthUserSid(self string) (result string, err error) {
	obj, err := client.APICall("session.get_auth_user_sid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetValidationTime: Get the validation_time field of the given session. */
func (client *XenClient) SessionGetValidationTime(self string) (result time.Time, err error) {
	obj, err := client.APICall("session.get_validation_time", self)
	if err != nil {
		return
	}
	result = obj.(time.Time)
	return
}

/* GetSubject: Get the subject field of the given session. */
func (client *XenClient) SessionGetSubject(self string) (result string, err error) {
	obj, err := client.APICall("session.get_subject", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetIsLocalSuperuser: Get the is_local_superuser field of the given session. */
func (client *XenClient) SessionGetIsLocalSuperuser(self string) (result bool, err error) {
	obj, err := client.APICall("session.get_is_local_superuser", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetOtherConfig: Get the other_config field of the given session. */
func (client *XenClient) SessionGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("session.get_other_config", self)
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

/* GetPool: Get the pool field of the given session. */
func (client *XenClient) SessionGetPool(self string) (result bool, err error) {
	obj, err := client.APICall("session.get_pool", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetLastActive: Get the last_active field of the given session. */
func (client *XenClient) SessionGetLastActive(self string) (result time.Time, err error) {
	obj, err := client.APICall("session.get_last_active", self)
	if err != nil {
		return
	}
	result = obj.(time.Time)
	return
}

/* GetThisUser: Get the this_user field of the given session. */
func (client *XenClient) SessionGetThisUser(self string) (result string, err error) {
	obj, err := client.APICall("session.get_this_user", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetThisHost: Get the this_host field of the given session. */
func (client *XenClient) SessionGetThisHost(self string) (result string, err error) {
	obj, err := client.APICall("session.get_this_host", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given session. */
func (client *XenClient) SessionGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("session.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the session instance with the specified UUID. */
func (client *XenClient) SessionGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("session.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given session. */
func (client *XenClient) SessionGetRecord(self string) (result Session, err error) {
	obj, err := client.APICall("session.get_record", self)
	if err != nil {
		return
	}

	result = *ToSession(obj.(interface{}))

	return
}

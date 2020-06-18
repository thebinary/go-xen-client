// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w VMPP.gen.go

package go_xen_client

import (
	"reflect"
	"strconv"
	"time"

	"github.com/nilshell/xmlrpc"
)

//VMPP: VM Protection Policy
type VMPP struct {
	Uuid                 string                // Unique identifier/object reference
	NameLabel            string                // a human-readable name
	NameDescription      string                // a notes field containing human-readable description
	IsPolicyEnabled      bool                  // enable or disable this policy
	BackupType           VmppBackupType        // type of the backup sub-policy
	BackupRetentionValue int                   // maximum number of backups that should be stored at any time
	BackupFrequency      VmppBackupFrequency   // frequency of the backup schedule
	BackupSchedule       map[string]string     // schedule of the backup containing 'hour', 'min', 'days'. Date/time-related information is in Local Timezone
	IsBackupRunning      bool                  // true if this protection policy's backup is running
	BackupLastRunTime    time.Time             // time of the last backup
	ArchiveTargetType    VmppArchiveTargetType // type of the archive target config
	ArchiveTargetConfig  map[string]string     // configuration for the archive, including its 'location', 'username', 'password'
	ArchiveFrequency     VmppArchiveFrequency  // frequency of the archive schedule
	ArchiveSchedule      map[string]string     // schedule of the archive containing 'hour', 'min', 'days'. Date/time-related information is in Local Timezone
	IsArchiveRunning     bool                  // true if this protection policy's archive is running
	ArchiveLastRunTime   time.Time             // time of the last archive
	VMs                  []string              // all VMs attached to this protection policy
	IsAlarmEnabled       bool                  // true if alarm is enabled for this policy
	AlarmConfig          map[string]string     // configuration for the alarm
	RecentAlerts         []string              // recent alerts

}

func FromVMPPToXml(VMPP *VMPP) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] = VMPP.Uuid

	result["name_label"] = VMPP.NameLabel

	result["name_description"] = VMPP.NameDescription

	result["is_policy_enabled"] = VMPP.IsPolicyEnabled

	result["backup_type"] = VMPP.BackupType.String()

	result["backup_retention_value"] = strconv.Itoa(VMPP.BackupRetentionValue)

	result["backup_frequency"] = VMPP.BackupFrequency.String()

	backup_schedule := make(xmlrpc.Struct)
	for key, value := range VMPP.BackupSchedule {
		backup_schedule[key] = value
	}
	result["backup_schedule"] = backup_schedule

	result["is_backup_running"] = VMPP.IsBackupRunning

	result["backup_last_run_time"] = VMPP.BackupLastRunTime

	result["archive_target_type"] = VMPP.ArchiveTargetType.String()

	archive_target_config := make(xmlrpc.Struct)
	for key, value := range VMPP.ArchiveTargetConfig {
		archive_target_config[key] = value
	}
	result["archive_target_config"] = archive_target_config

	result["archive_frequency"] = VMPP.ArchiveFrequency.String()

	archive_schedule := make(xmlrpc.Struct)
	for key, value := range VMPP.ArchiveSchedule {
		archive_schedule[key] = value
	}
	result["archive_schedule"] = archive_schedule

	result["is_archive_running"] = VMPP.IsArchiveRunning

	result["archive_last_run_time"] = VMPP.ArchiveLastRunTime

	result["VMs"] = VMPP.VMs

	result["is_alarm_enabled"] = VMPP.IsAlarmEnabled

	alarm_config := make(xmlrpc.Struct)
	for key, value := range VMPP.AlarmConfig {
		alarm_config[key] = value
	}
	result["alarm_config"] = alarm_config

	result["recent_alerts"] = VMPP.RecentAlerts

	return result
}

func ToVMPP(obj interface{}) (resultObj *VMPP) {

	objValue := reflect.ValueOf(obj)
	resultObj = &VMPP{}

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
		case "is_policy_enabled":
			if v, ok := keyValue.(bool); ok {
				resultObj.IsPolicyEnabled = v
			}
		case "backup_type":
			if v, ok := keyValue.(VmppBackupType); ok {
				resultObj.BackupType = v
			}
		case "backup_retention_value":
			if v, ok := keyValue.(int); ok {
				resultObj.BackupRetentionValue = v
			}
		case "backup_frequency":
			if v, ok := keyValue.(VmppBackupFrequency); ok {
				resultObj.BackupFrequency = v
			}
		case "backup_schedule":

			resultObj.BackupSchedule = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.BackupSchedule[mapKeyName] = v
				} else {
					resultObj.BackupSchedule[mapKeyName] = ""
				}
			}
		case "is_backup_running":
			if v, ok := keyValue.(bool); ok {
				resultObj.IsBackupRunning = v
			}
		case "backup_last_run_time":
			if v, ok := keyValue.(time.Time); ok {
				resultObj.BackupLastRunTime = v
			}
		case "archive_target_type":
			if v, ok := keyValue.(VmppArchiveTargetType); ok {
				resultObj.ArchiveTargetType = v
			}
		case "archive_target_config":

			resultObj.ArchiveTargetConfig = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.ArchiveTargetConfig[mapKeyName] = v
				} else {
					resultObj.ArchiveTargetConfig[mapKeyName] = ""
				}
			}
		case "archive_frequency":
			if v, ok := keyValue.(VmppArchiveFrequency); ok {
				resultObj.ArchiveFrequency = v
			}
		case "archive_schedule":

			resultObj.ArchiveSchedule = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.ArchiveSchedule[mapKeyName] = v
				} else {
					resultObj.ArchiveSchedule[mapKeyName] = ""
				}
			}
		case "is_archive_running":
			if v, ok := keyValue.(bool); ok {
				resultObj.IsArchiveRunning = v
			}
		case "archive_last_run_time":
			if v, ok := keyValue.(time.Time); ok {
				resultObj.ArchiveLastRunTime = v
			}
		case "VMs":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.VMs = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.VMs[i] = v
					}
				}
			}
		case "is_alarm_enabled":
			if v, ok := keyValue.(bool); ok {
				resultObj.IsAlarmEnabled = v
			}
		case "alarm_config":

			resultObj.AlarmConfig = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.AlarmConfig[mapKeyName] = v
				} else {
					resultObj.AlarmConfig[mapKeyName] = ""
				}
			}
		case "recent_alerts":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.RecentAlerts = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.RecentAlerts[i] = v
					}
				}
			}
		}
	}

	return resultObj
}

/* GetAllRecords: Return a map of VMPP references to VMPP records for all VMPPs known to the system. */
func (client *XenClient) VMPPGetAllRecords() (result map[string]VMPP, err error) {
	obj, err := client.APICall("VMPP.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]VMPP{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToVMPP(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the VMPPs known to the system. */
func (client *XenClient) VMPPGetAll() (result []string, err error) {
	obj, err := client.APICall("VMPP.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* SetArchiveLastRunTime:  */
func (client *XenClient) VMPPSetArchiveLastRunTime(self string, value time.Time) (err error) {
	_, err = client.APICall("VMPP.set_archive_last_run_time", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetBackupLastRunTime:  */
func (client *XenClient) VMPPSetBackupLastRunTime(self string, value time.Time) (err error) {
	_, err = client.APICall("VMPP.set_backup_last_run_time", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveFromAlarmConfig:  */
func (client *XenClient) VMPPRemoveFromAlarmConfig(self string, key string) (err error) {
	_, err = client.APICall("VMPP.remove_from_alarm_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveFromArchiveSchedule:  */
func (client *XenClient) VMPPRemoveFromArchiveSchedule(self string, key string) (err error) {
	_, err = client.APICall("VMPP.remove_from_archive_schedule", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveFromArchiveTargetConfig:  */
func (client *XenClient) VMPPRemoveFromArchiveTargetConfig(self string, key string) (err error) {
	_, err = client.APICall("VMPP.remove_from_archive_target_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveFromBackupSchedule:  */
func (client *XenClient) VMPPRemoveFromBackupSchedule(self string, key string) (err error) {
	_, err = client.APICall("VMPP.remove_from_backup_schedule", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToAlarmConfig:  */
func (client *XenClient) VMPPAddToAlarmConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("VMPP.add_to_alarm_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToArchiveSchedule:  */
func (client *XenClient) VMPPAddToArchiveSchedule(self string, key string, value string) (err error) {
	_, err = client.APICall("VMPP.add_to_archive_schedule", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToArchiveTargetConfig:  */
func (client *XenClient) VMPPAddToArchiveTargetConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("VMPP.add_to_archive_target_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToBackupSchedule:  */
func (client *XenClient) VMPPAddToBackupSchedule(self string, key string, value string) (err error) {
	_, err = client.APICall("VMPP.add_to_backup_schedule", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetAlarmConfig:  */
func (client *XenClient) VMPPSetAlarmConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("VMPP.set_alarm_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetIsAlarmEnabled: Set the value of the is_alarm_enabled field */
func (client *XenClient) VMPPSetIsAlarmEnabled(self string, value bool) (err error) {
	_, err = client.APICall("VMPP.set_is_alarm_enabled", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetArchiveTargetConfig:  */
func (client *XenClient) VMPPSetArchiveTargetConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("VMPP.set_archive_target_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetArchiveTargetType: Set the value of the archive_target_config_type field */
func (client *XenClient) VMPPSetArchiveTargetType(self string, value VmppArchiveTargetType) (err error) {
	_, err = client.APICall("VMPP.set_archive_target_type", self, value.String())
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetArchiveSchedule:  */
func (client *XenClient) VMPPSetArchiveSchedule(self string, value map[string]string) (err error) {
	_, err = client.APICall("VMPP.set_archive_schedule", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetArchiveFrequency: Set the value of the archive_frequency field */
func (client *XenClient) VMPPSetArchiveFrequency(self string, value VmppArchiveFrequency) (err error) {
	_, err = client.APICall("VMPP.set_archive_frequency", self, value.String())
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetBackupSchedule:  */
func (client *XenClient) VMPPSetBackupSchedule(self string, value map[string]string) (err error) {
	_, err = client.APICall("VMPP.set_backup_schedule", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetBackupFrequency: Set the value of the backup_frequency field */
func (client *XenClient) VMPPSetBackupFrequency(self string, value VmppBackupFrequency) (err error) {
	_, err = client.APICall("VMPP.set_backup_frequency", self, value.String())
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetBackupRetentionValue:  */
func (client *XenClient) VMPPSetBackupRetentionValue(self string, value int) (err error) {
	_, err = client.APICall("VMPP.set_backup_retention_value", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetAlerts: This call fetches a history of alerts for a given protection policy */
func (client *XenClient) VMPPGetAlerts(vmpp string, hours_from_now int) (result []string, err error) {
	obj, err := client.APICall("VMPP.get_alerts", vmpp, hours_from_now)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* ArchiveNow: This call archives the snapshot provided as a parameter */
func (client *XenClient) VMPPArchiveNow(snapshot string) (result string, err error) {
	obj, err := client.APICall("VMPP.archive_now", snapshot)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* ProtectNow: This call executes the protection policy immediately */
func (client *XenClient) VMPPProtectNow(vmpp string) (result string, err error) {
	obj, err := client.APICall("VMPP.protect_now", vmpp)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* SetBackupType: Set the backup_type field of the given VMPP. */
func (client *XenClient) VMPPSetBackupType(self string, value VmppBackupType) (err error) {
	_, err = client.APICall("VMPP.set_backup_type", self, value.String())
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetIsPolicyEnabled: Set the is_policy_enabled field of the given VMPP. */
func (client *XenClient) VMPPSetIsPolicyEnabled(self string, value bool) (err error) {
	_, err = client.APICall("VMPP.set_is_policy_enabled", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetNameDescription: Set the name/description field of the given VMPP. */
func (client *XenClient) VMPPSetNameDescription(self string, value string) (err error) {
	_, err = client.APICall("VMPP.set_name_description", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetNameLabel: Set the name/label field of the given VMPP. */
func (client *XenClient) VMPPSetNameLabel(self string, value string) (err error) {
	_, err = client.APICall("VMPP.set_name_label", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetRecentAlerts: Get the recent_alerts field of the given VMPP. */
func (client *XenClient) VMPPGetRecentAlerts(self string) (result []string, err error) {
	obj, err := client.APICall("VMPP.get_recent_alerts", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetAlarmConfig: Get the alarm_config field of the given VMPP. */
func (client *XenClient) VMPPGetAlarmConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VMPP.get_alarm_config", self)
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

/* GetIsAlarmEnabled: Get the is_alarm_enabled field of the given VMPP. */
func (client *XenClient) VMPPGetIsAlarmEnabled(self string) (result bool, err error) {
	obj, err := client.APICall("VMPP.get_is_alarm_enabled", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetVMs: Get the VMs field of the given VMPP. */
func (client *XenClient) VMPPGetVMs(self string) (result []string, err error) {
	obj, err := client.APICall("VMPP.get_VMs", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetArchiveLastRunTime: Get the archive_last_run_time field of the given VMPP. */
func (client *XenClient) VMPPGetArchiveLastRunTime(self string) (result time.Time, err error) {
	obj, err := client.APICall("VMPP.get_archive_last_run_time", self)
	if err != nil {
		return
	}
	result = obj.(time.Time)
	return
}

/* GetIsArchiveRunning: Get the is_archive_running field of the given VMPP. */
func (client *XenClient) VMPPGetIsArchiveRunning(self string) (result bool, err error) {
	obj, err := client.APICall("VMPP.get_is_archive_running", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetArchiveSchedule: Get the archive_schedule field of the given VMPP. */
func (client *XenClient) VMPPGetArchiveSchedule(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VMPP.get_archive_schedule", self)
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

/* GetArchiveFrequency: Get the archive_frequency field of the given VMPP. */
func (client *XenClient) VMPPGetArchiveFrequency(self string) (result VmppArchiveFrequency, err error) {
	obj, err := client.APICall("VMPP.get_archive_frequency", self)
	if err != nil {
		return
	}

	result = ToVmppArchiveFrequency(obj.(string))

	return
}

/* GetArchiveTargetConfig: Get the archive_target_config field of the given VMPP. */
func (client *XenClient) VMPPGetArchiveTargetConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VMPP.get_archive_target_config", self)
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

/* GetArchiveTargetType: Get the archive_target_type field of the given VMPP. */
func (client *XenClient) VMPPGetArchiveTargetType(self string) (result VmppArchiveTargetType, err error) {
	obj, err := client.APICall("VMPP.get_archive_target_type", self)
	if err != nil {
		return
	}

	result = ToVmppArchiveTargetType(obj.(string))

	return
}

/* GetBackupLastRunTime: Get the backup_last_run_time field of the given VMPP. */
func (client *XenClient) VMPPGetBackupLastRunTime(self string) (result time.Time, err error) {
	obj, err := client.APICall("VMPP.get_backup_last_run_time", self)
	if err != nil {
		return
	}
	result = obj.(time.Time)
	return
}

/* GetIsBackupRunning: Get the is_backup_running field of the given VMPP. */
func (client *XenClient) VMPPGetIsBackupRunning(self string) (result bool, err error) {
	obj, err := client.APICall("VMPP.get_is_backup_running", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetBackupSchedule: Get the backup_schedule field of the given VMPP. */
func (client *XenClient) VMPPGetBackupSchedule(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VMPP.get_backup_schedule", self)
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

/* GetBackupFrequency: Get the backup_frequency field of the given VMPP. */
func (client *XenClient) VMPPGetBackupFrequency(self string) (result VmppBackupFrequency, err error) {
	obj, err := client.APICall("VMPP.get_backup_frequency", self)
	if err != nil {
		return
	}

	result = ToVmppBackupFrequency(obj.(string))

	return
}

/* GetBackupRetentionValue: Get the backup_retention_value field of the given VMPP. */
func (client *XenClient) VMPPGetBackupRetentionValue(self string) (result int, err error) {
	obj, err := client.APICall("VMPP.get_backup_retention_value", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetBackupType: Get the backup_type field of the given VMPP. */
func (client *XenClient) VMPPGetBackupType(self string) (result VmppBackupType, err error) {
	obj, err := client.APICall("VMPP.get_backup_type", self)
	if err != nil {
		return
	}

	result = ToVmppBackupType(obj.(string))

	return
}

/* GetIsPolicyEnabled: Get the is_policy_enabled field of the given VMPP. */
func (client *XenClient) VMPPGetIsPolicyEnabled(self string) (result bool, err error) {
	obj, err := client.APICall("VMPP.get_is_policy_enabled", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetNameDescription: Get the name/description field of the given VMPP. */
func (client *XenClient) VMPPGetNameDescription(self string) (result string, err error) {
	obj, err := client.APICall("VMPP.get_name_description", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetNameLabel: Get the name/label field of the given VMPP. */
func (client *XenClient) VMPPGetNameLabel(self string) (result string, err error) {
	obj, err := client.APICall("VMPP.get_name_label", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given VMPP. */
func (client *XenClient) VMPPGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("VMPP.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByNameLabel: Get all the VMPP instances with the given label. */
func (client *XenClient) VMPPGetByNameLabel(label string) (result []string, err error) {
	obj, err := client.APICall("VMPP.get_by_name_label", label)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* Destroy: Destroy the specified VMPP instance. */
func (client *XenClient) VMPPDestroy(self string) (err error) {
	_, err = client.APICall("VMPP.destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Create: Create a new VMPP instance, and return its handle.
The constructor args are: name_label, name_description, is_policy_enabled, backup_type, backup_retention_value, backup_frequency, backup_schedule, archive_target_type, archive_target_config, archive_frequency, archive_schedule, is_alarm_enabled, alarm_config (* = non-optional). */
func (client *XenClient) VMPPCreate(args VMPP) (result string, err error) {
	obj, err := client.APICall("VMPP.create", FromVMPPToXml(&args))
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the VMPP instance with the specified UUID. */
func (client *XenClient) VMPPGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("VMPP.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given VMPP. */
func (client *XenClient) VMPPGetRecord(self string) (result VMPP, err error) {
	obj, err := client.APICall("VMPP.get_record", self)
	if err != nil {
		return
	}

	result = *ToVMPP(obj.(interface{}))

	return
}

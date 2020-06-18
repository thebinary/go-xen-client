// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w Pool.gen.go

package go_xen_client

import (
	"reflect"
	"strconv"

	"github.com/nilshell/xmlrpc"
)

//Pool: Pool-wide information
type Pool struct {
	Uuid                     string                           // Unique identifier/object reference
	NameLabel                string                           // Short name
	NameDescription          string                           // Description
	Master                   string                           // The host that is pool master
	DefaultSR                string                           // Default SR for VDIs
	SuspendImageSR           string                           // The SR in which VDIs for suspend images are created
	CrashDumpSR              string                           // The SR in which VDIs for crash dumps are created
	OtherConfig              map[string]string                // additional configuration
	HaEnabled                bool                             // true if HA is enabled on the pool, false otherwise
	HaConfiguration          map[string]string                // The current HA configuration
	HaStatefiles             []string                         // HA statefile VDIs in use
	HaHostFailuresToTolerate int                              // Number of host failures to tolerate before the Pool is declared to be overcommitted
	HaPlanExistsFor          int                              // Number of future host failures we have managed to find a plan for. Once this reaches zero any future host failures will cause the failure of protected VMs.
	HaAllowOvercommit        bool                             // If set to false then operations which would cause the Pool to become overcommitted will be blocked.
	HaOvercommitted          bool                             // True if the Pool is considered to be overcommitted i.e. if there exist insufficient physical resources to tolerate the configured number of host failures
	Blobs                    map[string]string                // Binary blobs associated with this pool
	Tags                     []string                         // user-specified tags for categorization purposes
	GuiConfig                map[string]string                // gui-specific configuration for pool
	HealthCheckConfig        map[string]string                // Configuration for the automatic health check feature
	WlbUrl                   string                           // Url for the configured workload balancing host
	WlbUsername              string                           // Username for accessing the workload balancing host
	WlbEnabled               bool                             // true if workload balancing is enabled on the pool, false otherwise
	WlbVerifyCert            bool                             // true if communication with the WLB server should enforce SSL certificate verification.
	RedoLogEnabled           bool                             // true a redo-log is to be used other than when HA is enabled, false otherwise
	RedoLogVdi               string                           // indicates the VDI to use for the redo-log other than when HA is enabled
	VswitchController        string                           // address of the vswitch controller
	Restrictions             map[string]string                // Pool-wide restrictions currently in effect
	MetadataVDIs             []string                         // The set of currently known metadata VDIs for this pool
	HaClusterStack           string                           // The HA cluster stack that is currently in use. Only valid when HA is enabled.
	AllowedOperations        []PoolAllowedOperations          // list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	CurrentOperations        map[string]PoolAllowedOperations // links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	GuestAgentConfig         map[string]string                // Pool-wide guest agent configuration information
	CpuInfo                  map[string]string                // Details about the physical CPUs on the pool
	PolicyNoVendorDevice     bool                             // The pool-wide policy for clients on whether to use the vendor device or not on newly created VMs. This field will also be consulted if the 'has_vendor_device' field is not specified in the VM.create call.
	LivePatchingDisabled     bool                             // The pool-wide flag to show if the live patching feauture is disabled or not.
	IgmpSnoopingEnabled      bool                             // true if IGMP snooping is enabled in the pool, false otherwise.
	UefiCertificates         string                           // The UEFI certificates allowing Secure Boot

}

func FromPoolToXml(pool *Pool) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] =

		pool.Uuid

	result["name_label"] =

		pool.NameLabel

	result["name_description"] =

		pool.NameDescription

	result["master"] =

		pool.Master

	result["default_SR"] =

		pool.DefaultSR

	result["suspend_image_SR"] =

		pool.SuspendImageSR

	result["crash_dump_SR"] =

		pool.CrashDumpSR

	other_config := make(xmlrpc.Struct)
	for key, value := range pool.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	result["ha_enabled"] =

		pool.HaEnabled

	ha_configuration := make(xmlrpc.Struct)
	for key, value := range pool.HaConfiguration {
		ha_configuration[key] = value
	}
	result["ha_configuration"] = ha_configuration

	result["ha_statefiles"] =

		pool.HaStatefiles

	result["ha_host_failures_to_tolerate"] =

		strconv.Itoa(pool.HaHostFailuresToTolerate)

	result["ha_plan_exists_for"] =

		strconv.Itoa(pool.HaPlanExistsFor)

	result["ha_allow_overcommit"] =

		pool.HaAllowOvercommit

	result["ha_overcommitted"] =

		pool.HaOvercommitted

	blobs := make(xmlrpc.Struct)
	for key, value := range pool.Blobs {
		blobs[key] = value
	}
	result["blobs"] = blobs

	result["tags"] =

		pool.Tags

	gui_config := make(xmlrpc.Struct)
	for key, value := range pool.GuiConfig {
		gui_config[key] = value
	}
	result["gui_config"] = gui_config

	health_check_config := make(xmlrpc.Struct)
	for key, value := range pool.HealthCheckConfig {
		health_check_config[key] = value
	}
	result["health_check_config"] = health_check_config

	result["wlb_url"] =

		pool.WlbUrl

	result["wlb_username"] =

		pool.WlbUsername

	result["wlb_enabled"] =

		pool.WlbEnabled

	result["wlb_verify_cert"] =

		pool.WlbVerifyCert

	result["redo_log_enabled"] =

		pool.RedoLogEnabled

	result["redo_log_vdi"] =

		pool.RedoLogVdi

	result["vswitch_controller"] =

		pool.VswitchController

	restrictions := make(xmlrpc.Struct)
	for key, value := range pool.Restrictions {
		restrictions[key] = value
	}
	result["restrictions"] = restrictions

	result["metadata_VDIs"] =

		pool.MetadataVDIs

	result["ha_cluster_stack"] =

		pool.HaClusterStack

	result["allowed_operations"] =

		pool.AllowedOperations

	current_operations := make(xmlrpc.Struct)
	for key, value := range pool.CurrentOperations {
		current_operations[key] = value
	}
	result["current_operations"] = current_operations

	guest_agent_config := make(xmlrpc.Struct)
	for key, value := range pool.GuestAgentConfig {
		guest_agent_config[key] = value
	}
	result["guest_agent_config"] = guest_agent_config

	cpu_info := make(xmlrpc.Struct)
	for key, value := range pool.CpuInfo {
		cpu_info[key] = value
	}
	result["cpu_info"] = cpu_info

	result["policy_no_vendor_device"] =

		pool.PolicyNoVendorDevice

	result["live_patching_disabled"] =

		pool.LivePatchingDisabled

	result["igmp_snooping_enabled"] =

		pool.IgmpSnoopingEnabled

	result["uefi_certificates"] =

		pool.UefiCertificates

	return result
}

func ToPool(obj interface{}) (resultObj *Pool) {

	objValue := reflect.ValueOf(obj)
	resultObj = &Pool{}

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
		case "master":
			if v, ok := keyValue.(string); ok {
				resultObj.Master = v
			}
		case "default_SR":
			if v, ok := keyValue.(string); ok {
				resultObj.DefaultSR = v
			}
		case "suspend_image_SR":
			if v, ok := keyValue.(string); ok {
				resultObj.SuspendImageSR = v
			}
		case "crash_dump_SR":
			if v, ok := keyValue.(string); ok {
				resultObj.CrashDumpSR = v
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
		case "ha_enabled":
			if v, ok := keyValue.(bool); ok {
				resultObj.HaEnabled = v
			}
		case "ha_configuration":

			resultObj.HaConfiguration = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.HaConfiguration[mapKeyName] = v
				} else {
					resultObj.HaConfiguration[mapKeyName] = ""
				}
			}
		case "ha_statefiles":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.HaStatefiles = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.HaStatefiles[i] = v
					}
				}
			}
		case "ha_host_failures_to_tolerate":
			if v, ok := keyValue.(int); ok {
				resultObj.HaHostFailuresToTolerate = v
			}
		case "ha_plan_exists_for":
			if v, ok := keyValue.(int); ok {
				resultObj.HaPlanExistsFor = v
			}
		case "ha_allow_overcommit":
			if v, ok := keyValue.(bool); ok {
				resultObj.HaAllowOvercommit = v
			}
		case "ha_overcommitted":
			if v, ok := keyValue.(bool); ok {
				resultObj.HaOvercommitted = v
			}
		case "blobs":

			resultObj.Blobs = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.Blobs[mapKeyName] = v
				} else {
					resultObj.Blobs[mapKeyName] = ""
				}
			}
		case "tags":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.Tags = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.Tags[i] = v
					}
				}
			}
		case "gui_config":

			resultObj.GuiConfig = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.GuiConfig[mapKeyName] = v
				} else {
					resultObj.GuiConfig[mapKeyName] = ""
				}
			}
		case "health_check_config":

			resultObj.HealthCheckConfig = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.HealthCheckConfig[mapKeyName] = v
				} else {
					resultObj.HealthCheckConfig[mapKeyName] = ""
				}
			}
		case "wlb_url":
			if v, ok := keyValue.(string); ok {
				resultObj.WlbUrl = v
			}
		case "wlb_username":
			if v, ok := keyValue.(string); ok {
				resultObj.WlbUsername = v
			}
		case "wlb_enabled":
			if v, ok := keyValue.(bool); ok {
				resultObj.WlbEnabled = v
			}
		case "wlb_verify_cert":
			if v, ok := keyValue.(bool); ok {
				resultObj.WlbVerifyCert = v
			}
		case "redo_log_enabled":
			if v, ok := keyValue.(bool); ok {
				resultObj.RedoLogEnabled = v
			}
		case "redo_log_vdi":
			if v, ok := keyValue.(string); ok {
				resultObj.RedoLogVdi = v
			}
		case "vswitch_controller":
			if v, ok := keyValue.(string); ok {
				resultObj.VswitchController = v
			}
		case "restrictions":

			resultObj.Restrictions = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.Restrictions[mapKeyName] = v
				} else {
					resultObj.Restrictions[mapKeyName] = ""
				}
			}
		case "metadata_VDIs":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.MetadataVDIs = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.MetadataVDIs[i] = v
					}
				}
			}
		case "ha_cluster_stack":
			if v, ok := keyValue.(string); ok {
				resultObj.HaClusterStack = v
			}
		case "allowed_operations":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.AllowedOperations = make([]PoolAllowedOperations, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(PoolAllowedOperations); ok {
						resultObj.AllowedOperations[i] = v
					}
				}
			}
		case "current_operations":

			resultObj.CurrentOperations = map[string]PoolAllowedOperations{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.CurrentOperations[mapKeyName] = ToPoolAllowedOperations(v)
				} else {
					resultObj.CurrentOperations[mapKeyName] = 0
				}
			}
		case "guest_agent_config":

			resultObj.GuestAgentConfig = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.GuestAgentConfig[mapKeyName] = v
				} else {
					resultObj.GuestAgentConfig[mapKeyName] = ""
				}
			}
		case "cpu_info":

			resultObj.CpuInfo = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.CpuInfo[mapKeyName] = v
				} else {
					resultObj.CpuInfo[mapKeyName] = ""
				}
			}
		case "policy_no_vendor_device":
			if v, ok := keyValue.(bool); ok {
				resultObj.PolicyNoVendorDevice = v
			}
		case "live_patching_disabled":
			if v, ok := keyValue.(bool); ok {
				resultObj.LivePatchingDisabled = v
			}
		case "igmp_snooping_enabled":
			if v, ok := keyValue.(bool); ok {
				resultObj.IgmpSnoopingEnabled = v
			}
		case "uefi_certificates":
			if v, ok := keyValue.(string); ok {
				resultObj.UefiCertificates = v
			}
		}
	}

	return resultObj
}

/* GetAllRecords: Return a map of pool references to pool records for all pools known to the system. */
func (client *XenClient) PoolGetAllRecords() (result map[string]Pool, err error) {
	obj, err := client.APICall("pool.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]Pool{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToPool(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the pools known to the system. */
func (client *XenClient) PoolGetAll() (result []string, err error) {
	obj, err := client.APICall("pool.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* RemoveFromGuestAgentConfig: Remove a key-value pair from the pool-wide guest agent configuration */
func (client *XenClient) PoolRemoveFromGuestAgentConfig(self string, key string) (err error) {
	_, err = client.APICall("pool.remove_from_guest_agent_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToGuestAgentConfig: Add a key-value pair to the pool-wide guest agent configuration */
func (client *XenClient) PoolAddToGuestAgentConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("pool.add_to_guest_agent_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* HasExtension: Return true if the extension is available on the pool */
func (client *XenClient) PoolHasExtension(self string, name string) (result bool, err error) {
	obj, err := client.APICall("pool.has_extension", self, name)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* SetIgmpSnoopingEnabled: Enable or disable IGMP Snooping on the pool. */
func (client *XenClient) PoolSetIgmpSnoopingEnabled(self string, value bool) (err error) {
	_, err = client.APICall("pool.set_igmp_snooping_enabled", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* DisableSslLegacy: Sets ssl_legacy true on each host, pool-master last. See Host.ssl_legacy and Host.set_ssl_legacy. */
func (client *XenClient) PoolDisableSslLegacy(self string) (err error) {
	_, err = client.APICall("pool.disable_ssl_legacy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* EnableSslLegacy: Sets ssl_legacy true on each host, pool-master last. See Host.ssl_legacy and Host.set_ssl_legacy. */
func (client *XenClient) PoolEnableSslLegacy(self string) (err error) {
	_, err = client.APICall("pool.enable_ssl_legacy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* ApplyEdition: Apply an edition to all hosts in the pool */
func (client *XenClient) PoolApplyEdition(self string, edition string) (err error) {
	_, err = client.APICall("pool.apply_edition", self, edition)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetLicenseState: This call returns the license state for the pool */
func (client *XenClient) PoolGetLicenseState(self string) (result map[string]string, err error) {
	obj, err := client.APICall("pool.get_license_state", self)
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

/* DisableLocalStorageCaching: This call disables pool-wide local storage caching */
func (client *XenClient) PoolDisableLocalStorageCaching(self string) (err error) {
	_, err = client.APICall("pool.disable_local_storage_caching", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* EnableLocalStorageCaching: This call attempts to enable pool-wide local storage caching */
func (client *XenClient) PoolEnableLocalStorageCaching(self string) (err error) {
	_, err = client.APICall("pool.enable_local_storage_caching", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* TestArchiveTarget: This call tests if a location is valid */
func (client *XenClient) PoolTestArchiveTarget(self string, config map[string]string) (result string, err error) {
	obj, err := client.APICall("pool.test_archive_target", self, config)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* SetVswitchController: Set the IP address of the vswitch controller. */
func (client *XenClient) PoolSetVswitchController(address string) (err error) {
	_, err = client.APICall("pool.set_vswitch_controller", address)
	if err != nil {
		return
	}
	// no return result
	return
}

/* DisableRedoLog: Disable the redo log if in use, unless HA is enabled. */
func (client *XenClient) PoolDisableRedoLog() (err error) {
	_, err = client.APICall("pool.disable_redo_log")
	if err != nil {
		return
	}
	// no return result
	return
}

/* EnableRedoLog: Enable the redo log on the given SR and start using it, unless HA is enabled. */
func (client *XenClient) PoolEnableRedoLog(sr string) (err error) {
	_, err = client.APICall("pool.enable_redo_log", sr)
	if err != nil {
		return
	}
	// no return result
	return
}

/* CertificateSync: Sync SSL certificates from master to slaves. */
func (client *XenClient) PoolCertificateSync() (err error) {
	_, err = client.APICall("pool.certificate_sync")
	if err != nil {
		return
	}
	// no return result
	return
}

/* CrlList: List all installed SSL certificate revocation lists. */
func (client *XenClient) PoolCrlList() (result []string, err error) {
	obj, err := client.APICall("pool.crl_list")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* CrlUninstall: Remove an SSL certificate revocation list. */
func (client *XenClient) PoolCrlUninstall(name string) (err error) {
	_, err = client.APICall("pool.crl_uninstall", name)
	if err != nil {
		return
	}
	// no return result
	return
}

/* CrlInstall: Install an SSL certificate revocation list, pool-wide. */
func (client *XenClient) PoolCrlInstall(name string, cert string) (err error) {
	_, err = client.APICall("pool.crl_install", name, cert)
	if err != nil {
		return
	}
	// no return result
	return
}

/* CertificateList: List all installed SSL certificates. */
func (client *XenClient) PoolCertificateList() (result []string, err error) {
	obj, err := client.APICall("pool.certificate_list")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* CertificateUninstall: Remove an SSL certificate. */
func (client *XenClient) PoolCertificateUninstall(name string) (err error) {
	_, err = client.APICall("pool.certificate_uninstall", name)
	if err != nil {
		return
	}
	// no return result
	return
}

/* CertificateInstall: Install an SSL certificate pool-wide. */
func (client *XenClient) PoolCertificateInstall(name string, cert string) (err error) {
	_, err = client.APICall("pool.certificate_install", name, cert)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SendTestPost: Send the given body to the given host and port, using HTTPS, and print the response.  This is used for debugging the SSL layer. */
func (client *XenClient) PoolSendTestPost(host string, port int, body string) (result string, err error) {
	obj, err := client.APICall("pool.send_test_post", host, port, body)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* RetrieveWlbRecommendations: Retrieves vm migrate recommendations for the pool from the workload balancing server */
func (client *XenClient) PoolRetrieveWlbRecommendations() (result map[string][]string, err error) {
	obj, err := client.APICall("pool.retrieve_wlb_recommendations")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string][]string{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		interimObj := obj.Interface().([]interface{})
		interimResult := make([]string, len(interimObj))
		for i, interimValue := range interimObj {
			interimResult[i] = interimValue.(string)
		}
		result[key.String()] = interimResult
	}

	return
}

/* RetrieveWlbConfiguration: Retrieves the pool optimization criteria from the workload balancing server */
func (client *XenClient) PoolRetrieveWlbConfiguration() (result map[string]string, err error) {
	obj, err := client.APICall("pool.retrieve_wlb_configuration")
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

/* SendWlbConfiguration: Sets the pool optimization criteria for the workload balancing server */
func (client *XenClient) PoolSendWlbConfiguration(config map[string]string) (err error) {
	_, err = client.APICall("pool.send_wlb_configuration", config)
	if err != nil {
		return
	}
	// no return result
	return
}

/* DeconfigureWlb: Permanently deconfigures workload balancing monitoring on this pool */
func (client *XenClient) PoolDeconfigureWlb() (err error) {
	_, err = client.APICall("pool.deconfigure_wlb")
	if err != nil {
		return
	}
	// no return result
	return
}

/* InitializeWlb: Initializes workload balancing monitoring on this pool with the specified wlb server */
func (client *XenClient) PoolInitializeWlb(wlb_url string, wlb_username string, wlb_password string, xenserver_username string, xenserver_password string) (err error) {
	_, err = client.APICall("pool.initialize_wlb", wlb_url, wlb_username, wlb_password, xenserver_username, xenserver_password)
	if err != nil {
		return
	}
	// no return result
	return
}

/* DetectNonhomogeneousExternalAuth: This call asynchronously detects if the external authentication configuration in any slave is different from that in the master and raises appropriate alerts */
func (client *XenClient) PoolDetectNonhomogeneousExternalAuth(pool string) (err error) {
	_, err = client.APICall("pool.detect_nonhomogeneous_external_auth", pool)
	if err != nil {
		return
	}
	// no return result
	return
}

/* DisableExternalAuth: This call disables external authentication on all the hosts of the pool */
func (client *XenClient) PoolDisableExternalAuth(pool string, config map[string]string) (err error) {
	_, err = client.APICall("pool.disable_external_auth", pool, config)
	if err != nil {
		return
	}
	// no return result
	return
}

/* EnableExternalAuth: This call enables external authentication on all the hosts of the pool */
func (client *XenClient) PoolEnableExternalAuth(pool string, config map[string]string, service_name string, auth_type string) (err error) {
	_, err = client.APICall("pool.enable_external_auth", pool, config, service_name, auth_type)
	if err != nil {
		return
	}
	// no return result
	return
}

/* CreateNewBlob: Create a placeholder for a named binary blob of data that is associated with this pool */
func (client *XenClient) PoolCreateNewBlob(pool string, name string, mime_type string, public bool) (result string, err error) {
	obj, err := client.APICall("pool.create_new_blob", pool, name, mime_type, public)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* SetHaHostFailuresToTolerate: Set the maximum number of host failures to consider in the HA VM restart planner */
func (client *XenClient) PoolSetHaHostFailuresToTolerate(self string, value int) (err error) {
	_, err = client.APICall("pool.set_ha_host_failures_to_tolerate", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* HaComputeVmFailoverPlan: Return a VM failover plan assuming a given subset of hosts fail */
func (client *XenClient) PoolHaComputeVmFailoverPlan(failed_hosts []string, failed_vms []string) (result map[string]string, err error) {
	obj, err := client.APICall("pool.ha_compute_vm_failover_plan", failed_hosts, failed_vms)
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

/* HaComputeHypotheticalMaxHostFailuresToTolerate: Returns the maximum number of host failures we could tolerate before we would be unable to restart the provided VMs */
func (client *XenClient) PoolHaComputeHypotheticalMaxHostFailuresToTolerate(configuration map[string]string) (result int, err error) {
	obj, err := client.APICall("pool.ha_compute_hypothetical_max_host_failures_to_tolerate", configuration)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* HaComputeMaxHostFailuresToTolerate: Returns the maximum number of host failures we could tolerate before we would be unable to restart configured VMs */
func (client *XenClient) PoolHaComputeMaxHostFailuresToTolerate() (result int, err error) {
	obj, err := client.APICall("pool.ha_compute_max_host_failures_to_tolerate")
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* HaFailoverPlanExists: Returns true if a VM failover plan exists for up to 'n' host failures */
func (client *XenClient) PoolHaFailoverPlanExists(n int) (result bool, err error) {
	obj, err := client.APICall("pool.ha_failover_plan_exists", n)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* HaPreventRestartsFor: When this call returns the VM restart logic will not run for the requested number of seconds. If the argument is zero then the restart thread is immediately unblocked */
func (client *XenClient) PoolHaPreventRestartsFor(seconds int) (err error) {
	_, err = client.APICall("pool.ha_prevent_restarts_for", seconds)
	if err != nil {
		return
	}
	// no return result
	return
}

/* DesignateNewMaster: Perform an orderly handover of the role of master to the referenced host. */
func (client *XenClient) PoolDesignateNewMaster(host string) (err error) {
	_, err = client.APICall("pool.designate_new_master", host)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SyncDatabase: Forcibly synchronise the database now */
func (client *XenClient) PoolSyncDatabase() (err error) {
	_, err = client.APICall("pool.sync_database")
	if err != nil {
		return
	}
	// no return result
	return
}

/* DisableHa: Turn off High Availability mode */
func (client *XenClient) PoolDisableHa() (err error) {
	_, err = client.APICall("pool.disable_ha")
	if err != nil {
		return
	}
	// no return result
	return
}

/* EnableHa: Turn on High Availability mode */
func (client *XenClient) PoolEnableHa(heartbeat_srs []string, configuration map[string]string) (err error) {
	_, err = client.APICall("pool.enable_ha", heartbeat_srs, configuration)
	if err != nil {
		return
	}
	// no return result
	return
}

/* CreateVLANFromPIF: Create a pool-wide VLAN by taking the PIF. */
func (client *XenClient) PoolCreateVLANFromPIF(pif string, network string, VLAN int) (result []string, err error) {
	obj, err := client.APICall("pool.create_VLAN_from_PIF", pif, network, VLAN)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* ManagementReconfigure: Reconfigure the management network interface for all Hosts in the Pool */
func (client *XenClient) PoolManagementReconfigure(network string) (err error) {
	_, err = client.APICall("pool.management_reconfigure", network)
	if err != nil {
		return
	}
	// no return result
	return
}

/* CreateVLAN: Create PIFs, mapping a network to the same physical interface/VLAN on each host. This call is deprecated: use Pool.create_VLAN_from_PIF instead. */
func (client *XenClient) PoolCreateVLAN(device string, network string, VLAN int) (result []string, err error) {
	obj, err := client.APICall("pool.create_VLAN", device, network, VLAN)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* RecoverSlaves: Instruct a pool master, M, to try and contact its slaves and, if slaves are in emergency mode, reset their master address to M. */
func (client *XenClient) PoolRecoverSlaves() (result []string, err error) {
	obj, err := client.APICall("pool.recover_slaves")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* EmergencyResetMaster: Instruct a slave already in a pool that the master has changed */
func (client *XenClient) PoolEmergencyResetMaster(master_address string) (err error) {
	_, err = client.APICall("pool.emergency_reset_master", master_address)
	if err != nil {
		return
	}
	// no return result
	return
}

/* EmergencyTransitionToMaster: Instruct host that's currently a slave to transition to being master */
func (client *XenClient) PoolEmergencyTransitionToMaster() (err error) {
	_, err = client.APICall("pool.emergency_transition_to_master")
	if err != nil {
		return
	}
	// no return result
	return
}

/* Eject: Instruct a pool master to eject a host from the pool */
func (client *XenClient) PoolEject(host string) (err error) {
	_, err = client.APICall("pool.eject", host)
	if err != nil {
		return
	}
	// no return result
	return
}

/* JoinForce: Instruct host to join a new pool */
func (client *XenClient) PoolJoinForce(master_address string, master_username string, master_password string) (err error) {
	_, err = client.APICall("pool.join_force", master_address, master_username, master_password)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Join: Instruct host to join a new pool */
func (client *XenClient) PoolJoin(master_address string, master_username string, master_password string) (err error) {
	_, err = client.APICall("pool.join", master_address, master_username, master_password)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetUefiCertificates: Set the uefi_certificates field of the given pool. */
func (client *XenClient) PoolSetUefiCertificates(self string, value string) (err error) {
	_, err = client.APICall("pool.set_uefi_certificates", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetLivePatchingDisabled: Set the live_patching_disabled field of the given pool. */
func (client *XenClient) PoolSetLivePatchingDisabled(self string, value bool) (err error) {
	_, err = client.APICall("pool.set_live_patching_disabled", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetPolicyNoVendorDevice: Set the policy_no_vendor_device field of the given pool. */
func (client *XenClient) PoolSetPolicyNoVendorDevice(self string, value bool) (err error) {
	_, err = client.APICall("pool.set_policy_no_vendor_device", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetWlbVerifyCert: Set the wlb_verify_cert field of the given pool. */
func (client *XenClient) PoolSetWlbVerifyCert(self string, value bool) (err error) {
	_, err = client.APICall("pool.set_wlb_verify_cert", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetWlbEnabled: Set the wlb_enabled field of the given pool. */
func (client *XenClient) PoolSetWlbEnabled(self string, value bool) (err error) {
	_, err = client.APICall("pool.set_wlb_enabled", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveFromHealthCheckConfig: Remove the given key and its corresponding value from the health_check_config field of the given pool.  If the key is not in that Map, then do nothing. */
func (client *XenClient) PoolRemoveFromHealthCheckConfig(self string, key string) (err error) {
	_, err = client.APICall("pool.remove_from_health_check_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToHealthCheckConfig: Add the given key-value pair to the health_check_config field of the given pool. */
func (client *XenClient) PoolAddToHealthCheckConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("pool.add_to_health_check_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetHealthCheckConfig: Set the health_check_config field of the given pool. */
func (client *XenClient) PoolSetHealthCheckConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("pool.set_health_check_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveFromGuiConfig: Remove the given key and its corresponding value from the gui_config field of the given pool.  If the key is not in that Map, then do nothing. */
func (client *XenClient) PoolRemoveFromGuiConfig(self string, key string) (err error) {
	_, err = client.APICall("pool.remove_from_gui_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToGuiConfig: Add the given key-value pair to the gui_config field of the given pool. */
func (client *XenClient) PoolAddToGuiConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("pool.add_to_gui_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetGuiConfig: Set the gui_config field of the given pool. */
func (client *XenClient) PoolSetGuiConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("pool.set_gui_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveTags: Remove the given value from the tags field of the given pool.  If the value is not in that Set, then do nothing. */
func (client *XenClient) PoolRemoveTags(self string, value string) (err error) {
	_, err = client.APICall("pool.remove_tags", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddTags: Add the given value to the tags field of the given pool.  If the value is already in that Set, then do nothing. */
func (client *XenClient) PoolAddTags(self string, value string) (err error) {
	_, err = client.APICall("pool.add_tags", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetTags: Set the tags field of the given pool. */
func (client *XenClient) PoolSetTags(self string, value []string) (err error) {
	_, err = client.APICall("pool.set_tags", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetHaAllowOvercommit: Set the ha_allow_overcommit field of the given pool. */
func (client *XenClient) PoolSetHaAllowOvercommit(self string, value bool) (err error) {
	_, err = client.APICall("pool.set_ha_allow_overcommit", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given pool.  If the key is not in that Map, then do nothing. */
func (client *XenClient) PoolRemoveFromOtherConfig(self string, key string) (err error) {
	_, err = client.APICall("pool.remove_from_other_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToOtherConfig: Add the given key-value pair to the other_config field of the given pool. */
func (client *XenClient) PoolAddToOtherConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("pool.add_to_other_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOtherConfig: Set the other_config field of the given pool. */
func (client *XenClient) PoolSetOtherConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("pool.set_other_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetCrashDumpSR: Set the crash_dump_SR field of the given pool. */
func (client *XenClient) PoolSetCrashDumpSR(self string, value string) (err error) {
	_, err = client.APICall("pool.set_crash_dump_SR", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetSuspendImageSR: Set the suspend_image_SR field of the given pool. */
func (client *XenClient) PoolSetSuspendImageSR(self string, value string) (err error) {
	_, err = client.APICall("pool.set_suspend_image_SR", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetDefaultSR: Set the default_SR field of the given pool. */
func (client *XenClient) PoolSetDefaultSR(self string, value string) (err error) {
	_, err = client.APICall("pool.set_default_SR", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetNameDescription: Set the name_description field of the given pool. */
func (client *XenClient) PoolSetNameDescription(self string, value string) (err error) {
	_, err = client.APICall("pool.set_name_description", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetNameLabel: Set the name_label field of the given pool. */
func (client *XenClient) PoolSetNameLabel(self string, value string) (err error) {
	_, err = client.APICall("pool.set_name_label", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetUefiCertificates: Get the uefi_certificates field of the given pool. */
func (client *XenClient) PoolGetUefiCertificates(self string) (result string, err error) {
	obj, err := client.APICall("pool.get_uefi_certificates", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetIgmpSnoopingEnabled: Get the igmp_snooping_enabled field of the given pool. */
func (client *XenClient) PoolGetIgmpSnoopingEnabled(self string) (result bool, err error) {
	obj, err := client.APICall("pool.get_igmp_snooping_enabled", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetLivePatchingDisabled: Get the live_patching_disabled field of the given pool. */
func (client *XenClient) PoolGetLivePatchingDisabled(self string) (result bool, err error) {
	obj, err := client.APICall("pool.get_live_patching_disabled", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetPolicyNoVendorDevice: Get the policy_no_vendor_device field of the given pool. */
func (client *XenClient) PoolGetPolicyNoVendorDevice(self string) (result bool, err error) {
	obj, err := client.APICall("pool.get_policy_no_vendor_device", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetCpuInfo: Get the cpu_info field of the given pool. */
func (client *XenClient) PoolGetCpuInfo(self string) (result map[string]string, err error) {
	obj, err := client.APICall("pool.get_cpu_info", self)
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

/* GetGuestAgentConfig: Get the guest_agent_config field of the given pool. */
func (client *XenClient) PoolGetGuestAgentConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("pool.get_guest_agent_config", self)
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

/* GetCurrentOperations: Get the current_operations field of the given pool. */
func (client *XenClient) PoolGetCurrentOperations(self string) (result map[string]PoolAllowedOperations, err error) {
	obj, err := client.APICall("pool.get_current_operations", self)
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]PoolAllowedOperations{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToPoolAllowedOperations(obj.String())
		result[key.String()] = mapObj
	}

	return
}

/* GetAllowedOperations: Get the allowed_operations field of the given pool. */
func (client *XenClient) PoolGetAllowedOperations(self string) (result []PoolAllowedOperations, err error) {
	obj, err := client.APICall("pool.get_allowed_operations", self)
	if err != nil {
		return
	}

	result = make([]PoolAllowedOperations, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = ToPoolAllowedOperations(value.(string))
	}

	return
}

/* GetHaClusterStack: Get the ha_cluster_stack field of the given pool. */
func (client *XenClient) PoolGetHaClusterStack(self string) (result string, err error) {
	obj, err := client.APICall("pool.get_ha_cluster_stack", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetMetadataVDIs: Get the metadata_VDIs field of the given pool. */
func (client *XenClient) PoolGetMetadataVDIs(self string) (result []string, err error) {
	obj, err := client.APICall("pool.get_metadata_VDIs", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetRestrictions: Get the restrictions field of the given pool. */
func (client *XenClient) PoolGetRestrictions(self string) (result map[string]string, err error) {
	obj, err := client.APICall("pool.get_restrictions", self)
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

/* GetVswitchController: Get the vswitch_controller field of the given pool. */
func (client *XenClient) PoolGetVswitchController(self string) (result string, err error) {
	obj, err := client.APICall("pool.get_vswitch_controller", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRedoLogVdi: Get the redo_log_vdi field of the given pool. */
func (client *XenClient) PoolGetRedoLogVdi(self string) (result string, err error) {
	obj, err := client.APICall("pool.get_redo_log_vdi", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRedoLogEnabled: Get the redo_log_enabled field of the given pool. */
func (client *XenClient) PoolGetRedoLogEnabled(self string) (result bool, err error) {
	obj, err := client.APICall("pool.get_redo_log_enabled", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetWlbVerifyCert: Get the wlb_verify_cert field of the given pool. */
func (client *XenClient) PoolGetWlbVerifyCert(self string) (result bool, err error) {
	obj, err := client.APICall("pool.get_wlb_verify_cert", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetWlbEnabled: Get the wlb_enabled field of the given pool. */
func (client *XenClient) PoolGetWlbEnabled(self string) (result bool, err error) {
	obj, err := client.APICall("pool.get_wlb_enabled", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetWlbUsername: Get the wlb_username field of the given pool. */
func (client *XenClient) PoolGetWlbUsername(self string) (result string, err error) {
	obj, err := client.APICall("pool.get_wlb_username", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetWlbUrl: Get the wlb_url field of the given pool. */
func (client *XenClient) PoolGetWlbUrl(self string) (result string, err error) {
	obj, err := client.APICall("pool.get_wlb_url", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetHealthCheckConfig: Get the health_check_config field of the given pool. */
func (client *XenClient) PoolGetHealthCheckConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("pool.get_health_check_config", self)
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

/* GetGuiConfig: Get the gui_config field of the given pool. */
func (client *XenClient) PoolGetGuiConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("pool.get_gui_config", self)
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

/* GetTags: Get the tags field of the given pool. */
func (client *XenClient) PoolGetTags(self string) (result []string, err error) {
	obj, err := client.APICall("pool.get_tags", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetBlobs: Get the blobs field of the given pool. */
func (client *XenClient) PoolGetBlobs(self string) (result map[string]string, err error) {
	obj, err := client.APICall("pool.get_blobs", self)
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

/* GetHaOvercommitted: Get the ha_overcommitted field of the given pool. */
func (client *XenClient) PoolGetHaOvercommitted(self string) (result bool, err error) {
	obj, err := client.APICall("pool.get_ha_overcommitted", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetHaAllowOvercommit: Get the ha_allow_overcommit field of the given pool. */
func (client *XenClient) PoolGetHaAllowOvercommit(self string) (result bool, err error) {
	obj, err := client.APICall("pool.get_ha_allow_overcommit", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetHaPlanExistsFor: Get the ha_plan_exists_for field of the given pool. */
func (client *XenClient) PoolGetHaPlanExistsFor(self string) (result int, err error) {
	obj, err := client.APICall("pool.get_ha_plan_exists_for", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetHaHostFailuresToTolerate: Get the ha_host_failures_to_tolerate field of the given pool. */
func (client *XenClient) PoolGetHaHostFailuresToTolerate(self string) (result int, err error) {
	obj, err := client.APICall("pool.get_ha_host_failures_to_tolerate", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetHaStatefiles: Get the ha_statefiles field of the given pool. */
func (client *XenClient) PoolGetHaStatefiles(self string) (result []string, err error) {
	obj, err := client.APICall("pool.get_ha_statefiles", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetHaConfiguration: Get the ha_configuration field of the given pool. */
func (client *XenClient) PoolGetHaConfiguration(self string) (result map[string]string, err error) {
	obj, err := client.APICall("pool.get_ha_configuration", self)
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

/* GetHaEnabled: Get the ha_enabled field of the given pool. */
func (client *XenClient) PoolGetHaEnabled(self string) (result bool, err error) {
	obj, err := client.APICall("pool.get_ha_enabled", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetOtherConfig: Get the other_config field of the given pool. */
func (client *XenClient) PoolGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("pool.get_other_config", self)
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

/* GetCrashDumpSR: Get the crash_dump_SR field of the given pool. */
func (client *XenClient) PoolGetCrashDumpSR(self string) (result string, err error) {
	obj, err := client.APICall("pool.get_crash_dump_SR", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetSuspendImageSR: Get the suspend_image_SR field of the given pool. */
func (client *XenClient) PoolGetSuspendImageSR(self string) (result string, err error) {
	obj, err := client.APICall("pool.get_suspend_image_SR", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetDefaultSR: Get the default_SR field of the given pool. */
func (client *XenClient) PoolGetDefaultSR(self string) (result string, err error) {
	obj, err := client.APICall("pool.get_default_SR", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetMaster: Get the master field of the given pool. */
func (client *XenClient) PoolGetMaster(self string) (result string, err error) {
	obj, err := client.APICall("pool.get_master", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetNameDescription: Get the name_description field of the given pool. */
func (client *XenClient) PoolGetNameDescription(self string) (result string, err error) {
	obj, err := client.APICall("pool.get_name_description", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetNameLabel: Get the name_label field of the given pool. */
func (client *XenClient) PoolGetNameLabel(self string) (result string, err error) {
	obj, err := client.APICall("pool.get_name_label", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given pool. */
func (client *XenClient) PoolGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("pool.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the pool instance with the specified UUID. */
func (client *XenClient) PoolGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("pool.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given pool. */
func (client *XenClient) PoolGetRecord(self string) (result Pool, err error) {
	obj, err := client.APICall("pool.get_record", self)
	if err != nil {
		return
	}

	result = *ToPool(obj.(interface{}))

	return
}

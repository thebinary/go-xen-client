// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w Host.gen.go

package go_xen_client

import (
	"reflect"
	"strconv"
	"time"

	"github.com/nilshell/xmlrpc"
)

//Host: A physical host
type Host struct {
	Uuid                            string                           // Unique identifier/object reference
	NameLabel                       string                           // a human-readable name
	NameDescription                 string                           // a notes field containing human-readable description
	MemoryOverhead                  int                              // Virtualization memory overhead (bytes).
	AllowedOperations               []HostAllowedOperations          // list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	CurrentOperations               map[string]HostAllowedOperations // links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	APIVersionMajor                 int                              // major version number
	APIVersionMinor                 int                              // minor version number
	APIVersionVendor                string                           // identification of vendor
	APIVersionVendorImplementation  map[string]string                // details of vendor implementation
	Enabled                         bool                             // True if the host is currently enabled
	SoftwareVersion                 map[string]string                // version strings
	OtherConfig                     map[string]string                // additional configuration
	Capabilities                    []string                         // Xen capabilities
	CpuConfiguration                map[string]string                // The CPU configuration on this host.  May contain keys such as "nr_nodes", "sockets_per_node", "cores_per_socket", or "threads_per_core"
	SchedPolicy                     string                           // Scheduler policy currently in force on this host
	SupportedBootloaders            []string                         // a list of the bootloaders installed on the machine
	ResidentVMs                     []string                         // list of VMs currently resident on host
	Logging                         map[string]string                // logging configuration
	PIFs                            []string                         // physical network interfaces
	SuspendImageSr                  string                           // The SR in which VDIs for suspend images are created
	CrashDumpSr                     string                           // The SR in which VDIs for crash dumps are created
	Crashdumps                      []string                         // Set of host crash dumps
	Patches                         []string                         // Set of host patches
	Updates                         []string                         // Set of updates
	PBDs                            []string                         // physical blockdevices
	HostCPUs                        []string                         // The physical CPUs on this host
	CpuInfo                         map[string]string                // Details about the physical CPUs on this host
	Hostname                        string                           // The hostname of this host
	Address                         string                           // The address by which this host can be contacted from any other host in the pool
	Metrics                         string                           // metrics associated with this host
	LicenseParams                   map[string]string                // State of the current license
	HaStatefiles                    []string                         // The set of statefiles accessible from this host
	HaNetworkPeers                  []string                         // The set of hosts visible via the network from this host
	Blobs                           map[string]string                // Binary blobs associated with this host
	Tags                            []string                         // user-specified tags for categorization purposes
	ExternalAuthType                string                           // type of external authentication service configured; empty if none configured.
	ExternalAuthServiceName         string                           // name of external authentication service configured; empty if none configured.
	ExternalAuthConfiguration       map[string]string                // configuration specific to external authentication service
	Edition                         string                           // Product edition
	LicenseServer                   map[string]string                // Contact information of the license server
	BiosStrings                     map[string]string                // BIOS strings
	PowerOnMode                     string                           // The power on mode
	PowerOnConfig                   map[string]string                // The power on config
	LocalCacheSr                    string                           // The SR that is used as a local cache
	ChipsetInfo                     map[string]string                // Information about chipset features
	PCIs                            []string                         // List of PCI devices in the host
	PGPUs                           []string                         // List of physical GPUs in the host
	PUSBs                           []string                         // List of physical USBs in the host
	SslLegacy                       bool                             // Allow SSLv3 protocol and ciphersuites as used by older server versions. This controls both incoming and outgoing connections. When this is set to a different value, the host immediately restarts its SSL/TLS listening service; typically this takes less than a second but existing connections to it will be broken. API login sessions will remain valid.
	GuestVCPUsParams                map[string]string                // VCPUs params to apply to all resident guests
	Display                         HostDisplay                      // indicates whether the host is configured to output its console to a physical display device
	VirtualHardwarePlatformVersions []int                            // The set of versions of the virtual hardware platform that the host can offer to its guests
	ControlDomain                   string                           // The control domain (domain 0)
	UpdatesRequiringReboot          []string                         // List of updates which require reboot
	Features                        []string                         // List of features available on this host
	IscsiIqn                        string                           // The initiator IQN for the host
	Multipathing                    bool                             // Specifies whether multipathing is enabled
	UefiCertificates                string                           // The UEFI certificates allowing Secure Boot

}

func FromHostToXml(host *Host) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] =

		host.Uuid

	result["name_label"] =

		host.NameLabel

	result["name_description"] =

		host.NameDescription

	result["memory_overhead"] =

		strconv.Itoa(host.MemoryOverhead)

	result["allowed_operations"] =

		host.AllowedOperations

	current_operations := make(xmlrpc.Struct)
	for key, value := range host.CurrentOperations {
		current_operations[key] = value
	}
	result["current_operations"] = current_operations

	result["API_version_major"] =

		strconv.Itoa(host.APIVersionMajor)

	result["API_version_minor"] =

		strconv.Itoa(host.APIVersionMinor)

	result["API_version_vendor"] =

		host.APIVersionVendor

	API_version_vendor_implementation := make(xmlrpc.Struct)
	for key, value := range host.APIVersionVendorImplementation {
		API_version_vendor_implementation[key] = value
	}
	result["API_version_vendor_implementation"] = API_version_vendor_implementation

	result["enabled"] =

		host.Enabled

	software_version := make(xmlrpc.Struct)
	for key, value := range host.SoftwareVersion {
		software_version[key] = value
	}
	result["software_version"] = software_version

	other_config := make(xmlrpc.Struct)
	for key, value := range host.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	result["capabilities"] =

		host.Capabilities

	cpu_configuration := make(xmlrpc.Struct)
	for key, value := range host.CpuConfiguration {
		cpu_configuration[key] = value
	}
	result["cpu_configuration"] = cpu_configuration

	result["sched_policy"] =

		host.SchedPolicy

	result["supported_bootloaders"] =

		host.SupportedBootloaders

	result["resident_VMs"] =

		host.ResidentVMs

	logging := make(xmlrpc.Struct)
	for key, value := range host.Logging {
		logging[key] = value
	}
	result["logging"] = logging

	result["PIFs"] =

		host.PIFs

	result["suspend_image_sr"] =

		host.SuspendImageSr

	result["crash_dump_sr"] =

		host.CrashDumpSr

	result["crashdumps"] =

		host.Crashdumps

	result["patches"] =

		host.Patches

	result["updates"] =

		host.Updates

	result["PBDs"] =

		host.PBDs

	result["host_CPUs"] =

		host.HostCPUs

	cpu_info := make(xmlrpc.Struct)
	for key, value := range host.CpuInfo {
		cpu_info[key] = value
	}
	result["cpu_info"] = cpu_info

	result["hostname"] =

		host.Hostname

	result["address"] =

		host.Address

	result["metrics"] =

		host.Metrics

	license_params := make(xmlrpc.Struct)
	for key, value := range host.LicenseParams {
		license_params[key] = value
	}
	result["license_params"] = license_params

	result["ha_statefiles"] =

		host.HaStatefiles

	result["ha_network_peers"] =

		host.HaNetworkPeers

	blobs := make(xmlrpc.Struct)
	for key, value := range host.Blobs {
		blobs[key] = value
	}
	result["blobs"] = blobs

	result["tags"] =

		host.Tags

	result["external_auth_type"] =

		host.ExternalAuthType

	result["external_auth_service_name"] =

		host.ExternalAuthServiceName

	external_auth_configuration := make(xmlrpc.Struct)
	for key, value := range host.ExternalAuthConfiguration {
		external_auth_configuration[key] = value
	}
	result["external_auth_configuration"] = external_auth_configuration

	result["edition"] =

		host.Edition

	license_server := make(xmlrpc.Struct)
	for key, value := range host.LicenseServer {
		license_server[key] = value
	}
	result["license_server"] = license_server

	bios_strings := make(xmlrpc.Struct)
	for key, value := range host.BiosStrings {
		bios_strings[key] = value
	}
	result["bios_strings"] = bios_strings

	result["power_on_mode"] =

		host.PowerOnMode

	power_on_config := make(xmlrpc.Struct)
	for key, value := range host.PowerOnConfig {
		power_on_config[key] = value
	}
	result["power_on_config"] = power_on_config

	result["local_cache_sr"] =

		host.LocalCacheSr

	chipset_info := make(xmlrpc.Struct)
	for key, value := range host.ChipsetInfo {
		chipset_info[key] = value
	}
	result["chipset_info"] = chipset_info

	result["PCIs"] =

		host.PCIs

	result["PGPUs"] =

		host.PGPUs

	result["PUSBs"] =

		host.PUSBs

	result["ssl_legacy"] =

		host.SslLegacy

	guest_VCPUs_params := make(xmlrpc.Struct)
	for key, value := range host.GuestVCPUsParams {
		guest_VCPUs_params[key] = value
	}
	result["guest_VCPUs_params"] = guest_VCPUs_params

	result["display"] =

		host.Display.String()

	result["virtual_hardware_platform_versions"] =

		host.VirtualHardwarePlatformVersions

	result["control_domain"] =

		host.ControlDomain

	result["updates_requiring_reboot"] =

		host.UpdatesRequiringReboot

	result["features"] =

		host.Features

	result["iscsi_iqn"] =

		host.IscsiIqn

	result["multipathing"] =

		host.Multipathing

	result["uefi_certificates"] =

		host.UefiCertificates

	return result
}

func ToHost(obj interface{}) (resultObj *Host) {

	objValue := reflect.ValueOf(obj)
	resultObj = &Host{}

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
		case "memory_overhead":
			if v, ok := keyValue.(int); ok {
				resultObj.MemoryOverhead = v
			}
		case "allowed_operations":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.AllowedOperations = make([]HostAllowedOperations, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(HostAllowedOperations); ok {
						resultObj.AllowedOperations[i] = v
					}
				}
			}
		case "current_operations":

			resultObj.CurrentOperations = map[string]HostAllowedOperations{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.CurrentOperations[mapKeyName] = ToHostAllowedOperations(v)
				} else {
					resultObj.CurrentOperations[mapKeyName] = 0
				}
			}
		case "API_version_major":
			if v, ok := keyValue.(int); ok {
				resultObj.APIVersionMajor = v
			}
		case "API_version_minor":
			if v, ok := keyValue.(int); ok {
				resultObj.APIVersionMinor = v
			}
		case "API_version_vendor":
			if v, ok := keyValue.(string); ok {
				resultObj.APIVersionVendor = v
			}
		case "API_version_vendor_implementation":

			resultObj.APIVersionVendorImplementation = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.APIVersionVendorImplementation[mapKeyName] = v
				} else {
					resultObj.APIVersionVendorImplementation[mapKeyName] = ""
				}
			}
		case "enabled":
			if v, ok := keyValue.(bool); ok {
				resultObj.Enabled = v
			}
		case "software_version":

			resultObj.SoftwareVersion = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.SoftwareVersion[mapKeyName] = v
				} else {
					resultObj.SoftwareVersion[mapKeyName] = ""
				}
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
		case "capabilities":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.Capabilities = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.Capabilities[i] = v
					}
				}
			}
		case "cpu_configuration":

			resultObj.CpuConfiguration = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.CpuConfiguration[mapKeyName] = v
				} else {
					resultObj.CpuConfiguration[mapKeyName] = ""
				}
			}
		case "sched_policy":
			if v, ok := keyValue.(string); ok {
				resultObj.SchedPolicy = v
			}
		case "supported_bootloaders":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.SupportedBootloaders = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.SupportedBootloaders[i] = v
					}
				}
			}
		case "resident_VMs":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.ResidentVMs = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.ResidentVMs[i] = v
					}
				}
			}
		case "logging":

			resultObj.Logging = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.Logging[mapKeyName] = v
				} else {
					resultObj.Logging[mapKeyName] = ""
				}
			}
		case "PIFs":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.PIFs = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.PIFs[i] = v
					}
				}
			}
		case "suspend_image_sr":
			if v, ok := keyValue.(string); ok {
				resultObj.SuspendImageSr = v
			}
		case "crash_dump_sr":
			if v, ok := keyValue.(string); ok {
				resultObj.CrashDumpSr = v
			}
		case "crashdumps":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.Crashdumps = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.Crashdumps[i] = v
					}
				}
			}
		case "patches":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.Patches = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.Patches[i] = v
					}
				}
			}
		case "updates":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.Updates = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.Updates[i] = v
					}
				}
			}
		case "PBDs":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.PBDs = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.PBDs[i] = v
					}
				}
			}
		case "host_CPUs":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.HostCPUs = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.HostCPUs[i] = v
					}
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
		case "hostname":
			if v, ok := keyValue.(string); ok {
				resultObj.Hostname = v
			}
		case "address":
			if v, ok := keyValue.(string); ok {
				resultObj.Address = v
			}
		case "metrics":
			if v, ok := keyValue.(string); ok {
				resultObj.Metrics = v
			}
		case "license_params":

			resultObj.LicenseParams = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.LicenseParams[mapKeyName] = v
				} else {
					resultObj.LicenseParams[mapKeyName] = ""
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
		case "ha_network_peers":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.HaNetworkPeers = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.HaNetworkPeers[i] = v
					}
				}
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
		case "external_auth_type":
			if v, ok := keyValue.(string); ok {
				resultObj.ExternalAuthType = v
			}
		case "external_auth_service_name":
			if v, ok := keyValue.(string); ok {
				resultObj.ExternalAuthServiceName = v
			}
		case "external_auth_configuration":

			resultObj.ExternalAuthConfiguration = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.ExternalAuthConfiguration[mapKeyName] = v
				} else {
					resultObj.ExternalAuthConfiguration[mapKeyName] = ""
				}
			}
		case "edition":
			if v, ok := keyValue.(string); ok {
				resultObj.Edition = v
			}
		case "license_server":

			resultObj.LicenseServer = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.LicenseServer[mapKeyName] = v
				} else {
					resultObj.LicenseServer[mapKeyName] = ""
				}
			}
		case "bios_strings":

			resultObj.BiosStrings = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.BiosStrings[mapKeyName] = v
				} else {
					resultObj.BiosStrings[mapKeyName] = ""
				}
			}
		case "power_on_mode":
			if v, ok := keyValue.(string); ok {
				resultObj.PowerOnMode = v
			}
		case "power_on_config":

			resultObj.PowerOnConfig = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.PowerOnConfig[mapKeyName] = v
				} else {
					resultObj.PowerOnConfig[mapKeyName] = ""
				}
			}
		case "local_cache_sr":
			if v, ok := keyValue.(string); ok {
				resultObj.LocalCacheSr = v
			}
		case "chipset_info":

			resultObj.ChipsetInfo = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.ChipsetInfo[mapKeyName] = v
				} else {
					resultObj.ChipsetInfo[mapKeyName] = ""
				}
			}
		case "PCIs":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.PCIs = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.PCIs[i] = v
					}
				}
			}
		case "PGPUs":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.PGPUs = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.PGPUs[i] = v
					}
				}
			}
		case "PUSBs":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.PUSBs = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.PUSBs[i] = v
					}
				}
			}
		case "ssl_legacy":
			if v, ok := keyValue.(bool); ok {
				resultObj.SslLegacy = v
			}
		case "guest_VCPUs_params":

			resultObj.GuestVCPUsParams = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.GuestVCPUsParams[mapKeyName] = v
				} else {
					resultObj.GuestVCPUsParams[mapKeyName] = ""
				}
			}
		case "display":
			if v, ok := keyValue.(HostDisplay); ok {
				resultObj.Display = v
			}
		case "virtual_hardware_platform_versions":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.VirtualHardwarePlatformVersions = make([]int, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(int); ok {
						resultObj.VirtualHardwarePlatformVersions[i] = v
					}
				}
			}
		case "control_domain":
			if v, ok := keyValue.(string); ok {
				resultObj.ControlDomain = v
			}
		case "updates_requiring_reboot":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.UpdatesRequiringReboot = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.UpdatesRequiringReboot[i] = v
					}
				}
			}
		case "features":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.Features = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.Features[i] = v
					}
				}
			}
		case "iscsi_iqn":
			if v, ok := keyValue.(string); ok {
				resultObj.IscsiIqn = v
			}
		case "multipathing":
			if v, ok := keyValue.(bool); ok {
				resultObj.Multipathing = v
			}
		case "uefi_certificates":
			if v, ok := keyValue.(string); ok {
				resultObj.UefiCertificates = v
			}
		}
	}

	return resultObj
}

/* GetAllRecords: Return a map of host references to host records for all hosts known to the system. */
func (client *XenClient) HostGetAllRecords() (result map[string]Host, err error) {
	obj, err := client.APICall("host.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]Host{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToHost(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the hosts known to the system. */
func (client *XenClient) HostGetAll() (result []string, err error) {
	obj, err := client.APICall("host.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* SetUefiCertificates: Sets the UEFI certificates on a host */
func (client *XenClient) HostSetUefiCertificates(host string, value string) (err error) {
	_, err = client.APICall("host.set_uefi_certificates", host, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetMultipathing: Specifies whether multipathing is enabled */
func (client *XenClient) HostSetMultipathing(host string, value bool) (err error) {
	_, err = client.APICall("host.set_multipathing", host, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetIscsiIqn: Sets the initiator IQN for the host */
func (client *XenClient) HostSetIscsiIqn(host string, value string) (err error) {
	_, err = client.APICall("host.set_iscsi_iqn", host, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetSslLegacy: Enable/disable SSLv3 for interoperability with older server versions. When this is set to a different value, the host immediately restarts its SSL/TLS listening service; typically this takes less than a second but existing connections to it will be broken. API login sessions will remain valid. */
func (client *XenClient) HostSetSslLegacy(self string, value bool) (err error) {
	_, err = client.APICall("host.set_ssl_legacy", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* DisableDisplay: Disable console output to the physical display device next time this host boots */
func (client *XenClient) HostDisableDisplay(host string) (result HostDisplay, err error) {
	obj, err := client.APICall("host.disable_display", host)
	if err != nil {
		return
	}

	result = ToHostDisplay(obj.(string))

	return
}

/* EnableDisplay: Enable console output to the physical display device next time this host boots */
func (client *XenClient) HostEnableDisplay(host string) (result HostDisplay, err error) {
	obj, err := client.APICall("host.enable_display", host)
	if err != nil {
		return
	}

	result = ToHostDisplay(obj.(string))

	return
}

/* DeclareDead: Declare that a host is dead. This is a dangerous operation, and should only be called if the administrator is absolutely sure the host is definitely dead */
func (client *XenClient) HostDeclareDead(host string) (err error) {
	_, err = client.APICall("host.declare_dead", host)
	if err != nil {
		return
	}
	// no return result
	return
}

/* MigrateReceive: Prepare to receive a VM, returning a token which can be passed to VM.migrate. */
func (client *XenClient) HostMigrateReceive(host string, network string, options map[string]string) (result map[string]string, err error) {
	obj, err := client.APICall("host.migrate_receive", host, network, options)
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

/* DisableLocalStorageCaching: Disable the use of a local SR for caching purposes */
func (client *XenClient) HostDisableLocalStorageCaching(host string) (err error) {
	_, err = client.APICall("host.disable_local_storage_caching", host)
	if err != nil {
		return
	}
	// no return result
	return
}

/* EnableLocalStorageCaching: Enable the use of a local SR for caching purposes */
func (client *XenClient) HostEnableLocalStorageCaching(host string, sr string) (err error) {
	_, err = client.APICall("host.enable_local_storage_caching", host, sr)
	if err != nil {
		return
	}
	// no return result
	return
}

/* ResetCpuFeatures: Remove the feature mask, such that after a reboot all features of the CPU are enabled. */
func (client *XenClient) HostResetCpuFeatures(host string) (err error) {
	_, err = client.APICall("host.reset_cpu_features", host)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetCpuFeatures: Set the CPU features to be used after a reboot, if the given features string is valid. */
func (client *XenClient) HostSetCpuFeatures(host string, features string) (err error) {
	_, err = client.APICall("host.set_cpu_features", host, features)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetPowerOnMode: Set the power-on-mode, host, user and password  */
func (client *XenClient) HostSetPowerOnMode(self string, power_on_mode string, power_on_config map[string]string) (err error) {
	_, err = client.APICall("host.set_power_on_mode", self, power_on_mode, power_on_config)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RefreshPackInfo: Refresh the list of installed Supplemental Packs. */
func (client *XenClient) HostRefreshPackInfo(host string) (err error) {
	_, err = client.APICall("host.refresh_pack_info", host)
	if err != nil {
		return
	}
	// no return result
	return
}

/* ApplyEdition: Change to another edition, or reactivate the current edition after a license has expired. This may be subject to the successful checkout of an appropriate license. */
func (client *XenClient) HostApplyEdition(host string, edition string, force bool) (err error) {
	_, err = client.APICall("host.apply_edition", host, edition, force)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetServerCertificate: Get the installed server public TLS certificate. */
func (client *XenClient) HostGetServerCertificate(host string) (result string, err error) {
	obj, err := client.APICall("host.get_server_certificate", host)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* RetrieveWlbEvacuateRecommendations: Retrieves recommended host migrations to perform when evacuating the host from the wlb server. If a VM cannot be migrated from the host the reason is listed instead of a recommendation. */
func (client *XenClient) HostRetrieveWlbEvacuateRecommendations(self string) (result map[string][]string, err error) {
	obj, err := client.APICall("host.retrieve_wlb_evacuate_recommendations", self)
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

/* DisableExternalAuth: This call disables external authentication on the local host */
func (client *XenClient) HostDisableExternalAuth(host string, config map[string]string) (err error) {
	_, err = client.APICall("host.disable_external_auth", host, config)
	if err != nil {
		return
	}
	// no return result
	return
}

/* EnableExternalAuth: This call enables external authentication on a host */
func (client *XenClient) HostEnableExternalAuth(host string, config map[string]string, service_name string, auth_type string) (err error) {
	_, err = client.APICall("host.enable_external_auth", host, config, service_name, auth_type)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetServerLocaltime: This call queries the host's clock for the current time in the host's local timezone */
func (client *XenClient) HostGetServerLocaltime(host string) (result time.Time, err error) {
	obj, err := client.APICall("host.get_server_localtime", host)
	if err != nil {
		return
	}
	result = obj.(time.Time)
	return
}

/* GetServertime: This call queries the host's clock for the current time */
func (client *XenClient) HostGetServertime(host string) (result time.Time, err error) {
	obj, err := client.APICall("host.get_servertime", host)
	if err != nil {
		return
	}
	result = obj.(time.Time)
	return
}

/* CallExtension: Call an API extension on this host */
func (client *XenClient) HostCallExtension(host string, call string) (result string, err error) {
	obj, err := client.APICall("host.call_extension", host, call)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* HasExtension: Return true if the extension is available on the host */
func (client *XenClient) HostHasExtension(host string, name string) (result bool, err error) {
	obj, err := client.APICall("host.has_extension", host, name)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* CallPlugin: Call an API plugin on this host */
func (client *XenClient) HostCallPlugin(host string, plugin string, fn string, args map[string]string) (result string, err error) {
	obj, err := client.APICall("host.call_plugin", host, plugin, fn, args)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* CreateNewBlob: Create a placeholder for a named binary blob of data that is associated with this host */
func (client *XenClient) HostCreateNewBlob(host string, name string, mime_type string, public bool) (result string, err error) {
	obj, err := client.APICall("host.create_new_blob", host, name, mime_type, public)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* BackupRrds: This causes the RRDs to be backed up to the master */
func (client *XenClient) HostBackupRrds(host string, delay float32) (err error) {
	_, err = client.APICall("host.backup_rrds", host, delay)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SyncData: This causes the synchronisation of the non-database data (messages, RRDs and so on) stored on the master to be synchronised with the host */
func (client *XenClient) HostSyncData(host string) (err error) {
	_, err = client.APICall("host.sync_data", host)
	if err != nil {
		return
	}
	// no return result
	return
}

/* ComputeMemoryOverhead: Computes the virtualization memory overhead of a host. */
func (client *XenClient) HostComputeMemoryOverhead(host string) (result int, err error) {
	obj, err := client.APICall("host.compute_memory_overhead", host)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* ComputeFreeMemory: Computes the amount of free memory on the host. */
func (client *XenClient) HostComputeFreeMemory(host string) (result int, err error) {
	obj, err := client.APICall("host.compute_free_memory", host)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* SetHostnameLive: Sets the host name to the specified string.  Both the API and lower-level system hostname are changed immediately. */
func (client *XenClient) HostSetHostnameLive(host string, hostname string) (err error) {
	_, err = client.APICall("host.set_hostname_live", host, hostname)
	if err != nil {
		return
	}
	// no return result
	return
}

/* ShutdownAgent: Shuts the agent down after a 10 second pause. WARNING: this is a dangerous operation. Any operations in progress will be aborted, and unrecoverable data loss may occur. The caller is responsible for ensuring that there are no operations in progress when this method is called. */
func (client *XenClient) HostShutdownAgent() (err error) {
	_, err = client.APICall("host.shutdown_agent")
	if err != nil {
		return
	}
	// no return result
	return
}

/* RestartAgent: Restarts the agent after a 10 second pause. WARNING: this is a dangerous operation. Any operations in progress will be aborted, and unrecoverable data loss may occur. The caller is responsible for ensuring that there are no operations in progress when this method is called. */
func (client *XenClient) HostRestartAgent(host string) (err error) {
	_, err = client.APICall("host.restart_agent", host)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetSystemStatusCapabilities:  */
func (client *XenClient) HostGetSystemStatusCapabilities(host string) (result string, err error) {
	obj, err := client.APICall("host.get_system_status_capabilities", host)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetManagementInterface: Returns the management interface for the specified host */
func (client *XenClient) HostGetManagementInterface(host string) (result string, err error) {
	obj, err := client.APICall("host.get_management_interface", host)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* ManagementDisable: Disable the management network interface */
func (client *XenClient) HostManagementDisable() (err error) {
	_, err = client.APICall("host.management_disable")
	if err != nil {
		return
	}
	// no return result
	return
}

/* LocalManagementReconfigure: Reconfigure the management network interface. Should only be used if Host.management_reconfigure is impossible because the network configuration is broken. */
func (client *XenClient) HostLocalManagementReconfigure(iface string) (err error) {
	_, err = client.APICall("host.local_management_reconfigure", iface)
	if err != nil {
		return
	}
	// no return result
	return
}

/* ManagementReconfigure: Reconfigure the management network interface */
func (client *XenClient) HostManagementReconfigure(pif string) (err error) {
	_, err = client.APICall("host.management_reconfigure", pif)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SyslogReconfigure: Re-configure syslog logging */
func (client *XenClient) HostSyslogReconfigure(host string) (err error) {
	_, err = client.APICall("host.syslog_reconfigure", host)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Evacuate: Migrate all VMs off of this host, where possible. */
func (client *XenClient) HostEvacuate(host string) (err error) {
	_, err = client.APICall("host.evacuate", host)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetUncooperativeResidentVMs: Return a set of VMs which are not co-operating with the host's memory control system */
func (client *XenClient) HostGetUncooperativeResidentVMs(self string) (result []string, err error) {
	obj, err := client.APICall("host.get_uncooperative_resident_VMs", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetVmsWhichPreventEvacuation: Return a set of VMs which prevent the host being evacuated, with per-VM error codes */
func (client *XenClient) HostGetVmsWhichPreventEvacuation(self string) (result map[string][]string, err error) {
	obj, err := client.APICall("host.get_vms_which_prevent_evacuation", self)
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

/* AssertCanEvacuate: Check this host can be evacuated. */
func (client *XenClient) HostAssertCanEvacuate(host string) (err error) {
	_, err = client.APICall("host.assert_can_evacuate", host)
	if err != nil {
		return
	}
	// no return result
	return
}

/* ForgetDataSourceArchives: Forget the recorded statistics related to the specified data source */
func (client *XenClient) HostForgetDataSourceArchives(host string, data_source string) (err error) {
	_, err = client.APICall("host.forget_data_source_archives", host, data_source)
	if err != nil {
		return
	}
	// no return result
	return
}

/* QueryDataSource: Query the latest value of the specified data source */
func (client *XenClient) HostQueryDataSource(host string, data_source string) (result float32, err error) {
	obj, err := client.APICall("host.query_data_source", host, data_source)
	if err != nil {
		return
	}
	result = obj.(float32)
	return
}

/* RecordDataSource: Start recording the specified data source */
func (client *XenClient) HostRecordDataSource(host string, data_source string) (err error) {
	_, err = client.APICall("host.record_data_source", host, data_source)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetDataSources:  */
func (client *XenClient) HostGetDataSources(host string) (result []DataSource, err error) {
	obj, err := client.APICall("host.get_data_sources", host)
	if err != nil {
		return
	}

	result = make([]DataSource, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = *ToDataSource(value)
	}

	return
}

/* EmergencyHaDisable: This call disables HA on the local host. This should only be used with extreme care. */
func (client *XenClient) HostEmergencyHaDisable(soft bool) (err error) {
	_, err = client.APICall("host.emergency_ha_disable", soft)
	if err != nil {
		return
	}
	// no return result
	return
}

/* PowerOn: Attempt to power-on the host (if the capability exists). */
func (client *XenClient) HostPowerOn(host string) (err error) {
	_, err = client.APICall("host.power_on", host)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Destroy: Destroy specified host record in database */
func (client *XenClient) HostDestroy(self string) (err error) {
	_, err = client.APICall("host.destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* LicenseRemove: Remove any license file from the specified host, and switch that host to the unlicensed edition */
func (client *XenClient) HostLicenseRemove(host string) (err error) {
	_, err = client.APICall("host.license_remove", host)
	if err != nil {
		return
	}
	// no return result
	return
}

/* LicenseAdd: Apply a new license to a host */
func (client *XenClient) HostLicenseAdd(host string, contents string) (err error) {
	_, err = client.APICall("host.license_add", host, contents)
	if err != nil {
		return
	}
	// no return result
	return
}

/* LicenseApply: Apply a new license to a host */
func (client *XenClient) HostLicenseApply(host string, contents string) (err error) {
	_, err = client.APICall("host.license_apply", host, contents)
	if err != nil {
		return
	}
	// no return result
	return
}

/* ListMethods: List all supported methods */
func (client *XenClient) HostListMethods() (result []string, err error) {
	obj, err := client.APICall("host.list_methods")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* BugreportUpload: Run xen-bugtool --yestoall and upload the output to support */
func (client *XenClient) HostBugreportUpload(host string, url string, options map[string]string) (err error) {
	_, err = client.APICall("host.bugreport_upload", host, url, options)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SendDebugKeys: Inject the given string as debugging keys into Xen */
func (client *XenClient) HostSendDebugKeys(host string, keys string) (err error) {
	_, err = client.APICall("host.send_debug_keys", host, keys)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetLog: Get the host's log file */
func (client *XenClient) HostGetLog(host string) (result string, err error) {
	obj, err := client.APICall("host.get_log", host)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* DmesgClear: Get the host xen dmesg, and clear the buffer. */
func (client *XenClient) HostDmesgClear(host string) (result string, err error) {
	obj, err := client.APICall("host.dmesg_clear", host)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* Dmesg: Get the host xen dmesg. */
func (client *XenClient) HostDmesg(host string) (result string, err error) {
	obj, err := client.APICall("host.dmesg", host)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* Reboot: Reboot the host. (This function can only be called if there are no currently running VMs on the host and it is disabled.) */
func (client *XenClient) HostReboot(host string) (err error) {
	_, err = client.APICall("host.reboot", host)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Shutdown: Shutdown the host. (This function can only be called if there are no currently running VMs on the host and it is disabled.) */
func (client *XenClient) HostShutdown(host string) (err error) {
	_, err = client.APICall("host.shutdown", host)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Enable: Puts the host into a state in which new VMs can be started. */
func (client *XenClient) HostEnable(host string) (err error) {
	_, err = client.APICall("host.enable", host)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Disable: Puts the host into a state in which no new VMs can be started. Currently active VMs on the host continue to execute. */
func (client *XenClient) HostDisable(host string) (err error) {
	_, err = client.APICall("host.disable", host)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetDisplay: Set the display field of the given host. */
func (client *XenClient) HostSetDisplay(self string, value HostDisplay) (err error) {
	_, err = client.APICall("host.set_display", self, value.String())
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveFromGuestVCPUsParams: Remove the given key and its corresponding value from the guest_VCPUs_params field of the given host.  If the key is not in that Map, then do nothing. */
func (client *XenClient) HostRemoveFromGuestVCPUsParams(self string, key string) (err error) {
	_, err = client.APICall("host.remove_from_guest_VCPUs_params", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToGuestVCPUsParams: Add the given key-value pair to the guest_VCPUs_params field of the given host. */
func (client *XenClient) HostAddToGuestVCPUsParams(self string, key string, value string) (err error) {
	_, err = client.APICall("host.add_to_guest_VCPUs_params", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetGuestVCPUsParams: Set the guest_VCPUs_params field of the given host. */
func (client *XenClient) HostSetGuestVCPUsParams(self string, value map[string]string) (err error) {
	_, err = client.APICall("host.set_guest_VCPUs_params", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveFromLicenseServer: Remove the given key and its corresponding value from the license_server field of the given host.  If the key is not in that Map, then do nothing. */
func (client *XenClient) HostRemoveFromLicenseServer(self string, key string) (err error) {
	_, err = client.APICall("host.remove_from_license_server", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToLicenseServer: Add the given key-value pair to the license_server field of the given host. */
func (client *XenClient) HostAddToLicenseServer(self string, key string, value string) (err error) {
	_, err = client.APICall("host.add_to_license_server", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetLicenseServer: Set the license_server field of the given host. */
func (client *XenClient) HostSetLicenseServer(self string, value map[string]string) (err error) {
	_, err = client.APICall("host.set_license_server", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveTags: Remove the given value from the tags field of the given host.  If the value is not in that Set, then do nothing. */
func (client *XenClient) HostRemoveTags(self string, value string) (err error) {
	_, err = client.APICall("host.remove_tags", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddTags: Add the given value to the tags field of the given host.  If the value is already in that Set, then do nothing. */
func (client *XenClient) HostAddTags(self string, value string) (err error) {
	_, err = client.APICall("host.add_tags", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetTags: Set the tags field of the given host. */
func (client *XenClient) HostSetTags(self string, value []string) (err error) {
	_, err = client.APICall("host.set_tags", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetAddress: Set the address field of the given host. */
func (client *XenClient) HostSetAddress(self string, value string) (err error) {
	_, err = client.APICall("host.set_address", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetHostname: Set the hostname field of the given host. */
func (client *XenClient) HostSetHostname(self string, value string) (err error) {
	_, err = client.APICall("host.set_hostname", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetCrashDumpSr: Set the crash_dump_sr field of the given host. */
func (client *XenClient) HostSetCrashDumpSr(self string, value string) (err error) {
	_, err = client.APICall("host.set_crash_dump_sr", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetSuspendImageSr: Set the suspend_image_sr field of the given host. */
func (client *XenClient) HostSetSuspendImageSr(self string, value string) (err error) {
	_, err = client.APICall("host.set_suspend_image_sr", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveFromLogging: Remove the given key and its corresponding value from the logging field of the given host.  If the key is not in that Map, then do nothing. */
func (client *XenClient) HostRemoveFromLogging(self string, key string) (err error) {
	_, err = client.APICall("host.remove_from_logging", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToLogging: Add the given key-value pair to the logging field of the given host. */
func (client *XenClient) HostAddToLogging(self string, key string, value string) (err error) {
	_, err = client.APICall("host.add_to_logging", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetLogging: Set the logging field of the given host. */
func (client *XenClient) HostSetLogging(self string, value map[string]string) (err error) {
	_, err = client.APICall("host.set_logging", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given host.  If the key is not in that Map, then do nothing. */
func (client *XenClient) HostRemoveFromOtherConfig(self string, key string) (err error) {
	_, err = client.APICall("host.remove_from_other_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToOtherConfig: Add the given key-value pair to the other_config field of the given host. */
func (client *XenClient) HostAddToOtherConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("host.add_to_other_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOtherConfig: Set the other_config field of the given host. */
func (client *XenClient) HostSetOtherConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("host.set_other_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetNameDescription: Set the name/description field of the given host. */
func (client *XenClient) HostSetNameDescription(self string, value string) (err error) {
	_, err = client.APICall("host.set_name_description", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetNameLabel: Set the name/label field of the given host. */
func (client *XenClient) HostSetNameLabel(self string, value string) (err error) {
	_, err = client.APICall("host.set_name_label", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetUefiCertificates: Get the uefi_certificates field of the given host. */
func (client *XenClient) HostGetUefiCertificates(self string) (result string, err error) {
	obj, err := client.APICall("host.get_uefi_certificates", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetMultipathing: Get the multipathing field of the given host. */
func (client *XenClient) HostGetMultipathing(self string) (result bool, err error) {
	obj, err := client.APICall("host.get_multipathing", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetIscsiIqn: Get the iscsi_iqn field of the given host. */
func (client *XenClient) HostGetIscsiIqn(self string) (result string, err error) {
	obj, err := client.APICall("host.get_iscsi_iqn", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetFeatures: Get the features field of the given host. */
func (client *XenClient) HostGetFeatures(self string) (result []string, err error) {
	obj, err := client.APICall("host.get_features", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetUpdatesRequiringReboot: Get the updates_requiring_reboot field of the given host. */
func (client *XenClient) HostGetUpdatesRequiringReboot(self string) (result []string, err error) {
	obj, err := client.APICall("host.get_updates_requiring_reboot", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetControlDomain: Get the control_domain field of the given host. */
func (client *XenClient) HostGetControlDomain(self string) (result string, err error) {
	obj, err := client.APICall("host.get_control_domain", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetVirtualHardwarePlatformVersions: Get the virtual_hardware_platform_versions field of the given host. */
func (client *XenClient) HostGetVirtualHardwarePlatformVersions(self string) (result []int, err error) {
	obj, err := client.APICall("host.get_virtual_hardware_platform_versions", self)
	if err != nil {
		return
	}

	result = make([]int, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(int)
	}

	return
}

/* GetDisplay: Get the display field of the given host. */
func (client *XenClient) HostGetDisplay(self string) (result HostDisplay, err error) {
	obj, err := client.APICall("host.get_display", self)
	if err != nil {
		return
	}

	result = ToHostDisplay(obj.(string))

	return
}

/* GetGuestVCPUsParams: Get the guest_VCPUs_params field of the given host. */
func (client *XenClient) HostGetGuestVCPUsParams(self string) (result map[string]string, err error) {
	obj, err := client.APICall("host.get_guest_VCPUs_params", self)
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

/* GetSslLegacy: Get the ssl_legacy field of the given host. */
func (client *XenClient) HostGetSslLegacy(self string) (result bool, err error) {
	obj, err := client.APICall("host.get_ssl_legacy", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetPUSBs: Get the PUSBs field of the given host. */
func (client *XenClient) HostGetPUSBs(self string) (result []string, err error) {
	obj, err := client.APICall("host.get_PUSBs", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetPGPUs: Get the PGPUs field of the given host. */
func (client *XenClient) HostGetPGPUs(self string) (result []string, err error) {
	obj, err := client.APICall("host.get_PGPUs", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetPCIs: Get the PCIs field of the given host. */
func (client *XenClient) HostGetPCIs(self string) (result []string, err error) {
	obj, err := client.APICall("host.get_PCIs", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetChipsetInfo: Get the chipset_info field of the given host. */
func (client *XenClient) HostGetChipsetInfo(self string) (result map[string]string, err error) {
	obj, err := client.APICall("host.get_chipset_info", self)
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

/* GetLocalCacheSr: Get the local_cache_sr field of the given host. */
func (client *XenClient) HostGetLocalCacheSr(self string) (result string, err error) {
	obj, err := client.APICall("host.get_local_cache_sr", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetPowerOnConfig: Get the power_on_config field of the given host. */
func (client *XenClient) HostGetPowerOnConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("host.get_power_on_config", self)
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

/* GetPowerOnMode: Get the power_on_mode field of the given host. */
func (client *XenClient) HostGetPowerOnMode(self string) (result string, err error) {
	obj, err := client.APICall("host.get_power_on_mode", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetBiosStrings: Get the bios_strings field of the given host. */
func (client *XenClient) HostGetBiosStrings(self string) (result map[string]string, err error) {
	obj, err := client.APICall("host.get_bios_strings", self)
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

/* GetLicenseServer: Get the license_server field of the given host. */
func (client *XenClient) HostGetLicenseServer(self string) (result map[string]string, err error) {
	obj, err := client.APICall("host.get_license_server", self)
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

/* GetEdition: Get the edition field of the given host. */
func (client *XenClient) HostGetEdition(self string) (result string, err error) {
	obj, err := client.APICall("host.get_edition", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetExternalAuthConfiguration: Get the external_auth_configuration field of the given host. */
func (client *XenClient) HostGetExternalAuthConfiguration(self string) (result map[string]string, err error) {
	obj, err := client.APICall("host.get_external_auth_configuration", self)
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

/* GetExternalAuthServiceName: Get the external_auth_service_name field of the given host. */
func (client *XenClient) HostGetExternalAuthServiceName(self string) (result string, err error) {
	obj, err := client.APICall("host.get_external_auth_service_name", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetExternalAuthType: Get the external_auth_type field of the given host. */
func (client *XenClient) HostGetExternalAuthType(self string) (result string, err error) {
	obj, err := client.APICall("host.get_external_auth_type", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetTags: Get the tags field of the given host. */
func (client *XenClient) HostGetTags(self string) (result []string, err error) {
	obj, err := client.APICall("host.get_tags", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetBlobs: Get the blobs field of the given host. */
func (client *XenClient) HostGetBlobs(self string) (result map[string]string, err error) {
	obj, err := client.APICall("host.get_blobs", self)
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

/* GetHaNetworkPeers: Get the ha_network_peers field of the given host. */
func (client *XenClient) HostGetHaNetworkPeers(self string) (result []string, err error) {
	obj, err := client.APICall("host.get_ha_network_peers", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetHaStatefiles: Get the ha_statefiles field of the given host. */
func (client *XenClient) HostGetHaStatefiles(self string) (result []string, err error) {
	obj, err := client.APICall("host.get_ha_statefiles", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetLicenseParams: Get the license_params field of the given host. */
func (client *XenClient) HostGetLicenseParams(self string) (result map[string]string, err error) {
	obj, err := client.APICall("host.get_license_params", self)
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

/* GetMetrics: Get the metrics field of the given host. */
func (client *XenClient) HostGetMetrics(self string) (result string, err error) {
	obj, err := client.APICall("host.get_metrics", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetAddress: Get the address field of the given host. */
func (client *XenClient) HostGetAddress(self string) (result string, err error) {
	obj, err := client.APICall("host.get_address", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetHostname: Get the hostname field of the given host. */
func (client *XenClient) HostGetHostname(self string) (result string, err error) {
	obj, err := client.APICall("host.get_hostname", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetCpuInfo: Get the cpu_info field of the given host. */
func (client *XenClient) HostGetCpuInfo(self string) (result map[string]string, err error) {
	obj, err := client.APICall("host.get_cpu_info", self)
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

/* GetHostCPUs: Get the host_CPUs field of the given host. */
func (client *XenClient) HostGetHostCPUs(self string) (result []string, err error) {
	obj, err := client.APICall("host.get_host_CPUs", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetPBDs: Get the PBDs field of the given host. */
func (client *XenClient) HostGetPBDs(self string) (result []string, err error) {
	obj, err := client.APICall("host.get_PBDs", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetUpdates: Get the updates field of the given host. */
func (client *XenClient) HostGetUpdates(self string) (result []string, err error) {
	obj, err := client.APICall("host.get_updates", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetPatches: Get the patches field of the given host. */
func (client *XenClient) HostGetPatches(self string) (result []string, err error) {
	obj, err := client.APICall("host.get_patches", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetCrashdumps: Get the crashdumps field of the given host. */
func (client *XenClient) HostGetCrashdumps(self string) (result []string, err error) {
	obj, err := client.APICall("host.get_crashdumps", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetCrashDumpSr: Get the crash_dump_sr field of the given host. */
func (client *XenClient) HostGetCrashDumpSr(self string) (result string, err error) {
	obj, err := client.APICall("host.get_crash_dump_sr", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetSuspendImageSr: Get the suspend_image_sr field of the given host. */
func (client *XenClient) HostGetSuspendImageSr(self string) (result string, err error) {
	obj, err := client.APICall("host.get_suspend_image_sr", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetPIFs: Get the PIFs field of the given host. */
func (client *XenClient) HostGetPIFs(self string) (result []string, err error) {
	obj, err := client.APICall("host.get_PIFs", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetLogging: Get the logging field of the given host. */
func (client *XenClient) HostGetLogging(self string) (result map[string]string, err error) {
	obj, err := client.APICall("host.get_logging", self)
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

/* GetResidentVMs: Get the resident_VMs field of the given host. */
func (client *XenClient) HostGetResidentVMs(self string) (result []string, err error) {
	obj, err := client.APICall("host.get_resident_VMs", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetSupportedBootloaders: Get the supported_bootloaders field of the given host. */
func (client *XenClient) HostGetSupportedBootloaders(self string) (result []string, err error) {
	obj, err := client.APICall("host.get_supported_bootloaders", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetSchedPolicy: Get the sched_policy field of the given host. */
func (client *XenClient) HostGetSchedPolicy(self string) (result string, err error) {
	obj, err := client.APICall("host.get_sched_policy", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetCpuConfiguration: Get the cpu_configuration field of the given host. */
func (client *XenClient) HostGetCpuConfiguration(self string) (result map[string]string, err error) {
	obj, err := client.APICall("host.get_cpu_configuration", self)
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

/* GetCapabilities: Get the capabilities field of the given host. */
func (client *XenClient) HostGetCapabilities(self string) (result []string, err error) {
	obj, err := client.APICall("host.get_capabilities", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetOtherConfig: Get the other_config field of the given host. */
func (client *XenClient) HostGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("host.get_other_config", self)
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

/* GetSoftwareVersion: Get the software_version field of the given host. */
func (client *XenClient) HostGetSoftwareVersion(self string) (result map[string]string, err error) {
	obj, err := client.APICall("host.get_software_version", self)
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

/* GetEnabled: Get the enabled field of the given host. */
func (client *XenClient) HostGetEnabled(self string) (result bool, err error) {
	obj, err := client.APICall("host.get_enabled", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetAPIVersionVendorImplementation: Get the API_version/vendor_implementation field of the given host. */
func (client *XenClient) HostGetAPIVersionVendorImplementation(self string) (result map[string]string, err error) {
	obj, err := client.APICall("host.get_API_version_vendor_implementation", self)
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

/* GetAPIVersionVendor: Get the API_version/vendor field of the given host. */
func (client *XenClient) HostGetAPIVersionVendor(self string) (result string, err error) {
	obj, err := client.APICall("host.get_API_version_vendor", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetAPIVersionMinor: Get the API_version/minor field of the given host. */
func (client *XenClient) HostGetAPIVersionMinor(self string) (result int, err error) {
	obj, err := client.APICall("host.get_API_version_minor", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetAPIVersionMajor: Get the API_version/major field of the given host. */
func (client *XenClient) HostGetAPIVersionMajor(self string) (result int, err error) {
	obj, err := client.APICall("host.get_API_version_major", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetCurrentOperations: Get the current_operations field of the given host. */
func (client *XenClient) HostGetCurrentOperations(self string) (result map[string]HostAllowedOperations, err error) {
	obj, err := client.APICall("host.get_current_operations", self)
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]HostAllowedOperations{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToHostAllowedOperations(obj.String())
		result[key.String()] = mapObj
	}

	return
}

/* GetAllowedOperations: Get the allowed_operations field of the given host. */
func (client *XenClient) HostGetAllowedOperations(self string) (result []HostAllowedOperations, err error) {
	obj, err := client.APICall("host.get_allowed_operations", self)
	if err != nil {
		return
	}

	result = make([]HostAllowedOperations, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = ToHostAllowedOperations(value.(string))
	}

	return
}

/* GetMemoryOverhead: Get the memory/overhead field of the given host. */
func (client *XenClient) HostGetMemoryOverhead(self string) (result int, err error) {
	obj, err := client.APICall("host.get_memory_overhead", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetNameDescription: Get the name/description field of the given host. */
func (client *XenClient) HostGetNameDescription(self string) (result string, err error) {
	obj, err := client.APICall("host.get_name_description", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetNameLabel: Get the name/label field of the given host. */
func (client *XenClient) HostGetNameLabel(self string) (result string, err error) {
	obj, err := client.APICall("host.get_name_label", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given host. */
func (client *XenClient) HostGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("host.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByNameLabel: Get all the host instances with the given label. */
func (client *XenClient) HostGetByNameLabel(label string) (result []string, err error) {
	obj, err := client.APICall("host.get_by_name_label", label)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetByUuid: Get a reference to the host instance with the specified UUID. */
func (client *XenClient) HostGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("host.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given host. */
func (client *XenClient) HostGetRecord(self string) (result Host, err error) {
	obj, err := client.APICall("host.get_record", self)
	if err != nil {
		return
	}

	result = *ToHost(obj.(interface{}))

	return
}

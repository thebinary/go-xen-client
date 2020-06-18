// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w VM.gen.go

package go_xen_client

import (
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/nilshell/xmlrpc"
)

//VM: A virtual machine (or 'guest').
type VM struct {
	Uuid                    string                  // Unique identifier/object reference
	AllowedOperations       []VmOperations          // list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	CurrentOperations       map[string]VmOperations // links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	PowerState              VmPowerState            // Current power state of the machine
	NameLabel               string                  // a human-readable name
	NameDescription         string                  // a notes field containing human-readable description
	UserVersion             int                     // Creators of VMs and templates may store version information here.
	IsATemplate             bool                    // true if this is a template. Template VMs can never be started, they are used only for cloning other VMs
	IsDefaultTemplate       bool                    // true if this is a default template. Default template VMs can never be started or migrated, they are used only for cloning other VMs
	SuspendVDI              string                  // The VDI that a suspend image is stored on. (Only has meaning if VM is currently suspended)
	ResidentOn              string                  // the host the VM is currently resident on
	Affinity                string                  // A host which the VM has some affinity for (or NULL). This is used as a hint to the start call when it decides where to run the VM. Resource constraints may cause the VM to be started elsewhere.
	MemoryOverhead          int                     // Virtualization memory overhead (bytes).
	MemoryTarget            int                     // Dynamically-set memory target (bytes). The value of this field indicates the current target for memory available to this VM.
	MemoryStaticMax         int                     // Statically-set (i.e. absolute) maximum (bytes). The value of this field at VM start time acts as a hard limit of the amount of memory a guest can use. New values only take effect on reboot.
	MemoryDynamicMax        int                     // Dynamic maximum (bytes)
	MemoryDynamicMin        int                     // Dynamic minimum (bytes)
	MemoryStaticMin         int                     // Statically-set (i.e. absolute) mininum (bytes). The value of this field indicates the least amount of memory this VM can boot with without crashing.
	VCPUsParams             map[string]string       // configuration parameters for the selected VCPU policy
	VCPUsMax                int                     // Max number of VCPUs
	VCPUsAtStartup          int                     // Boot number of VCPUs
	ActionsAfterShutdown    OnNormalExit            // action to take after the guest has shutdown itself
	ActionsAfterReboot      OnNormalExit            // action to take after the guest has rebooted itself
	ActionsAfterCrash       OnCrashBehaviour        // action to take if the guest crashes
	Consoles                []string                // virtual console devices
	VIFs                    []string                // virtual network interfaces
	VBDs                    []string                // virtual block devices
	VUSBs                   []string                // vitual usb devices
	CrashDumps              []string                // crash dumps associated with this VM
	VTPMs                   []string                // virtual TPMs
	PVBootloader            string                  // name of or path to bootloader
	PVKernel                string                  // path to the kernel
	PVRamdisk               string                  // path to the initrd
	PVArgs                  string                  // kernel command-line arguments
	PVBootloaderArgs        string                  // miscellaneous arguments for the bootloader
	PVLegacyArgs            string                  // to make Zurich guests boot
	HVMBootPolicy           string                  // HVM boot policy
	HVMBootParams           map[string]string       // HVM boot params
	HVMShadowMultiplier     float32                 // multiplier applied to the amount of shadow that will be made available to the guest
	Platform                map[string]string       // platform-specific configuration
	PCIBus                  string                  // PCI bus path for pass-through devices
	OtherConfig             map[string]string       // additional configuration
	Domid                   int                     // domain ID (if available, -1 otherwise)
	Domarch                 string                  // Domain architecture (if available, null string otherwise)
	LastBootCPUFlags        map[string]string       // describes the CPU flags on which the VM was last booted
	IsControlDomain         bool                    // true if this is a control domain (domain 0 or a driver domain)
	Metrics                 string                  // metrics associated with this VM
	GuestMetrics            string                  // metrics associated with the running guest
	LastBootedRecord        string                  // marshalled value containing VM record at time of last boot, updated dynamically to reflect the runtime state of the domain
	Recommendations         string                  // An XML specification of recommended values and ranges for properties of this VM
	XenstoreData            map[string]string       // data to be inserted into the xenstore tree (/local/domain/<domid>/vm-data) after the VM is created.
	HaAlwaysRun             bool                    // if true then the system will attempt to keep the VM running as much as possible.
	HaRestartPriority       string                  // has possible values: "best-effort" meaning "try to restart this VM if possible but don't consider the Pool to be overcommitted if this is not possible"; "restart" meaning "this VM should be restarted"; "" meaning "do not try to restart this VM"
	IsASnapshot             bool                    // true if this is a snapshot. Snapshotted VMs can never be started, they are used only for cloning other VMs
	SnapshotOf              string                  // Ref pointing to the VM this snapshot is of.
	Snapshots               []string                // List pointing to all the VM snapshots.
	SnapshotTime            time.Time               // Date/time when this snapshot was created.
	TransportableSnapshotId string                  // Transportable ID of the snapshot VM
	Blobs                   map[string]string       // Binary blobs associated with this VM
	Tags                    []string                // user-specified tags for categorization purposes
	BlockedOperations       map[VmOperations]string // List of operations which have been explicitly blocked and an error code
	SnapshotInfo            map[string]string       // Human-readable information concerning this snapshot
	SnapshotMetadata        string                  // Encoded information about the VM's metadata this is a snapshot of
	Parent                  string                  // Ref pointing to the parent of this VM
	Children                []string                // List pointing to all the children of this VM
	BiosStrings             map[string]string       // BIOS strings
	ProtectionPolicy        string                  // Ref pointing to a protection policy for this VM
	IsSnapshotFromVmpp      bool                    // true if this snapshot was created by the protection policy
	SnapshotSchedule        string                  // Ref pointing to a snapshot schedule for this VM
	IsVmssSnapshot          bool                    // true if this snapshot was created by the snapshot schedule
	Appliance               string                  // the appliance to which this VM belongs
	StartDelay              int                     // The delay to wait before proceeding to the next order in the startup sequence (seconds)
	ShutdownDelay           int                     // The delay to wait before proceeding to the next order in the shutdown sequence (seconds)
	Order                   int                     // The point in the startup or shutdown sequence at which this VM will be started
	VGPUs                   []string                // Virtual GPUs
	AttachedPCIs            []string                // Currently passed-through PCI devices
	SuspendSR               string                  // The SR on which a suspend image is stored
	Version                 int                     // The number of times this VM has been recovered
	GenerationId            string                  // Generation ID of the VM
	HardwarePlatformVersion int                     // The host virtual hardware platform version the VM can run on
	HasVendorDevice         bool                    // When an HVM guest starts, this controls the presence of the emulated C000 PCI device which triggers Windows Update to fetch or update PV drivers.
	RequiresReboot          bool                    // Indicates whether a VM requires a reboot in order to update its configuration, e.g. its memory allocation.
	ReferenceLabel          string                  // Textual reference to the template used to create a VM. This can be used by clients in need of an immutable reference to the template since the latter's uuid and name_label may change, for example, after a package installation or upgrade.
	DomainType              DomainType              // The type of domain that will be created when the VM is started
	NVRAM                   map[string]string       // initial value for guest NVRAM (containing UEFI variables, etc). Cannot be changed while the VM is running

}

func FromVMToXml(VM *VM) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] = VM.Uuid

	result["allowed_operations"] = VM.AllowedOperations

	current_operations := make(xmlrpc.Struct)
	for key, value := range VM.CurrentOperations {
		current_operations[key] = value
	}
	result["current_operations"] = current_operations

	result["power_state"] = VM.PowerState.String()

	result["name_label"] = VM.NameLabel

	result["name_description"] = VM.NameDescription

	result["user_version"] = strconv.Itoa(VM.UserVersion)

	result["is_a_template"] = VM.IsATemplate

	result["is_default_template"] = VM.IsDefaultTemplate

	result["suspend_VDI"] = VM.SuspendVDI

	result["resident_on"] = VM.ResidentOn

	result["affinity"] = VM.Affinity

	result["memory_overhead"] = strconv.Itoa(VM.MemoryOverhead)

	result["memory_target"] = strconv.Itoa(VM.MemoryTarget)

	result["memory_static_max"] = strconv.Itoa(VM.MemoryStaticMax)

	result["memory_dynamic_max"] = strconv.Itoa(VM.MemoryDynamicMax)

	result["memory_dynamic_min"] = strconv.Itoa(VM.MemoryDynamicMin)

	result["memory_static_min"] = strconv.Itoa(VM.MemoryStaticMin)

	VCPUs_params := make(xmlrpc.Struct)
	for key, value := range VM.VCPUsParams {
		VCPUs_params[key] = value
	}
	result["VCPUs_params"] = VCPUs_params

	result["VCPUs_max"] = strconv.Itoa(VM.VCPUsMax)

	result["VCPUs_at_startup"] = strconv.Itoa(VM.VCPUsAtStartup)

	result["actions_after_shutdown"] = VM.ActionsAfterShutdown.String()

	result["actions_after_reboot"] = VM.ActionsAfterReboot.String()

	result["actions_after_crash"] = VM.ActionsAfterCrash.String()

	result["consoles"] = VM.Consoles

	result["VIFs"] = VM.VIFs

	result["VBDs"] = VM.VBDs

	result["VUSBs"] = VM.VUSBs

	result["crash_dumps"] = VM.CrashDumps

	result["VTPMs"] = VM.VTPMs

	result["PV_bootloader"] = VM.PVBootloader

	result["PV_kernel"] = VM.PVKernel

	result["PV_ramdisk"] = VM.PVRamdisk

	result["PV_args"] = VM.PVArgs

	result["PV_bootloader_args"] = VM.PVBootloaderArgs

	result["PV_legacy_args"] = VM.PVLegacyArgs

	result["HVM_boot_policy"] = VM.HVMBootPolicy

	HVM_boot_params := make(xmlrpc.Struct)
	for key, value := range VM.HVMBootParams {
		HVM_boot_params[key] = value
	}
	result["HVM_boot_params"] = HVM_boot_params

	result["HVM_shadow_multiplier"] = VM.HVMShadowMultiplier

	platform := make(xmlrpc.Struct)
	for key, value := range VM.Platform {
		platform[key] = value
	}
	result["platform"] = platform

	result["PCI_bus"] = VM.PCIBus

	other_config := make(xmlrpc.Struct)
	for key, value := range VM.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	result["domid"] = strconv.Itoa(VM.Domid)

	result["domarch"] = VM.Domarch

	last_boot_CPU_flags := make(xmlrpc.Struct)
	for key, value := range VM.LastBootCPUFlags {
		last_boot_CPU_flags[key] = value
	}
	result["last_boot_CPU_flags"] = last_boot_CPU_flags

	result["is_control_domain"] = VM.IsControlDomain

	result["metrics"] = VM.Metrics

	result["guest_metrics"] = VM.GuestMetrics

	result["last_booted_record"] = VM.LastBootedRecord

	result["recommendations"] = VM.Recommendations

	xenstore_data := make(xmlrpc.Struct)
	for key, value := range VM.XenstoreData {
		xenstore_data[key] = value
	}
	result["xenstore_data"] = xenstore_data

	result["ha_always_run"] = VM.HaAlwaysRun

	result["ha_restart_priority"] = VM.HaRestartPriority

	result["is_a_snapshot"] = VM.IsASnapshot

	result["snapshot_of"] = VM.SnapshotOf

	result["snapshots"] = VM.Snapshots

	result["snapshot_time"] = VM.SnapshotTime

	result["transportable_snapshot_id"] = VM.TransportableSnapshotId

	blobs := make(xmlrpc.Struct)
	for key, value := range VM.Blobs {
		blobs[key] = value
	}
	result["blobs"] = blobs

	result["tags"] = VM.Tags

	blocked_operations := make(xmlrpc.Struct)
	for key, value := range VM.BlockedOperations {
		blocked_operations[key.String()] = value
	}
	result["blocked_operations"] = blocked_operations

	snapshot_info := make(xmlrpc.Struct)
	for key, value := range VM.SnapshotInfo {
		snapshot_info[key] = value
	}
	result["snapshot_info"] = snapshot_info

	result["snapshot_metadata"] = VM.SnapshotMetadata

	result["parent"] = VM.Parent

	result["children"] = VM.Children

	bios_strings := make(xmlrpc.Struct)
	for key, value := range VM.BiosStrings {
		bios_strings[key] = value
	}
	result["bios_strings"] = bios_strings

	result["protection_policy"] = VM.ProtectionPolicy

	result["is_snapshot_from_vmpp"] = VM.IsSnapshotFromVmpp

	result["snapshot_schedule"] = VM.SnapshotSchedule

	result["is_vmss_snapshot"] = VM.IsVmssSnapshot

	result["appliance"] = VM.Appliance

	result["start_delay"] = strconv.Itoa(VM.StartDelay)

	result["shutdown_delay"] = strconv.Itoa(VM.ShutdownDelay)

	result["order"] = strconv.Itoa(VM.Order)

	result["VGPUs"] = VM.VGPUs

	result["attached_PCIs"] = VM.AttachedPCIs

	result["suspend_SR"] = VM.SuspendSR

	result["version"] = strconv.Itoa(VM.Version)

	result["generation_id"] = VM.GenerationId

	result["hardware_platform_version"] = strconv.Itoa(VM.HardwarePlatformVersion)

	result["has_vendor_device"] = VM.HasVendorDevice

	result["requires_reboot"] = VM.RequiresReboot

	result["reference_label"] = VM.ReferenceLabel

	result["domain_type"] = VM.DomainType.String()

	NVRAM := make(xmlrpc.Struct)
	for key, value := range VM.NVRAM {
		NVRAM[key] = value
	}
	result["NVRAM"] = NVRAM

	return result
}

func ToVM(obj interface{}) (resultObj *VM) {

	objValue := reflect.ValueOf(obj)
	resultObj = &VM{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.Uuid = v
			}
		case "allowed_operations":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.AllowedOperations = make([]VmOperations, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(VmOperations); ok {
						resultObj.AllowedOperations[i] = v
					}
				}
			}
		case "current_operations":

			resultObj.CurrentOperations = map[string]VmOperations{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.CurrentOperations[mapKeyName] = ToVmOperations(v)
				} else {
					resultObj.CurrentOperations[mapKeyName] = 0
				}
			}
		case "power_state":
			if v, ok := keyValue.(VmPowerState); ok {
				resultObj.PowerState = v
			}
		case "name_label":
			if v, ok := keyValue.(string); ok {
				resultObj.NameLabel = v
			}
		case "name_description":
			if v, ok := keyValue.(string); ok {
				resultObj.NameDescription = v
			}
		case "user_version":
			if v, ok := keyValue.(int); ok {
				resultObj.UserVersion = v
			}
		case "is_a_template":
			if v, ok := keyValue.(bool); ok {
				resultObj.IsATemplate = v
			}
		case "is_default_template":
			if v, ok := keyValue.(bool); ok {
				resultObj.IsDefaultTemplate = v
			}
		case "suspend_VDI":
			if v, ok := keyValue.(string); ok {
				resultObj.SuspendVDI = v
			}
		case "resident_on":
			if v, ok := keyValue.(string); ok {
				resultObj.ResidentOn = v
			}
		case "affinity":
			if v, ok := keyValue.(string); ok {
				resultObj.Affinity = v
			}
		case "memory_overhead":
			if v, ok := keyValue.(int); ok {
				resultObj.MemoryOverhead = v
			}
		case "memory_target":
			if v, ok := keyValue.(int); ok {
				resultObj.MemoryTarget = v
			}
		case "memory_static_max":
			if v, ok := keyValue.(int); ok {
				resultObj.MemoryStaticMax = v
			}
		case "memory_dynamic_max":
			if v, ok := keyValue.(int); ok {
				resultObj.MemoryDynamicMax = v
			}
		case "memory_dynamic_min":
			if v, ok := keyValue.(int); ok {
				resultObj.MemoryDynamicMin = v
			}
		case "memory_static_min":
			if v, ok := keyValue.(int); ok {
				resultObj.MemoryStaticMin = v
			}
		case "VCPUs_params":

			resultObj.VCPUsParams = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.VCPUsParams[mapKeyName] = v
				} else {
					resultObj.VCPUsParams[mapKeyName] = ""
				}
			}
		case "VCPUs_max":
			if v, ok := keyValue.(int); ok {
				resultObj.VCPUsMax = v
			}
		case "VCPUs_at_startup":
			if v, ok := keyValue.(int); ok {
				resultObj.VCPUsAtStartup = v
			}
		case "actions_after_shutdown":
			if v, ok := keyValue.(OnNormalExit); ok {
				resultObj.ActionsAfterShutdown = v
			}
		case "actions_after_reboot":
			if v, ok := keyValue.(OnNormalExit); ok {
				resultObj.ActionsAfterReboot = v
			}
		case "actions_after_crash":
			if v, ok := keyValue.(OnCrashBehaviour); ok {
				resultObj.ActionsAfterCrash = v
			}
		case "consoles":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.Consoles = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.Consoles[i] = v
					}
				}
			}
		case "VIFs":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.VIFs = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.VIFs[i] = v
					}
				}
			}
		case "VBDs":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.VBDs = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.VBDs[i] = v
					}
				}
			}
		case "VUSBs":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.VUSBs = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.VUSBs[i] = v
					}
				}
			}
		case "crash_dumps":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.CrashDumps = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.CrashDumps[i] = v
					}
				}
			}
		case "VTPMs":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.VTPMs = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.VTPMs[i] = v
					}
				}
			}
		case "PV_bootloader":
			if v, ok := keyValue.(string); ok {
				resultObj.PVBootloader = v
			}
		case "PV_kernel":
			if v, ok := keyValue.(string); ok {
				resultObj.PVKernel = v
			}
		case "PV_ramdisk":
			if v, ok := keyValue.(string); ok {
				resultObj.PVRamdisk = v
			}
		case "PV_args":
			if v, ok := keyValue.(string); ok {
				resultObj.PVArgs = v
			}
		case "PV_bootloader_args":
			if v, ok := keyValue.(string); ok {
				resultObj.PVBootloaderArgs = v
			}
		case "PV_legacy_args":
			if v, ok := keyValue.(string); ok {
				resultObj.PVLegacyArgs = v
			}
		case "HVM_boot_policy":
			if v, ok := keyValue.(string); ok {
				resultObj.HVMBootPolicy = v
			}
		case "HVM_boot_params":

			resultObj.HVMBootParams = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.HVMBootParams[mapKeyName] = v
				} else {
					resultObj.HVMBootParams[mapKeyName] = ""
				}
			}
		case "HVM_shadow_multiplier":
			if v, ok := keyValue.(float32); ok {
				resultObj.HVMShadowMultiplier = v
			}
		case "platform":

			resultObj.Platform = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.Platform[mapKeyName] = v
				} else {
					resultObj.Platform[mapKeyName] = ""
				}
			}
		case "PCI_bus":
			if v, ok := keyValue.(string); ok {
				resultObj.PCIBus = v
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
		case "domid":
			if v, ok := keyValue.(int); ok {
				resultObj.Domid = v
			}
		case "domarch":
			if v, ok := keyValue.(string); ok {
				resultObj.Domarch = v
			}
		case "last_boot_CPU_flags":

			resultObj.LastBootCPUFlags = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.LastBootCPUFlags[mapKeyName] = v
				} else {
					resultObj.LastBootCPUFlags[mapKeyName] = ""
				}
			}
		case "is_control_domain":
			if v, ok := keyValue.(bool); ok {
				resultObj.IsControlDomain = v
			}
		case "metrics":
			if v, ok := keyValue.(string); ok {
				resultObj.Metrics = v
			}
		case "guest_metrics":
			if v, ok := keyValue.(string); ok {
				resultObj.GuestMetrics = v
			}
		case "last_booted_record":
			if v, ok := keyValue.(string); ok {
				resultObj.LastBootedRecord = v
			}
		case "recommendations":
			if v, ok := keyValue.(string); ok {
				resultObj.Recommendations = v
			}
		case "xenstore_data":

			resultObj.XenstoreData = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.XenstoreData[mapKeyName] = v
				} else {
					resultObj.XenstoreData[mapKeyName] = ""
				}
			}
		case "ha_always_run":
			if v, ok := keyValue.(bool); ok {
				resultObj.HaAlwaysRun = v
			}
		case "ha_restart_priority":
			if v, ok := keyValue.(string); ok {
				resultObj.HaRestartPriority = v
			}
		case "is_a_snapshot":
			if v, ok := keyValue.(bool); ok {
				resultObj.IsASnapshot = v
			}
		case "snapshot_of":
			if v, ok := keyValue.(string); ok {
				resultObj.SnapshotOf = v
			}
		case "snapshots":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.Snapshots = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.Snapshots[i] = v
					}
				}
			}
		case "snapshot_time":
			if v, ok := keyValue.(time.Time); ok {
				resultObj.SnapshotTime = v
			}
		case "transportable_snapshot_id":
			if v, ok := keyValue.(string); ok {
				resultObj.TransportableSnapshotId = v
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
		case "blocked_operations":

		case "snapshot_info":

			resultObj.SnapshotInfo = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.SnapshotInfo[mapKeyName] = v
				} else {
					resultObj.SnapshotInfo[mapKeyName] = ""
				}
			}
		case "snapshot_metadata":
			if v, ok := keyValue.(string); ok {
				resultObj.SnapshotMetadata = v
			}
		case "parent":
			if v, ok := keyValue.(string); ok {
				resultObj.Parent = v
			}
		case "children":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.Children = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.Children[i] = v
					}
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
		case "protection_policy":
			if v, ok := keyValue.(string); ok {
				resultObj.ProtectionPolicy = v
			}
		case "is_snapshot_from_vmpp":
			if v, ok := keyValue.(bool); ok {
				resultObj.IsSnapshotFromVmpp = v
			}
		case "snapshot_schedule":
			if v, ok := keyValue.(string); ok {
				resultObj.SnapshotSchedule = v
			}
		case "is_vmss_snapshot":
			if v, ok := keyValue.(bool); ok {
				resultObj.IsVmssSnapshot = v
			}
		case "appliance":
			if v, ok := keyValue.(string); ok {
				resultObj.Appliance = v
			}
		case "start_delay":
			if v, ok := keyValue.(int); ok {
				resultObj.StartDelay = v
			}
		case "shutdown_delay":
			if v, ok := keyValue.(int); ok {
				resultObj.ShutdownDelay = v
			}
		case "order":
			if v, ok := keyValue.(int); ok {
				resultObj.Order = v
			}
		case "VGPUs":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.VGPUs = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.VGPUs[i] = v
					}
				}
			}
		case "attached_PCIs":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.AttachedPCIs = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.AttachedPCIs[i] = v
					}
				}
			}
		case "suspend_SR":
			if v, ok := keyValue.(string); ok {
				resultObj.SuspendSR = v
			}
		case "version":
			if v, ok := keyValue.(int); ok {
				resultObj.Version = v
			}
		case "generation_id":
			if v, ok := keyValue.(string); ok {
				resultObj.GenerationId = v
			}
		case "hardware_platform_version":
			if v, ok := keyValue.(int); ok {
				resultObj.HardwarePlatformVersion = v
			}
		case "has_vendor_device":
			if v, ok := keyValue.(bool); ok {
				resultObj.HasVendorDevice = v
			}
		case "requires_reboot":
			if v, ok := keyValue.(bool); ok {
				resultObj.RequiresReboot = v
			}
		case "reference_label":
			if v, ok := keyValue.(string); ok {
				resultObj.ReferenceLabel = v
			}
		case "domain_type":
			if v, ok := keyValue.(DomainType); ok {
				resultObj.DomainType = v
			}
		case "NVRAM":

			resultObj.NVRAM = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.NVRAM[mapKeyName] = v
				} else {
					resultObj.NVRAM[mapKeyName] = ""
				}
			}
		}
	}

	return resultObj
}

/* GetAllRecords: Return a map of VM references to VM records for all VMs known to the system. */
func (client *XenClient) VMGetAllRecords() (result map[string]VM, err error) {
	obj, err := client.APICall("VM.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]VM{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToVM(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the VMs known to the system. */
func (client *XenClient) VMGetAll() (result []string, err error) {
	obj, err := client.APICall("VM.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* SetHVMBootPolicy: Set the VM.HVM_boot_policy field of the given VM, which will take effect when it is next started */
func (client *XenClient) VMSetHVMBootPolicy(self string, value string) (err error) {
	_, err = client.APICall("VM.set_HVM_boot_policy", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetDomainType: Set the VM.domain_type field of the given VM, which will take effect when it is next started */
func (client *XenClient) VMSetDomainType(self string, value DomainType) (err error) {
	_, err = client.APICall("VM.set_domain_type", self, value.String())
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetActionsAfterCrash: Sets the actions_after_crash parameter */
func (client *XenClient) VMSetActionsAfterCrash(self string, value OnCrashBehaviour) (err error) {
	_, err = client.APICall("VM.set_actions_after_crash", self, value.String())
	if err != nil {
		return
	}
	// no return result
	return
}

/* Import: Import an XVA from a URI */
func (client *XenClient) VMImport(url string, sr string, full_restore bool, force bool) (result []string, err error) {
	obj, err := client.APICall("VM.import", url, sr, full_restore, force)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* SetHasVendorDevice: Controls whether, when the VM starts in HVM mode, its virtual hardware will include the emulated PCI device for which drivers may be available through Windows Update. Usually this should never be changed on a VM on which Windows has been installed: changing it on such a VM is likely to lead to a crash on next start. */
func (client *XenClient) VMSetHasVendorDevice(self string, value bool) (err error) {
	_, err = client.APICall("VM.set_has_vendor_device", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* CallPlugin: Call an API plugin on this vm */
func (client *XenClient) VMCallPlugin(vm string, plugin string, fn string, args map[string]string) (result string, err error) {
	obj, err := client.APICall("VM.call_plugin", vm, plugin, fn, args)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* QueryServices: Query the system services advertised by this VM and register them. This can only be applied to a system domain. */
func (client *XenClient) VMQueryServices(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VM.query_services", self)
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

/* SetAppliance: Assign this VM to an appliance. */
func (client *XenClient) VMSetAppliance(self string, value string) (err error) {
	_, err = client.APICall("VM.set_appliance", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* ImportConvert: Import using a conversion service. */
func (client *XenClient) VMImportConvert(xtype string, username string, password string, sr string, remote_config map[string]string) (err error) {
	_, err = client.APICall("VM.import_convert", xtype, username, password, sr, remote_config)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Recover: Recover the VM */
func (client *XenClient) VMRecover(self string, session_to string, force bool) (err error) {
	_, err = client.APICall("VM.recover", self, session_to, force)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetSRsRequiredForRecovery: List all the SR's that are required for the VM to be recovered */
func (client *XenClient) VMGetSRsRequiredForRecovery(self string, session_to string) (result []string, err error) {
	obj, err := client.APICall("VM.get_SRs_required_for_recovery", self, session_to)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* AssertCanBeRecovered: Assert whether all SRs required to recover this VM are available. */
func (client *XenClient) VMAssertCanBeRecovered(self string, session_to string) (err error) {
	_, err = client.APICall("VM.assert_can_be_recovered", self, session_to)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetSuspendVDI: Set this VM's suspend VDI, which must be indentical to its current one */
func (client *XenClient) VMSetSuspendVDI(self string, value string) (err error) {
	_, err = client.APICall("VM.set_suspend_VDI", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOrder: Set this VM's boot order */
func (client *XenClient) VMSetOrder(self string, value int) (err error) {
	_, err = client.APICall("VM.set_order", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetShutdownDelay: Set this VM's shutdown delay in seconds */
func (client *XenClient) VMSetShutdownDelay(self string, value int) (err error) {
	_, err = client.APICall("VM.set_shutdown_delay", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetStartDelay: Set this VM's start delay in seconds */
func (client *XenClient) VMSetStartDelay(self string, value int) (err error) {
	_, err = client.APICall("VM.set_start_delay", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetSnapshotSchedule: Set the value of the snapshot schedule field */
func (client *XenClient) VMSetSnapshotSchedule(self string, value string) (err error) {
	_, err = client.APICall("VM.set_snapshot_schedule", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetProtectionPolicy: Set the value of the protection_policy field */
func (client *XenClient) VMSetProtectionPolicy(self string, value string) (err error) {
	_, err = client.APICall("VM.set_protection_policy", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* CopyBiosStrings: Copy the BIOS strings from the given host to this VM */
func (client *XenClient) VMCopyBiosStrings(vm string, host string) (err error) {
	_, err = client.APICall("VM.copy_bios_strings", vm, host)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetBiosStrings: Set custom BIOS strings to this VM. VM will be given a default set of BIOS strings, only some of which can be overridden by the supplied values. Allowed keys are: 'bios-vendor', 'bios-version', 'system-manufacturer', 'system-product-name', 'system-version', 'system-serial-number', 'enclosure-asset-tag', 'baseboard-manufacturer', 'baseboard-product-name', 'baseboard-version', 'baseboard-serial-number', 'baseboard-asset-tag', 'baseboard-location-in-chassis', 'enclosure-asset-tag' */
func (client *XenClient) VMSetBiosStrings(self string, value map[string]string) (err error) {
	_, err = client.APICall("VM.set_bios_strings", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RetrieveWlbRecommendations: Returns mapping of hosts to ratings, indicating the suitability of starting the VM at that location according to wlb. Rating is replaced with an error if the VM cannot boot there. */
func (client *XenClient) VMRetrieveWlbRecommendations(vm string) (result map[string][]string, err error) {
	obj, err := client.APICall("VM.retrieve_wlb_recommendations", vm)
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

/* AssertAgile: Returns an error if the VM is not considered agile e.g. because it is tied to a resource local to a host */
func (client *XenClient) VMAssertAgile(self string) (err error) {
	_, err = client.APICall("VM.assert_agile", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* CreateNewBlob: Create a placeholder for a named binary blob of data that is associated with this VM */
func (client *XenClient) VMCreateNewBlob(vm string, name string, mime_type string, public bool) (result string, err error) {
	obj, err := client.APICall("VM.create_new_blob", vm, name, mime_type, public)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* AssertCanBootHere: Returns an error if the VM could not boot on this host for some reason */
func (client *XenClient) VMAssertCanBootHere(self string, host string) (err error) {
	_, err = client.APICall("VM.assert_can_boot_here", self, host)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetPossibleHosts: Return the list of hosts on which this VM may run. */
func (client *XenClient) VMGetPossibleHosts(vm string) (result []string, err error) {
	obj, err := client.APICall("VM.get_possible_hosts", vm)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetAllowedVIFDevices: Returns a list of the allowed values that a VIF device field can take */
func (client *XenClient) VMGetAllowedVIFDevices(vm string) (result []string, err error) {
	obj, err := client.APICall("VM.get_allowed_VIF_devices", vm)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetAllowedVBDDevices: Returns a list of the allowed values that a VBD device field can take */
func (client *XenClient) VMGetAllowedVBDDevices(vm string) (result []string, err error) {
	obj, err := client.APICall("VM.get_allowed_VBD_devices", vm)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* UpdateAllowedOperations: Recomputes the list of acceptable operations */
func (client *XenClient) VMUpdateAllowedOperations(self string) (err error) {
	_, err = client.APICall("VM.update_allowed_operations", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AssertOperationValid: Check to see whether this operation is acceptable in the current state of the system, raising an error if the operation is invalid for some reason */
func (client *XenClient) VMAssertOperationValid(self string, op VmOperations) (err error) {
	_, err = client.APICall("VM.assert_operation_valid", self, op.String())
	if err != nil {
		return
	}
	// no return result
	return
}

/* ForgetDataSourceArchives: Forget the recorded statistics related to the specified data source */
func (client *XenClient) VMForgetDataSourceArchives(self string, data_source string) (err error) {
	_, err = client.APICall("VM.forget_data_source_archives", self, data_source)
	if err != nil {
		return
	}
	// no return result
	return
}

/* QueryDataSource: Query the latest value of the specified data source */
func (client *XenClient) VMQueryDataSource(self string, data_source string) (result float32, err error) {
	obj, err := client.APICall("VM.query_data_source", self, data_source)
	if err != nil {
		return
	}
	result = obj.(float32)
	return
}

/* RecordDataSource: Start recording the specified data source */
func (client *XenClient) VMRecordDataSource(self string, data_source string) (err error) {
	_, err = client.APICall("VM.record_data_source", self, data_source)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetDataSources:  */
func (client *XenClient) VMGetDataSources(self string) (result []DataSource, err error) {
	obj, err := client.APICall("VM.get_data_sources", self)
	if err != nil {
		return
	}

	result = make([]DataSource, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = *ToDataSource(value)
	}

	return
}

/* GetBootRecord: Returns a record describing the VM's dynamic state, initialised when the VM boots and updated to reflect runtime configuration changes e.g. CPU hotplug */
func (client *XenClient) VMGetBootRecord(self string) (result VM, err error) {
	obj, err := client.APICall("VM.get_boot_record", self)
	if err != nil {
		return
	}

	result = *ToVM(obj.(interface{}))

	return
}

/* AssertCanMigrate: Assert whether a VM can be migrated to the specified destination. */
func (client *XenClient) VMAssertCanMigrate(vm string, dest map[string]string, live bool, vdi_map map[string]string, vif_map map[string]string, options map[string]string, vgpu_map map[string]string) (err error) {
	_, err = client.APICall("VM.assert_can_migrate", vm, dest, live, vdi_map, vif_map, options, vgpu_map)
	if err != nil {
		return
	}
	// no return result
	return
}

/* MigrateSend: Migrate the VM to another host.  This can only be called when the specified VM is in the Running state. */
func (client *XenClient) VMMigrateSend(vm string, dest map[string]string, live bool, vdi_map map[string]string, vif_map map[string]string, options map[string]string, vgpu_map map[string]string) (result string, err error) {
	obj, err := client.APICall("VM.migrate_send", vm, dest, live, vdi_map, vif_map, options, vgpu_map)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* MaximiseMemory: Returns the maximum amount of guest memory which will fit, together with overheads, in the supplied amount of physical memory. If 'exact' is true then an exact calculation is performed using the VM's current settings. If 'exact' is false then a more conservative approximation is used */
func (client *XenClient) VMMaximiseMemory(self string, total int, approximate bool) (result int, err error) {
	obj, err := client.APICall("VM.maximise_memory", self, total, approximate)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* SendTrigger: Send the named trigger to this VM.  This can only be called when the specified VM is in the Running state. */
func (client *XenClient) VMSendTrigger(vm string, trigger string) (err error) {
	_, err = client.APICall("VM.send_trigger", vm, trigger)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SendSysrq: Send the given key as a sysrq to this VM.  The key is specified as a single character (a String of length 1).  This can only be called when the specified VM is in the Running state. */
func (client *XenClient) VMSendSysrq(vm string, key string) (err error) {
	_, err = client.APICall("VM.send_sysrq", vm, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetVCPUsAtStartup: Set the number of startup VCPUs for a halted VM */
func (client *XenClient) VMSetVCPUsAtStartup(self string, value int) (err error) {
	_, err = client.APICall("VM.set_VCPUs_at_startup", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetVCPUsMax: Set the maximum number of VCPUs for a halted VM */
func (client *XenClient) VMSetVCPUsMax(self string, value int) (err error) {
	_, err = client.APICall("VM.set_VCPUs_max", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetShadowMultiplierLive: Set the shadow memory multiplier on a running VM */
func (client *XenClient) VMSetShadowMultiplierLive(self string, multiplier float32) (err error) {
	_, err = client.APICall("VM.set_shadow_multiplier_live", self, multiplier)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetHVMShadowMultiplier: Set the shadow memory multiplier on a halted VM */
func (client *XenClient) VMSetHVMShadowMultiplier(self string, value float32) (err error) {
	_, err = client.APICall("VM.set_HVM_shadow_multiplier", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetCooperative: Return true if the VM is currently 'co-operative' i.e. is expected to reach a balloon target and actually has done */
func (client *XenClient) VMGetCooperative(self string) (result bool, err error) {
	obj, err := client.APICall("VM.get_cooperative", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* WaitMemoryTargetLive: Wait for a running VM to reach its current memory target */
func (client *XenClient) VMWaitMemoryTargetLive(self string) (err error) {
	_, err = client.APICall("VM.wait_memory_target_live", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetMemoryTargetLive: Set the memory target for a running VM */
func (client *XenClient) VMSetMemoryTargetLive(self string, target int) (err error) {
	_, err = client.APICall("VM.set_memory_target_live", self, target)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetMemory: Set the memory allocation of this VM. Sets all of memory_static_max, memory_dynamic_min, and memory_dynamic_max to the given value, and leaves memory_static_min untouched. */
func (client *XenClient) VMSetMemory(self string, value int) (err error) {
	_, err = client.APICall("VM.set_memory", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetMemoryLimits: Set the memory limits of this VM. */
func (client *XenClient) VMSetMemoryLimits(self string, static_min int, static_max int, dynamic_min int, dynamic_max int) (err error) {
	_, err = client.APICall("VM.set_memory_limits", self, static_min, static_max, dynamic_min, dynamic_max)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetMemoryStaticRange: Set the static (ie boot-time) range of virtual memory that the VM is allowed to use. */
func (client *XenClient) VMSetMemoryStaticRange(self string, min int, max int) (err error) {
	_, err = client.APICall("VM.set_memory_static_range", self, min, max)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetMemoryStaticMin: Set the value of the memory_static_min field */
func (client *XenClient) VMSetMemoryStaticMin(self string, value int) (err error) {
	_, err = client.APICall("VM.set_memory_static_min", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetMemoryStaticMax: Set the value of the memory_static_max field */
func (client *XenClient) VMSetMemoryStaticMax(self string, value int) (err error) {
	_, err = client.APICall("VM.set_memory_static_max", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetMemoryDynamicRange: Set the minimum and maximum amounts of physical memory the VM is allowed to use. */
func (client *XenClient) VMSetMemoryDynamicRange(self string, min int, max int) (err error) {
	_, err = client.APICall("VM.set_memory_dynamic_range", self, min, max)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetMemoryDynamicMin: Set the value of the memory_dynamic_min field */
func (client *XenClient) VMSetMemoryDynamicMin(self string, value int) (err error) {
	_, err = client.APICall("VM.set_memory_dynamic_min", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetMemoryDynamicMax: Set the value of the memory_dynamic_max field */
func (client *XenClient) VMSetMemoryDynamicMax(self string, value int) (err error) {
	_, err = client.APICall("VM.set_memory_dynamic_max", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* ComputeMemoryOverhead: Computes the virtualization memory overhead of a VM. */
func (client *XenClient) VMComputeMemoryOverhead(vm string) (result int, err error) {
	obj, err := client.APICall("VM.compute_memory_overhead", vm)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* SetHaAlwaysRun: Set the value of the ha_always_run */
func (client *XenClient) VMSetHaAlwaysRun(self string, value bool) (err error) {
	_, err = client.APICall("VM.set_ha_always_run", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetHaRestartPriority: Set the value of the ha_restart_priority field */
func (client *XenClient) VMSetHaRestartPriority(self string, value string) (err error) {
	_, err = client.APICall("VM.set_ha_restart_priority", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveFromNVRAM:  */
func (client *XenClient) VMRemoveFromNVRAM(self string, key string) (err error) {
	_, err = client.APICall("VM.remove_from_NVRAM", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToNVRAM:  */
func (client *XenClient) VMAddToNVRAM(self string, key string, value string) (err error) {
	_, err = client.APICall("VM.add_to_NVRAM", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetNVRAM:  */
func (client *XenClient) VMSetNVRAM(self string, value map[string]string) (err error) {
	_, err = client.APICall("VM.set_NVRAM", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToVCPUsParamsLive: Add the given key-value pair to VM.VCPUs_params, and apply that value on the running VM */
func (client *XenClient) VMAddToVCPUsParamsLive(self string, key string, value string) (err error) {
	_, err = client.APICall("VM.add_to_VCPUs_params_live", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetVCPUsNumberLive: Set the number of VCPUs for a running VM */
func (client *XenClient) VMSetVCPUsNumberLive(self string, nvcpu int) (err error) {
	_, err = client.APICall("VM.set_VCPUs_number_live", self, nvcpu)
	if err != nil {
		return
	}
	// no return result
	return
}

/* PoolMigrate: Migrate a VM to another Host. */
func (client *XenClient) VMPoolMigrate(vm string, host string, options map[string]string) (err error) {
	_, err = client.APICall("VM.pool_migrate", vm, host, options)
	if err != nil {
		return
	}
	// no return result
	return
}

/* ResumeOn: Awaken the specified VM and resume it on a particular Host.  This can only be called when the specified VM is in the Suspended state. */
func (client *XenClient) VMResumeOn(vm string, host string, start_paused bool, force bool) (err error) {
	_, err = client.APICall("VM.resume_on", vm, host, start_paused, force)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Resume: Awaken the specified VM and resume it.  This can only be called when the specified VM is in the Suspended state. */
func (client *XenClient) VMResume(vm string, start_paused bool, force bool) (err error) {
	_, err = client.APICall("VM.resume", vm, start_paused, force)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Suspend: Suspend the specified VM to disk.  This can only be called when the specified VM is in the Running state. */
func (client *XenClient) VMSuspend(vm string) (err error) {
	_, err = client.APICall("VM.suspend", vm)
	if err != nil {
		return
	}
	// no return result
	return
}

/* HardReboot: Stop executing the specified VM without attempting a clean shutdown and immediately restart the VM. */
func (client *XenClient) VMHardReboot(vm string) (err error) {
	_, err = client.APICall("VM.hard_reboot", vm)
	if err != nil {
		return
	}
	// no return result
	return
}

/* PowerStateReset: Reset the power-state of the VM to halted in the database only. (Used to recover from slave failures in pooling scenarios by resetting the power-states of VMs running on dead slaves to halted.) This is a potentially dangerous operation; use with care. */
func (client *XenClient) VMPowerStateReset(vm string) (err error) {
	_, err = client.APICall("VM.power_state_reset", vm)
	if err != nil {
		return
	}
	// no return result
	return
}

/* HardShutdown: Stop executing the specified VM without attempting a clean shutdown. */
func (client *XenClient) VMHardShutdown(vm string) (err error) {
	_, err = client.APICall("VM.hard_shutdown", vm)
	if err != nil {
		return
	}
	// no return result
	return
}

/* CleanReboot: Attempt to cleanly shutdown the specified VM (Note: this may not be supported---e.g. if a guest agent is not installed). This can only be called when the specified VM is in the Running state. */
func (client *XenClient) VMCleanReboot(vm string) (err error) {
	_, err = client.APICall("VM.clean_reboot", vm)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Shutdown: Attempts to first clean shutdown a VM and if it should fail then perform a hard shutdown on it. */
func (client *XenClient) VMShutdown(vm string) (err error) {
	_, err = client.APICall("VM.shutdown", vm)
	if err != nil {
		return
	}
	// no return result
	return
}

/* CleanShutdown: Attempt to cleanly shutdown the specified VM. (Note: this may not be supported---e.g. if a guest agent is not installed). This can only be called when the specified VM is in the Running state. */
func (client *XenClient) VMCleanShutdown(vm string) (err error) {
	_, err = client.APICall("VM.clean_shutdown", vm)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Unpause: Resume the specified VM. This can only be called when the specified VM is in the Paused state. */
func (client *XenClient) VMUnpause(vm string) (err error) {
	_, err = client.APICall("VM.unpause", vm)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Pause: Pause the specified VM. This can only be called when the specified VM is in the Running state. */
func (client *XenClient) VMPause(vm string) (err error) {
	_, err = client.APICall("VM.pause", vm)
	if err != nil {
		return
	}
	// no return result
	return
}

/* StartOn: Start the specified VM on a particular host.  This function can only be called with the VM is in the Halted State. */
func (client *XenClient) VMStartOn(vm string, host string, start_paused bool, force bool) (err error) {
	_, err = client.APICall("VM.start_on", vm, host, start_paused, force)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Start: Start the specified VM.  This function can only be called with the VM is in the Halted State. */
func (client *XenClient) VMStart(vm string, start_paused bool, force bool) (err error) {
	_, err = client.APICall("VM.start", vm, start_paused, force)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Provision: Inspects the disk configuration contained within the VM's other_config, creates VDIs and VBDs and then executes any applicable post-install script. */
func (client *XenClient) VMProvision(vm string) (err error) {
	_, err = client.APICall("VM.provision", vm)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Checkpoint: Checkpoints the specified VM, making a new VM. Checkpoint automatically exploits the capabilities of the underlying storage repository in which the VM's disk images are stored (e.g. Copy on Write) and saves the memory image as well. */
func (client *XenClient) VMCheckpoint(vm string, new_name string) (result string, err error) {
	obj, err := client.APICall("VM.checkpoint", vm, new_name)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* Revert: Reverts the specified VM to a previous state. */
func (client *XenClient) VMRevert(snapshot string) (err error) {
	_, err = client.APICall("VM.revert", snapshot)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Copy: Copied the specified VM, making a new VM. Unlike clone, copy does not exploits the capabilities of the underlying storage repository in which the VM's disk images are stored. Instead, copy guarantees that the disk images of the newly created VM will be 'full disks' - i.e. not part of a CoW chain.  This function can only be called when the VM is in the Halted State. */
func (client *XenClient) VMCopy(vm string, new_name string, sr string) (result string, err error) {
	obj, err := client.APICall("VM.copy", vm, new_name, sr)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* Clone: Clones the specified VM, making a new VM. Clone automatically exploits the capabilities of the underlying storage repository in which the VM's disk images are stored (e.g. Copy on Write).   This function can only be called when the VM is in the Halted State. */
func (client *XenClient) VMClone(vm string, new_name string) (result string, err error) {
	obj, err := client.APICall("VM.clone", vm, new_name)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* SnapshotWithQuiesce: Snapshots the specified VM with quiesce, making a new VM. Snapshot automatically exploits the capabilities of the underlying storage repository in which the VM's disk images are stored (e.g. Copy on Write). */
func (client *XenClient) VMSnapshotWithQuiesce(vm string, new_name string) (result string, err error) {
	obj, err := client.APICall("VM.snapshot_with_quiesce", vm, new_name)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* Snapshot: Snapshots the specified VM, making a new VM. Snapshot automatically exploits the capabilities of the underlying storage repository in which the VM's disk images are stored (e.g. Copy on Write). */
func (client *XenClient) VMSnapshot(vm string, new_name string) (result string, err error) {
	obj, err := client.APICall("VM.snapshot", vm, new_name)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* SetHardwarePlatformVersion: Set the hardware_platform_version field of the given VM. */
func (client *XenClient) VMSetHardwarePlatformVersion(self string, value int) (err error) {
	_, err = client.APICall("VM.set_hardware_platform_version", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetSuspendSR: Set the suspend_SR field of the given VM. */
func (client *XenClient) VMSetSuspendSR(self string, value string) (err error) {
	_, err = client.APICall("VM.set_suspend_SR", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveFromBlockedOperations: Remove the given key and its corresponding value from the blocked_operations field of the given VM.  If the key is not in that Map, then do nothing. */
func (client *XenClient) VMRemoveFromBlockedOperations(self string, key VmOperations) (err error) {
	_, err = client.APICall("VM.remove_from_blocked_operations", self, key.String())
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToBlockedOperations: Add the given key-value pair to the blocked_operations field of the given VM. */
func (client *XenClient) VMAddToBlockedOperations(self string, key VmOperations, value string) (err error) {
	_, err = client.APICall("VM.add_to_blocked_operations", self, key.String(), value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetBlockedOperations: Set the blocked_operations field of the given VM. */
func (client *XenClient) VMSetBlockedOperations(self string, value map[VmOperations]string) (err error) {
	_, err = client.APICall("VM.set_blocked_operations", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveTags: Remove the given value from the tags field of the given VM.  If the value is not in that Set, then do nothing. */
func (client *XenClient) VMRemoveTags(self string, value string) (err error) {
	_, err = client.APICall("VM.remove_tags", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddTags: Add the given value to the tags field of the given VM.  If the value is already in that Set, then do nothing. */
func (client *XenClient) VMAddTags(self string, value string) (err error) {
	_, err = client.APICall("VM.add_tags", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetTags: Set the tags field of the given VM. */
func (client *XenClient) VMSetTags(self string, value []string) (err error) {
	_, err = client.APICall("VM.set_tags", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveFromXenstoreData: Remove the given key and its corresponding value from the xenstore_data field of the given VM.  If the key is not in that Map, then do nothing. */
func (client *XenClient) VMRemoveFromXenstoreData(self string, key string) (err error) {
	_, err = client.APICall("VM.remove_from_xenstore_data", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToXenstoreData: Add the given key-value pair to the xenstore_data field of the given VM. */
func (client *XenClient) VMAddToXenstoreData(self string, key string, value string) (err error) {
	_, err = client.APICall("VM.add_to_xenstore_data", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetXenstoreData: Set the xenstore_data field of the given VM. */
func (client *XenClient) VMSetXenstoreData(self string, value map[string]string) (err error) {
	_, err = client.APICall("VM.set_xenstore_data", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetRecommendations: Set the recommendations field of the given VM. */
func (client *XenClient) VMSetRecommendations(self string, value string) (err error) {
	_, err = client.APICall("VM.set_recommendations", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given VM.  If the key is not in that Map, then do nothing. */
func (client *XenClient) VMRemoveFromOtherConfig(self string, key string) (err error) {
	_, err = client.APICall("VM.remove_from_other_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToOtherConfig: Add the given key-value pair to the other_config field of the given VM. */
func (client *XenClient) VMAddToOtherConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("VM.add_to_other_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOtherConfig: Set the other_config field of the given VM. */
func (client *XenClient) VMSetOtherConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("VM.set_other_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetPCIBus: Set the PCI_bus field of the given VM. */
func (client *XenClient) VMSetPCIBus(self string, value string) (err error) {
	_, err = client.APICall("VM.set_PCI_bus", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveFromPlatform: Remove the given key and its corresponding value from the platform field of the given VM.  If the key is not in that Map, then do nothing. */
func (client *XenClient) VMRemoveFromPlatform(self string, key string) (err error) {
	_, err = client.APICall("VM.remove_from_platform", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToPlatform: Add the given key-value pair to the platform field of the given VM. */
func (client *XenClient) VMAddToPlatform(self string, key string, value string) (err error) {
	_, err = client.APICall("VM.add_to_platform", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetPlatform: Set the platform field of the given VM. */
func (client *XenClient) VMSetPlatform(self string, value map[string]string) (err error) {
	_, err = client.APICall("VM.set_platform", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveFromHVMBootParams: Remove the given key and its corresponding value from the HVM/boot_params field of the given VM.  If the key is not in that Map, then do nothing. */
func (client *XenClient) VMRemoveFromHVMBootParams(self string, key string) (err error) {
	_, err = client.APICall("VM.remove_from_HVM_boot_params", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToHVMBootParams: Add the given key-value pair to the HVM/boot_params field of the given VM. */
func (client *XenClient) VMAddToHVMBootParams(self string, key string, value string) (err error) {
	_, err = client.APICall("VM.add_to_HVM_boot_params", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetHVMBootParams: Set the HVM/boot_params field of the given VM. */
func (client *XenClient) VMSetHVMBootParams(self string, value map[string]string) (err error) {
	_, err = client.APICall("VM.set_HVM_boot_params", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetPVLegacyArgs: Set the PV/legacy_args field of the given VM. */
func (client *XenClient) VMSetPVLegacyArgs(self string, value string) (err error) {
	_, err = client.APICall("VM.set_PV_legacy_args", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetPVBootloaderArgs: Set the PV/bootloader_args field of the given VM. */
func (client *XenClient) VMSetPVBootloaderArgs(self string, value string) (err error) {
	_, err = client.APICall("VM.set_PV_bootloader_args", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetPVArgs: Set the PV/args field of the given VM. */
func (client *XenClient) VMSetPVArgs(self string, value string) (err error) {
	_, err = client.APICall("VM.set_PV_args", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetPVRamdisk: Set the PV/ramdisk field of the given VM. */
func (client *XenClient) VMSetPVRamdisk(self string, value string) (err error) {
	_, err = client.APICall("VM.set_PV_ramdisk", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetPVKernel: Set the PV/kernel field of the given VM. */
func (client *XenClient) VMSetPVKernel(self string, value string) (err error) {
	_, err = client.APICall("VM.set_PV_kernel", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetPVBootloader: Set the PV/bootloader field of the given VM. */
func (client *XenClient) VMSetPVBootloader(self string, value string) (err error) {
	_, err = client.APICall("VM.set_PV_bootloader", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetActionsAfterReboot: Set the actions/after_reboot field of the given VM. */
func (client *XenClient) VMSetActionsAfterReboot(self string, value OnNormalExit) (err error) {
	_, err = client.APICall("VM.set_actions_after_reboot", self, value.String())
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetActionsAfterShutdown: Set the actions/after_shutdown field of the given VM. */
func (client *XenClient) VMSetActionsAfterShutdown(self string, value OnNormalExit) (err error) {
	_, err = client.APICall("VM.set_actions_after_shutdown", self, value.String())
	if err != nil {
		return
	}
	// no return result
	return
}

/* RemoveFromVCPUsParams: Remove the given key and its corresponding value from the VCPUs/params field of the given VM.  If the key is not in that Map, then do nothing. */
func (client *XenClient) VMRemoveFromVCPUsParams(self string, key string) (err error) {
	_, err = client.APICall("VM.remove_from_VCPUs_params", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToVCPUsParams: Add the given key-value pair to the VCPUs/params field of the given VM. */
func (client *XenClient) VMAddToVCPUsParams(self string, key string, value string) (err error) {
	_, err = client.APICall("VM.add_to_VCPUs_params", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetVCPUsParams: Set the VCPUs/params field of the given VM. */
func (client *XenClient) VMSetVCPUsParams(self string, value map[string]string) (err error) {
	_, err = client.APICall("VM.set_VCPUs_params", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetAffinity: Set the affinity field of the given VM. */
func (client *XenClient) VMSetAffinity(self string, value string) (err error) {
	_, err = client.APICall("VM.set_affinity", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetIsATemplate: Set the is_a_template field of the given VM. */
func (client *XenClient) VMSetIsATemplate(self string, value bool) (err error) {
	_, err = client.APICall("VM.set_is_a_template", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetUserVersion: Set the user_version field of the given VM. */
func (client *XenClient) VMSetUserVersion(self string, value int) (err error) {
	_, err = client.APICall("VM.set_user_version", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetNameDescription: Set the name/description field of the given VM. */
func (client *XenClient) VMSetNameDescription(self string, value string) (err error) {
	_, err = client.APICall("VM.set_name_description", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetNameLabel: Set the name/label field of the given VM. */
func (client *XenClient) VMSetNameLabel(self string, value string) (err error) {
	_, err = client.APICall("VM.set_name_label", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetNVRAM: Get the NVRAM field of the given VM. */
func (client *XenClient) VMGetNVRAM(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VM.get_NVRAM", self)
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

/* GetDomainType: Get the domain_type field of the given VM. */
func (client *XenClient) VMGetDomainType(self string) (result DomainType, err error) {
	obj, err := client.APICall("VM.get_domain_type", self)
	if err != nil {
		return
	}

	result = ToDomainType(obj.(string))

	return
}

/* GetReferenceLabel: Get the reference_label field of the given VM. */
func (client *XenClient) VMGetReferenceLabel(self string) (result string, err error) {
	obj, err := client.APICall("VM.get_reference_label", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRequiresReboot: Get the requires_reboot field of the given VM. */
func (client *XenClient) VMGetRequiresReboot(self string) (result bool, err error) {
	obj, err := client.APICall("VM.get_requires_reboot", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetHasVendorDevice: Get the has_vendor_device field of the given VM. */
func (client *XenClient) VMGetHasVendorDevice(self string) (result bool, err error) {
	obj, err := client.APICall("VM.get_has_vendor_device", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetHardwarePlatformVersion: Get the hardware_platform_version field of the given VM. */
func (client *XenClient) VMGetHardwarePlatformVersion(self string) (result int, err error) {
	obj, err := client.APICall("VM.get_hardware_platform_version", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetGenerationId: Get the generation_id field of the given VM. */
func (client *XenClient) VMGetGenerationId(self string) (result string, err error) {
	obj, err := client.APICall("VM.get_generation_id", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetVersion: Get the version field of the given VM. */
func (client *XenClient) VMGetVersion(self string) (result int, err error) {
	obj, err := client.APICall("VM.get_version", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetSuspendSR: Get the suspend_SR field of the given VM. */
func (client *XenClient) VMGetSuspendSR(self string) (result string, err error) {
	obj, err := client.APICall("VM.get_suspend_SR", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetAttachedPCIs: Get the attached_PCIs field of the given VM. */
func (client *XenClient) VMGetAttachedPCIs(self string) (result []string, err error) {
	obj, err := client.APICall("VM.get_attached_PCIs", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetVGPUs: Get the VGPUs field of the given VM. */
func (client *XenClient) VMGetVGPUs(self string) (result []string, err error) {
	obj, err := client.APICall("VM.get_VGPUs", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetOrder: Get the order field of the given VM. */
func (client *XenClient) VMGetOrder(self string) (result int, err error) {
	obj, err := client.APICall("VM.get_order", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetShutdownDelay: Get the shutdown_delay field of the given VM. */
func (client *XenClient) VMGetShutdownDelay(self string) (result int, err error) {
	obj, err := client.APICall("VM.get_shutdown_delay", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetStartDelay: Get the start_delay field of the given VM. */
func (client *XenClient) VMGetStartDelay(self string) (result int, err error) {
	obj, err := client.APICall("VM.get_start_delay", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetAppliance: Get the appliance field of the given VM. */
func (client *XenClient) VMGetAppliance(self string) (result string, err error) {
	obj, err := client.APICall("VM.get_appliance", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetIsVmssSnapshot: Get the is_vmss_snapshot field of the given VM. */
func (client *XenClient) VMGetIsVmssSnapshot(self string) (result bool, err error) {
	obj, err := client.APICall("VM.get_is_vmss_snapshot", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetSnapshotSchedule: Get the snapshot_schedule field of the given VM. */
func (client *XenClient) VMGetSnapshotSchedule(self string) (result string, err error) {
	obj, err := client.APICall("VM.get_snapshot_schedule", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetIsSnapshotFromVmpp: Get the is_snapshot_from_vmpp field of the given VM. */
func (client *XenClient) VMGetIsSnapshotFromVmpp(self string) (result bool, err error) {
	obj, err := client.APICall("VM.get_is_snapshot_from_vmpp", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetProtectionPolicy: Get the protection_policy field of the given VM. */
func (client *XenClient) VMGetProtectionPolicy(self string) (result string, err error) {
	obj, err := client.APICall("VM.get_protection_policy", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetBiosStrings: Get the bios_strings field of the given VM. */
func (client *XenClient) VMGetBiosStrings(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VM.get_bios_strings", self)
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

/* GetChildren: Get the children field of the given VM. */
func (client *XenClient) VMGetChildren(self string) (result []string, err error) {
	obj, err := client.APICall("VM.get_children", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetParent: Get the parent field of the given VM. */
func (client *XenClient) VMGetParent(self string) (result string, err error) {
	obj, err := client.APICall("VM.get_parent", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetSnapshotMetadata: Get the snapshot_metadata field of the given VM. */
func (client *XenClient) VMGetSnapshotMetadata(self string) (result string, err error) {
	obj, err := client.APICall("VM.get_snapshot_metadata", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetSnapshotInfo: Get the snapshot_info field of the given VM. */
func (client *XenClient) VMGetSnapshotInfo(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VM.get_snapshot_info", self)
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

/* GetBlockedOperations: Get the blocked_operations field of the given VM. */
func (client *XenClient) VMGetBlockedOperations(self string) (result map[VmOperations]string, err error) {
	obj, err := client.APICall("VM.get_blocked_operations", self)
	if err != nil {
		return
	}
	//result conversion not implemented yet
	log.Printf("%+v", obj)
	return
}

/* GetTags: Get the tags field of the given VM. */
func (client *XenClient) VMGetTags(self string) (result []string, err error) {
	obj, err := client.APICall("VM.get_tags", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetBlobs: Get the blobs field of the given VM. */
func (client *XenClient) VMGetBlobs(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VM.get_blobs", self)
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

/* GetTransportableSnapshotId: Get the transportable_snapshot_id field of the given VM. */
func (client *XenClient) VMGetTransportableSnapshotId(self string) (result string, err error) {
	obj, err := client.APICall("VM.get_transportable_snapshot_id", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetSnapshotTime: Get the snapshot_time field of the given VM. */
func (client *XenClient) VMGetSnapshotTime(self string) (result time.Time, err error) {
	obj, err := client.APICall("VM.get_snapshot_time", self)
	if err != nil {
		return
	}
	result = obj.(time.Time)
	return
}

/* GetSnapshots: Get the snapshots field of the given VM. */
func (client *XenClient) VMGetSnapshots(self string) (result []string, err error) {
	obj, err := client.APICall("VM.get_snapshots", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetSnapshotOf: Get the snapshot_of field of the given VM. */
func (client *XenClient) VMGetSnapshotOf(self string) (result string, err error) {
	obj, err := client.APICall("VM.get_snapshot_of", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetIsASnapshot: Get the is_a_snapshot field of the given VM. */
func (client *XenClient) VMGetIsASnapshot(self string) (result bool, err error) {
	obj, err := client.APICall("VM.get_is_a_snapshot", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetHaRestartPriority: Get the ha_restart_priority field of the given VM. */
func (client *XenClient) VMGetHaRestartPriority(self string) (result string, err error) {
	obj, err := client.APICall("VM.get_ha_restart_priority", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetHaAlwaysRun: Get the ha_always_run field of the given VM. */
func (client *XenClient) VMGetHaAlwaysRun(self string) (result bool, err error) {
	obj, err := client.APICall("VM.get_ha_always_run", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetXenstoreData: Get the xenstore_data field of the given VM. */
func (client *XenClient) VMGetXenstoreData(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VM.get_xenstore_data", self)
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

/* GetRecommendations: Get the recommendations field of the given VM. */
func (client *XenClient) VMGetRecommendations(self string) (result string, err error) {
	obj, err := client.APICall("VM.get_recommendations", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetLastBootedRecord: Get the last_booted_record field of the given VM. */
func (client *XenClient) VMGetLastBootedRecord(self string) (result string, err error) {
	obj, err := client.APICall("VM.get_last_booted_record", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetGuestMetrics: Get the guest_metrics field of the given VM. */
func (client *XenClient) VMGetGuestMetrics(self string) (result string, err error) {
	obj, err := client.APICall("VM.get_guest_metrics", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetMetrics: Get the metrics field of the given VM. */
func (client *XenClient) VMGetMetrics(self string) (result string, err error) {
	obj, err := client.APICall("VM.get_metrics", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetIsControlDomain: Get the is_control_domain field of the given VM. */
func (client *XenClient) VMGetIsControlDomain(self string) (result bool, err error) {
	obj, err := client.APICall("VM.get_is_control_domain", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetLastBootCPUFlags: Get the last_boot_CPU_flags field of the given VM. */
func (client *XenClient) VMGetLastBootCPUFlags(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VM.get_last_boot_CPU_flags", self)
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

/* GetDomarch: Get the domarch field of the given VM. */
func (client *XenClient) VMGetDomarch(self string) (result string, err error) {
	obj, err := client.APICall("VM.get_domarch", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetDomid: Get the domid field of the given VM. */
func (client *XenClient) VMGetDomid(self string) (result int, err error) {
	obj, err := client.APICall("VM.get_domid", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetOtherConfig: Get the other_config field of the given VM. */
func (client *XenClient) VMGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VM.get_other_config", self)
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

/* GetPCIBus: Get the PCI_bus field of the given VM. */
func (client *XenClient) VMGetPCIBus(self string) (result string, err error) {
	obj, err := client.APICall("VM.get_PCI_bus", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetPlatform: Get the platform field of the given VM. */
func (client *XenClient) VMGetPlatform(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VM.get_platform", self)
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

/* GetHVMShadowMultiplier: Get the HVM/shadow_multiplier field of the given VM. */
func (client *XenClient) VMGetHVMShadowMultiplier(self string) (result float32, err error) {
	obj, err := client.APICall("VM.get_HVM_shadow_multiplier", self)
	if err != nil {
		return
	}
	result = obj.(float32)
	return
}

/* GetHVMBootParams: Get the HVM/boot_params field of the given VM. */
func (client *XenClient) VMGetHVMBootParams(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VM.get_HVM_boot_params", self)
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

/* GetHVMBootPolicy: Get the HVM/boot_policy field of the given VM. */
func (client *XenClient) VMGetHVMBootPolicy(self string) (result string, err error) {
	obj, err := client.APICall("VM.get_HVM_boot_policy", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetPVLegacyArgs: Get the PV/legacy_args field of the given VM. */
func (client *XenClient) VMGetPVLegacyArgs(self string) (result string, err error) {
	obj, err := client.APICall("VM.get_PV_legacy_args", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetPVBootloaderArgs: Get the PV/bootloader_args field of the given VM. */
func (client *XenClient) VMGetPVBootloaderArgs(self string) (result string, err error) {
	obj, err := client.APICall("VM.get_PV_bootloader_args", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetPVArgs: Get the PV/args field of the given VM. */
func (client *XenClient) VMGetPVArgs(self string) (result string, err error) {
	obj, err := client.APICall("VM.get_PV_args", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetPVRamdisk: Get the PV/ramdisk field of the given VM. */
func (client *XenClient) VMGetPVRamdisk(self string) (result string, err error) {
	obj, err := client.APICall("VM.get_PV_ramdisk", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetPVKernel: Get the PV/kernel field of the given VM. */
func (client *XenClient) VMGetPVKernel(self string) (result string, err error) {
	obj, err := client.APICall("VM.get_PV_kernel", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetPVBootloader: Get the PV/bootloader field of the given VM. */
func (client *XenClient) VMGetPVBootloader(self string) (result string, err error) {
	obj, err := client.APICall("VM.get_PV_bootloader", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetVTPMs: Get the VTPMs field of the given VM. */
func (client *XenClient) VMGetVTPMs(self string) (result []string, err error) {
	obj, err := client.APICall("VM.get_VTPMs", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetCrashDumps: Get the crash_dumps field of the given VM. */
func (client *XenClient) VMGetCrashDumps(self string) (result []string, err error) {
	obj, err := client.APICall("VM.get_crash_dumps", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetVUSBs: Get the VUSBs field of the given VM. */
func (client *XenClient) VMGetVUSBs(self string) (result []string, err error) {
	obj, err := client.APICall("VM.get_VUSBs", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetVBDs: Get the VBDs field of the given VM. */
func (client *XenClient) VMGetVBDs(self string) (result []string, err error) {
	obj, err := client.APICall("VM.get_VBDs", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetVIFs: Get the VIFs field of the given VM. */
func (client *XenClient) VMGetVIFs(self string) (result []string, err error) {
	obj, err := client.APICall("VM.get_VIFs", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetConsoles: Get the consoles field of the given VM. */
func (client *XenClient) VMGetConsoles(self string) (result []string, err error) {
	obj, err := client.APICall("VM.get_consoles", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetActionsAfterCrash: Get the actions/after_crash field of the given VM. */
func (client *XenClient) VMGetActionsAfterCrash(self string) (result OnCrashBehaviour, err error) {
	obj, err := client.APICall("VM.get_actions_after_crash", self)
	if err != nil {
		return
	}

	result = ToOnCrashBehaviour(obj.(string))

	return
}

/* GetActionsAfterReboot: Get the actions/after_reboot field of the given VM. */
func (client *XenClient) VMGetActionsAfterReboot(self string) (result OnNormalExit, err error) {
	obj, err := client.APICall("VM.get_actions_after_reboot", self)
	if err != nil {
		return
	}

	result = ToOnNormalExit(obj.(string))

	return
}

/* GetActionsAfterShutdown: Get the actions/after_shutdown field of the given VM. */
func (client *XenClient) VMGetActionsAfterShutdown(self string) (result OnNormalExit, err error) {
	obj, err := client.APICall("VM.get_actions_after_shutdown", self)
	if err != nil {
		return
	}

	result = ToOnNormalExit(obj.(string))

	return
}

/* GetVCPUsAtStartup: Get the VCPUs/at_startup field of the given VM. */
func (client *XenClient) VMGetVCPUsAtStartup(self string) (result int, err error) {
	obj, err := client.APICall("VM.get_VCPUs_at_startup", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetVCPUsMax: Get the VCPUs/max field of the given VM. */
func (client *XenClient) VMGetVCPUsMax(self string) (result int, err error) {
	obj, err := client.APICall("VM.get_VCPUs_max", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetVCPUsParams: Get the VCPUs/params field of the given VM. */
func (client *XenClient) VMGetVCPUsParams(self string) (result map[string]string, err error) {
	obj, err := client.APICall("VM.get_VCPUs_params", self)
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

/* GetMemoryStaticMin: Get the memory/static_min field of the given VM. */
func (client *XenClient) VMGetMemoryStaticMin(self string) (result int, err error) {
	obj, err := client.APICall("VM.get_memory_static_min", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetMemoryDynamicMin: Get the memory/dynamic_min field of the given VM. */
func (client *XenClient) VMGetMemoryDynamicMin(self string) (result int, err error) {
	obj, err := client.APICall("VM.get_memory_dynamic_min", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetMemoryDynamicMax: Get the memory/dynamic_max field of the given VM. */
func (client *XenClient) VMGetMemoryDynamicMax(self string) (result int, err error) {
	obj, err := client.APICall("VM.get_memory_dynamic_max", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetMemoryStaticMax: Get the memory/static_max field of the given VM. */
func (client *XenClient) VMGetMemoryStaticMax(self string) (result int, err error) {
	obj, err := client.APICall("VM.get_memory_static_max", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetMemoryTarget: Get the memory/target field of the given VM. */
func (client *XenClient) VMGetMemoryTarget(self string) (result int, err error) {
	obj, err := client.APICall("VM.get_memory_target", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetMemoryOverhead: Get the memory/overhead field of the given VM. */
func (client *XenClient) VMGetMemoryOverhead(self string) (result int, err error) {
	obj, err := client.APICall("VM.get_memory_overhead", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetAffinity: Get the affinity field of the given VM. */
func (client *XenClient) VMGetAffinity(self string) (result string, err error) {
	obj, err := client.APICall("VM.get_affinity", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetResidentOn: Get the resident_on field of the given VM. */
func (client *XenClient) VMGetResidentOn(self string) (result string, err error) {
	obj, err := client.APICall("VM.get_resident_on", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetSuspendVDI: Get the suspend_VDI field of the given VM. */
func (client *XenClient) VMGetSuspendVDI(self string) (result string, err error) {
	obj, err := client.APICall("VM.get_suspend_VDI", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetIsDefaultTemplate: Get the is_default_template field of the given VM. */
func (client *XenClient) VMGetIsDefaultTemplate(self string) (result bool, err error) {
	obj, err := client.APICall("VM.get_is_default_template", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetIsATemplate: Get the is_a_template field of the given VM. */
func (client *XenClient) VMGetIsATemplate(self string) (result bool, err error) {
	obj, err := client.APICall("VM.get_is_a_template", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetUserVersion: Get the user_version field of the given VM. */
func (client *XenClient) VMGetUserVersion(self string) (result int, err error) {
	obj, err := client.APICall("VM.get_user_version", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetNameDescription: Get the name/description field of the given VM. */
func (client *XenClient) VMGetNameDescription(self string) (result string, err error) {
	obj, err := client.APICall("VM.get_name_description", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetNameLabel: Get the name/label field of the given VM. */
func (client *XenClient) VMGetNameLabel(self string) (result string, err error) {
	obj, err := client.APICall("VM.get_name_label", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetPowerState: Get the power_state field of the given VM. */
func (client *XenClient) VMGetPowerState(self string) (result VmPowerState, err error) {
	obj, err := client.APICall("VM.get_power_state", self)
	if err != nil {
		return
	}

	result = ToVmPowerState(obj.(string))

	return
}

/* GetCurrentOperations: Get the current_operations field of the given VM. */
func (client *XenClient) VMGetCurrentOperations(self string) (result map[string]VmOperations, err error) {
	obj, err := client.APICall("VM.get_current_operations", self)
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]VmOperations{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToVmOperations(obj.String())
		result[key.String()] = mapObj
	}

	return
}

/* GetAllowedOperations: Get the allowed_operations field of the given VM. */
func (client *XenClient) VMGetAllowedOperations(self string) (result []VmOperations, err error) {
	obj, err := client.APICall("VM.get_allowed_operations", self)
	if err != nil {
		return
	}

	result = make([]VmOperations, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = ToVmOperations(value.(string))
	}

	return
}

/* GetUuid: Get the uuid field of the given VM. */
func (client *XenClient) VMGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("VM.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByNameLabel: Get all the VM instances with the given label. */
func (client *XenClient) VMGetByNameLabel(label string) (result []string, err error) {
	obj, err := client.APICall("VM.get_by_name_label", label)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* Destroy: Destroy the specified VM.  The VM is completely removed from the system.  This function can only be called when the VM is in the Halted State. */
func (client *XenClient) VMDestroy(self string) (err error) {
	_, err = client.APICall("VM.destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Create: NOT RECOMMENDED! VM.clone or VM.copy (or VM.import) is a better choice in almost all situations. The standard way to obtain a new VM is to call VM.clone on a template VM, then call VM.provision on the new clone. Caution: if VM.create is used and then the new VM is attached to a virtual disc that has an operating system already installed, then there is no guarantee that the operating system will boot and run. Any software that calls VM.create on a future version of this API may fail or give unexpected results. For example this could happen if an additional parameter were added to VM.create. VM.create is intended only for use in the automatic creation of the system VM templates. It creates a new VM instance, and returns its handle.
The constructor args are: name_label, name_description, user_version*, is_a_template*, affinity*, memory_target, memory_static_max*, memory_dynamic_max*, memory_dynamic_min*, memory_static_min*, VCPUs_params*, VCPUs_max*, VCPUs_at_startup*, actions_after_shutdown*, actions_after_reboot*, actions_after_crash*, PV_bootloader*, PV_kernel*, PV_ramdisk*, PV_args*, PV_bootloader_args*, PV_legacy_args*, HVM_boot_policy*, HVM_boot_params*, HVM_shadow_multiplier, platform*, PCI_bus*, other_config*, recommendations*, xenstore_data, ha_always_run, ha_restart_priority, tags, blocked_operations, protection_policy, is_snapshot_from_vmpp, snapshot_schedule, is_vmss_snapshot, appliance, start_delay, shutdown_delay, order, suspend_SR, version, generation_id, hardware_platform_version, has_vendor_device, reference_label, domain_type, NVRAM (* = non-optional). */
func (client *XenClient) VMCreate(args VM) (result string, err error) {
	obj, err := client.APICall("VM.create", FromVMToXml(&args))
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the VM instance with the specified UUID. */
func (client *XenClient) VMGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("VM.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given VM. */
func (client *XenClient) VMGetRecord(self string) (result VM, err error) {
	obj, err := client.APICall("VM.get_record", self)
	if err != nil {
		return
	}

	result = *ToVM(obj.(interface{}))

	return
}

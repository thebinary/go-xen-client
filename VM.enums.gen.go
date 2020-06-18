// This is a generated file. DO NOT EDIT manually.

package go_xen_client

//VmPowerState: vm_power_state enum type
type VmPowerState int
const(
   VmPowerStateHalted VmPowerState = iota //VM is offline and not using any resources  
   VmPowerStatePaused //All resources have been allocated but the VM itself is paused and its vCPUs are not running  
   VmPowerStateRunning //Running  
   VmPowerStateSuspended //VM state has been saved to disk and it is nolonger running. Note that disks remain in-use while the VM is suspended.  
)

func (e VmPowerState) String() string {
	switch e {
		
	case 0:
			return "Halted"
		
	case 1:
			return "Paused"
		
	case 2:
			return "Running"
		
	case 3:
			return "Suspended"
		
		default:
			return ""
	}

}

func ToVmPowerState(strValue string) VmPowerState {
	switch strValue { 
		case "Halted":
			return 0
		case "Paused":
			return 1
		case "Running":
			return 2
		case "Suspended":
			return 3
		default:
			return -1
	}
}

//OnNormalExit: on_normal_exit enum type
type OnNormalExit int
const(
   OnNormalExitDestroy OnNormalExit = iota //destroy the VM state  
   OnNormalExitRestart //restart the VM  
)

func (e OnNormalExit) String() string {
	switch e {
		
	case 0:
			return "destroy"
		
	case 1:
			return "restart"
		
		default:
			return ""
	}

}

func ToOnNormalExit(strValue string) OnNormalExit {
	switch strValue { 
		case "destroy":
			return 0
		case "restart":
			return 1
		default:
			return -1
	}
}

//VmOperations: vm_operations enum type
type VmOperations int
const(
   VmOperationsSnapshot VmOperations = iota //refers to the operation "snapshot"  
   VmOperationsClone //refers to the operation "clone"  
   VmOperationsCopy //refers to the operation "copy"  
   VmOperationsCreateTemplate //refers to the operation "create_template"  
   VmOperationsRevert //refers to the operation "revert"  
   VmOperationsCheckpoint //refers to the operation "checkpoint"  
   VmOperationsSnapshotWithQuiesce //refers to the operation "snapshot_with_quiesce"  
   VmOperationsProvision //refers to the operation "provision"  
   VmOperationsStart //refers to the operation "start"  
   VmOperationsStartOn //refers to the operation "start_on"  
   VmOperationsPause //refers to the operation "pause"  
   VmOperationsUnpause //refers to the operation "unpause"  
   VmOperationsCleanShutdown //refers to the operation "clean_shutdown"  
   VmOperationsCleanReboot //refers to the operation "clean_reboot"  
   VmOperationsHardShutdown //refers to the operation "hard_shutdown"  
   VmOperationsPowerStateReset //refers to the operation "power_state_reset"  
   VmOperationsHardReboot //refers to the operation "hard_reboot"  
   VmOperationsSuspend //refers to the operation "suspend"  
   VmOperationsCsvm //refers to the operation "csvm"  
   VmOperationsResume //refers to the operation "resume"  
   VmOperationsResumeOn //refers to the operation "resume_on"  
   VmOperationsPoolMigrate //refers to the operation "pool_migrate"  
   VmOperationsMigrateSend //refers to the operation "migrate_send"  
   VmOperationsGetBootRecord //refers to the operation "get_boot_record"  
   VmOperationsSendSysrq //refers to the operation "send_sysrq"  
   VmOperationsSendTrigger //refers to the operation "send_trigger"  
   VmOperationsQueryServices //refers to the operation "query_services"  
   VmOperationsShutdown //refers to the operation "shutdown"  
   VmOperationsCallPlugin //refers to the operation "call_plugin"  
   VmOperationsChangingMemoryLive //Changing the memory settings  
   VmOperationsAwaitingMemoryLive //Waiting for the memory settings to change  
   VmOperationsChangingDynamicRange //Changing the memory dynamic range  
   VmOperationsChangingStaticRange //Changing the memory static range  
   VmOperationsChangingMemoryLimits //Changing the memory limits  
   VmOperationsChangingShadowMemory //Changing the shadow memory for a halted VM.  
   VmOperationsChangingShadowMemoryLive //Changing the shadow memory for a running VM.  
   VmOperationsChangingVCPUs //Changing VCPU settings for a halted VM.  
   VmOperationsChangingVCPUsLive //Changing VCPU settings for a running VM.  
   VmOperationsChangingNVRAM //Changing NVRAM for a halted VM.  
   VmOperationsAssertOperationValid //  
   VmOperationsDataSourceOp //Add, remove, query or list data sources  
   VmOperationsUpdateAllowedOperations //  
   VmOperationsMakeIntoTemplate //Turning this VM into a template  
   VmOperationsImport //importing a VM from a network stream  
   VmOperationsExport //exporting a VM to a network stream  
   VmOperationsMetadataExport //exporting VM metadata to a network stream  
   VmOperationsReverting //Reverting the VM to a previous snapshotted state  
   VmOperationsDestroy //refers to the act of uninstalling the VM  
)

func (e VmOperations) String() string {
	switch e {
		
	case 0:
			return "snapshot"
		
	case 1:
			return "clone"
		
	case 2:
			return "copy"
		
	case 3:
			return "create_template"
		
	case 4:
			return "revert"
		
	case 5:
			return "checkpoint"
		
	case 6:
			return "snapshot_with_quiesce"
		
	case 7:
			return "provision"
		
	case 8:
			return "start"
		
	case 9:
			return "start_on"
		
	case 10:
			return "pause"
		
	case 11:
			return "unpause"
		
	case 12:
			return "clean_shutdown"
		
	case 13:
			return "clean_reboot"
		
	case 14:
			return "hard_shutdown"
		
	case 15:
			return "power_state_reset"
		
	case 16:
			return "hard_reboot"
		
	case 17:
			return "suspend"
		
	case 18:
			return "csvm"
		
	case 19:
			return "resume"
		
	case 20:
			return "resume_on"
		
	case 21:
			return "pool_migrate"
		
	case 22:
			return "migrate_send"
		
	case 23:
			return "get_boot_record"
		
	case 24:
			return "send_sysrq"
		
	case 25:
			return "send_trigger"
		
	case 26:
			return "query_services"
		
	case 27:
			return "shutdown"
		
	case 28:
			return "call_plugin"
		
	case 29:
			return "changing_memory_live"
		
	case 30:
			return "awaiting_memory_live"
		
	case 31:
			return "changing_dynamic_range"
		
	case 32:
			return "changing_static_range"
		
	case 33:
			return "changing_memory_limits"
		
	case 34:
			return "changing_shadow_memory"
		
	case 35:
			return "changing_shadow_memory_live"
		
	case 36:
			return "changing_VCPUs"
		
	case 37:
			return "changing_VCPUs_live"
		
	case 38:
			return "changing_NVRAM"
		
	case 39:
			return "assert_operation_valid"
		
	case 40:
			return "data_source_op"
		
	case 41:
			return "update_allowed_operations"
		
	case 42:
			return "make_into_template"
		
	case 43:
			return "import"
		
	case 44:
			return "export"
		
	case 45:
			return "metadata_export"
		
	case 46:
			return "reverting"
		
	case 47:
			return "destroy"
		
		default:
			return ""
	}

}

func ToVmOperations(strValue string) VmOperations {
	switch strValue { 
		case "snapshot":
			return 0
		case "clone":
			return 1
		case "copy":
			return 2
		case "create_template":
			return 3
		case "revert":
			return 4
		case "checkpoint":
			return 5
		case "snapshot_with_quiesce":
			return 6
		case "provision":
			return 7
		case "start":
			return 8
		case "start_on":
			return 9
		case "pause":
			return 10
		case "unpause":
			return 11
		case "clean_shutdown":
			return 12
		case "clean_reboot":
			return 13
		case "hard_shutdown":
			return 14
		case "power_state_reset":
			return 15
		case "hard_reboot":
			return 16
		case "suspend":
			return 17
		case "csvm":
			return 18
		case "resume":
			return 19
		case "resume_on":
			return 20
		case "pool_migrate":
			return 21
		case "migrate_send":
			return 22
		case "get_boot_record":
			return 23
		case "send_sysrq":
			return 24
		case "send_trigger":
			return 25
		case "query_services":
			return 26
		case "shutdown":
			return 27
		case "call_plugin":
			return 28
		case "changing_memory_live":
			return 29
		case "awaiting_memory_live":
			return 30
		case "changing_dynamic_range":
			return 31
		case "changing_static_range":
			return 32
		case "changing_memory_limits":
			return 33
		case "changing_shadow_memory":
			return 34
		case "changing_shadow_memory_live":
			return 35
		case "changing_VCPUs":
			return 36
		case "changing_VCPUs_live":
			return 37
		case "changing_NVRAM":
			return 38
		case "assert_operation_valid":
			return 39
		case "data_source_op":
			return 40
		case "update_allowed_operations":
			return 41
		case "make_into_template":
			return 42
		case "import":
			return 43
		case "export":
			return 44
		case "metadata_export":
			return 45
		case "reverting":
			return 46
		case "destroy":
			return 47
		default:
			return -1
	}
}

//OnCrashBehaviour: on_crash_behaviour enum type
type OnCrashBehaviour int
const(
   OnCrashBehaviourDestroy OnCrashBehaviour = iota //destroy the VM state  
   OnCrashBehaviourCoredumpAndDestroy //record a coredump and then destroy the VM state  
   OnCrashBehaviourRestart //restart the VM  
   OnCrashBehaviourCoredumpAndRestart //record a coredump and then restart the VM  
   OnCrashBehaviourPreserve //leave the crashed VM paused  
   OnCrashBehaviourRenameRestart //rename the crashed VM and start a new copy  
)

func (e OnCrashBehaviour) String() string {
	switch e {
		
	case 0:
			return "destroy"
		
	case 1:
			return "coredump_and_destroy"
		
	case 2:
			return "restart"
		
	case 3:
			return "coredump_and_restart"
		
	case 4:
			return "preserve"
		
	case 5:
			return "rename_restart"
		
		default:
			return ""
	}

}

func ToOnCrashBehaviour(strValue string) OnCrashBehaviour {
	switch strValue { 
		case "destroy":
			return 0
		case "coredump_and_destroy":
			return 1
		case "restart":
			return 2
		case "coredump_and_restart":
			return 3
		case "preserve":
			return 4
		case "rename_restart":
			return 5
		default:
			return -1
	}
}

//DomainType: domain_type enum type
type DomainType int
const(
   DomainTypeHvm DomainType = iota //HVM; Fully Virtualised  
   DomainTypePv //PV: Paravirtualised  
   DomainTypePvInPvh //PV inside a PVH container  
   DomainTypeUnspecified //Not specified or unknown domain type  
)

func (e DomainType) String() string {
	switch e {
		
	case 0:
			return "hvm"
		
	case 1:
			return "pv"
		
	case 2:
			return "pv_in_pvh"
		
	case 3:
			return "unspecified"
		
		default:
			return ""
	}

}

func ToDomainType(strValue string) DomainType {
	switch strValue { 
		case "hvm":
			return 0
		case "pv":
			return 1
		case "pv_in_pvh":
			return 2
		case "unspecified":
			return 3
		default:
			return -1
	}
}


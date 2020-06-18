// This is a generated file. DO NOT EDIT manually.

package go_xen_client

//HostAllowedOperations: host_allowed_operations enum type
type HostAllowedOperations int
const(
   HostAllowedOperationsProvision HostAllowedOperations = iota //Indicates this host is able to provision another VM  
   HostAllowedOperationsEvacuate //Indicates this host is evacuating  
   HostAllowedOperationsShutdown //Indicates this host is in the process of shutting itself down  
   HostAllowedOperationsReboot //Indicates this host is in the process of rebooting  
   HostAllowedOperationsPowerOn //Indicates this host is in the process of being powered on  
   HostAllowedOperationsVmStart //This host is starting a VM  
   HostAllowedOperationsVmResume //This host is resuming a VM  
   HostAllowedOperationsVmMigrate //This host is the migration target of a VM  
)

func (e HostAllowedOperations) String() string {
	switch e {
		
	case 0:
			return "provision"
		
	case 1:
			return "evacuate"
		
	case 2:
			return "shutdown"
		
	case 3:
			return "reboot"
		
	case 4:
			return "power_on"
		
	case 5:
			return "vm_start"
		
	case 6:
			return "vm_resume"
		
	case 7:
			return "vm_migrate"
		
		default:
			return ""
	}

}

func ToHostAllowedOperations(strValue string) HostAllowedOperations {
	switch strValue { 
		case "provision":
			return 0
		case "evacuate":
			return 1
		case "shutdown":
			return 2
		case "reboot":
			return 3
		case "power_on":
			return 4
		case "vm_start":
			return 5
		case "vm_resume":
			return 6
		case "vm_migrate":
			return 7
		default:
			return -1
	}
}

//HostDisplay: host_display enum type
type HostDisplay int
const(
   HostDisplayEnabled HostDisplay = iota //This host is outputting its console to a physical display device  
   HostDisplayDisableOnReboot //The host will stop outputting its console to a physical display device on next boot  
   HostDisplayDisabled //This host is not outputting its console to a physical display device  
   HostDisplayEnableOnReboot //The host will start outputting its console to a physical display device on next boot  
)

func (e HostDisplay) String() string {
	switch e {
		
	case 0:
			return "enabled"
		
	case 1:
			return "disable_on_reboot"
		
	case 2:
			return "disabled"
		
	case 3:
			return "enable_on_reboot"
		
		default:
			return ""
	}

}

func ToHostDisplay(strValue string) HostDisplay {
	switch strValue { 
		case "enabled":
			return 0
		case "disable_on_reboot":
			return 1
		case "disabled":
			return 2
		case "enable_on_reboot":
			return 3
		default:
			return -1
	}
}


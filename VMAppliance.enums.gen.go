// This is a generated file. DO NOT EDIT manually.

package go_xen_client

//VmApplianceOperation: vm_appliance_operation enum type
type VmApplianceOperation int
const(
   VmApplianceOperationStart VmApplianceOperation = iota //Start  
   VmApplianceOperationCleanShutdown //Clean shutdown  
   VmApplianceOperationHardShutdown //Hard shutdown  
   VmApplianceOperationShutdown //Shutdown  
)

func (e VmApplianceOperation) String() string {
	switch e {
		
	case 0:
			return "start"
		
	case 1:
			return "clean_shutdown"
		
	case 2:
			return "hard_shutdown"
		
	case 3:
			return "shutdown"
		
		default:
			return ""
	}

}

func ToVmApplianceOperation(strValue string) VmApplianceOperation {
	switch strValue { 
		case "start":
			return 0
		case "clean_shutdown":
			return 1
		case "hard_shutdown":
			return 2
		case "shutdown":
			return 3
		default:
			return -1
	}
}


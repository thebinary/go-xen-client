// This is a generated file. DO NOT EDIT manually.

package go_xen_client

//PgpuDom0Access: pgpu_dom0_access enum type
type PgpuDom0Access int
const(
   PgpuDom0AccessEnabled PgpuDom0Access = iota //dom0 can access this device as normal  
   PgpuDom0AccessDisableOnReboot //On host reboot dom0 will be blocked from accessing this device  
   PgpuDom0AccessDisabled //dom0 cannot access this device  
   PgpuDom0AccessEnableOnReboot //On host reboot dom0 will be allowed to access this device  
)

func (e PgpuDom0Access) String() string {
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

func ToPgpuDom0Access(strValue string) PgpuDom0Access {
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


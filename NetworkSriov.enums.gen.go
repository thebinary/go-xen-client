// This is a generated file. DO NOT EDIT manually.

package go_xen_client

//SriovConfigurationMode: sriov_configuration_mode enum type
type SriovConfigurationMode int
const(
   SriovConfigurationModeSysfs SriovConfigurationMode = iota //Configure network sriov by sysfs, do not need reboot  
   SriovConfigurationModeModprobe //Configure network sriov by modprobe, need reboot  
   SriovConfigurationModeUnknown //Unknown mode  
)

func (e SriovConfigurationMode) String() string {
	switch e {
		
	case 0:
			return "sysfs"
		
	case 1:
			return "modprobe"
		
	case 2:
			return "unknown"
		
		default:
			return ""
	}

}

func ToSriovConfigurationMode(strValue string) SriovConfigurationMode {
	switch strValue { 
		case "sysfs":
			return 0
		case "modprobe":
			return 1
		case "unknown":
			return 2
		default:
			return -1
	}
}


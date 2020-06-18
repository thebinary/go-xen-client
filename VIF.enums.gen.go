// This is a generated file. DO NOT EDIT manually.

package go_xen_client

//VifOperations: vif_operations enum type
type VifOperations int
const(
   VifOperationsAttach VifOperations = iota //Attempting to attach this VIF to a VM  
   VifOperationsPlug //Attempting to hotplug this VIF  
   VifOperationsUnplug //Attempting to hot unplug this VIF  
)

func (e VifOperations) String() string {
	switch e {
		
	case 0:
			return "attach"
		
	case 1:
			return "plug"
		
	case 2:
			return "unplug"
		
		default:
			return ""
	}

}

func ToVifOperations(strValue string) VifOperations {
	switch strValue { 
		case "attach":
			return 0
		case "plug":
			return 1
		case "unplug":
			return 2
		default:
			return -1
	}
}

//VifLockingMode: vif_locking_mode enum type
type VifLockingMode int
const(
   VifLockingModeNetworkDefault VifLockingMode = iota //No specific configuration set - default network policy applies  
   VifLockingModeLocked //Only traffic to a specific MAC and a list of IPv4 or IPv6 addresses is permitted  
   VifLockingModeUnlocked //All traffic is permitted  
   VifLockingModeDisabled //No traffic is permitted  
)

func (e VifLockingMode) String() string {
	switch e {
		
	case 0:
			return "network_default"
		
	case 1:
			return "locked"
		
	case 2:
			return "unlocked"
		
	case 3:
			return "disabled"
		
		default:
			return ""
	}

}

func ToVifLockingMode(strValue string) VifLockingMode {
	switch strValue { 
		case "network_default":
			return 0
		case "locked":
			return 1
		case "unlocked":
			return 2
		case "disabled":
			return 3
		default:
			return -1
	}
}

//VifIpv4ConfigurationMode: vif_ipv4_configuration_mode enum type
type VifIpv4ConfigurationMode int
const(
   VifIpv4ConfigurationModeNone VifIpv4ConfigurationMode = iota //Follow the default IPv4 configuration of the guest (this is guest-dependent)  
   VifIpv4ConfigurationModeStatic //Static IPv4 address configuration  
)

func (e VifIpv4ConfigurationMode) String() string {
	switch e {
		
	case 0:
			return "None"
		
	case 1:
			return "Static"
		
		default:
			return ""
	}

}

func ToVifIpv4ConfigurationMode(strValue string) VifIpv4ConfigurationMode {
	switch strValue { 
		case "None":
			return 0
		case "Static":
			return 1
		default:
			return -1
	}
}

//VifIpv6ConfigurationMode: vif_ipv6_configuration_mode enum type
type VifIpv6ConfigurationMode int
const(
   VifIpv6ConfigurationModeNone VifIpv6ConfigurationMode = iota //Follow the default IPv6 configuration of the guest (this is guest-dependent)  
   VifIpv6ConfigurationModeStatic //Static IPv6 address configuration  
)

func (e VifIpv6ConfigurationMode) String() string {
	switch e {
		
	case 0:
			return "None"
		
	case 1:
			return "Static"
		
		default:
			return ""
	}

}

func ToVifIpv6ConfigurationMode(strValue string) VifIpv6ConfigurationMode {
	switch strValue { 
		case "None":
			return 0
		case "Static":
			return 1
		default:
			return -1
	}
}


// This is a generated file. DO NOT EDIT manually.

package go_xen_client

//PifIgmpStatus: pif_igmp_status enum type
type PifIgmpStatus int
const(
   PifIgmpStatusEnabled PifIgmpStatus = iota //IGMP Snooping is enabled in the corresponding backend bridge.'  
   PifIgmpStatusDisabled //IGMP Snooping is disabled in the corresponding backend bridge.'  
   PifIgmpStatusUnknown //IGMP snooping status is unknown. If this is a VLAN master, then please consult the underlying VLAN slave PIF.  
)

func (e PifIgmpStatus) String() string {
	switch e {
		
	case 0:
			return "enabled"
		
	case 1:
			return "disabled"
		
	case 2:
			return "unknown"
		
		default:
			return ""
	}

}

func ToPifIgmpStatus(strValue string) PifIgmpStatus {
	switch strValue { 
		case "enabled":
			return 0
		case "disabled":
			return 1
		case "unknown":
			return 2
		default:
			return -1
	}
}

//IpConfigurationMode: ip_configuration_mode enum type
type IpConfigurationMode int
const(
   IpConfigurationModeNone IpConfigurationMode = iota //Do not acquire an IP address  
   IpConfigurationModeDHCP //Acquire an IP address by DHCP  
   IpConfigurationModeStatic //Static IP address configuration  
)

func (e IpConfigurationMode) String() string {
	switch e {
		
	case 0:
			return "None"
		
	case 1:
			return "DHCP"
		
	case 2:
			return "Static"
		
		default:
			return ""
	}

}

func ToIpConfigurationMode(strValue string) IpConfigurationMode {
	switch strValue { 
		case "None":
			return 0
		case "DHCP":
			return 1
		case "Static":
			return 2
		default:
			return -1
	}
}

//Ipv6ConfigurationMode: ipv6_configuration_mode enum type
type Ipv6ConfigurationMode int
const(
   Ipv6ConfigurationModeNone Ipv6ConfigurationMode = iota //Do not acquire an IPv6 address  
   Ipv6ConfigurationModeDHCP //Acquire an IPv6 address by DHCP  
   Ipv6ConfigurationModeStatic //Static IPv6 address configuration  
   Ipv6ConfigurationModeAutoconf //Router assigned prefix delegation IPv6 allocation  
)

func (e Ipv6ConfigurationMode) String() string {
	switch e {
		
	case 0:
			return "None"
		
	case 1:
			return "DHCP"
		
	case 2:
			return "Static"
		
	case 3:
			return "Autoconf"
		
		default:
			return ""
	}

}

func ToIpv6ConfigurationMode(strValue string) Ipv6ConfigurationMode {
	switch strValue { 
		case "None":
			return 0
		case "DHCP":
			return 1
		case "Static":
			return 2
		case "Autoconf":
			return 3
		default:
			return -1
	}
}

//PrimaryAddressType: primary_address_type enum type
type PrimaryAddressType int
const(
   PrimaryAddressTypeIPv4 PrimaryAddressType = iota //Primary address is the IPv4 address  
   PrimaryAddressTypeIPv6 //Primary address is the IPv6 address  
)

func (e PrimaryAddressType) String() string {
	switch e {
		
	case 0:
			return "IPv4"
		
	case 1:
			return "IPv6"
		
		default:
			return ""
	}

}

func ToPrimaryAddressType(strValue string) PrimaryAddressType {
	switch strValue { 
		case "IPv4":
			return 0
		case "IPv6":
			return 1
		default:
			return -1
	}
}


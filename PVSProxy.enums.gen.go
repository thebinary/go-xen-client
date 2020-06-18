// This is a generated file. DO NOT EDIT manually.

package go_xen_client

//PvsProxyStatus: pvs_proxy_status enum type
type PvsProxyStatus int
const(
   PvsProxyStatusStopped PvsProxyStatus = iota //The proxy is not currently running  
   PvsProxyStatusInitialised //The proxy is setup but has not yet cached anything  
   PvsProxyStatusCaching //The proxy is currently caching data  
   PvsProxyStatusIncompatibleWriteCacheMode //The PVS device is configured to use an incompatible write-cache mode  
   PvsProxyStatusIncompatibleProtocolVersion //The PVS protocol in use is not compatible with the PVS proxy  
)

func (e PvsProxyStatus) String() string {
	switch e {
		
	case 0:
			return "stopped"
		
	case 1:
			return "initialised"
		
	case 2:
			return "caching"
		
	case 3:
			return "incompatible_write_cache_mode"
		
	case 4:
			return "incompatible_protocol_version"
		
		default:
			return ""
	}

}

func ToPvsProxyStatus(strValue string) PvsProxyStatus {
	switch strValue { 
		case "stopped":
			return 0
		case "initialised":
			return 1
		case "caching":
			return 2
		case "incompatible_write_cache_mode":
			return 3
		case "incompatible_protocol_version":
			return 4
		default:
			return -1
	}
}


// This is a generated file. DO NOT EDIT manually.

package go_xen_client

//NetworkOperations: network_operations enum type
type NetworkOperations int
const(
   NetworkOperationsAttaching NetworkOperations = iota //Indicates this network is attaching to a VIF or PIF  
)

func (e NetworkOperations) String() string {
	switch e {
		
	case 0:
			return "attaching"
		
		default:
			return ""
	}

}

func ToNetworkOperations(strValue string) NetworkOperations {
	switch strValue { 
		case "attaching":
			return 0
		default:
			return -1
	}
}

//NetworkDefaultLockingMode: network_default_locking_mode enum type
type NetworkDefaultLockingMode int
const(
   NetworkDefaultLockingModeUnlocked NetworkDefaultLockingMode = iota //Treat all VIFs on this network with locking_mode = 'default' as if they have locking_mode = 'unlocked'  
   NetworkDefaultLockingModeDisabled //Treat all VIFs on this network with locking_mode = 'default' as if they have locking_mode = 'disabled'  
)

func (e NetworkDefaultLockingMode) String() string {
	switch e {
		
	case 0:
			return "unlocked"
		
	case 1:
			return "disabled"
		
		default:
			return ""
	}

}

func ToNetworkDefaultLockingMode(strValue string) NetworkDefaultLockingMode {
	switch strValue { 
		case "unlocked":
			return 0
		case "disabled":
			return 1
		default:
			return -1
	}
}

//NetworkPurpose: network_purpose enum type
type NetworkPurpose int
const(
   NetworkPurposeNbd NetworkPurpose = iota //Network Block Device service using TLS  
   NetworkPurposeInsecureNbd //Network Block Device service without integrity or confidentiality: NOT RECOMMENDED  
)

func (e NetworkPurpose) String() string {
	switch e {
		
	case 0:
			return "nbd"
		
	case 1:
			return "insecure_nbd"
		
		default:
			return ""
	}

}

func ToNetworkPurpose(strValue string) NetworkPurpose {
	switch strValue { 
		case "nbd":
			return 0
		case "insecure_nbd":
			return 1
		default:
			return -1
	}
}


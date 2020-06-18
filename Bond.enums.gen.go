// This is a generated file. DO NOT EDIT manually.

package go_xen_client

//BondMode: bond_mode enum type
type BondMode int
const(
   BondModeBalanceSlb BondMode = iota //Source-level balancing  
   BondModeActiveBackup //Active/passive bonding: only one NIC is carrying traffic  
   BondModeLacp //Link aggregation control protocol  
)

func (e BondMode) String() string {
	switch e {
		
	case 0:
			return "balance-slb"
		
	case 1:
			return "active-backup"
		
	case 2:
			return "lacp"
		
		default:
			return ""
	}

}

func ToBondMode(strValue string) BondMode {
	switch strValue { 
		case "balance-slb":
			return 0
		case "active-backup":
			return 1
		case "lacp":
			return 2
		default:
			return -1
	}
}


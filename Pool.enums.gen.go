// This is a generated file. DO NOT EDIT manually.

package go_xen_client

//PoolAllowedOperations: pool_allowed_operations enum type
type PoolAllowedOperations int
const(
   PoolAllowedOperationsHaEnable PoolAllowedOperations = iota //Indicates this pool is in the process of enabling HA  
   PoolAllowedOperationsHaDisable //Indicates this pool is in the process of disabling HA  
   PoolAllowedOperationsClusterCreate //Indicates this pool is in the process of creating a cluster  
)

func (e PoolAllowedOperations) String() string {
	switch e {
		
	case 0:
			return "ha_enable"
		
	case 1:
			return "ha_disable"
		
	case 2:
			return "cluster_create"
		
		default:
			return ""
	}

}

func ToPoolAllowedOperations(strValue string) PoolAllowedOperations {
	switch strValue { 
		case "ha_enable":
			return 0
		case "ha_disable":
			return 1
		case "cluster_create":
			return 2
		default:
			return -1
	}
}


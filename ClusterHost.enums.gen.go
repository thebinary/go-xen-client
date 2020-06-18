// This is a generated file. DO NOT EDIT manually.

package go_xen_client

//ClusterHostOperation: cluster_host_operation enum type
type ClusterHostOperation int
const(
   ClusterHostOperationEnable ClusterHostOperation = iota //enabling cluster membership on a particular host  
   ClusterHostOperationDisable //disabling cluster membership on a particular host  
   ClusterHostOperationDestroy //completely destroying a cluster host  
)

func (e ClusterHostOperation) String() string {
	switch e {
		
	case 0:
			return "enable"
		
	case 1:
			return "disable"
		
	case 2:
			return "destroy"
		
		default:
			return ""
	}

}

func ToClusterHostOperation(strValue string) ClusterHostOperation {
	switch strValue { 
		case "enable":
			return 0
		case "disable":
			return 1
		case "destroy":
			return 2
		default:
			return -1
	}
}


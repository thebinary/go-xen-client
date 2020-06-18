// This is a generated file. DO NOT EDIT manually.

package go_xen_client

//ClusterOperation: cluster_operation enum type
type ClusterOperation int
const(
   ClusterOperationAdd ClusterOperation = iota //adding a new member to the cluster  
   ClusterOperationRemove //removing a member from the cluster  
   ClusterOperationEnable //enabling any cluster member  
   ClusterOperationDisable //disabling any cluster member  
   ClusterOperationDestroy //completely destroying a cluster  
)

func (e ClusterOperation) String() string {
	switch e {
		
	case 0:
			return "add"
		
	case 1:
			return "remove"
		
	case 2:
			return "enable"
		
	case 3:
			return "disable"
		
	case 4:
			return "destroy"
		
		default:
			return ""
	}

}

func ToClusterOperation(strValue string) ClusterOperation {
	switch strValue { 
		case "add":
			return 0
		case "remove":
			return 1
		case "enable":
			return 2
		case "disable":
			return 3
		case "destroy":
			return 4
		default:
			return -1
	}
}


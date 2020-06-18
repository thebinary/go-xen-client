// This is a generated file. DO NOT EDIT manually.

package go_xen_client

//TaskAllowedOperations: task_allowed_operations enum type
type TaskAllowedOperations int
const(
   TaskAllowedOperationsCancel TaskAllowedOperations = iota //refers to the operation "cancel"  
   TaskAllowedOperationsDestroy //refers to the operation "destroy"  
)

func (e TaskAllowedOperations) String() string {
	switch e {
		
	case 0:
			return "cancel"
		
	case 1:
			return "destroy"
		
		default:
			return ""
	}

}

func ToTaskAllowedOperations(strValue string) TaskAllowedOperations {
	switch strValue { 
		case "cancel":
			return 0
		case "destroy":
			return 1
		default:
			return -1
	}
}

//TaskStatusType: task_status_type enum type
type TaskStatusType int
const(
   TaskStatusTypePending TaskStatusType = iota //task is in progress  
   TaskStatusTypeSuccess //task was completed successfully  
   TaskStatusTypeFailure //task has failed  
   TaskStatusTypeCancelling //task is being cancelled  
   TaskStatusTypeCancelled //task has been cancelled  
)

func (e TaskStatusType) String() string {
	switch e {
		
	case 0:
			return "pending"
		
	case 1:
			return "success"
		
	case 2:
			return "failure"
		
	case 3:
			return "cancelling"
		
	case 4:
			return "cancelled"
		
		default:
			return ""
	}

}

func ToTaskStatusType(strValue string) TaskStatusType {
	switch strValue { 
		case "pending":
			return 0
		case "success":
			return 1
		case "failure":
			return 2
		case "cancelling":
			return 3
		case "cancelled":
			return 4
		default:
			return -1
	}
}


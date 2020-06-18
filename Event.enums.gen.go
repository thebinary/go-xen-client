// This is a generated file. DO NOT EDIT manually.

package go_xen_client

//EventOperation: event_operation enum type
type EventOperation int
const(
   EventOperationAdd EventOperation = iota //An object has been created  
   EventOperationDel //An object has been deleted  
   EventOperationMod //An object has been modified  
)

func (e EventOperation) String() string {
	switch e {
		
	case 0:
			return "add"
		
	case 1:
			return "del"
		
	case 2:
			return "mod"
		
		default:
			return ""
	}

}

func ToEventOperation(strValue string) EventOperation {
	switch strValue { 
		case "add":
			return 0
		case "del":
			return 1
		case "mod":
			return 2
		default:
			return -1
	}
}


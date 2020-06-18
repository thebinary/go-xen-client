// This is a generated file. DO NOT EDIT manually.

package go_xen_client

//VusbOperations: vusb_operations enum type
type VusbOperations int
const(
   VusbOperationsAttach VusbOperations = iota //Attempting to attach this VUSB to a VM  
   VusbOperationsPlug //Attempting to plug this VUSB into a VM  
   VusbOperationsUnplug //Attempting to hot unplug this VUSB  
)

func (e VusbOperations) String() string {
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

func ToVusbOperations(strValue string) VusbOperations {
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


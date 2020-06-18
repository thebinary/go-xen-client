// This is a generated file. DO NOT EDIT manually.

package go_xen_client

//Cls: cls enum type
type Cls int
const(
   ClsVM Cls = iota //VM  
   ClsHost //Host  
   ClsSR //SR  
   ClsPool //Pool  
   ClsVMPP //VMPP  
   ClsVMSS //VMSS  
   ClsPVSProxy //PVS_proxy  
   ClsVDI //VDI  
)

func (e Cls) String() string {
	switch e {
		
	case 0:
			return "VM"
		
	case 1:
			return "Host"
		
	case 2:
			return "SR"
		
	case 3:
			return "Pool"
		
	case 4:
			return "VMPP"
		
	case 5:
			return "VMSS"
		
	case 6:
			return "PVS_proxy"
		
	case 7:
			return "VDI"
		
		default:
			return ""
	}

}

func ToCls(strValue string) Cls {
	switch strValue { 
		case "VM":
			return 0
		case "Host":
			return 1
		case "SR":
			return 2
		case "Pool":
			return 3
		case "VMPP":
			return 4
		case "VMSS":
			return 5
		case "PVS_proxy":
			return 6
		case "VDI":
			return 7
		default:
			return -1
	}
}


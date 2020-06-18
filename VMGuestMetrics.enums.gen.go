// This is a generated file. DO NOT EDIT manually.

package go_xen_client

//TristateType: tristate_type enum type
type TristateType int
const(
   TristateTypeYes TristateType = iota //Known to be true  
   TristateTypeNo //Known to be false  
   TristateTypeUnspecified //Unknown or unspecified  
)

func (e TristateType) String() string {
	switch e {
		
	case 0:
			return "yes"
		
	case 1:
			return "no"
		
	case 2:
			return "unspecified"
		
		default:
			return ""
	}

}

func ToTristateType(strValue string) TristateType {
	switch strValue { 
		case "yes":
			return 0
		case "no":
			return 1
		case "unspecified":
			return 2
		default:
			return -1
	}
}


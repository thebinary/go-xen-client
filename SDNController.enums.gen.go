// This is a generated file. DO NOT EDIT manually.

package go_xen_client

//SdnControllerProtocol: sdn_controller_protocol enum type
type SdnControllerProtocol int
const(
   SdnControllerProtocolSsl SdnControllerProtocol = iota //Active ssl connection  
   SdnControllerProtocolPssl //Passive ssl connection  
)

func (e SdnControllerProtocol) String() string {
	switch e {
		
	case 0:
			return "ssl"
		
	case 1:
			return "pssl"
		
		default:
			return ""
	}

}

func ToSdnControllerProtocol(strValue string) SdnControllerProtocol {
	switch strValue { 
		case "ssl":
			return 0
		case "pssl":
			return 1
		default:
			return -1
	}
}


// This is a generated file. DO NOT EDIT manually.

package go_xen_client

//SrHealth: sr_health enum type
type SrHealth int
const(
   SrHealthHealthy SrHealth = iota //Storage is fully available  
   SrHealthRecovering //Storage is busy recovering, e.g. rebuilding mirrors.  
)

func (e SrHealth) String() string {
	switch e {
		
	case 0:
			return "healthy"
		
	case 1:
			return "recovering"
		
		default:
			return ""
	}

}

func ToSrHealth(strValue string) SrHealth {
	switch strValue { 
		case "healthy":
			return 0
		case "recovering":
			return 1
		default:
			return -1
	}
}


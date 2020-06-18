// This is a generated file. DO NOT EDIT manually.

package go_xen_client

//AfterApplyGuidance: after_apply_guidance enum type
type AfterApplyGuidance int
const(
   AfterApplyGuidanceRestartHVM AfterApplyGuidance = iota //This patch requires HVM guests to be restarted once applied.  
   AfterApplyGuidanceRestartPV //This patch requires PV guests to be restarted once applied.  
   AfterApplyGuidanceRestartHost //This patch requires the host to be restarted once applied.  
   AfterApplyGuidanceRestartXAPI //This patch requires XAPI to be restarted once applied.  
)

func (e AfterApplyGuidance) String() string {
	switch e {
		
	case 0:
			return "restartHVM"
		
	case 1:
			return "restartPV"
		
	case 2:
			return "restartHost"
		
	case 3:
			return "restartXAPI"
		
		default:
			return ""
	}

}

func ToAfterApplyGuidance(strValue string) AfterApplyGuidance {
	switch strValue { 
		case "restartHVM":
			return 0
		case "restartPV":
			return 1
		case "restartHost":
			return 2
		case "restartXAPI":
			return 3
		default:
			return -1
	}
}


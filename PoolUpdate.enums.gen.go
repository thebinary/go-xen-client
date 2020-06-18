// This is a generated file. DO NOT EDIT manually.

package go_xen_client

//UpdateAfterApplyGuidance: update_after_apply_guidance enum type
type UpdateAfterApplyGuidance int
const(
   UpdateAfterApplyGuidanceRestartHVM UpdateAfterApplyGuidance = iota //This update requires HVM guests to be restarted once applied.  
   UpdateAfterApplyGuidanceRestartPV //This update requires PV guests to be restarted once applied.  
   UpdateAfterApplyGuidanceRestartHost //This update requires the host to be restarted once applied.  
   UpdateAfterApplyGuidanceRestartXAPI //This update requires XAPI to be restarted once applied.  
)

func (e UpdateAfterApplyGuidance) String() string {
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

func ToUpdateAfterApplyGuidance(strValue string) UpdateAfterApplyGuidance {
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

//LivepatchStatus: livepatch_status enum type
type LivepatchStatus int
const(
   LivepatchStatusOkLivepatchComplete LivepatchStatus = iota //An applicable live patch exists for every required component  
   LivepatchStatusOkLivepatchIncomplete //An applicable live patch exists but it is not sufficient  
   LivepatchStatusOk //There is no applicable live patch  
)

func (e LivepatchStatus) String() string {
	switch e {
		
	case 0:
			return "ok_livepatch_complete"
		
	case 1:
			return "ok_livepatch_incomplete"
		
	case 2:
			return "ok"
		
		default:
			return ""
	}

}

func ToLivepatchStatus(strValue string) LivepatchStatus {
	switch strValue { 
		case "ok_livepatch_complete":
			return 0
		case "ok_livepatch_incomplete":
			return 1
		case "ok":
			return 2
		default:
			return -1
	}
}


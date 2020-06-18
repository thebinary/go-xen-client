// This is a generated file. DO NOT EDIT manually.

package go_xen_client

//VmssFrequency: vmss_frequency enum type
type VmssFrequency int
const(
   VmssFrequencyHourly VmssFrequency = iota //Hourly snapshots  
   VmssFrequencyDaily //Daily snapshots  
   VmssFrequencyWeekly //Weekly snapshots  
)

func (e VmssFrequency) String() string {
	switch e {
		
	case 0:
			return "hourly"
		
	case 1:
			return "daily"
		
	case 2:
			return "weekly"
		
		default:
			return ""
	}

}

func ToVmssFrequency(strValue string) VmssFrequency {
	switch strValue { 
		case "hourly":
			return 0
		case "daily":
			return 1
		case "weekly":
			return 2
		default:
			return -1
	}
}

//VmssType: vmss_type enum type
type VmssType int
const(
   VmssTypeSnapshot VmssType = iota //The snapshot is a disk snapshot  
   VmssTypeCheckpoint //The snapshot is a checkpoint  
   VmssTypeSnapshotWithQuiesce //The snapshot is a VSS  
)

func (e VmssType) String() string {
	switch e {
		
	case 0:
			return "snapshot"
		
	case 1:
			return "checkpoint"
		
	case 2:
			return "snapshot_with_quiesce"
		
		default:
			return ""
	}

}

func ToVmssType(strValue string) VmssType {
	switch strValue { 
		case "snapshot":
			return 0
		case "checkpoint":
			return 1
		case "snapshot_with_quiesce":
			return 2
		default:
			return -1
	}
}


// This is a generated file. DO NOT EDIT manually.

package go_xen_client

//VbdOperations: vbd_operations enum type
type VbdOperations int
const(
   VbdOperationsAttach VbdOperations = iota //Attempting to attach this VBD to a VM  
   VbdOperationsEject //Attempting to eject the media from this VBD  
   VbdOperationsInsert //Attempting to insert new media into this VBD  
   VbdOperationsPlug //Attempting to hotplug this VBD  
   VbdOperationsUnplug //Attempting to hot unplug this VBD  
   VbdOperationsUnplugForce //Attempting to forcibly unplug this VBD  
   VbdOperationsPause //Attempting to pause a block device backend  
   VbdOperationsUnpause //Attempting to unpause a block device backend  
)

func (e VbdOperations) String() string {
	switch e {
		
	case 0:
			return "attach"
		
	case 1:
			return "eject"
		
	case 2:
			return "insert"
		
	case 3:
			return "plug"
		
	case 4:
			return "unplug"
		
	case 5:
			return "unplug_force"
		
	case 6:
			return "pause"
		
	case 7:
			return "unpause"
		
		default:
			return ""
	}

}

func ToVbdOperations(strValue string) VbdOperations {
	switch strValue { 
		case "attach":
			return 0
		case "eject":
			return 1
		case "insert":
			return 2
		case "plug":
			return 3
		case "unplug":
			return 4
		case "unplug_force":
			return 5
		case "pause":
			return 6
		case "unpause":
			return 7
		default:
			return -1
	}
}

//VbdType: vbd_type enum type
type VbdType int
const(
   VbdTypeCD VbdType = iota //VBD will appear to guest as CD  
   VbdTypeDisk //VBD will appear to guest as disk  
   VbdTypeFloppy //VBD will appear as a floppy  
)

func (e VbdType) String() string {
	switch e {
		
	case 0:
			return "CD"
		
	case 1:
			return "Disk"
		
	case 2:
			return "Floppy"
		
		default:
			return ""
	}

}

func ToVbdType(strValue string) VbdType {
	switch strValue { 
		case "CD":
			return 0
		case "Disk":
			return 1
		case "Floppy":
			return 2
		default:
			return -1
	}
}

//VbdMode: vbd_mode enum type
type VbdMode int
const(
   VbdModeRO VbdMode = iota //only read-only access will be allowed  
   VbdModeRW //read-write access will be allowed  
)

func (e VbdMode) String() string {
	switch e {
		
	case 0:
			return "RO"
		
	case 1:
			return "RW"
		
		default:
			return ""
	}

}

func ToVbdMode(strValue string) VbdMode {
	switch strValue { 
		case "RO":
			return 0
		case "RW":
			return 1
		default:
			return -1
	}
}


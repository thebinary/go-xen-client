// This is a generated file. DO NOT EDIT manually.

package go_xen_client

//VdiOperations: vdi_operations enum type
type VdiOperations int
const(
   VdiOperationsClone VdiOperations = iota //Cloning the VDI  
   VdiOperationsCopy //Copying the VDI  
   VdiOperationsResize //Resizing the VDI  
   VdiOperationsResizeOnline //Resizing the VDI which may or may not be online  
   VdiOperationsSnapshot //Snapshotting the VDI  
   VdiOperationsMirror //Mirroring the VDI  
   VdiOperationsDestroy //Destroying the VDI  
   VdiOperationsForget //Forget about the VDI  
   VdiOperationsUpdate //Refreshing the fields of the VDI  
   VdiOperationsForceUnlock //Forcibly unlocking the VDI  
   VdiOperationsGenerateConfig //Generating static configuration  
   VdiOperationsEnableCbt //Enabling changed block tracking for a VDI  
   VdiOperationsDisableCbt //Disabling changed block tracking for a VDI  
   VdiOperationsDataDestroy //Deleting the data of the VDI  
   VdiOperationsListChangedBlocks //Exporting a bitmap that shows the changed blocks between two VDIs  
   VdiOperationsSetOnBoot //Setting the on_boot field of the VDI  
   VdiOperationsBlocked //Operations on this VDI are temporarily blocked  
)

func (e VdiOperations) String() string {
	switch e {
		
	case 0:
			return "clone"
		
	case 1:
			return "copy"
		
	case 2:
			return "resize"
		
	case 3:
			return "resize_online"
		
	case 4:
			return "snapshot"
		
	case 5:
			return "mirror"
		
	case 6:
			return "destroy"
		
	case 7:
			return "forget"
		
	case 8:
			return "update"
		
	case 9:
			return "force_unlock"
		
	case 10:
			return "generate_config"
		
	case 11:
			return "enable_cbt"
		
	case 12:
			return "disable_cbt"
		
	case 13:
			return "data_destroy"
		
	case 14:
			return "list_changed_blocks"
		
	case 15:
			return "set_on_boot"
		
	case 16:
			return "blocked"
		
		default:
			return ""
	}

}

func ToVdiOperations(strValue string) VdiOperations {
	switch strValue { 
		case "clone":
			return 0
		case "copy":
			return 1
		case "resize":
			return 2
		case "resize_online":
			return 3
		case "snapshot":
			return 4
		case "mirror":
			return 5
		case "destroy":
			return 6
		case "forget":
			return 7
		case "update":
			return 8
		case "force_unlock":
			return 9
		case "generate_config":
			return 10
		case "enable_cbt":
			return 11
		case "disable_cbt":
			return 12
		case "data_destroy":
			return 13
		case "list_changed_blocks":
			return 14
		case "set_on_boot":
			return 15
		case "blocked":
			return 16
		default:
			return -1
	}
}

//VdiType: vdi_type enum type
type VdiType int
const(
   VdiTypeSystem VdiType = iota //a disk that may be replaced on upgrade  
   VdiTypeUser //a disk that is always preserved on upgrade  
   VdiTypeEphemeral //a disk that may be reformatted on upgrade  
   VdiTypeSuspend //a disk that stores a suspend image  
   VdiTypeCrashdump //a disk that stores VM crashdump information  
   VdiTypeHaStatefile //a disk used for HA storage heartbeating  
   VdiTypeMetadata //a disk used for HA Pool metadata  
   VdiTypeRedoLog //a disk used for a general metadata redo-log  
   VdiTypeRrd //a disk that stores SR-level RRDs  
   VdiTypePvsCache //a disk that stores PVS cache data  
   VdiTypeCbtMetadata //Metadata about a snapshot VDI that has been deleted: the set of blocks that changed between some previous version of the disk and the version tracked by the snapshot.  
)

func (e VdiType) String() string {
	switch e {
		
	case 0:
			return "system"
		
	case 1:
			return "user"
		
	case 2:
			return "ephemeral"
		
	case 3:
			return "suspend"
		
	case 4:
			return "crashdump"
		
	case 5:
			return "ha_statefile"
		
	case 6:
			return "metadata"
		
	case 7:
			return "redo_log"
		
	case 8:
			return "rrd"
		
	case 9:
			return "pvs_cache"
		
	case 10:
			return "cbt_metadata"
		
		default:
			return ""
	}

}

func ToVdiType(strValue string) VdiType {
	switch strValue { 
		case "system":
			return 0
		case "user":
			return 1
		case "ephemeral":
			return 2
		case "suspend":
			return 3
		case "crashdump":
			return 4
		case "ha_statefile":
			return 5
		case "metadata":
			return 6
		case "redo_log":
			return 7
		case "rrd":
			return 8
		case "pvs_cache":
			return 9
		case "cbt_metadata":
			return 10
		default:
			return -1
	}
}

//OnBoot: on_boot enum type
type OnBoot int
const(
   OnBootReset OnBoot = iota //When a VM containing this VDI is started, the contents of the VDI are reset to the state they were in when this flag was last set.  
   OnBootPersist //Standard behaviour.  
)

func (e OnBoot) String() string {
	switch e {
		
	case 0:
			return "reset"
		
	case 1:
			return "persist"
		
		default:
			return ""
	}

}

func ToOnBoot(strValue string) OnBoot {
	switch strValue { 
		case "reset":
			return 0
		case "persist":
			return 1
		default:
			return -1
	}
}


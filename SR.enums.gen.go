// This is a generated file. DO NOT EDIT manually.

package go_xen_client

//StorageOperations: storage_operations enum type
type StorageOperations int
const(
   StorageOperationsScan StorageOperations = iota //Scanning backends for new or deleted VDIs  
   StorageOperationsDestroy //Destroying the SR  
   StorageOperationsForget //Forgetting about SR  
   StorageOperationsPlug //Plugging a PBD into this SR  
   StorageOperationsUnplug //Unplugging a PBD from this SR  
   StorageOperationsUpdate //Refresh the fields on the SR  
   StorageOperationsVdiCreate //Creating a new VDI  
   StorageOperationsVdiIntroduce //Introducing a new VDI  
   StorageOperationsVdiDestroy //Destroying a VDI  
   StorageOperationsVdiResize //Resizing a VDI  
   StorageOperationsVdiClone //Cloneing a VDI  
   StorageOperationsVdiSnapshot //Snapshotting a VDI  
   StorageOperationsVdiMirror //Mirroring a VDI  
   StorageOperationsVdiEnableCbt //Enabling changed block tracking for a VDI  
   StorageOperationsVdiDisableCbt //Disabling changed block tracking for a VDI  
   StorageOperationsVdiDataDestroy //Deleting the data of the VDI  
   StorageOperationsVdiListChangedBlocks //Exporting a bitmap that shows the changed blocks between two VDIs  
   StorageOperationsVdiSetOnBoot //Setting the on_boot field of the VDI  
   StorageOperationsPbdCreate //Creating a PBD for this SR  
   StorageOperationsPbdDestroy //Destroying one of this SR's PBDs  
)

func (e StorageOperations) String() string {
	switch e {
		
	case 0:
			return "scan"
		
	case 1:
			return "destroy"
		
	case 2:
			return "forget"
		
	case 3:
			return "plug"
		
	case 4:
			return "unplug"
		
	case 5:
			return "update"
		
	case 6:
			return "vdi_create"
		
	case 7:
			return "vdi_introduce"
		
	case 8:
			return "vdi_destroy"
		
	case 9:
			return "vdi_resize"
		
	case 10:
			return "vdi_clone"
		
	case 11:
			return "vdi_snapshot"
		
	case 12:
			return "vdi_mirror"
		
	case 13:
			return "vdi_enable_cbt"
		
	case 14:
			return "vdi_disable_cbt"
		
	case 15:
			return "vdi_data_destroy"
		
	case 16:
			return "vdi_list_changed_blocks"
		
	case 17:
			return "vdi_set_on_boot"
		
	case 18:
			return "pbd_create"
		
	case 19:
			return "pbd_destroy"
		
		default:
			return ""
	}

}

func ToStorageOperations(strValue string) StorageOperations {
	switch strValue { 
		case "scan":
			return 0
		case "destroy":
			return 1
		case "forget":
			return 2
		case "plug":
			return 3
		case "unplug":
			return 4
		case "update":
			return 5
		case "vdi_create":
			return 6
		case "vdi_introduce":
			return 7
		case "vdi_destroy":
			return 8
		case "vdi_resize":
			return 9
		case "vdi_clone":
			return 10
		case "vdi_snapshot":
			return 11
		case "vdi_mirror":
			return 12
		case "vdi_enable_cbt":
			return 13
		case "vdi_disable_cbt":
			return 14
		case "vdi_data_destroy":
			return 15
		case "vdi_list_changed_blocks":
			return 16
		case "vdi_set_on_boot":
			return 17
		case "pbd_create":
			return 18
		case "pbd_destroy":
			return 19
		default:
			return -1
	}
}


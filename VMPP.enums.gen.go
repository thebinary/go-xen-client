// This is a generated file. DO NOT EDIT manually.

package go_xen_client

//VmppBackupType: vmpp_backup_type enum type
type VmppBackupType int
const(
   VmppBackupTypeSnapshot VmppBackupType = iota //The backup is a snapshot  
   VmppBackupTypeCheckpoint //The backup is a checkpoint  
)

func (e VmppBackupType) String() string {
	switch e {
		
	case 0:
			return "snapshot"
		
	case 1:
			return "checkpoint"
		
		default:
			return ""
	}

}

func ToVmppBackupType(strValue string) VmppBackupType {
	switch strValue { 
		case "snapshot":
			return 0
		case "checkpoint":
			return 1
		default:
			return -1
	}
}

//VmppBackupFrequency: vmpp_backup_frequency enum type
type VmppBackupFrequency int
const(
   VmppBackupFrequencyHourly VmppBackupFrequency = iota //Hourly backups  
   VmppBackupFrequencyDaily //Daily backups  
   VmppBackupFrequencyWeekly //Weekly backups  
)

func (e VmppBackupFrequency) String() string {
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

func ToVmppBackupFrequency(strValue string) VmppBackupFrequency {
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

//VmppArchiveFrequency: vmpp_archive_frequency enum type
type VmppArchiveFrequency int
const(
   VmppArchiveFrequencyNever VmppArchiveFrequency = iota //Never archive  
   VmppArchiveFrequencyAlwaysAfterBackup //Archive after backup  
   VmppArchiveFrequencyDaily //Daily archives  
   VmppArchiveFrequencyWeekly //Weekly backups  
)

func (e VmppArchiveFrequency) String() string {
	switch e {
		
	case 0:
			return "never"
		
	case 1:
			return "always_after_backup"
		
	case 2:
			return "daily"
		
	case 3:
			return "weekly"
		
		default:
			return ""
	}

}

func ToVmppArchiveFrequency(strValue string) VmppArchiveFrequency {
	switch strValue { 
		case "never":
			return 0
		case "always_after_backup":
			return 1
		case "daily":
			return 2
		case "weekly":
			return 3
		default:
			return -1
	}
}

//VmppArchiveTargetType: vmpp_archive_target_type enum type
type VmppArchiveTargetType int
const(
   VmppArchiveTargetTypeNone VmppArchiveTargetType = iota //No target config  
   VmppArchiveTargetTypeCifs //CIFS target config  
   VmppArchiveTargetTypeNfs //NFS target config  
)

func (e VmppArchiveTargetType) String() string {
	switch e {
		
	case 0:
			return "none"
		
	case 1:
			return "cifs"
		
	case 2:
			return "nfs"
		
		default:
			return ""
	}

}

func ToVmppArchiveTargetType(strValue string) VmppArchiveTargetType {
	switch strValue { 
		case "none":
			return 0
		case "cifs":
			return 1
		case "nfs":
			return 2
		default:
			return -1
	}
}


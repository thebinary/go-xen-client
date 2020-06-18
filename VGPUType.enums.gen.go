// This is a generated file. DO NOT EDIT manually.

package go_xen_client

//VgpuTypeImplementation: vgpu_type_implementation enum type
type VgpuTypeImplementation int
const(
   VgpuTypeImplementationPassthrough VgpuTypeImplementation = iota //Pass through an entire physical GPU to a guest  
   VgpuTypeImplementationNvidia //vGPU using NVIDIA hardware  
   VgpuTypeImplementationGvtG //vGPU using Intel GVT-g  
   VgpuTypeImplementationMxgpu //vGPU using AMD MxGPU  
)

func (e VgpuTypeImplementation) String() string {
	switch e {
		
	case 0:
			return "passthrough"
		
	case 1:
			return "nvidia"
		
	case 2:
			return "gvt_g"
		
	case 3:
			return "mxgpu"
		
		default:
			return ""
	}

}

func ToVgpuTypeImplementation(strValue string) VgpuTypeImplementation {
	switch strValue { 
		case "passthrough":
			return 0
		case "nvidia":
			return 1
		case "gvt_g":
			return 2
		case "mxgpu":
			return 3
		default:
			return -1
	}
}


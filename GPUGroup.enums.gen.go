// This is a generated file. DO NOT EDIT manually.

package go_xen_client

//AllocationAlgorithm: allocation_algorithm enum type
type AllocationAlgorithm int
const(
   AllocationAlgorithmBreadthFirst AllocationAlgorithm = iota //vGPUs of a given type are allocated evenly across supporting pGPUs.  
   AllocationAlgorithmDepthFirst //vGPUs of a given type are allocated on supporting pGPUs until they are full.  
)

func (e AllocationAlgorithm) String() string {
	switch e {
		
	case 0:
			return "breadth_first"
		
	case 1:
			return "depth_first"
		
		default:
			return ""
	}

}

func ToAllocationAlgorithm(strValue string) AllocationAlgorithm {
	switch strValue { 
		case "breadth_first":
			return 0
		case "depth_first":
			return 1
		default:
			return -1
	}
}


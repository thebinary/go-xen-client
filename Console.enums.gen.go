// This is a generated file. DO NOT EDIT manually.

package go_xen_client

//ConsoleProtocol: console_protocol enum type
type ConsoleProtocol int
const(
   ConsoleProtocolVt100 ConsoleProtocol = iota //VT100 terminal  
   ConsoleProtocolRfb //Remote FrameBuffer protocol (as used in VNC)  
   ConsoleProtocolRdp //Remote Desktop Protocol  
)

func (e ConsoleProtocol) String() string {
	switch e {
		
	case 0:
			return "vt100"
		
	case 1:
			return "rfb"
		
	case 2:
			return "rdp"
		
		default:
			return ""
	}

}

func ToConsoleProtocol(strValue string) ConsoleProtocol {
	switch strValue { 
		case "vt100":
			return 0
		case "rfb":
			return 1
		case "rdp":
			return 2
		default:
			return -1
	}
}


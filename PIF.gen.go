// This is a generated file. DO NOT EDIT manually.
//go:generate goimports -w PIF.gen.go

package go_xen_client

import (
	"reflect"
	"strconv"

	"github.com/nilshell/xmlrpc"
)

//PIF: A physical network interface (note separate VLANs are represented as several PIFs)
type PIF struct {
	Uuid                  string                // Unique identifier/object reference
	Device                string                // machine-readable name of the interface (e.g. eth0)
	Network               string                // virtual network to which this pif is connected
	Host                  string                // physical machine to which this pif is connected
	MAC                   string                // ethernet MAC address of physical interface
	MTU                   int                   // MTU in octets
	VLAN                  int                   // VLAN tag for all traffic passing through this interface
	Metrics               string                // metrics associated with this PIF
	Physical              bool                  // true if this represents a physical network interface
	CurrentlyAttached     bool                  // true if this interface is online
	IpConfigurationMode   IpConfigurationMode   // Sets if and how this interface gets an IP address
	IP                    string                // IP address
	Netmask               string                // IP netmask
	Gateway               string                // IP gateway
	DNS                   string                // Comma separated list of the IP addresses of the DNS servers to use
	BondSlaveOf           string                // Indicates which bond this interface is part of
	BondMasterOf          []string              // Indicates this PIF represents the results of a bond
	VLANMasterOf          string                // Indicates wich VLAN this interface receives untagged traffic from
	VLANSlaveOf           []string              // Indicates which VLANs this interface transmits tagged traffic to
	Management            bool                  // Indicates whether the control software is listening for connections on this interface
	OtherConfig           map[string]string     // Additional configuration
	DisallowUnplug        bool                  // Prevent this PIF from being unplugged; set this to notify the management tool-stack that the PIF has a special use and should not be unplugged under any circumstances (e.g. because you're running storage traffic over it)
	TunnelAccessPIFOf     []string              // Indicates to which tunnel this PIF gives access
	TunnelTransportPIFOf  []string              // Indicates to which tunnel this PIF provides transport
	Ipv6ConfigurationMode Ipv6ConfigurationMode // Sets if and how this interface gets an IPv6 address
	IPv6                  []string              // IPv6 address
	Ipv6Gateway           string                // IPv6 gateway
	PrimaryAddressType    PrimaryAddressType    // Which protocol should define the primary address of this interface
	Managed               bool                  // Indicates whether the interface is managed by xapi. If it is not, then xapi will not configure the interface, the commands PIF.plug/unplug/reconfigure_ip(v6) cannot be used, nor can the interface be bonded or have VLANs based on top through xapi.
	Properties            map[string]string     // Additional configuration properties for the interface.
	Capabilities          []string              // Additional capabilities on the interface.
	IgmpSnoopingStatus    PifIgmpStatus         // The IGMP snooping status of the corresponding network bridge
	SriovPhysicalPIFOf    []string              // Indicates which network_sriov this interface is physical of
	SriovLogicalPIFOf     []string              // Indicates which network_sriov this interface is logical of
	PCI                   string                // Link to underlying PCI device

}

func FromPIFToXml(PIF *PIF) (result xmlrpc.Struct) {
	result = make(xmlrpc.Struct)

	result["uuid"] = PIF.Uuid

	result["device"] = PIF.Device

	result["network"] = PIF.Network

	result["host"] = PIF.Host

	result["MAC"] = PIF.MAC

	result["MTU"] = strconv.Itoa(PIF.MTU)

	result["VLAN"] = strconv.Itoa(PIF.VLAN)

	result["metrics"] = PIF.Metrics

	result["physical"] = PIF.Physical

	result["currently_attached"] = PIF.CurrentlyAttached

	result["ip_configuration_mode"] = PIF.IpConfigurationMode.String()

	result["IP"] = PIF.IP

	result["netmask"] = PIF.Netmask

	result["gateway"] = PIF.Gateway

	result["DNS"] = PIF.DNS

	result["bond_slave_of"] = PIF.BondSlaveOf

	result["bond_master_of"] = PIF.BondMasterOf

	result["VLAN_master_of"] = PIF.VLANMasterOf

	result["VLAN_slave_of"] = PIF.VLANSlaveOf

	result["management"] = PIF.Management

	other_config := make(xmlrpc.Struct)
	for key, value := range PIF.OtherConfig {
		other_config[key] = value
	}
	result["other_config"] = other_config

	result["disallow_unplug"] = PIF.DisallowUnplug

	result["tunnel_access_PIF_of"] = PIF.TunnelAccessPIFOf

	result["tunnel_transport_PIF_of"] = PIF.TunnelTransportPIFOf

	result["ipv6_configuration_mode"] = PIF.Ipv6ConfigurationMode.String()

	result["IPv6"] = PIF.IPv6

	result["ipv6_gateway"] = PIF.Ipv6Gateway

	result["primary_address_type"] = PIF.PrimaryAddressType.String()

	result["managed"] = PIF.Managed

	properties := make(xmlrpc.Struct)
	for key, value := range PIF.Properties {
		properties[key] = value
	}
	result["properties"] = properties

	result["capabilities"] = PIF.Capabilities

	result["igmp_snooping_status"] = PIF.IgmpSnoopingStatus.String()

	result["sriov_physical_PIF_of"] = PIF.SriovPhysicalPIFOf

	result["sriov_logical_PIF_of"] = PIF.SriovLogicalPIFOf

	result["PCI"] = PIF.PCI

	return result
}

func ToPIF(obj interface{}) (resultObj *PIF) {

	objValue := reflect.ValueOf(obj)
	resultObj = &PIF{}

	for _, oKey := range objValue.MapKeys() {
		keyName := oKey.String()
		keyValue := objValue.MapIndex(oKey).Interface()
		switch keyName {
		case "uuid":
			if v, ok := keyValue.(string); ok {
				resultObj.Uuid = v
			}
		case "device":
			if v, ok := keyValue.(string); ok {
				resultObj.Device = v
			}
		case "network":
			if v, ok := keyValue.(string); ok {
				resultObj.Network = v
			}
		case "host":
			if v, ok := keyValue.(string); ok {
				resultObj.Host = v
			}
		case "MAC":
			if v, ok := keyValue.(string); ok {
				resultObj.MAC = v
			}
		case "MTU":
			if v, ok := keyValue.(int); ok {
				resultObj.MTU = v
			}
		case "VLAN":
			if v, ok := keyValue.(int); ok {
				resultObj.VLAN = v
			}
		case "metrics":
			if v, ok := keyValue.(string); ok {
				resultObj.Metrics = v
			}
		case "physical":
			if v, ok := keyValue.(bool); ok {
				resultObj.Physical = v
			}
		case "currently_attached":
			if v, ok := keyValue.(bool); ok {
				resultObj.CurrentlyAttached = v
			}
		case "ip_configuration_mode":
			if v, ok := keyValue.(IpConfigurationMode); ok {
				resultObj.IpConfigurationMode = v
			}
		case "IP":
			if v, ok := keyValue.(string); ok {
				resultObj.IP = v
			}
		case "netmask":
			if v, ok := keyValue.(string); ok {
				resultObj.Netmask = v
			}
		case "gateway":
			if v, ok := keyValue.(string); ok {
				resultObj.Gateway = v
			}
		case "DNS":
			if v, ok := keyValue.(string); ok {
				resultObj.DNS = v
			}
		case "bond_slave_of":
			if v, ok := keyValue.(string); ok {
				resultObj.BondSlaveOf = v
			}
		case "bond_master_of":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.BondMasterOf = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.BondMasterOf[i] = v
					}
				}
			}
		case "VLAN_master_of":
			if v, ok := keyValue.(string); ok {
				resultObj.VLANMasterOf = v
			}
		case "VLAN_slave_of":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.VLANSlaveOf = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.VLANSlaveOf[i] = v
					}
				}
			}
		case "management":
			if v, ok := keyValue.(bool); ok {
				resultObj.Management = v
			}
		case "other_config":

			resultObj.OtherConfig = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.OtherConfig[mapKeyName] = v
				} else {
					resultObj.OtherConfig[mapKeyName] = ""
				}
			}
		case "disallow_unplug":
			if v, ok := keyValue.(bool); ok {
				resultObj.DisallowUnplug = v
			}
		case "tunnel_access_PIF_of":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.TunnelAccessPIFOf = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.TunnelAccessPIFOf[i] = v
					}
				}
			}
		case "tunnel_transport_PIF_of":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.TunnelTransportPIFOf = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.TunnelTransportPIFOf[i] = v
					}
				}
			}
		case "ipv6_configuration_mode":
			if v, ok := keyValue.(Ipv6ConfigurationMode); ok {
				resultObj.Ipv6ConfigurationMode = v
			}
		case "IPv6":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.IPv6 = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.IPv6[i] = v
					}
				}
			}
		case "ipv6_gateway":
			if v, ok := keyValue.(string); ok {
				resultObj.Ipv6Gateway = v
			}
		case "primary_address_type":
			if v, ok := keyValue.(PrimaryAddressType); ok {
				resultObj.PrimaryAddressType = v
			}
		case "managed":
			if v, ok := keyValue.(bool); ok {
				resultObj.Managed = v
			}
		case "properties":

			resultObj.Properties = map[string]string{}
			interimMap := reflect.ValueOf(keyValue).MapKeys()
			for _, mapKey := range interimMap {
				mapKeyName := mapKey.String()
				mapKeyValue := reflect.ValueOf(keyValue).MapIndex(mapKey).Interface()
				if v, ok := mapKeyValue.(string); ok {
					resultObj.Properties[mapKeyName] = v
				} else {
					resultObj.Properties[mapKeyName] = ""
				}
			}
		case "capabilities":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.Capabilities = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.Capabilities[i] = v
					}
				}
			}
		case "igmp_snooping_status":
			if v, ok := keyValue.(PifIgmpStatus); ok {
				resultObj.IgmpSnoopingStatus = v
			}
		case "sriov_physical_PIF_of":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.SriovPhysicalPIFOf = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.SriovPhysicalPIFOf[i] = v
					}
				}
			}
		case "sriov_logical_PIF_of":
			if interim, ok := keyValue.([]interface{}); ok {
				resultObj.SriovLogicalPIFOf = make([]string, len(interim))
				for i, interimValue := range interim {
					if v, ok := interimValue.(string); ok {
						resultObj.SriovLogicalPIFOf[i] = v
					}
				}
			}
		case "PCI":
			if v, ok := keyValue.(string); ok {
				resultObj.PCI = v
			}
		}
	}

	return resultObj
}

/* GetAllRecords: Return a map of PIF references to PIF records for all PIFs known to the system. */
func (client *XenClient) PIFGetAllRecords() (result map[string]PIF, err error) {
	obj, err := client.APICall("PIF.get_all_records")
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]PIF{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		mapObj := ToPIF(obj.Interface())
		result[key.String()] = *mapObj
	}

	return
}

/* GetAll: Return a list of all the PIFs known to the system. */
func (client *XenClient) PIFGetAll() (result []string, err error) {
	obj, err := client.APICall("PIF.get_all")
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* SetProperty: Set the value of a property of the PIF */
func (client *XenClient) PIFSetProperty(self string, name string, value string) (err error) {
	_, err = client.APICall("PIF.set_property", self, name, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* DbForget: Destroy a PIF database record. */
func (client *XenClient) PIFDbForget(self string) (err error) {
	_, err = client.APICall("PIF.db_forget", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* DbIntroduce: Create a new PIF record in the database only */
func (client *XenClient) PIFDbIntroduce(device string, network string, host string, MAC string, MTU int, VLAN int, physical bool, ip_configuration_mode IpConfigurationMode, IP string, netmask string, gateway string, DNS string, bond_slave_of string, VLAN_master_of string, management bool, other_config map[string]string, disallow_unplug bool, ipv6_configuration_mode Ipv6ConfigurationMode, IPv6 []string, ipv6_gateway string, primary_address_type PrimaryAddressType, managed bool, properties map[string]string) (result string, err error) {
	obj, err := client.APICall("PIF.db_introduce", device, network, host, MAC, MTU, VLAN, physical, ip_configuration_mode.String(), IP, netmask, gateway, DNS, bond_slave_of, VLAN_master_of, management, other_config, disallow_unplug, ipv6_configuration_mode.String(), IPv6, ipv6_gateway, primary_address_type.String(), managed, properties)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* Plug: Attempt to bring up a physical interface */
func (client *XenClient) PIFPlug(self string) (err error) {
	_, err = client.APICall("PIF.plug", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetDisallowUnplug: Set whether unplugging the PIF is allowed */
func (client *XenClient) PIFSetDisallowUnplug(self string, value bool) (err error) {
	_, err = client.APICall("PIF.set_disallow_unplug", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Unplug: Attempt to bring down a physical interface */
func (client *XenClient) PIFUnplug(self string) (err error) {
	_, err = client.APICall("PIF.unplug", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Forget: Destroy the PIF object matching a particular network interface */
func (client *XenClient) PIFForget(self string) (err error) {
	_, err = client.APICall("PIF.forget", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Introduce: Create a PIF object matching a particular network interface */
func (client *XenClient) PIFIntroduce(host string, MAC string, device string, managed bool) (result string, err error) {
	obj, err := client.APICall("PIF.introduce", host, MAC, device, managed)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* Scan: Scan for physical interfaces on a host and create PIF objects to represent them */
func (client *XenClient) PIFScan(host string) (err error) {
	_, err = client.APICall("PIF.scan", host)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetPrimaryAddressType: Change the primary address type used by this PIF */
func (client *XenClient) PIFSetPrimaryAddressType(self string, primary_address_type PrimaryAddressType) (err error) {
	_, err = client.APICall("PIF.set_primary_address_type", self, primary_address_type.String())
	if err != nil {
		return
	}
	// no return result
	return
}

/* ReconfigureIpv6: Reconfigure the IPv6 address settings for this interface */
func (client *XenClient) PIFReconfigureIpv6(self string, mode Ipv6ConfigurationMode, IPv6 string, gateway string, DNS string) (err error) {
	_, err = client.APICall("PIF.reconfigure_ipv6", self, mode.String(), IPv6, gateway, DNS)
	if err != nil {
		return
	}
	// no return result
	return
}

/* ReconfigureIp: Reconfigure the IP address settings for this interface */
func (client *XenClient) PIFReconfigureIp(self string, mode IpConfigurationMode, IP string, netmask string, gateway string, DNS string) (err error) {
	_, err = client.APICall("PIF.reconfigure_ip", self, mode.String(), IP, netmask, gateway, DNS)
	if err != nil {
		return
	}
	// no return result
	return
}

/* Destroy: Destroy the PIF object (provided it is a VLAN interface). This call is deprecated: use VLAN.destroy or Bond.destroy instead */
func (client *XenClient) PIFDestroy(self string) (err error) {
	_, err = client.APICall("PIF.destroy", self)
	if err != nil {
		return
	}
	// no return result
	return
}

/* CreateVLAN: Create a VLAN interface from an existing physical interface. This call is deprecated: use VLAN.create instead */
func (client *XenClient) PIFCreateVLAN(device string, network string, host string, VLAN int) (result string, err error) {
	obj, err := client.APICall("PIF.create_VLAN", device, network, host, VLAN)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* RemoveFromOtherConfig: Remove the given key and its corresponding value from the other_config field of the given PIF.  If the key is not in that Map, then do nothing. */
func (client *XenClient) PIFRemoveFromOtherConfig(self string, key string) (err error) {
	_, err = client.APICall("PIF.remove_from_other_config", self, key)
	if err != nil {
		return
	}
	// no return result
	return
}

/* AddToOtherConfig: Add the given key-value pair to the other_config field of the given PIF. */
func (client *XenClient) PIFAddToOtherConfig(self string, key string, value string) (err error) {
	_, err = client.APICall("PIF.add_to_other_config", self, key, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* SetOtherConfig: Set the other_config field of the given PIF. */
func (client *XenClient) PIFSetOtherConfig(self string, value map[string]string) (err error) {
	_, err = client.APICall("PIF.set_other_config", self, value)
	if err != nil {
		return
	}
	// no return result
	return
}

/* GetPCI: Get the PCI field of the given PIF. */
func (client *XenClient) PIFGetPCI(self string) (result string, err error) {
	obj, err := client.APICall("PIF.get_PCI", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetSriovLogicalPIFOf: Get the sriov_logical_PIF_of field of the given PIF. */
func (client *XenClient) PIFGetSriovLogicalPIFOf(self string) (result []string, err error) {
	obj, err := client.APICall("PIF.get_sriov_logical_PIF_of", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetSriovPhysicalPIFOf: Get the sriov_physical_PIF_of field of the given PIF. */
func (client *XenClient) PIFGetSriovPhysicalPIFOf(self string) (result []string, err error) {
	obj, err := client.APICall("PIF.get_sriov_physical_PIF_of", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetIgmpSnoopingStatus: Get the igmp_snooping_status field of the given PIF. */
func (client *XenClient) PIFGetIgmpSnoopingStatus(self string) (result PifIgmpStatus, err error) {
	obj, err := client.APICall("PIF.get_igmp_snooping_status", self)
	if err != nil {
		return
	}

	result = ToPifIgmpStatus(obj.(string))

	return
}

/* GetCapabilities: Get the capabilities field of the given PIF. */
func (client *XenClient) PIFGetCapabilities(self string) (result []string, err error) {
	obj, err := client.APICall("PIF.get_capabilities", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetProperties: Get the properties field of the given PIF. */
func (client *XenClient) PIFGetProperties(self string) (result map[string]string, err error) {
	obj, err := client.APICall("PIF.get_properties", self)
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]string{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		result[key.String()] = obj.String()
	}

	return
}

/* GetManaged: Get the managed field of the given PIF. */
func (client *XenClient) PIFGetManaged(self string) (result bool, err error) {
	obj, err := client.APICall("PIF.get_managed", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetPrimaryAddressType: Get the primary_address_type field of the given PIF. */
func (client *XenClient) PIFGetPrimaryAddressType(self string) (result PrimaryAddressType, err error) {
	obj, err := client.APICall("PIF.get_primary_address_type", self)
	if err != nil {
		return
	}

	result = ToPrimaryAddressType(obj.(string))

	return
}

/* GetIpv6Gateway: Get the ipv6_gateway field of the given PIF. */
func (client *XenClient) PIFGetIpv6Gateway(self string) (result string, err error) {
	obj, err := client.APICall("PIF.get_ipv6_gateway", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetIPv6: Get the IPv6 field of the given PIF. */
func (client *XenClient) PIFGetIPv6(self string) (result []string, err error) {
	obj, err := client.APICall("PIF.get_IPv6", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetIpv6ConfigurationMode: Get the ipv6_configuration_mode field of the given PIF. */
func (client *XenClient) PIFGetIpv6ConfigurationMode(self string) (result Ipv6ConfigurationMode, err error) {
	obj, err := client.APICall("PIF.get_ipv6_configuration_mode", self)
	if err != nil {
		return
	}

	result = ToIpv6ConfigurationMode(obj.(string))

	return
}

/* GetTunnelTransportPIFOf: Get the tunnel_transport_PIF_of field of the given PIF. */
func (client *XenClient) PIFGetTunnelTransportPIFOf(self string) (result []string, err error) {
	obj, err := client.APICall("PIF.get_tunnel_transport_PIF_of", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetTunnelAccessPIFOf: Get the tunnel_access_PIF_of field of the given PIF. */
func (client *XenClient) PIFGetTunnelAccessPIFOf(self string) (result []string, err error) {
	obj, err := client.APICall("PIF.get_tunnel_access_PIF_of", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetDisallowUnplug: Get the disallow_unplug field of the given PIF. */
func (client *XenClient) PIFGetDisallowUnplug(self string) (result bool, err error) {
	obj, err := client.APICall("PIF.get_disallow_unplug", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetOtherConfig: Get the other_config field of the given PIF. */
func (client *XenClient) PIFGetOtherConfig(self string) (result map[string]string, err error) {
	obj, err := client.APICall("PIF.get_other_config", self)
	if err != nil {
		return
	}

	interim := reflect.ValueOf(obj)
	result = map[string]string{}
	for _, key := range interim.MapKeys() {
		obj := interim.MapIndex(key)
		result[key.String()] = obj.String()
	}

	return
}

/* GetManagement: Get the management field of the given PIF. */
func (client *XenClient) PIFGetManagement(self string) (result bool, err error) {
	obj, err := client.APICall("PIF.get_management", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetVLANSlaveOf: Get the VLAN_slave_of field of the given PIF. */
func (client *XenClient) PIFGetVLANSlaveOf(self string) (result []string, err error) {
	obj, err := client.APICall("PIF.get_VLAN_slave_of", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetVLANMasterOf: Get the VLAN_master_of field of the given PIF. */
func (client *XenClient) PIFGetVLANMasterOf(self string) (result string, err error) {
	obj, err := client.APICall("PIF.get_VLAN_master_of", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetBondMasterOf: Get the bond_master_of field of the given PIF. */
func (client *XenClient) PIFGetBondMasterOf(self string) (result []string, err error) {
	obj, err := client.APICall("PIF.get_bond_master_of", self)
	if err != nil {
		return
	}

	result = make([]string, len(obj.([]interface{})))
	for i, value := range obj.([]interface{}) {
		result[i] = value.(string)
	}

	return
}

/* GetBondSlaveOf: Get the bond_slave_of field of the given PIF. */
func (client *XenClient) PIFGetBondSlaveOf(self string) (result string, err error) {
	obj, err := client.APICall("PIF.get_bond_slave_of", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetDNS: Get the DNS field of the given PIF. */
func (client *XenClient) PIFGetDNS(self string) (result string, err error) {
	obj, err := client.APICall("PIF.get_DNS", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetGateway: Get the gateway field of the given PIF. */
func (client *XenClient) PIFGetGateway(self string) (result string, err error) {
	obj, err := client.APICall("PIF.get_gateway", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetNetmask: Get the netmask field of the given PIF. */
func (client *XenClient) PIFGetNetmask(self string) (result string, err error) {
	obj, err := client.APICall("PIF.get_netmask", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetIP: Get the IP field of the given PIF. */
func (client *XenClient) PIFGetIP(self string) (result string, err error) {
	obj, err := client.APICall("PIF.get_IP", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetIpConfigurationMode: Get the ip_configuration_mode field of the given PIF. */
func (client *XenClient) PIFGetIpConfigurationMode(self string) (result IpConfigurationMode, err error) {
	obj, err := client.APICall("PIF.get_ip_configuration_mode", self)
	if err != nil {
		return
	}

	result = ToIpConfigurationMode(obj.(string))

	return
}

/* GetCurrentlyAttached: Get the currently_attached field of the given PIF. */
func (client *XenClient) PIFGetCurrentlyAttached(self string) (result bool, err error) {
	obj, err := client.APICall("PIF.get_currently_attached", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetPhysical: Get the physical field of the given PIF. */
func (client *XenClient) PIFGetPhysical(self string) (result bool, err error) {
	obj, err := client.APICall("PIF.get_physical", self)
	if err != nil {
		return
	}
	result = obj.(bool)
	return
}

/* GetMetrics: Get the metrics field of the given PIF. */
func (client *XenClient) PIFGetMetrics(self string) (result string, err error) {
	obj, err := client.APICall("PIF.get_metrics", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetVLAN: Get the VLAN field of the given PIF. */
func (client *XenClient) PIFGetVLAN(self string) (result int, err error) {
	obj, err := client.APICall("PIF.get_VLAN", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetMTU: Get the MTU field of the given PIF. */
func (client *XenClient) PIFGetMTU(self string) (result int, err error) {
	obj, err := client.APICall("PIF.get_MTU", self)
	if err != nil {
		return
	}
	result = obj.(int)
	return
}

/* GetMAC: Get the MAC field of the given PIF. */
func (client *XenClient) PIFGetMAC(self string) (result string, err error) {
	obj, err := client.APICall("PIF.get_MAC", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetHost: Get the host field of the given PIF. */
func (client *XenClient) PIFGetHost(self string) (result string, err error) {
	obj, err := client.APICall("PIF.get_host", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetNetwork: Get the network field of the given PIF. */
func (client *XenClient) PIFGetNetwork(self string) (result string, err error) {
	obj, err := client.APICall("PIF.get_network", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetDevice: Get the device field of the given PIF. */
func (client *XenClient) PIFGetDevice(self string) (result string, err error) {
	obj, err := client.APICall("PIF.get_device", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetUuid: Get the uuid field of the given PIF. */
func (client *XenClient) PIFGetUuid(self string) (result string, err error) {
	obj, err := client.APICall("PIF.get_uuid", self)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetByUuid: Get a reference to the PIF instance with the specified UUID. */
func (client *XenClient) PIFGetByUuid(uuid string) (result string, err error) {
	obj, err := client.APICall("PIF.get_by_uuid", uuid)
	if err != nil {
		return
	}
	result = obj.(string)
	return
}

/* GetRecord: Get a record containing the current state of the given PIF. */
func (client *XenClient) PIFGetRecord(self string) (result PIF, err error) {
	obj, err := client.APICall("PIF.get_record", self)
	if err != nil {
		return
	}

	result = *ToPIF(obj.(interface{}))

	return
}

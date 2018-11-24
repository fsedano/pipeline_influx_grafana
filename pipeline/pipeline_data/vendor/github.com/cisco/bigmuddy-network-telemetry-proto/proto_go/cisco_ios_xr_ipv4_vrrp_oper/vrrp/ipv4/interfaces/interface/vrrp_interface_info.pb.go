// Code generated by protoc-gen-go.
// source: vrrp_interface_info.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv4_vrrp_oper_vrrp_ipv4_interfaces_interface is a generated protocol buffer package.

It is generated from these files:
	vrrp_interface_info.proto

It has these top-level messages:
	VrrpInterfaceInfo_KEYS
	VrrpInterfaceInfo
*/
package cisco_ios_xr_ipv4_vrrp_oper_vrrp_ipv4_interfaces_interface

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// VRRP Interface statistics
type VrrpInterfaceInfo_KEYS struct {
	InterfaceName string `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
}

func (m *VrrpInterfaceInfo_KEYS) Reset()                    { *m = VrrpInterfaceInfo_KEYS{} }
func (m *VrrpInterfaceInfo_KEYS) String() string            { return proto.CompactTextString(m) }
func (*VrrpInterfaceInfo_KEYS) ProtoMessage()               {}
func (*VrrpInterfaceInfo_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *VrrpInterfaceInfo_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type VrrpInterfaceInfo struct {
	// IM Interface
	Interface string `protobuf:"bytes,50,opt,name=interface" json:"interface,omitempty"`
	// Invalid checksum
	InvalidChecksumCount uint32 `protobuf:"varint,51,opt,name=invalid_checksum_count,json=invalidChecksumCount" json:"invalid_checksum_count,omitempty"`
	// Unknown/unsupported version
	InvalidVersionCount uint32 `protobuf:"varint,52,opt,name=invalid_version_count,json=invalidVersionCount" json:"invalid_version_count,omitempty"`
	// Invalid vrID
	InvalidVridCount uint32 `protobuf:"varint,53,opt,name=invalid_vrid_count,json=invalidVridCount" json:"invalid_vrid_count,omitempty"`
	// Bad packet lengths
	InvalidPacketLengthCount uint32 `protobuf:"varint,54,opt,name=invalid_packet_length_count,json=invalidPacketLengthCount" json:"invalid_packet_length_count,omitempty"`
}

func (m *VrrpInterfaceInfo) Reset()                    { *m = VrrpInterfaceInfo{} }
func (m *VrrpInterfaceInfo) String() string            { return proto.CompactTextString(m) }
func (*VrrpInterfaceInfo) ProtoMessage()               {}
func (*VrrpInterfaceInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *VrrpInterfaceInfo) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *VrrpInterfaceInfo) GetInvalidChecksumCount() uint32 {
	if m != nil {
		return m.InvalidChecksumCount
	}
	return 0
}

func (m *VrrpInterfaceInfo) GetInvalidVersionCount() uint32 {
	if m != nil {
		return m.InvalidVersionCount
	}
	return 0
}

func (m *VrrpInterfaceInfo) GetInvalidVridCount() uint32 {
	if m != nil {
		return m.InvalidVridCount
	}
	return 0
}

func (m *VrrpInterfaceInfo) GetInvalidPacketLengthCount() uint32 {
	if m != nil {
		return m.InvalidPacketLengthCount
	}
	return 0
}

func init() {
	proto.RegisterType((*VrrpInterfaceInfo_KEYS)(nil), "cisco_ios_xr_ipv4_vrrp_oper.vrrp.ipv4.interfaces.interface.vrrp_interface_info_KEYS")
	proto.RegisterType((*VrrpInterfaceInfo)(nil), "cisco_ios_xr_ipv4_vrrp_oper.vrrp.ipv4.interfaces.interface.vrrp_interface_info")
}

func init() { proto.RegisterFile("vrrp_interface_info.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xd1, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0x06, 0x70, 0xea, 0x41, 0x58, 0x60, 0x22, 0x99, 0x4a, 0x44, 0x0f, 0x63, 0x20, 0xec, 0x20,
	0x3d, 0x6c, 0xd5, 0x83, 0xe0, 0x41, 0x86, 0x27, 0x45, 0x64, 0x82, 0xe0, 0xe9, 0x11, 0xb3, 0xcc,
	0x85, 0xad, 0x49, 0x78, 0xc9, 0x8a, 0x7f, 0x83, 0x7f, 0xb5, 0xf4, 0x35, 0x6d, 0x2f, 0xbb, 0x85,
	0xef, 0xfb, 0x7e, 0xaf, 0x87, 0xb2, 0xcb, 0x0a, 0xd1, 0x83, 0xb1, 0x51, 0xe3, 0x5a, 0x2a, 0x0d,
	0xc6, 0xae, 0x5d, 0xee, 0xd1, 0x45, 0xc7, 0x1f, 0x94, 0x09, 0xca, 0x81, 0x71, 0x01, 0x7e, 0x11,
	0x8c, 0xaf, 0x0a, 0xa0, 0xb1, 0xf3, 0x1a, 0xf3, 0xfa, 0x95, 0xd7, 0x59, 0xde, 0xd9, 0xd0, 0x3f,
	0x27, 0x4f, 0x4c, 0x1c, 0x38, 0x0c, 0x2f, 0xcf, 0x5f, 0x1f, 0xfc, 0x86, 0x9d, 0xf4, 0xb1, 0x95,
	0xa5, 0x16, 0xd9, 0x38, 0x9b, 0x0e, 0x96, 0xc3, 0x2e, 0x7d, 0x93, 0xa5, 0x9e, 0xfc, 0x1d, 0xb1,
	0xd1, 0x81, 0x1b, 0xfc, 0x9a, 0x0d, 0xba, 0x44, 0xcc, 0x48, 0xf6, 0x01, 0x2f, 0xd8, 0x85, 0xb1,
	0x95, 0xdc, 0x99, 0x15, 0xa8, 0x8d, 0x56, 0xdb, 0xb0, 0x2f, 0x41, 0xb9, 0xbd, 0x8d, 0x62, 0x3e,
	0xce, 0xa6, 0xc3, 0xe5, 0x59, 0x6a, 0x17, 0xa9, 0x5c, 0xd4, 0x1d, 0x9f, 0xb1, 0xf3, 0x56, 0x55,
	0x1a, 0x83, 0x71, 0x36, 0xa1, 0x82, 0xd0, 0x28, 0x95, 0x9f, 0x4d, 0xd7, 0x98, 0x5b, 0xc6, 0x3b,
	0x83, 0xf5, 0xe7, 0x08, 0xdc, 0x11, 0x38, 0x6d, 0x01, 0x9a, 0x55, 0xb3, 0x7e, 0x64, 0x57, 0xed,
	0xda, 0x4b, 0xb5, 0xd5, 0x11, 0x76, 0xda, 0xfe, 0xc4, 0x4d, 0x62, 0xf7, 0xc4, 0x44, 0x9a, 0xbc,
	0xd3, 0xe2, 0x95, 0x06, 0xc4, 0xbf, 0x8f, 0xe9, 0x97, 0xcc, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x35, 0xf5, 0x2c, 0x23, 0xaf, 0x01, 0x00, 0x00,
}

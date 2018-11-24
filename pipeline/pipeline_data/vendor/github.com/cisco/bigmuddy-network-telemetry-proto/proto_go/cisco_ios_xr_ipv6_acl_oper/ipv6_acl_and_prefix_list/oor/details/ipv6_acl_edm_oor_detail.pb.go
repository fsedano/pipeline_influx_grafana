// Code generated by protoc-gen-go.
// source: ipv6_acl_edm_oor_detail.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv6_acl_oper_ipv6_acl_and_prefix_list_oor_details is a generated protocol buffer package.

It is generated from these files:
	ipv6_acl_edm_oor_detail.proto

It has these top-level messages:
	Ipv6AclEdmOorDetail_KEYS
	Ipv6AclEdmOorDetail
*/
package cisco_ios_xr_ipv6_acl_oper_ipv6_acl_and_prefix_list_oor_details

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

// Oor deatil config BAG
type Ipv6AclEdmOorDetail_KEYS struct {
}

func (m *Ipv6AclEdmOorDetail_KEYS) Reset()                    { *m = Ipv6AclEdmOorDetail_KEYS{} }
func (m *Ipv6AclEdmOorDetail_KEYS) String() string            { return proto.CompactTextString(m) }
func (*Ipv6AclEdmOorDetail_KEYS) ProtoMessage()               {}
func (*Ipv6AclEdmOorDetail_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Ipv6AclEdmOorDetail struct {
	// default max configurable acls
	IsDefaultMaximumConfigurableAcLs uint32 `protobuf:"varint,50,opt,name=is_default_maximum_configurable_ac_ls,json=isDefaultMaximumConfigurableAcLs" json:"is_default_maximum_configurable_ac_ls,omitempty"`
	// default max configurable aces
	IsDefaultMaximumConfigurableAcEs uint32 `protobuf:"varint,51,opt,name=is_default_maximum_configurable_ac_es,json=isDefaultMaximumConfigurableAcEs" json:"is_default_maximum_configurable_ac_es,omitempty"`
	// Current configured acls
	IsCurrentConfiguredAcLs uint32 `protobuf:"varint,52,opt,name=is_current_configured_ac_ls,json=isCurrentConfiguredAcLs" json:"is_current_configured_ac_ls,omitempty"`
	// Current configured aces
	IsCurrentConfiguredAces uint32 `protobuf:"varint,53,opt,name=is_current_configured_aces,json=isCurrentConfiguredAces" json:"is_current_configured_aces,omitempty"`
	// Current max configurable acls
	IsCurrentMaximumConfigurableAcls uint32 `protobuf:"varint,54,opt,name=is_current_maximum_configurable_acls,json=isCurrentMaximumConfigurableAcls" json:"is_current_maximum_configurable_acls,omitempty"`
	// Current max configurable aces
	IsCurrentMaximumConfigurableAces uint32 `protobuf:"varint,55,opt,name=is_current_maximum_configurable_aces,json=isCurrentMaximumConfigurableAces" json:"is_current_maximum_configurable_aces,omitempty"`
	// max configurable acls
	IsMaximumConfigurableAcLs uint32 `protobuf:"varint,56,opt,name=is_maximum_configurable_ac_ls,json=isMaximumConfigurableAcLs" json:"is_maximum_configurable_ac_ls,omitempty"`
	// max configurable aces
	IsMaximumConfigurableAcEs uint32 `protobuf:"varint,57,opt,name=is_maximum_configurable_ac_es,json=isMaximumConfigurableAcEs" json:"is_maximum_configurable_ac_es,omitempty"`
}

func (m *Ipv6AclEdmOorDetail) Reset()                    { *m = Ipv6AclEdmOorDetail{} }
func (m *Ipv6AclEdmOorDetail) String() string            { return proto.CompactTextString(m) }
func (*Ipv6AclEdmOorDetail) ProtoMessage()               {}
func (*Ipv6AclEdmOorDetail) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ipv6AclEdmOorDetail) GetIsDefaultMaximumConfigurableAcLs() uint32 {
	if m != nil {
		return m.IsDefaultMaximumConfigurableAcLs
	}
	return 0
}

func (m *Ipv6AclEdmOorDetail) GetIsDefaultMaximumConfigurableAcEs() uint32 {
	if m != nil {
		return m.IsDefaultMaximumConfigurableAcEs
	}
	return 0
}

func (m *Ipv6AclEdmOorDetail) GetIsCurrentConfiguredAcLs() uint32 {
	if m != nil {
		return m.IsCurrentConfiguredAcLs
	}
	return 0
}

func (m *Ipv6AclEdmOorDetail) GetIsCurrentConfiguredAces() uint32 {
	if m != nil {
		return m.IsCurrentConfiguredAces
	}
	return 0
}

func (m *Ipv6AclEdmOorDetail) GetIsCurrentMaximumConfigurableAcls() uint32 {
	if m != nil {
		return m.IsCurrentMaximumConfigurableAcls
	}
	return 0
}

func (m *Ipv6AclEdmOorDetail) GetIsCurrentMaximumConfigurableAces() uint32 {
	if m != nil {
		return m.IsCurrentMaximumConfigurableAces
	}
	return 0
}

func (m *Ipv6AclEdmOorDetail) GetIsMaximumConfigurableAcLs() uint32 {
	if m != nil {
		return m.IsMaximumConfigurableAcLs
	}
	return 0
}

func (m *Ipv6AclEdmOorDetail) GetIsMaximumConfigurableAcEs() uint32 {
	if m != nil {
		return m.IsMaximumConfigurableAcEs
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv6AclEdmOorDetail_KEYS)(nil), "cisco_ios_xr_ipv6_acl_oper.ipv6_acl_and_prefix_list.oor.details.ipv6_acl_edm_oor_detail_KEYS")
	proto.RegisterType((*Ipv6AclEdmOorDetail)(nil), "cisco_ios_xr_ipv6_acl_oper.ipv6_acl_and_prefix_list.oor.details.ipv6_acl_edm_oor_detail")
}

func init() { proto.RegisterFile("ipv6_acl_edm_oor_detail.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x3d, 0x4b, 0xc4, 0x40,
	0x10, 0x86, 0x11, 0xc4, 0x62, 0xc1, 0x26, 0xcd, 0xad, 0x1f, 0x27, 0xc7, 0xa1, 0x60, 0x95, 0xc2,
	0xd3, 0x53, 0x51, 0x50, 0x89, 0xa9, 0xfc, 0x02, 0xad, 0xac, 0x86, 0xbd, 0xcd, 0x44, 0x06, 0x36,
	0x99, 0xb0, 0x93, 0x48, 0x7e, 0xab, 0xbf, 0x46, 0x48, 0xce, 0x70, 0x85, 0x09, 0x29, 0x17, 0x9e,
	0xf7, 0xd9, 0x79, 0x99, 0x51, 0x53, 0x2a, 0xbe, 0x97, 0x60, 0xac, 0x03, 0x4c, 0x32, 0x60, 0xf6,
	0x90, 0x60, 0x69, 0xc8, 0x85, 0x85, 0xe7, 0x92, 0x83, 0x3b, 0x4b, 0x62, 0x19, 0x88, 0x05, 0x6a,
	0x0f, 0x1d, 0xcb, 0x05, 0xfa, 0xb0, 0x7b, 0x99, 0x3c, 0x81, 0xc2, 0x63, 0x4a, 0x35, 0x38, 0x92,
	0x32, 0x64, 0xf6, 0x61, 0x6b, 0x91, 0xf9, 0x91, 0x3a, 0xec, 0xf9, 0x01, 0x9e, 0xe2, 0xcf, 0x8f,
	0xf9, 0xcf, 0xb6, 0x9a, 0xf4, 0x00, 0xc1, 0x9b, 0x3a, 0x21, 0x81, 0x04, 0x53, 0x53, 0xb9, 0x12,
	0x32, 0x53, 0x53, 0x56, 0x65, 0x60, 0x39, 0x4f, 0xe9, 0xab, 0xf2, 0x66, 0xe5, 0x10, 0x8c, 0x05,
	0x27, 0xfa, 0x6c, 0xb6, 0x75, 0xba, 0xfb, 0x3e, 0x23, 0x79, 0x6c, 0xd9, 0x97, 0x16, 0x8d, 0x36,
	0xc8, 0x07, 0xfb, 0x2c, 0x23, 0x85, 0x28, 0x7a, 0x31, 0x46, 0x18, 0x4b, 0x70, 0xab, 0x0e, 0x48,
	0xc0, 0x56, 0xde, 0x63, 0x5e, 0x76, 0x22, 0x4c, 0xd6, 0x73, 0x9d, 0x37, 0x9a, 0x09, 0x49, 0xd4,
	0x12, 0x51, 0x07, 0x34, 0xe3, 0xdc, 0xa8, 0xfd, 0xbe, 0x34, 0x8a, 0xbe, 0x18, 0x08, 0xa3, 0x04,
	0xaf, 0xea, 0x78, 0x23, 0xdc, 0xd3, 0xc5, 0x89, 0x5e, 0xfe, 0x55, 0x59, 0x6b, 0xfe, 0xad, 0xe2,
	0x46, 0xfa, 0x50, 0xf4, 0xe5, 0x18, 0x1f, 0x4a, 0x70, 0xaf, 0xa6, 0x24, 0x43, 0x4b, 0xbb, 0x6a,
	0x44, 0x7b, 0x24, 0x7d, 0xdb, 0x1a, 0x36, 0xa0, 0xe8, 0xeb, 0x41, 0x43, 0x2c, 0xab, 0x9d, 0xe6,
	0x88, 0x17, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x06, 0x3d, 0xa5, 0x90, 0xe5, 0x02, 0x00, 0x00,
}

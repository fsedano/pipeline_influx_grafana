// Code generated by protoc-gen-go.
// source: ipv6_rib_edm_advert.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ip_rib_ipv6_oper_ipv6_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert is a generated protocol buffer package.

It is generated from these files:
	ipv6_rib_edm_advert.proto

It has these top-level messages:
	Ipv6RibEdmAdvert_KEYS
	Ipv6RibEdmAdvert
	Ipv6RibEdmAdvertItem
*/
package cisco_ios_xr_ip_rib_ipv6_oper_ipv6_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert

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

// Route advertisement information
type Ipv6RibEdmAdvert_KEYS struct {
	VrfName        string `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	AfName         string `protobuf:"bytes,2,opt,name=af_name,json=afName" json:"af_name,omitempty"`
	SafName        string `protobuf:"bytes,3,opt,name=saf_name,json=safName" json:"saf_name,omitempty"`
	RouteTableName string `protobuf:"bytes,4,opt,name=route_table_name,json=routeTableName" json:"route_table_name,omitempty"`
	Address        string `protobuf:"bytes,5,opt,name=address" json:"address,omitempty"`
	PrefixLength   uint32 `protobuf:"varint,6,opt,name=prefix_length,json=prefixLength" json:"prefix_length,omitempty"`
}

func (m *Ipv6RibEdmAdvert_KEYS) Reset()                    { *m = Ipv6RibEdmAdvert_KEYS{} }
func (m *Ipv6RibEdmAdvert_KEYS) String() string            { return proto.CompactTextString(m) }
func (*Ipv6RibEdmAdvert_KEYS) ProtoMessage()               {}
func (*Ipv6RibEdmAdvert_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ipv6RibEdmAdvert_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv6RibEdmAdvert_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *Ipv6RibEdmAdvert_KEYS) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *Ipv6RibEdmAdvert_KEYS) GetRouteTableName() string {
	if m != nil {
		return m.RouteTableName
	}
	return ""
}

func (m *Ipv6RibEdmAdvert_KEYS) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Ipv6RibEdmAdvert_KEYS) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

type Ipv6RibEdmAdvert struct {
	// Next advertising proto
	Ipv6RibEdmAdvert []*Ipv6RibEdmAdvertItem `protobuf:"bytes,50,rep,name=ipv6_rib_edm_advert,json=ipv6RibEdmAdvert" json:"ipv6_rib_edm_advert,omitempty"`
}

func (m *Ipv6RibEdmAdvert) Reset()                    { *m = Ipv6RibEdmAdvert{} }
func (m *Ipv6RibEdmAdvert) String() string            { return proto.CompactTextString(m) }
func (*Ipv6RibEdmAdvert) ProtoMessage()               {}
func (*Ipv6RibEdmAdvert) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ipv6RibEdmAdvert) GetIpv6RibEdmAdvert() []*Ipv6RibEdmAdvertItem {
	if m != nil {
		return m.Ipv6RibEdmAdvert
	}
	return nil
}

type Ipv6RibEdmAdvertItem struct {
	// Protocol advertising the route
	ProtocolId uint32 `protobuf:"varint,1,opt,name=protocol_id,json=protocolId" json:"protocol_id,omitempty"`
	//   Client advertising the route
	ClientId uint32 `protobuf:"varint,2,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	// Number of extended communities
	NumberOfExtendedCommunities uint32 `protobuf:"varint,3,opt,name=number_of_extended_communities,json=numberOfExtendedCommunities" json:"number_of_extended_communities,omitempty"`
	// Extended communities
	ExtendedCommunities []byte `protobuf:"bytes,4,opt,name=extended_communities,json=extendedCommunities,proto3" json:"extended_communities,omitempty"`
	// OSPF area-id flags
	ProtocolOpaqueFlags uint32 `protobuf:"varint,5,opt,name=protocol_opaque_flags,json=protocolOpaqueFlags" json:"protocol_opaque_flags,omitempty"`
	// OSPF area-id
	ProtocolOpaque uint32 `protobuf:"varint,6,opt,name=protocol_opaque,json=protocolOpaque" json:"protocol_opaque,omitempty"`
	// Protocol code
	Code int32 `protobuf:"zigzag32,7,opt,name=code" json:"code,omitempty"`
	// Instance name
	InstanceName []byte `protobuf:"bytes,8,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
}

func (m *Ipv6RibEdmAdvertItem) Reset()                    { *m = Ipv6RibEdmAdvertItem{} }
func (m *Ipv6RibEdmAdvertItem) String() string            { return proto.CompactTextString(m) }
func (*Ipv6RibEdmAdvertItem) ProtoMessage()               {}
func (*Ipv6RibEdmAdvertItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Ipv6RibEdmAdvertItem) GetProtocolId() uint32 {
	if m != nil {
		return m.ProtocolId
	}
	return 0
}

func (m *Ipv6RibEdmAdvertItem) GetClientId() uint32 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *Ipv6RibEdmAdvertItem) GetNumberOfExtendedCommunities() uint32 {
	if m != nil {
		return m.NumberOfExtendedCommunities
	}
	return 0
}

func (m *Ipv6RibEdmAdvertItem) GetExtendedCommunities() []byte {
	if m != nil {
		return m.ExtendedCommunities
	}
	return nil
}

func (m *Ipv6RibEdmAdvertItem) GetProtocolOpaqueFlags() uint32 {
	if m != nil {
		return m.ProtocolOpaqueFlags
	}
	return 0
}

func (m *Ipv6RibEdmAdvertItem) GetProtocolOpaque() uint32 {
	if m != nil {
		return m.ProtocolOpaque
	}
	return 0
}

func (m *Ipv6RibEdmAdvertItem) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Ipv6RibEdmAdvertItem) GetInstanceName() []byte {
	if m != nil {
		return m.InstanceName
	}
	return nil
}

func init() {
	proto.RegisterType((*Ipv6RibEdmAdvert_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.adverts.advert.ipv6_rib_edm_advert_KEYS")
	proto.RegisterType((*Ipv6RibEdmAdvert)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.adverts.advert.ipv6_rib_edm_advert")
	proto.RegisterType((*Ipv6RibEdmAdvertItem)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.adverts.advert.ipv6_rib_edm_advert_item")
}

func init() { proto.RegisterFile("ipv6_rib_edm_advert.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 458 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x92, 0xd1, 0x8a, 0xd4, 0x30,
	0x18, 0x85, 0xe9, 0xec, 0x3a, 0x33, 0x9b, 0x9d, 0xae, 0x6b, 0x46, 0xb1, 0xc3, 0x82, 0x0e, 0xeb,
	0x85, 0xbd, 0x0a, 0x38, 0x82, 0xf7, 0xb2, 0x8c, 0xb0, 0x28, 0x2e, 0x44, 0x6f, 0xbc, 0x0a, 0x69,
	0xf3, 0x67, 0x0d, 0xb4, 0x4d, 0x4d, 0x32, 0x65, 0x7c, 0x00, 0x9f, 0xc3, 0x47, 0x12, 0x7c, 0x08,
	0x7d, 0x0d, 0x49, 0xd2, 0xae, 0xb8, 0xd4, 0xfb, 0xbd, 0x69, 0xda, 0xf3, 0x9d, 0x13, 0xfe, 0x72,
	0x7e, 0xb4, 0x52, 0x6d, 0xf7, 0x8a, 0x19, 0x55, 0x30, 0x10, 0x35, 0xe3, 0xa2, 0x03, 0xe3, 0x48,
	0x6b, 0xb4, 0xd3, 0xf8, 0x5b, 0x52, 0x2a, 0x5b, 0x6a, 0xa6, 0xb4, 0x65, 0x7b, 0xc3, 0x54, 0x1b,
	0x6c, 0xc1, 0xaf, 0x5b, 0x30, 0xe4, 0x26, 0x69, 0x9d, 0x28, 0xbe, 0x92, 0xce, 0x48, 0xeb, 0x1f,
	0x84, 0x4b, 0x4b, 0xb8, 0x24, 0xd6, 0x9f, 0x96, 0x4b, 0xd2, 0x07, 0x8d, 0xde, 0x39, 0x60, 0x8e,
	0x17, 0x15, 0xb0, 0x86, 0xd7, 0x60, 0xff, 0x07, 0x48, 0x9c, 0xc2, 0xf6, 0xe7, 0xf9, 0xcf, 0x04,
	0x65, 0x23, 0x53, 0xb2, 0xb7, 0xdb, 0x4f, 0x1f, 0xf0, 0x0a, 0xcd, 0x3b, 0x23, 0x43, 0x2e, 0x4b,
	0xd6, 0x49, 0x7e, 0x44, 0x67, 0x9d, 0x91, 0xef, 0x79, 0x0d, 0xf8, 0x31, 0x9a, 0xf1, 0x9e, 0x4c,
	0x02, 0x99, 0xf2, 0x08, 0x56, 0x68, 0x6e, 0x07, 0x72, 0x10, 0x33, 0xb6, 0x47, 0x39, 0x3a, 0xbd,
	0x3d, 0x4e, 0x76, 0x18, 0x2c, 0x27, 0x41, 0xff, 0xe8, 0xe5, 0xe0, 0xcc, 0xd0, 0x8c, 0x0b, 0x61,
	0xc0, 0xda, 0xec, 0x5e, 0xbc, 0xa3, 0xff, 0xc4, 0xcf, 0x50, 0xda, 0x1a, 0x90, 0x6a, 0xcf, 0x2a,
	0x68, 0xae, 0xdd, 0xe7, 0x6c, 0xba, 0x4e, 0xf2, 0x94, 0x2e, 0xa2, 0xf8, 0x2e, 0x68, 0xe7, 0xbf,
	0x13, 0xb4, 0x1c, 0xf9, 0x29, 0xfc, 0x63, 0x5c, 0xcf, 0x36, 0xeb, 0x83, 0xfc, 0x78, 0xf3, 0x3d,
	0x21, 0x77, 0xa2, 0x13, 0x32, 0xd6, 0x87, 0x72, 0x50, 0xd3, 0x53, 0x4f, 0xa8, 0x2a, 0xb6, 0xa2,
	0x7e, 0x1d, 0xeb, 0xfb, 0x35, 0x19, 0xaf, 0xcf, 0xdb, 0xf1, 0x53, 0x74, 0x1c, 0x96, 0xad, 0xd4,
	0x15, 0x53, 0x22, 0x34, 0x98, 0x52, 0x34, 0x48, 0x97, 0x02, 0x9f, 0xa1, 0xa3, 0xb2, 0x52, 0xd0,
	0x38, 0x8f, 0x27, 0x01, 0xcf, 0xa3, 0x70, 0x29, 0xf0, 0x05, 0x7a, 0xd2, 0xec, 0xea, 0x02, 0x0c,
	0xd3, 0x92, 0xc1, 0xde, 0x41, 0x23, 0x40, 0xb0, 0x52, 0xd7, 0xf5, 0xae, 0x51, 0x4e, 0x81, 0x0d,
	0xf5, 0xa6, 0xf4, 0x2c, 0xba, 0xae, 0xe4, 0xb6, 0xf7, 0x5c, 0xfc, 0xb5, 0xe0, 0x17, 0xe8, 0xe1,
	0x68, 0xd4, 0xd7, 0xbe, 0xa0, 0x4b, 0x18, 0x89, 0x6c, 0xd0, 0xa3, 0x9b, 0xa9, 0x75, 0xcb, 0xbf,
	0xec, 0x80, 0xc9, 0x8a, 0x5f, 0xc7, 0x4d, 0x48, 0xe9, 0x72, 0x80, 0x57, 0x81, 0xbd, 0xf1, 0x08,
	0x3f, 0x47, 0xf7, 0x6f, 0x65, 0xfa, 0xbd, 0x38, 0xf9, 0xd7, 0x8d, 0x31, 0x3a, 0x2c, 0xb5, 0x80,
	0x6c, 0xb6, 0x4e, 0xf2, 0x07, 0x34, 0xbc, 0xfb, 0x95, 0x52, 0x8d, 0x75, 0xbc, 0x29, 0xfb, 0x9d,
	0x9c, 0x87, 0xe1, 0x16, 0x83, 0xe8, 0x37, 0xb2, 0x98, 0x86, 0x8b, 0x5e, 0xfe, 0x09, 0x00, 0x00,
	0xff, 0xff, 0xc6, 0x8e, 0xb7, 0xdd, 0xd3, 0x03, 0x00, 0x00,
}
